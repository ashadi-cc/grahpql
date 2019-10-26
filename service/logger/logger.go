package logger

import (
	"gql-ashadi/service/config"
	"os"

	"github.com/op/go-logging"
)

var logger *logging.Logger

//GetLogger get logger
func GetLogger() *logging.Logger {
	if logger == nil {
		config := config.GetConfig()
		backend := logging.NewLogBackend(os.Stderr, "", 0)
		format := logging.MustStringFormatter(config.LogFormat)
		backendFormatter := logging.NewBackendFormatter(backend, format)

		backendLeveled := logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.INFO, "")
		if config.DebugMode {
			backendLeveled.SetLevel(logging.DEBUG, "")
		}

		logging.SetBackend(backendLeveled)
		logger = logging.MustGetLogger(config.AppName)
	}

	return logger
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Warning(args ...interface{}) {
	GetLogger().Warning(args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func Critical(args ...interface{}) {
	GetLogger().Critical(args...)
}
