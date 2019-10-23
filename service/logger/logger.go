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
