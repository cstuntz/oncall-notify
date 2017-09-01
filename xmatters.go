package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	config "github.com/cstuntz/oncall-notify/config"
)

// GetXmatters Calls XMatters API to get the current list of on-call users
func GetXmatters(url string, auth string) config.XMattersRes {

	// Make new client, set custom timeout
	xClient := http.Client{
		Timeout: time.Second * 5,
	}
	// Create the request with our custom auth header
	req, err := http.NewRequest(http.MethodGet, url, nil)
	failOnErr(err, "Failed to initialize request")
	req.Header.Set("Authorization", auth)

	// Perform the http call with our custom info
	res, getErr := xClient.Do(req)
	failOnErr(getErr, "Failed to get XMatters URL")
	defer res.Body.Close()

	// Read the body of the response into a variable
	body, readErr := ioutil.ReadAll(res.Body)
	failOnErr(readErr, "Failed to read XMatters response body")

	// Convert the JSON response into a struct
	var xRes config.XMattersRes
	jsonErr := json.Unmarshal(body, &xRes)
	failOnErr(jsonErr, "Failed to unmarshal json response")

	return xRes
}
