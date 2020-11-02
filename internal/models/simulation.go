package models

import "time"

// Simulation representation of simulation into struct
type Simulation struct {
	StartTime 	time.Time 	`json:"startTime"`
	Data 		[]Data 		`json:"data"`
}
