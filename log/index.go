package log

import (
	"os"

	"github.com/caffeines/choto/log/hooks"

	"github.com/sirupsen/logrus"
)

var defLogger *logrus.Logger

// SetupLog ...
func SetupLog() {
	defLogger = logrus.New()
	defLogger.Out = os.Stdout
	defLogger.AddHook(hooks.NewHook())
}

// Log returns user defined logger
func Log() *logrus.Logger {
	return defLogger
}
