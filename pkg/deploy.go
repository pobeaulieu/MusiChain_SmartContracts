package pkg

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
	metaWrapper "musichain/pkg/abigen/metadata"
	"musichain/pkg/abigen/sale"
	"os"

	baseContractWrapper "musichain/pkg/abigen/Base"
)

func Deploy(privateKey string) {
	metaAddress := deployNewMetaDataContract(privateKey)
	baseAddress := deployNewBaseContract(privateKey, metaAddress)
	saleAddress := deployNewSaleContract(privateKey, baseAddress)

	f, err := os.Create(".env")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString("REACT_APP_METADATA_ADDRESS = " + metaAddress + "\n")
	_, err = f.WriteString("REACT_APP_BASE_ADDRESS = " + baseAddress + "\n")
	_, err = f.WriteString("REACT_APP_SALE_ADDRESS = " + saleAddress + "\n")

	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	f.Sync()
}

func deployNewMetaDataContract(privateKeyString string) string {
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

	address, tx, instance, err := metaWrapper.DeployMetadata(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	addressHex := address.Hex()

	_ = instance

	fmt.Println(addressHex)
	fmt.Println(tx.Hash().Hex())

	return addressHex
}

func deployNewBaseContract(privateKeyString string, contractAdress string) string {
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

	contract_Address_hex := common.HexToAddress(contractAdress)
	address, tx, instance, err := baseContractWrapper.DeployBase(auth, client, contract_Address_hex)
	if err != nil {
		log.Fatal(err)
	}

	addressHex := address.Hex()

	fmt.Println(addressHex)
	fmt.Println(tx.Hash().Hex())

	_ = instance

	return addressHex
}

func deployNewSaleContract(privateKeyString string, contract_Address string) string {
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

	_ = instance

	fmt.Println(addressHex)
	fmt.Println(tx.Hash().Hex())

	return addressHex
}
