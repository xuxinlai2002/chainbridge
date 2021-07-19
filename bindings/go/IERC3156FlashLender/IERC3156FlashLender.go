// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC3156FlashLender

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

// IERC3156FlashLenderABI is the input ABI used to generate the binding from.
const IERC3156FlashLenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC3156FlashLender is an auto generated Go binding around an Ethereum contract.
type IERC3156FlashLender struct {
	IERC3156FlashLenderCaller     // Read-only binding to the contract
	IERC3156FlashLenderTransactor // Write-only binding to the contract
	IERC3156FlashLenderFilterer   // Log filterer for contract events
}

// IERC3156FlashLenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC3156FlashLenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC3156FlashLenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC3156FlashLenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC3156FlashLenderSession struct {
	Contract     *IERC3156FlashLender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC3156FlashLenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC3156FlashLenderCallerSession struct {
	Contract *IERC3156FlashLenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC3156FlashLenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC3156FlashLenderTransactorSession struct {
	Contract     *IERC3156FlashLenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC3156FlashLenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC3156FlashLenderRaw struct {
	Contract *IERC3156FlashLender // Generic contract binding to access the raw methods on
}

// IERC3156FlashLenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC3156FlashLenderCallerRaw struct {
	Contract *IERC3156FlashLenderCaller // Generic read-only contract binding to access the raw methods on
}

// IERC3156FlashLenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC3156FlashLenderTransactorRaw struct {
	Contract *IERC3156FlashLenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC3156FlashLender creates a new instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLender(address common.Address, backend bind.ContractBackend) (*IERC3156FlashLender, error) {
	contract, err := bindIERC3156FlashLender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLender{IERC3156FlashLenderCaller: IERC3156FlashLenderCaller{contract: contract}, IERC3156FlashLenderTransactor: IERC3156FlashLenderTransactor{contract: contract}, IERC3156FlashLenderFilterer: IERC3156FlashLenderFilterer{contract: contract}}, nil
}

// NewIERC3156FlashLenderCaller creates a new read-only instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderCaller(address common.Address, caller bind.ContractCaller) (*IERC3156FlashLenderCaller, error) {
	contract, err := bindIERC3156FlashLender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderCaller{contract: contract}, nil
}

// NewIERC3156FlashLenderTransactor creates a new write-only instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC3156FlashLenderTransactor, error) {
	contract, err := bindIERC3156FlashLender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderTransactor{contract: contract}, nil
}

// NewIERC3156FlashLenderFilterer creates a new log filterer instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC3156FlashLenderFilterer, error) {
	contract, err := bindIERC3156FlashLender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderFilterer{contract: contract}, nil
}

// bindIERC3156FlashLender binds a generic wrapper to an already deployed contract.
func bindIERC3156FlashLender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC3156FlashLenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashLender *IERC3156FlashLenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashLender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.contract.Transact(opts, method, params...)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCaller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC3156FlashLender.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.FlashFee(&_IERC3156FlashLender.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.FlashFee(&_IERC3156FlashLender.CallOpts, token, amount)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC3156FlashLender.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.MaxFlashLoan(&_IERC3156FlashLender.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.MaxFlashLoan(&_IERC3156FlashLender.CallOpts, token)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.FlashLoan(&_IERC3156FlashLender.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.FlashLoan(&_IERC3156FlashLender.TransactOpts, receiver, token, amount, data)
}
