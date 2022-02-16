package point

type Coordinate struct {
	X int `json:"x" db:"x"`
	Y int `json:"y" db:"y"`
}

type CoordinateDistance struct {
	From     Coordinate `json:"from"`
	To       Coordinate `json:"to"`
	Distance int        `json:"distance"`
}

type ParamCoordinate struct {
	X        int `query:"x" db:"x"`
	Y        int `query:"y" db:"y"`
	Distance int `query:"distance"`
}