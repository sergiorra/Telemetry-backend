package main

import (
	"fmt"
	"github.com/sergiorra/Telemetry-backend/internal/config"
	"github.com/sergiorra/Telemetry-backend/internal/repository/jsonfile"
	"github.com/sergiorra/Telemetry-backend/internal/server"
	"log"
	"net/http"
)


func main() {
	config := config.New()
	repo := jsonfile.NewRepository(config.Server.SimfileDir)
	s := server.New(repo, config)

	httpAddr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Fatal(http.ListenAndServe(httpAddr, s.Router()))
}

