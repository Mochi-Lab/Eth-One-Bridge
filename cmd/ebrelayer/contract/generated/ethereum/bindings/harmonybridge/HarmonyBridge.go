// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HarmonyBridge

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

// HarmonyBridgeABI is the input ABI used to generate the binding from.
const HarmonyBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonySender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthLogNewUnlockClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"EthLogUnlockCompleted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"HARMONYBRIDGE_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"contractBridgeBank\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeRegistry\",\"outputs\":[{\"internalType\":\"contractBridgeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"completeUnlockClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"isUnlockClaimActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"isUnlockClaimValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonySender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"newUnlockClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockClaimCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unlockClaims\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"harmonySender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originalValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumHarmonyBridge.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"contractValset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HarmonyBridgeBin is the compiled bytecode used for deploying new contracts.
var HarmonyBridgeBin = "0x60806040526000805534801561001457600080fd5b50611cec806100246000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80637dc0d1d01161008c578063b1493ce011610066578063b1493ce01461046b578063b1c51a8014610489578063c4d66de8146104a7578063cf93b56c146104eb576100cf565b80637dc0d1d0146103a95780637f54af0c146103f35780639e67206d1461043d576100cf565b80630e41f373146100d45780631a86e3771461011e578063316be17114610164578063570ca735146101ae57806360cf6648146101f85780637c5d759014610286575b600080fd5b6100dc610531565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61014a6004803603602081101561013457600080fd5b8101908080359060200190929190505050610557565b604051808215151515815260200191505060405180910390f35b61016c61059d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101b66105c3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102846004803603608081101561020e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105e9565b005b6102b26004803603602081101561029c57600080fd5b8101908080359060200190929190505050610da6565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182600381111561039057fe5b60ff168152602001965050505050505060405180910390f35b6103b1610e6f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103fb610e95565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104696004803603602081101561045357600080fd5b8101908080359060200190929190505050610ebb565b005b610473611059565b6040518082815260200191505060405180910390f35b61049161105f565b6040518082815260200191505060405180910390f35b6104e9600480360360208110156104bd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611064565b005b6105176004803603602081101561050157600080fd5b8101908080359060200190929190505050611529565b604051808215151515815260200191505060405180910390f35b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001600381111561056657fe5b603a600084815260200190815260200160002060050160009054906101000a900460ff16600381111561059557fe5b149050919050565b603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663833b1fce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561066957600080fd5b505afa15801561067d573d6000803e3d6000fd5b505050506040513d602081101561069357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141580156107975750600073ffffffffffffffffffffffffffffffffffffffff16603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d33f397e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561074357600080fd5b505afa158015610757573d6000803e3d6000fd5b505050506040513d602081101561076d57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614155b6107ec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526046815260200180611be36046913960600191505060405180910390fd5b603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561088b57600080fd5b505afa15801561089f573d6000803e3d6000fd5b505050506040513d60208110156108b557600080fd5b8101908080519060200190929190505050610938576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b4bfd9a783836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060206040518083038186803b1580156109df57600080fd5b505afa1580156109f3573d6000803e3d6000fd5b505050506040513d6020811015610a0957600080fd5b8101908080519060200190929190505050610a6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180611c57603a913960400191505060405180910390fd5b610a77611b49565b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160016003811115610b0557fe5b8152509050610b20600160345461164290919063ffffffff16565b60348190555080603a6000603454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050160006101000a81548160ff02191690836003811115610c8657fe5b02179055509050507f4389bb697e92204e405d58c4811114444ca51e675f1258844e426ce971bf4c566034548686338787604051808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405180910390a15050505050565b603a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff16905086565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80610ec581610557565b610f37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f556e6c6f636b20636c61696d206973206e6f742061637469766500000000000081525060200191505060405180910390fd5b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fdd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180611c916027913960400191505060405180910390fd5b6002603a600084815260200190815260200160002060050160006101000a81548160ff0219169083600381111561101057fe5b021790555061101e826116ca565b7f56bad42424f5c90e6bfd5e97f55f32bec9a642653ef665d2febfaa6b69e765dc826040518082815260200191505060405180910390a15050565b60345481565b600181565b600061106e611b2f565b9050600160009054906101000a900460ff168061108f575061108e611b38565b5b8061109b575060005481115b6110f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611c29602e913960400191505060405180910390fd5b6000600160009054906101000a900460ff16159050801561112d5760018060006101000a81548160ff021916908315150217905550816000819055505b82603860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b1580156111d657600080fd5b505afa1580156111ea573d6000803e3d6000fd5b505050506040513d602081101561120057600080fd5b8101908080519060200190929190505050603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fd168ec6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112b957600080fd5b505afa1580156112cd573d6000803e3d6000fd5b505050506040513d60208110156112e357600080fd5b8101908080519060200190929190505050603660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d33f397e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561139c57600080fd5b505afa1580156113b0573d6000803e3d6000fd5b505050506040513d60208110156113c657600080fd5b8101908080519060200190929190505050603760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663833b1fce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561147f57600080fd5b505afa158015611493573d6000803e3d6000fd5b505050506040513d60208110156114a957600080fd5b8101908080519060200190929190505050603960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006034819055508015611524576000600160006101000a81548160ff0219169083151502179055505b505050565b6000603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c603a600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561160057600080fd5b505afa158015611614573d6000803e3d6000fd5b505050506040513d602081101561162a57600080fd5b81019080805190602001909291905050509050919050565b6000808284019050838110156116c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6116d2611b49565b603a60008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820160009054906101000a900460ff16600381111561187057fe5b600381111561187b57fe5b815250509050603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da228a9b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156118e957600080fd5b505afa1580156118fd573d6000803e3d6000fd5b505050506040513d602081101561191357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16816060015173ffffffffffffffffffffffffffffffffffffffff161415611a2957603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed581fd3826020015183608001516040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611a0c57600080fd5b505af1158015611a20573d6000803e3d6000fd5b50505050611b2b565b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633c298e788260200151836060015184608001516040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611b1257600080fd5b505af1158015611b26573d6000803e3d6000fd5b505050505b5050565b60006001905090565b600080303b90506000811491505090565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160006003811115611bdc57fe5b8152509056fe546865204f70657261746f72206d7573742073657420746865206f7261636c6520616e64206272696467652062616e6b20666f72206272696467652061637469766174696f6e436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65644e6f7420656e6f756768206c6f636b65642061737365747320746f20636f6d706c657465207468652070726f706f7365642070726f70686563794f6e6c7920746865204f7261636c65206d617920636f6d706c6574652070726f70686563696573a265627a7a72315820c59ee2191154bdea31d17994f93b58aa0715f82e8181d9efd946d64e57cf32f364736f6c63430005110032"

