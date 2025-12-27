
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

## access the smart contracts on alchemy via golang
all steps
1. apply an account on alchemy and you should create an app
2. write a simple solidity smart contract on remix
3. choose the sepolia test network and set network address (your app access address on alchemy) on metaMask
4. deploy smart contract and Environment choose sepolia-testnet-metaMask (you should create a smart contract account)
5. get a smart contract address (Deployed Contracts) 
6. get ABI(json file) and generate xxx.go file vie abigen tool
7. write a golang progress to access smart contract

### smart contract solidity
```solidity
// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Storage {

    uint256 number;

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number = num;
    }

    /**
     * @dev Return value 
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256){
        return number;
    }
}
```

### golang program
1. access address
2. smart contract


# package

## common

## types

## account and account abi