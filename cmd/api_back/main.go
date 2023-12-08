package main

import (
	"Wetao/internal/api_back/app/bootstrap"
	"Wetao/internal/database"
	"Wetao/pkg"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	//Env
	pkg.EnvLoadInit()
	pkg.LoggerInit()
	//WebApp
	app := bootstrap.HttpAppServer()
	bootstrap.HttpServiceProvider(app)
	_, err := database.InitDB()
	if err != nil {
		log.Print(err)
	}

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "9761"
	}
	//
	appAddress := appHost + ":" + appPort
	//Listen
	log.Err(app.Listen(appAddress))
}
