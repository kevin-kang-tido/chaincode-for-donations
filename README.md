
```
### List all the function that's chainecode handle Admin 

1.GetallDonation 
2.QueryDonationByDonorName // get by name 
3.GetDonationHistory  // check the history of the donation 
4.DonationExists // Checks specific donation record exists on the ledger

### List all the function that's chainecode handle Peer  

1.GetallDonation // all donation make by peer (client)
2.QueryDonationByDonorName // get by name 
3.GetDonationHistoryByName  // check the history of the donation my peers 
4.CreateDonation  // create donation peers 

```


```
### go start to project with golang 

go mod init your-project-name
go mod tidy

```

```
peer lifecycle chaincode package chaincode-donation.tar.gz --path ./chaincode-donation --label chaincode-donation_v1
peer lifecycle chaincode install chaincode-donation.tar.gz

```
