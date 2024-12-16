package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0x1Da7CbE2c41ee853ef9dAF9E2a97Cf4899D151DA")
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())
}
