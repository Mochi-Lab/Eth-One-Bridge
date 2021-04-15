// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Valset

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

// ValsetABI is the input ABI used to generate the binding from.
const ValsetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"EthLogValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"EthLogValidatorPowerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"EthLogValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"EthLogValsetReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"EthLogValsetUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALSET_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorPower\",\"type\":\"uint256\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentValsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"getValidatorPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeRegistry\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_initValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_initPowers\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_valsetVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"recoverGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newValidatorPower\",\"type\":\"uint256\"}],\"name\":\"updateValidatorPower\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValsetBin is the compiled bytecode used for deploying new contracts.
var ValsetBin = "0x60806040526000805534801561001457600080fd5b50611ad8806100246000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638d56c37d11610097578063db3ad22c11610066578063db3ad22c146105f8578063ee0b296e14610616578063fc6c1f0214610634578063ff1d57521461068257610100565b80638d56c37d146105045780639bdafcb314610522578063b5672be314610568578063b872c523146105b657610100565b806340a141ff116100d357806340a141ff146102d2578063473691a414610316578063570ca7351461036e578063788cf92f146103b857610100565b80630f43a6771461010557806319045a25146101235780632e75293b1461022857806340550a1c14610276575b600080fd5b61010d6107ee565b6040518082815260200191505060405180910390f35b6101e66004803603604081101561013957600080fd5b81019080803590602001909291908035906020019064010000000081111561016057600080fd5b82018360208201111561017257600080fd5b8035906020019184600183028401116401000000008311171561019457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107f4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102746004803603604081101561023e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610810565b005b6102b86004803603602081101561028c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610aa2565b604051808215151515815260200191505060405180910390f35b610314600480360360208110156102e857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b30565b005b6103586004803603602081101561032c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ddf565b6040518082815260200191505060405180910390f35b610376610e60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610502600480360360408110156103ce57600080fd5b81019080803590602001906401000000008111156103eb57600080fd5b8201836020820111156103fd57600080fd5b8035906020019184602083028401116401000000008311171561041f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561047f57600080fd5b82018360208201111561049157600080fd5b803590602001918460208302840111640100000000831117156104b357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610e86565b005b61050c611057565b6040518082815260200191505060405180910390f35b61054e6004803603602081101561053857600080fd5b810190808035906020019092919050505061105d565b604051808215151515815260200191505060405180910390f35b6105b46004803603604081101561057e57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061107d565b005b6105e2600480360360208110156105cc57600080fd5b810190808035906020019092919050505061123b565b6040518082815260200191505060405180910390f35b610600611253565b6040518082815260200191505060405180910390f35b61061e611259565b6040518082815260200191505060405180910390f35b6106806004803603604081101561064a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061125e565b005b6107ec6004803603606081101561069857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156106d557600080fd5b8201836020820111156106e757600080fd5b8035906020019184602083028401116401000000008311171561070957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561076957600080fd5b82018360208201111561077b57600080fd5b8035906020019184602083028401116401000000008311171561079d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505061132f565b005b60375481565b6000610808610802846114f4565b8361154c565b905092915050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600060365483604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506038600082815260200190815260200160002060009054906101000a900460ff166109ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611a47602e913960400191505060405180910390fd5b6000603960008381526020019081526020016000205490506109da8160355461162e90919063ffffffff16565b6035819055506109f58360355461167890919063ffffffff16565b6035819055508260396000848152602001908152602001600020819055507fc998ed8ee0bce4bf08f1078680045b0dee1d4051e0ffa493f67e47a4187bb2dc8484603654603754603554604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a150505050565b60008060365483604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506038600082815260200190815260200160002060009054906101000a900460ff16915050919050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b600060365482604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506038600082815260200190815260200160002060009054906101000a900460ff16610ccd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806119f16021913960400191505060405180910390fd5b610ce3600160375461162e90919063ffffffff16565b603781905550610d11603960008381526020019081526020016000205460355461162e90919063ffffffff16565b6035819055506038600082815260200190815260200160002060006101000a81549060ff021916905560396000828152602001908152602001600020600090557f123272f4071b21d474981554217c22d548c109443a4222f8e225bd8ae837f891826000603654603754603554604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a15050565b60008060365483604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506039600082815260200190815260200160002054915050919050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f49576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b8051825114610fa3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180611a75602f913960400191505060405180910390fd5b610fab611700565b60008090505b825181101561100557610fea838281518110610fc957fe5b6020026020010151838381518110610fdd57fe5b602002602001015161177b565b610ffe60018261167890919063ffffffff16565b9050610fb1565b507fa263e6a7a37f9afb3cb30027a71f0f16aa7936bab718860b78f2e799aaa0e24960365460375460355460405180848152602001838152602001828152602001935050505060405180910390a15050565b60365481565b60386020528060005260406000206000915054906101000a900460ff1681565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611140576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b603654821061119a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526035815260200180611a126035913960400191505060405180910390fd5b60008282604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506038600082815260200190815260200160002060006101000a81549060ff02191690556039600082815260200190815260200160002060009055505050565b60396020528060005260406000206000915090505481565b60355481565b600181565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611321576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b61132b828261177b565b5050565b60006113396118e8565b9050600160009054906101000a900460ff168061135a57506113596118f1565b5b80611366575060005481115b6113bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001806119c3602e913960400191505060405180910390fd5b6000600160009054906101000a900460ff1615905080156113f85760018060006101000a81548160ff021916908315150217905550816000819055505b8473ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b15801561143e57600080fd5b505afa158015611452573d6000803e3d6000fd5b505050506040513d602081101561146857600080fd5b8101908080519060200190929190505050603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006036819055506114cb8484610e86565b80156114ed576000600160006101000a81548160ff0219169083151502179055505b5050505050565b60008160405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c01828152602001915050604051602081830303815290604052805190602001209050919050565b60008060008060418551146115675760009350505050611628565b6020850151925060408501519150606085015160001a9050601b8160ff16101561159257601b810190505b601b8160ff16141580156115aa5750601c8160ff1614155b156115bb5760009350505050611628565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611618573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600061167083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611902565b905092915050565b6000808284019050838110156116f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b611716600160365461167890919063ffffffff16565b603681905550600060378190555060006035819055507fc6ac622caa95b6a11c41d9b0f1539d060de8883801e4efe5f2b20af915d6ecaf60365460375460355460405180848152602001838152602001828152602001935050505060405180910390a1565b600060365483604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012090506117f5600160375461167890919063ffffffff16565b6037819055506118108260355461167890919063ffffffff16565b60358190555060016038600083815260200190815260200160002060006101000a81548160ff0219169083151502179055508160396000838152602001908152602001600020819055507ffac068c28f4356ccf9adf606e50ebad8b3b0a44915e6cf7627ae28c3416f5d958383603654603754603554604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a1505050565b60006001905090565b600080303b90506000811491505090565b60008383111582906119af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611974578082015181840152602081019050611959565b50505050905090810190601f1680156119a15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050939250505056fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a656443616e206f6e6c792072656d6f7665206163746976652076616c646961746f7273476173207265636f76657279206f6e6c7920616c6c6f77656420666f722070726576696f75732076616c696461746f72207365747343616e206f6e6c79207570646174652074686520706f776572206f66206163746976652076616c646961746f727345766572792076616c696461746f72206d7573742068617665206120636f72726573706f6e64696e6720706f776572a265627a7a72315820c06989f78db0d1c75f455cb162d60eb734e0f33b48b8b14c37200d9220570e0664736f6c63430005110032"

