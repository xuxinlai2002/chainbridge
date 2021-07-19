// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ToString

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

// ToStringABI is the input ABI used to generate the binding from.
const ToStringABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addressToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toStringBase\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"uintToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ToStringBin is the compiled bytecode used for deploying new contracts.
var ToStringBin = "0x608060405234801561001057600080fd5b506108b0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806339614e4f1461005c5780635e57966d146101905780638d1b75e41461024d5780639201de5514610381578063e939567914610428575b600080fd5b6101156004803603602081101561007257600080fd5b810190808035906020019064010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506104cf565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561015557808201518184015260208101905061013a565b50505050905090810190601f1680156101825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101d2600480360360208110156101a657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061054b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102125780820151818401526020810190506101f7565b50505050905090810190601f16801561023f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103066004803603602081101561026357600080fd5b810190808035906020019064010000000081111561028057600080fd5b82018360208201111561029257600080fd5b803590602001918460018302840111640100000000831117156102b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610595565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561034657808201518184015260208101905061032b565b50505050905090810190601f1680156103735780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103ad6004803603602081101561039757600080fd5b8101908080359060200190929190505050610818565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103ed5780820151818401526020810190506103d2565b50505050905090810190601f16801561041a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104546004803603602081101561043e57600080fd5b8101908080359060200190929190505050610849565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610494578082015181840152602081019050610479565b50505050905090810190601f1680156104c15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6060610544826040516020018082805190602001908083835b6020831061050b57805182526020820191506020810190506020830392506104e8565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052610595565b9050919050565b606061058e82604051602001808273ffffffffffffffffffffffffffffffffffffffff1660601b8152601401915050604051602081830303815290604052610595565b9050919050565b6060806040518060400160405280601081526020017f303132333435363738396162636465660000000000000000000000000000000081525090506060600284510260020167ffffffffffffffff811180156105f057600080fd5b506040519080825280601f01601f1916602001820160405280156106235781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061065457fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106106b157fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b845181101561080d578260048683815181106106fb57fe5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c60f81c60ff168151811061073a57fe5b602001015160f81c60f81b82600283026002018151811061075757fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535082600f60f81b86838151811061079857fe5b602001015160f81c60f81b1660f81c60ff16815181106107b457fe5b602001015160f81c60f81b8260028302600301815181106107d157fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506106e3565b508092505050919050565b60606108428260405160200180828152602001915050604051602081830303815290604052610595565b9050919050565b60606108738260405160200180828152602001915050604051602081830303815290604052610595565b905091905056fea2646970667358221220d4d5e8de3fa376f37d15a06a76395cde994431dce9b5606ba82e8fc599fdf78a64736f6c634300060c0033"

