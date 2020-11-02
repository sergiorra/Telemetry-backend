package models

import "time"

// Simulation representation of simulation into struct
type Simulation struct {
	StartTime 	time.Time 	`jsonfile:"startTime"`
	Data 		[]Data 		`jsonfile:"data"`
}
