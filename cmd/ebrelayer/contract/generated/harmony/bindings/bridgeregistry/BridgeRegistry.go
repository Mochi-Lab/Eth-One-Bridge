// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeRegistry

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/harmony-one/harmony/accounts/abi"
	"github.com/harmony-one/harmony/accounts/abi/bind"
	"github.com/harmony-one/harmony/core/types"
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

// BridgeRegistryABI is the input ABI used to generate the binding from.
const BridgeRegistryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"HmyLogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ethereumBridge\",\"type\":\"address\"}],\"name\":\"HmyLogHarmonyBridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"HmyLogOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"}],\"name\":\"HmyLogValsetSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethereumBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBridgeBank\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthereumBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumBridge\",\"type\":\"address\"}],\"name\":\"setEthereumBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"}],\"name\":\"setValset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeRegistryBin is the compiled bytecode used for deploying new contracts.
var BridgeRegistryBin = "0x608060405234801561001057600080fd5b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c90806100616000396000f3fe608060405234801561001057600080fd5b50600436106100e95760003560e01c80637f54af0c1161008c578063960bc2ee11610066578063960bc2ee146103ba5780639fd168ec14610404578063d33f397e1461044e578063e7f43c6814610498576100e9565b80637f54af0c146102e2578063814c92c31461032c578063833b1fce14610370576100e9565b80633d051a9e116100c85780633d051a9e146101c6578063570ca7351461020a5780637adbf973146102545780637dc0d1d014610298576100e9565b80625e09ca146100ee5780630e41f373146101325780631cf86a631461017c575b600080fd5b6101306004803603602081101561010457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104e2565b005b61013a61064b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610184610671565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610208600480360360208110156101dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610696565b005b610212610800565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102966004803603602081101561026a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610826565b005b6102a06109b2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102ea6109d8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61036e6004803603602081101561034257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109fe565b005b610378610b8a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103c2610bb4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61040c610bdd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610456610c07565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104a0610c31565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105a5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd3869bed92142314dc7f6cd8e15f19f72f14034c42e74a8b170d39b42fa1b4ac81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610759576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f31a3f52a07e4eaa601f0c839a0b72df88e5a49ae52a6aefcf04eeac96c72f9aa81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fad9ae01008c260a11bf3fa6e389527b6a3a3fb1ed015c1937a9087d7c6dbc18d600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fa4ce696f3b5c3ecec852cb66da1f78522d9a3b5a98f2e3428eeca4ce0931df86600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509056fea265627a7a723158200a9ab13185d8dd5a0dd7a0c6c85926db51f2887285547d1e1459e40a2124e7cd64736f6c63430005110032"

// DeployBridgeRegistry deploys a new Ethereum contract, binding an instance of BridgeRegistry to it.
func DeployBridgeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeRegistry{BridgeRegistryCaller: BridgeRegistryCaller{contract: contract}, BridgeRegistryTransactor: BridgeRegistryTransactor{contract: contract}, BridgeRegistryFilterer: BridgeRegistryFilterer{contract: contract}}, nil
}

// BridgeRegistry is an auto generated Go binding around an Ethereum contract.
type BridgeRegistry struct {
	BridgeRegistryCaller     // Read-only binding to the contract
	BridgeRegistryTransactor // Write-only binding to the contract
	BridgeRegistryFilterer   // Log filterer for contract events
}

// BridgeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRegistrySession struct {
	Contract     *BridgeRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRegistryCallerSession struct {
	Contract *BridgeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRegistryTransactorSession struct {
	Contract     *BridgeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRegistryRaw struct {
	Contract *BridgeRegistry // Generic contract binding to access the raw methods on
}

// BridgeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRegistryCallerRaw struct {
	Contract *BridgeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRegistryTransactorRaw struct {
	Contract *BridgeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeRegistry creates a new instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistry(address common.Address, backend bind.ContractBackend) (*BridgeRegistry, error) {
	contract, err := bindBridgeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistry{BridgeRegistryCaller: BridgeRegistryCaller{contract: contract}, BridgeRegistryTransactor: BridgeRegistryTransactor{contract: contract}, BridgeRegistryFilterer: BridgeRegistryFilterer{contract: contract}}, nil
}

// NewBridgeRegistryCaller creates a new read-only instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryCaller(address common.Address, caller bind.ContractCaller) (*BridgeRegistryCaller, error) {
	contract, err := bindBridgeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryCaller{contract: contract}, nil
}

// NewBridgeRegistryTransactor creates a new write-only instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRegistryTransactor, error) {
	contract, err := bindBridgeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryTransactor{contract: contract}, nil
}

// NewBridgeRegistryFilterer creates a new log filterer instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRegistryFilterer, error) {
	contract, err := bindBridgeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryFilterer{contract: contract}, nil
}