// DeployHarmonyBridge deploys a new Ethereum contract, binding an instance of HarmonyBridge to it.
func DeployHarmonyBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HarmonyBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(HarmonyBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HarmonyBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HarmonyBridge{HarmonyBridgeCaller: HarmonyBridgeCaller{contract: contract}, HarmonyBridgeTransactor: HarmonyBridgeTransactor{contract: contract}, HarmonyBridgeFilterer: HarmonyBridgeFilterer{contract: contract}}, nil
}

// HarmonyBridge is an auto generated Go binding around an Ethereum contract.
type HarmonyBridge struct {
	HarmonyBridgeCaller     // Read-only binding to the contract
	HarmonyBridgeTransactor // Write-only binding to the contract
	HarmonyBridgeFilterer   // Log filterer for contract events
}

// HarmonyBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HarmonyBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarmonyBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HarmonyBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarmonyBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HarmonyBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HarmonyBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HarmonyBridgeSession struct {
	Contract     *HarmonyBridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HarmonyBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HarmonyBridgeCallerSession struct {
	Contract *HarmonyBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HarmonyBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HarmonyBridgeTransactorSession struct {
	Contract     *HarmonyBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HarmonyBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HarmonyBridgeRaw struct {
	Contract *HarmonyBridge // Generic contract binding to access the raw methods on
}

// HarmonyBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HarmonyBridgeCallerRaw struct {
	Contract *HarmonyBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// HarmonyBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HarmonyBridgeTransactorRaw struct {
	Contract *HarmonyBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHarmonyBridge creates a new instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridge(address common.Address, backend bind.ContractBackend) (*HarmonyBridge, error) {
	contract, err := bindHarmonyBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridge{HarmonyBridgeCaller: HarmonyBridgeCaller{contract: contract}, HarmonyBridgeTransactor: HarmonyBridgeTransactor{contract: contract}, HarmonyBridgeFilterer: HarmonyBridgeFilterer{contract: contract}}, nil
}

// NewHarmonyBridgeCaller creates a new read-only instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridgeCaller(address common.Address, caller bind.ContractCaller) (*HarmonyBridgeCaller, error) {
	contract, err := bindHarmonyBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeCaller{contract: contract}, nil
}

// NewHarmonyBridgeTransactor creates a new write-only instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*HarmonyBridgeTransactor, error) {
	contract, err := bindHarmonyBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeTransactor{contract: contract}, nil
}

