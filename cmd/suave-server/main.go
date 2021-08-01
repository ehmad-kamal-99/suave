package main

import (
	"strconv"

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

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
