package main

import (
	"fmt"
	"net/http"

	sh "azam-akram/go/utils/http_utils/server/server_handler"
)

const LISTENING_PORT = 8989

func main() {
	fmt.Printf("Server is listening at port %d\n", LISTENING_PORT)

	h := sh.NewHTTPServer()
	http.HandleFunc("/", h.Handle)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", LISTENING_PORT), nil); err != nil {
		fmt.Println(err)
	}
}
