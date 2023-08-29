package handlers

import (
	"kibogo/app/libraries"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {

	resp := &libraries.Response{
		Status:  1,
		Message: "Hello World!",
	}

	return c.JSON(http.StatusOK, resp)
}
