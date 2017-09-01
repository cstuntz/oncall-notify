package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/cstuntz/oncall-notify/config"
)

func main() {

	xMattersBaseURL := os.Getenv("XM_BASE_URL")
	xMattersAuth := os.Getenv("XM_AUTH")
	slackURL := os.Getenv("SLACK_URL")

	// Get the list of all users we care about
	userList := config.LoadUsers()

	for _, user := range userList.Users {
		// Create the url path with the user's group
		xMattersPath := fmt.Sprintf("/api/xm/1/on-call?groups=%s", user.XMattersGroupID)

		// Get the group on-call info from XMatters API
		url := xMattersBaseURL + xMattersPath
		xRes := GetXmatters(url, xMattersAuth)

		// If you're on call, send the Slack message
		if isOnCall(user.XMattersID, &xRes) {
			tempID := user.SlackID
			log.Printf("Notifying user: %s\n", tempID)
			// Format the message so it's in the JSON format Slack is expecting
			message := fmt.Sprintf("{\"text\": \"You're on call!\", \"channel\": \"%s\"}", tempID)
			err := SendMessage(message, slackURL)
			failOnErr(err, "Failed to send slack message")
		}
	}

}

func isOnCall(id string, x *config.XMattersRes) bool {

	// Does the current user ID equal the current on-call ID?
	if id == x.Data[0].Members.Data[0].Member.ID {
		return true
	}

	return false
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}