// DeployValset deploys a new Ethereum contract, binding an instance of Valset to it.
func DeployValset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Valset, error) {
	parsed, err := abi.JSON(strings.NewReader(ValsetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValsetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Valset{ValsetCaller: ValsetCaller{contract: contract}, ValsetTransactor: ValsetTransactor{contract: contract}, ValsetFilterer: ValsetFilterer{contract: contract}}, nil
}

// Valset is an auto generated Go binding around an Ethereum contract.
type Valset struct {
	ValsetCaller     // Read-only binding to the contract
	ValsetTransactor // Write-only binding to the contract
	ValsetFilterer   // Log filterer for contract events
}

// ValsetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValsetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValsetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValsetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValsetSession struct {
	Contract     *Valset           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValsetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValsetCallerSession struct {
	Contract *ValsetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValsetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValsetTransactorSession struct {
	Contract     *ValsetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValsetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValsetRaw struct {
	Contract *Valset // Generic contract binding to access the raw methods on
}

// ValsetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValsetCallerRaw struct {
	Contract *ValsetCaller // Generic read-only contract binding to access the raw methods on
}

// ValsetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValsetTransactorRaw struct {
	Contract *ValsetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValset creates a new instance of Valset, bound to a specific deployed contract.
func NewValset(address common.Address, backend bind.ContractBackend) (*Valset, error) {
	contract, err := bindValset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Valset{ValsetCaller: ValsetCaller{contract: contract}, ValsetTransactor: ValsetTransactor{contract: contract}, ValsetFilterer: ValsetFilterer{contract: contract}}, nil
}

// NewValsetCaller creates a new read-only instance of Valset, bound to a specific deployed contract.
func NewValsetCaller(address common.Address, caller bind.ContractCaller) (*ValsetCaller, error) {
	contract, err := bindValset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValsetCaller{contract: contract}, nil
}

// NewValsetTransactor creates a new write-only instance of Valset, bound to a specific deployed contract.
func NewValsetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValsetTransactor, error) {
	contract, err := bindValset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValsetTransactor{contract: contract}, nil
}

// NewValsetFilterer creates a new log filterer instance of Valset, bound to a specific deployed contract.
func NewValsetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValsetFilterer, error) {
	contract, err := bindValset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValsetFilterer{contract: contract}, nil
}

// bindValset binds a generic wrapper to an already deployed contract.
func bindValset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValsetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Valset *ValsetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Valset.Contract.ValsetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Valset *ValsetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Valset.Contract.ValsetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Valset *ValsetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Valset.Contract.ValsetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Valset *ValsetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Valset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Valset *ValsetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Valset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Valset *ValsetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Valset.Contract.contract.Transact(opts, method, params...)
}

