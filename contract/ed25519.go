// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ed25519

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

// Ed25519ABI is the input ABI used to generate the binding from.
const Ed25519ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"scalarmult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Ed25519Bin is the compiled bytecode used for deploying new contracts.
const Ed25519Bin = `608060405234801561001057600080fd5b50610522806100206000396000f3fe608060405234801561001057600080fd5b5060043610610047577c01000000000000000000000000000000000000000000000000000000006000350463e49cf911811461004c575b600080fd5b6100696004803603602081101561006257600080fd5b5035610082565b6040805192835260208301919091528051918290030190f35b60008061008d61048e565b61009561048e565b7f216936d3cd6e53fec0a4e231fdd6dc5c692cc7609525a7b2c9562d608f25d51a82527f66666666666666666666666666666666666666666666666666666666666666586020808401919091526001604080850182905260008452918301819052908201525b600085111561013357846001166001141561011d5761011a8183610177565b90505b60029094049361012c826102f5565b91506100fb565b6000610142826040015161042c565b9050601360ff60020a0382518290098252601360ff60020a038183602001510960208301819052915194509092505050915091565b61017f61048e565b6101876104b0565b601360ff60020a0383604001518560400151098152601360ff60020a03815180096020820152601360ff60020a0383518551096040820152601360ff60020a0383602001518560200151096060820152601360ff60020a038082606001518360400151097f52036cee2b6ffe738cc740797779e89800700a4d4141d8ab75eb4dca135978a3096080820152601360ff60020a038160800151601360ff60020a030382602001510860a0820152601360ff60020a03816080015182602001510860c0820152601360ff60020a03806060830151601360ff60020a03908103906040850151601360ff60020a0390810390601360ff60020a0360208a01518a5108601360ff60020a0360208c01518c5108090808601360ff60020a0360a0840151845109098252601360ff60020a03808260400151836060015108601360ff60020a0360c0840151845109096020830152601360ff60020a038160c001518260a001510960408301525092915050565b6102fd61048e565b6103056104b0565b601360ff60020a0360208401518451088152601360ff60020a03815180096020820152601360ff60020a03835180096040820152601360ff60020a0360208401518009606082018190526040820151601360ff60020a03908103608084018190529091900860a0820152601360ff60020a036040840151800960e0820152601360ff60020a03808260e00151600209601360ff60020a03038260a001510860c0820152601360ff60020a0360c0820151601360ff60020a036060840151601360ff60020a03908103908560400151601360ff60020a030386602001510808098252601360ff60020a03808260600151601360ff60020a03038360800151088260a00151096020830152601360ff60020a038160c001518260a0015109604083015250919050565b6000806002601360ff60020a030390506000601360ff60020a03905060405160208152602080820152602060408201528460608201528260808201528160a082015260208160c0836005600019fa151561048557600080fd5b51949350505050565b6060604051908101604052806000815260200160008152602001600081525090565b610100604051908101604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152509056fea165627a7a723058201948ced63f3428e0046b54a4bf5ea23a66a5110a90c32660acd64acd5291b72b0029`

// DeployEd25519 deploys a new Ethereum contract, binding an instance of Ed25519 to it.
func DeployEd25519(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ed25519, error) {
	parsed, err := abi.JSON(strings.NewReader(Ed25519ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Ed25519Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ed25519{Ed25519Caller: Ed25519Caller{contract: contract}, Ed25519Transactor: Ed25519Transactor{contract: contract}, Ed25519Filterer: Ed25519Filterer{contract: contract}}, nil
}

// Ed25519 is an auto generated Go binding around an Ethereum contract.
type Ed25519 struct {
	Ed25519Caller     // Read-only binding to the contract
	Ed25519Transactor // Write-only binding to the contract
	Ed25519Filterer   // Log filterer for contract events
}

// Ed25519Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ed25519Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ed25519Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ed25519Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ed25519Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ed25519Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ed25519Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ed25519Session struct {
	Contract     *Ed25519          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ed25519CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ed25519CallerSession struct {
	Contract *Ed25519Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Ed25519TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ed25519TransactorSession struct {
	Contract     *Ed25519Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Ed25519Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ed25519Raw struct {
	Contract *Ed25519 // Generic contract binding to access the raw methods on
}

// Ed25519CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ed25519CallerRaw struct {
	Contract *Ed25519Caller // Generic read-only contract binding to access the raw methods on
}

// Ed25519TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ed25519TransactorRaw struct {
	Contract *Ed25519Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEd25519 creates a new instance of Ed25519, bound to a specific deployed contract.
func NewEd25519(address common.Address, backend bind.ContractBackend) (*Ed25519, error) {
	contract, err := bindEd25519(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ed25519{Ed25519Caller: Ed25519Caller{contract: contract}, Ed25519Transactor: Ed25519Transactor{contract: contract}, Ed25519Filterer: Ed25519Filterer{contract: contract}}, nil
}

// NewEd25519Caller creates a new read-only instance of Ed25519, bound to a specific deployed contract.
func NewEd25519Caller(address common.Address, caller bind.ContractCaller) (*Ed25519Caller, error) {
	contract, err := bindEd25519(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ed25519Caller{contract: contract}, nil
}

// NewEd25519Transactor creates a new write-only instance of Ed25519, bound to a specific deployed contract.
func NewEd25519Transactor(address common.Address, transactor bind.ContractTransactor) (*Ed25519Transactor, error) {
	contract, err := bindEd25519(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ed25519Transactor{contract: contract}, nil
}

// NewEd25519Filterer creates a new log filterer instance of Ed25519, bound to a specific deployed contract.
func NewEd25519Filterer(address common.Address, filterer bind.ContractFilterer) (*Ed25519Filterer, error) {
	contract, err := bindEd25519(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ed25519Filterer{contract: contract}, nil
}

// bindEd25519 binds a generic wrapper to an already deployed contract.
func bindEd25519(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ed25519ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ed25519 *Ed25519Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ed25519.Contract.Ed25519Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ed25519 *Ed25519Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ed25519.Contract.Ed25519Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ed25519 *Ed25519Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ed25519.Contract.Ed25519Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ed25519 *Ed25519CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ed25519.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ed25519 *Ed25519TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ed25519.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ed25519 *Ed25519TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ed25519.Contract.contract.Transact(opts, method, params...)
}

// Scalarmult is a free data retrieval call binding the contract method 0xe49cf911.
//
// Solidity: function scalarmult(uint256 s) constant returns(uint256, uint256)
func (_Ed25519 *Ed25519Caller) Scalarmult(opts *bind.CallOpts, s *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Ed25519.contract.Call(opts, out, "scalarmult", s)
	return *ret0, *ret1, err
}

// Scalarmult is a free data retrieval call binding the contract method 0xe49cf911.
//
// Solidity: function scalarmult(uint256 s) constant returns(uint256, uint256)
func (_Ed25519 *Ed25519Session) Scalarmult(s *big.Int) (*big.Int, *big.Int, error) {
	return _Ed25519.Contract.Scalarmult(&_Ed25519.CallOpts, s)
}

// Scalarmult is a free data retrieval call binding the contract method 0xe49cf911.
//
// Solidity: function scalarmult(uint256 s) constant returns(uint256, uint256)
func (_Ed25519 *Ed25519CallerSession) Scalarmult(s *big.Int) (*big.Int, *big.Int, error) {
	return _Ed25519.Contract.Scalarmult(&_Ed25519.CallOpts, s)
}
