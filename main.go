package main

import (
	"fmt"
	"log"

	"github.com/cartesian-api/config"
	migration "github.com/cartesian-api/infrastructure/migration"

	"github.com/cartesian-api/point"
	"github.com/cartesian-api/utils/api"
	"github.com/cartesian-api/utils/database"
)

func main() {
	instance := database.MustGetByFile(config.POSTGRES_FILE)

	err := migration.RunMigrations(instance)
	if err != nil {
		err = migration.DownMigrations(instance)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(err)
		return
	}

	err = migration.LoadFilePoint(config.POINT_FILE)
	if err != nil {
		log.Println(err)
		return
	}

	api.Make()
	api.UseCustomHTTPErrorHandler()
	api.ProvideEchoInstance(point.AddRoutes)
	api.Run()
}
