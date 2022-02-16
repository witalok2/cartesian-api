package point

import (
	"github.com/cartesian-api/utils/handler"
	"github.com/labstack/echo/v4"
)

func FindByFilterHandler(c echo.Context) (err error) {
	params := *c.Get(handler.PARAMETERS).(*ParamCoordinate)

	point, err := FindByCoordinate(params)
	if err != nil {
		return err
	}

	return c.JSON(200, point)
}

func CreateMultipleCoordinateHnadler(c echo.Context) (err error) {
	coordinate := *c.Get(handler.PARAMETERS).(*[]Coordinate)

	err = CreateMultipleCoordinate(coordinate)
	if err != nil {
		return err
	}

	return c.NoContent(201)
}
