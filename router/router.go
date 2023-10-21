package router

// import (
// 	"go-rest-api/inject"

// 	"github.com/labstack/echo/v4"
// )

// func NewRouter(i *inject.Injector) *echo.Echo {
// 	e := echo.New()
// 	e.POST("/signup", uc.SignUp)
// 	e.POST("/login", uc.Login)
// 	e.POST("/logout", uc.Logout)
// 	return e
// }
import (
	"go-rest-api/inject"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()
	i := inject.NewInjector(db)
	route(e, &i)
	return e
}

func route(e *echo.Echo, i *inject.Injector) {
	e.POST("/signup", i.UserControllerSinghUp())
	// e.POST("/login", i.Login)
	// e.POST("/logout", i.Logout)
}
