package models

// StatusResponse representation of statusResponse into struct
type StatusResponse struct {
	Kind string 	`jsonfile:"kind"`
	Data Status 	`jsonfile:"data"`
}
