// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve25519

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

// Curve25519ABI is the input ABI used to generate the binding from.
const Curve25519ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"privKey\",\"type\":\"uint256\"}],\"name\":\"publicKey\",\"outputs\":[{\"name\":\"qx\",\"type\":\"uint256\"},{\"name\":\"qy\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"privKey\",\"type\":\"uint256\"},{\"name\":\"pubX\",\"type\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\"}],\"name\":\"deriveKey\",\"outputs\":[{\"name\":\"qx\",\"type\":\"uint256\"},{\"name\":\"qy\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Curve25519Bin is the compiled bytecode used for deploying new contracts.
const Curve25519Bin = `608060405234801561001057600080fd5b506104c5806100206000396000f3fe608060405234801561001057600080fd5b5060043610610052577c010000000000000000000000000000000000000000000000000000000060003504638940aebe8114610057578063e241c1d91461008d575b600080fd5b6100746004803603602081101561006d57600080fd5b50356100b6565b6040805192835260208301919091528051918290030190f35b610074600480360360608110156100a357600080fd5b5080359060208101359060400135610120565b60008060008060006100ec8660097f20ae19a1b8a086b4e01edd2c7748d14c923d4d7e6d7c61b229e9c5a27eced3d9600161016c565b919450925090506100fc816101ef565b9050601360ff60020a038184099450601360ff60020a038183099350505050915091565b6000806000806000610135888888600161016c565b91945092509050610145816101ef565b9050601360ff60020a038184099450601360ff60020a038183099350505050935093915050565b600080808686868684806001861515610197575060009850889750600196506101e595505050505050565b86156101d95760018716156101bc576101b4838383898989610248565b919450925090505b6002870496506101cd8686866103dc565b91975095509350610197565b91985096509450505050505b9450945094915050565b6000806001601360ff60020a0384835b811561023d57818381151561021057fe5b04905083601360ff60020a0380868409601360ff60020a03038708909550935090918183029003906101ff565b509295945050505050565b60008080808080808c15801561025c57508b155b1561027357898989965096509650505050506103d0565b8915801561027f575088155b15610296578c8c8c965096509650505050506103d0565b898d1480156102a45750888c145b156102fa576102b58d8c8f8e6103fe565b90945092506102c88484600360016103fe565b90945092506102dd848462076d066001610425565b90945092506102f08c8c600260016103fe565b909250905061031d565b61030689898e8e61044e565b90945092506103178a898f8e61044e565b90925090505b61032984848484610473565b909450925061033a848481816103fe565b909750915061034b87838f8e61044e565b909750915061035c87838c8b61044e565b909750915061036d8d8c898561044e565b909650905061037e868286866103fe565b909650905061038f86828e8e61044e565b90965090508181146103c757601360ff60020a038188099650601360ff60020a038287099550601360ff60020a0381830994506103cb565b8194505b505050505b96509650969350505050565b60008060006103ef868686898989610248565b91989097509095509350505050565b600080601360ff60020a03848709601360ff60020a035b8487099097909650945050505050565b600080601360ff60020a0380868609601360ff60020a035b88860908601360ff60020a03610415565b600080601360ff60020a03808686601360ff60020a030309601360ff60020a0361043d565b600080601360ff60020a03838709601360ff60020a03858709909790965094505050505056fea165627a7a72305820d18c3ca39005a1c84edb9c69a932222e737a271bcaebdff55b9bf2c79a444d1a0029`

