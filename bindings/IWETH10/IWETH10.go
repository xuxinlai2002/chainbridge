// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IWETH10

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

// IWETH10ABI is the input ABI used to generate the binding from.
const IWETH10ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositToAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWETH10 is an auto generated Go binding around an Ethereum contract.
type IWETH10 struct {
	IWETH10Caller     // Read-only binding to the contract
	IWETH10Transactor // Write-only binding to the contract
	IWETH10Filterer   // Log filterer for contract events
}

// IWETH10Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH10Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH10Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH10Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH10Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH10Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH10Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH10Session struct {
	Contract     *IWETH10          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH10CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH10CallerSession struct {
	Contract *IWETH10Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IWETH10TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH10TransactorSession struct {
	Contract     *IWETH10Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IWETH10Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH10Raw struct {
	Contract *IWETH10 // Generic contract binding to access the raw methods on
}

// IWETH10CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH10CallerRaw struct {
	Contract *IWETH10Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH10TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH10TransactorRaw struct {
	Contract *IWETH10Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH10 creates a new instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10(address common.Address, backend bind.ContractBackend) (*IWETH10, error) {
	contract, err := bindIWETH10(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH10{IWETH10Caller: IWETH10Caller{contract: contract}, IWETH10Transactor: IWETH10Transactor{contract: contract}, IWETH10Filterer: IWETH10Filterer{contract: contract}}, nil
}

// NewIWETH10Caller creates a new read-only instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10Caller(address common.Address, caller bind.ContractCaller) (*IWETH10Caller, error) {
	contract, err := bindIWETH10(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH10Caller{contract: contract}, nil
}

// NewIWETH10Transactor creates a new write-only instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH10Transactor, error) {
	contract, err := bindIWETH10(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH10Transactor{contract: contract}, nil
}

// NewIWETH10Filterer creates a new log filterer instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH10Filterer, error) {
	contract, err := bindIWETH10(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH10Filterer{contract: contract}, nil
}

// bindIWETH10 binds a generic wrapper to an already deployed contract.
func bindIWETH10(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETH10ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH10 *IWETH10Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH10.Contract.IWETH10Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH10 *IWETH10Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH10.Contract.IWETH10Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH10 *IWETH10Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH10.Contract.IWETH10Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH10 *IWETH10CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH10.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH10 *IWETH10TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH10.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH10 *IWETH10TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH10.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IWETH10 *IWETH10Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IWETH10 *IWETH10Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _IWETH10.Contract.DOMAINSEPARATOR(&_IWETH10.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IWETH10 *IWETH10CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IWETH10.Contract.DOMAINSEPARATOR(&_IWETH10.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH10 *IWETH10Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH10 *IWETH10Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Allowance(&_IWETH10.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Allowance(&_IWETH10.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH10 *IWETH10Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH10 *IWETH10Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH10.Contract.BalanceOf(&_IWETH10.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH10.Contract.BalanceOf(&_IWETH10.CallOpts, account)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IWETH10 *IWETH10Caller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IWETH10 *IWETH10Session) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IWETH10.Contract.FlashFee(&_IWETH10.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IWETH10.Contract.FlashFee(&_IWETH10.CallOpts, token, amount)
}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_IWETH10 *IWETH10Caller) FlashMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "flashMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_IWETH10 *IWETH10Session) FlashMinted() (*big.Int, error) {
	return _IWETH10.Contract.FlashMinted(&_IWETH10.CallOpts)
}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) FlashMinted() (*big.Int, error) {
	return _IWETH10.Contract.FlashMinted(&_IWETH10.CallOpts)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IWETH10 *IWETH10Caller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IWETH10 *IWETH10Session) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IWETH10.Contract.MaxFlashLoan(&_IWETH10.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IWETH10.Contract.MaxFlashLoan(&_IWETH10.CallOpts, token)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IWETH10 *IWETH10Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IWETH10 *IWETH10Session) Nonces(owner common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Nonces(&_IWETH10.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Nonces(&_IWETH10.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH10 *IWETH10Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH10 *IWETH10Session) TotalSupply() (*big.Int, error) {
	return _IWETH10.Contract.TotalSupply(&_IWETH10.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH10.Contract.TotalSupply(&_IWETH10.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Approve(&_IWETH10.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Approve(&_IWETH10.TransactOpts, spender, amount)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Transactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "approveAndCall", spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Session) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.ApproveAndCall(&_IWETH10.TransactOpts, spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.ApproveAndCall(&_IWETH10.TransactOpts, spender, value, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH10 *IWETH10Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH10 *IWETH10Session) Deposit() (*types.Transaction, error) {
	return _IWETH10.Contract.Deposit(&_IWETH10.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH10 *IWETH10TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH10.Contract.Deposit(&_IWETH10.TransactOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_IWETH10 *IWETH10Transactor) DepositTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "depositTo", to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_IWETH10 *IWETH10Session) DepositTo(to common.Address) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositTo(&_IWETH10.TransactOpts, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_IWETH10 *IWETH10TransactorSession) DepositTo(to common.Address) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositTo(&_IWETH10.TransactOpts, to)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool)
func (_IWETH10 *IWETH10Transactor) DepositToAndCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "depositToAndCall", to, data)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool)
func (_IWETH10 *IWETH10Session) DepositToAndCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositToAndCall(&_IWETH10.TransactOpts, to, data)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool)
func (_IWETH10 *IWETH10TransactorSession) DepositToAndCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositToAndCall(&_IWETH10.TransactOpts, to, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IWETH10 *IWETH10Transactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IWETH10 *IWETH10Session) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.FlashLoan(&_IWETH10.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.FlashLoan(&_IWETH10.TransactOpts, receiver, token, amount, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IWETH10 *IWETH10Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IWETH10 *IWETH10Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IWETH10.Contract.Permit(&_IWETH10.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IWETH10 *IWETH10TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IWETH10.Contract.Permit(&_IWETH10.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Transfer(&_IWETH10.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Transfer(&_IWETH10.TransactOpts, recipient, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Session) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferAndCall(&_IWETH10.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferAndCall(&_IWETH10.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferFrom(&_IWETH10.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferFrom(&_IWETH10.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_IWETH10 *IWETH10Transactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_IWETH10 *IWETH10Session) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Withdraw(&_IWETH10.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_IWETH10 *IWETH10TransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Withdraw(&_IWETH10.TransactOpts, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_IWETH10 *IWETH10Transactor) WithdrawFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "withdrawFrom", from, to, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_IWETH10 *IWETH10Session) WithdrawFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawFrom(&_IWETH10.TransactOpts, from, to, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_IWETH10 *IWETH10TransactorSession) WithdrawFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawFrom(&_IWETH10.TransactOpts, from, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_IWETH10 *IWETH10Transactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "withdrawTo", to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_IWETH10 *IWETH10Session) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawTo(&_IWETH10.TransactOpts, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_IWETH10 *IWETH10TransactorSession) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawTo(&_IWETH10.TransactOpts, to, value)
}

// IWETH10ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH10 contract.
type IWETH10ApprovalIterator struct {
	Event *IWETH10Approval // Event containing the contract specifics and raw log

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
func (it *IWETH10ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH10Approval)
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
		it.Event = new(IWETH10Approval)
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
func (it *IWETH10ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH10ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH10Approval represents a Approval event raised by the IWETH10 contract.
type IWETH10Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH10 *IWETH10Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETH10ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH10.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETH10ApprovalIterator{contract: _IWETH10.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH10 *IWETH10Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETH10Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH10.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH10Approval)
				if err := _IWETH10.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH10 *IWETH10Filterer) ParseApproval(log types.Log) (*IWETH10Approval, error) {
	event := new(IWETH10Approval)
	if err := _IWETH10.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH10TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH10 contract.
type IWETH10TransferIterator struct {
	Event *IWETH10Transfer // Event containing the contract specifics and raw log

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
func (it *IWETH10TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH10Transfer)
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
		it.Event = new(IWETH10Transfer)
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
func (it *IWETH10TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH10TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH10Transfer represents a Transfer event raised by the IWETH10 contract.
type IWETH10Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH10 *IWETH10Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETH10TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH10.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETH10TransferIterator{contract: _IWETH10.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH10 *IWETH10Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETH10Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH10.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH10Transfer)
				if err := _IWETH10.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH10 *IWETH10Filterer) ParseTransfer(log types.Log) (*IWETH10Transfer, error) {
	event := new(IWETH10Transfer)
	if err := _IWETH10.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