// NewHarmonyBridgeFilterer creates a new log filterer instance of HarmonyBridge, bound to a specific deployed contract.
func NewHarmonyBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*HarmonyBridgeFilterer, error) {
	contract, err := bindHarmonyBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeFilterer{contract: contract}, nil
}

// bindHarmonyBridge binds a generic wrapper to an already deployed contract.
func bindHarmonyBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HarmonyBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarmonyBridge *HarmonyBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HarmonyBridge.Contract.HarmonyBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarmonyBridge *HarmonyBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.HarmonyBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarmonyBridge *HarmonyBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.HarmonyBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HarmonyBridge *HarmonyBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HarmonyBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HarmonyBridge *HarmonyBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HarmonyBridge *HarmonyBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.contract.Transact(opts, method, params...)
}

// HARMONYBRIDGEREVISION is a free data retrieval call binding the contract method 0xb1c51a80.
//
// Solidity: function HARMONYBRIDGE_REVISION() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeCaller) HARMONYBRIDGEREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "HARMONYBRIDGE_REVISION")
	return *ret0, err
}

// HARMONYBRIDGEREVISION is a free data retrieval call binding the contract method 0xb1c51a80.
//
// Solidity: function HARMONYBRIDGE_REVISION() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeSession) HARMONYBRIDGEREVISION() (*big.Int, error) {
	return _HarmonyBridge.Contract.HARMONYBRIDGEREVISION(&_HarmonyBridge.CallOpts)
}

