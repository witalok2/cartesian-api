package point

import (
	"math"
	"sort"

	"github.com/cartesian-api/config"
	"github.com/cartesian-api/utils/database"
)

func CreateMultipleCoordinate(coordinate []Coordinate) error {
	db := database.MustGetByFile(config.POSTGRES_FILE)

	err := createMultipleCoordinateDB(coordinate, db)
	if err != nil {
		return err
	}

	return nil
}

func FindByCoordinate(params ParamCoordinate) (coordinateDistance []CoordinateDistance, err error) {
	db := database.MustGetByFile(config.POSTGRES_FILE)

	coordinates, err := findByCoordinateDB(db)
	if err != nil {
		return
	}

	for _, coordinate := range coordinates {
		init := []int{params.X, params.Y}
		final := []int{coordinate.X, coordinate.Y}

		distance := manhattanDistance(init, final)
		if distance <= params.Distance {
			operation := CoordinateDistance{
				From: Coordinate{
					X: params.X,
					Y: params.Y,
				},
				To: Coordinate{
					X: coordinate.X,
					Y: coordinate.Y,
				},
				Distance: distance,
			}
			coordinateDistance = append(coordinateDistance, operation)
		}
	}

	sort.Slice(coordinateDistance, func(i, j int) bool {
		return coordinateDistance[i].Distance < coordinateDistance[j].Distance
	})

	return coordinateDistance, nil
}

func manhattanDistance(x, y []int) int {
	var distance int

	if len(y) != len(x) {
		return 0
	}

	distance = 0

	for i := 0; i < len(x); i += 1 {
		distance += int(math.Abs(float64(y[i] - x[i])))
	}

	return distance
}
