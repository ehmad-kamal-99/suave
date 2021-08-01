package config

import (
	"github.com/spf13/viper"
)

// keys for username availability service configuration.
const (
	DBHost       = "db.name"
	DBPort       = "db.port"
	DBPassword   = "db.password"
	DBClientName = "db.client.name"
	DBFilterName = "db.filter.name"

	ServerHost = "server.host"
	ServerPort = "server.port"
	LogLevel   = "log.level"
)

func init() {
	// env variables for database.
	_ = viper.BindEnv(DBHost, "DB_HOST")
	_ = viper.BindEnv(DBPort, "DB_PORT")
	_ = viper.BindEnv(DBPassword, "DB_PASSWORD")
	_ = viper.BindEnv(DBClientName, "DB_CLIENT_NAME")
	_ = viper.BindEnv(DBFilterName, "DB_FILTER_NAME")

	// env var for main server.
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")
	_ = viper.BindEnv(LogLevel, "LOG_LEVEL")

	// defaults for redis-bloom.
	viper.SetDefault(DBHost, "localhost")
	viper.SetDefault(DBPort, "6379")
	viper.SetDefault(DBPassword, "")
	viper.SetDefault(DBClientName, "go-bloom-client")
	viper.SetDefault(DBFilterName, "username")

	// defaults for main server.
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
	viper.SetDefault(LogLevel, "debug")
}
