package main

import (
	"Wetao/internal/api_back/app/bootstrap"
	"Wetao/internal/database"
	"Wetao/pkg"
	"github.com/rs/zerolog/log"
)

func main() {
	pkg.LoggerInit()
	//Env
	pkg.EnvLoadInit()
	//WebApp
	app := bootstrap.HttpAppServer()
	bootstrap.HttpServiceProvider(app)
	_, err := database.InitDB()
	if err != nil {
		log.Print(err)
	}
	//Listen
	log.Err(app.Listen(":9761"))
}
