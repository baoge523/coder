
# hello-eth

## gas
gasPrice >= baseGas + gasTipPrice

gasPrice: means the max price you can pay

baseGas: execution in EVM, the base gas

gasTipPrice: (priority) tip for validator

```go
price, err := client.SuggestGasPrice(ctx)

baseFee, err := client.BlobBaseFee(ctx)

tipCap, err := client.SuggestGasTipCap(ctx)
```

gasLimit: a number of gas, not price

sender need pay: total price = (process num < gas limit) * ((baseGas + gasTipCap) < gasPrice)

## block
```go
// If `number` is nil, the latest known block is returned.
blockInfo, err := client.BlockByNumber(context.Background(), nil)

// just get block header info
blockHeader, err := client.HeaderByNumber(context.Background(), blockInfo.Number())
headerHash, err := client.HeaderByHash(context.Background(), blockInfo.Hash())
```

## transaction

### getTransaction
```go
// tx transaction info ; pending  is pending status
tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(transID))
```

### sendTransaction

 - primary key (sender account)
 - to account address
 - sender account address
 - value (amount of ETH)
 - gasLimit
 - chainID
 - nonce (sender)
 - tipCap
 - feeCap

```go
    sk  := crypto.ToECDSAUnsafe(common.FromHex(SK))
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
	cl.SendTransaction(context.Background(), signedTx)
```

learn types and crypto package is required


## event (logs)

transaction logs
```go
    eventSignatureBytes := []byte("Transfer(address,address,uint256)")
	eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)  // hash the content

	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(2)),
		ToBlock:   blockNumberBig,
		Topics: [][]common.Hash{
			{eventSignaturehash},
		},
	}

	logs, err := client.FilterLogs(context.Background(), q)
```

the event log is smart contract event (ABI)

such as:
    - transfer (Transfer(address,address,uint256))  hash (0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef)
    - other ...

## ABI
   transfer
   ```json
   [
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "name": "from",
        "type": "address"
      },
      {
        "indexed": true,
        "name": "to",
        "type": "address"
      },
      {
        "indexed": false,
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "Transfer",
    "type": "event"
  }
]
   ```
  

## install abigen
```linux
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

 abigen --abi=abi/storage.abi --pkg=storage --type=Storage --out=storage.go
```

# package

## common

## types

## account and account abi