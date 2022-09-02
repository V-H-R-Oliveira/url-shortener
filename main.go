package main

import (
	"log"
	"url-shortener/1.0/config"
	"url-shortener/1.0/database"
	"url-shortener/1.0/routes"
	"url-shortener/1.0/server"
	"url-shortener/1.0/utils"
)

func main() {
	isProd, err := utils.IsProd()

	if err != nil {
		log.Fatal(err)
	}

	var db database.Repository

	if isProd {
		db = database.NewRedisDatabase()
	} else {
		db = database.NewMemoryDatabase()
	}

	envPort, err := utils.GetEnvVariable("PORT")

	if err != nil {
		envPort = config.DEFAULT_SERVER_PORT
	}

	server.ListenServer(server.CreateServer(routes.CreateAppRouter(db), envPort))
}