// VALSETREVISION is a free data retrieval call binding the contract method 0xee0b296e.
//
// Solidity: function VALSET_REVISION() view returns(uint256)
func (_Valset *ValsetCaller) VALSETREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "VALSET_REVISION")
	return *ret0, err
}

// VALSETREVISION is a free data retrieval call binding the contract method 0xee0b296e.
//
// Solidity: function VALSET_REVISION() view returns(uint256)
func (_Valset *ValsetSession) VALSETREVISION() (*big.Int, error) {
	return _Valset.Contract.VALSETREVISION(&_Valset.CallOpts)
}

// VALSETREVISION is a free data retrieval call binding the contract method 0xee0b296e.
//
// Solidity: function VALSET_REVISION() view returns(uint256)
func (_Valset *ValsetCallerSession) VALSETREVISION() (*big.Int, error) {
	return _Valset.Contract.VALSETREVISION(&_Valset.CallOpts)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_Valset *ValsetCaller) CurrentValsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "currentValsetVersion")
	return *ret0, err
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_Valset *ValsetSession) CurrentValsetVersion() (*big.Int, error) {
	return _Valset.Contract.CurrentValsetVersion(&_Valset.CallOpts)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_Valset *ValsetCallerSession) CurrentValsetVersion() (*big.Int, error) {
	return _Valset.Contract.CurrentValsetVersion(&_Valset.CallOpts)
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_Valset *ValsetCaller) GetValidatorPower(opts *bind.CallOpts, _validatorAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "getValidatorPower", _validatorAddress)
	return *ret0, err
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_Valset *ValsetSession) GetValidatorPower(_validatorAddress common.Address) (*big.Int, error) {
	return _Valset.Contract.GetValidatorPower(&_Valset.CallOpts, _validatorAddress)
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_Valset *ValsetCallerSession) GetValidatorPower(_validatorAddress common.Address) (*big.Int, error) {
	return _Valset.Contract.GetValidatorPower(&_Valset.CallOpts, _validatorAddress)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_Valset *ValsetCaller) IsActiveValidator(opts *bind.CallOpts, _validatorAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "isActiveValidator", _validatorAddress)
	return *ret0, err
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_Valset *ValsetSession) IsActiveValidator(_validatorAddress common.Address) (bool, error) {
	return _Valset.Contract.IsActiveValidator(&_Valset.CallOpts, _validatorAddress)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_Valset *ValsetCallerSession) IsActiveValidator(_validatorAddress common.Address) (bool, error) {
	return _Valset.Contract.IsActiveValidator(&_Valset.CallOpts, _validatorAddress)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Valset *ValsetCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Valset *ValsetSession) Operator() (common.Address, error) {
	return _Valset.Contract.Operator(&_Valset.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Valset *ValsetCallerSession) Operator() (common.Address, error) {
	return _Valset.Contract.Operator(&_Valset.CallOpts)
}

// Powers is a free data retrieval call binding the contract method 0xb872c523.
//
// Solidity: function powers(bytes32 ) view returns(uint256)
func (_Valset *ValsetCaller) Powers(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "powers", arg0)
	return *ret0, err
}

