package jsonresp

import "github.com/labstack/echo"

type response struct {
	Response string `json:"response"`
}

// New returns an Echo error message in proper JSON format
func New(c echo.Context, httpStatus int, message string) error {
	response := &response{
		Response: message,
	}

	return c.JSON(httpStatus, *response)
}
