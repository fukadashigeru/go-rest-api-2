package controller

import (
	icontoller "go-rest-api/domain/controller"
	"go-rest-api/domain/model"
	iusecase "go-rest-api/domain/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type taskController struct {
	tu iusecase.ITaskUsecase
}

func NewTaskController(tu iusecase.ITaskUsecase) icontoller.ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := ctx.Param("taskId")
	taskId, _ := strconv.Atoi(id) //何してる？
	// uinitとかfloat64とか何してる？
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := model.Task{} //何してる？
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := ctx.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := ctx.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
