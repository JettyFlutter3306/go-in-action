package net

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	address := "http://5lmh.com"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	response, err := client.Get(address)

	if err != nil {
		fmt.Println("Bugs~")
	}

	fmt.Println(response.Body)

	request, err := http.NewRequest("GET", address, nil)

	request.Header.Add("If-None-Match", `W/"wyzzy"`)

	response1, err := client.Do(request)

	if err != nil {
		fmt.Println("bugs")
	}

	fmt.Println(response1.Body)

}
