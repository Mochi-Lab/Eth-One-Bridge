// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EthereumBridge

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

// EthereumBridgeABI is the input ABI used to generate the binding from.
const EthereumBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ethereumSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"HmyLogNewUnlockClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"HmyLogUnlockCompleted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"HARMONYBRIDGE_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"contractBridgeBank\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeRegistry\",\"outputs\":[{\"internalType\":\"contractBridgeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"completeUnlockClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"isUnlockClaimActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"isUnlockClaimValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumSender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"newUnlockClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockClaimCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unlockClaims\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ethereumSender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originalValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumEthereumBridge.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"contractValset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthereumBridgeBin is the compiled bytecode used for deploying new contracts.
var EthereumBridgeBin = "0x60806040526000805534801561001457600080fd5b50611cec806100246000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80637dc0d1d01161008c578063b1493ce011610066578063b1493ce01461046b578063b1c51a8014610489578063c4d66de8146104a7578063cf93b56c146104eb576100cf565b80637dc0d1d0146103a95780637f54af0c146103f35780639e67206d1461043d576100cf565b80630e41f373146100d45780631a86e3771461011e578063316be17114610164578063570ca735146101ae57806360cf6648146101f85780637c5d759014610286575b600080fd5b6100dc610531565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61014a6004803603602081101561013457600080fd5b8101908080359060200190929190505050610557565b604051808215151515815260200191505060405180910390f35b61016c61059d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101b66105c3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102846004803603608081101561020e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105e9565b005b6102b26004803603602081101561029c57600080fd5b8101908080359060200190929190505050610da6565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182600381111561039057fe5b60ff168152602001965050505050505060405180910390f35b6103b1610e6f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103fb610e95565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104696004803603602081101561045357600080fd5b8101908080359060200190929190505050610ebb565b005b610473611059565b6040518082815260200191505060405180910390f35b61049161105f565b6040518082815260200191505060405180910390f35b6104e9600480360360208110156104bd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611064565b005b6105176004803603602081101561050157600080fd5b8101908080359060200190929190505050611529565b604051808215151515815260200191505060405180910390f35b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001600381111561056657fe5b603a600084815260200190815260200160002060050160009054906101000a900460ff16600381111561059557fe5b149050919050565b603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663833b1fce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561066957600080fd5b505afa15801561067d573d6000803e3d6000fd5b505050506040513d602081101561069357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141580156107975750600073ffffffffffffffffffffffffffffffffffffffff16603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d33f397e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561074357600080fd5b505afa158015610757573d6000803e3d6000fd5b505050506040513d602081101561076d57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614155b6107ec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526046815260200180611be36046913960600191505060405180910390fd5b603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561088b57600080fd5b505afa15801561089f573d6000803e3d6000fd5b505050506040513d60208110156108b557600080fd5b8101908080519060200190929190505050610938576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b4bfd9a783836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060206040518083038186803b1580156109df57600080fd5b505afa1580156109f3573d6000803e3d6000fd5b505050506040513d6020811015610a0957600080fd5b8101908080519060200190929190505050610a6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180611c57603a913960400191505060405180910390fd5b610a77611b49565b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160016003811115610b0557fe5b8152509050610b20600160345461164290919063ffffffff16565b60348190555080603a6000603454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050160006101000a81548160ff02191690836003811115610c8657fe5b02179055509050507faeac38f4f561543d773d127aadc54cc1c83684be65b0ecf15dcc0232fadf14496034548686338787604051808781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405180910390a15050505050565b603a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff16905086565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80610ec581610557565b610f37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f556e6c6f636b20636c61696d206973206e6f742061637469766500000000000081525060200191505060405180910390fd5b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fdd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180611c916027913960400191505060405180910390fd5b6002603a600084815260200190815260200160002060050160006101000a81548160ff0219169083600381111561101057fe5b021790555061101e826116ca565b7face8053805be17e29b4a9ca915b3a6d15949a53b5733a1fb027be2bdff393cfb826040518082815260200191505060405180910390a15050565b60345481565b600181565b600061106e611b2f565b9050600160009054906101000a900460ff168061108f575061108e611b38565b5b8061109b575060005481115b6110f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611c29602e913960400191505060405180910390fd5b6000600160009054906101000a900460ff16159050801561112d5760018060006101000a81548160ff021916908315150217905550816000819055505b82603860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b1580156111d657600080fd5b505afa1580156111ea573d6000803e3d6000fd5b505050506040513d602081101561120057600080fd5b8101908080519060200190929190505050603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fd168ec6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112b957600080fd5b505afa1580156112cd573d6000803e3d6000fd5b505050506040513d60208110156112e357600080fd5b8101908080519060200190929190505050603660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d33f397e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561139c57600080fd5b505afa1580156113b0573d6000803e3d6000fd5b505050506040513d60208110156113c657600080fd5b8101908080519060200190929190505050603760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663833b1fce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561147f57600080fd5b505afa158015611493573d6000803e3d6000fd5b505050506040513d60208110156114a957600080fd5b8101908080519060200190929190505050603960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006034819055508015611524576000600160006101000a81548160ff0219169083151502179055505b505050565b6000603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c603a600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561160057600080fd5b505afa158015611614573d6000803e3d6000fd5b505050506040513d602081101561162a57600080fd5b81019080805190602001909291905050509050919050565b6000808284019050838110156116c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6116d2611b49565b603a60008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820160009054906101000a900460ff16600381111561187057fe5b600381111561187b57fe5b815250509050603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639d05974b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156118e957600080fd5b505afa1580156118fd573d6000803e3d6000fd5b505050506040513d602081101561191357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16816060015173ffffffffffffffffffffffffffffffffffffffff161415611a2957603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166367c2c018826020015183608001516040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611a0c57600080fd5b505af1158015611a20573d6000803e3d6000fd5b50505050611b2b565b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633c298e788260200151836060015184608001516040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611b1257600080fd5b505af1158015611b26573d6000803e3d6000fd5b505050505b5050565b60006001905090565b600080303b90506000811491505090565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160006003811115611bdc57fe5b8152509056fe546865204f70657261746f72206d7573742073657420746865206f7261636c6520616e64206272696467652062616e6b20666f72206272696467652061637469766174696f6e436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65644e6f7420656e6f756768206c6f636b65642061737365747320746f20636f6d706c657465207468652070726f706f7365642070726f70686563794f6e6c7920746865204f7261636c65206d617920636f6d706c6574652070726f70686563696573a265627a7a72315820e6164ae65fe6fa4f204654df5b4c0a8e1157d9754ce8c538af11a19efd902c6a64736f6c63430005110032"

