package redisbloom

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ehmad-kamal-99/suave/config"
)

func log() *logger.Entry {
	level, err := logger.ParseLevel(viper.GetString(config.LogLevel))
	if err != nil {
		logger.SetLevel(logger.DebugLevel)
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "redis",
	})
}
