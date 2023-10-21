package inject

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func (i *Injector) UserControllerSignUp() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).SignUp
}

func (i *Injector) UserControllerLogin() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).Login
}

func (i *Injector) UserControllerLogout() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).Logout
}
