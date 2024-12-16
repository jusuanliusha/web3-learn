package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	///Users/zhangwei/Library/Ethereum/geth.ipc
	//"http://localhost:8545"
	//"https://cloudflare-eth.com"
	client, err := ethclient.Dial("/Users/zhangwei/Library/Ethereum/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
