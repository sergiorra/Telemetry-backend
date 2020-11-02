package main

import (
	"log"
	"net/http"

	"github.com/sergiorra/Telemetry-backend/internal/repository/jsonfile"
	"github.com/sergiorra/Telemetry-backend/internal/server"
)


func main() {
	repo := jsonfile.NewRepository("internal/data/simfile.json")
	s := server.New(repo)
	log.Fatal(http.ListenAndServe(":3000", s.Router()))
}

