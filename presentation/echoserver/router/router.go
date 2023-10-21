package router

import (
	"go-rest-api/presentation/echoserver/inject"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()
	injector := inject.NewInjector(db)
	route(e, &injector)
	return e
}

func route(e *echo.Echo, injector *inject.Injector) {
	e.POST("/signup", injector.UserControllerSignUp())
	e.POST("/login", injector.UserControllerLogin())
	e.POST("/logout", injector.UserControllerLogout())
}
