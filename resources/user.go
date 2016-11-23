package resources

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"

	"github.com/g0dgarden/ddd-api/users"
	infra "github.com/g0dgarden/ddd-api/infrastructures"
)

func NewUser() *userResource {
	return &userResource{
		user: users.NewRepository(),
	}
}

type userResource struct {
	user users.Repository
}

func (u *userResource) GetUsers(exec infra.Executor) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		users, err := u.user.GetUsers(exec)
		if err != nil {
			return echo.NewHTTPError(fasthttp.StatusNotFound, "User does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, users)
	}
}

func (u *userResource) GetUser(exec infra.Executor) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
		user, err := u.user.GetUser(exec, id)
		if err != nil {
			return echo.NewHTTPError(fasthttp.StatusNotFound, "User does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, user)
	}
}
