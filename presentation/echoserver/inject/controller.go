package inject

import (
	"go-rest-api/presentation/echoserver/controller"

	"github.com/labstack/echo/v4"
)

func (i *Injector) SignUp() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).SignUp
}

func (i *Injector) Login() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).Login
}

func (i *Injector) Logout() echo.HandlerFunc {
	return controller.NewUserController(i.UserUsecase()).Logout
}

func (i *Injector) GetAllTasks() echo.HandlerFunc {
	return controller.NewTaskController(i.TaskUsecase()).GetAllTasks
}

func (i *Injector) GetTaskById() echo.HandlerFunc {
	return controller.NewTaskController(i.TaskUsecase()).GetTaskById
}

func (i *Injector) CreateTask() echo.HandlerFunc {
	return controller.NewTaskController(i.TaskUsecase()).CreateTask
}

func (i *Injector) UpdateTask() echo.HandlerFunc {
	return controller.NewTaskController(i.TaskUsecase()).UpdateTask
}

func (i *Injector) DeleteTask() echo.HandlerFunc {
	return controller.NewTaskController(i.TaskUsecase()).DeleteTask
}
