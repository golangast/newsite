package routes

import (
	"github.com/golangast/newsite/handler/home"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)

}
