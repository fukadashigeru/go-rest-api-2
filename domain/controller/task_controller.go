package icontoller

import (
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(ctx echo.Context) error
	GetTaskById(ctx echo.Context) error
	CreateTask(ctx echo.Context) error
	UpdateTask(ctx echo.Context) error
	DeleteTask(ctx echo.Context) error
}
