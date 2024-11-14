package main

import (
	"chaincode-donation/contract"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	donationContract := new(contract.DonationContract)
	chaincode, err := contractapi.NewChaincode(donationContract)
	// if create the donation got the error 
	if err != nil {
		log.Panicf("Error creating donation chaincode: %v", err)
	}
    // if chaincode start the errro 
	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting donation chaincode: %v", err)
	}
}


