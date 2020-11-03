package jsonfile

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sergiorra/Telemetry-backend/internal/models"
	. "github.com/sergiorra/Telemetry-backend/internal/repository"
)

// repository representation of repository into struct
type repository struct {
	fileName string
}

// NewRepository initialize jsonfile repository
func NewRepository(fileName string) SimulationRepo {
	return &repository{
		fileName: fileName,
	}
}

// GetSimulation fetch simulation data from json file
func (r *repository) GetSimulation() (*models.Simulation, error) {
	jsonFile, err := os.Open(r.fileName)
	if err != nil {
		return &models.Simulation{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &models.Simulation{}, err
	}

	var simulation models.Simulation
	json.Unmarshal(byteValue, &simulation)

	return &simulation, nil
}
