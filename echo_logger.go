package logger

import (
	"net/http"
	"strconv"
	"sync"
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
		pool    *sync.Pool
	}
)

// EchoLogger returns a middleware that logs HTTP requests.
func (logger Logger) EchoLogger() echo.MiddlewareFunc {
	return logger.loggerWithConfig(LoggerConfig{
		Skipper: middleware.DefaultSkipper,
	})
}

func (logger Logger) loggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {

	config.pool = &sync.Pool{
		New: func() interface{} {
			return initFields()
		},
	}

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

			fields := config.pool.Get().([]zap.Field)
			resetFields(fields)
			defer config.pool.Put(fields)

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields[0].String = id
			fields[1].String = req.Host
			fields[2].String = req.RequestURI
			fields[3].String = req.Method
			p := req.URL.Path
			if p == "" {
				p = "/"
			}
			fields[4].String = p
			fields[5].String = stop.Sub(start).String()
			cl := req.Header.Get(echo.HeaderContentLength)
			if cl == "" {
				cl = "0"
			}
			fields[6].String = cl
			fields[7].String = strconv.FormatInt(res.Size, 10)
			if err != nil {
				he, _ := err.(*echo.HTTPError)
				fields[8].Integer = int64(he.Code)
				logger.With(fields...).Error(he.Internal.Error())
				return
			}

			fields[8].Integer = int64(res.Status)

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

func initFields() []zap.Field {
	fields := make([]zap.Field, 9)
	fields[0] = zap.Field{
		Key:    "append",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[1] = zap.Field{
		Key:    "host",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[2] = zap.Field{
		Key:    "uri",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[3] = zap.Field{
		Key:    "method",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[4] = zap.Field{
		Key:    "source",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[5] = zap.Field{
		Key:    "latency_human",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[6] = zap.Field{
		Key:    "bytes_in",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[7] = zap.Field{
		Key:    "bytes_out",
		String: "",
		Type:   zapcore.StringType,
	}
	fields[8] = zap.Field{
		Key:     "status",
		Integer: 0,
		Type:    zapcore.Int64Type,
	}
	return fields
}

func resetFields(fields []zap.Field) {
	for _, field := range fields {
		field.String = ""
		field.Integer = 0
	}
}
