package models

// StatusResponse representation of statusResponse into struct
type StatusResponse struct {
	Kind string 	`json:"kind"`
	Data Status 	`json:"data"`
}
