// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package history_contract

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
)

// HistoryContractHistory is an auto generated low-level Go binding around an user-defined struct.
type HistoryContractHistory struct {
	Timestamp *big.Int
	Data      [32]byte
}

// HistoryContractMetaData contains all meta data concerning the HistoryContract contract.
var HistoryContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"recordIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"internalType\":\"structHistoryContract.History\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HistoryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use HistoryContractMetaData.ABI instead.
var HistoryContractABI = HistoryContractMetaData.ABI

// HistoryContract is an auto generated Go binding around an Ethereum contract.
type HistoryContract struct {
	HistoryContractCaller     // Read-only binding to the contract
	HistoryContractTransactor // Write-only binding to the contract
	HistoryContractFilterer   // Log filterer for contract events
}

// HistoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type HistoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HistoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HistoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HistoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HistoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HistoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HistoryContractSession struct {
	Contract     *HistoryContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HistoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HistoryContractCallerSession struct {
	Contract *HistoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HistoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HistoryContractTransactorSession struct {
	Contract     *HistoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HistoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type HistoryContractRaw struct {
	Contract *HistoryContract // Generic contract binding to access the raw methods on
}

// HistoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HistoryContractCallerRaw struct {
	Contract *HistoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// HistoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HistoryContractTransactorRaw struct {
	Contract *HistoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHistoryContract creates a new instance of HistoryContract, bound to a specific deployed contract.
func NewHistoryContract(address common.Address, backend bind.ContractBackend) (*HistoryContract, error) {
	contract, err := bindHistoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HistoryContract{HistoryContractCaller: HistoryContractCaller{contract: contract}, HistoryContractTransactor: HistoryContractTransactor{contract: contract}, HistoryContractFilterer: HistoryContractFilterer{contract: contract}}, nil
}

// NewHistoryContractCaller creates a new read-only instance of HistoryContract, bound to a specific deployed contract.
func NewHistoryContractCaller(address common.Address, caller bind.ContractCaller) (*HistoryContractCaller, error) {
	contract, err := bindHistoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HistoryContractCaller{contract: contract}, nil
}

// NewHistoryContractTransactor creates a new write-only instance of HistoryContract, bound to a specific deployed contract.
func NewHistoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*HistoryContractTransactor, error) {
	contract, err := bindHistoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HistoryContractTransactor{contract: contract}, nil
}

// NewHistoryContractFilterer creates a new log filterer instance of HistoryContract, bound to a specific deployed contract.
func NewHistoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*HistoryContractFilterer, error) {
	contract, err := bindHistoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HistoryContractFilterer{contract: contract}, nil
}

// bindHistoryContract binds a generic wrapper to an already deployed contract.
func bindHistoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HistoryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HistoryContract *HistoryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HistoryContract.Contract.HistoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HistoryContract *HistoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HistoryContract.Contract.HistoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HistoryContract *HistoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HistoryContract.Contract.HistoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HistoryContract *HistoryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HistoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HistoryContract *HistoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HistoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HistoryContract *HistoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HistoryContract.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 _index) view returns((uint256,bytes32))
func (_HistoryContract *HistoryContractCaller) GetData(opts *bind.CallOpts, _index *big.Int) (HistoryContractHistory, error) {
	var out []interface{}
	err := _HistoryContract.contract.Call(opts, &out, "getData", _index)

	if err != nil {
		return *new(HistoryContractHistory), err
	}

	out0 := *abi.ConvertType(out[0], new(HistoryContractHistory)).(*HistoryContractHistory)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 _index) view returns((uint256,bytes32))
func (_HistoryContract *HistoryContractSession) GetData(_index *big.Int) (HistoryContractHistory, error) {
	return _HistoryContract.Contract.GetData(&_HistoryContract.CallOpts, _index)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 _index) view returns((uint256,bytes32))
