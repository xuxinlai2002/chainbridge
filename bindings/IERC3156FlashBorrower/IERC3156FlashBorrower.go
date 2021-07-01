// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC3156FlashBorrower

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

// IERC3156FlashBorrowerABI is the input ABI used to generate the binding from.
const IERC3156FlashBorrowerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC3156FlashBorrower is an auto generated Go binding around an Ethereum contract.
type IERC3156FlashBorrower struct {
	IERC3156FlashBorrowerCaller     // Read-only binding to the contract
	IERC3156FlashBorrowerTransactor // Write-only binding to the contract
	IERC3156FlashBorrowerFilterer   // Log filterer for contract events
}

// IERC3156FlashBorrowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC3156FlashBorrowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC3156FlashBorrowerSession struct {
	Contract     *IERC3156FlashBorrower // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IERC3156FlashBorrowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC3156FlashBorrowerCallerSession struct {
	Contract *IERC3156FlashBorrowerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IERC3156FlashBorrowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC3156FlashBorrowerTransactorSession struct {
	Contract     *IERC3156FlashBorrowerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IERC3156FlashBorrowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC3156FlashBorrowerRaw struct {
	Contract *IERC3156FlashBorrower // Generic contract binding to access the raw methods on
}

// IERC3156FlashBorrowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerCallerRaw struct {
	Contract *IERC3156FlashBorrowerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC3156FlashBorrowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerTransactorRaw struct {
	Contract *IERC3156FlashBorrowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC3156FlashBorrower creates a new instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrower(address common.Address, backend bind.ContractBackend) (*IERC3156FlashBorrower, error) {
	contract, err := bindIERC3156FlashBorrower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrower{IERC3156FlashBorrowerCaller: IERC3156FlashBorrowerCaller{contract: contract}, IERC3156FlashBorrowerTransactor: IERC3156FlashBorrowerTransactor{contract: contract}, IERC3156FlashBorrowerFilterer: IERC3156FlashBorrowerFilterer{contract: contract}}, nil
}

// NewIERC3156FlashBorrowerCaller creates a new read-only instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerCaller(address common.Address, caller bind.ContractCaller) (*IERC3156FlashBorrowerCaller, error) {
	contract, err := bindIERC3156FlashBorrower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerCaller{contract: contract}, nil
}

// NewIERC3156FlashBorrowerTransactor creates a new write-only instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC3156FlashBorrowerTransactor, error) {
	contract, err := bindIERC3156FlashBorrower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerTransactor{contract: contract}, nil
}

// NewIERC3156FlashBorrowerFilterer creates a new log filterer instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC3156FlashBorrowerFilterer, error) {
	contract, err := bindIERC3156FlashBorrower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerFilterer{contract: contract}, nil
}

// bindIERC3156FlashBorrower binds a generic wrapper to an already deployed contract.
func bindIERC3156FlashBorrower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC3156FlashBorrowerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashBorrower.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.contract.Transact(opts, method, params...)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactor) OnFlashLoan(opts *bind.TransactOpts, initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.contract.Transact(opts, "onFlashLoan", initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.OnFlashLoan(&_IERC3156FlashBorrower.TransactOpts, initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.OnFlashLoan(&_IERC3156FlashBorrower.TransactOpts, initiator, token, amount, fee, data)
}
