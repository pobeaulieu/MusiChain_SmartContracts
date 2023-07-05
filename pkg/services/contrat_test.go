package services

import (
	"testing"
)

var contract_base_Address = "0x7dd3E696a820F6a6b439e8481FAc12eAd6Cedd5D"
var contract_sale_Address = "0x02479718614701B3946EAFeBd1991149c3cc4818"

const minterAdress = "0x0db1AbC5c61f709e64C8b9bfeF44C4e6e2784301"

var private_buyer_key = "3d820822284da292ccd381967ddfb4b4faedb17a5da4c5f37f7c824366d7c64f"
var buyer_adress = "0x39222FC95d11809590eB080306f281dbD4E90D4d"

func TestWeb3(t *testing.T) {
	//deployNewBaseContract()
	//deployNewSaleContract(contract_base_Address)
	//mintToken(1, 10, contract_base_Address)
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
