package router

import (
	"go-rest-api/presentation/echoserver/inject"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
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
	e.POST("/signup", injector.SignUp())
	e.POST("/login", injector.Login())
	e.POST("/logout", injector.Logout())
	// ①の件：この役目は何？
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", injector.GetAllTasks())
	t.GET("/:taskId", injector.GetTaskById())
	t.POST("", injector.CreateTask())
	t.PUT("/:taskId", injector.UpdateTask())
	t.DELETE("/:taskId", injector.DeleteTask())

}
