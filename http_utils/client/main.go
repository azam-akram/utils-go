package main

import (
	"azam-akram/go/utils/logger"
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

	logger.GetLogger().PrintKeyValue("Client:main", "Error", err)
	logger.GetLogger().PrintKeyValue("Client:main", "StatusCode", resp.StatusCode())
	logger.GetLogger().PrintKeyValue("Client:main", "Status", resp.Status())
	logger.GetLogger().PrintKeyValue("Client:main", "resp", resp)
}
