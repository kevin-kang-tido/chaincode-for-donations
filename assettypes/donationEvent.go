package assettypes


import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// DonationEvent  Define the Donation asset for the event in donation systems
// define the donation type 
var DonationEvent = assets.AssetType{
	Tag:         "donationEvent",
    Label:       "Donation Event",
    Description: "- The pupose of the Donation Event is Educate Poor young Genretion",
    
	// Props of the donation event 
	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "eventID",
			Label:    "Event ID",
			DataType: "string",
			// Only org1MSP can write
			Writers:  []string{"org1MSP"},  //  make this dynamic later 
		},
		{
			Tag:      "eventName",
			Label:    "Event Name",
			DataType: "string",
			Required: true,
			Writers:  []string{"org1MSP"}, // Writing restricted to org1MSP      // Readable by all organizations
		},
		{
			Tag:      "recipient",
			Label:    "Recipient Organization",
			DataType: "string",
			Required: true,
			Writers:  []string{"org1MSP"}, // Writing restricted to org1MSP      // Readable by all organizations
		},
		{
			Tag:      "description",
			Label:    "Event Description",
			DataType: "string",
			Required: false,
			Writers:  []string{"org1MSP"},
		},
		{
			Tag:      "timestamp",
			Label:    "Timestamp",
			DataType: "datetime",
			Required: true,
			Writers:  []string{"org1MSP"},
		},
	},
	

}