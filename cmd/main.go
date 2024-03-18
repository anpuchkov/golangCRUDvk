package main

import (
	"context"
	"go.uber.org/zap"
	"log"
	"net/http"
	"vkFilmoteka/configs"
	"vkFilmoteka/database/psql"
	db2 "vkFilmoteka/server"
	"vkFilmoteka/server/logs"
	"vkFilmoteka/server/middleware"
)

func main() {
	ctx := context.Background()
	cfg, err := configs.ConfigInit()
	if err != nil {
		log.Println("failed to initialize configs: ", err)
		return
	}
	logger := logs.InitLogger()
	if err != nil {
		log.Println("failed to initialize logger: ", err)
		return
	}

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Println("failed to sync logger: ", err)
		}
	}(logger)

	database, err := psql.InitPostgresConnection(ctx, *cfg.DBConfig)
	if err != nil {
		log.Println("unable to connect to the database: ", err)
		return
	}

	defer database.Close()

	db2.SetupRoutes(database, logger)

	s := http.Server{
		Addr:    ":8080",
		Handler: middleware.LoggingMiddleware(logger, http.DefaultServeMux),
	}

	s.ListenAndServe()

}
