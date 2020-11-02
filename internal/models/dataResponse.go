package models

// DataResponse representation of dataResponse into struct
type DataResponse struct {
	Kind string 	`jsonfile:"kind"`
	Data Data 		`jsonfile:"data"`
}
