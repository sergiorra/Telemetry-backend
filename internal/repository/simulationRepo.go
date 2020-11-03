package repository

import "github.com/sergiorra/Telemetry-backend/internal/models"

type SimulationRepo interface {
	GetSimulation() (*models.Simulation, error)
}