// HARMONYBRIDGEREVISION is a free data retrieval call binding the contract method 0xb1c51a80.
//
// Solidity: function HARMONYBRIDGE_REVISION() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeCallerSession) HARMONYBRIDGEREVISION() (*big.Int, error) {
	return _HarmonyBridge.Contract.HARMONYBRIDGEREVISION(&_HarmonyBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) BridgeBank() (common.Address, error) {
	return _HarmonyBridge.Contract.BridgeBank(&_HarmonyBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) BridgeBank() (common.Address, error) {
	return _HarmonyBridge.Contract.BridgeBank(&_HarmonyBridge.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) BridgeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "bridgeRegistry")
	return *ret0, err
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) BridgeRegistry() (common.Address, error) {
	return _HarmonyBridge.Contract.BridgeRegistry(&_HarmonyBridge.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) BridgeRegistry() (common.Address, error) {
	return _HarmonyBridge.Contract.BridgeRegistry(&_HarmonyBridge.CallOpts)
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCaller) IsUnlockClaimActive(opts *bind.CallOpts, _unlockID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "isUnlockClaimActive", _unlockID)
	return *ret0, err
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeSession) IsUnlockClaimActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCallerSession) IsUnlockClaimActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCaller) IsUnlockClaimValidatorActive(opts *bind.CallOpts, _unlockID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "isUnlockClaimValidatorActive", _unlockID)
	return *ret0, err
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeSession) IsUnlockClaimValidatorActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimValidatorActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_HarmonyBridge *HarmonyBridgeCallerSession) IsUnlockClaimValidatorActive(_unlockID *big.Int) (bool, error) {
	return _HarmonyBridge.Contract.IsUnlockClaimValidatorActive(&_HarmonyBridge.CallOpts, _unlockID)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) Operator() (common.Address, error) {
	return _HarmonyBridge.Contract.Operator(&_HarmonyBridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) Operator() (common.Address, error) {
	return _HarmonyBridge.Contract.Operator(&_HarmonyBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) Oracle() (common.Address, error) {
	return _HarmonyBridge.Contract.Oracle(&_HarmonyBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) Oracle() (common.Address, error) {
	return _HarmonyBridge.Contract.Oracle(&_HarmonyBridge.CallOpts)
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeCaller) UnlockClaimCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "unlockClaimCount")
	return *ret0, err
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeSession) UnlockClaimCount() (*big.Int, error) {
	return _HarmonyBridge.Contract.UnlockClaimCount(&_HarmonyBridge.CallOpts)
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_HarmonyBridge *HarmonyBridgeCallerSession) UnlockClaimCount() (*big.Int, error) {
	return _HarmonyBridge.Contract.UnlockClaimCount(&_HarmonyBridge.CallOpts)
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address harmonySender, address ethereumReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_HarmonyBridge *HarmonyBridgeCaller) UnlockClaims(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HarmonySender     common.Address
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	ret := new(struct {
		HarmonySender     common.Address
		EthereumReceiver  common.Address
		OriginalValidator common.Address
		Token             common.Address
		Amount            *big.Int
		Status            uint8
	})
	out := ret
	err := _HarmonyBridge.contract.Call(opts, out, "unlockClaims", arg0)
	return *ret, err
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address harmonySender, address ethereumReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_HarmonyBridge *HarmonyBridgeSession) UnlockClaims(arg0 *big.Int) (struct {
	HarmonySender     common.Address
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	return _HarmonyBridge.Contract.UnlockClaims(&_HarmonyBridge.CallOpts, arg0)
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address harmonySender, address ethereumReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_HarmonyBridge *HarmonyBridgeCallerSession) UnlockClaims(arg0 *big.Int) (struct {
	HarmonySender     common.Address
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	return _HarmonyBridge.Contract.UnlockClaims(&_HarmonyBridge.CallOpts, arg0)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HarmonyBridge.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_HarmonyBridge *HarmonyBridgeSession) Valset() (common.Address, error) {
	return _HarmonyBridge.Contract.Valset(&_HarmonyBridge.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_HarmonyBridge *HarmonyBridgeCallerSession) Valset() (common.Address, error) {
	return _HarmonyBridge.Contract.Valset(&_HarmonyBridge.CallOpts)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) CompleteUnlockClaim(opts *bind.TransactOpts, _unlockID *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "completeUnlockClaim", _unlockID)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_HarmonyBridge *HarmonyBridgeSession) CompleteUnlockClaim(_unlockID *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.CompleteUnlockClaim(&_HarmonyBridge.TransactOpts, _unlockID)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) CompleteUnlockClaim(_unlockID *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.CompleteUnlockClaim(&_HarmonyBridge.TransactOpts, _unlockID)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeRegistry) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) Initialize(opts *bind.TransactOpts, _bridgeRegistry common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "initialize", _bridgeRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeRegistry) returns()
func (_HarmonyBridge *HarmonyBridgeSession) Initialize(_bridgeRegistry common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.Initialize(&_HarmonyBridge.TransactOpts, _bridgeRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeRegistry) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) Initialize(_bridgeRegistry common.Address) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.Initialize(&_HarmonyBridge.TransactOpts, _bridgeRegistry)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _harmonySender, address _ethereumReceiver, address _token, uint256 _amount) returns()
func (_HarmonyBridge *HarmonyBridgeTransactor) NewUnlockClaim(opts *bind.TransactOpts, _harmonySender common.Address, _ethereumReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.contract.Transact(opts, "newUnlockClaim", _harmonySender, _ethereumReceiver, _token, _amount)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _harmonySender, address _ethereumReceiver, address _token, uint256 _amount) returns()
func (_HarmonyBridge *HarmonyBridgeSession) NewUnlockClaim(_harmonySender common.Address, _ethereumReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.NewUnlockClaim(&_HarmonyBridge.TransactOpts, _harmonySender, _ethereumReceiver, _token, _amount)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _harmonySender, address _ethereumReceiver, address _token, uint256 _amount) returns()
func (_HarmonyBridge *HarmonyBridgeTransactorSession) NewUnlockClaim(_harmonySender common.Address, _ethereumReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HarmonyBridge.Contract.NewUnlockClaim(&_HarmonyBridge.TransactOpts, _harmonySender, _ethereumReceiver, _token, _amount)
}

// HarmonyBridgeEthLogNewUnlockClaimIterator is returned from FilterEthLogNewUnlockClaim and is used to iterate over the raw logs and unpacked data for EthLogNewUnlockClaim events raised by the HarmonyBridge contract.
type HarmonyBridgeEthLogNewUnlockClaimIterator struct {
	Event *HarmonyBridgeEthLogNewUnlockClaim // Event containing the contract specifics and raw log

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
func (it *HarmonyBridgeEthLogNewUnlockClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarmonyBridgeEthLogNewUnlockClaim)
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
		it.Event = new(HarmonyBridgeEthLogNewUnlockClaim)
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
func (it *HarmonyBridgeEthLogNewUnlockClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarmonyBridgeEthLogNewUnlockClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarmonyBridgeEthLogNewUnlockClaim represents a EthLogNewUnlockClaim event raised by the HarmonyBridge contract.
type HarmonyBridgeEthLogNewUnlockClaim struct {
	UnlockID         *big.Int
	HarmonySender    common.Address
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEthLogNewUnlockClaim is a free log retrieval operation binding the contract event 0x4389bb697e92204e405d58c4811114444ca51e675f1258844e426ce971bf4c56.
//
// Solidity: event EthLogNewUnlockClaim(uint256 _unlockID, address _harmonySender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_HarmonyBridge *HarmonyBridgeFilterer) FilterEthLogNewUnlockClaim(opts *bind.FilterOpts) (*HarmonyBridgeEthLogNewUnlockClaimIterator, error) {

	logs, sub, err := _HarmonyBridge.contract.FilterLogs(opts, "EthLogNewUnlockClaim")
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeEthLogNewUnlockClaimIterator{contract: _HarmonyBridge.contract, event: "EthLogNewUnlockClaim", logs: logs, sub: sub}, nil
}

// WatchEthLogNewUnlockClaim is a free log subscription operation binding the contract event 0x4389bb697e92204e405d58c4811114444ca51e675f1258844e426ce971bf4c56.
//
// Solidity: event EthLogNewUnlockClaim(uint256 _unlockID, address _harmonySender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_HarmonyBridge *HarmonyBridgeFilterer) WatchEthLogNewUnlockClaim(opts *bind.WatchOpts, sink chan<- *HarmonyBridgeEthLogNewUnlockClaim) (event.Subscription, error) {

	logs, sub, err := _HarmonyBridge.contract.WatchLogs(opts, "EthLogNewUnlockClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarmonyBridgeEthLogNewUnlockClaim)
				if err := _HarmonyBridge.contract.UnpackLog(event, "EthLogNewUnlockClaim", log); err != nil {
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

// ParseEthLogNewUnlockClaim is a log parse operation binding the contract event 0x4389bb697e92204e405d58c4811114444ca51e675f1258844e426ce971bf4c56.
//
// Solidity: event EthLogNewUnlockClaim(uint256 _unlockID, address _harmonySender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_HarmonyBridge *HarmonyBridgeFilterer) ParseEthLogNewUnlockClaim(log types.Log) (*HarmonyBridgeEthLogNewUnlockClaim, error) {
	event := new(HarmonyBridgeEthLogNewUnlockClaim)
	if err := _HarmonyBridge.contract.UnpackLog(event, "EthLogNewUnlockClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HarmonyBridgeEthLogUnlockCompletedIterator is returned from FilterEthLogUnlockCompleted and is used to iterate over the raw logs and unpacked data for EthLogUnlockCompleted events raised by the HarmonyBridge contract.
type HarmonyBridgeEthLogUnlockCompletedIterator struct {
	Event *HarmonyBridgeEthLogUnlockCompleted // Event containing the contract specifics and raw log

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
func (it *HarmonyBridgeEthLogUnlockCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HarmonyBridgeEthLogUnlockCompleted)
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
		it.Event = new(HarmonyBridgeEthLogUnlockCompleted)
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
func (it *HarmonyBridgeEthLogUnlockCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HarmonyBridgeEthLogUnlockCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HarmonyBridgeEthLogUnlockCompleted represents a EthLogUnlockCompleted event raised by the HarmonyBridge contract.
type HarmonyBridgeEthLogUnlockCompleted struct {
	UnlockID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthLogUnlockCompleted is a free log retrieval operation binding the contract event 0x56bad42424f5c90e6bfd5e97f55f32bec9a642653ef665d2febfaa6b69e765dc.
//
// Solidity: event EthLogUnlockCompleted(uint256 _unlockID)
func (_HarmonyBridge *HarmonyBridgeFilterer) FilterEthLogUnlockCompleted(opts *bind.FilterOpts) (*HarmonyBridgeEthLogUnlockCompletedIterator, error) {

	logs, sub, err := _HarmonyBridge.contract.FilterLogs(opts, "EthLogUnlockCompleted")
	if err != nil {
		return nil, err
	}
	return &HarmonyBridgeEthLogUnlockCompletedIterator{contract: _HarmonyBridge.contract, event: "EthLogUnlockCompleted", logs: logs, sub: sub}, nil
}

// WatchEthLogUnlockCompleted is a free log subscription operation binding the contract event 0x56bad42424f5c90e6bfd5e97f55f32bec9a642653ef665d2febfaa6b69e765dc.
//
// Solidity: event EthLogUnlockCompleted(uint256 _unlockID)
func (_HarmonyBridge *HarmonyBridgeFilterer) WatchEthLogUnlockCompleted(opts *bind.WatchOpts, sink chan<- *HarmonyBridgeEthLogUnlockCompleted) (event.Subscription, error) {

	logs, sub, err := _HarmonyBridge.contract.WatchLogs(opts, "EthLogUnlockCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HarmonyBridgeEthLogUnlockCompleted)
				if err := _HarmonyBridge.contract.UnpackLog(event, "EthLogUnlockCompleted", log); err != nil {
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

// ParseEthLogUnlockCompleted is a log parse operation binding the contract event 0x56bad42424f5c90e6bfd5e97f55f32bec9a642653ef665d2febfaa6b69e765dc.
//
// Solidity: event EthLogUnlockCompleted(uint256 _unlockID)
func (_HarmonyBridge *HarmonyBridgeFilterer) ParseEthLogUnlockCompleted(log types.Log) (*HarmonyBridgeEthLogUnlockCompleted, error) {
	event := new(HarmonyBridgeEthLogUnlockCompleted)
	if err := _HarmonyBridge.contract.UnpackLog(event, "EthLogUnlockCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}