func (_HistoryContract *HistoryContractCallerSession) GetData(_index *big.Int) (HistoryContractHistory, error) {
	return _HistoryContract.Contract.GetData(&_HistoryContract.CallOpts, _index)
}

// StoreData is a paid mutator transaction binding the contract method 0x3fc906c5.
//
// Solidity: function storeData(bytes32 _data) returns()
func (_HistoryContract *HistoryContractTransactor) StoreData(opts *bind.TransactOpts, _data [32]byte) (*types.Transaction, error) {
	return _HistoryContract.contract.Transact(opts, "storeData", _data)
}

// StoreData is a paid mutator transaction binding the contract method 0x3fc906c5.
//
// Solidity: function storeData(bytes32 _data) returns()
func (_HistoryContract *HistoryContractSession) StoreData(_data [32]byte) (*types.Transaction, error) {
	return _HistoryContract.Contract.StoreData(&_HistoryContract.TransactOpts, _data)
}

// StoreData is a paid mutator transaction binding the contract method 0x3fc906c5.
//
// Solidity: function storeData(bytes32 _data) returns()
func (_HistoryContract *HistoryContractTransactorSession) StoreData(_data [32]byte) (*types.Transaction, error) {
	return _HistoryContract.Contract.StoreData(&_HistoryContract.TransactOpts, _data)
}

// HistoryContractDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the HistoryContract contract.
type HistoryContractDataStoredIterator struct {
	Event *HistoryContractDataStored // Event containing the contract specifics and raw log

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
func (it *HistoryContractDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryContractDataStored)
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
		it.Event = new(HistoryContractDataStored)
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
func (it *HistoryContractDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryContractDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryContractDataStored represents a DataStored event raised by the HistoryContract contract.
type HistoryContractDataStored struct {
	User        common.Address
	RecordIndex *big.Int
	Data        [32]byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0xcb5c0173ec1aa1ca5096d7dcb185ee453c0d56841672539c22346a8611cff6f1.
//
// Solidity: event DataStored(address indexed user, uint256 indexed recordIndex, bytes32 data, uint256 timestamp)
func (_HistoryContract *HistoryContractFilterer) FilterDataStored(opts *bind.FilterOpts, user []common.Address, recordIndex []*big.Int) (*HistoryContractDataStoredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var recordIndexRule []interface{}
	for _, recordIndexItem := range recordIndex {
		recordIndexRule = append(recordIndexRule, recordIndexItem)
	}

	logs, sub, err := _HistoryContract.contract.FilterLogs(opts, "DataStored", userRule, recordIndexRule)
	if err != nil {
		return nil, err
	}
	return &HistoryContractDataStoredIterator{contract: _HistoryContract.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0xcb5c0173ec1aa1ca5096d7dcb185ee453c0d56841672539c22346a8611cff6f1.
//
// Solidity: event DataStored(address indexed user, uint256 indexed recordIndex, bytes32 data, uint256 timestamp)
func (_HistoryContract *HistoryContractFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *HistoryContractDataStored, user []common.Address, recordIndex []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var recordIndexRule []interface{}
	for _, recordIndexItem := range recordIndex {
		recordIndexRule = append(recordIndexRule, recordIndexItem)
	}

	logs, sub, err := _HistoryContract.contract.WatchLogs(opts, "DataStored", userRule, recordIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryContractDataStored)
				if err := _HistoryContract.contract.UnpackLog(event, "DataStored", log); err != nil {
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

// ParseDataStored is a log parse operation binding the contract event 0xcb5c0173ec1aa1ca5096d7dcb185ee453c0d56841672539c22346a8611cff6f1.
//
// Solidity: event DataStored(address indexed user, uint256 indexed recordIndex, bytes32 data, uint256 timestamp)
func (_HistoryContract *HistoryContractFilterer) ParseDataStored(log types.Log) (*HistoryContractDataStored, error) {
	event := new(HistoryContractDataStored)
	if err := _HistoryContract.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
