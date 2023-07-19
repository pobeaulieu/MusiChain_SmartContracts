package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"musichain/pkg/services/abigen/sale"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	baseContractWrapper "musichain/pkg/services/abigen/Base"
)

func deployNewBaseContract(privateKeyString string) {
	// Connect to Ganache
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to Ganache: %v", err)
	}

	// Read private key
	privateKey, err := crypto.HexToECDSA(privateKeyString)
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := baseContractWrapper.DeployBase(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	addressHex := address.Hex()

	fmt.Println(addressHex)
	fmt.Println(tx.Hash().Hex())

	f, err := os.Create("base_address.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(addressHex + "\n")
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	f.Sync()

	_ = instance
}

func deployNewSaleContract(privateKeyString string, contract_Address string) {
	// Connect to Ganache
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to Ganache: %v", err)
	}

	// Read private key
	privateKey, err := crypto.HexToECDSA(privateKeyString)
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	contract_Address_hex := common.HexToAddress(contract_Address)
	address, tx, instance, err := sale.DeploySale(auth, client, contract_Address_hex)
	if err != nil {
		log.Fatal(err)
	}

	addressHex := address.Hex()

	fmt.Println(addressHex)
	fmt.Println(tx.Hash().Hex())

	f, err := os.Create("sale_address.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(addressHex + "\n")
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	f.Sync()

	_ = instance
}

func mintToken(privateKeyString string, tokenID int64, amountt int64, minterAdress string) {
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
	privateKey, err := crypto.HexToECDSA(privateKeyString)
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

	amount := big.NewInt(amountt)
	data := []byte{}

	// Call the mint function
	tx, err := contract.Mint(auth, "exemple", amount, data)
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

func listTokenForSale(privateKeyString string, baseContractAdress string, marketplaceAddress string, tokenId, price, amount *big.Int) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Configure the transactor
	privateKey, err := crypto.HexToECDSA(privateKeyString)
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

func getOwnerOfToken(tokenID int64, contract_Address string) {
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

	// Get the owner of the token
	owner, err := contract.GetOwnerOfToken(&bind.CallOpts{}, big.NewInt(tokenID))
	if err != nil {
		log.Fatalf("Failed to retrieve the owner of the token: %v", err)
	}

	// Log the owner
	fmt.Printf("Owner of token %d: %s\n", tokenID, owner.Hex())
}

func getOwnedTokens(contractAddress, ownerAddress string) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// The address of the deployed contract
	contractAddr := common.HexToAddress(contractAddress)

	// Create a new instance of the contract
	contract, err := baseContractWrapper.NewBase(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to create a new instance of the contract: %v", err)
	}

	// The address of the owner
	ownerAddr := common.HexToAddress(ownerAddress)

	// Call the getOwnedTokens function in the contract
	tokenIds, err := contract.GetOwnedTokens(nil, ownerAddr)
	if err != nil {
		log.Fatalf("Failed to retrieve the owned tokens: %v", err)
	}

	// Log the owned token IDs
	fmt.Printf("Owned tokens for address %s:\n", ownerAddress)
	for _, tokenId := range tokenIds {
		fmt.Printf("- Token ID: %s\n", tokenId.String())
	}
}

func getContractBaseAddress() string {
	bytes, err := ioutil.ReadFile("base_address.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return strings.TrimSpace(string(bytes))
}

func getContractSaleAddress() string {
	bytes, err := ioutil.ReadFile("sale_address.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return strings.TrimSpace(string(bytes))
}
