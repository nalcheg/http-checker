package main

import (
	"log"
	"os"
	"strconv"

	"github.com/nalcheg/http-checker/application"
	"github.com/nalcheg/http-checker/repository"
)

func main() {
	wcEnv := os.Getenv("WORKERS_COUNT")
	wc, err := strconv.ParseInt(wcEnv, 10, 64)
	if err != nil {
		wc = 10
	}

	repo, err := repository.NewRepository(os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}

	app, err := application.NewApplication(repo, wc)
	if err != nil {
		panic(err)
	}

	if err := app.Start(os.Getenv("CRON_STRING")); err != nil {
		log.Fatal(err)
	}

	log.Print("application started")
	select {}
}
