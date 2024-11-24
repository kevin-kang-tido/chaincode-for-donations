package models

// Define the donation entities
type Donation struct {
	ID              string  `json:"id"`
	DonationEventID string  `json:"donationEventID"`
	Donor           string  `json:"donor"`
	Amount          float32 `json:"amount"`
	Message         string  `json:"message"`
	Recipient       string  `json:"recipient"`
	Timestamp       string  `json:"timestamp"`
}

// Define the DonationEvent
type DonationEvent struct {
	ID           string      `json:"id"`
	EventName    string      `json:"eventName"`
	Recipient    string      `json:"recipient"`
	Description  string      `json:"description"`
	Timestamp    string      `json:"timestamp"`
	Organization string      `json:"organization"`
	Donations    []Donation  `json:"donations"`
}

// Add a method to initialize Donations to an empty array if it's nil.
func (e *DonationEvent) InitializeDefaults() {
    if e.Donations == nil {
        e.Donations = []Donation{}
    }
}