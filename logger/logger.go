package logger

import (
	"fmt"
	"net/url"
	"os"

	"go.uber.org/zap"

	"github.com/ZeroTechh/VelocityCore/config"
)

func newWinFileSink(u *url.URL) (zap.Sink, error) {
	// Remove leading slash left by url.Parse()
	return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

// getLogFilePath is used to get file directory based on filename
func getLogFilePath(fileName string) string {
	zap.RegisterSink("winfile", newWinFileSink)
	return config.LogDir + fileName
}

// GetLogger returns a logger
func GetLogger(fileName string, debug bool) *zap.Logger {
	var cfg zap.Config

	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.Encoding = "json"

	cfg.OutputPaths = []string{
		getLogFilePath(fileName),
	}

	cfg.ErrorOutputPaths = []string{
		getLogFilePath("errors.log"),
	}

	logger, err := cfg.Build()
	defer logger.Sync()
	if err != nil {
		panic(fmt.Sprintf("Cant Load Logger Because Of Error %+v", err))
	}

	return logger
}
