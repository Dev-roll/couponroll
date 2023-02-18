package main

import (
	"github.com/Dev-roll/couponroll/handler"
	"github.com/Dev-roll/couponroll/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(session.Middleware((model.SessionStore)))

	// auth
	e.POST("/auth/register", handler.Register)
	e.POST("/auth/login", handler.Login)

	// stores
	e.POST("/stores", handler.PostStore)
	e.GET("/stores/:store_id", handler.GetStore)
	e.PATCH("/stores/:store_id", handler.PatchStore)
	e.DELETE("/stores/:store_id", handler.DeleteStore)

	// coupons
	e.POST("/coupons", handler.PostCoupon)
	e.GET("/coupons/:coupon_id", handler.GetCoupon)
	e.PATCH("/coupons/:coupon_id", handler.PatchCoupon)
	e.DELETE("/coupons/:coupon_id", handler.DeleteCoupon)

	// user_coupons
	e.POST("/user_coupons", handler.PostUserCoupons)
	e.GET("/user_coupons", handler.GetUserCoupon)
	e.PATCH("/user_coupons/:user_coupon_id", handler.PatchUserCoupon)
	e.DELETE("/user_coupons", handler.DeleteUserCoupons)
	e.DELETE("/user_coupons/:user_coupon_id", handler.DeleteUserCoupon)

	g := e.Group("")
	g.Use(handler.CheckLogin)

	// users
	g.GET("/users/me", handler.GetMe)
	g.PATCH("/users/me", handler.PatchMe)

	g.GET("/users/me/stores", handler.GetMyStores)
	g.POST("/users/me/stores", handler.PostMyStore)
	g.DELETE("/users/me/stores/:store_id", handler.DeleteMyStore)

	g.GET("/users/me/coupons", handler.GetMyAllCoupons)
	g.GET("/users/me/coupons/:store_id", handler.GetMyCoupons)
	g.DELETE("/users/me/coupons/:coupon_id", handler.DeleteMyCoupon)

	// TODO: Does not work correctly
	g.PUT("/users/me/password", handler.PutPassword)

	e.Logger.Fatal(e.Start(":1323"))
}
