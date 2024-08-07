package main

import (
	"log"
	"olympy/auth-service/api"
	"olympy/auth-service/internal/config"
	"olympy/auth-service/internal/service"
	"olympy/auth-service/internal/storage"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/k0kubun/pp"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	migrateDB()

	storage, err := storage.NewAuthService(configs)
	if err != nil {
		log.Fatal(err)
	}
	serr := service.NewAuthServiceServer(storage)
	api := api.New(serr)
	go serr.StartRabbitMQConsumer()

	log.Fatal(api.RUN(configs))
}

func migrateDB() {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable PG_URL not declared or empty")
	}

	m, err := migrate.New("file://migrations", databaseURL)
	if err != nil {
		log.Fatalf("Migrate: error creating migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Printf("Migrate: no change")
		} else {
			log.Fatalf("Migrate: error applying migrations: %v", err)
		}
	} else {
		pp.Printf("Migrate: up success")
	}
}
