package inject

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func (i *Injector) UserControllerSinghUp() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).SignUp
}
