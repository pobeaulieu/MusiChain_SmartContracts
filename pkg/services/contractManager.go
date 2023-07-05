package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"

	baseContractWrapper "musichain/pkg/services/abigen/Base"
)

const privateKey = "00756bd467db1d44d1c896101f9a292b075a337b0acbb94820f28fcca57d9d55"

func deployNewBaseContract() {
	// Connect to Ganache
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to Ganache: %v", err)
	}

	// Read private key
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create a new authorized transactor
	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := baseContractWrapper.DeployBase(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func mintToken(tokenID int64, amountt int64, contract_Address string) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// The address of the ERC1155 contract deployed
	contractAddress := common.HexToAddress(contract_Address)

	// Load the contract (make sure to use your actual ABI)
	contract, err := baseContractWrapper.NewBase(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	// Configure the transactor
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Ganache default chain ID
	if err != nil {
		log.Fatal(err)
	}

	// Set gas price and limit
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(20000000000)

	tokenId := big.NewInt(tokenID)
	amount := big.NewInt(amountt)
	data := []byte{}

	// Call the mint function
	tx, err := contract.Mint(auth, tokenId, amount, data)
	if err != nil {
		log.Fatalf("Failed to mint token: %v", err)
	}

	fmt.Printf("Minted tokens! Transaction hash: %s\n", tx.Hash().Hex())
}

func checkBalance(tokenID int64, contract_Address string, recipient_Address string) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// The address of the deployed contract
	contractAddress := common.HexToAddress(contract_Address)

	// Create a new instance of the contract
	contract, err := baseContractWrapper.NewBase(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create a new instance of the contract: %v", err)
	}

	// Specify the recipient's address and the token ID
	recipientAddress := common.HexToAddress(recipient_Address)
	tokenId := big.NewInt(tokenID) // specify the correct token ID

	// Query the balance of the specific token for the recipient address
	balance, err := contract.BalanceOf(&bind.CallOpts{}, recipientAddress, tokenId)
	if err != nil {
		log.Fatalf("Failed to retrieve token balance: %v", err)
	}

	// Log the balance
	fmt.Printf("Balance of token %d for address %s: %d\n", tokenId, recipientAddress.Hex(), balance)
}

//
//func deployNewSaleContract(contract_Address string) {
//	// Connect to Ganache
//	client, err := ethclient.Dial("http://localhost:7545")
//	if err != nil {
//		log.Fatalf("Failed to connect to Ganache: %v", err)
//	}
//
//	// Read private key
//	privateKey, err := crypto.HexToECDSA(privateKey)
//	if err != nil {
//		log.Fatalf("Failed to parse private key: %v", err)
//	}
//
//	auth := bind.NewKeyedTransactor(privateKey)
//	auth.Value = big.NewInt(0)            // in wei
//	auth.GasLimit = uint64(6721975)       // reduce this to a reasonable amount
//	auth.GasPrice = big.NewInt(863252572) // 20 Gwei, as provided by Ganache
//	// Read the contract bytecode
//	bytecode, err := ioutil.ReadFile("build/contracts_tokenSale_sol_TokenSale.bin")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Read and parse the ABI from file
//	abiFile, err := ioutil.ReadFile("build/contracts_tokenSale_sol_TokenSale.abi")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var parsedABI abi.ABI
//	if err := json.Unmarshal(abiFile, &parsedABI); err != nil {
//		log.Fatal(err)
//	}
//
//	input := append([]byte(nil), bytecode...) // This could be contract constructor parameters, if necessary
//
//	tokenContractAddress := common.HexToAddress(contract_Address)
//	tokenPrice := big.NewInt(10000000000000000) // Example token price
//
//	address, tx, _, err := bind.DeployContract(auth, parsedABI, input, client, tokenContractAddress, tokenPrice)
//	if err != nil {
//		log.Fatalf("Failed to deploy new token contract: %v", err)
//	}
//
//	fmt.Printf("Contract deployed to address: %s\n", address.Hex())
//	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
//}
//
//func buyToken(tokenID int64, number int64) {
//	client, err := ethclient.Dial("http://localhost:7545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA(privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	auth := bind.NewKeyedTransactor(privateKey)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)
//	auth.GasLimit = uint64(300000)
//	auth.GasPrice = gasPrice
//
//	// Token sale contract address
//	tokenSaleAddress := common.HexToAddress("your-token-sale-contract-address")
//	instance, err := sale.NewTokenSale(tokenSaleAddress, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Replace with your token ID, number of tokens and any necessary data
//	tokenId := big.NewInt(tokenID)
//	numberOfTokens := big.NewInt(number)
//	data := []byte{} // Data can be empty or any necessary data according to your contract
//
//	tx, err := instance.BuyTokens(auth, tokenId, numberOfTokens, data)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Purchased tokens: tx=%s\n", tx.Hash().Hex())
//}