// DeployEthereumBridge deploys a new Ethereum contract, binding an instance of EthereumBridge to it.
func DeployEthereumBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthereumBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthereumBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthereumBridge{EthereumBridgeCaller: EthereumBridgeCaller{contract: contract}, EthereumBridgeTransactor: EthereumBridgeTransactor{contract: contract}, EthereumBridgeFilterer: EthereumBridgeFilterer{contract: contract}}, nil
}

// EthereumBridge is an auto generated Go binding around an Ethereum contract.
type EthereumBridge struct {
	EthereumBridgeCaller     // Read-only binding to the contract
	EthereumBridgeTransactor // Write-only binding to the contract
	EthereumBridgeFilterer   // Log filterer for contract events
}

// EthereumBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumBridgeSession struct {
	Contract     *EthereumBridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumBridgeCallerSession struct {
	Contract *EthereumBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EthereumBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumBridgeTransactorSession struct {
	Contract     *EthereumBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EthereumBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumBridgeRaw struct {
	Contract *EthereumBridge // Generic contract binding to access the raw methods on
}

// EthereumBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumBridgeCallerRaw struct {
	Contract *EthereumBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumBridgeTransactorRaw struct {
	Contract *EthereumBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereumBridge creates a new instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridge(address common.Address, backend bind.ContractBackend) (*EthereumBridge, error) {
	contract, err := bindEthereumBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthereumBridge{EthereumBridgeCaller: EthereumBridgeCaller{contract: contract}, EthereumBridgeTransactor: EthereumBridgeTransactor{contract: contract}, EthereumBridgeFilterer: EthereumBridgeFilterer{contract: contract}}, nil
}

// NewEthereumBridgeCaller creates a new read-only instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridgeCaller(address common.Address, caller bind.ContractCaller) (*EthereumBridgeCaller, error) {
	contract, err := bindEthereumBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeCaller{contract: contract}, nil
}

// NewEthereumBridgeTransactor creates a new write-only instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumBridgeTransactor, error) {
	contract, err := bindEthereumBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTransactor{contract: contract}, nil
}

// NewEthereumBridgeFilterer creates a new log filterer instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumBridgeFilterer, error) {
	contract, err := bindEthereumBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeFilterer{contract: contract}, nil
}

