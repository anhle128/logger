package logger

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/labstack/gommon/log"
)

type Logger struct {
	service string
	env     string
	zap     *zap.Logger
	sugar   *zap.SugaredLogger
	out     io.Writer
	level   log.Lvl
}

// Init logger
func Init(service string, evn string) (*Logger, error) {

	var zapLogger *zap.Logger
	var err error
	if evn == "production" {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		zapLogger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		))
	} else {
		zapLogger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	sugar := zapLogger.Sugar()
	return &Logger{
		service: service,
		env:     evn,
		sugar:   sugar,
		zap:     zapLogger,
	}, nil
}

// GetEchoLogFormat - get echo log format
func (l Logger) GetEchoLogFormat() string {
	format := `{"level":"info","ts":"${time_custom}","id":"${id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out},"service":"%s"}` + "\n"
	return fmt.Sprintf(format, l.service)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func (l Logger) Sync() error {
	return l.zap.Sync()
}
