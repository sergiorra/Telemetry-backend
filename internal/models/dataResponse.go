package models

// DataResponse representation of dataResponse into struct
type DataResponse struct {
	Kind string 	`json:"kind"`
	Data Data 		`json:"data"`
}