// bindEthereumBridge binds a generic wrapper to an already deployed contract.
func bindEthereumBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumBridge *EthereumBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthereumBridge.Contract.EthereumBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumBridge *EthereumBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.Contract.EthereumBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumBridge *EthereumBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumBridge.Contract.EthereumBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumBridge *EthereumBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthereumBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumBridge *EthereumBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumBridge *EthereumBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumBridge.Contract.contract.Transact(opts, method, params...)
}

// HARMONYBRIDGEREVISION is a free data retrieval call binding the contract method 0xb1c51a80.
//
// Solidity: function HARMONYBRIDGE_REVISION() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) HARMONYBRIDGEREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "HARMONYBRIDGE_REVISION")
	return *ret0, err
}

// HARMONYBRIDGEREVISION is a free data retrieval call binding the contract method 0xb1c51a80.
//
// Solidity: function HARMONYBRIDGE_REVISION() view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) HARMONYBRIDGEREVISION() (*big.Int, error) {
	return _EthereumBridge.Contract.HARMONYBRIDGEREVISION(&_EthereumBridge.CallOpts)
}

// HARMONYBRIDGEREVISION is a free data retrieval call binding the contract method 0xb1c51a80.
//
// Solidity: function HARMONYBRIDGE_REVISION() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) HARMONYBRIDGEREVISION() (*big.Int, error) {
	return _EthereumBridge.Contract.HARMONYBRIDGEREVISION(&_EthereumBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "bridgeBank")
	return *ret0, err
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) BridgeBank() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeBank(&_EthereumBridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) BridgeBank() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeBank(&_EthereumBridge.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) BridgeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "bridgeRegistry")
	return *ret0, err
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) BridgeRegistry() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeRegistry(&_EthereumBridge.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) BridgeRegistry() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeRegistry(&_EthereumBridge.CallOpts)
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_EthereumBridge *EthereumBridgeCaller) IsUnlockClaimActive(opts *bind.CallOpts, _unlockID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "isUnlockClaimActive", _unlockID)
	return *ret0, err
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_EthereumBridge *EthereumBridgeSession) IsUnlockClaimActive(_unlockID *big.Int) (bool, error) {
	return _EthereumBridge.Contract.IsUnlockClaimActive(&_EthereumBridge.CallOpts, _unlockID)
}

// IsUnlockClaimActive is a free data retrieval call binding the contract method 0x1a86e377.
//
// Solidity: function isUnlockClaimActive(uint256 _unlockID) view returns(bool)
func (_EthereumBridge *EthereumBridgeCallerSession) IsUnlockClaimActive(_unlockID *big.Int) (bool, error) {
	return _EthereumBridge.Contract.IsUnlockClaimActive(&_EthereumBridge.CallOpts, _unlockID)
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_EthereumBridge *EthereumBridgeCaller) IsUnlockClaimValidatorActive(opts *bind.CallOpts, _unlockID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "isUnlockClaimValidatorActive", _unlockID)
	return *ret0, err
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_EthereumBridge *EthereumBridgeSession) IsUnlockClaimValidatorActive(_unlockID *big.Int) (bool, error) {
	return _EthereumBridge.Contract.IsUnlockClaimValidatorActive(&_EthereumBridge.CallOpts, _unlockID)
}

