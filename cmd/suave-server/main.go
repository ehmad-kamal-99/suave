package main

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/ehmad-kamal-99/suave"
	"github.com/ehmad-kamal-99/suave/config"
	"github.com/ehmad-kamal-99/suave/gen/restapi"
	"github.com/ehmad-kamal-99/suave/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt := runtime.NewRuntime()

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}

	server.ConfigureAPI()

	done := make(chan bool)

	go gracefulShutdown(server, rt, done)

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-done
}

func gracefulShutdown(server *restapi.Server, rt *runtime.Runtime, done chan<- bool) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTSTP, syscall.SIGTERM)

	<-quit

	log().Info("server is shutting down...")

	if err := server.Shutdown(); err != nil {
		logger.Warnf("could not gracefully shutdown the server: %+v", err)
	}

	log().Info("Closing redis-bloom connection")

	if err := rt.Service().Close(); err != nil {
		logger.Warnf("could not gracefully shutdown the mongo client: %+v", err)
	}

	close(done)
}

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
		"package": "suave-main",
	})
}
