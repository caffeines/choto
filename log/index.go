package log

import (
	"fmt"
	"os"

	"github.com/caffeines/choto/log/hooks"

	"github.com/sirupsen/logrus"
)

var defLogger *logrus.Logger

// SetupLog ...
func SetupLog() {
	defLogger = logrus.New()
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defLogger.SetOutput(f)
	// defLogger.Out = os.Stdout
	defLogger.AddHook(hooks.NewHook())
}

// Log returns user defined logger
func Log() *logrus.Logger {
	return defLogger
}