// bindBridgeRegistry binds a generic wrapper to an already deployed contract.
func bindBridgeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRegistry *BridgeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeRegistry.Contract.BridgeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRegistry *BridgeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.BridgeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRegistry *BridgeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.BridgeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRegistry *BridgeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRegistry *BridgeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRegistry *BridgeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) BridgeBank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeBank(&_BridgeRegistry.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) BridgeBank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeBank(&_BridgeRegistry.CallOpts)
}

// EthereumBridge is a free data retrieval call binding the contract method 0x1cf86a63.
//
// Solidity: function ethereumBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) EthereumBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "ethereumBridge")
	return *ret0, err
}

// EthereumBridge is a free data retrieval call binding the contract method 0x1cf86a63.
//
// Solidity: function ethereumBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) EthereumBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.EthereumBridge(&_BridgeRegistry.CallOpts)
}

// EthereumBridge is a free data retrieval call binding the contract method 0x1cf86a63.
//
// Solidity: function ethereumBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) EthereumBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.EthereumBridge(&_BridgeRegistry.CallOpts)
}

// GetBridgeBank is a free data retrieval call binding the contract method 0xd33f397e.
//
// Solidity: function getBridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) GetBridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "getBridgeBank")
	return *ret0, err
}

// GetBridgeBank is a free data retrieval call binding the contract method 0xd33f397e.
//
// Solidity: function getBridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) GetBridgeBank() (common.Address, error) {
	return _BridgeRegistry.Contract.GetBridgeBank(&_BridgeRegistry.CallOpts)
}

// GetBridgeBank is a free data retrieval call binding the contract method 0xd33f397e.
//
// Solidity: function getBridgeBank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) GetBridgeBank() (common.Address, error) {
	return _BridgeRegistry.Contract.GetBridgeBank(&_BridgeRegistry.CallOpts)
}

// GetEthereumBridge is a free data retrieval call binding the contract method 0x960bc2ee.
//
// Solidity: function getEthereumBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) GetEthereumBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "getEthereumBridge")
	return *ret0, err
}

// GetEthereumBridge is a free data retrieval call binding the contract method 0x960bc2ee.
//
// Solidity: function getEthereumBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) GetEthereumBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.GetEthereumBridge(&_BridgeRegistry.CallOpts)
}

// GetEthereumBridge is a free data retrieval call binding the contract method 0x960bc2ee.
//
// Solidity: function getEthereumBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) GetEthereumBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.GetEthereumBridge(&_BridgeRegistry.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "getOperator")
	return *ret0, err
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) GetOperator() (common.Address, error) {
	return _BridgeRegistry.Contract.GetOperator(&_BridgeRegistry.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) GetOperator() (common.Address, error) {
	return _BridgeRegistry.Contract.GetOperator(&_BridgeRegistry.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "getOracle")
	return *ret0, err
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) GetOracle() (common.Address, error) {
	return _BridgeRegistry.Contract.GetOracle(&_BridgeRegistry.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) GetOracle() (common.Address, error) {
	return _BridgeRegistry.Contract.GetOracle(&_BridgeRegistry.CallOpts)
}

// GetValset is a free data retrieval call binding the contract method 0x9fd168ec.
//
// Solidity: function getValset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) GetValset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "getValset")
	return *ret0, err
}

// GetValset is a free data retrieval call binding the contract method 0x9fd168ec.
//
// Solidity: function getValset() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) GetValset() (common.Address, error) {
	return _BridgeRegistry.Contract.GetValset(&_BridgeRegistry.CallOpts)
}

