package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
	"web3-go/go-eth/tools"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

// go get github.com/ethereum/go-ethereum

var (
	address              = "" // 0.01
	address2             = "" // 0
	smartContractAddress = "" // smart contract

	SK            = ""
	ethSepoliaURL = "http://127.0.0.1:8545"
)

func init() {
	secure, err := tools.LoadSecureFile()
	if err != nil {
		panic(err)
	}
	address = secure.Accounts[0].Address
	address2 = secure.Accounts[1].Address
	smartContractAddress = secure.Accounts[2].Address
	SK = secure.Accounts[0].SK
	ethSepoliaURL = secure.EthSepoliaURL
}

func main() {
	// ipc and host:ip(local node)
	client, err := ethclient.Dial(ethSepoliaURL)
	if err != nil {
		panic(err)
	}
	//gasInfo(client)
	err = sendTransaction(client)
	if err != nil {
		panic(err)
	}
	//getTransactions(client)

	//blockInfo(client)

}

func findInfo(client *ethclient.Client) {
	ctx := context.Background()
	// get blockChain's current cumber
	number, err := client.BlockNumber(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Block number:", number)

	fee, err := client.BlobBaseFee(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fee basefee:", fee)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Chain ID:", chainID)

	nonce1, err := client.NonceAt(ctx, common.HexToAddress(address), nil)

	if err != nil {
		panic(err)
	}
	fmt.Println("Nonce:", nonce1)

	balance, err := client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Balance:", balance)

	balance2, err := client.BalanceAt(ctx, common.HexToAddress(address2), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Balance2:", balance2)

	blockInfo, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Block Info:", blockInfo)
	blockInfo2, err := client.BlockByHash(ctx, blockInfo.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println("Block Info2:", blockInfo2)

	codeAt, err := client.CodeAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Code At:", codeAt)

	baseFee, err := client.BlobBaseFee(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Base Fee:", baseFee)
	networkID, err := client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Network ID:", networkID)

	count, err := client.PeerCount(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Peer Count:", count)

	price, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Suggested GasPrice:", price)
	tipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("TipCap:", tipCap)
}

// sendTransaction sends a transaction with 1 ETH to a specified address.
func sendTransaction(cl *ethclient.Client) error {
	var (
		sk       = crypto.ToECDSAUnsafe(common.FromHex(SK))
		to       = common.HexToAddress(smartContractAddress)
		value    = new(big.Int).Mul(big.NewInt(10000000), big.NewInt(params.GWei))
		sender   = common.HexToAddress(address)
		gasLimit = uint64(21000)
	)

	// Retrieve the chainid (needed for signer)
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return err
	}
	// Retrieve the pending nonce
	nonce, err := cl.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return err
	}
	// Get suggested gas price
	tipCap, _ := cl.SuggestGasTipCap(context.Background())
	feeCap, _ := cl.SuggestGasPrice(context.Background())
	// Create a new transaction
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainid,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     value,
			Data:      nil,
		})
	// Sign the transaction using our keys
	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainid), sk)
	// Send the transaction to our node
	return cl.SendTransaction(context.Background(), signedTx)
}

func queryEvents(client *ethclient.Client) {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	blockNumberBig := big.NewInt(int64(blockNumber))
	fmt.Println("Block number:", blockNumberBig)

	eventSignatureBytes := []byte("Transfer(address,address,uint256)")
	eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)

	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(1)),
		ToBlock:   blockNumberBig,
		Topics: [][]common.Hash{
			{eventSignaturehash},
		},
	}

	logs, err := client.FilterLogs(context.Background(), q)
	if err != nil {
		fmt.Println("Failed to FilterLogs:", err)
		return
	}
	for _, l := range logs {
		marshal, _ := json.Marshal(l)
		fmt.Printf("%s \n", string(marshal))
	}
}

func queryEvent2(client *ethclient.Client) {
	// obtain block number
	number, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	blockNumberBig := big.NewInt(int64(number))
	eventSignatureBytes := []byte("Transfer(address,address,uint256)")
	eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)
	queryParam := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(2)),
		ToBlock:   blockNumberBig,
		Topics: [][]common.Hash{
			{eventSignaturehash},
		},
	}
	logs, err := client.FilterLogs(context.Background(), queryParam)
	if err != nil {
		fmt.Println("Failed to FilterLogs:", err)
		return
	}
	for _, l := range logs {
		for topic := range l.Topics {
			fmt.Printf("topic %v\n", topic)
		}
	}
}

func getTransactions(client *ethclient.Client) {
	transID := "0x491ba6828c57fdfd524535f1d39cb016d0575063714b579a73763ba76c8335a4"

	tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(transID))
	if err != nil {
		fmt.Println("Failed to retrieve transaction:", err)
		return
	}
	fmt.Println("is pending: ", pending)

	marshal, _ := json.Marshal(tx)
	fmt.Println("tx:", string(marshal))
}

