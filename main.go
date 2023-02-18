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

	g := e.Group("")
	g.Use(handler.CheckLogin)

	// users
	g.GET("/users/me", handler.GetMe)
	g.PATCH("/users/me", handler.PatchMe)
	// Does not work correctly
	g.PUT("/users/me/password", handler.PutPassword)

	e.Logger.Fatal(e.Start(":1323"))
}
