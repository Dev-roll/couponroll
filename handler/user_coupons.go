package handler

import (
	"net/http"

	"github.com/Dev-roll/couponroll/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	PostUserCouponRequest struct {
		UserID   uuid.UUID `json:"user_id"`
		CouponID uuid.UUID `json:"coupon_id"`
	}
)

type (
	PatchUserCouponRequest struct {
		ID     uuid.UUID `json:"id"`
		IsUsed bool      `json:"is_used"`
	}
)

func PostUserCoupon(c echo.Context) error {
	var req PostUserCouponRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userCoupon := model.UserCoupon{
		UserID:   req.UserID,
		CouponID: req.CouponID,
		IsUsed:   false,
	}

	model.CreateUserCoupon(userCoupon)

	return c.NoContent(http.StatusCreated)
}

func GetUserCoupon(c echo.Context) error {
	userCouponID := c.Param("user_coupon_id")
	userCoupon := model.GetUserCouponFromID(model.UserCoupon{ID: uuid.MustParse(userCouponID)})
	return c.JSON(http.StatusOK, userCoupon)
}

func PatchUserCoupon(c echo.Context) error {
	userCouponID := c.Param("user_coupon_id")

	var req PatchUserCouponRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userCoupon := model.UserCoupon{
		IsUsed: req.IsUsed,
	}

	model.UpdateUserCoupon(uuid.MustParse(userCouponID), userCoupon)

	return c.NoContent(http.StatusNoContent)
}

func DeleteUserCoupon(c echo.Context) error {
	userCouponID := c.Param("user_coupon_id")
	model.DeleteUserCoupon(model.UserCoupon{ID: uuid.MustParse(userCouponID)})
	return c.NoContent(http.StatusNoContent)
}

func GetMyAllCoupons(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	userCoupons := model.GetUserCouponsFromUserID(userID)
	return c.JSON(http.StatusOK, userCoupons)
}

func GetMyCouponsByStore(c echo.Context) error {
	userID := c.Get("userID").(uuid.UUID)
	storeID := c.Param("store_id")
	userCoupons := model.GetUserCouponsFromUserIDByStore(userID, uuid.MustParse(storeID))
	return c.JSON(http.StatusOK, userCoupons)
}

func DeleteMyCoupon(c echo.Context) error {
	userCouponID := c.Param("user_coupon_id")

	userCoupon := model.GetUserCouponFromID(model.UserCoupon{ID: uuid.MustParse(userCouponID)})
	if userCoupon.UserID != c.Get("userID").(uuid.UUID) {
		return c.String(http.StatusForbidden, "forbidden")
	}

	model.DeleteUserCoupon(model.UserCoupon{ID: uuid.MustParse(userCouponID)})

	return c.NoContent(http.StatusNoContent)
}
