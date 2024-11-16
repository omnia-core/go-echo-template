package redoc

import (
	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

func Middleware() echo.MiddlewareFunc {
	doc := redoc.Redoc{
		Title:       "Go Echo REST API",
		Description: "Go Echo REST API Description",
		SpecFile:    "docs/swagger.json",
		SpecPath:    "redoc/swagger.json",
		DocsPath:    "/redoc",
	}

	return echoredoc.New(doc)
}
