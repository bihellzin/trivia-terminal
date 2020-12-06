package apiConnection

import (
	"net/http"
)

func GenerateAPIURL() *http.Response {
	resp, _ := http.Get("https://opentdb.com/api.php?amount=10")

	return resp
}
