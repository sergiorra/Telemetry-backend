package models

import "time"

// Data representation of data into struct
type Data struct {
	Time  	time.Time 	`jsonfile:"time"`
	Gear  	string 		`jsonfile:"gear"`
	Rpm   	int    		`jsonfile:"rpm"`
	Speed 	int    		`jsonfile:"speed"`
}
