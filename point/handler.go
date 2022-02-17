package point

import (
	"github.com/cartesian-api/utils/handler"
	"github.com/labstack/echo/v4"
)

func FindByCoordinateHandler(c echo.Context) (err error) {
	params := *c.Get(handler.PARAMETERS).(*ParamCoordinate)

	point, err := FindByCoordinate(params)
	if err != nil {
		return err
	}

	return c.JSON(200, point)
}
