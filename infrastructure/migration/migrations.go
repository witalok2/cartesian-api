package migrations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cartesian-api/point"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sqlx.DB) error {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	err = goose.Up(db.DB, path+"/infrastructure/migration")
	fmt.Println(err)
	return err
}

func DownMigrations(db *sqlx.DB) error {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	err = goose.Down(db.DB, path+"/infrastructure/migration")
	if err != nil {
		log.Println(err)
	}
	return err
}

func LoadFilePoint(path string) error {
	points := []point.Coordinate{}

	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err
	}

	defer jsonFile.Close()

	byteValueJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(byteValueJSON, &points)
	if err != nil {
		log.Println(err)
		return err
	}

	err = point.CreateMultipleCoordinate(points)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
