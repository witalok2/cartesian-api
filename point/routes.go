package point

import (
	"github.com/cartesian-api/utils/handler"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	auth := e.Group("api/v1/")

	auth.GET("points", FindByFilterHandler, handler.MiddlewareBindAndValidate(&ParamCoordinate{}))
}
