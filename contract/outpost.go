// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package outpost

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OutpostABI is the input ABI used to generate the binding from.
const OutpostABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"pings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paidPings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"paidPing\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// OutpostBin is the compiled bytecode used for deploying new contracts.
const OutpostBin = `608060405260008055600060015534801561001957600080fd5b5060f8806100286000396000f3fe6080604052600436106056577c010000000000000000000000000000000000000000000000000000000060003504631e81ccb28114605b57806331de25e514607f5780635c36b186146091578063aebeefdf1460a5575b600080fd5b348015606657600080fd5b50606d60ab565b60408051918252519081900360200190f35b348015608a57600080fd5b50606d60b1565b348015609c57600080fd5b5060a360b7565b005b60a360c2565b60005481565b60015481565b600080546001019055565b600180548101905556fea165627a7a723058207b4bd5e0abaee29ceba300c23cdca96949737e5e56e842ed53ace64363f0a3220029`

// DeployOutpost deploys a new Ethereum contract, binding an instance of Outpost to it.
func DeployOutpost(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Outpost, error) {
	parsed, err := abi.JSON(strings.NewReader(OutpostABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutpostBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outpost{OutpostCaller: OutpostCaller{contract: contract}, OutpostTransactor: OutpostTransactor{contract: contract}, OutpostFilterer: OutpostFilterer{contract: contract}}, nil
}

// Outpost is an auto generated Go binding around an Ethereum contract.
type Outpost struct {
	OutpostCaller     // Read-only binding to the contract
	OutpostTransactor // Write-only binding to the contract
	OutpostFilterer   // Log filterer for contract events
}

// OutpostCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutpostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutpostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutpostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutpostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutpostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutpostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutpostSession struct {
	Contract     *Outpost          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutpostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutpostCallerSession struct {
	Contract *OutpostCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OutpostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutpostTransactorSession struct {
	Contract     *OutpostTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OutpostRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutpostRaw struct {
	Contract *Outpost // Generic contract binding to access the raw methods on
}

// OutpostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutpostCallerRaw struct {
	Contract *OutpostCaller // Generic read-only contract binding to access the raw methods on
}

// OutpostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutpostTransactorRaw struct {
	Contract *OutpostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutpost creates a new instance of Outpost, bound to a specific deployed contract.
func NewOutpost(address common.Address, backend bind.ContractBackend) (*Outpost, error) {
	contract, err := bindOutpost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outpost{OutpostCaller: OutpostCaller{contract: contract}, OutpostTransactor: OutpostTransactor{contract: contract}, OutpostFilterer: OutpostFilterer{contract: contract}}, nil
}

// NewOutpostCaller creates a new read-only instance of Outpost, bound to a specific deployed contract.
func NewOutpostCaller(address common.Address, caller bind.ContractCaller) (*OutpostCaller, error) {
	contract, err := bindOutpost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutpostCaller{contract: contract}, nil
}

// NewOutpostTransactor creates a new write-only instance of Outpost, bound to a specific deployed contract.
func NewOutpostTransactor(address common.Address, transactor bind.ContractTransactor) (*OutpostTransactor, error) {
	contract, err := bindOutpost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutpostTransactor{contract: contract}, nil
}

// NewOutpostFilterer creates a new log filterer instance of Outpost, bound to a specific deployed contract.
func NewOutpostFilterer(address common.Address, filterer bind.ContractFilterer) (*OutpostFilterer, error) {
	contract, err := bindOutpost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutpostFilterer{contract: contract}, nil
}

// bindOutpost binds a generic wrapper to an already deployed contract.
func bindOutpost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutpostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outpost *OutpostRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Outpost.Contract.OutpostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outpost *OutpostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outpost.Contract.OutpostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outpost *OutpostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outpost.Contract.OutpostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outpost *OutpostCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Outpost.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outpost *OutpostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outpost.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outpost *OutpostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outpost.Contract.contract.Transact(opts, method, params...)
}

// PaidPings is a free data retrieval call binding the contract method 0x31de25e5.
//
// Solidity: function paidPings() constant returns(uint256)
func (_Outpost *OutpostCaller) PaidPings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Outpost.contract.Call(opts, out, "paidPings")
	return *ret0, err
}

// PaidPings is a free data retrieval call binding the contract method 0x31de25e5.
//
// Solidity: function paidPings() constant returns(uint256)
func (_Outpost *OutpostSession) PaidPings() (*big.Int, error) {
	return _Outpost.Contract.PaidPings(&_Outpost.CallOpts)
}

// PaidPings is a free data retrieval call binding the contract method 0x31de25e5.
//
// Solidity: function paidPings() constant returns(uint256)
func (_Outpost *OutpostCallerSession) PaidPings() (*big.Int, error) {
	return _Outpost.Contract.PaidPings(&_Outpost.CallOpts)
}

// Pings is a free data retrieval call binding the contract method 0x1e81ccb2.
//
// Solidity: function pings() constant returns(uint256)
func (_Outpost *OutpostCaller) Pings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Outpost.contract.Call(opts, out, "pings")
	return *ret0, err
}

// Pings is a free data retrieval call binding the contract method 0x1e81ccb2.
//
// Solidity: function pings() constant returns(uint256)
func (_Outpost *OutpostSession) Pings() (*big.Int, error) {
	return _Outpost.Contract.Pings(&_Outpost.CallOpts)
}

// Pings is a free data retrieval call binding the contract method 0x1e81ccb2.
//
// Solidity: function pings() constant returns(uint256)
func (_Outpost *OutpostCallerSession) Pings() (*big.Int, error) {
	return _Outpost.Contract.Pings(&_Outpost.CallOpts)
}

// PaidPing is a paid mutator transaction binding the contract method 0xaebeefdf.
//
// Solidity: function paidPing() returns()
func (_Outpost *OutpostTransactor) PaidPing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outpost.contract.Transact(opts, "paidPing")
}

// PaidPing is a paid mutator transaction binding the contract method 0xaebeefdf.
//
// Solidity: function paidPing() returns()
func (_Outpost *OutpostSession) PaidPing() (*types.Transaction, error) {
	return _Outpost.Contract.PaidPing(&_Outpost.TransactOpts)
}

// PaidPing is a paid mutator transaction binding the contract method 0xaebeefdf.
//
// Solidity: function paidPing() returns()
func (_Outpost *OutpostTransactorSession) PaidPing() (*types.Transaction, error) {
	return _Outpost.Contract.PaidPing(&_Outpost.TransactOpts)
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_Outpost *OutpostTransactor) Ping(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outpost.contract.Transact(opts, "ping")
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_Outpost *OutpostSession) Ping() (*types.Transaction, error) {
	return _Outpost.Contract.Ping(&_Outpost.TransactOpts)
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_Outpost *OutpostTransactorSession) Ping() (*types.Transaction, error) {
	return _Outpost.Contract.Ping(&_Outpost.TransactOpts)
}
