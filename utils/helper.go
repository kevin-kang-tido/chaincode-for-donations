package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp"

	// "github.com/golang/protobuf/proto"
	"fmt"
)

// ToJSON converts an object to JSON string
func ToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// FromJSON converts JSON  to  object
func FromJSON(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}


func GetCreatorMSPID(ctx contractapi.TransactionContextInterface) (string, error) {
    creator, err := ctx.GetStub().GetCreator()
    if err != nil {
        return "", fmt.Errorf("error getting creator: %v", err)
    }

    sID := &msp.SerializedIdentity{}
    if err := proto.Unmarshal(creator, sID); err != nil {
        return "", fmt.Errorf("error unmarshaling SerializedIdentity: %v", err)
    }

    return sID.Mspid, nil
}

func IsValidRecipientForMSP(recipient, mspID string) bool {
    // Add logic to check recipient's organization

    return true // Replace with actual validation
}

// func ToJSON(data interface{}) ([]byte, error) {
//     return json.Marshal(data)
// }
// generate the The uniques id 
func GenereteHashID(donor,donationEventID string) string{
    data := donor + donationEventID 
    hash := sha512.Sum384([]byte(data))
    return hex.EncodeToString(hash[:])
}