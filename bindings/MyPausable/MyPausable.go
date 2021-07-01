// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MyPausable

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

// MyPausableABI is the input ABI used to generate the binding from.
const MyPausableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MyPausableBin is the compiled bytecode used for deploying new contracts.
var MyPausableBin = "0x6080604052348015600f57600080fd5b5060008060006101000a81548160ff0219169083151502179055506097806100386000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80635c975abb14602d575b600080fd5b6033604b565b60405180821515815260200191505060405180910390f35b60008060009054906101000a900460ff1690509056fea264697066735822122055f20f1ee959950b58046cbe060e2ed17f4bba19f4a0219edd6be63d6822ac6464736f6c63430007060033"

// DeployMyPausable deploys a new Ethereum contract, binding an instance of MyPausable to it.
func DeployMyPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyPausable, error) {
	parsed, err := abi.JSON(strings.NewReader(MyPausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyPausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyPausable{MyPausableCaller: MyPausableCaller{contract: contract}, MyPausableTransactor: MyPausableTransactor{contract: contract}, MyPausableFilterer: MyPausableFilterer{contract: contract}}, nil
}

// MyPausable is an auto generated Go binding around an Ethereum contract.
type MyPausable struct {
	MyPausableCaller     // Read-only binding to the contract
	MyPausableTransactor // Write-only binding to the contract
	MyPausableFilterer   // Log filterer for contract events
}

// MyPausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyPausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyPausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyPausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyPausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyPausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyPausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyPausableSession struct {
	Contract     *MyPausable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyPausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyPausableCallerSession struct {
	Contract *MyPausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyPausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyPausableTransactorSession struct {
	Contract     *MyPausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyPausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyPausableRaw struct {
	Contract *MyPausable // Generic contract binding to access the raw methods on
}

// MyPausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyPausableCallerRaw struct {
	Contract *MyPausableCaller // Generic read-only contract binding to access the raw methods on
}

// MyPausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyPausableTransactorRaw struct {
	Contract *MyPausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyPausable creates a new instance of MyPausable, bound to a specific deployed contract.
func NewMyPausable(address common.Address, backend bind.ContractBackend) (*MyPausable, error) {
	contract, err := bindMyPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyPausable{MyPausableCaller: MyPausableCaller{contract: contract}, MyPausableTransactor: MyPausableTransactor{contract: contract}, MyPausableFilterer: MyPausableFilterer{contract: contract}}, nil
}

// NewMyPausableCaller creates a new read-only instance of MyPausable, bound to a specific deployed contract.
func NewMyPausableCaller(address common.Address, caller bind.ContractCaller) (*MyPausableCaller, error) {
	contract, err := bindMyPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyPausableCaller{contract: contract}, nil
}

// NewMyPausableTransactor creates a new write-only instance of MyPausable, bound to a specific deployed contract.
func NewMyPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*MyPausableTransactor, error) {
	contract, err := bindMyPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyPausableTransactor{contract: contract}, nil
}

// NewMyPausableFilterer creates a new log filterer instance of MyPausable, bound to a specific deployed contract.
func NewMyPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*MyPausableFilterer, error) {
	contract, err := bindMyPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyPausableFilterer{contract: contract}, nil
}

// bindMyPausable binds a generic wrapper to an already deployed contract.
func bindMyPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyPausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyPausable *MyPausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyPausable.Contract.MyPausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyPausable *MyPausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyPausable.Contract.MyPausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyPausable *MyPausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyPausable.Contract.MyPausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyPausable *MyPausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyPausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyPausable *MyPausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyPausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyPausable *MyPausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyPausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MyPausable *MyPausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MyPausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MyPausable *MyPausableSession) Paused() (bool, error) {
	return _MyPausable.Contract.Paused(&_MyPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MyPausable *MyPausableCallerSession) Paused() (bool, error) {
	return _MyPausable.Contract.Paused(&_MyPausable.CallOpts)
}

// MyPausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MyPausable contract.
type MyPausablePausedIterator struct {
	Event *MyPausablePaused // Event containing the contract specifics and raw log

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
func (it *MyPausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyPausablePaused)
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
		it.Event = new(MyPausablePaused)
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
func (it *MyPausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyPausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyPausablePaused represents a Paused event raised by the MyPausable contract.
type MyPausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MyPausable *MyPausableFilterer) FilterPaused(opts *bind.FilterOpts) (*MyPausablePausedIterator, error) {

	logs, sub, err := _MyPausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MyPausablePausedIterator{contract: _MyPausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MyPausable *MyPausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MyPausablePaused) (event.Subscription, error) {

	logs, sub, err := _MyPausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyPausablePaused)
				if err := _MyPausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MyPausable *MyPausableFilterer) ParsePaused(log types.Log) (*MyPausablePaused, error) {
	event := new(MyPausablePaused)
	if err := _MyPausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyPausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MyPausable contract.
type MyPausableUnpausedIterator struct {
	Event *MyPausableUnpaused // Event containing the contract specifics and raw log

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
func (it *MyPausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyPausableUnpaused)
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
		it.Event = new(MyPausableUnpaused)
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
func (it *MyPausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyPausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyPausableUnpaused represents a Unpaused event raised by the MyPausable contract.
type MyPausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MyPausable *MyPausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MyPausableUnpausedIterator, error) {

	logs, sub, err := _MyPausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MyPausableUnpausedIterator{contract: _MyPausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MyPausable *MyPausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MyPausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _MyPausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyPausableUnpaused)
				if err := _MyPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MyPausable *MyPausableFilterer) ParseUnpaused(log types.Log) (*MyPausableUnpaused, error) {
	event := new(MyPausableUnpaused)
	if err := _MyPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
