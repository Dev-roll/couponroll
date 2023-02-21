package handler

import (
	"net/http"
	"time"

	"github.com/Dev-roll/couponroll/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	PostCouponRequest struct {
		StoreID     string    `json:"store_id,omitempty" form:"store_id"`
		Name        string    `json:"name,omitempty" form:"name"`
		Description string    `json:"description,omitempty" form:"description"`
		ImgUrl      string    `json:"img_url,omitempty" form:"img_url"`
		Discount    int       `json:"discount,omitempty" form:"discount"`
		CreatorID   string    `json:"creator_id,omitempty" form:"creator_id"`
		ExpiresAt   time.Time `json:"expires_at,omitempty" form:"expires_at"`
	}
)

func PostCoupon(c echo.Context) error {
	var req PostCouponRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	coupon := model.Coupon{
		StoreID:     uuid.MustParse(req.StoreID),
		Name:        req.Name,
		Description: req.Description,
		ImgUrl:      req.ImgUrl,
		Discount:    req.Discount,
		CreatorID:   uuid.MustParse(req.CreatorID),
		ExpiresAt:   req.ExpiresAt,
	}

	model.CreateCoupon(coupon)

	return c.NoContent(http.StatusCreated)
}

func GetCoupon(c echo.Context) error {
	couponID := c.Param("coupon_id")
	coupon := model.GetCouponFromID(model.Coupon{ID: uuid.MustParse(couponID)})
	return c.JSON(http.StatusOK, coupon)
}

func GetCoupons(c echo.Context) error {
	coupons := model.GetCoupons()
	return c.JSON(http.StatusOK, coupons)
}

func PatchCoupon(c echo.Context) error {
	couponID := c.Param("coupon_id")

	var req PostCouponRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	coupon := model.Coupon{
		StoreID:     uuid.MustParse(req.StoreID),
		Name:        req.Name,
		Description: req.Description,
		ImgUrl:      req.ImgUrl,
		Discount:    req.Discount,
		CreatorID:   uuid.MustParse(req.CreatorID),
		ExpiresAt:   req.ExpiresAt,
	}

	model.UpdateCoupon(uuid.MustParse(couponID), coupon)

	return c.NoContent(http.StatusOK)
}

func DeleteCoupon(c echo.Context) error {
	couponID := c.Param("coupon_id")
	model.DeleteCoupon(model.Coupon{ID: uuid.MustParse(couponID)})
	return c.NoContent(http.StatusOK)
}
