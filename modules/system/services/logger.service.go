package services

import (
	"encoding/json"
	"fmt"

	config "github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLevel(level string) zapcore.Level {

	switch level {
	case "debug":
		{
			return -1
		}
	case "info":
		{
			return 0
		}
	case "warn":
		{
			return 1
		}
	case "error":
		{
			return 2
		}
	default:
		{
			return 2
		}
	}
}

func NewAppLogger() *zap.Logger {
	level := config.GetString("logging.level")

	format := config.GetString("logging.format")

	rawJSON := []byte(fmt.Sprintf(`{
		"level": "%s",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoding": "%s",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`, level, format))
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	return logger

}