func blockInfo(client *ethclient.Client) {

	blockInfo, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	fmt.Printf("block number: %d, tx len %d, uncles len %d, hash %v, receiveAt %v\n",
		blockInfo.Number(), len(blockInfo.Transactions()), len(blockInfo.Uncles()), blockInfo.Hash(), blockInfo.ReceivedAt)

	blockHeader, err := client.HeaderByNumber(context.Background(), blockInfo.Number())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
	}

	fmt.Println("block header:", blockHeader)

	headerHash, err := client.HeaderByHash(context.Background(), blockInfo.Hash())
	if err != nil {
		fmt.Println("Failed to retrieve block hash:", err)
	}
	marshal, _ := json.Marshal(headerHash)
	fmt.Println("block header by hash:", string(marshal))

}

func gasInfo(client *ethclient.Client) {
	ctx := context.Background()
	networkID, err := client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Network ID:", networkID)

	//count, err := client.PeerCount(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Peer Count:", count)
	fee, err := client.BlobBaseFee(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fee basefee:", fee)

	price, err := client.SuggestGasPrice(ctx) // max price for
	if err != nil {
		panic(err)
	}
	fmt.Println("Suggested GasPrice:", price)
	tipCap, err := client.SuggestGasTipCap(ctx) // priority unit:gwei
	if err != nil {
		panic(err)
	}
	fmt.Println("TipCap:", tipCap)

	// gasPrice >= baseFee + gasTipCap(priority)
}

// subscribeToLogs  subscribe event logs
func subscribeToLogs(client *ethclient.Client) {

	ctx := context.Background()

	msgReceiver := make(chan types.Log, 100)

	number, err := client.BlockNumber(ctx)
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	theLatestBlockNumber := big.NewInt(int64(number))

	subscribeQuery := ethereum.FilterQuery{
		FromBlock: theLatestBlockNumber,
		ToBlock:   nil, // always receive the most newly logs
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		},
	}
	logs, err := client.SubscribeFilterLogs(ctx, subscribeQuery, msgReceiver)
	if err != nil {
		fmt.Println("Failed to subscribe to logs:", err)
		return
	}
	defer func() {
		// unsubscribe
		logs.Unsubscribe()
	}()

	go func() {
		fmt.Println("start subscribing to logs")
		for {
			select {
			case <-ctx.Done():
				return
			case log := <-msgReceiver:
				ParseTransferLog(log)
			}
		}
	}()

	select {
	case <-ctx.Done():
	case err = <-logs.Err():
		fmt.Println("Failed to subscribe to logs:", err)
	case <-time.After(15 * time.Second):
		fmt.Println("close after 15s")
	}
}

