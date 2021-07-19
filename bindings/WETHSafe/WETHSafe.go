// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WETHSafe

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

// WETHSafeABI is the input ABI used to generate the binding from.
const WETHSafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WETHSafeBin is the compiled bytecode used for deploying new contracts.
var WETHSafeBin = "0x608060405234801561001057600080fd5b5060b28061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063783b6c5414602d575b600080fd5b607660048036036040811015604157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506078565b005b505056fea2646970667358221220905c07256294a3de8ffb09eb1ea8df56f01bdfc8f99322b2f4a5ccdd64e2a6c864736f6c634300060c0033"

// DeployWETHSafe deploys a new Ethereum contract, binding an instance of WETHSafe to it.
func DeployWETHSafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WETHSafe, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHSafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WETHSafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WETHSafe{WETHSafeCaller: WETHSafeCaller{contract: contract}, WETHSafeTransactor: WETHSafeTransactor{contract: contract}, WETHSafeFilterer: WETHSafeFilterer{contract: contract}}, nil
}

// WETHSafe is an auto generated Go binding around an Ethereum contract.
type WETHSafe struct {
	WETHSafeCaller     // Read-only binding to the contract
	WETHSafeTransactor // Write-only binding to the contract
	WETHSafeFilterer   // Log filterer for contract events
}

// WETHSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHSafeSession struct {
	Contract     *WETHSafe         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHSafeCallerSession struct {
	Contract *WETHSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WETHSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHSafeTransactorSession struct {
	Contract     *WETHSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WETHSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHSafeRaw struct {
	Contract *WETHSafe // Generic contract binding to access the raw methods on
}

// WETHSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHSafeCallerRaw struct {
	Contract *WETHSafeCaller // Generic read-only contract binding to access the raw methods on
}

// WETHSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHSafeTransactorRaw struct {
	Contract *WETHSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETHSafe creates a new instance of WETHSafe, bound to a specific deployed contract.
func NewWETHSafe(address common.Address, backend bind.ContractBackend) (*WETHSafe, error) {
	contract, err := bindWETHSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETHSafe{WETHSafeCaller: WETHSafeCaller{contract: contract}, WETHSafeTransactor: WETHSafeTransactor{contract: contract}, WETHSafeFilterer: WETHSafeFilterer{contract: contract}}, nil
}

// NewWETHSafeCaller creates a new read-only instance of WETHSafe, bound to a specific deployed contract.
func NewWETHSafeCaller(address common.Address, caller bind.ContractCaller) (*WETHSafeCaller, error) {
	contract, err := bindWETHSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHSafeCaller{contract: contract}, nil
}

// NewWETHSafeTransactor creates a new write-only instance of WETHSafe, bound to a specific deployed contract.
func NewWETHSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHSafeTransactor, error) {
	contract, err := bindWETHSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHSafeTransactor{contract: contract}, nil
}

// NewWETHSafeFilterer creates a new log filterer instance of WETHSafe, bound to a specific deployed contract.
func NewWETHSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHSafeFilterer, error) {
	contract, err := bindWETHSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHSafeFilterer{contract: contract}, nil
}

// bindWETHSafe binds a generic wrapper to an already deployed contract.
func bindWETHSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHSafe *WETHSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHSafe.Contract.WETHSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHSafe *WETHSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHSafe.Contract.WETHSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHSafe *WETHSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHSafe.Contract.WETHSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHSafe *WETHSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHSafe *WETHSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHSafe *WETHSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHSafe.Contract.contract.Transact(opts, method, params...)
}

// TransferWETH is a paid mutator transaction binding the contract method 0x783b6c54.
//
// Solidity: function transferWETH(address to, uint256 amount) returns()
func (_WETHSafe *WETHSafeTransactor) TransferWETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHSafe.contract.Transact(opts, "transferWETH", to, amount)
}

// TransferWETH is a paid mutator transaction binding the contract method 0x783b6c54.
//
// Solidity: function transferWETH(address to, uint256 amount) returns()
func (_WETHSafe *WETHSafeSession) TransferWETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHSafe.Contract.TransferWETH(&_WETHSafe.TransactOpts, to, amount)
}

// TransferWETH is a paid mutator transaction binding the contract method 0x783b6c54.
//
// Solidity: function transferWETH(address to, uint256 amount) returns()
func (_WETHSafe *WETHSafeTransactorSession) TransferWETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHSafe.Contract.TransferWETH(&_WETHSafe.TransactOpts, to, amount)
}
