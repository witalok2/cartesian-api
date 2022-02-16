package point

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

func createMultipleCoordinateDB(coordinates []Coordinate, db *sqlx.DB) error {
	query := `INSERT INTO coordinate (x, y) VALUES `

	args := []interface{}{}
	for i, s := range coordinates {
		args = append(args, s.X, s.Y)

		numFields := 2
		n := i * numFields

		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1]
	query += " ON CONFLICT (x, y) DO NOTHING;"

	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func findByCoordinateDB(db *sqlx.DB) (coordinate []Coordinate, err error) {
	query := `SELECT x, y FROM coordinate WHERE TRUE `

	err = db.Select(&coordinate, query)
	if err != nil {
		return nil, err
	}

	return coordinate, nil
}
