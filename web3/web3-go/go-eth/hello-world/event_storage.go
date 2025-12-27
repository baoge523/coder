// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EventStorageMetaData contains all meta data concerning the EventStorage contract.
var EventStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_id\",\"type\":\"int256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"str\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"display\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EventStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use EventStorageMetaData.ABI instead.
var EventStorageABI = EventStorageMetaData.ABI

// EventStorage is an auto generated Go binding around an Ethereum contract.
type EventStorage struct {
	EventStorageCaller     // Read-only binding to the contract
	EventStorageTransactor // Write-only binding to the contract
	EventStorageFilterer   // Log filterer for contract events
}

// EventStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventStorageSession struct {
	Contract     *EventStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventStorageCallerSession struct {
	Contract *EventStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventStorageTransactorSession struct {
	Contract     *EventStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventStorageRaw struct {
	Contract *EventStorage // Generic contract binding to access the raw methods on
}

// EventStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventStorageCallerRaw struct {
	Contract *EventStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EventStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventStorageTransactorRaw struct {
	Contract *EventStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventStorage creates a new instance of EventStorage, bound to a specific deployed contract.
func NewEventStorage(address common.Address, backend bind.ContractBackend) (*EventStorage, error) {
	contract, err := bindEventStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventStorage{EventStorageCaller: EventStorageCaller{contract: contract}, EventStorageTransactor: EventStorageTransactor{contract: contract}, EventStorageFilterer: EventStorageFilterer{contract: contract}}, nil
}

// NewEventStorageCaller creates a new read-only instance of EventStorage, bound to a specific deployed contract.
func NewEventStorageCaller(address common.Address, caller bind.ContractCaller) (*EventStorageCaller, error) {
	contract, err := bindEventStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventStorageCaller{contract: contract}, nil
}

// NewEventStorageTransactor creates a new write-only instance of EventStorage, bound to a specific deployed contract.
func NewEventStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EventStorageTransactor, error) {
	contract, err := bindEventStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventStorageTransactor{contract: contract}, nil
}

// NewEventStorageFilterer creates a new log filterer instance of EventStorage, bound to a specific deployed contract.
func NewEventStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EventStorageFilterer, error) {
	contract, err := bindEventStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventStorageFilterer{contract: contract}, nil
}

// bindEventStorage binds a generic wrapper to an already deployed contract.
func bindEventStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventStorage *EventStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventStorage.Contract.EventStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventStorage *EventStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventStorage.Contract.EventStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventStorage *EventStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventStorage.Contract.EventStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventStorage *EventStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventStorage *EventStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventStorage *EventStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventStorage.Contract.contract.Transact(opts, method, params...)
}

// Display is a free data retrieval call binding the contract method 0x0c1b7c1e.
//
// Solidity: function display() view returns(int256, uint256)
func (_EventStorage *EventStorageCaller) Display(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EventStorage.contract.Call(opts, &out, "display")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Display is a free data retrieval call binding the contract method 0x0c1b7c1e.
//
// Solidity: function display() view returns(int256, uint256)
func (_EventStorage *EventStorageSession) Display() (*big.Int, *big.Int, error) {
	return _EventStorage.Contract.Display(&_EventStorage.CallOpts)
}

// Display is a free data retrieval call binding the contract method 0x0c1b7c1e.
//
// Solidity: function display() view returns(int256, uint256)
func (_EventStorage *EventStorageCallerSession) Display() (*big.Int, *big.Int, error) {
	return _EventStorage.Contract.Display(&_EventStorage.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf04991f0.
//
// Solidity: function deposit(int256 _id) payable returns()
func (_EventStorage *EventStorageTransactor) Deposit(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _EventStorage.contract.Transact(opts, "deposit", _id)
}

// Deposit is a paid mutator transaction binding the contract method 0xf04991f0.
//
// Solidity: function deposit(int256 _id) payable returns()
func (_EventStorage *EventStorageSession) Deposit(_id *big.Int) (*types.Transaction, error) {
	return _EventStorage.Contract.Deposit(&_EventStorage.TransactOpts, _id)
}

// Deposit is a paid mutator transaction binding the contract method 0xf04991f0.
//
// Solidity: function deposit(int256 _id) payable returns()
func (_EventStorage *EventStorageTransactorSession) Deposit(_id *big.Int) (*types.Transaction, error) {
	return _EventStorage.Contract.Deposit(&_EventStorage.TransactOpts, _id)
}

// EventStorageDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EventStorage contract.
type EventStorageDepositIterator struct {
	Event *EventStorageDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventStorageDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventStorageDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventStorageDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventStorageDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventStorageDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventStorageDeposit represents a Deposit event raised by the EventStorage contract.
type EventStorageDeposit struct {
	From  common.Address
	Str   *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4706859488cc8327c56b7f7198045caedd94daed8ba07bc9317fd56deaaa7f0a.
//
// Solidity: event Deposit(address indexed from, int256 indexed str, uint256 value)
func (_EventStorage *EventStorageFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, str []*big.Int) (*EventStorageDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var strRule []interface{}
	for _, strItem := range str {
		strRule = append(strRule, strItem)
	}

	logs, sub, err := _EventStorage.contract.FilterLogs(opts, "Deposit", fromRule, strRule)
	if err != nil {
		return nil, err
	}
	return &EventStorageDepositIterator{contract: _EventStorage.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4706859488cc8327c56b7f7198045caedd94daed8ba07bc9317fd56deaaa7f0a.
//
// Solidity: event Deposit(address indexed from, int256 indexed str, uint256 value)
func (_EventStorage *EventStorageFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EventStorageDeposit, from []common.Address, str []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var strRule []interface{}
	for _, strItem := range str {
		strRule = append(strRule, strItem)
	}

	logs, sub, err := _EventStorage.contract.WatchLogs(opts, "Deposit", fromRule, strRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventStorageDeposit)
				if err := _EventStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x4706859488cc8327c56b7f7198045caedd94daed8ba07bc9317fd56deaaa7f0a.
//
// Solidity: event Deposit(address indexed from, int256 indexed str, uint256 value)
func (_EventStorage *EventStorageFilterer) ParseDeposit(log types.Log) (*EventStorageDeposit, error) {
	event := new(EventStorageDeposit)
	if err := _EventStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
