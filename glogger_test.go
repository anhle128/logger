package logger_test

import (
	"testing"

	"github.com/anhle128/logger"
)

func TestLog(t *testing.T) {
	log, err := logger.Init("test", "production")
	if err != nil {
		t.Error(err)
		return
	}
	log.Info("ahihihihihihiihi")
	log.Infow(
		"this is message",
		"key1", "value1",
		"key2", "value2",
	)
}
