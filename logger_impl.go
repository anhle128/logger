package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// With - with
func (l Logger) With(fields ...zap.Field) Logger {
	l.zap = l.zap.With(fields...)
	return l
}

//
// ──────────────────────────────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: E N T R Y   P R I N T   F A M I L Y   F U N C T I O N S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────────────────────
//

// Debug - implement from Ilogger
func (l Logger) Debug(mgs string) {
	l.zap.Debug(mgs, zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Info - implement from Ilogger
func (l Logger) Info(mgs string) {
	l.zap.Info(mgs, zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Warn - implement from Ilogger
func (l Logger) Warn(mgs string) {
	l.zap.Warn(mgs, zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Error - implement from Ilogger
func (l Logger) Error(mgs string) {
	l.zap.Error(mgs, zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Fatal - implement from Ilogger
func (l Logger) Fatal(mgs string) {
	l.zap.Fatal(mgs, zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Panic - implement from Ilogger
func (l Logger) Panic(mgs string) {
	l.zap.Panic(mgs, zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

//
// ────────────────────────────────────────────────────────────────────────────────────────────────── II ──────────
//   :::::: E N T R Y   P R I N T F   F A M I L Y   F U N C T I O N S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────────────────────────────────
//

// Debugf - implement from Ilogger
func (l Logger) Debugf(format string, args ...interface{}) {
	l.zap.Debug(fmt.Sprintf(format, args...), zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Infof - implement from Ilogger
func (l Logger) Infof(format string, args ...interface{}) {
	l.zap.Info(fmt.Sprintf(format, args...), zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Warnf - implement from Ilogger
func (l Logger) Warnf(format string, args ...interface{}) {
	l.zap.Warn(fmt.Sprintf(format, args...), zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Errorf - implement from Ilogger
func (l Logger) Errorf(format string, args ...interface{}) {
	l.zap.Error(fmt.Sprintf(format, args...), zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Fatalf - implement from Ilogger
func (l Logger) Fatalf(format string, args ...interface{}) {
	l.zap.Fatal(fmt.Sprintf(format, args...), zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

// Panicf - implement from Ilogger
func (l Logger) Panicf(format string, args ...interface{}) {
	l.zap.Panic(fmt.Sprintf(format, args...), zap.Field{
		Key:    "server",
		String: l.service,
		Type:   zapcore.StringType,
	})
}

//
// ──────────────────────────────────────────────────────────────────────────────────────────────────── III ──────────
//   :::::: E N T R Y   P R I N T L N   F A M I L Y   F U N C T I O N S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────────────────────────
//

// Debugw - implement from Ilogger
func (l Logger) Debugw(mgs string, keysAndValues ...interface{}) {
	l.sugar.With("server", l.service).Debugw(mgs, keysAndValues...)
}

// Infow - implement from Ilogger
func (l Logger) Infow(mgs string, keysAndValues ...interface{}) {
	l.sugar.With("server", l.service).Infow(mgs, keysAndValues...)
}

// Warnw - implement from Ilogger
func (l Logger) Warnw(mgs string, keysAndValues ...interface{}) {
	l.sugar.With("server", l.service).Warnw(mgs, keysAndValues...)
}

// Errorw - implement from Ilogger
func (l Logger) Errorw(mgs string, keysAndValues ...interface{}) {
	l.sugar.With("server", l.service).Errorw(mgs, keysAndValues...)
}

// Fatalw - implement from Ilogger
func (l Logger) Fatalw(mgs string, keysAndValues ...interface{}) {
	l.sugar.With("server", l.service).Fatalw(mgs, keysAndValues...)
}

// Panicw - implement from Ilogger
func (l Logger) Panicw(mgs string, keysAndValues ...interface{}) {
	l.sugar.With("server", l.service).Panicw(mgs, keysAndValues...)
}
