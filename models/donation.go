package models


// define the donation entities 
type Donation struct {
	ID         string `json:"id"`
	Donor      string `json:"donor"`
	Amount     int    `json:"amount"`
	Message    string `json:"message"`
	Recipient  string `json:"recipient"`
	Timestamp  string `json:"timestamp"`
}