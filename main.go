package main

import (
	"azam-akram/go/utils/json_utils"
	"azam-akram/go/utils/time_utils"
	"azam-akram/go/utils/url_utils"
)

func main() {
	jsonHandler := json_utils.JsonHandler{}
	jsonHandler.DisplayAllJsonHandlers()

	time_utils.DemoTime("2022-12-31T00:00:00Z")

	url_utils.DemoUrlEncode()
}
