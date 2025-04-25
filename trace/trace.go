// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

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

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLock.Op\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"Assert\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Locked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumLock.Op\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"used\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"withdrawl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610927806100206000396000f3fe60806040526004361061003f5760003560e01c8063137658381461004457806331672eeb1461006057806344427e3014610089578063aa17ab1b146100c6575b600080fd5b61005e60048036038101906100599190610501565b610103565b005b34801561006c57600080fd5b506100876004803603810190610082919061052e565b610291565b005b34801561009557600080fd5b506100b060048036038101906100ab91906105cc565b610475565b6040516100bd9190610612565b60405180910390f35b3480156100d257600080fd5b506100ed60048036038101906100e89190610652565b61048d565b6040516100fa91906106ad565b60405180910390f35b60003490506000339050816000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461015b91906106f7565b92505081905550600160008060018111156101795761017861072b565b5b600181111561018b5761018a61072b565b5b815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156101d657600080fd5b60018060008060018111156101ee576101ed61072b565b5b6001811115610200576101ff61072b565b5b815260200190815260200160002060008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f87d092cc8ff72d1243337cb7a4e178708d98aad3ff50ff5ec6abc6457239f70381848443426000604051610284969594939291906107c0565b60405180910390a1505050565b60003390508267ffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102ee9190610821565b925050819055508073ffffffffffffffffffffffffffffffffffffffff166108fc8467ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015610345573d6000803e3d6000fd5b506001600060018081111561035d5761035c61072b565b5b600181111561036f5761036e61072b565b5b815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156103ba57600080fd5b60018060006001808111156103d2576103d161072b565b5b60018111156103e4576103e361072b565b5b815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f87d092cc8ff72d1243337cb7a4e178708d98aad3ff50ff5ec6abc6457239f7038183854342600160405161046896959493929190610890565b60405180910390a1505050565b60006020528060005260406000206000915090505481565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600080fd5b600067ffffffffffffffff82169050919050565b6104de816104c1565b81146104e957600080fd5b50565b6000813590506104fb816104d5565b92915050565b600060208284031215610517576105166104bc565b5b6000610525848285016104ec565b91505092915050565b60008060408385031215610545576105446104bc565b5b6000610553858286016104ec565b9250506020610564858286016104ec565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105998261056e565b9050919050565b6105a98161058e565b81146105b457600080fd5b50565b6000813590506105c6816105a0565b92915050565b6000602082840312156105e2576105e16104bc565b5b60006105f0848285016105b7565b91505092915050565b6000819050919050565b61060c816105f9565b82525050565b60006020820190506106276000830184610603565b92915050565b6002811061063a57600080fd5b50565b60008135905061064c8161062d565b92915050565b60008060408385031215610669576106686104bc565b5b60006106778582860161063d565b9250506020610688858286016104ec565b9150509250929050565b60008115159050919050565b6106a781610692565b82525050565b60006020820190506106c2600083018461069e565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610702826105f9565b915061070d836105f9565b9250828201905080821115610725576107246106c8565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6107638161058e565b82525050565b610772816104c1565b82525050565b600281106107895761078861072b565b5b50565b600081905061079a82610778565b919050565b60006107aa8261078c565b9050919050565b6107ba8161079f565b82525050565b600060c0820190506107d5600083018961075a565b6107e26020830188610769565b6107ef6040830187610603565b6107fc6060830186610603565b6108096080830185610603565b61081660a08301846107b1565b979650505050505050565b600061082c826105f9565b9150610837836105f9565b925082820390508181111561084f5761084e6106c8565b5b92915050565b6000819050919050565b600061087a610875610870846104c1565b610855565b6105f9565b9050919050565b61088a8161085f565b82525050565b600060c0820190506108a5600083018961075a565b6108b26020830188610769565b6108bf6040830187610881565b6108cc6060830186610603565b6108d96080830185610603565b6108e660a08301846107b1565b97965050505050505056fea26469706673582212204fb8ecaf3c1995f1a2e4fbe32029b464df8672f46941d6e54770ac56b345710f64736f6c63430008130033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// Locked is a free data retrieval call binding the contract method 0x44427e30.
//
// Solidity: function Locked(address ) view returns(uint256)
func (_Trace *TraceCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "Locked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0x44427e30.
//
// Solidity: function Locked(address ) view returns(uint256)
func (_Trace *TraceSession) Locked(arg0 common.Address) (*big.Int, error) {
	return _Trace.Contract.Locked(&_Trace.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0x44427e30.
//
// Solidity: function Locked(address ) view returns(uint256)
func (_Trace *TraceCallerSession) Locked(arg0 common.Address) (*big.Int, error) {
	return _Trace.Contract.Locked(&_Trace.CallOpts, arg0)
}

// Used is a free data retrieval call binding the contract method 0xaa17ab1b.
//
// Solidity: function used(uint8 , uint64 ) view returns(bool)
func (_Trace *TraceCaller) Used(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (bool, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "used", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Used is a free data retrieval call binding the contract method 0xaa17ab1b.
//
// Solidity: function used(uint8 , uint64 ) view returns(bool)
func (_Trace *TraceSession) Used(arg0 uint8, arg1 uint64) (bool, error) {
	return _Trace.Contract.Used(&_Trace.CallOpts, arg0, arg1)
}

// Used is a free data retrieval call binding the contract method 0xaa17ab1b.
//
// Solidity: function used(uint8 , uint64 ) view returns(bool)
func (_Trace *TraceCallerSession) Used(arg0 uint8, arg1 uint64) (bool, error) {
	return _Trace.Contract.Used(&_Trace.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x13765838.
//
// Solidity: function deposit(uint64 nonce) payable returns()
func (_Trace *TraceTransactor) Deposit(opts *bind.TransactOpts, nonce uint64) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "deposit", nonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x13765838.
//
// Solidity: function deposit(uint64 nonce) payable returns()
func (_Trace *TraceSession) Deposit(nonce uint64) (*types.Transaction, error) {
	return _Trace.Contract.Deposit(&_Trace.TransactOpts, nonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x13765838.
//
// Solidity: function deposit(uint64 nonce) payable returns()
func (_Trace *TraceTransactorSession) Deposit(nonce uint64) (*types.Transaction, error) {
	return _Trace.Contract.Deposit(&_Trace.TransactOpts, nonce)
}

// Withdrawl is a paid mutator transaction binding the contract method 0x31672eeb.
//
// Solidity: function withdrawl(uint64 amount, uint64 nonce) returns()
func (_Trace *TraceTransactor) Withdrawl(opts *bind.TransactOpts, amount uint64, nonce uint64) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "withdrawl", amount, nonce)
}

// Withdrawl is a paid mutator transaction binding the contract method 0x31672eeb.
//
// Solidity: function withdrawl(uint64 amount, uint64 nonce) returns()
func (_Trace *TraceSession) Withdrawl(amount uint64, nonce uint64) (*types.Transaction, error) {
	return _Trace.Contract.Withdrawl(&_Trace.TransactOpts, amount, nonce)
}

// Withdrawl is a paid mutator transaction binding the contract method 0x31672eeb.
//
// Solidity: function withdrawl(uint64 amount, uint64 nonce) returns()
func (_Trace *TraceTransactorSession) Withdrawl(amount uint64, nonce uint64) (*types.Transaction, error) {
	return _Trace.Contract.Withdrawl(&_Trace.TransactOpts, amount, nonce)
}

// TraceAssertIterator is returned from FilterAssert and is used to iterate over the raw logs and unpacked data for Assert events raised by the Trace contract.
type TraceAssertIterator struct {
	Event *TraceAssert // Event containing the contract specifics and raw log

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
func (it *TraceAssertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraceAssert)
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
		it.Event = new(TraceAssert)
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
func (it *TraceAssertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraceAssertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraceAssert represents a Assert event raised by the Trace contract.
type TraceAssert struct {
	From      common.Address
	Nonce     uint64
	Amount    *big.Int
	Blocknum  *big.Int
	Timestamp *big.Int
	Op        uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssert is a free log retrieval operation binding the contract event 0x87d092cc8ff72d1243337cb7a4e178708d98aad3ff50ff5ec6abc6457239f703.
//
// Solidity: event Assert(address from, uint64 nonce, uint256 amount, uint256 blocknum, uint256 timestamp, uint8 op)
func (_Trace *TraceFilterer) FilterAssert(opts *bind.FilterOpts) (*TraceAssertIterator, error) {

	logs, sub, err := _Trace.contract.FilterLogs(opts, "Assert")
	if err != nil {
		return nil, err
	}
	return &TraceAssertIterator{contract: _Trace.contract, event: "Assert", logs: logs, sub: sub}, nil
}

// WatchAssert is a free log subscription operation binding the contract event 0x87d092cc8ff72d1243337cb7a4e178708d98aad3ff50ff5ec6abc6457239f703.
//
// Solidity: event Assert(address from, uint64 nonce, uint256 amount, uint256 blocknum, uint256 timestamp, uint8 op)
func (_Trace *TraceFilterer) WatchAssert(opts *bind.WatchOpts, sink chan<- *TraceAssert) (event.Subscription, error) {

	logs, sub, err := _Trace.contract.WatchLogs(opts, "Assert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraceAssert)
				if err := _Trace.contract.UnpackLog(event, "Assert", log); err != nil {
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

// ParseAssert is a log parse operation binding the contract event 0x87d092cc8ff72d1243337cb7a4e178708d98aad3ff50ff5ec6abc6457239f703.
//
// Solidity: event Assert(address from, uint64 nonce, uint256 amount, uint256 blocknum, uint256 timestamp, uint8 op)
func (_Trace *TraceFilterer) ParseAssert(log types.Log) (*TraceAssert, error) {
	event := new(TraceAssert)
	if err := _Trace.contract.UnpackLog(event, "Assert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
