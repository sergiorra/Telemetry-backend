package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Simulation struct {
	StartTime 	time.Time 		`json:"startTime"`
	Data 		[]Data 		`json:"data"`
}

type Data struct {
	Time  time.Time `json:"time"`
	Gear  string `json:"gear"`
	Rpm   int    `json:"rpm"`
	Speed int    `json:"speed"`
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

	fmt.Println("Simulation startTime: ", simulation.StartTime)
	for i := 0; i < len(simulation.Data); i++ {
		fmt.Println("Data time: ", simulation.Data[i].Time)
		fmt.Println("Data gear: ", simulation.Data[i].Gear)
		fmt.Println("Data rpm: ", simulation.Data[i].Rpm)
		fmt.Println("Data speed: ", simulation.Data[i].Speed)
	}

}
