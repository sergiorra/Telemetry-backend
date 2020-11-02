package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sergiorra/Telemetry-backend/internal/models"
	"github.com/sergiorra/Telemetry-backend/internal/repository"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type server struct {
	router 	http.Handler
	repo 	repository.SimulationRepo
}

type Server interface {
	Router() http.Handler
}

func New(repo repository.SimulationRepo) Server {
	a := &server{repo: repo}
	router(a)
	return a
}

func router(s *server) {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("internal/static"))).Methods(http.MethodGet)
	r.HandleFunc("/replay", s.replay)

	s.router = r
}

func (s *server) Router() http.Handler {
	return s.router
}

var upgrader = websocket.Upgrader{}

func (s *server) replay(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error opening websocket connection", http.StatusBadRequest)
	}
	defer ws.Close()

    simulation, err := s.repo.GetSimulation()
	if err != nil {
		http.Error(w, "Error getting simulation data", http.StatusInternalServerError)
	}

	incomingCommands := make(chan models.Command)
	play, stop, reset := make(chan bool),make(chan bool),make(chan bool)

	go s.readCommands(ws, incomingCommands)
	go s.controlCommands(ws, simulation, play, stop, reset)

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

func (s *server) controlCommands(ws *websocket.Conn, simulation *models.Simulation, play <-chan bool, stop <-chan bool, reset <-chan bool) {
	step := 0
	isSending := false
	currentTime := simulation.StartTime
	for {
		select {
		case <-play:
			isSending = true
			go s.sendData(ws, simulation, &step, &currentTime, &isSending)
		case <-stop:
			isSending = false
			statusResponse := &models.StatusResponse{
				Kind: "status",
				Data: models.Status{
					Status: "stop",
				},
			}
			resBytes, _ := json.Marshal(*statusResponse)
			_ = ws.WriteMessage(0, resBytes)
		case <-reset:
			isSending = false
			step = 0
			currentTime = simulation.StartTime
			dataResponse := &models.DataResponse{
				Kind: "data",
				Data: simulation.Data[step],
			}
			_ = ws.WriteJSON(dataResponse)
		}
	}
}

func (s *server) sendData(ws *websocket.Conn, simulation *models.Simulation, step *int, currentTime *time.Time, isSending *bool) {
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
	}
}

func (s *server) readCommands(ws *websocket.Conn, incomingCommands chan<- models.Command) {
	for {
		var command models.Command
		_, message, _ := ws.ReadMessage()
		_ = json.Unmarshal(message, &command)
		incomingCommands <- command
	}
}