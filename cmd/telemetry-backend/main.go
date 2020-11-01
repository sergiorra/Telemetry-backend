package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Simulation struct {
	StartTime 	time.Time 	`json:"startTime"`
	Data 		[]Data 		`json:"data"`
}

type Data struct {
	Time  	time.Time 	`json:"time"`
	Gear  	string 		`json:"gear"`
	Rpm   	int    		`json:"rpm"`
	Speed 	int    		`json:"speed"`
}

func main() {
	jsonFile, err := os.Open("internal/data/simfile.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var simulation Simulation

	json.Unmarshal(byteValue, &simulation)

	http.Handle("/", http.FileServer(http.Dir("internal/static")))
	http.ListenAndServe(":3000", nil)

}

