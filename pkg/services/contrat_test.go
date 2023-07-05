package services

import "testing"

var contract_Address = "0x6dc715370602Ef180Ee13b826225757e5F006e37"

const recipientAddress = "0xA55a12bcEd288BD428fF4E66e5c4EFF973D84ae2"

func TestWeb3(t *testing.T) {
	//deployNewBaseContract()
	mintToken(3, 10, contract_Address)
	checkBalance(1, contract_Address, recipientAddress)
	checkBalance(2, contract_Address, recipientAddress)
	checkBalance(3, contract_Address, recipientAddress)
	//deployNewSaleContract(contract_Address)
	//buyToken(1, 10)
	//checkBalance(1)
}
