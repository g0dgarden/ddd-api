package routers

import (
	"github.com/labstack/echo"

	"github.com/g0dgarden/ddd-api"
	"github.com/g0dgarden/ddd-api/resources"
)

func Init() *echo.Echo {
	user := resources.NewUser()
	e := echo.New()
	v1 := e.Group("/v1")
	{
		v1.GET("/users", user.GetUsers(api.DBConn))
		v1.GET("/users/:id", user.GetUser(api.DBConn))
	}
	return e
}
