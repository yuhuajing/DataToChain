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

// WorkRecordWorkDetails is an auto generated low-level Go binding around an user-defined struct.
type WorkRecordWorkDetails struct {
	WorkLocation string
	WorkContent  string
	Person       string
	Advice       string
	WorkTime     string
	Remarks      string
	ImagesUrl    string
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workLocation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workContent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"person\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"advice\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imagesUrl\",\"type\":\"string\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"}],\"name\":\"getRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"workLocation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workContent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"person\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"advice\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imagesUrl\",\"type\":\"string\"}],\"internalType\":\"structWorkRecord.WorkDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610c16806100a56000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806311dd88451461005c578063715018a6146100855780638da5cb5b1461008f578063abafeb1e146100aa578063f2fde38b146100bd575b600080fd5b61006f61006a3660046107a8565b6100d0565b60405161007c9190610835565b60405180910390f35b61008d61053e565b005b6000546040516001600160a01b03909116815260200161007c565b61008d6100b836600461090b565b610552565b61008d6100cb366004610a4b565b610645565b6101106040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b6001826040516101209190610a7b565b90815260200160405180910390206040518060e001604052908160008201805461014990610a97565b80601f016020809104026020016040519081016040528092919081815260200182805461017590610a97565b80156101c25780601f10610197576101008083540402835291602001916101c2565b820191906000526020600020905b8154815290600101906020018083116101a557829003601f168201915b505050505081526020016001820180546101db90610a97565b80601f016020809104026020016040519081016040528092919081815260200182805461020790610a97565b80156102545780601f1061022957610100808354040283529160200191610254565b820191906000526020600020905b81548152906001019060200180831161023757829003601f168201915b5050505050815260200160028201805461026d90610a97565b80601f016020809104026020016040519081016040528092919081815260200182805461029990610a97565b80156102e65780601f106102bb576101008083540402835291602001916102e6565b820191906000526020600020905b8154815290600101906020018083116102c957829003601f168201915b505050505081526020016003820180546102ff90610a97565b80601f016020809104026020016040519081016040528092919081815260200182805461032b90610a97565b80156103785780601f1061034d57610100808354040283529160200191610378565b820191906000526020600020905b81548152906001019060200180831161035b57829003601f168201915b5050505050815260200160048201805461039190610a97565b80601f01602080910402602001604051908101604052809291908181526020018280546103bd90610a97565b801561040a5780601f106103df5761010080835404028352916020019161040a565b820191906000526020600020905b8154815290600101906020018083116103ed57829003601f168201915b5050505050815260200160058201805461042390610a97565b80601f016020809104026020016040519081016040528092919081815260200182805461044f90610a97565b801561049c5780601f106104715761010080835404028352916020019161049c565b820191906000526020600020905b81548152906001019060200180831161047f57829003601f168201915b505050505081526020016006820180546104b590610a97565b80601f01602080910402602001604051908101604052809291908181526020018280546104e190610a97565b801561052e5780601f106105035761010080835404028352916020019161052e565b820191906000526020600020905b81548152906001019060200180831161051157829003601f168201915b5050505050815250509050919050565b610546610688565b61055060006106b5565b565b61055a610688565b6040518060e001604052808881526020018781526020018681526020018581526020018481526020018381526020018281525060018960405161059d9190610a7b565b908152604051908190036020019020815181906105ba9082610b20565b50602082015160018201906105cf9082610b20565b50604082015160028201906105e49082610b20565b50606082015160038201906105f99082610b20565b506080820151600482019061060e9082610b20565b5060a082015160058201906106239082610b20565b5060c082015160068201906106389082610b20565b5050505050505050505050565b61064d610688565b6001600160a01b03811661067c57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b610685816106b5565b50565b6000546001600160a01b031633146105505760405163118cdaa760e01b8152336004820152602401610673565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261072c57600080fd5b813567ffffffffffffffff8082111561074757610747610705565b604051601f8301601f19908116603f0116810190828211818310171561076f5761076f610705565b8160405283815286602085880101111561078857600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156107ba57600080fd5b813567ffffffffffffffff8111156107d157600080fd5b6107dd8482850161071b565b949350505050565b60005b838110156108005781810151838201526020016107e8565b50506000910152565b600081518084526108218160208601602086016107e5565b601f01601f19169290920160200192915050565b602081526000825160e06020840152610852610100840182610809565b90506020840151601f19808584030160408601526108708383610809565b9250604086015191508085840301606086015261088d8383610809565b925060608601519150808584030160808601526108aa8383610809565b925060808601519150808584030160a08601526108c78383610809565b925060a08601519150808584030160c08601526108e48383610809565b925060c08601519150808584030160e0860152506109028282610809565b95945050505050565b600080600080600080600080610100898b03121561092857600080fd5b883567ffffffffffffffff8082111561094057600080fd5b61094c8c838d0161071b565b995060208b013591508082111561096257600080fd5b61096e8c838d0161071b565b985060408b013591508082111561098457600080fd5b6109908c838d0161071b565b975060608b01359150808211156109a657600080fd5b6109b28c838d0161071b565b965060808b01359150808211156109c857600080fd5b6109d48c838d0161071b565b955060a08b01359150808211156109ea57600080fd5b6109f68c838d0161071b565b945060c08b0135915080821115610a0c57600080fd5b610a188c838d0161071b565b935060e08b0135915080821115610a2e57600080fd5b50610a3b8b828c0161071b565b9150509295985092959890939650565b600060208284031215610a5d57600080fd5b81356001600160a01b0381168114610a7457600080fd5b9392505050565b60008251610a8d8184602087016107e5565b9190910192915050565b600181811c90821680610aab57607f821691505b602082108103610acb57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610b1b57600081815260208120601f850160051c81016020861015610af85750805b601f850160051c820191505b81811015610b1757828155600101610b04565b5050505b505050565b815167ffffffffffffffff811115610b3a57610b3a610705565b610b4e81610b488454610a97565b84610ad1565b602080601f831160018114610b835760008415610b6b5750858301515b600019600386901b1c1916600185901b178555610b17565b600085815260208120601f198616915b82811015610bb257888601518255948401946001909101908401610b93565b5085821015610bd05787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212208ad206f52f6f71c5bfa11169bff4a8caf78d92b9bfbfbb13b90cad1b9880640b64736f6c63430008140033",
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

