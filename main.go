package main

import (
	"azam-akram/go/utils/json_utils"
	"fmt"
)

func main() {
	fmt.Println("Starting ..")

	jsonHandler := json_utils.JsonHandler{}
	jsonHandler.DisplayAllJsonHandlers()
}
