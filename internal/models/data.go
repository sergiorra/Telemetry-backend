package models

import "time"

// Data representation of data into struct
type Data struct {
	Time  	time.Time 	`json:"time"`
	Gear  	string 		`json:"gear"`
	Rpm   	int    		`json:"rpm"`
	Speed 	int    		`json:"speed"`
}
