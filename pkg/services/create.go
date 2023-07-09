package services

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func CreateToken(client *rpc.Client, contractAddress common.Address, tokenID *big.Int, recipientAddress common.Address) error {
	callData := []byte(fmt.Sprintf("0x6352211e000000000000000000000000%s%s", tokenID.Text(16), recipientAddress.Hex()))

	var result string
	err := client.Call(&result, "eth_sendTransaction", []interface{}{
		map[string]interface{}{
			"to":   contractAddress,
			"data": callData,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	client, err := rpc.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x2c73ea7405a6Fc96cc9e9c410050cFFf22B425B2")
	recipientAddress := common.HexToAddress("0x95b6db7a509453b1b60b6471e6ba237a9425c1dfa31a389e8524de4bb9f3a7da")
	tokenID := big.NewInt(123)

	err = CreateToken(client, contractAddress, tokenID, recipientAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Token created successfully!")
}
