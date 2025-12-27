package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// demo for smart contract in go

type ContractCaller struct {
	ctx context.Context

	chainID *big.Int
	// client eth client
	client *ethclient.Client
	// primaryKey ecdsa primary key
	primaryKey *ecdsa.PrivateKey
	// fromAddress
	fromAddress common.Address
}

func NewContractCaller() *ContractCaller {
	// connect node
	cli, err := ethclient.Dial(ethSepoliaURL)
	if err != nil {
		panic(err)
	}
	// obtain primary key
	primaryKey, err := crypto.HexToECDSA(smartContractOwnerSK)
	if err != nil {
		panic(err)
	}
	// obtain public key
	publicKey := primaryKey.Public()
	// convert crypto.PublicKey to ecdsa.PublicKey
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	ctx := context.Background()

	chain, err := cli.ChainID(ctx)
	if err != nil {
		panic(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &ContractCaller{
		client:      cli,
		primaryKey:  primaryKey,
		fromAddress: fromAddress,
		ctx:         ctx,
		chainID:     chain,
	}

}

func (cc *ContractCaller) CallViewFunction(contractAddress common.Address) error {

	storage, err := NewStorage(contractAddress, cc.client)
	if err != nil {
		return err
	}
	retrieve, err := storage.Retrieve(&bind.CallOpts{
		From:    cc.fromAddress,
		Context: cc.ctx, // cancel or timeout
	})
	if err != nil {
		return fmt.Errorf("storage retrieve failed: %v", err)
	}
	fmt.Printf("retrieve: %d\n", retrieve.Int64())
	return nil
}

func (cc *ContractCaller) CallWriteFunction(contractAddress common.Address) error {
	storage, err := NewStorage(contractAddress, cc.client)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(cc.primaryKey, cc.chainID)
	if err != nil {
		return err
	}
	tipCap, _ := cc.client.SuggestGasTipCap(context.Background())
	feeCap, _ := cc.client.SuggestGasPrice(context.Background())
	gasLimit := cc.estimateGas(contractAddress, big.NewInt(11))
	fmt.Printf("gasLimit: %d\n", gasLimit)
	opts := bind.TransactOpts{
		Signer:    auth.Signer,
		From:      cc.fromAddress,
		GasLimit:  gasLimit * 2, // 注意这里智能合约比简单转账需要更多gas，所以这里会需要更大的gasLimit，不然会导致交易失败
		GasFeeCap: feeCap,
		GasTipCap: tipCap,
	}

	trans, err := storage.Store(&opts, big.NewInt(11))
	if err != nil {
		return fmt.Errorf("storage store failed: %v", err)
	}
	marshal, _ := json.Marshal(trans)
	fmt.Printf("store trans: %s\n", marshal)

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), cc.client, trans)
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

func (cc *ContractCaller) estimateGas(contractAddress common.Address, newValue *big.Int) uint64 {

	msg := ethereum.CallMsg{
		//From: cc.fromAddress,
		//To: &contractAddress,
	}
	gas, err := cc.client.EstimateGas(cc.ctx, msg)
	if err != nil {
		panic(err)
	}
	return gas
}

// 编码函数调用数据
func (cc *ContractCaller) encodeSetValueData(newValue *big.Int) []byte {
	// setValue(uint256) 的函数选择器
	functionSelector := []byte{0x20, 0x96, 0x52, 0x55} // 前4字节

	// 编码参数（uint256 是32字节）
	paddedValue := common.LeftPadBytes(newValue.Bytes(), 32)

	return append(functionSelector, paddedValue...)
}
