package routes

import (
	"kibogo/app/handlers"
	"kibogo/app/libraries"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Api(route *echo.Group) {

	// auth api route
	route.POST("/auth/user", handlers.UserAuth)

	// application api route
	appsRoute := route.Group("/apps/v1")
	appsRoute.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		Validator: func(key string, c echo.Context) (bool, error) {
			validate := libraries.CheckAuthKey(key)
			return validate, nil
		},
	}))
	appsRoute.GET("/test/handler", handlers.HelloWorld)
