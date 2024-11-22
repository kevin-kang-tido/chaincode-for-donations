package models


// define the donation entities
type Donation struct {
	ID         string `json:"id"`
	DonationEventID string `json:donationEventID`
	Donor      string `json:"donor"`
	Amount     float32 `json:"amount"`
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