// DeployToString deploys a new Ethereum contract, binding an instance of ToString to it.
func DeployToString(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ToString, error) {
	parsed, err := abi.JSON(strings.NewReader(ToStringABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ToStringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ToString{ToStringCaller: ToStringCaller{contract: contract}, ToStringTransactor: ToStringTransactor{contract: contract}, ToStringFilterer: ToStringFilterer{contract: contract}}, nil
}

// ToString is an auto generated Go binding around an Ethereum contract.
type ToString struct {
	ToStringCaller     // Read-only binding to the contract
	ToStringTransactor // Write-only binding to the contract
	ToStringFilterer   // Log filterer for contract events
}

// ToStringCaller is an auto generated read-only Go binding around an Ethereum contract.
type ToStringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToStringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ToStringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToStringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ToStringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToStringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ToStringSession struct {
	Contract     *ToString         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ToStringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ToStringCallerSession struct {
	Contract *ToStringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ToStringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ToStringTransactorSession struct {
	Contract     *ToStringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ToStringRaw is an auto generated low-level Go binding around an Ethereum contract.
type ToStringRaw struct {
	Contract *ToString // Generic contract binding to access the raw methods on
}

// ToStringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ToStringCallerRaw struct {
	Contract *ToStringCaller // Generic read-only contract binding to access the raw methods on
}

// ToStringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ToStringTransactorRaw struct {
	Contract *ToStringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToString creates a new instance of ToString, bound to a specific deployed contract.
func NewToString(address common.Address, backend bind.ContractBackend) (*ToString, error) {
	contract, err := bindToString(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ToString{ToStringCaller: ToStringCaller{contract: contract}, ToStringTransactor: ToStringTransactor{contract: contract}, ToStringFilterer: ToStringFilterer{contract: contract}}, nil
}

// NewToStringCaller creates a new read-only instance of ToString, bound to a specific deployed contract.
func NewToStringCaller(address common.Address, caller bind.ContractCaller) (*ToStringCaller, error) {
	contract, err := bindToString(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ToStringCaller{contract: contract}, nil
}

// NewToStringTransactor creates a new write-only instance of ToString, bound to a specific deployed contract.
func NewToStringTransactor(address common.Address, transactor bind.ContractTransactor) (*ToStringTransactor, error) {
	contract, err := bindToString(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ToStringTransactor{contract: contract}, nil
}

// NewToStringFilterer creates a new log filterer instance of ToString, bound to a specific deployed contract.
func NewToStringFilterer(address common.Address, filterer bind.ContractFilterer) (*ToStringFilterer, error) {
	contract, err := bindToString(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ToStringFilterer{contract: contract}, nil
}

// bindToString binds a generic wrapper to an already deployed contract.
func bindToString(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ToStringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ToString *ToStringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ToString.Contract.ToStringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ToString *ToStringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToString.Contract.ToStringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToString *ToStringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ToString.Contract.ToStringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ToString *ToStringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ToString.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ToString *ToStringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToString.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToString *ToStringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ToString.Contract.contract.Transact(opts, method, params...)
}

// AddressToString is a free data retrieval call binding the contract method 0x5e57966d.
//
// Solidity: function addressToString(address account) pure returns(string)
func (_ToString *ToStringCaller) AddressToString(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _ToString.contract.Call(opts, &out, "addressToString", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToString is a free data retrieval call binding the contract method 0x5e57966d.
//
// Solidity: function addressToString(address account) pure returns(string)
func (_ToString *ToStringSession) AddressToString(account common.Address) (string, error) {
	return _ToString.Contract.AddressToString(&_ToString.CallOpts, account)
}

// AddressToString is a free data retrieval call binding the contract method 0x5e57966d.
//
// Solidity: function addressToString(address account) pure returns(string)
func (_ToString *ToStringCallerSession) AddressToString(account common.Address) (string, error) {
	return _ToString.Contract.AddressToString(&_ToString.CallOpts, account)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 value) pure returns(string)
func (_ToString *ToStringCaller) Bytes32ToString(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _ToString.contract.Call(opts, &out, "bytes32ToString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 value) pure returns(string)
func (_ToString *ToStringSession) Bytes32ToString(value [32]byte) (string, error) {
	return _ToString.Contract.Bytes32ToString(&_ToString.CallOpts, value)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 value) pure returns(string)
func (_ToString *ToStringCallerSession) Bytes32ToString(value [32]byte) (string, error) {
	return _ToString.Contract.Bytes32ToString(&_ToString.CallOpts, value)
}

// BytesToString is a free data retrieval call binding the contract method 0x39614e4f.
//
// Solidity: function bytesToString(bytes value) pure returns(string)
func (_ToString *ToStringCaller) BytesToString(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _ToString.contract.Call(opts, &out, "bytesToString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BytesToString is a free data retrieval call binding the contract method 0x39614e4f.
//
// Solidity: function bytesToString(bytes value) pure returns(string)
func (_ToString *ToStringSession) BytesToString(value []byte) (string, error) {
	return _ToString.Contract.BytesToString(&_ToString.CallOpts, value)
}

// BytesToString is a free data retrieval call binding the contract method 0x39614e4f.
//
// Solidity: function bytesToString(bytes value) pure returns(string)
func (_ToString *ToStringCallerSession) BytesToString(value []byte) (string, error) {
	return _ToString.Contract.BytesToString(&_ToString.CallOpts, value)
}

// ToStringBase is a free data retrieval call binding the contract method 0x8d1b75e4.
//
// Solidity: function toStringBase(bytes data) pure returns(string)
func (_ToString *ToStringCaller) ToStringBase(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _ToString.contract.Call(opts, &out, "toStringBase", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToStringBase is a free data retrieval call binding the contract method 0x8d1b75e4.
//
// Solidity: function toStringBase(bytes data) pure returns(string)
func (_ToString *ToStringSession) ToStringBase(data []byte) (string, error) {
	return _ToString.Contract.ToStringBase(&_ToString.CallOpts, data)
}

// ToStringBase is a free data retrieval call binding the contract method 0x8d1b75e4.
//
// Solidity: function toStringBase(bytes data) pure returns(string)
func (_ToString *ToStringCallerSession) ToStringBase(data []byte) (string, error) {
	return _ToString.Contract.ToStringBase(&_ToString.CallOpts, data)
}

// UintToString is a free data retrieval call binding the contract method 0xe9395679.
//
// Solidity: function uintToString(uint256 value) pure returns(string)
func (_ToString *ToStringCaller) UintToString(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _ToString.contract.Call(opts, &out, "uintToString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UintToString is a free data retrieval call binding the contract method 0xe9395679.
//
// Solidity: function uintToString(uint256 value) pure returns(string)
func (_ToString *ToStringSession) UintToString(value *big.Int) (string, error) {
	return _ToString.Contract.UintToString(&_ToString.CallOpts, value)
}

// UintToString is a free data retrieval call binding the contract method 0xe9395679.
//
// Solidity: function uintToString(uint256 value) pure returns(string)
func (_ToString *ToStringCallerSession) UintToString(value *big.Int) (string, error) {
	return _ToString.Contract.UintToString(&_ToString.CallOpts, value)
}
