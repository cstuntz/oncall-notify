package config

// XMattersRes Struct to hold the response from XMatters
type XMattersRes struct {
	Count int `json:"count"`
	Data  []struct {
		End   string `json:"end"`
		Group struct {
			ID    string `json:"id"`
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
			RecipientType string `json:"recipientType"`
			TargetName    string `json:"targetName"`
		} `json:"group"`
		Members struct {
			Count int `json:"count"`
			Data  []struct {
				Delay          int    `json:"delay"`
				EscalationType string `json:"escalationType"`
				Member         struct {
					ExternallyOwned bool   `json:"externallyOwned"`
					FirstName       string `json:"firstName"`
					ID              string `json:"id"`
					Language        string `json:"language"`
					LastName        string `json:"lastName"`
					Links           struct {
						Self string `json:"self"`
					} `json:"links"`
					PhoneLogin    string `json:"phoneLogin"`
					RecipientType string `json:"recipientType"`
					Site          struct {
						ID    string `json:"id"`
						Links struct {
							Self string `json:"self"`
						} `json:"links"`
						Name string `json:"name"`
					} `json:"site"`
					Status     string `json:"status"`
					TargetName string `json:"targetName"`
					Timezone   string `json:"timezone"`
					WebLogin   string `json:"webLogin"`
				} `json:"member"`
				Position int `json:"position"`
			} `json:"data"`
			Links struct {
				Next string `json:"next"`
				Self string `json:"self"`
			} `json:"links"`
			Total int `json:"total"`
		} `json:"members"`
		Shift struct {
			ID    string `json:"id"`
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
		} `json:"shift"`
		Start string `json:"start"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Total int `json:"total"`
}

// UserList Struct to hold the list of users that want to be alerted
type UserList struct {
	Users []User `json:"users"`
}

// User Single user's information for xmatters and slack
type User struct {
	XMattersID      string `json:"xMattersID"`
	XMattersGroupID string `json:"xMattersGroupID"`
	SlackID         string `json:"slackName"`
}
