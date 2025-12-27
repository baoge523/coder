package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func (cc *ContractCaller) CallEventDisplay(contractAddress common.Address) error {
	storage, err := NewEventStorage(contractAddress, cc.client)
	if err != nil {
		return err
	}
	id, value, err := storage.Display(&bind.CallOpts{
		From:    cc.fromAddress,
		Context: cc.ctx, // cancel or timeout
	})
	if err != nil {
		return err
	}
	fmt.Println(id)
	fmt.Println(value)
	return nil
}

func (cc *ContractCaller) CallEventDeposit(contractAddress common.Address) error {
	storage, err := NewEventStorage(contractAddress, cc.client)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(cc.primaryKey, cc.chainID)
	if err != nil {
		return err
	}
	tipCap, _ := cc.client.SuggestGasTipCap(cc.ctx)
	feeCap, _ := cc.client.SuggestGasPrice(cc.ctx)
	gasLimit := cc.estimateGas(contractAddress, big.NewInt(11))
	fmt.Printf("gasLimit: %d\n", gasLimit)
	opts := bind.TransactOpts{
		Signer:    auth.Signer,
		From:      cc.fromAddress,
		GasLimit:  gasLimit * 2, // 注意这里智能合约比简单转账需要更多gas，所以这里会需要更大的gasLimit，不然会导致交易失败
		GasFeeCap: feeCap,
		GasTipCap: tipCap,
		Value:     big.NewInt(params.GWei),
	}

	trans, err := storage.Deposit(&opts, big.NewInt(11))
	if err != nil {
		return err
	}
	// 等待交易确认
	receipt, err := bind.WaitMined(cc.ctx, cc.client, trans)
	if err != nil {
		panic(err)
	}
	// 检查交易状态
	if receipt.Status != 1 {
		fmt.Println("交易执行失败，状态:", receipt.Status)
	} else {
		fmt.Println("交易成功确认，区块:", receipt.BlockNumber)
	}
	return nil

}

func (cc *ContractCaller) obtainEventLogs() error {

	blockNumber, err := cc.client.BlockNumber(cc.ctx)
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return err
	}
	blockNumberBig := big.NewInt(int64(blockNumber))
	fmt.Println("Block number:", blockNumberBig)

	eventSignatureBytes := []byte("Deposit(address,int256,uint256)")
	eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)

	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(1)),
		ToBlock:   blockNumberBig,
		Topics: [][]common.Hash{
			{eventSignaturehash},
		},
	}

	logs, err := cc.client.FilterLogs(cc.ctx, q)
	if err != nil {
		fmt.Println("Failed to FilterLogs:", err)
		return err
	}
	for _, l := range logs {
		marshal, _ := json.Marshal(l)
		fmt.Printf("%s \n", string(marshal))
	}
	return nil
}
