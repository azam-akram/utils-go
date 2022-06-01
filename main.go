package main

import (
	jsonHandling "azam-akram/go/utils/json_utils"
	"fmt"
)

func main() {
	fmt.Println("Starting ..")

	jsonHandler := jsonHandling.JsonHandler{}
	jsonHandler.DisplayAllJsonHandlers()
}
