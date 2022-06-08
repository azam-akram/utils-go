package main

import (
	"azam-akram/go/utils/logger"
	"encoding/json"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

var counter = 0

func handler(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().PrintKeyValue("Server::handler", "counter", counter)

	if counter < 3 {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		counter++
		return
	}

	logger.GetLogger().PrintKeyValue("handler: recovered from problem", "counter", counter)

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

func main() {
	logger.GetLogger().PrintKeyValue("main: Server is listening at", "port", 8989)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8989", nil))
}
