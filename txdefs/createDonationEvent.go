package txdefs

import (
	"encoding/json"
	"fmt"

	// "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
	"github.com/hyperledger/fabric-protos-go/msp"
)

//  Tasaction Definition
//  create the Donation Event
//  CreateDonationEvent defines a transaction for creating new donation events.
var CreateDonationEvent = tx.Transaction{
    Tag:         "createDonationEvent",
    Label:       "Create Donation Event",
    Description: "Create a new event in the donation system.",
    Method:      "POST",

    Callers: []accesscontrol.Caller{
        // {
		// 	// make this dynamic later 
        //     MSP: "org1MSP", // Only org1MSP can create the Event in the  org
        //     OU:  "admin",
        // },
        // {
        //     MSP: "org2MSP",
        //     OU:  "admin",
        // },
    },

    Args: []tx.Argument{
        {
            Tag:         "eventID",
            Label:       "Event ID",
            Description: "Unique identifier for the event.",
            DataType:    "string",
            Required:    true,
        },
        {
            Tag:         "recipient",
            Label:       "Recipient",
            Description: "Recipient creating the event.",
            DataType:    "string",
            Required:    true,
        },
        {
            Tag:         "eventDescription",
            Label:       "Event Description",
            Description: "Description of the event.",
            DataType:    "string",
            Required:    true,
        },
        {
            Tag:         "timestamp",
            Label:       "Timestamp",
            Description: "Event creation timestamp.",
            DataType:    "datetime",
            Required:    true,
        },
    },

    Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
        eventID := req["eventID"].(string)
        recipient := req["recipient"].(string)
        eventDescription := req["eventDescription"].(string)
        timestamp := req["timestamp"].(string)

        // Retrieve client's MSP ID. form  our api 
        // clientMSPID, err := stub.GetClientMSPID()
        // Retrieve client's MSP ID from the transaction creator's certificate
    creator, err := stub.Stub.GetCreator()
    
    if err != nil {
        return nil, errors.WrapError(err, "Failed to get transaction creator")
    }

    clientMSPID, err := extractMSPIDFromCreator(creator)
    if err != nil {
        return nil, errors.WrapError(err, "Failed to extract MSP ID from creator")
    }

        // Create the DonationEvent asset.
        donationEventMap := map[string]interface{}{
            "@assetType":      "donationEvent",
            "eventID":         eventID,
            "recipient":       recipient,
            "eventDescription": eventDescription,
            "organization":    clientMSPID,
            "timestamp":       timestamp,
        }

        donationEvent, err := assets.NewAsset(donationEventMap)
        if err != nil {
            return nil, errors.WrapError(err, "Failed to create a new donation event")
        }

        // Save the event on the blockchain.
        _, err = donationEvent.PutNew(stub)
        if err != nil {
            return nil, errors.WrapError(err, "Error saving event on blockchain")
        }

        eventJSON, err := json.Marshal(donationEvent)
        if err != nil {
            return nil, errors.WrapError(err, "Failed to encode event to JSON")
        }

        return eventJSON, nil
    },
}

// Helper function to extract organization from recipient.
func extractMSPIDFromCreator(creator []byte) (string, error) {
    sID := &msp.SerializedIdentity{}
    if err := proto.Unmarshal(creator, sID); err != nil {
        return "", fmt.Errorf("failed to unmarshal SerializedIdentity: %w", err)
    }
    return sID.Mspid, nil
}
