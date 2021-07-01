// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITransferReceiver

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ITransferReceiverABI is the input ABI used to generate the binding from.
const ITransferReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITransferReceiver is an auto generated Go binding around an Ethereum contract.
type ITransferReceiver struct {
	ITransferReceiverCaller     // Read-only binding to the contract
	ITransferReceiverTransactor // Write-only binding to the contract
	ITransferReceiverFilterer   // Log filterer for contract events
}

// ITransferReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITransferReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITransferReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITransferReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITransferReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITransferReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITransferReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITransferReceiverSession struct {
	Contract     *ITransferReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ITransferReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITransferReceiverCallerSession struct {
	Contract *ITransferReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ITransferReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITransferReceiverTransactorSession struct {
	Contract     *ITransferReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ITransferReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITransferReceiverRaw struct {
	Contract *ITransferReceiver // Generic contract binding to access the raw methods on
}

// ITransferReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITransferReceiverCallerRaw struct {
	Contract *ITransferReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ITransferReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITransferReceiverTransactorRaw struct {
	Contract *ITransferReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITransferReceiver creates a new instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiver(address common.Address, backend bind.ContractBackend) (*ITransferReceiver, error) {
	contract, err := bindITransferReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiver{ITransferReceiverCaller: ITransferReceiverCaller{contract: contract}, ITransferReceiverTransactor: ITransferReceiverTransactor{contract: contract}, ITransferReceiverFilterer: ITransferReceiverFilterer{contract: contract}}, nil
}

// NewITransferReceiverCaller creates a new read-only instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiverCaller(address common.Address, caller bind.ContractCaller) (*ITransferReceiverCaller, error) {
	contract, err := bindITransferReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiverCaller{contract: contract}, nil
}

// NewITransferReceiverTransactor creates a new write-only instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ITransferReceiverTransactor, error) {
	contract, err := bindITransferReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiverTransactor{contract: contract}, nil
}

// NewITransferReceiverFilterer creates a new log filterer instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ITransferReceiverFilterer, error) {
	contract, err := bindITransferReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiverFilterer{contract: contract}, nil
}

// bindITransferReceiver binds a generic wrapper to an already deployed contract.
func bindITransferReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITransferReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITransferReceiver *ITransferReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITransferReceiver.Contract.ITransferReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITransferReceiver *ITransferReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.ITransferReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITransferReceiver *ITransferReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.ITransferReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITransferReceiver *ITransferReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITransferReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITransferReceiver *ITransferReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITransferReceiver *ITransferReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes ) returns(bool)
func (_ITransferReceiver *ITransferReceiverTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _ITransferReceiver.contract.Transact(opts, "onTokenTransfer", arg0, arg1, arg2)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes ) returns(bool)
func (_ITransferReceiver *ITransferReceiverSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.OnTokenTransfer(&_ITransferReceiver.TransactOpts, arg0, arg1, arg2)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes ) returns(bool)
func (_ITransferReceiver *ITransferReceiverTransactorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.OnTokenTransfer(&_ITransferReceiver.TransactOpts, arg0, arg1, arg2)
}
