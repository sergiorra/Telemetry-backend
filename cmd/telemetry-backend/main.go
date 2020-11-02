package main

import (
	"encoding/json"
	"fmt"
	"github.com/sergiorra/Telemetry-backend/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)


var simulation models.Simulation
var upgrader = websocket.Upgrader{}


func replay(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error opening websocket connection", http.StatusBadRequest)
	}
	defer ws.Close()

	incomingCommands := make(chan models.Command)
	play, stop, reset := make(chan bool),make(chan bool),make(chan bool)

	go readCommands(ws, incomingCommands)
	go control(ws, play, stop, reset)

	for {
		select {
		case nextCommand := <-incomingCommands:
			switch nextCommand.Status {
			case "play":
				play <- true
			case "stop":
				stop <- true
			case "reset":
				reset <- true
			}
		}
	}

}

func control(ws *websocket.Conn, play <-chan bool, stop <-chan bool, reset <-chan bool) {
	step := 0
	isSending := false
	currentTime := simulation.StartTime
	for {
		select {
		case <-play:
			isSending = true
			go sendData(ws, &step, &currentTime, &isSending)
		case <-stop:
			response := &models.StatusResponse{
				Kind: "status",
				Data: models.Status{
					Status: "stop",
				},
			}
			res, _ := json.Marshal(*response)
			ws.WriteMessage(0, res)
			isSending = false
		case <-reset:
			step = 0
			isSending = false
			currentTime = simulation.StartTime
		}
	}
}

func sendData(ws *websocket.Conn, step *int, currentTime *time.Time, isSending *bool) {
	for *isSending {
		nextTime := simulation.Data[*step].Time
		countdown := nextTime.Sub(*currentTime).Milliseconds()
		if countdown > 0 {
			time.Sleep(time.Duration(countdown) * time.Millisecond)
		}
		dataResponse := &models.DataResponse{
			Kind: "data",
			Data: simulation.Data[*step],
		}
		_ = ws.WriteJSON(dataResponse)
		*currentTime = simulation.Data[*step].Time
		*step++
		if !(*isSending) {
			break
		}
	}
}

func readCommands(ws *websocket.Conn, incomingCommands chan<- models.Command) {
	for {
		var command models.Command
		_, message, _ := ws.ReadMessage()
		_ = json.Unmarshal(message, &command)
		incomingCommands <- command
	}
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

	json.Unmarshal(byteValue, &simulation)

	http.Handle("/", http.FileServer(http.Dir("internal/static")))
	http.HandleFunc("/replay", replay)
	log.Fatal(http.ListenAndServe(":3000", nil))

	/* repo := jsonfile.NewRepository("internal/data/simfile.json")
	s := server.New(repo)
	log.Fatal(http.ListenAndServe(":3000", s.Router()))*/

}

