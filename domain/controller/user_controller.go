package icontoller

import "github.com/labstack/echo/v4"

type IUserController interface {
	SignUp(ctx echo.Context) error
	Login(ctx echo.Context) error
	Logout(ctx echo.Context) error
}
