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
	"musichain/pkg/services/abigen/sale"

	baseContractWrapper "musichain/pkg/services/abigen/Base"
)

const privateKey = "567a218541559b46d9570da9cc75e48d195f7181ea231e85a22643e051f4f1c3"

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

func deployNewSaleContract(contract_Address string) {
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

	contract_Address_hex := common.HexToAddress(contract_Address)
	address, tx, instance, err := sale.DeploySale(auth, client, contract_Address_hex)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func mintToken(tokenID int64, amountt int64, minterAdress string) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// The address of the ERC1155 contract deployed
	contractAddress := common.HexToAddress(minterAdress)

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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
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

func listTokenForSale(baseContractAdress string, marketplaceAddress string, tokenId, price, amount *big.Int) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Configure the transactor
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	baseContractAdress_hex := common.HexToAddress(baseContractAdress)
	baseContract, err := baseContractWrapper.NewBase(baseContractAdress_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new instance of the marketplace contract
	marketplaceAddress_hex := common.HexToAddress(marketplaceAddress)
	_, err = baseContract.SetApprovalForAll(auth, marketplaceAddress_hex, true)
	if err != nil {
		log.Fatal(err)
	}
	marketplace, err := sale.NewSale(marketplaceAddress_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	// List a token for sale
	_, err = marketplace.ListToken(auth, tokenId, price, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d token %s listed at price %d\n", amount, tokenId, price)
}

func checkTokenListed(marketplaceAddress string) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Create a new instance of the marketplace contract
	marketplaceAddressHex := common.HexToAddress(marketplaceAddress)
	marketplace, err := sale.NewSale(marketplaceAddressHex, client)
	if err != nil {
		log.Fatal(err)
	}

	// Call the getListings function
	listings, err := marketplace.GetListings(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print out the listings
	for i, listing := range listings {
		fmt.Printf("Listing %d: seller = %s, tokenId = %d, price = %d, amount = %d\n",
			i, listing.Seller.Hex(), listing.TokenId, listing.Price, listing.Amount)
	}
}

func buyTokenFromListing(private_buyer_key string, marketplaceAddress string, listingId, price, amount *big.Int) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Configure the transactor
	privateKey, err := crypto.HexToECDSA(private_buyer_key)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Create a new instance of the marketplace contract
	marketplaceAddress_hex := common.HexToAddress(marketplaceAddress)
	marketplace, err := sale.NewSale(marketplaceAddress_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	// Buy a token
	auth.Value = price
	_, err = marketplace.BuyToken(auth, listingId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d token %d bought \n", amount, listingId)
}
