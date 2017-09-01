package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// LoadUsers Reads in userlist.json, returns array of userids
func LoadUsers() UserList {
	// Read in the file
	body, err := ioutil.ReadFile("config/userslist.json")
	failOnErr(err, "Failed to read config file")

	// Convert the file to a struct
	var userList UserList
	jsonErr := json.Unmarshal(body, &userList)
	failOnErr(jsonErr, "Failed to unmarshal config file into json struct")

	return userList
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}
