package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	retryCount              = 5
	retryMinWaitTimeSeconds = 5
	retryMaxWaitTimeSeconds = 15
)

func main() {
	client := resty.New().SetRetryCount(retryCount).
		SetRetryWaitTime(retryMinWaitTimeSeconds * time.Second).
		SetRetryMaxWaitTime(retryMaxWaitTimeSeconds * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() == http.StatusRequestTimeout ||
					r.StatusCode() >= http.StatusInternalServerError
			},
		)

	resp, err := client.R().Get("http://localhost:8989/")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
}
