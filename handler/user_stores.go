package handler

import (
	"net/http"

	"github.com/Dev-roll/couponroll/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	PostUserStoreRequest struct {
		ID uuid.UUID `json:"id"`
	}
)

func GetUsers(c echo.Context) error {
	storeID := c.Param("store_id")
	users := model.GetUsersFromStoreID(uuid.MustParse(storeID))
	return c.JSON(http.StatusOK, users)
}

func GetMyStores(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	stores := model.GetStoresFromUserID(userID)
	return c.JSON(http.StatusOK, stores)
}

func PutMyStore(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)

	var req PostUserStoreRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userStore := model.UserStores{
		UserID:  userID,
		StoreID: req.ID,
	}

	model.CreateUserStore(userStore)

	return c.NoContent(http.StatusCreated)
}

func DeleteMyStore(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	storeID := c.Param("store_id")

	model.DeleteUserStore(userID, model.UserStores{StoreID: uuid.MustParse(storeID)})

	return c.NoContent(http.StatusNoContent)
}
