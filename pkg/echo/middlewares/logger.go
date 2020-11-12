package middlewares

import (
	"strconv"
	"time"

	"github.com/myugen/hexagonal-go-architecture/pkg/logger"

	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}

			log := logger.Log().WithFields(map[string]interface{}{
				"module":     "api",
				"id":         "id",
				"remote_ip":  c.RealIP(),
				"host":       req.Host,
				"method":     req.Method,
				"uri":        req.RequestURI,
				"user_agent": req.UserAgent(),
				"status":     res.Status,
				"latency":    stop.Sub(start).String(),
				"bytes_in":   reqSize,
				"bytes_out":  strconv.FormatInt(res.Size, 10),
			})
			if err != nil {
				log.WithField("error", err)
			}
			if c.QueryParams().Encode() != "" {
				log.WithField("query", c.QueryParams().Encode())
			}
			if req.Referer() != "" {
				log.WithField("referer", req.Referer())
			}
			log.Debugf("%s - %s", req.Method, req.RequestURI)
			return err
		}
	}
}
