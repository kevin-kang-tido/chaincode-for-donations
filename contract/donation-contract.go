package contract

import (
	"chaincode-donation/models"
	"chaincode-donation/utils"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DonationContract struct {
	contractapi.Contract
}

// InitLedger adds a base set of donations to the ledger
// InitLedger just for testing 
func (dc *DonationContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	donations := []models.Donation{
		{ID: "donation1", Donor: "Alice", Amount: 100, Message: "For education", Recipient: "CharityA", Timestamp: "2023-10-10T10:00:00Z"},
		{ID: "donation2", Donor: "Bob", Amount: 150, Message: "Medical aid", Recipient: "CharityB", Timestamp: "2023-10-11T11:30:00Z"},
		{ID: "donation3", Donor: "Charlie", Amount: 200, Message: "Food and shelter", Recipient: "CharityC", Timestamp: "2023-10-12T14:15:00Z"},
		{ID: "donation4", Donor: "Dana", Amount: 250, Message: "Water sanitation", Recipient: "CharityD", Timestamp: "2023-10-13T16:45:00Z"},
		{ID: "donation5", Donor: "Eve", Amount: 300, Message: "Emergency relief", Recipient: "CharityE", Timestamp: "2023-10-14T18:00:00Z"},
	}

	for _, donation := range donations {
		// donationJSON, err := json.Marshal(donation)
	    donationJSON, err := utils.ToJSON(donation)
		if err != nil {
			return fmt.Errorf("failed to marshal donation: %v", err)
		}

		err = ctx.GetStub().PutState(donation.ID, donationJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state: %v", err)
		}
	}

	return nil
}


// CreateDonation registers a new donation in CouchDB
func (dc *DonationContract) CreateDonation(ctx contractapi.TransactionContextInterface, id, donor string, amount int, message, recipient, timestamp string) error {
	exists, err := dc.DonationExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Donation with ID %s already exists", id)
	}

	donation := models.Donation{
		ID:         id,
		Donor:      donor,
		Amount:     amount,
		Message:    message,
		Recipient:  recipient,
		Timestamp:  timestamp,
	}

	donationJSON, err := json.Marshal(donation)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, donationJSON)
}

// ReadDonation retrieves a donation from CouchDB by ID
func (dc *DonationContract) ReadDonation(ctx contractapi.TransactionContextInterface, id string) (*models.Donation, error) {
	donationJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read donation: %v", err)
	}
	if donationJSON == nil {
		return nil, fmt.Errorf("donation %s does not exist", id)
	}

	var donation models.Donation
	err = json.Unmarshal(donationJSON, &donation)
	if err != nil {
		return nil, err
	}

	return &donation, nil
}

// UpdateDonation allows updating existing donation details
func (dc *DonationContract) UpdateDonation(ctx contractapi.TransactionContextInterface, id, donor string, amount int, message, recipient, timestamp string) error {
	exists, err := dc.DonationExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("donation %s does not exist", id)
	}

	donation := models.Donation{
		ID:         id,
		Donor:      donor,
		Amount:     amount,
		Message:    message,
		Recipient:  recipient,
		Timestamp:  timestamp,
	}

	donationJSON, err := json.Marshal(donation)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, donationJSON)
}

// DeleteDonation removes a donation entry from CouchDB
func (dc *DonationContract) DeleteDonation(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := dc.DonationExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("donation %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// DonationExists checks if a donation exists in CouchDB by ID
func (dc *DonationContract) DonationExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	donationJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, err
	}
	return donationJSON != nil, nil
}
// get function to get all Donatin data : 
// GetAllDonations retrieves all donations from CouchDB
func (dc *DonationContract) GetAllDonations(ctx contractapi.TransactionContextInterface) ([]*models.Donation, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve donations: %v", err)
	}
	defer resultsIterator.Close()

	var donations []*models.Donation
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var donation models.Donation
		err = json.Unmarshal(queryResponse.Value, &donation)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal donation: %v", err)
		}

		donations = append(donations, &donation)
	}

	return donations, nil
}



