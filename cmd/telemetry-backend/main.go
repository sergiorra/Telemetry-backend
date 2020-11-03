package main

import (
	"log"
	"net/http"

	"github.com/sergiorra/Telemetry-backend/internal/config"
	"github.com/sergiorra/Telemetry-backend/internal/repository/jsonfile"
	"github.com/sergiorra/Telemetry-backend/internal/server"
)


func main() {
	config := config.New()
	repo := jsonfile.NewRepository(config.Server.SimfileDir)
	s := server.New(repo, config)

	log.Fatal(http.ListenAndServe(config.Server.Host + ":" + config.Server.Port, s.Router()))
}

