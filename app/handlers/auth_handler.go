package handlers

import (
	"kibogo/app/libraries"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	// UserRequest ...
	UserAuthRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// ResponseData ...
	UserAuthResponseData struct {
		Token    string `json:"token"`
		ExpireIn int    `json:"expires_in"`
	}
)

func UserAuth(c echo.Context) (err error) {
	var response *libraries.Response
	u := new(UserAuthRequest)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}

	auth := libraries.UserCredential{
		Username: u.Username,
		Password: u.Password,
		Ttl:      3600,
	}
	result, ok := auth.CreateToken()

	if ok {
		response = &libraries.Response{
			Status:  1,
			Message: "Success",
			Data: UserAuthResponseData{
				Token:    result,
				ExpireIn: 3600,
			},
		}
	} else {
		response = &libraries.Response{
			Status:  0,
			Message: result,
		}
	}

	return c.JSON(http.StatusOK, response)
}
