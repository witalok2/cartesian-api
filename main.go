package main

import (
	"github.com/cartesian-api/point"
	"github.com/cartesian-api/utils/api"
)

func main() {
	api.Make()
	api.UseCustomHTTPErrorHandler()
	api.ProvideEchoInstance(point.AddRoutes)
	api.Run()
}
