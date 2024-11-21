package main

import (
	"chaincode-donation/contract"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	// create a instance of the donatin which get the chaincode struct from contract 
	donationContract := new(contract.DonationContract)
	// using contrapapi function to create new chaincode  
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


