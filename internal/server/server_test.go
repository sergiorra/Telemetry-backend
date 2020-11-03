package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sergiorra/Telemetry-backend/internal/config"
	"github.com/sergiorra/Telemetry-backend/internal/models"
	"github.com/sergiorra/Telemetry-backend/internal/repository/jsonfile"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLoadFrontendFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	s := buildServer()
	resRecorder := httptest.NewRecorder()

	s.Router().ServeHTTP(resRecorder, req)

	res := resRecorder.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
}

func TestWebsocketConnection(t *testing.T) {
	wsURL := url.URL{Scheme: "ws", Host: "0.0.0.0:3000", Path: "/replay"}
	c, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err != nil {
		t.Fatalf("Cannot connect to the websocket %v", err)
	}
	c.Close()
}

func TestWebsocketCommands(t *testing.T) {
	wsURL := url.URL{Scheme: "ws", Host: "0.0.0.0:3000", Path: "/replay"}
	c, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err != nil {
		t.Errorf("Cannot connect to the websocket %v", err)
	}
	defer c.Close()
	done := make(chan bool)

	go func() {
		defer close(done)
		tm, message, err := c.ReadMessage()
		if err != nil {
			t.Fatalf("Error in the message reception: %v (type %v)", err, tm)
		}
		responseGot, responseExpected := models.DataResponse{}, models.DataResponse{}
		response := []byte(`{"kind":"data","data":{"time":"09:01:00.011","gear":"1","rpm":10895,"speed":0}}`)
		json.Unmarshal(message, &responseGot)
		json.Unmarshal(response, &responseExpected)
		if responseGot != responseExpected {
			t.Fatal("Response received should be the same as the response expected")
		}
		done <- true
	}()
	<-done
}

func buildServer() Server {
	config := &config.Config{
		Server: config.ServerConfig{
			Host: "0.0.0.0",
			Port: "3000",
			PublicDir: "../static",
			SimfileDir: "../data/simfile.json",
		},
	}
	repo := jsonfile.NewRepository(config.Server.SimfileDir)
	return New(repo, config)
}