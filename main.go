package main

import (
	"kibogo/app/libraries"
	"kibogo/app/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // initiate echo

	// get debug status
	debug, err := strconv.ParseBool(libraries.Env("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}
	e.Debug = debug

	e.Static("assets", "public/assets") // static assets

	libraries.NewDB()                                                    // load DB
	e.Validator = &libraries.CustomValidator{Validator: validator.New()} // load validator

	e.Any("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome, Everyone!")
	})

	routes.Api(e.Group("/api"))

	e.Logger.Fatal(e.Start(":1323"))
}