func subscribeNewHeader(client *ethclient.Client) {
	ctx := context.Background()
	receiveNewHeader := make(chan *types.Header, 100)
	head, err := client.SubscribeNewHead(ctx, receiveNewHeader)
	if err != nil {
		fmt.Println("Failed to subscribe to new head:", err)
		return
	}
	defer func() {
		// unsubscribe
		head.Unsubscribe()
	}()

	go func() {
		fmt.Println("start subscribing to new header")
		for {
			select {
			case <-ctx.Done():
				return
			case log := <-receiveNewHeader:
				marshal, _ := json.Marshal(log)
				fmt.Println("new header:", string(marshal))
			}
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("ctx done")
	case err := <-head.Err():
		fmt.Println("Failed to subscribe to new head:", err)
	case <-time.After(15 * time.Second):
		fmt.Println("close after 15s")
	}
}

func subscribeTransactionsReceipt(client *ethclient.Client) {
	ctx := context.Background()

	receiveMsg := make(chan []*types.Receipt, 100)

	query := ethereum.TransactionReceiptsQuery{}

	receipts, err := client.SubscribeTransactionReceipts(ctx, &query, receiveMsg)
	if err != nil {
		fmt.Println("Failed to subscribe to receipts:", err)
		return
	}
	defer func() {
		// unsubscribe
		receipts.Unsubscribe()
	}()

	go func() {
		fmt.Println("start subscribing to new header")
		for {
			select {
			case <-ctx.Done():
				return
			case log := <-receiveMsg:

				if len(log) != 0 {
					for _, receipt := range log {
						marshal, _ := json.Marshal(receipt)
						fmt.Println("subscribe transaction receipts:", string(marshal))
					}

				} else {
					fmt.Println("log receipt len == 0, ignore receipt")
				}

			}
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("ctx done")
	case err := <-receipts.Err():
		fmt.Println("Failed to subscribe to new head:", err)
	case <-time.After(15 * time.Second):
		fmt.Println("close after 15s")
	}

}

// estimateGas demo
func estimateGas(client *ethclient.Client) {
	ctx := context.Background()

	gas, err := client.EstimateGas(ctx, ethereum.CallMsg{})
	if err != nil {
		fmt.Println("Failed to estimate gas:", err)
		return
	}
	fmt.Println("estimate gas:", gas)

	gas, err = client.EstimateGasAtBlock(ctx, ethereum.CallMsg{}, nil)
	if err != nil {
		fmt.Println("Failed to estimate gas:", err)
		return
	}
	fmt.Println("estimate gas at block:", gas)

	gas, err = client.EstimateGasAtBlockHash(ctx, ethereum.CallMsg{}, common.HexToHash("block hash string"))
	if err != nil {
		fmt.Println("Failed to estimate gas:", err)
		return
	}
	fmt.Println("estimate gas at block hash:", gas)

}

// transaction
func transaction(client *ethclient.Client) {
	ctx := context.Background()

	// get the number of the block
	count, err := client.TransactionCount(ctx, common.HexToHash("block hash string"))
	if err != nil {
		fmt.Println("Failed to transaction count:", err)
		return
	}
	fmt.Println("transaction count:", count)

	tx, pending, err := client.TransactionByHash(ctx, common.HexToHash("transaction hash string"))
	if err != nil {
		fmt.Println("Failed to transaction by hash:", err)
		return
	}
	fmt.Println("pending transaction:", pending) // current transaction status, is pending status ?
	fmt.Println("transaction by hash:", tx)

	// index: the index of the transaction in block (sort number)
	transactionInfo, err := client.TransactionInBlock(ctx, common.HexToHash("transaction in block's hash string"), 1)

	if err != nil {
		fmt.Println("Failed to transaction in block:", err)
		return
	}
	fmt.Println("transaction in block:", transactionInfo)

	chainID, _ := client.ChainID(ctx)
	sender, err := types.Sender(types.NewEIP155Signer(chainID), tx)
	if err == nil {
		fmt.Println("sender address:", sender.Hex())
		return
	}

	fmt.Printf("sender address: %v\n", sender)

	// get transaction sender, you can also use client.TransactionInBlock(...), or client.TransactionByHash(...) tx.
	sender, err = client.TransactionSender(ctx, &types.Transaction{}, common.HexToHash("block hash string"), 1)
	if err != nil {
		fmt.Println("Failed to transaction sender:", err)
		return
	}
	fmt.Println("sender transaction sender:", sender)

	// receipt represents the results of a transaction.
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash("transaction hash string, Note that the receipt is not available for pending transactions"))
	if err != nil {
		fmt.Println("Failed to transaction receipt:", err)
		return
	}
	fmt.Println("transaction receipt:", receipt)
}

// all kinds of the sendTransactions
func sendTransactionMethod(client *ethclient.Client) {
	ctx := context.Background()

	chainID := new(big.Int).SetInt64(11155111)
	nonce := uint64(1)
	// Get suggested gas price
	tipCap, _ := client.SuggestGasTipCap(context.Background())
	feeCap, _ := client.SuggestGasPrice(context.Background())
	gasLimit := uint64(21000)
	to := common.HexToAddress("to hash address string")
	// ETH
	value := big.NewInt(1)

	// this use DynamicFeeTx
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
		Gas:       gasLimit,
		To:        &to,
		Value:     value,
		Data:      nil,
	})
	// sign tx
	singTx, err2 := types.SignTx(tx, types.NewLondonSigner(chainID), crypto.ToECDSAUnsafe(common.FromHex(SK)))

	if err2 != nil {
		fmt.Println("Failed to sign tx:", err2)
		return
	}

	err := client.SendTransaction(ctx, singTx)

	if err != nil {
		fmt.Println("Failed to send tx:", err)
		return
	}

	timeout := 5 * time.Second

	syncReceipt, err2 := client.SendTransactionSync(ctx, singTx, &timeout)
	if err2 != nil {
		fmt.Println("Failed to send sync tx:", err2)
		return
	}
	fmt.Println("send sync tx receipt:", syncReceipt)

	rawSyncReceipt, err2 := client.SendRawTransactionSync(ctx, nil, &timeout)
	if err2 != nil {
		fmt.Println("Failed to send sync tx:", err2)
		return
	}
	fmt.Println("send sync tx receipt:", rawSyncReceipt)

}

// contract
func contract(client *ethclient.Client) {

	ctx := context.Background()

	toAddress := common.HexToAddress("to address hash string")
	tipCap, _ := client.SuggestGasTipCap(context.Background())
	feeCap, _ := client.SuggestGasPrice(context.Background())
	callMsg := ethereum.CallMsg{
		From:      common.HexToAddress("from address hash string"), // sender address
		To:        &toAddress,                                      // contract address
		Gas:       uint64(21000),                                   // gas limit
		GasPrice:  new(big.Int).SetInt64(params.GWei),              // todo  base fee ??
		GasFeeCap: feeCap,                                          // the max price for a pear
		GasTipCap: tipCap,                                          // tip gas
		Value:     big.NewInt(params.Wei),                          // amount of wei sent along with the call
		Data:      nil,                                             // usually ABI method invocation
	}

	callContract, err := client.CallContract(ctx, callMsg, nil)

	if err != nil {
		fmt.Println("Failed to call contract:", err)
		return
	}

	fmt.Println("call contract:", callContract)

	result, err := client.CallContractAtHash(ctx, callMsg, common.HexToHash("block hash string"))
	if err != nil {
		fmt.Println("Failed to call contract:", err)
		return
	}
	fmt.Println("call contract hash:", result)
}