// GetValset is a free data retrieval call binding the contract method 0x9fd168ec.
//
// Solidity: function getValset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) GetValset() (common.Address, error) {
	return _BridgeRegistry.Contract.GetValset(&_BridgeRegistry.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Operator() (common.Address, error) {
	return _BridgeRegistry.Contract.Operator(&_BridgeRegistry.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Operator() (common.Address, error) {
	return _BridgeRegistry.Contract.Operator(&_BridgeRegistry.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Oracle() (common.Address, error) {
	return _BridgeRegistry.Contract.Oracle(&_BridgeRegistry.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Oracle() (common.Address, error) {
	return _BridgeRegistry.Contract.Oracle(&_BridgeRegistry.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Valset() (common.Address, error) {
	return _BridgeRegistry.Contract.Valset(&_BridgeRegistry.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Valset() (common.Address, error) {
	return _BridgeRegistry.Contract.Valset(&_BridgeRegistry.CallOpts)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_BridgeRegistry *BridgeRegistryTransactor) SetBridgeBank(opts *bind.TransactOpts, _bridgeBank common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.contract.Transact(opts, "setBridgeBank", _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_BridgeRegistry *BridgeRegistrySession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetBridgeBank(&_BridgeRegistry.TransactOpts, _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_BridgeRegistry *BridgeRegistryTransactorSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetBridgeBank(&_BridgeRegistry.TransactOpts, _bridgeBank)
}

// SetEthereumBridge is a paid mutator transaction binding the contract method 0x005e09ca.
//
// Solidity: function setEthereumBridge(address _ethereumBridge) returns()
func (_BridgeRegistry *BridgeRegistryTransactor) SetEthereumBridge(opts *bind.TransactOpts, _ethereumBridge common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.contract.Transact(opts, "setEthereumBridge", _ethereumBridge)
}

// SetEthereumBridge is a paid mutator transaction binding the contract method 0x005e09ca.
//
// Solidity: function setEthereumBridge(address _ethereumBridge) returns()
func (_BridgeRegistry *BridgeRegistrySession) SetEthereumBridge(_ethereumBridge common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetEthereumBridge(&_BridgeRegistry.TransactOpts, _ethereumBridge)
}

// SetEthereumBridge is a paid mutator transaction binding the contract method 0x005e09ca.
//
// Solidity: function setEthereumBridge(address _ethereumBridge) returns()
func (_BridgeRegistry *BridgeRegistryTransactorSession) SetEthereumBridge(_ethereumBridge common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetEthereumBridge(&_BridgeRegistry.TransactOpts, _ethereumBridge)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_BridgeRegistry *BridgeRegistryTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_BridgeRegistry *BridgeRegistrySession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetOracle(&_BridgeRegistry.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_BridgeRegistry *BridgeRegistryTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetOracle(&_BridgeRegistry.TransactOpts, _oracle)
}

// SetValset is a paid mutator transaction binding the contract method 0x3d051a9e.
//
// Solidity: function setValset(address _valset) returns()
func (_BridgeRegistry *BridgeRegistryTransactor) SetValset(opts *bind.TransactOpts, _valset common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.contract.Transact(opts, "setValset", _valset)
}

// SetValset is a paid mutator transaction binding the contract method 0x3d051a9e.
//
// Solidity: function setValset(address _valset) returns()
func (_BridgeRegistry *BridgeRegistrySession) SetValset(_valset common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetValset(&_BridgeRegistry.TransactOpts, _valset)
}

// SetValset is a paid mutator transaction binding the contract method 0x3d051a9e.
//
// Solidity: function setValset(address _valset) returns()
func (_BridgeRegistry *BridgeRegistryTransactorSession) SetValset(_valset common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetValset(&_BridgeRegistry.TransactOpts, _valset)
}

// BridgeRegistryHmyLogBridgeBankSetIterator is returned from FilterHmyLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for HmyLogBridgeBankSet events raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogBridgeBankSetIterator struct {
	Event *BridgeRegistryHmyLogBridgeBankSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryHmyLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryHmyLogBridgeBankSet)
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
		it.Event = new(BridgeRegistryHmyLogBridgeBankSet)
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
func (it *BridgeRegistryHmyLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryHmyLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryHmyLogBridgeBankSet represents a HmyLogBridgeBankSet event raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHmyLogBridgeBankSet is a free log retrieval operation binding the contract event 0xa4ce696f3b5c3ecec852cb66da1f78522d9a3b5a98f2e3428eeca4ce0931df86.
//
// Solidity: event HmyLogBridgeBankSet(address _bridgeBank)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterHmyLogBridgeBankSet(opts *bind.FilterOpts) (*BridgeRegistryHmyLogBridgeBankSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "HmyLogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryHmyLogBridgeBankSetIterator{contract: _BridgeRegistry.contract, event: "HmyLogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchHmyLogBridgeBankSet is a free log subscription operation binding the contract event 0xa4ce696f3b5c3ecec852cb66da1f78522d9a3b5a98f2e3428eeca4ce0931df86.
//
// Solidity: event HmyLogBridgeBankSet(address _bridgeBank)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchHmyLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryHmyLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "HmyLogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryHmyLogBridgeBankSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogBridgeBankSet", log); err != nil {
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

// ParseHmyLogBridgeBankSet is a log parse operation binding the contract event 0xa4ce696f3b5c3ecec852cb66da1f78522d9a3b5a98f2e3428eeca4ce0931df86.
//
// Solidity: event HmyLogBridgeBankSet(address _bridgeBank)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseHmyLogBridgeBankSet(log types.Log) (*BridgeRegistryHmyLogBridgeBankSet, error) {
	event := new(BridgeRegistryHmyLogBridgeBankSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogBridgeBankSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRegistryHmyLogHarmonyBridgeSetIterator is returned from FilterHmyLogHarmonyBridgeSet and is used to iterate over the raw logs and unpacked data for HmyLogHarmonyBridgeSet events raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogHarmonyBridgeSetIterator struct {
	Event *BridgeRegistryHmyLogHarmonyBridgeSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryHmyLogHarmonyBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryHmyLogHarmonyBridgeSet)
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
		it.Event = new(BridgeRegistryHmyLogHarmonyBridgeSet)
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
func (it *BridgeRegistryHmyLogHarmonyBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryHmyLogHarmonyBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryHmyLogHarmonyBridgeSet represents a HmyLogHarmonyBridgeSet event raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogHarmonyBridgeSet struct {
	EthereumBridge common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterHmyLogHarmonyBridgeSet is a free log retrieval operation binding the contract event 0xd3869bed92142314dc7f6cd8e15f19f72f14034c42e74a8b170d39b42fa1b4ac.
//
// Solidity: event HmyLogHarmonyBridgeSet(address _ethereumBridge)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterHmyLogHarmonyBridgeSet(opts *bind.FilterOpts) (*BridgeRegistryHmyLogHarmonyBridgeSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "HmyLogHarmonyBridgeSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryHmyLogHarmonyBridgeSetIterator{contract: _BridgeRegistry.contract, event: "HmyLogHarmonyBridgeSet", logs: logs, sub: sub}, nil
}

// WatchHmyLogHarmonyBridgeSet is a free log subscription operation binding the contract event 0xd3869bed92142314dc7f6cd8e15f19f72f14034c42e74a8b170d39b42fa1b4ac.
//
// Solidity: event HmyLogHarmonyBridgeSet(address _ethereumBridge)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchHmyLogHarmonyBridgeSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryHmyLogHarmonyBridgeSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "HmyLogHarmonyBridgeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryHmyLogHarmonyBridgeSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogHarmonyBridgeSet", log); err != nil {
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

// ParseHmyLogHarmonyBridgeSet is a log parse operation binding the contract event 0xd3869bed92142314dc7f6cd8e15f19f72f14034c42e74a8b170d39b42fa1b4ac.
//
// Solidity: event HmyLogHarmonyBridgeSet(address _ethereumBridge)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseHmyLogHarmonyBridgeSet(log types.Log) (*BridgeRegistryHmyLogHarmonyBridgeSet, error) {
	event := new(BridgeRegistryHmyLogHarmonyBridgeSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogHarmonyBridgeSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRegistryHmyLogOracleSetIterator is returned from FilterHmyLogOracleSet and is used to iterate over the raw logs and unpacked data for HmyLogOracleSet events raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogOracleSetIterator struct {
	Event *BridgeRegistryHmyLogOracleSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryHmyLogOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryHmyLogOracleSet)
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
		it.Event = new(BridgeRegistryHmyLogOracleSet)
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
func (it *BridgeRegistryHmyLogOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryHmyLogOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryHmyLogOracleSet represents a HmyLogOracleSet event raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHmyLogOracleSet is a free log retrieval operation binding the contract event 0xad9ae01008c260a11bf3fa6e389527b6a3a3fb1ed015c1937a9087d7c6dbc18d.
//
// Solidity: event HmyLogOracleSet(address _oracle)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterHmyLogOracleSet(opts *bind.FilterOpts) (*BridgeRegistryHmyLogOracleSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "HmyLogOracleSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryHmyLogOracleSetIterator{contract: _BridgeRegistry.contract, event: "HmyLogOracleSet", logs: logs, sub: sub}, nil
}

// WatchHmyLogOracleSet is a free log subscription operation binding the contract event 0xad9ae01008c260a11bf3fa6e389527b6a3a3fb1ed015c1937a9087d7c6dbc18d.
//
// Solidity: event HmyLogOracleSet(address _oracle)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchHmyLogOracleSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryHmyLogOracleSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "HmyLogOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryHmyLogOracleSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogOracleSet", log); err != nil {
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

// ParseHmyLogOracleSet is a log parse operation binding the contract event 0xad9ae01008c260a11bf3fa6e389527b6a3a3fb1ed015c1937a9087d7c6dbc18d.
//
// Solidity: event HmyLogOracleSet(address _oracle)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseHmyLogOracleSet(log types.Log) (*BridgeRegistryHmyLogOracleSet, error) {
	event := new(BridgeRegistryHmyLogOracleSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogOracleSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRegistryHmyLogValsetSetIterator is returned from FilterHmyLogValsetSet and is used to iterate over the raw logs and unpacked data for HmyLogValsetSet events raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogValsetSetIterator struct {
	Event *BridgeRegistryHmyLogValsetSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryHmyLogValsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryHmyLogValsetSet)
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
		it.Event = new(BridgeRegistryHmyLogValsetSet)
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
func (it *BridgeRegistryHmyLogValsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryHmyLogValsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryHmyLogValsetSet represents a HmyLogValsetSet event raised by the BridgeRegistry contract.
type BridgeRegistryHmyLogValsetSet struct {
	Valset common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHmyLogValsetSet is a free log retrieval operation binding the contract event 0x31a3f52a07e4eaa601f0c839a0b72df88e5a49ae52a6aefcf04eeac96c72f9aa.
//
// Solidity: event HmyLogValsetSet(address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterHmyLogValsetSet(opts *bind.FilterOpts) (*BridgeRegistryHmyLogValsetSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "HmyLogValsetSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryHmyLogValsetSetIterator{contract: _BridgeRegistry.contract, event: "HmyLogValsetSet", logs: logs, sub: sub}, nil
}

// WatchHmyLogValsetSet is a free log subscription operation binding the contract event 0x31a3f52a07e4eaa601f0c839a0b72df88e5a49ae52a6aefcf04eeac96c72f9aa.
//
// Solidity: event HmyLogValsetSet(address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchHmyLogValsetSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryHmyLogValsetSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "HmyLogValsetSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryHmyLogValsetSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogValsetSet", log); err != nil {
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

// ParseHmyLogValsetSet is a log parse operation binding the contract event 0x31a3f52a07e4eaa601f0c839a0b72df88e5a49ae52a6aefcf04eeac96c72f9aa.
//
// Solidity: event HmyLogValsetSet(address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseHmyLogValsetSet(log types.Log) (*BridgeRegistryHmyLogValsetSet, error) {
	event := new(BridgeRegistryHmyLogValsetSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "HmyLogValsetSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
