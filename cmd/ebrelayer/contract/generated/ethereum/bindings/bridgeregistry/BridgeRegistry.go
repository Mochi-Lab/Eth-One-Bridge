// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeRegistry

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

// BridgeRegistryABI is the input ABI used to generate the binding from.
const BridgeRegistryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"EthLogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonyBridge\",\"type\":\"address\"}],\"name\":\"EthLogHarmonyBridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"EthLogOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"}],\"name\":\"EthLogValsetSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBridgeBank\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHarmonyBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"harmonyBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyBridge\",\"type\":\"address\"}],\"name\":\"setHarmonyBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"}],\"name\":\"setValset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeRegistryBin is the compiled bytecode used for deploying new contracts.
var BridgeRegistryBin = "0x608060405234801561001057600080fd5b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c91806100616000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063814c92c31161008c5780639fd168ec116100665780639fd168ec146103c1578063d33f397e1461040b578063e7f43c6814610455578063f586612d1461049f576100ea565b8063814c92c3146102e9578063833b1fce1461032d578063972ded0914610377576100ea565b8063570ca735116100c8578063570ca735146101c75780637adbf973146102115780637dc0d1d0146102555780637f54af0c1461029f576100ea565b80630e41f373146100ef5780633d051a9e14610139578063476dc25c1461017d575b600080fd5b6100f76104e3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61017b6004803603602081101561014f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610509565b005b610185610673565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101cf61069c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102536004803603602081101561022757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106c2565b005b61025d61084e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102a7610874565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61032b600480360360208110156102ff57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061089a565b005b610335610a26565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61037f610a50565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103c9610a75565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610413610a9f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61045d610ac9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104e1600480360360208110156104b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610af3565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f87440d8040bc9fb5af27bfa874ca43f4b96b8b524659be11298a090d4477622681604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610785576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f9ffd222950e4db02bc92f4d13e818aa185bdaa667d849024c5ba4d96e2db9677600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe3cac451a52739458bf231fa2a2dacca7e1bc0567481d729b0c962acafb66faf600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bb6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f75349d6f5fa7b7f9ef4e394c23f391adc18ccdea3e1ae9cf51d538f4bf23759981604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15056fea265627a7a72315820a8384d9ad111ec3830d6b167ea433ecc8a6d50ad2291ba65f8b11a16d8f265e864736f6c63430005110032"

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

// GetHarmonyBridge is a free data retrieval call binding the contract method 0x476dc25c.
//
// Solidity: function getHarmonyBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) GetHarmonyBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "getHarmonyBridge")
	return *ret0, err
}

// GetHarmonyBridge is a free data retrieval call binding the contract method 0x476dc25c.
//
// Solidity: function getHarmonyBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) GetHarmonyBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.GetHarmonyBridge(&_BridgeRegistry.CallOpts)
}