// IsUnlockClaimValidatorActive is a free data retrieval call binding the contract method 0xcf93b56c.
//
// Solidity: function isUnlockClaimValidatorActive(uint256 _unlockID) view returns(bool)
func (_EthereumBridge *EthereumBridgeCallerSession) IsUnlockClaimValidatorActive(_unlockID *big.Int) (bool, error) {
	return _EthereumBridge.Contract.IsUnlockClaimValidatorActive(&_EthereumBridge.CallOpts, _unlockID)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) Operator() (common.Address, error) {
	return _EthereumBridge.Contract.Operator(&_EthereumBridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) Operator() (common.Address, error) {
	return _EthereumBridge.Contract.Operator(&_EthereumBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) Oracle() (common.Address, error) {
	return _EthereumBridge.Contract.Oracle(&_EthereumBridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) Oracle() (common.Address, error) {
	return _EthereumBridge.Contract.Oracle(&_EthereumBridge.CallOpts)
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) UnlockClaimCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "unlockClaimCount")
	return *ret0, err
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) UnlockClaimCount() (*big.Int, error) {
	return _EthereumBridge.Contract.UnlockClaimCount(&_EthereumBridge.CallOpts)
}

// UnlockClaimCount is a free data retrieval call binding the contract method 0xb1493ce0.
//
// Solidity: function unlockClaimCount() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) UnlockClaimCount() (*big.Int, error) {
	return _EthereumBridge.Contract.UnlockClaimCount(&_EthereumBridge.CallOpts)
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address ethereumSender, address harmonyReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_EthereumBridge *EthereumBridgeCaller) UnlockClaims(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EthereumSender    common.Address
	HarmonyReceiver   common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	ret := new(struct {
		EthereumSender    common.Address
		HarmonyReceiver   common.Address
		OriginalValidator common.Address
		Token             common.Address
		Amount            *big.Int
		Status            uint8
	})
	out := ret
	err := _EthereumBridge.contract.Call(opts, out, "unlockClaims", arg0)
	return *ret, err
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address ethereumSender, address harmonyReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_EthereumBridge *EthereumBridgeSession) UnlockClaims(arg0 *big.Int) (struct {
	EthereumSender    common.Address
	HarmonyReceiver   common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	return _EthereumBridge.Contract.UnlockClaims(&_EthereumBridge.CallOpts, arg0)
}

// UnlockClaims is a free data retrieval call binding the contract method 0x7c5d7590.
//
// Solidity: function unlockClaims(uint256 ) view returns(address ethereumSender, address harmonyReceiver, address originalValidator, address token, uint256 amount, uint8 status)
func (_EthereumBridge *EthereumBridgeCallerSession) UnlockClaims(arg0 *big.Int) (struct {
	EthereumSender    common.Address
	HarmonyReceiver   common.Address
	OriginalValidator common.Address
	Token             common.Address
	Amount            *big.Int
	Status            uint8
}, error) {
	return _EthereumBridge.Contract.UnlockClaims(&_EthereumBridge.CallOpts, arg0)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthereumBridge.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) Valset() (common.Address, error) {
	return _EthereumBridge.Contract.Valset(&_EthereumBridge.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) Valset() (common.Address, error) {
	return _EthereumBridge.Contract.Valset(&_EthereumBridge.CallOpts)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_EthereumBridge *EthereumBridgeTransactor) CompleteUnlockClaim(opts *bind.TransactOpts, _unlockID *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "completeUnlockClaim", _unlockID)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_EthereumBridge *EthereumBridgeSession) CompleteUnlockClaim(_unlockID *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.CompleteUnlockClaim(&_EthereumBridge.TransactOpts, _unlockID)
}

// CompleteUnlockClaim is a paid mutator transaction binding the contract method 0x9e67206d.
//
// Solidity: function completeUnlockClaim(uint256 _unlockID) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) CompleteUnlockClaim(_unlockID *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.CompleteUnlockClaim(&_EthereumBridge.TransactOpts, _unlockID)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeRegistry) returns()
func (_EthereumBridge *EthereumBridgeTransactor) Initialize(opts *bind.TransactOpts, _bridgeRegistry common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "initialize", _bridgeRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeRegistry) returns()
func (_EthereumBridge *EthereumBridgeSession) Initialize(_bridgeRegistry common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.Initialize(&_EthereumBridge.TransactOpts, _bridgeRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeRegistry) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) Initialize(_bridgeRegistry common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.Initialize(&_EthereumBridge.TransactOpts, _bridgeRegistry)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _ethereumSender, address _harmonyReceiver, address _token, uint256 _amount) returns()
func (_EthereumBridge *EthereumBridgeTransactor) NewUnlockClaim(opts *bind.TransactOpts, _ethereumSender common.Address, _harmonyReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "newUnlockClaim", _ethereumSender, _harmonyReceiver, _token, _amount)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _ethereumSender, address _harmonyReceiver, address _token, uint256 _amount) returns()
func (_EthereumBridge *EthereumBridgeSession) NewUnlockClaim(_ethereumSender common.Address, _harmonyReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.NewUnlockClaim(&_EthereumBridge.TransactOpts, _ethereumSender, _harmonyReceiver, _token, _amount)
}

