// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC2612

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

// IERC2612ABI is the input ABI used to generate the binding from.
const IERC2612ABI = "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC2612 is an auto generated Go binding around an Ethereum contract.
type IERC2612 struct {
	IERC2612Caller     // Read-only binding to the contract
	IERC2612Transactor // Write-only binding to the contract
	IERC2612Filterer   // Log filterer for contract events
}

// IERC2612Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC2612Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2612Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC2612Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2612Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC2612Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2612Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC2612Session struct {
	Contract     *IERC2612         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC2612CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC2612CallerSession struct {
	Contract *IERC2612Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC2612TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC2612TransactorSession struct {
	Contract     *IERC2612Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC2612Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC2612Raw struct {
	Contract *IERC2612 // Generic contract binding to access the raw methods on
}

// IERC2612CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC2612CallerRaw struct {
	Contract *IERC2612Caller // Generic read-only contract binding to access the raw methods on
}

// IERC2612TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC2612TransactorRaw struct {
	Contract *IERC2612Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC2612 creates a new instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612(address common.Address, backend bind.ContractBackend) (*IERC2612, error) {
	contract, err := bindIERC2612(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC2612{IERC2612Caller: IERC2612Caller{contract: contract}, IERC2612Transactor: IERC2612Transactor{contract: contract}, IERC2612Filterer: IERC2612Filterer{contract: contract}}, nil
}

// NewIERC2612Caller creates a new read-only instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612Caller(address common.Address, caller bind.ContractCaller) (*IERC2612Caller, error) {
	contract, err := bindIERC2612(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC2612Caller{contract: contract}, nil
}

// NewIERC2612Transactor creates a new write-only instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC2612Transactor, error) {
	contract, err := bindIERC2612(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC2612Transactor{contract: contract}, nil
}

// NewIERC2612Filterer creates a new log filterer instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC2612Filterer, error) {
	contract, err := bindIERC2612(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC2612Filterer{contract: contract}, nil
}

// bindIERC2612 binds a generic wrapper to an already deployed contract.
func bindIERC2612(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC2612ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC2612 *IERC2612Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC2612.Contract.IERC2612Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC2612 *IERC2612Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC2612.Contract.IERC2612Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC2612 *IERC2612Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC2612.Contract.IERC2612Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC2612 *IERC2612CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC2612.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC2612 *IERC2612TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC2612.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC2612 *IERC2612TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC2612.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC2612 *IERC2612Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC2612.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC2612 *IERC2612Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC2612.Contract.DOMAINSEPARATOR(&_IERC2612.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC2612 *IERC2612CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC2612.Contract.DOMAINSEPARATOR(&_IERC2612.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC2612 *IERC2612Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC2612.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC2612 *IERC2612Session) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC2612.Contract.Nonces(&_IERC2612.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC2612 *IERC2612CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC2612.Contract.Nonces(&_IERC2612.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC2612 *IERC2612Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC2612.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC2612 *IERC2612Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC2612.Contract.Permit(&_IERC2612.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC2612 *IERC2612TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC2612.Contract.Permit(&_IERC2612.TransactOpts, owner, spender, value, deadline, v, r, s)
}
