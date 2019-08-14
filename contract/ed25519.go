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
const Ed25519ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"scalarMultBase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Ed25519Bin is the compiled bytecode used for deploying new contracts.
var Ed25519Bin = "0x608060405234801561001057600080fd5b5061052d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c4f4912b14610030575b600080fd5b61004d6004803603602081101561004657600080fd5b5035610066565b6040805192835260208301919091528051918290030190f35b600080610071610492565b610079610492565b7f216936d3cd6e53fec0a4e231fdd6dc5c692cc7609525a7b2c9562d608f25d51a82527f66666666666666666666666666666666666666666666666666666666666666586020808401919091526001604080850182905260008452918301819052908201525b84156101155784600116600114156100fe576100fb8183610159565b90505b600185901c945061010e826102ef565b91506100df565b60006101248260400151610432565b90506013600160ff1b03825182900982526013600160ff1b038183602001510960208301819052915194509092505050915091565b610161610492565b6101696104b3565b6013600160ff1b03836040015185604001510981526013600160ff1b038151800960208201526013600160ff1b03835185510960408201526013600160ff1b03836020015185602001510960608201526013600160ff1b038082606001518360400151097f52036cee2b6ffe738cc740797779e89800700a4d4141d8ab75eb4dca135978a30960808201526013600160ff1b0381608001516013600160ff1b030382602001510860a08201526013600160ff1b03816080015182602001510860c08201526013600160ff1b038082606001516013600160ff1b03036013600160ff1b038061025357fe5b84604001516013600160ff1b03036013600160ff1b038061027057fe5b6013600160ff1b0360208a01518a51086013600160ff1b0360208c01518c51080908086013600160ff1b0360a08401518451090982526013600160ff1b038082604001518360600151086013600160ff1b0360c08401518451090960208301526013600160ff1b038160c001518260a001510960408301525092915050565b6102f7610492565b6102ff6104b3565b6013600160ff1b03602084015184510881526013600160ff1b038151800960208201526013600160ff1b038351800960408201526013600160ff1b03602084015180096060820181905260408201516013600160ff1b03908103608084018190529091900860a08201526013600160ff1b036040840151800960e08201526013600160ff1b03808260e001516002096013600160ff1b03038260a001510860c08201526013600160ff1b0360c08201516013600160ff1b0383606001516013600160ff1b03036013600160ff1b03806103d457fe5b85604001516013600160ff1b0303866020015108080982526013600160ff1b038082606001516013600160ff1b03038360800151088260a001510960208301526013600160ff1b038160c001518260a0015109604083015250919050565b60008060026013600160ff1b0303905060006013600160ff1b03905060405160208152602080820152602060408201528460608201528260808201528160a082015260208160c0836005600019fa61048957600080fd5b51949350505050565b60405180606001604052806000815260200160008152602001600081525090565b6040518061010001604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152509056fea265627a7a72305820cde07ce38f4d43ec416cdca6f4acbc8967d408616f9433a7bd48450aa5f9eab664736f6c634300050a0032"

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

// ScalarMultBase is a free data retrieval call binding the contract method 0xc4f4912b.
//
// Solidity: function scalarMultBase(uint256 s) constant returns(uint256, uint256)
func (_Ed25519 *Ed25519Caller) ScalarMultBase(opts *bind.CallOpts, s *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Ed25519.contract.Call(opts, out, "scalarMultBase", s)
	return *ret0, *ret1, err
}

// ScalarMultBase is a free data retrieval call binding the contract method 0xc4f4912b.
//
// Solidity: function scalarMultBase(uint256 s) constant returns(uint256, uint256)
func (_Ed25519 *Ed25519Session) ScalarMultBase(s *big.Int) (*big.Int, *big.Int, error) {
	return _Ed25519.Contract.ScalarMultBase(&_Ed25519.CallOpts, s)
}

// ScalarMultBase is a free data retrieval call binding the contract method 0xc4f4912b.
//
// Solidity: function scalarMultBase(uint256 s) constant returns(uint256, uint256)
func (_Ed25519 *Ed25519CallerSession) ScalarMultBase(s *big.Int) (*big.Int, *big.Int, error) {
	return _Ed25519.Contract.ScalarMultBase(&_Ed25519.CallOpts, s)
}