// NewUnlockClaim is a paid mutator transaction binding the contract method 0x60cf6648.
//
// Solidity: function newUnlockClaim(address _ethereumSender, address _harmonyReceiver, address _token, uint256 _amount) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) NewUnlockClaim(_ethereumSender common.Address, _harmonyReceiver common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.NewUnlockClaim(&_EthereumBridge.TransactOpts, _ethereumSender, _harmonyReceiver, _token, _amount)
}

// EthereumBridgeHmyLogNewUnlockClaimIterator is returned from FilterHmyLogNewUnlockClaim and is used to iterate over the raw logs and unpacked data for HmyLogNewUnlockClaim events raised by the EthereumBridge contract.
type EthereumBridgeHmyLogNewUnlockClaimIterator struct {
	Event *EthereumBridgeHmyLogNewUnlockClaim // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeHmyLogNewUnlockClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeHmyLogNewUnlockClaim)
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
		it.Event = new(EthereumBridgeHmyLogNewUnlockClaim)
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
func (it *EthereumBridgeHmyLogNewUnlockClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeHmyLogNewUnlockClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeHmyLogNewUnlockClaim represents a HmyLogNewUnlockClaim event raised by the EthereumBridge contract.
type EthereumBridgeHmyLogNewUnlockClaim struct {
	UnlockID         *big.Int
	EthereumSender   common.Address
	HarmonyReceiver  common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHmyLogNewUnlockClaim is a free log retrieval operation binding the contract event 0xaeac38f4f561543d773d127aadc54cc1c83684be65b0ecf15dcc0232fadf1449.
//
// Solidity: event HmyLogNewUnlockClaim(uint256 _unlockID, address _ethereumSender, address _harmonyReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_EthereumBridge *EthereumBridgeFilterer) FilterHmyLogNewUnlockClaim(opts *bind.FilterOpts) (*EthereumBridgeHmyLogNewUnlockClaimIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "HmyLogNewUnlockClaim")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeHmyLogNewUnlockClaimIterator{contract: _EthereumBridge.contract, event: "HmyLogNewUnlockClaim", logs: logs, sub: sub}, nil
}

// WatchHmyLogNewUnlockClaim is a free log subscription operation binding the contract event 0xaeac38f4f561543d773d127aadc54cc1c83684be65b0ecf15dcc0232fadf1449.
//
// Solidity: event HmyLogNewUnlockClaim(uint256 _unlockID, address _ethereumSender, address _harmonyReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_EthereumBridge *EthereumBridgeFilterer) WatchHmyLogNewUnlockClaim(opts *bind.WatchOpts, sink chan<- *EthereumBridgeHmyLogNewUnlockClaim) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "HmyLogNewUnlockClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeHmyLogNewUnlockClaim)
				if err := _EthereumBridge.contract.UnpackLog(event, "HmyLogNewUnlockClaim", log); err != nil {
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

// ParseHmyLogNewUnlockClaim is a log parse operation binding the contract event 0xaeac38f4f561543d773d127aadc54cc1c83684be65b0ecf15dcc0232fadf1449.
//
// Solidity: event HmyLogNewUnlockClaim(uint256 _unlockID, address _ethereumSender, address _harmonyReceiver, address _validatorAddress, address _tokenAddress, uint256 _amount)
func (_EthereumBridge *EthereumBridgeFilterer) ParseHmyLogNewUnlockClaim(log types.Log) (*EthereumBridgeHmyLogNewUnlockClaim, error) {
	event := new(EthereumBridgeHmyLogNewUnlockClaim)
	if err := _EthereumBridge.contract.UnpackLog(event, "HmyLogNewUnlockClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthereumBridgeHmyLogUnlockCompletedIterator is returned from FilterHmyLogUnlockCompleted and is used to iterate over the raw logs and unpacked data for HmyLogUnlockCompleted events raised by the EthereumBridge contract.
type EthereumBridgeHmyLogUnlockCompletedIterator struct {
	Event *EthereumBridgeHmyLogUnlockCompleted // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeHmyLogUnlockCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeHmyLogUnlockCompleted)
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
		it.Event = new(EthereumBridgeHmyLogUnlockCompleted)
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
func (it *EthereumBridgeHmyLogUnlockCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeHmyLogUnlockCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeHmyLogUnlockCompleted represents a HmyLogUnlockCompleted event raised by the EthereumBridge contract.
type EthereumBridgeHmyLogUnlockCompleted struct {
	UnlockID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterHmyLogUnlockCompleted is a free log retrieval operation binding the contract event 0xace8053805be17e29b4a9ca915b3a6d15949a53b5733a1fb027be2bdff393cfb.
//
// Solidity: event HmyLogUnlockCompleted(uint256 _unlockID)
func (_EthereumBridge *EthereumBridgeFilterer) FilterHmyLogUnlockCompleted(opts *bind.FilterOpts) (*EthereumBridgeHmyLogUnlockCompletedIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "HmyLogUnlockCompleted")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeHmyLogUnlockCompletedIterator{contract: _EthereumBridge.contract, event: "HmyLogUnlockCompleted", logs: logs, sub: sub}, nil
}

// WatchHmyLogUnlockCompleted is a free log subscription operation binding the contract event 0xace8053805be17e29b4a9ca915b3a6d15949a53b5733a1fb027be2bdff393cfb.
//
// Solidity: event HmyLogUnlockCompleted(uint256 _unlockID)
func (_EthereumBridge *EthereumBridgeFilterer) WatchHmyLogUnlockCompleted(opts *bind.WatchOpts, sink chan<- *EthereumBridgeHmyLogUnlockCompleted) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "HmyLogUnlockCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeHmyLogUnlockCompleted)
				if err := _EthereumBridge.contract.UnpackLog(event, "HmyLogUnlockCompleted", log); err != nil {
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

// ParseHmyLogUnlockCompleted is a log parse operation binding the contract event 0xace8053805be17e29b4a9ca915b3a6d15949a53b5733a1fb027be2bdff393cfb.
//
// Solidity: event HmyLogUnlockCompleted(uint256 _unlockID)
func (_EthereumBridge *EthereumBridgeFilterer) ParseHmyLogUnlockCompleted(log types.Log) (*EthereumBridgeHmyLogUnlockCompleted, error) {
	event := new(EthereumBridgeHmyLogUnlockCompleted)
	if err := _EthereumBridge.contract.UnpackLog(event, "HmyLogUnlockCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}
