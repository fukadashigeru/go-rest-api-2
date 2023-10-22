package controller

import (
	icontoller "go-rest-api/domain/controller"
	"go-rest-api/domain/model"
	iusecase "go-rest-api/domain/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type userController struct {
	uu iusecase.IUserUsecase
}

func NewUserController(uu iusecase.IUserUsecase) icontoller.IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)
}

func (uc *userController) Logout(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)
}
