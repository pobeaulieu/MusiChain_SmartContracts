package services

import "testing"

var contract_Address = "0xEd6303f10913eE5a3ef16A351CD0a3a145Ee4437"

const recipientAddress = "0x78870f63471f3152b78dd57A01fE5E9A494D898A\nBALANCE\n"

func TestWeb3(t *testing.T) {
	deployNewTokenContract()
	//mintToken(1, 100, contract_Address, recipientAddress)
	//checkBalance(1, contract_Address, recipientAddress)
	//deployNewSaleContract(contract_Address)
	//buyToken(1, 10)
	//checkBalance(1)
}