// Powers is a free data retrieval call binding the contract method 0xb872c523.
//
// Solidity: function powers(bytes32 ) view returns(uint256)
func (_Valset *ValsetSession) Powers(arg0 [32]byte) (*big.Int, error) {
	return _Valset.Contract.Powers(&_Valset.CallOpts, arg0)
}

// Powers is a free data retrieval call binding the contract method 0xb872c523.
//
// Solidity: function powers(bytes32 ) view returns(uint256)
func (_Valset *ValsetCallerSession) Powers(arg0 [32]byte) (*big.Int, error) {
	return _Valset.Contract.Powers(&_Valset.CallOpts, arg0)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _message, bytes _signature) pure returns(address)
func (_Valset *ValsetCaller) Recover(opts *bind.CallOpts, _message [32]byte, _signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "recover", _message, _signature)
	return *ret0, err
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _message, bytes _signature) pure returns(address)
func (_Valset *ValsetSession) Recover(_message [32]byte, _signature []byte) (common.Address, error) {
	return _Valset.Contract.Recover(&_Valset.CallOpts, _message, _signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _message, bytes _signature) pure returns(address)
func (_Valset *ValsetCallerSession) Recover(_message [32]byte, _signature []byte) (common.Address, error) {
	return _Valset.Contract.Recover(&_Valset.CallOpts, _message, _signature)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Valset *ValsetCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "totalPower")
	return *ret0, err
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Valset *ValsetSession) TotalPower() (*big.Int, error) {
	return _Valset.Contract.TotalPower(&_Valset.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Valset *ValsetCallerSession) TotalPower() (*big.Int, error) {
	return _Valset.Contract.TotalPower(&_Valset.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Valset *ValsetCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "validatorCount")
	return *ret0, err
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Valset *ValsetSession) ValidatorCount() (*big.Int, error) {
	return _Valset.Contract.ValidatorCount(&_Valset.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Valset *ValsetCallerSession) ValidatorCount() (*big.Int, error) {
	return _Valset.Contract.ValidatorCount(&_Valset.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(bool)
func (_Valset *ValsetCaller) Validators(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Valset.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(bool)
func (_Valset *ValsetSession) Validators(arg0 [32]byte) (bool, error) {
	return _Valset.Contract.Validators(&_Valset.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(bool)
func (_Valset *ValsetCallerSession) Validators(arg0 [32]byte) (bool, error) {
	return _Valset.Contract.Validators(&_Valset.CallOpts, arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_Valset *ValsetTransactor) AddValidator(opts *bind.TransactOpts, _validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "addValidator", _validatorAddress, _validatorPower)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_Valset *ValsetSession) AddValidator(_validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.AddValidator(&_Valset.TransactOpts, _validatorAddress, _validatorPower)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_Valset *ValsetTransactorSession) AddValidator(_validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.AddValidator(&_Valset.TransactOpts, _validatorAddress, _validatorPower)
}

// Initialize is a paid mutator transaction binding the contract method 0xff1d5752.
//
// Solidity: function initialize(address _bridgeRegistry, address[] _initValidators, uint256[] _initPowers) returns()
func (_Valset *ValsetTransactor) Initialize(opts *bind.TransactOpts, _bridgeRegistry common.Address, _initValidators []common.Address, _initPowers []*big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "initialize", _bridgeRegistry, _initValidators, _initPowers)
}

// Initialize is a paid mutator transaction binding the contract method 0xff1d5752.
//
// Solidity: function initialize(address _bridgeRegistry, address[] _initValidators, uint256[] _initPowers) returns()
func (_Valset *ValsetSession) Initialize(_bridgeRegistry common.Address, _initValidators []common.Address, _initPowers []*big.Int) (*types.Transaction, error) {
	return _Valset.Contract.Initialize(&_Valset.TransactOpts, _bridgeRegistry, _initValidators, _initPowers)
}

// Initialize is a paid mutator transaction binding the contract method 0xff1d5752.
//
// Solidity: function initialize(address _bridgeRegistry, address[] _initValidators, uint256[] _initPowers) returns()
func (_Valset *ValsetTransactorSession) Initialize(_bridgeRegistry common.Address, _initValidators []common.Address, _initPowers []*big.Int) (*types.Transaction, error) {
	return _Valset.Contract.Initialize(&_Valset.TransactOpts, _bridgeRegistry, _initValidators, _initPowers)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_Valset *ValsetTransactor) RecoverGas(opts *bind.TransactOpts, _valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "recoverGas", _valsetVersion, _validatorAddress)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_Valset *ValsetSession) RecoverGas(_valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RecoverGas(&_Valset.TransactOpts, _valsetVersion, _validatorAddress)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_Valset *ValsetTransactorSession) RecoverGas(_valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RecoverGas(&_Valset.TransactOpts, _valsetVersion, _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_Valset *ValsetTransactor) RemoveValidator(opts *bind.TransactOpts, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "removeValidator", _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_Valset *ValsetSession) RemoveValidator(_validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RemoveValidator(&_Valset.TransactOpts, _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_Valset *ValsetTransactorSession) RemoveValidator(_validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RemoveValidator(&_Valset.TransactOpts, _validatorAddress)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_Valset *ValsetTransactor) UpdateValidatorPower(opts *bind.TransactOpts, _validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "updateValidatorPower", _validatorAddress, _newValidatorPower)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_Valset *ValsetSession) UpdateValidatorPower(_validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValidatorPower(&_Valset.TransactOpts, _validatorAddress, _newValidatorPower)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_Valset *ValsetTransactorSession) UpdateValidatorPower(_validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValidatorPower(&_Valset.TransactOpts, _validatorAddress, _newValidatorPower)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_Valset *ValsetTransactor) UpdateValset(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "updateValset", _validators, _powers)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_Valset *ValsetSession) UpdateValset(_validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValset(&_Valset.TransactOpts, _validators, _powers)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_Valset *ValsetTransactorSession) UpdateValset(_validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValset(&_Valset.TransactOpts, _validators, _powers)
}

// ValsetEthLogValidatorAddedIterator is returned from FilterEthLogValidatorAdded and is used to iterate over the raw logs and unpacked data for EthLogValidatorAdded events raised by the Valset contract.
type ValsetEthLogValidatorAddedIterator struct {
	Event *ValsetEthLogValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ValsetEthLogValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetEthLogValidatorAdded)
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
		it.Event = new(ValsetEthLogValidatorAdded)
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
func (it *ValsetEthLogValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetEthLogValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetEthLogValidatorAdded represents a EthLogValidatorAdded event raised by the Valset contract.
type ValsetEthLogValidatorAdded struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEthLogValidatorAdded is a free log retrieval operation binding the contract event 0xfac068c28f4356ccf9adf606e50ebad8b3b0a44915e6cf7627ae28c3416f5d95.
//
// Solidity: event EthLogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterEthLogValidatorAdded(opts *bind.FilterOpts) (*ValsetEthLogValidatorAddedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "EthLogValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &ValsetEthLogValidatorAddedIterator{contract: _Valset.contract, event: "EthLogValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchEthLogValidatorAdded is a free log subscription operation binding the contract event 0xfac068c28f4356ccf9adf606e50ebad8b3b0a44915e6cf7627ae28c3416f5d95.
//
// Solidity: event EthLogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchEthLogValidatorAdded(opts *bind.WatchOpts, sink chan<- *ValsetEthLogValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "EthLogValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetEthLogValidatorAdded)
				if err := _Valset.contract.UnpackLog(event, "EthLogValidatorAdded", log); err != nil {
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

// ParseEthLogValidatorAdded is a log parse operation binding the contract event 0xfac068c28f4356ccf9adf606e50ebad8b3b0a44915e6cf7627ae28c3416f5d95.
//
// Solidity: event EthLogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseEthLogValidatorAdded(log types.Log) (*ValsetEthLogValidatorAdded, error) {
	event := new(ValsetEthLogValidatorAdded)
	if err := _Valset.contract.UnpackLog(event, "EthLogValidatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValsetEthLogValidatorPowerUpdatedIterator is returned from FilterEthLogValidatorPowerUpdated and is used to iterate over the raw logs and unpacked data for EthLogValidatorPowerUpdated events raised by the Valset contract.
type ValsetEthLogValidatorPowerUpdatedIterator struct {
	Event *ValsetEthLogValidatorPowerUpdated // Event containing the contract specifics and raw log

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
func (it *ValsetEthLogValidatorPowerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetEthLogValidatorPowerUpdated)
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
		it.Event = new(ValsetEthLogValidatorPowerUpdated)
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
func (it *ValsetEthLogValidatorPowerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetEthLogValidatorPowerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetEthLogValidatorPowerUpdated represents a EthLogValidatorPowerUpdated event raised by the Valset contract.
type ValsetEthLogValidatorPowerUpdated struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEthLogValidatorPowerUpdated is a free log retrieval operation binding the contract event 0xc998ed8ee0bce4bf08f1078680045b0dee1d4051e0ffa493f67e47a4187bb2dc.
//
// Solidity: event EthLogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterEthLogValidatorPowerUpdated(opts *bind.FilterOpts) (*ValsetEthLogValidatorPowerUpdatedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "EthLogValidatorPowerUpdated")
	if err != nil {
		return nil, err
	}
	return &ValsetEthLogValidatorPowerUpdatedIterator{contract: _Valset.contract, event: "EthLogValidatorPowerUpdated", logs: logs, sub: sub}, nil
}

// WatchEthLogValidatorPowerUpdated is a free log subscription operation binding the contract event 0xc998ed8ee0bce4bf08f1078680045b0dee1d4051e0ffa493f67e47a4187bb2dc.
//
// Solidity: event EthLogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchEthLogValidatorPowerUpdated(opts *bind.WatchOpts, sink chan<- *ValsetEthLogValidatorPowerUpdated) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "EthLogValidatorPowerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetEthLogValidatorPowerUpdated)
				if err := _Valset.contract.UnpackLog(event, "EthLogValidatorPowerUpdated", log); err != nil {
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

// ParseEthLogValidatorPowerUpdated is a log parse operation binding the contract event 0xc998ed8ee0bce4bf08f1078680045b0dee1d4051e0ffa493f67e47a4187bb2dc.
//
// Solidity: event EthLogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseEthLogValidatorPowerUpdated(log types.Log) (*ValsetEthLogValidatorPowerUpdated, error) {
	event := new(ValsetEthLogValidatorPowerUpdated)
	if err := _Valset.contract.UnpackLog(event, "EthLogValidatorPowerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValsetEthLogValidatorRemovedIterator is returned from FilterEthLogValidatorRemoved and is used to iterate over the raw logs and unpacked data for EthLogValidatorRemoved events raised by the Valset contract.
type ValsetEthLogValidatorRemovedIterator struct {
	Event *ValsetEthLogValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ValsetEthLogValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetEthLogValidatorRemoved)
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
		it.Event = new(ValsetEthLogValidatorRemoved)
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
func (it *ValsetEthLogValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetEthLogValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetEthLogValidatorRemoved represents a EthLogValidatorRemoved event raised by the Valset contract.
type ValsetEthLogValidatorRemoved struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEthLogValidatorRemoved is a free log retrieval operation binding the contract event 0x123272f4071b21d474981554217c22d548c109443a4222f8e225bd8ae837f891.
//
// Solidity: event EthLogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterEthLogValidatorRemoved(opts *bind.FilterOpts) (*ValsetEthLogValidatorRemovedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "EthLogValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &ValsetEthLogValidatorRemovedIterator{contract: _Valset.contract, event: "EthLogValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchEthLogValidatorRemoved is a free log subscription operation binding the contract event 0x123272f4071b21d474981554217c22d548c109443a4222f8e225bd8ae837f891.
//
// Solidity: event EthLogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchEthLogValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValsetEthLogValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "EthLogValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetEthLogValidatorRemoved)
				if err := _Valset.contract.UnpackLog(event, "EthLogValidatorRemoved", log); err != nil {
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

// ParseEthLogValidatorRemoved is a log parse operation binding the contract event 0x123272f4071b21d474981554217c22d548c109443a4222f8e225bd8ae837f891.
//
// Solidity: event EthLogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseEthLogValidatorRemoved(log types.Log) (*ValsetEthLogValidatorRemoved, error) {
	event := new(ValsetEthLogValidatorRemoved)
	if err := _Valset.contract.UnpackLog(event, "EthLogValidatorRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValsetEthLogValsetResetIterator is returned from FilterEthLogValsetReset and is used to iterate over the raw logs and unpacked data for EthLogValsetReset events raised by the Valset contract.
type ValsetEthLogValsetResetIterator struct {
	Event *ValsetEthLogValsetReset // Event containing the contract specifics and raw log

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
func (it *ValsetEthLogValsetResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetEthLogValsetReset)
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
		it.Event = new(ValsetEthLogValsetReset)
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
func (it *ValsetEthLogValsetResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetEthLogValsetResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetEthLogValsetReset represents a EthLogValsetReset event raised by the Valset contract.
type ValsetEthLogValsetReset struct {
	NewValsetVersion *big.Int
	ValidatorCount   *big.Int
	TotalPower       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEthLogValsetReset is a free log retrieval operation binding the contract event 0xc6ac622caa95b6a11c41d9b0f1539d060de8883801e4efe5f2b20af915d6ecaf.
//
// Solidity: event EthLogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterEthLogValsetReset(opts *bind.FilterOpts) (*ValsetEthLogValsetResetIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "EthLogValsetReset")
	if err != nil {
		return nil, err
	}
	return &ValsetEthLogValsetResetIterator{contract: _Valset.contract, event: "EthLogValsetReset", logs: logs, sub: sub}, nil
}

// WatchEthLogValsetReset is a free log subscription operation binding the contract event 0xc6ac622caa95b6a11c41d9b0f1539d060de8883801e4efe5f2b20af915d6ecaf.
//
// Solidity: event EthLogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchEthLogValsetReset(opts *bind.WatchOpts, sink chan<- *ValsetEthLogValsetReset) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "EthLogValsetReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetEthLogValsetReset)
				if err := _Valset.contract.UnpackLog(event, "EthLogValsetReset", log); err != nil {
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

// ParseEthLogValsetReset is a log parse operation binding the contract event 0xc6ac622caa95b6a11c41d9b0f1539d060de8883801e4efe5f2b20af915d6ecaf.
//
// Solidity: event EthLogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseEthLogValsetReset(log types.Log) (*ValsetEthLogValsetReset, error) {
	event := new(ValsetEthLogValsetReset)
	if err := _Valset.contract.UnpackLog(event, "EthLogValsetReset", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValsetEthLogValsetUpdatedIterator is returned from FilterEthLogValsetUpdated and is used to iterate over the raw logs and unpacked data for EthLogValsetUpdated events raised by the Valset contract.
type ValsetEthLogValsetUpdatedIterator struct {
	Event *ValsetEthLogValsetUpdated // Event containing the contract specifics and raw log

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
func (it *ValsetEthLogValsetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetEthLogValsetUpdated)
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
		it.Event = new(ValsetEthLogValsetUpdated)
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
func (it *ValsetEthLogValsetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetEthLogValsetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetEthLogValsetUpdated represents a EthLogValsetUpdated event raised by the Valset contract.
type ValsetEthLogValsetUpdated struct {
	NewValsetVersion *big.Int
	ValidatorCount   *big.Int
	TotalPower       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEthLogValsetUpdated is a free log retrieval operation binding the contract event 0xa263e6a7a37f9afb3cb30027a71f0f16aa7936bab718860b78f2e799aaa0e249.
//
// Solidity: event EthLogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterEthLogValsetUpdated(opts *bind.FilterOpts) (*ValsetEthLogValsetUpdatedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "EthLogValsetUpdated")
	if err != nil {
		return nil, err
	}
	return &ValsetEthLogValsetUpdatedIterator{contract: _Valset.contract, event: "EthLogValsetUpdated", logs: logs, sub: sub}, nil
}

// WatchEthLogValsetUpdated is a free log subscription operation binding the contract event 0xa263e6a7a37f9afb3cb30027a71f0f16aa7936bab718860b78f2e799aaa0e249.
//
// Solidity: event EthLogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchEthLogValsetUpdated(opts *bind.WatchOpts, sink chan<- *ValsetEthLogValsetUpdated) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "EthLogValsetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetEthLogValsetUpdated)
				if err := _Valset.contract.UnpackLog(event, "EthLogValsetUpdated", log); err != nil {
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

// ParseEthLogValsetUpdated is a log parse operation binding the contract event 0xa263e6a7a37f9afb3cb30027a71f0f16aa7936bab718860b78f2e799aaa0e249.
//
// Solidity: event EthLogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseEthLogValsetUpdated(log types.Log) (*ValsetEthLogValsetUpdated, error) {
	event := new(ValsetEthLogValsetUpdated)
	if err := _Valset.contract.UnpackLog(event, "EthLogValsetUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