// GetHarmonyBridge is a free data retrieval call binding the contract method 0x476dc25c.
//
// Solidity: function getHarmonyBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) GetHarmonyBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.GetHarmonyBridge(&_BridgeRegistry.CallOpts)
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

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) HarmonyBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "harmonyBridge")
	return *ret0, err
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) HarmonyBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.HarmonyBridge(&_BridgeRegistry.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) HarmonyBridge() (common.Address, error) {
	return _BridgeRegistry.Contract.HarmonyBridge(&_BridgeRegistry.CallOpts)
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

// SetHarmonyBridge is a paid mutator transaction binding the contract method 0xf586612d.
//
// Solidity: function setHarmonyBridge(address _harmonyBridge) returns()
func (_BridgeRegistry *BridgeRegistryTransactor) SetHarmonyBridge(opts *bind.TransactOpts, _harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.contract.Transact(opts, "setHarmonyBridge", _harmonyBridge)
}

// SetHarmonyBridge is a paid mutator transaction binding the contract method 0xf586612d.
//
// Solidity: function setHarmonyBridge(address _harmonyBridge) returns()
func (_BridgeRegistry *BridgeRegistrySession) SetHarmonyBridge(_harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetHarmonyBridge(&_BridgeRegistry.TransactOpts, _harmonyBridge)
}

// SetHarmonyBridge is a paid mutator transaction binding the contract method 0xf586612d.
//
// Solidity: function setHarmonyBridge(address _harmonyBridge) returns()
func (_BridgeRegistry *BridgeRegistryTransactorSession) SetHarmonyBridge(_harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.SetHarmonyBridge(&_BridgeRegistry.TransactOpts, _harmonyBridge)
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

// BridgeRegistryEthLogBridgeBankSetIterator is returned from FilterEthLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for EthLogBridgeBankSet events raised by the BridgeRegistry contract.
type BridgeRegistryEthLogBridgeBankSetIterator struct {
	Event *BridgeRegistryEthLogBridgeBankSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryEthLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryEthLogBridgeBankSet)
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
		it.Event = new(BridgeRegistryEthLogBridgeBankSet)
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
func (it *BridgeRegistryEthLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryEthLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryEthLogBridgeBankSet represents a EthLogBridgeBankSet event raised by the BridgeRegistry contract.
type BridgeRegistryEthLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthLogBridgeBankSet is a free log retrieval operation binding the contract event 0xe3cac451a52739458bf231fa2a2dacca7e1bc0567481d729b0c962acafb66faf.
//
// Solidity: event EthLogBridgeBankSet(address _bridgeBank)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterEthLogBridgeBankSet(opts *bind.FilterOpts) (*BridgeRegistryEthLogBridgeBankSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "EthLogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryEthLogBridgeBankSetIterator{contract: _BridgeRegistry.contract, event: "EthLogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchEthLogBridgeBankSet is a free log subscription operation binding the contract event 0xe3cac451a52739458bf231fa2a2dacca7e1bc0567481d729b0c962acafb66faf.
//
// Solidity: event EthLogBridgeBankSet(address _bridgeBank)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchEthLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryEthLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "EthLogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryEthLogBridgeBankSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogBridgeBankSet", log); err != nil {
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

// ParseEthLogBridgeBankSet is a log parse operation binding the contract event 0xe3cac451a52739458bf231fa2a2dacca7e1bc0567481d729b0c962acafb66faf.
//
// Solidity: event EthLogBridgeBankSet(address _bridgeBank)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseEthLogBridgeBankSet(log types.Log) (*BridgeRegistryEthLogBridgeBankSet, error) {
	event := new(BridgeRegistryEthLogBridgeBankSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogBridgeBankSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRegistryEthLogHarmonyBridgeSetIterator is returned from FilterEthLogHarmonyBridgeSet and is used to iterate over the raw logs and unpacked data for EthLogHarmonyBridgeSet events raised by the BridgeRegistry contract.
type BridgeRegistryEthLogHarmonyBridgeSetIterator struct {
	Event *BridgeRegistryEthLogHarmonyBridgeSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryEthLogHarmonyBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryEthLogHarmonyBridgeSet)
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
		it.Event = new(BridgeRegistryEthLogHarmonyBridgeSet)
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
func (it *BridgeRegistryEthLogHarmonyBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryEthLogHarmonyBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryEthLogHarmonyBridgeSet represents a EthLogHarmonyBridgeSet event raised by the BridgeRegistry contract.
type BridgeRegistryEthLogHarmonyBridgeSet struct {
	HarmonyBridge common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthLogHarmonyBridgeSet is a free log retrieval operation binding the contract event 0x75349d6f5fa7b7f9ef4e394c23f391adc18ccdea3e1ae9cf51d538f4bf237599.
//
// Solidity: event EthLogHarmonyBridgeSet(address _harmonyBridge)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterEthLogHarmonyBridgeSet(opts *bind.FilterOpts) (*BridgeRegistryEthLogHarmonyBridgeSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "EthLogHarmonyBridgeSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryEthLogHarmonyBridgeSetIterator{contract: _BridgeRegistry.contract, event: "EthLogHarmonyBridgeSet", logs: logs, sub: sub}, nil
}

// WatchEthLogHarmonyBridgeSet is a free log subscription operation binding the contract event 0x75349d6f5fa7b7f9ef4e394c23f391adc18ccdea3e1ae9cf51d538f4bf237599.
//
// Solidity: event EthLogHarmonyBridgeSet(address _harmonyBridge)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchEthLogHarmonyBridgeSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryEthLogHarmonyBridgeSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "EthLogHarmonyBridgeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryEthLogHarmonyBridgeSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogHarmonyBridgeSet", log); err != nil {
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

// ParseEthLogHarmonyBridgeSet is a log parse operation binding the contract event 0x75349d6f5fa7b7f9ef4e394c23f391adc18ccdea3e1ae9cf51d538f4bf237599.
//
// Solidity: event EthLogHarmonyBridgeSet(address _harmonyBridge)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseEthLogHarmonyBridgeSet(log types.Log) (*BridgeRegistryEthLogHarmonyBridgeSet, error) {
	event := new(BridgeRegistryEthLogHarmonyBridgeSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogHarmonyBridgeSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRegistryEthLogOracleSetIterator is returned from FilterEthLogOracleSet and is used to iterate over the raw logs and unpacked data for EthLogOracleSet events raised by the BridgeRegistry contract.
type BridgeRegistryEthLogOracleSetIterator struct {
	Event *BridgeRegistryEthLogOracleSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryEthLogOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryEthLogOracleSet)
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
		it.Event = new(BridgeRegistryEthLogOracleSet)
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
func (it *BridgeRegistryEthLogOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryEthLogOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryEthLogOracleSet represents a EthLogOracleSet event raised by the BridgeRegistry contract.
type BridgeRegistryEthLogOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthLogOracleSet is a free log retrieval operation binding the contract event 0x9ffd222950e4db02bc92f4d13e818aa185bdaa667d849024c5ba4d96e2db9677.
//
// Solidity: event EthLogOracleSet(address _oracle)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterEthLogOracleSet(opts *bind.FilterOpts) (*BridgeRegistryEthLogOracleSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "EthLogOracleSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryEthLogOracleSetIterator{contract: _BridgeRegistry.contract, event: "EthLogOracleSet", logs: logs, sub: sub}, nil
}

// WatchEthLogOracleSet is a free log subscription operation binding the contract event 0x9ffd222950e4db02bc92f4d13e818aa185bdaa667d849024c5ba4d96e2db9677.
//
// Solidity: event EthLogOracleSet(address _oracle)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchEthLogOracleSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryEthLogOracleSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "EthLogOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryEthLogOracleSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogOracleSet", log); err != nil {
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

// ParseEthLogOracleSet is a log parse operation binding the contract event 0x9ffd222950e4db02bc92f4d13e818aa185bdaa667d849024c5ba4d96e2db9677.
//
// Solidity: event EthLogOracleSet(address _oracle)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseEthLogOracleSet(log types.Log) (*BridgeRegistryEthLogOracleSet, error) {
	event := new(BridgeRegistryEthLogOracleSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogOracleSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeRegistryEthLogValsetSetIterator is returned from FilterEthLogValsetSet and is used to iterate over the raw logs and unpacked data for EthLogValsetSet events raised by the BridgeRegistry contract.
type BridgeRegistryEthLogValsetSetIterator struct {
	Event *BridgeRegistryEthLogValsetSet // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryEthLogValsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryEthLogValsetSet)
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
		it.Event = new(BridgeRegistryEthLogValsetSet)
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
func (it *BridgeRegistryEthLogValsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryEthLogValsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryEthLogValsetSet represents a EthLogValsetSet event raised by the BridgeRegistry contract.
type BridgeRegistryEthLogValsetSet struct {
	Valset common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthLogValsetSet is a free log retrieval operation binding the contract event 0x87440d8040bc9fb5af27bfa874ca43f4b96b8b524659be11298a090d44776226.
//
// Solidity: event EthLogValsetSet(address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterEthLogValsetSet(opts *bind.FilterOpts) (*BridgeRegistryEthLogValsetSetIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "EthLogValsetSet")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryEthLogValsetSetIterator{contract: _BridgeRegistry.contract, event: "EthLogValsetSet", logs: logs, sub: sub}, nil
}

// WatchEthLogValsetSet is a free log subscription operation binding the contract event 0x87440d8040bc9fb5af27bfa874ca43f4b96b8b524659be11298a090d44776226.
//
// Solidity: event EthLogValsetSet(address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchEthLogValsetSet(opts *bind.WatchOpts, sink chan<- *BridgeRegistryEthLogValsetSet) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "EthLogValsetSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryEthLogValsetSet)
				if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogValsetSet", log); err != nil {
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

// ParseEthLogValsetSet is a log parse operation binding the contract event 0x87440d8040bc9fb5af27bfa874ca43f4b96b8b524659be11298a090d44776226.
//
// Solidity: event EthLogValsetSet(address _valset)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseEthLogValsetSet(log types.Log) (*BridgeRegistryEthLogValsetSet, error) {
	event := new(BridgeRegistryEthLogValsetSet)
	if err := _BridgeRegistry.contract.UnpackLog(event, "EthLogValsetSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
