package logger

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	// LoggerConfig defines the config for Logger middleware.
	LoggerConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper
	}
)

// EchoLogger returns a middleware that logs HTTP requests.
func (logger Logger) EchoLogger() echo.MiddlewareFunc {
	return logger.loggerWithConfig(LoggerConfig{
		Skipper: middleware.DefaultSkipper,
	})
}

func (logger Logger) loggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			// fields := make(map[string]interface{})
			fields := []zap.Field{}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			fields = append(fields, zap.Field{
				Key:    "append",
				String: id,
				Type:   zapcore.StringType,
			})
			fields = append(fields, zap.Field{
				Key:    "host",
				String: req.Host,
				Type:   zapcore.StringType,
			})
			fields = append(fields, zap.Field{
				Key:    "uri",
				String: req.RequestURI,
				Type:   zapcore.StringType,
			})

			fields = append(fields, zap.Field{
				Key:    "method",
				String: req.Method,
				Type:   zapcore.StringType,
			})

			p := req.URL.Path
			if p == "" {
				p = "/"
			}
			fields = append(fields, zap.Field{
				Key:    "source",
				String: p,
				Type:   zapcore.StringType,
			})

			fields = append(fields, zap.Field{
				Key:    "latency_human",
				String: stop.Sub(start).String(),
				Type:   zapcore.StringType,
			})

			cl := req.Header.Get(echo.HeaderContentLength)
			if cl == "" {
				cl = "0"
			}
			fields = append(fields, zap.Field{
				Key:    "bytes_in",
				String: cl,
				Type:   zapcore.StringType,
			})

			fields = append(fields, zap.Field{
				Key:    "bytes_out",
				String: strconv.FormatInt(res.Size, 10),
				Type:   zapcore.StringType,
			})

			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					fields = append(fields, zap.Field{
						Key:     "status",
						Integer: int64(he.Code),
						Type:    zapcore.Int64Type,
					})
					logger.With(fields...).Error(he.Error())
				} else {
					logger.With(fields...).Error(err.Error())
				}
				return
			}

			fields = append(fields, zap.Field{
				Key:     "status",
				Integer: int64(res.Status),
				Type:    zapcore.Int64Type,
			})

			if res.Status == http.StatusOK || res.Status == http.StatusNoContent {
				logger.With(fields...).Info("success")
			} else {
				// not log with OPTION and status 204
				if req.Method == http.MethodOptions && res.Status == http.StatusNoContent {
					return
				}

				messErr := res.Header().Get("internal_error")
				if len(messErr) == 0 {
					messErr = "internal server error"
				}
				res.Header().Del("internal_error")
				logger.With(fields...).Error(messErr)
			}
			return
		}
	}
}
