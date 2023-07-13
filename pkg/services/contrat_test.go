package services

import (
	"testing"
)

const privateKey = "383ac5298026b523038e452fc357a3e99341aceb59f9b260b5af2193a36fb8fc"

var contract_base_Address = getContractBaseAddress()
var contract_sale_Address = getContractSaleAddress()

var ownerAddress = "0x03e1ad05936c71c4DBeAa9b4da89000169703AB5"

const minterAdress = "0x03e1ad05936c71c4DBeAa9b4da89000169703AB5"

var private_buyer_key = "95049903b2fd23535299024c9ba71cc0edc3411710fb5ac38626b77af4ca3e02"
var buyer_adress = "0xaB2A69Cfa03A27Aea2D99E2FB79E0d7351F0103c"

func TestWeb3(t *testing.T) {
	//deployNewBaseContract(privateKey)
	//deployNewSaleContract(privateKey, contract_base_Address)
	//mintToken(privateKey, 1, 10, contract_base_Address)
	//getOwnerOfToken(55, contract_base_Address)
	//getOwnedTokens(contract_base_Address, ownerAddress)
	//checkBalance(1, contract_base_Address, minterAdress)
	//listTokenForSale(privateKey, contract_base_Address, contract_sale_Address, big.NewInt(int64(1)), new(big.Int).Mul(big.NewInt(int64(1)), big.NewInt(1e18)), big.NewInt(int64(50)))
	//checkTokenListed(contract_sale_Address)
	//checkBalance(1, contract_base_Address, minterAdress)
	//checkBalance(1, contract_base_Address, buyer_adress)
	//buyTokenFromListing(private_buyer_key, contract_sale_Address, big.NewInt(int64(0)), new(big.Int).Mul(big.NewInt(int64(1)), big.NewInt(1e18)), big.NewInt(int64(5)))
	//checkBalance(1, contract_base_Address, minterAdress)
	//checkBalance(1, contract_base_Address, buyer_adress)
}
