package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "jsls.work/ge/contracts/erc20"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383820),
		ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	LogApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)
	for _, vLog := range logs {
		fmt.Println("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Println("Log Index: %d\n", vLog.Index)
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			_, err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			transferEvent.From = common.HexToAddress(string(vLog.Topics[1].Hex()))
			transferEvent.To = common.HexToAddress(string(vLog.Topics[2].Hex()))

			fmt.Println("From: %s\n", transferEvent.From.Hex())
			fmt.Println("To: %s\n", transferEvent.To.Hex())
			fmt.Println("Tokens: %s\n", transferEvent.Tokens.String())

		case LogApprovalSigHash.Hex():
			fmt.Println("Log Name: Approval\n")
			var approvalEvent LogApproval
			_, err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			approvalEvent.TokenOwner = common.HexToAddress(string(vLog.Topics[1].Hex()))
			approvalEvent.Spender = common.HexToAddress(string(vLog.Topics[2].Hex()))

			fmt.Println("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Println("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Println("Tokens: %s\n", approvalEvent.Tokens.String())
		}
		fmt.Println("\n\n")
	}

}
