package web3

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
)

func PrintTest() {
	fmt.Println("Hello, World!")
}

// Prochaine tentative :https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9
// Exemple adapté de : https://github.com/chenzhijie/go-web3/blob/master/examples/contract/erc20.go et de https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings
//func deployNewContract() {
//
//	abiStr := `[{"inputs": [{"internalType": "uint256","name": "num","type": "uint256"},{"internalType": "uint256","name": "price","type": "uint256"},{"internalType": "uint256","name": "divi","type": "uint256"},{"internalType": "uint256","name": "initial","type": "uint256"}],"name": "store","outputs": [],"stateMutability": "nonpayable","type": "function"},{"inputs": [],"name": "retrieve","outputs": [{"internalType": "uint256","name": "","type": "uint256"},{"internalType": "uint256","name": "","type": "uint256"},{"internalType": "uint256","name": "","type": "uint256"},{"internalType": "uint256","name": "","type": "uint256"}],"stateMutability": "view","type": "function"}]`
//	// change to your rpc provider (this one is for Ganache)
//	rpcProvider := "HTTP://127.0.0.1:7545"
//	// set default account by private key
//	privateKey1 := "b4e49fdbfd4bf7989d121b9423cf371dfdc7d383b5cfdae98b27f9153de75f63"
//	bytecode := "0x" + "608060405234801561001057600080fd5b506101e6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b578063f58a28451461005c575b600080fd5b610043610078565b60405161005394939291906100d3565b60405180910390f35b61007660048036038101906100719190610149565b610098565b005b600080600080600054600154600254600354935093509350935090919293565b8360008190555082600181905550816002819055508060038190555050505050565b6000819050919050565b6100cd816100ba565b82525050565b60006080820190506100e860008301876100c4565b6100f560208301866100c4565b61010260408301856100c4565b61010f60608301846100c4565b95945050505050565b600080fd5b610126816100ba565b811461013157600080fd5b50565b6000813590506101438161011d565b92915050565b6000806000806080858703121561016357610162610118565b5b600061017187828801610134565b945050602061018287828801610134565b935050604061019387828801610134565b92505060606101a487828801610134565b9150509295919450925056fea26469706673582212207fa7ea02e09f6434989a440d717016a1778a5128a9adacd0b5df40e5890bd78a64736f6c63430008120033"
//
//	client, err := ethclient.Dial(rpcProvider)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA(privateKey1)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("error casting public key to ECDSA")
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
//	auth.Value = big.NewInt(0)     // in wei
//	auth.GasLimit = uint64(300000) // in units
//	auth.GasPrice = gasPrice
//
//	input := "1.0"
//	address, tx, instance, err := store.DeployStore(auth, client, input)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	chainId := int64(1337)
//	if err := web3.Eth.SetAccount(privateKey); err != nil {
//		panic(err)
//	}
//	web3.Eth.SetChainId(chainId)
//	tokenAddr := "0x623639bA2c3ffc7b7d8879e8F3e4CFC905f7A3C7"
//	contract, err := web3.Eth.NewContract(abiStr, tokenAddr)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Contract address: ", contract.Address())
//
//	// We need to change that to get the 1st element of retrieve
//	totalSupply, err := contract.Call("retrieve")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Total supply %v\n", totalSupply)

