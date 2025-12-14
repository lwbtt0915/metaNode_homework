// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountDecreasedByStep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountDecremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountIncreasedByStep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountReset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"}],\"name\":\"decreaseByStep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"}],\"name\":\"increaseByStep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Token *TokenCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Token *TokenSession) GetCount() (*big.Int, error) {
	return _Token.Contract.GetCount(&_Token.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Token *TokenCallerSession) GetCount() (*big.Int, error) {
	return _Token.Contract.GetCount(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// DecreaseByStep is a paid mutator transaction binding the contract method 0x30ced4ab.
//
// Solidity: function decreaseByStep(uint256 step) returns()
func (_Token *TokenTransactor) DecreaseByStep(opts *bind.TransactOpts, step *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseByStep", step)
}

// DecreaseByStep is a paid mutator transaction binding the contract method 0x30ced4ab.
//
// Solidity: function decreaseByStep(uint256 step) returns()
func (_Token *TokenSession) DecreaseByStep(step *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseByStep(&_Token.TransactOpts, step)
}

// DecreaseByStep is a paid mutator transaction binding the contract method 0x30ced4ab.
//
// Solidity: function decreaseByStep(uint256 step) returns()
func (_Token *TokenTransactorSession) DecreaseByStep(step *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseByStep(&_Token.TransactOpts, step)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Token *TokenTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Token *TokenSession) Decrement() (*types.Transaction, error) {
	return _Token.Contract.Decrement(&_Token.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Token *TokenTransactorSession) Decrement() (*types.Transaction, error) {
	return _Token.Contract.Decrement(&_Token.TransactOpts)
}

// IncreaseByStep is a paid mutator transaction binding the contract method 0x56e340e4.
//
// Solidity: function increaseByStep(uint256 step) returns()
func (_Token *TokenTransactor) IncreaseByStep(opts *bind.TransactOpts, step *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseByStep", step)
}

// IncreaseByStep is a paid mutator transaction binding the contract method 0x56e340e4.
//
// Solidity: function increaseByStep(uint256 step) returns()
func (_Token *TokenSession) IncreaseByStep(step *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseByStep(&_Token.TransactOpts, step)
}

// IncreaseByStep is a paid mutator transaction binding the contract method 0x56e340e4.
//
// Solidity: function increaseByStep(uint256 step) returns()
func (_Token *TokenTransactorSession) IncreaseByStep(step *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseByStep(&_Token.TransactOpts, step)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Token *TokenTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Token *TokenSession) Increment() (*types.Transaction, error) {
	return _Token.Contract.Increment(&_Token.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Token *TokenTransactorSession) Increment() (*types.Transaction, error) {
	return _Token.Contract.Increment(&_Token.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Token *TokenTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Token *TokenSession) Reset() (*types.Transaction, error) {
	return _Token.Contract.Reset(&_Token.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Token *TokenTransactorSession) Reset() (*types.Transaction, error) {
	return _Token.Contract.Reset(&_Token.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TokenCountDecreasedByStepIterator is returned from FilterCountDecreasedByStep and is used to iterate over the raw logs and unpacked data for CountDecreasedByStep events raised by the Token contract.
type TokenCountDecreasedByStepIterator struct {
	Event *TokenCountDecreasedByStep // Event containing the contract specifics and raw log

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
func (it *TokenCountDecreasedByStepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCountDecreasedByStep)
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
		it.Event = new(TokenCountDecreasedByStep)
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
func (it *TokenCountDecreasedByStepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCountDecreasedByStepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCountDecreasedByStep represents a CountDecreasedByStep event raised by the Token contract.
type TokenCountDecreasedByStep struct {
	Step     *big.Int
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountDecreasedByStep is a free log retrieval operation binding the contract event 0xc33af293bbb8717d1d4fd23b5c0a03fd71348605267d2954f343f766ac70281c.
//
// Solidity: event CountDecreasedByStep(uint256 step, uint256 newCount)
func (_Token *TokenFilterer) FilterCountDecreasedByStep(opts *bind.FilterOpts) (*TokenCountDecreasedByStepIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "CountDecreasedByStep")
	if err != nil {
		return nil, err
	}
	return &TokenCountDecreasedByStepIterator{contract: _Token.contract, event: "CountDecreasedByStep", logs: logs, sub: sub}, nil
}

// WatchCountDecreasedByStep is a free log subscription operation binding the contract event 0xc33af293bbb8717d1d4fd23b5c0a03fd71348605267d2954f343f766ac70281c.
//
// Solidity: event CountDecreasedByStep(uint256 step, uint256 newCount)
func (_Token *TokenFilterer) WatchCountDecreasedByStep(opts *bind.WatchOpts, sink chan<- *TokenCountDecreasedByStep) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "CountDecreasedByStep")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCountDecreasedByStep)
				if err := _Token.contract.UnpackLog(event, "CountDecreasedByStep", log); err != nil {
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

// ParseCountDecreasedByStep is a log parse operation binding the contract event 0xc33af293bbb8717d1d4fd23b5c0a03fd71348605267d2954f343f766ac70281c.
//
// Solidity: event CountDecreasedByStep(uint256 step, uint256 newCount)
func (_Token *TokenFilterer) ParseCountDecreasedByStep(log types.Log) (*TokenCountDecreasedByStep, error) {
	event := new(TokenCountDecreasedByStep)
	if err := _Token.contract.UnpackLog(event, "CountDecreasedByStep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCountDecrementedIterator is returned from FilterCountDecremented and is used to iterate over the raw logs and unpacked data for CountDecremented events raised by the Token contract.
type TokenCountDecrementedIterator struct {
	Event *TokenCountDecremented // Event containing the contract specifics and raw log

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
func (it *TokenCountDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCountDecremented)
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
		it.Event = new(TokenCountDecremented)
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
func (it *TokenCountDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCountDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCountDecremented represents a CountDecremented event raised by the Token contract.
type TokenCountDecremented struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountDecremented is a free log retrieval operation binding the contract event 0x36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af9506.
//
// Solidity: event CountDecremented(uint256 newCount)
func (_Token *TokenFilterer) FilterCountDecremented(opts *bind.FilterOpts) (*TokenCountDecrementedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "CountDecremented")
	if err != nil {
		return nil, err
	}
	return &TokenCountDecrementedIterator{contract: _Token.contract, event: "CountDecremented", logs: logs, sub: sub}, nil
}

// WatchCountDecremented is a free log subscription operation binding the contract event 0x36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af9506.
//
// Solidity: event CountDecremented(uint256 newCount)
func (_Token *TokenFilterer) WatchCountDecremented(opts *bind.WatchOpts, sink chan<- *TokenCountDecremented) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "CountDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCountDecremented)
				if err := _Token.contract.UnpackLog(event, "CountDecremented", log); err != nil {
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

// ParseCountDecremented is a log parse operation binding the contract event 0x36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af9506.
//
// Solidity: event CountDecremented(uint256 newCount)
func (_Token *TokenFilterer) ParseCountDecremented(log types.Log) (*TokenCountDecremented, error) {
	event := new(TokenCountDecremented)
	if err := _Token.contract.UnpackLog(event, "CountDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCountIncreasedByStepIterator is returned from FilterCountIncreasedByStep and is used to iterate over the raw logs and unpacked data for CountIncreasedByStep events raised by the Token contract.
type TokenCountIncreasedByStepIterator struct {
	Event *TokenCountIncreasedByStep // Event containing the contract specifics and raw log

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
func (it *TokenCountIncreasedByStepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCountIncreasedByStep)
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
		it.Event = new(TokenCountIncreasedByStep)
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
func (it *TokenCountIncreasedByStepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCountIncreasedByStepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCountIncreasedByStep represents a CountIncreasedByStep event raised by the Token contract.
type TokenCountIncreasedByStep struct {
	Step     *big.Int
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountIncreasedByStep is a free log retrieval operation binding the contract event 0xce5493a868a3452ea0b45c42ae587bb1a0377c0acc990ed1d22b154b04aa6cf6.
//
// Solidity: event CountIncreasedByStep(uint256 step, uint256 newCount)
func (_Token *TokenFilterer) FilterCountIncreasedByStep(opts *bind.FilterOpts) (*TokenCountIncreasedByStepIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "CountIncreasedByStep")
	if err != nil {
		return nil, err
	}
	return &TokenCountIncreasedByStepIterator{contract: _Token.contract, event: "CountIncreasedByStep", logs: logs, sub: sub}, nil
}

// WatchCountIncreasedByStep is a free log subscription operation binding the contract event 0xce5493a868a3452ea0b45c42ae587bb1a0377c0acc990ed1d22b154b04aa6cf6.
//
// Solidity: event CountIncreasedByStep(uint256 step, uint256 newCount)
func (_Token *TokenFilterer) WatchCountIncreasedByStep(opts *bind.WatchOpts, sink chan<- *TokenCountIncreasedByStep) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "CountIncreasedByStep")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCountIncreasedByStep)
				if err := _Token.contract.UnpackLog(event, "CountIncreasedByStep", log); err != nil {
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

// ParseCountIncreasedByStep is a log parse operation binding the contract event 0xce5493a868a3452ea0b45c42ae587bb1a0377c0acc990ed1d22b154b04aa6cf6.
//
// Solidity: event CountIncreasedByStep(uint256 step, uint256 newCount)
func (_Token *TokenFilterer) ParseCountIncreasedByStep(log types.Log) (*TokenCountIncreasedByStep, error) {
	event := new(TokenCountIncreasedByStep)
	if err := _Token.contract.UnpackLog(event, "CountIncreasedByStep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCountIncrementedIterator is returned from FilterCountIncremented and is used to iterate over the raw logs and unpacked data for CountIncremented events raised by the Token contract.
type TokenCountIncrementedIterator struct {
	Event *TokenCountIncremented // Event containing the contract specifics and raw log

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
func (it *TokenCountIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCountIncremented)
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
		it.Event = new(TokenCountIncremented)
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
func (it *TokenCountIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCountIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCountIncremented represents a CountIncremented event raised by the Token contract.
type TokenCountIncremented struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountIncremented is a free log retrieval operation binding the contract event 0x420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb.
//
// Solidity: event CountIncremented(uint256 newCount)
func (_Token *TokenFilterer) FilterCountIncremented(opts *bind.FilterOpts) (*TokenCountIncrementedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "CountIncremented")
	if err != nil {
		return nil, err
	}
	return &TokenCountIncrementedIterator{contract: _Token.contract, event: "CountIncremented", logs: logs, sub: sub}, nil
}

// WatchCountIncremented is a free log subscription operation binding the contract event 0x420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb.
//
// Solidity: event CountIncremented(uint256 newCount)
func (_Token *TokenFilterer) WatchCountIncremented(opts *bind.WatchOpts, sink chan<- *TokenCountIncremented) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "CountIncremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCountIncremented)
				if err := _Token.contract.UnpackLog(event, "CountIncremented", log); err != nil {
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

// ParseCountIncremented is a log parse operation binding the contract event 0x420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb.
//
// Solidity: event CountIncremented(uint256 newCount)
func (_Token *TokenFilterer) ParseCountIncremented(log types.Log) (*TokenCountIncremented, error) {
	event := new(TokenCountIncremented)
	if err := _Token.contract.UnpackLog(event, "CountIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCountResetIterator is returned from FilterCountReset and is used to iterate over the raw logs and unpacked data for CountReset events raised by the Token contract.
type TokenCountResetIterator struct {
	Event *TokenCountReset // Event containing the contract specifics and raw log

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
func (it *TokenCountResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCountReset)
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
		it.Event = new(TokenCountReset)
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
func (it *TokenCountResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCountResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCountReset represents a CountReset event raised by the Token contract.
type TokenCountReset struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountReset is a free log retrieval operation binding the contract event 0x5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f0223977.
//
// Solidity: event CountReset(uint256 newCount)
func (_Token *TokenFilterer) FilterCountReset(opts *bind.FilterOpts) (*TokenCountResetIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "CountReset")
	if err != nil {
		return nil, err
	}
	return &TokenCountResetIterator{contract: _Token.contract, event: "CountReset", logs: logs, sub: sub}, nil
}

// WatchCountReset is a free log subscription operation binding the contract event 0x5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f0223977.
//
// Solidity: event CountReset(uint256 newCount)
func (_Token *TokenFilterer) WatchCountReset(opts *bind.WatchOpts, sink chan<- *TokenCountReset) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "CountReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCountReset)
				if err := _Token.contract.UnpackLog(event, "CountReset", log); err != nil {
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

// ParseCountReset is a log parse operation binding the contract event 0x5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f0223977.
//
// Solidity: event CountReset(uint256 newCount)
func (_Token *TokenFilterer) ParseCountReset(log types.Log) (*TokenCountReset, error) {
	event := new(TokenCountReset)
	if err := _Token.contract.UnpackLog(event, "CountReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
