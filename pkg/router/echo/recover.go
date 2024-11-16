package echo

import (
	"net/http"
	"runtime/debug"

	"github.com/omnia-core/go-echo-template/pkg/log"

	"github.com/labstack/echo/v4"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//nolint:errcheck
			defer func() error {

				if r := recover(); r != nil {
					log.New().Errorf("echo middleware recover: %v\n%v", r, string(debug.Stack()))

					return c.JSON(http.StatusInternalServerError, "Internal Server Error")
				}
				return nil
			}()
			return next(c)
		}
	}
}
