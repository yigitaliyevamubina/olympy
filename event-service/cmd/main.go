package main

import (
	"COMPETITIONS/olympy/event-service/api"
	"COMPETITIONS/olympy/event-service/internal/config"
	service "COMPETITIONS/olympy/event-service/internal/service"
	"COMPETITIONS/olympy/event-service/internal/storage"
	"log"
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

	storage, err := storage.NewEventService(configs)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.NewEventService(*storage))

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
