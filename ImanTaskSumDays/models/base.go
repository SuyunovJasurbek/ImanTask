package models

// error struct
type Error struct {
	Error string `json:"error"`
}

// message struct
type Message struct {
	Token string `json:"token"`
}

// days struct
type DaysMessage struct {
	Days int `json:"days"`
}