package main

import (
	"net/http"
	"strings"
)

//SendMessage Send a message to the slack webhook
func SendMessage(msg string, url string) error {

	// Send the message to slack
	resp, err := http.Post(url, "application/json", strings.NewReader(msg))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
