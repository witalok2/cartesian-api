package api

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cartesian-api/utils/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	defaultPort *string
	debug       *bool
	echoServer  *echo.Echo
)

func init() {
	defaultPort = flag.String("port", "9000", "port for the service HTTP")
	debug = flag.Bool("debug", false, "mod of the debug")
}

func Make() *echo.Echo {
	flag.Parse()

	echoServer = echo.New()

	// Esconde o cabeçalho do Echo
	echoServer.HideBanner = true

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Recover())

	if *debug {
		echoServer.Debug = true
		echoServer.Use(middleware.Logger())
		//log.EnableDebug(true)
	}

	return echoServer
}

func ProvideEchoInstance(task func(e *echo.Echo)) {
	task(echoServer)
}

func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		port = *defaultPort
	}

	echoServer.Logger.Fatal(echoServer.Start(":" + port))
}

func Use(middleware ...echo.MiddlewareFunc) {
	echoServer.Use(middleware...)
}

func UseCustomHTTPErrorHandler() {
	echoServer.HTTPErrorHandler = CustomHTTPErrorHandler
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Internal)
		}
	} else {
		msg = http.StatusText(code)
	}

	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	} else {
		log.File(time.Now().Format("errors/2006/01/0215h.log"), err.Error())
	}
}
