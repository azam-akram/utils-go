package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

var counter = 0

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("counter", counter)

	if counter < 3 {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		counter++
		return
	}

	fmt.Println("Have recovered from problem")

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
	fmt.Println("Server is listening at 8989")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8989", nil))
}
