package point

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
)

func FindByCoordinate(params ParamCoordinate) (coordinateDistance []CoordinateDistance, err error) {
	coordinates, err := LoadFilePoint()
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

func LoadFilePoint() (coordinate []Coordinate, err error) {
	jsonFile, err := os.Open("point.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValueJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(byteValueJSON, &coordinate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return coordinate, nil
}
