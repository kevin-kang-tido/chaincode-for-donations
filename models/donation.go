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

// Define the DonatinEvent 
type DonationEvent struct {
    ID           string `json:"id"`
	EventName    string  `json:eventName`
    Recipient    string `json:"recipient"`
    Description  string `json:"description"`
    Timestamp    string `json:"timestamp"`
    Organization string `json:"organization"`
}