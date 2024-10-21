package main

import (
	"log"
	"os"

	"github.com/ribeirosaimon/shortify-read/config/server"
	"github.com/ribeirosaimon/shortify-read/internal/controller"
	"github.com/ribeirosaimon/shortify-read/internal/usecase"
	"github.com/ribeirosaimon/tooltip/tserver"
)

// @title Shortify
// @swagger: "2.0"
// @description Shortener api
// @termsOfService  http://swagger.io/terms/
// @contact.url    http://www.swagger.io/support
// @contact.email  suporte@swagger.io
// @host      localhost:8081
// @BasePath  /
func main() {
	if myEnv := os.Getenv("ENVIRONMENT"); myEnv != "" {
		tserver.StartEnv(tserver.Environment(myEnv))
	}

	urlUseCase := usecase.NewUrlRecord()

	server.NewServices(
		server.WithUrlUseCase(urlUseCase),
	)
	controller.Start()

	config := tserver.GetEnvironment()
	if config.Env == "" {
		log.Fatal("Environment variable not set")
	}
	tserver.NewServer(config)
}
