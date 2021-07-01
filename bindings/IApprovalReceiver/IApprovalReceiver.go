// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IApprovalReceiver

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

// IApprovalReceiverABI is the input ABI used to generate the binding from.
const IApprovalReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTokenApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IApprovalReceiver is an auto generated Go binding around an Ethereum contract.
type IApprovalReceiver struct {
	IApprovalReceiverCaller     // Read-only binding to the contract
	IApprovalReceiverTransactor // Write-only binding to the contract
	IApprovalReceiverFilterer   // Log filterer for contract events
}

// IApprovalReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IApprovalReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApprovalReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IApprovalReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApprovalReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IApprovalReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApprovalReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IApprovalReceiverSession struct {
	Contract     *IApprovalReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IApprovalReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IApprovalReceiverCallerSession struct {
	Contract *IApprovalReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IApprovalReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IApprovalReceiverTransactorSession struct {
	Contract     *IApprovalReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IApprovalReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IApprovalReceiverRaw struct {
	Contract *IApprovalReceiver // Generic contract binding to access the raw methods on
}

// IApprovalReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IApprovalReceiverCallerRaw struct {
	Contract *IApprovalReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IApprovalReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IApprovalReceiverTransactorRaw struct {
	Contract *IApprovalReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIApprovalReceiver creates a new instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiver(address common.Address, backend bind.ContractBackend) (*IApprovalReceiver, error) {
	contract, err := bindIApprovalReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiver{IApprovalReceiverCaller: IApprovalReceiverCaller{contract: contract}, IApprovalReceiverTransactor: IApprovalReceiverTransactor{contract: contract}, IApprovalReceiverFilterer: IApprovalReceiverFilterer{contract: contract}}, nil
}

// NewIApprovalReceiverCaller creates a new read-only instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiverCaller(address common.Address, caller bind.ContractCaller) (*IApprovalReceiverCaller, error) {
	contract, err := bindIApprovalReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiverCaller{contract: contract}, nil
}

// NewIApprovalReceiverTransactor creates a new write-only instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IApprovalReceiverTransactor, error) {
	contract, err := bindIApprovalReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiverTransactor{contract: contract}, nil
}

// NewIApprovalReceiverFilterer creates a new log filterer instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IApprovalReceiverFilterer, error) {
	contract, err := bindIApprovalReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiverFilterer{contract: contract}, nil
}

// bindIApprovalReceiver binds a generic wrapper to an already deployed contract.
func bindIApprovalReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IApprovalReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApprovalReceiver *IApprovalReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApprovalReceiver.Contract.IApprovalReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApprovalReceiver *IApprovalReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.IApprovalReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApprovalReceiver *IApprovalReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.IApprovalReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApprovalReceiver *IApprovalReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApprovalReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApprovalReceiver *IApprovalReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApprovalReceiver *IApprovalReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnTokenApproval is a paid mutator transaction binding the contract method 0x00ba451f.
//
// Solidity: function onTokenApproval(address , uint256 , bytes ) returns(bool)
func (_IApprovalReceiver *IApprovalReceiverTransactor) OnTokenApproval(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IApprovalReceiver.contract.Transact(opts, "onTokenApproval", arg0, arg1, arg2)
}

// OnTokenApproval is a paid mutator transaction binding the contract method 0x00ba451f.
//
// Solidity: function onTokenApproval(address , uint256 , bytes ) returns(bool)
func (_IApprovalReceiver *IApprovalReceiverSession) OnTokenApproval(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.OnTokenApproval(&_IApprovalReceiver.TransactOpts, arg0, arg1, arg2)
}

// OnTokenApproval is a paid mutator transaction binding the contract method 0x00ba451f.
//
// Solidity: function onTokenApproval(address , uint256 , bytes ) returns(bool)
func (_IApprovalReceiver *IApprovalReceiverTransactorSession) OnTokenApproval(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.OnTokenApproval(&_IApprovalReceiver.TransactOpts, arg0, arg1, arg2)
}
