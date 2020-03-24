package logger

import (
	"fmt"
)

// MockedLoggerImpl - mock for logger
type MockedLoggerImpl struct {
}

//
// ──────────────────────────────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: E N T R Y   P R I N T   F A M I L Y   F U N C T I O N S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────────────────────
//

// Debug - implement from Ilogger
func (l MockedLoggerImpl) Debug(mgs string) {
	fmt.Println(mgs)
}

// Info - implement from Ilogger
func (l MockedLoggerImpl) Info(mgs string) {
	fmt.Println(mgs)
}

// Warn - implement from Ilogger
func (l MockedLoggerImpl) Warn(mgs string) {
	fmt.Println(mgs)
}

// Error - implement from Ilogger
func (l MockedLoggerImpl) Error(mgs string) {
	fmt.Println(mgs)
}

// Fatal - implement from Ilogger
func (l MockedLoggerImpl) Fatal(mgs string) {
	fmt.Println(mgs)
}

// Panic - implement from Ilogger
func (l MockedLoggerImpl) Panic(mgs string) {
	fmt.Println(mgs)
}

//
// ────────────────────────────────────────────────────────────────────────────────────────────────── II ──────────
//   :::::: E N T R Y   P R I N T F   F A M I L Y   F U N C T I O N S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────────────────────────────────
//

// Debugf - implement from Ilogger
func (l MockedLoggerImpl) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Infof - implement from Ilogger
func (l MockedLoggerImpl) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Warnf - implement from Ilogger
func (l MockedLoggerImpl) Warnf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Errorf - implement from Ilogger
func (l MockedLoggerImpl) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Fatalf - implement from Ilogger
func (l MockedLoggerImpl) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Panicf - implement from Ilogger
func (l MockedLoggerImpl) Panicf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

//
// ──────────────────────────────────────────────────────────────────────────────────────────────────── III ──────────
//   :::::: E N T R Y   P R I N T L N   F A M I L Y   F U N C T I O N S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────────────────────────
//

// Debugw - implement from Ilogger
func (l MockedLoggerImpl) Debugw(mgs string, keysAndValues ...interface{}) {
	fmt.Println(mgs)
}

// Infow - implement from Ilogger
func (l MockedLoggerImpl) Infow(mgs string, keysAndValues ...interface{}) {
	fmt.Println(mgs)
}

// Warnw - implement from Ilogger
func (l MockedLoggerImpl) Warnw(mgs string, keysAndValues ...interface{}) {
	fmt.Println(mgs)
}

// Errorw - implement from Ilogger
func (l MockedLoggerImpl) Errorw(mgs string, keysAndValues ...interface{}) {
	fmt.Println(mgs)
}

// Fatalw - implement from Ilogger
func (l MockedLoggerImpl) Fatalw(mgs string, keysAndValues ...interface{}) {
	fmt.Println(mgs)
}

// Panicw - implement from Ilogger
func (l MockedLoggerImpl) Panicw(mgs string, keysAndValues ...interface{}) {
	fmt.Println(mgs)
}