// DeployCurve25519 deploys a new Ethereum contract, binding an instance of Curve25519 to it.
func DeployCurve25519(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Curve25519, error) {
	parsed, err := abi.JSON(strings.NewReader(Curve25519ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Curve25519Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Curve25519{Curve25519Caller: Curve25519Caller{contract: contract}, Curve25519Transactor: Curve25519Transactor{contract: contract}, Curve25519Filterer: Curve25519Filterer{contract: contract}}, nil
}

// Curve25519 is an auto generated Go binding around an Ethereum contract.
type Curve25519 struct {
	Curve25519Caller     // Read-only binding to the contract
	Curve25519Transactor // Write-only binding to the contract
	Curve25519Filterer   // Log filterer for contract events
}

// Curve25519Caller is an auto generated read-only Go binding around an Ethereum contract.
type Curve25519Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Curve25519Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Curve25519Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Curve25519Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Curve25519Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Curve25519Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Curve25519Session struct {
	Contract     *Curve25519       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Curve25519CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Curve25519CallerSession struct {
	Contract *Curve25519Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Curve25519TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Curve25519TransactorSession struct {
	Contract     *Curve25519Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Curve25519Raw is an auto generated low-level Go binding around an Ethereum contract.
type Curve25519Raw struct {
	Contract *Curve25519 // Generic contract binding to access the raw methods on
}

// Curve25519CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Curve25519CallerRaw struct {
	Contract *Curve25519Caller // Generic read-only contract binding to access the raw methods on
}

// Curve25519TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Curve25519TransactorRaw struct {
	Contract *Curve25519Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCurve25519 creates a new instance of Curve25519, bound to a specific deployed contract.
func NewCurve25519(address common.Address, backend bind.ContractBackend) (*Curve25519, error) {
	contract, err := bindCurve25519(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curve25519{Curve25519Caller: Curve25519Caller{contract: contract}, Curve25519Transactor: Curve25519Transactor{contract: contract}, Curve25519Filterer: Curve25519Filterer{contract: contract}}, nil
}

// NewCurve25519Caller creates a new read-only instance of Curve25519, bound to a specific deployed contract.
func NewCurve25519Caller(address common.Address, caller bind.ContractCaller) (*Curve25519Caller, error) {
	contract, err := bindCurve25519(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Curve25519Caller{contract: contract}, nil
}

// NewCurve25519Transactor creates a new write-only instance of Curve25519, bound to a specific deployed contract.
func NewCurve25519Transactor(address common.Address, transactor bind.ContractTransactor) (*Curve25519Transactor, error) {
	contract, err := bindCurve25519(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Curve25519Transactor{contract: contract}, nil
}

// NewCurve25519Filterer creates a new log filterer instance of Curve25519, bound to a specific deployed contract.
func NewCurve25519Filterer(address common.Address, filterer bind.ContractFilterer) (*Curve25519Filterer, error) {
	contract, err := bindCurve25519(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Curve25519Filterer{contract: contract}, nil
}

// bindCurve25519 binds a generic wrapper to an already deployed contract.
func bindCurve25519(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Curve25519ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curve25519 *Curve25519Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Curve25519.Contract.Curve25519Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curve25519 *Curve25519Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve25519.Contract.Curve25519Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curve25519 *Curve25519Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curve25519.Contract.Curve25519Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curve25519 *Curve25519CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Curve25519.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curve25519 *Curve25519TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve25519.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curve25519 *Curve25519TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curve25519.Contract.contract.Transact(opts, method, params...)
}

// DeriveKey is a free data retrieval call binding the contract method 0xe241c1d9.
//
// Solidity: function deriveKey(uint256 privKey, uint256 pubX, uint256 pubY) constant returns(uint256 qx, uint256 qy)
func (_Curve25519 *Curve25519Caller) DeriveKey(opts *bind.CallOpts, privKey *big.Int, pubX *big.Int, pubY *big.Int) (struct {
	Qx *big.Int
	Qy *big.Int
}, error) {
	ret := new(struct {
		Qx *big.Int
		Qy *big.Int
	})
	out := ret
	err := _Curve25519.contract.Call(opts, out, "deriveKey", privKey, pubX, pubY)
	return *ret, err
}

// DeriveKey is a free data retrieval call binding the contract method 0xe241c1d9.
//
// Solidity: function deriveKey(uint256 privKey, uint256 pubX, uint256 pubY) constant returns(uint256 qx, uint256 qy)
func (_Curve25519 *Curve25519Session) DeriveKey(privKey *big.Int, pubX *big.Int, pubY *big.Int) (struct {
	Qx *big.Int
	Qy *big.Int
}, error) {
	return _Curve25519.Contract.DeriveKey(&_Curve25519.CallOpts, privKey, pubX, pubY)
}

// DeriveKey is a free data retrieval call binding the contract method 0xe241c1d9.
//
// Solidity: function deriveKey(uint256 privKey, uint256 pubX, uint256 pubY) constant returns(uint256 qx, uint256 qy)
func (_Curve25519 *Curve25519CallerSession) DeriveKey(privKey *big.Int, pubX *big.Int, pubY *big.Int) (struct {
	Qx *big.Int
	Qy *big.Int
}, error) {
	return _Curve25519.Contract.DeriveKey(&_Curve25519.CallOpts, privKey, pubX, pubY)
}

// PublicKey is a free data retrieval call binding the contract method 0x8940aebe.
//
// Solidity: function publicKey(uint256 privKey) constant returns(uint256 qx, uint256 qy)
func (_Curve25519 *Curve25519Caller) PublicKey(opts *bind.CallOpts, privKey *big.Int) (struct {
	Qx *big.Int
	Qy *big.Int
}, error) {
	ret := new(struct {
		Qx *big.Int
		Qy *big.Int
	})
	out := ret
	err := _Curve25519.contract.Call(opts, out, "publicKey", privKey)
	return *ret, err
}

// PublicKey is a free data retrieval call binding the contract method 0x8940aebe.
//
// Solidity: function publicKey(uint256 privKey) constant returns(uint256 qx, uint256 qy)
func (_Curve25519 *Curve25519Session) PublicKey(privKey *big.Int) (struct {
	Qx *big.Int
	Qy *big.Int
}, error) {
	return _Curve25519.Contract.PublicKey(&_Curve25519.CallOpts, privKey)
}

// PublicKey is a free data retrieval call binding the contract method 0x8940aebe.
//
// Solidity: function publicKey(uint256 privKey) constant returns(uint256 qx, uint256 qy)
func (_Curve25519 *Curve25519CallerSession) PublicKey(privKey *big.Int) (struct {
	Qx *big.Int
	Qy *big.Int
}, error) {
	return _Curve25519.Contract.PublicKey(&_Curve25519.CallOpts, privKey)
}
