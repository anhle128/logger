package logger

// ILogger interface
type ILogger interface {
	Debug(mgs string)
	Info(mgs string)
	Warn(mgs string)
	Error(mgs string)
	Fatal(mgs string)
	Panic(mgs string)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debugw(mgs string, keysAndValues ...interface{})
	Infow(mgs string, keysAndValues ...interface{})
	Warnw(mgs string, keysAndValues ...interface{})
	Errorw(mgs string, keysAndValues ...interface{})
	Fatalw(mgs string, keysAndValues ...interface{})
	Panicw(mgs string, keysAndValues ...interface{})
}