// GetRecord is a free data retrieval call binding the contract method 0x11dd8845.
//
// Solidity: function getRecord(string number) view returns((string,string,string,string,string,string,string))
func (_Trace *TraceCaller) GetRecord(opts *bind.CallOpts, number string) (WorkRecordWorkDetails, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "getRecord", number)

	if err != nil {
		return *new(WorkRecordWorkDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkRecordWorkDetails)).(*WorkRecordWorkDetails)

	return out0, err

}

// GetRecord is a free data retrieval call binding the contract method 0x11dd8845.
//
// Solidity: function getRecord(string number) view returns((string,string,string,string,string,string,string))
func (_Trace *TraceSession) GetRecord(number string) (WorkRecordWorkDetails, error) {
	return _Trace.Contract.GetRecord(&_Trace.CallOpts, number)
}

// GetRecord is a free data retrieval call binding the contract method 0x11dd8845.
//
// Solidity: function getRecord(string number) view returns((string,string,string,string,string,string,string))
func (_Trace *TraceCallerSession) GetRecord(number string) (WorkRecordWorkDetails, error) {
	return _Trace.Contract.GetRecord(&_Trace.CallOpts, number)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceSession) Owner() (common.Address, error) {
	return _Trace.Contract.Owner(&_Trace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceCallerSession) Owner() (common.Address, error) {
	return _Trace.Contract.Owner(&_Trace.CallOpts)
}

// AddRecord is a paid mutator transaction binding the contract method 0xabafeb1e.
//
// Solidity: function addRecord(string number, string workLocation, string workContent, string person, string advice, string workTime, string remarks, string imagesUrl) returns()
func (_Trace *TraceTransactor) AddRecord(opts *bind.TransactOpts, number string, workLocation string, workContent string, person string, advice string, workTime string, remarks string, imagesUrl string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addRecord", number, workLocation, workContent, person, advice, workTime, remarks, imagesUrl)
}

// AddRecord is a paid mutator transaction binding the contract method 0xabafeb1e.
//
// Solidity: function addRecord(string number, string workLocation, string workContent, string person, string advice, string workTime, string remarks, string imagesUrl) returns()
func (_Trace *TraceSession) AddRecord(number string, workLocation string, workContent string, person string, advice string, workTime string, remarks string, imagesUrl string) (*types.Transaction, error) {
	return _Trace.Contract.AddRecord(&_Trace.TransactOpts, number, workLocation, workContent, person, advice, workTime, remarks, imagesUrl)
}

// AddRecord is a paid mutator transaction binding the contract method 0xabafeb1e.
//
// Solidity: function addRecord(string number, string workLocation, string workContent, string person, string advice, string workTime, string remarks, string imagesUrl) returns()
func (_Trace *TraceTransactorSession) AddRecord(number string, workLocation string, workContent string, person string, advice string, workTime string, remarks string, imagesUrl string) (*types.Transaction, error) {
	return _Trace.Contract.AddRecord(&_Trace.TransactOpts, number, workLocation, workContent, person, advice, workTime, remarks, imagesUrl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trace *TraceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trace *TraceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trace.Contract.RenounceOwnership(&_Trace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trace *TraceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trace.Contract.RenounceOwnership(&_Trace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trace *TraceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trace *TraceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trace.Contract.TransferOwnership(&_Trace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trace *TraceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trace.Contract.TransferOwnership(&_Trace.TransactOpts, newOwner)
}

// TraceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trace contract.
type TraceOwnershipTransferredIterator struct {
	Event *TraceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TraceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraceOwnershipTransferred)
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
		it.Event = new(TraceOwnershipTransferred)
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
func (it *TraceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraceOwnershipTransferred represents a OwnershipTransferred event raised by the Trace contract.
type TraceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trace *TraceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TraceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TraceOwnershipTransferredIterator{contract: _Trace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trace *TraceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TraceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraceOwnershipTransferred)
				if err := _Trace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trace *TraceFilterer) ParseOwnershipTransferred(log types.Log) (*TraceOwnershipTransferred, error) {
	event := new(TraceOwnershipTransferred)
	if err := _Trace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