//data, _ := contract.EncodeABI("balanceOf", web3.Eth.Address())
//fmt.Printf("%x\n", data)
//
//balance, err := contract.Call("balanceOf", web3.Eth.Address())
//if err != nil {
//	panic(err)
//}
//fmt.Printf("Balance of %v is %v\n", web3.Eth.Address(), balance)
//
//allowance, err := contract.Call("allowance", web3.Eth.Address(), "0x0000000000000000000000000000000000000002")
//if err != nil {
//	panic(err)
//}
//
//fmt.Printf("Allowance is %v\n", allowance)
//approveInputData, err := contract.Methods("approve").Inputs.Pack("0x0000000000000000000000000000000000000002", web3.Utils.ToWei("0.2"))
//if err != nil {
//	panic(err)
//}
//// fmt.Println(approveInputData)
//
//tokenAddress := common.HexToAddress(tokenAddr)
//
//call := &types.CallMsg{
//	From: web3.Eth.Address(),
//	To:   tokenAddress,
//	Data: approveInputData,
//	Gas:  types.NewCallMsgBigInt(big.NewInt(types.MAX_GAS_LIMIT)),
//}
//// fmt.Printf("call %v\n", call)
//gasLimit, err := web3.Eth.EstimateGas(call)
//if err != nil {
//	panic(err)
//}
//fmt.Printf("Estimate gas limit %v\n", gasLimit)
//nonce, err := web3.Eth.GetNonce(web3.Eth.Address(), nil)
//if err != nil {
//	panic(err)
//}
//txHash, err := web3.Eth.SyncSendRawTransaction(
//	common.HexToAddress(tokenAddr),
//	big.NewInt(0),
//	nonce,
//	gasLimit,
//	web3.Utils.ToGWei(1),
//	approveInputData,
//)
//if err != nil {
//	panic(err)
//}
//fmt.Printf("Send approve tx hash %v\n", txHash)
//}

//func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {
//
//	privateKey, err := crypto.HexToECDSA(accountAddress)
//	if err != nil {
//		panic(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		panic("invalid key")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//
//	//fetch the last use nonce of account
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("nounce=", nonce)
//	chainID, err := client.ChainID(context.Background())
//	if err != nil {
//		panic(err)
//	}
//
//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err != nil {
//		panic(err)
//	}
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)      // in wei
//	auth.GasLimit = uint64(6000000) // in units
//	auth.GasPrice = gasPrice
//
//	return auth
//}
//
//func deployNewContract() {
//	// address of etherum env
//	client, err := ethclient.Dial("http://127.0.0.1:7545")
//	if err != nil {
//		panic(err)
//	}
//
//	// create auth and transaction package for deploying smart contract
//	auth := getAccountAuth(client, "36cdaf7554327b0ae13eb64967559137f1183a901cdea654acbd736a6317c8ac")
//
//	//deploying smart contract
//	deployedContractAddress, tx, instance, err := api.DeployApi(auth, client) //api is redirected from api directory from our contract go file
//	if err != nil {
//		panic(err)
//	}
//	print(tx)
//	print(instance)
//
//	fmt.Println(deployedContractAddress.Hex()) // print deployed contract address
//}

func deployNewContract() {
	// Connect to Ganache
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to Ganache: %v", err)
	}

	// Read private key
	privateKey, err := crypto.HexToECDSA("36cdaf7554327b0ae13eb64967559137f1183a901cdea654acbd736a6317c8ac")
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Derive the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to derive ECDSA public key")
	}

	// Get the address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(7721975) // in units
	auth.GasPrice = gasPrice

	// Read the contract bytecode
	bytecode, err := ioutil.ReadFile("../build/contracts_musiChain_sol_MusiChain.bin")
	if err != nil {
		log.Fatal(err)
	}

	// Read and parse the ABI from file
	abiFile, err := ioutil.ReadFile("../build/contracts_musiChain_sol_MusiChain.abi")
	if err != nil {
		log.Fatal(err)
	}

	var parsedABI abi.ABI
	if err := json.Unmarshal(abiFile, &parsedABI); err != nil {
		log.Fatal(err)
	}

	input := append([]byte(nil), bytecode...) // This could be contract constructor parameters, if necessary

	// Deploy the contract
	tx := types.NewContractCreation(nonce, big.NewInt(0), uint64(300000), gasPrice, input)
	signer := types.NewEIP155Signer(big.NewInt(1337)) // assuming ChainID is 1337
	signedTx, _ := types.SignTx(tx, signer, privateKey)

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx sent: %s", signedTx.Hash().Hex()) // Output the transaction hash
}
