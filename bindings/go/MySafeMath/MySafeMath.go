// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MySafeMath

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

// MySafeMathABI is the input ABI used to generate the binding from.
const MySafeMathABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MySafeMathBin is the compiled bytecode used for deploying new contracts.
var MySafeMathBin = "0x608060405234801561001057600080fd5b506101bc806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063b67d77c514610030575b600080fd5b6100666004803603604081101561004657600080fd5b81019080803590602001909291908035906020019092919050505061007c565b6040518082815260200191505060405180910390f35b60006100be83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506100c6565b905092915050565b6000838311158290610173576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561013857808201518184015260208101905061011d565b50505050905090810190601f1680156101655780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050939250505056fea26469706673582212206ee820546a2ef6409afb0e16f8313a55c9ad77cc912fad53b7ccab3ec55639d564736f6c634300060c0033"

// DeployMySafeMath deploys a new Ethereum contract, binding an instance of MySafeMath to it.
func DeployMySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MySafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(MySafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MySafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MySafeMath{MySafeMathCaller: MySafeMathCaller{contract: contract}, MySafeMathTransactor: MySafeMathTransactor{contract: contract}, MySafeMathFilterer: MySafeMathFilterer{contract: contract}}, nil
}

// MySafeMath is an auto generated Go binding around an Ethereum contract.
type MySafeMath struct {
	MySafeMathCaller     // Read-only binding to the contract
	MySafeMathTransactor // Write-only binding to the contract
	MySafeMathFilterer   // Log filterer for contract events
}

// MySafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MySafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MySafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MySafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MySafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MySafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MySafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MySafeMathSession struct {
	Contract     *MySafeMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MySafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MySafeMathCallerSession struct {
	Contract *MySafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MySafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MySafeMathTransactorSession struct {
	Contract     *MySafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MySafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MySafeMathRaw struct {
	Contract *MySafeMath // Generic contract binding to access the raw methods on
}

// MySafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MySafeMathCallerRaw struct {
	Contract *MySafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// MySafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MySafeMathTransactorRaw struct {
	Contract *MySafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMySafeMath creates a new instance of MySafeMath, bound to a specific deployed contract.
func NewMySafeMath(address common.Address, backend bind.ContractBackend) (*MySafeMath, error) {
	contract, err := bindMySafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MySafeMath{MySafeMathCaller: MySafeMathCaller{contract: contract}, MySafeMathTransactor: MySafeMathTransactor{contract: contract}, MySafeMathFilterer: MySafeMathFilterer{contract: contract}}, nil
}

// NewMySafeMathCaller creates a new read-only instance of MySafeMath, bound to a specific deployed contract.
func NewMySafeMathCaller(address common.Address, caller bind.ContractCaller) (*MySafeMathCaller, error) {
	contract, err := bindMySafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MySafeMathCaller{contract: contract}, nil
}

// NewMySafeMathTransactor creates a new write-only instance of MySafeMath, bound to a specific deployed contract.
func NewMySafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MySafeMathTransactor, error) {
	contract, err := bindMySafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MySafeMathTransactor{contract: contract}, nil
}

// NewMySafeMathFilterer creates a new log filterer instance of MySafeMath, bound to a specific deployed contract.
func NewMySafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MySafeMathFilterer, error) {
	contract, err := bindMySafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MySafeMathFilterer{contract: contract}, nil
}

// bindMySafeMath binds a generic wrapper to an already deployed contract.
func bindMySafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MySafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MySafeMath *MySafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MySafeMath.Contract.MySafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MySafeMath *MySafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MySafeMath.Contract.MySafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MySafeMath *MySafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MySafeMath.Contract.MySafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MySafeMath *MySafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MySafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MySafeMath *MySafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MySafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MySafeMath *MySafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MySafeMath.Contract.contract.Transact(opts, method, params...)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_MySafeMath *MySafeMathCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MySafeMath.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_MySafeMath *MySafeMathSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MySafeMath.Contract.Sub(&_MySafeMath.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_MySafeMath *MySafeMathCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _MySafeMath.Contract.Sub(&_MySafeMath.CallOpts, a, b)
}
