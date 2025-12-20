package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/*
*

	{
	    "address": "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238",
	    "topics": [
	        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
	        "0x0000000000000000000000003c3380cdfb94dfeeaa41cad9f58254ae380d752d",
	        "0x000000000000000000000000157b434ed20aea9d647b5a7c308f0eb0b626f32e"
	    ],
	    "data": "0x00000000000000000000000000000000000000000000000000000000000f4240",
	    "blockNumber": "0x960611",
	    "transactionHash": "0x208123e31df6bf73032f03d2176320a6601f9f9bf58645753a86f531b4e81c50",
	    "transactionIndex": "0x4",
	    "blockHash": "0x3e2e72f1c56476833b85dbe399f1eadb6c0021c752d1d298d288e81a5a2d8626",
	    "blockTimestamp": "0x693d71d8",
	    "logIndex": "0x3",
	    "removed": false
	}
*/

func ConvertTransferLog(log string) (types.Log, error) {
	var tl types.Log
	if log == "" {
		return tl, fmt.Errorf("log is empty")
	}
	err := json.Unmarshal([]byte(log), &tl)
	if err != nil {
		return tl, err
	}
	return tl, nil
}

type TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func (te TransferEvent) String() string {
	return fmt.Sprintf("from %s \nto %s \nvalue %d\n", te.From.Hex(), te.To.Hex(), te.Value.Int64())
}

// ParseTransferLog  only parse the transfer log
func ParseTransferLog(log types.Log) {
	const transferEventABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

	transferABI, err := abi.JSON(strings.NewReader(transferEventABI))
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(transferABI)
	fmt.Printf("Transfer Abi %s \n", string(marshal))
	var transferEvent TransferEvent

	err = transferABI.UnpackIntoInterface(&transferEvent, "Transfer", log.Data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 解析 indexed 参数（topics[1] 和 topics[2]）
	if len(log.Topics) > 1 {
		transferEvent.From = common.BytesToAddress(log.Topics[1].Bytes())
	}
	if len(log.Topics) > 2 {
		transferEvent.To = common.BytesToAddress(log.Topics[2].Bytes())
	}
	fmt.Println(transferEvent)
}
