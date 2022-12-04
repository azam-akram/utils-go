package server_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const LISTENING_PORT = 8989

type Profile struct {
	Name    string
	Hobbies []string
}

var COUNTER = 0

const MAX_COUNTER = 4

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

var handler Handler

type HTTPServer struct{}

func NewHTTPServer() Handler {
	if handler == nil {
		handler = &HTTPServer{}
	}

	return handler
}

func (h HTTPServer) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTTPServer::Handler", "counter = ", COUNTER)

	if COUNTER < MAX_COUNTER {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		COUNTER++
		return
	}

	fmt.Printf("HTTPServer: recovered from problem after %d failed attempts", COUNTER)

	profile := &Profile{
		Name:    "User",
		Hobbies: []string{"Sports", "walk"},
	}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
