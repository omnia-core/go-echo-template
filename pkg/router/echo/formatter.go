package echo

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/omnia-core/go-echo-template/pkg/log"

	"github.com/labstack/echo/v4"
)

const (
	_maxRequestBodySize = 10 * 1024 // 10KB (setnry 에서는 10KB 이상의 body 를 보내지 않는다.)
)

var (
	ServerIp = "default ip"
	Commit   = "default commit"
	Branch   = "default branch"
)

func Formatter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			response := c.Response()

			var buf []byte
			if request.ContentLength < _maxRequestBodySize {
				buf, _ = io.ReadAll(request.Body)
				request.Body = io.NopCloser(bytes.NewBuffer(buf))
			}

			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			if c.Request().Method == http.MethodGet {
				switch path := c.Path(); path {
				case "/":
					return nil
				case "/ping":
					return nil
				}
			}

			bytesIn, _ := strconv.Atoi(request.Header.Get(echo.HeaderContentLength))

			log.Logger.WithFields(map[string]interface{}{
				"status":        response.Status,
				"remote_ip":     c.RealIP(),
				"server_ip":     ServerIp, //main.go 에서 선언한 전역변수
				"method":        request.Method,
				"uri":           request.RequestURI,
				"path":          request.URL.Path,
				"latency":       stop.Sub(start).Milliseconds(),
				"latency_human": stop.Sub(start).String(),
				"referer":       request.Referer(),
				"user_agent":    request.UserAgent(),
				"bytes_in":      bytesIn,
				"bytes_out":     response.Size,
				"request_body":  string(buf),
				"commit":        Commit,
				"branch":        Branch,
				"echo_path":     c.Path(),
			}).Info()

			return nil
		}
	}
}
