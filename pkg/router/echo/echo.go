package echo

import (
	"encoding/base64"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/omnia-core/go-echo-template/pkg/config"
	"github.com/omnia-core/go-echo-template/pkg/log"
	"github.com/omnia-core/go-echo-template/pkg/redoc"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func New(cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.Logger = log.GetEchoLogger()

	secret, err := base64.StdEncoding.DecodeString(cfg.JWTConfig.Secret)
	if err != nil {
		log.New().Fatalf("%v", err)
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Use(middleware.CORS())
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(
		echoSwagger.DocExpansion("none"),
		echoSwagger.PersistAuthorization(true),
	))
	e.Use(redoc.Middleware())
	// Middleware: Timeout after 60 seconds
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Second * 60,
		Skipper: func(c echo.Context) bool {
			switch {
			case strings.HasPrefix(c.Request().URL.Path, "/debug/pprof"):
				return true
			default:
				return false
			}
		},
	}))
	// Middleware: Limit request body size to 128MB
	e.Use(middleware.BodyLimit("128M"))
	// Middleware: Gzip compression
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{}))
	// Middleware: Formatter for logging
	e.Use(Formatter())
	// Middleware: Recover from panic
	e.Use(Recover())
	// Middleware: JWT
	e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			switch {
			case c.Request().URL.Path == "/" || c.Request().URL.Path == "/ping":
				return true
			default:
				return false
			}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		},
		SigningKey: secret,
		SuccessHandler: func(c echo.Context) {
			// TODO: implement
			// Get claims from JWT token and set it to echo context
		},
	}))
	return e
}
