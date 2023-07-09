package services

import (
	"testing"
)

var contract_base_Address = "0x3AFc6A5Bf7Ca257A1c002943386957c4E3d6D67e"
var contract_sale_Address = "0xbA5C8E2106fd59B76A93f4543195Ff9ADb735BD5"
var ownerAddress = "0cdA52d182Bc617B96caCE55839698780498e4b7"

const minterAdress = "0x6C7a0D6AF8e18cB0eE70D006dEF743457bf61efb"

var private_buyer_key = "3d820822284da292ccd381967ddfb4b4faedb17a5da4c5f37f7c824366d7c64f"
var buyer_adress = "0x39222FC95d11809590eB080306f281dbD4E90D4d"

func TestWeb3(t *testing.T) {
	//deployNewBaseContract()
	//deployNewSaleContract(contract_base_Address)
	//mintToken(2, 10, contract_base_Address)
	//getOwnerOfToken(55, contract_base_Address)
	getOwnedTokens(contract_base_Address, ownerAddress)
	//checkBalance(1, contract_base_Address, minterAdress)
	//checkBalance(2, contract_base_Address, minterAdress)
	//checkBalance(3, contract_base_Address, minterAdress)
	//checkBalance(4, contract_base_Address, minterAdress)
	//listTokenForSale(contract_base_Address, contract_sale_Address, big.NewInt(int64(4)), new(big.Int).Mul(big.NewInt(int64(1)), big.NewInt(1e18)), big.NewInt(int64(50)))
	//checkTokenListed(contract_sale_Address)
	//checkBalance(4, contract_base_Address, minterAdress)
	//checkBalance(4, contract_base_Address, buyer_adress)
	//buyTokenFromListing(private_buyer_key, contract_sale_Address, big.NewInt(int64(0)), new(big.Int).Mul(big.NewInt(int64(2)), big.NewInt(1e18)), big.NewInt(int64(5)))
	//checkBalance(4, contract_base_Address, minterAdress)
	//checkBalance(4, contract_base_Address, buyer_adress)
}
