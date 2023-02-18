package handler

import (
	"fmt"
	"net/http"

	"github.com/Dev-roll/couponroll/model"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	PostUserRequest struct {
		Name     string `json:"name,omitempty" form:"name" validate:"required,min=3,max=32"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8,max=32"`
		Role     string `json:"role,omitempty" form:"role"`
	}

	PatchUserRequest struct {
		Name  string `json:"name,omitempty" form:"name" validate:"required,min=3,max=32"`
		Email string `json:"email" form:"email" validate:"required,email"`
		Role  string `json:"role,omitempty" form:"role"`
	}

	PutPasswordRequest struct {
		OldPassword string `json:"old_password" form:"old_password" validate:"required,min=8,max=32"`
		NewPassword string `json:"new_password" form:"new_password" validate:"required,min=8,max=32"`
	}
)

func Register(c echo.Context) error {
	var req PostUserRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// TODO: validate
	if req.Email == "" || req.Password == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if model.GetUsersCountFromName(model.User{Name: req.Name}) > 0 {
		return c.String(http.StatusConflict, "user name already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	model.CreateUser(model.User{Name: req.Name, Email: req.Email, Password: string(hashedPass), Role: "user"})

	return c.NoContent(http.StatusCreated)
}

func Login(c echo.Context) error {
	var req PostUserRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := model.GetUserFromEmail(model.User{Email: req.Email})

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.NoContent(http.StatusForbidden)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}

	sess.Values["userName"] = user.Name
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusNoContent)
}

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}

		if sess.Values["userName"] == nil {
			return c.String(http.StatusForbidden, "please login")
		}

		userName := sess.Values["userName"].(string)
		user := model.GetUserFromName(model.User{Name: userName})

		c.Set("userID", user.ID)

		return next(c)
	}
}

func GetMe(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)

	user := model.GetUserFromID(model.User{ID: userID})

	res := model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

func PatchMe(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)

	var req PatchUserRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if req.Name == "" || req.Email == "" || req.Role == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if model.GetUsersCountFromName(model.User{Name: req.Name}) > 0 {
		return c.String(http.StatusConflict, "user name already exists")
	}

	user := model.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	model.UpdateUser(userID, user)

	return c.NoContent(http.StatusNoContent)
}

func PutPassword(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)

	var req PutPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := model.GetUserFromID(model.User{ID: userID})

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.NoContent(http.StatusForbidden)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	model.UpdatePassword(userID, model.User{Password: string(hashedPass)})
	Login(c)

	return c.NoContent(http.StatusNoContent)
}
