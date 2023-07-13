package services

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	baseContractWrapper "musichain/pkg/services/abigen/Base"
	"strings"
)

func sendSignedTx(signedTxHex string) (string, error) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	// Convert the hex string to bytes
	signedTxBytes, err := hex.DecodeString(signedTxHex)
	if err != nil {
		return "", fmt.Errorf("Failed to decode hex: %v", err)
	}

	// Convert the bytes to a transaction type
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(signedTxBytes, tx); err != nil {
		return "", fmt.Errorf("Failed to decode transaction: %v", err)
	}

	// Broadcast the transaction
	if err := client.SendTransaction(context.Background(), tx); err != nil {
		return "", fmt.Errorf("Failed to broadcast transaction: %v", err)
	}

	return tx.Hash().Hex(), nil
}

func prepareMintTransaction(tokenID int64, amountt int64, contractAddress string, minterAddress string) ([]byte, error) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(minterAddress))
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// Prepare the data to be sent with the transaction
	tokenId := big.NewInt(tokenID)
	amount := big.NewInt(amountt)
	data := []byte{}

	// Parse the ABI string to an ABI instance
	parsedABI, err := abi.JSON(strings.NewReader(baseContractWrapper.BaseMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse contract ABI: %v", err)
	}

	// Now you can call the Pack method
	input, err := parsedABI.Pack("mint", tokenId, amount, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to pack data for mint: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, common.HexToAddress(minterAddress), big.NewInt(0), 3000000, gasPrice, input)

	// Encode the transaction to RLP
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode transaction: %v", err)
	}

	return rawTxBytes, nil
}
