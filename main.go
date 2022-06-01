package main

import (
	jsonHandling "azam-akram/go/utils/json_handling"
	"fmt"
)

func main() {
	fmt.Println("Starting ..")

	jsonHandler := jsonHandling.JsonHandler{}
	jsonHandler.DisplayAllJsonHandlers()
}
