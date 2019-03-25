package logger

import (
	"fmt"

	"github.com/theonlyjohnny/jogger-go/jogger"
)

var (
	Log *jogger.Logger
)

func init() {
	var err error
	config := jogger.Config{
		AppName:    "dot-term-visualizer",
		LogLevel:   "debug",
		LogConsole: true,
		LogSyslog:  nil,
	}
	Log, err = jogger.CreateLogger(config)
	if err != nil {
		panic(fmt.Sprintf("Could not get logger: %s", err.Error()))
	}

}
