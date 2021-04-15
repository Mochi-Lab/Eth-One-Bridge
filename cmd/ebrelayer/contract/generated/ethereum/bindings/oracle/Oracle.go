// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Oracle

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

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"EthLogNewOracleClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyPowerCurrent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyPowerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"EthLogUnlockProcessed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ORACLE_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeRegistry\",\"outputs\":[{\"internalType\":\"contractBridgeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"checkBridgeUnlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"harmonyBridge\",\"outputs\":[{\"internalType\":\"contractHarmonyBridge\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasMadeClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_consensusThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"newOracleClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracleClaimValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockID\",\"type\":\"uint256\"}],\"name\":\"processBridgeUnlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"contractValset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OracleBin is the compiled bytecode used for deploying new contracts.
var OracleBin = "0x60806040526000805534801561001457600080fd5b50611dee806100246000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80637f54af0c116100715780637f54af0c14610316578063972ded0914610360578063996f5144146103aa578063a219763e146103c8578063cd6dc6871461042e578063f9b0b5b91461047c576100b4565b8063316be171146100b957806336e413411461010357806349f7c7971461017b578063568b3c4f146101a9578063570ca735146102785780635ecab647146102c2575b600080fd5b6100c161049a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101396004803603604081101561011957600080fd5b8101908080359060200190929190803590602001909291905050506104c0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101a76004803603602081101561019157600080fd5b810190808035906020019092919050505061050b565b005b610276600480360360608110156101bf57600080fd5b810190808035906020019092919080359060200190929190803590602001906401000000008111156101f057600080fd5b82018360208201111561020257600080fd5b8035906020019184600183028401116401000000008311171561022457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610857565b005b610280610f8b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102ee600480360360208110156102d857600080fd5b8101908080359060200190929190505050610fb1565b6040518084151515158152602001838152602001828152602001935050505060405180910390f35b61031e6112c2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103686112e8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103b261130e565b6040518082815260200191505060405180910390f35b610414600480360360408110156103de57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611313565b604051808215151515815260200191505060405180910390f35b61047a6004803603604081101561044457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611342565b005b61048461177d565b6040518082815260200191505060405180910390f35b603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603660205281600052604060002081815481106104d957fe5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156105aa57600080fd5b505afa1580156105be573d6000803e3d6000fd5b505050506040513d60208110156105d457600080fd5b8101908080519060200190929190505050610657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b8060011515603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a86e377836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156106cf57600080fd5b505afa1580156106e3573d6000803e3d6000fd5b505050506040513d60208110156106f957600080fd5b8101908080519060200190929190505050151514610762576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d815260200180611d5f602d913960400191505060405180910390fd5b600080600061077085611783565b925092509250826107cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526048815260200180611c976048913960600191505060405180910390fd5b6107d585611ade565b7f49d4d330587ca24c6c2c18500b1e3d931dea6d8011a7fd0b3404f4efa3ce91bb85838333604051808581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390a15050505050565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156108f657600080fd5b505afa15801561090a573d6000803e3d6000fd5b505050506040513d602081101561092057600080fd5b81019080805190602001909291905050506109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d75737420626520616e206163746976652076616c696461746f72000000000081525060200191505060405180910390fd5b8260011515603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a86e377836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610a1b57600080fd5b505afa158015610a2f573d6000803e3d6000fd5b505050506040513d6020811015610a4557600080fd5b8101908080519060200190929190505050151514610aae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d815260200180611d5f602d913960400191505060405180910390fd5b6000339050603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319045a2585856040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b48578082015181840152602081019050610b2d565b50505050905090810190601f168015610b755780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b158015610b9357600080fd5b505afa158015610ba7573d6000803e3d6000fd5b505050506040513d6020811015610bbd57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c6e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f496e76616c6964206d657373616765207369676e61747572652e00000000000081525060200191505060405180910390fd5b6037600086815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610d22576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180611d04603a913960400191505060405180910390fd5b60016037600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550603660008681526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f2044e3f2f4a0e17625f1a884d866d6cad0cc709957cdac54c4218bd39af180c485858386604051808581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610ea5578082015181840152602081019050610e8a565b50505050905090810190601f168015610ed25780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a16000806000610ef088611783565b9250925092508215610f8157610f0588611ade565b7f49d4d330587ca24c6c2c18500b1e3d931dea6d8011a7fd0b3404f4efa3ce91bb88838333604051808581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390a15b5050505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611079576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b8360011515603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a86e377836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156110f157600080fd5b505afa158015611105573d6000803e3d6000fd5b505050506040513d602081101561111b57600080fd5b8101908080519060200190929190505050151514611184576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d815260200180611d5f602d913960400191505060405180910390fd5b60011515603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a86e377876040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156111fb57600080fd5b505afa15801561120f573d6000803e3d6000fd5b505050506040513d602081101561122557600080fd5b81019080805190602001909291905050501515146112ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f43616e206f6e6c7920636865636b206163746976652070726f7068656369657381525060200191505060405180910390fd5b6112b485611783565b935093509350509193909250565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600181565b60376020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600061134c611b6e565b9050600160009054906101000a900460ff168061136d575061136c611b77565b5b80611379575060005481115b6113ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611d8c602e913960400191505060405180910390fd5b6000600160009054906101000a900460ff16159050801561140b5760018060006101000a81548160ff021916908315150217905550816000819055505b60008311611464576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611cdf6025913960400191505060405180910390fd5b83603860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b15801561150d57600080fd5b505afa158015611521573d6000803e3d6000fd5b505050506040513d602081101561153757600080fd5b8101908080519060200190929190505050603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fd168ec6040518163ffffffff1660e01b815260040160206040518083038186803b1580156115f057600080fd5b505afa158015611604573d6000803e3d6000fd5b505050506040513d602081101561161a57600080fd5b8101908080519060200190929190505050603960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663476dc25c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156116d357600080fd5b505afa1580156116e7573d6000803e3d6000fd5b505050506040513d60208110156116fd57600080fd5b8101908080519060200190929190505050603a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826035819055508015611777576000600160006101000a81548160ff0219169083151502179055505b50505050565b60355481565b600080600080600090506000603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663db3ad22c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117f757600080fd5b505afa15801561180b573d6000803e3d6000fd5b505050506040513d602081101561182157600080fd5b8101908080519060200190929190505050905060008090505b6036600088815260200190815260200160002080549050811015611a8f57600060366000898152602001908152602001600020828154811061187857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340550a1c826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561194457600080fd5b505afa158015611958573d6000803e3d6000fd5b505050506040513d602081101561196e57600080fd5b810190808051906020019092919050505015611a7357611a70603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663473691a4836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611a2657600080fd5b505afa158015611a3a573d6000803e3d6000fd5b505050506040513d6020811015611a5057600080fd5b810190808051906020019092919050505085611b8890919063ffffffff16565b93505b50611a88600182611b8890919063ffffffff16565b905061183a565b506000611aa760355483611c1090919063ffffffff16565b90506000611abf606485611c1090919063ffffffff16565b9050600082821015905080828497509750975050505050509193909250565b603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639e67206d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015611b5357600080fd5b505af1158015611b67573d6000803e3d6000fd5b5050505050565b60006001905090565b600080303b90506000811491505090565b600080828401905083811015611c06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600080831415611c235760009050611c90565b6000828402905082848281611c3457fe5b0414611c8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611d3e6021913960400191505060405180910390fd5b809150505b9291505056fe5468652063756d756c617469766520706f776572206f66207369676e61746f72792076616c696461746f727320646f6573206e6f74206d65657420746865207468726573686f6c64436f6e73656e737573207468726573686f6c64206d75737420626520706f7369746976652e43616e6e6f74206d616b65206475706c6963617465206f7261636c6520636c61696d732066726f6d207468652073616d6520616464726573732e536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7754686520756e6c6f636b206d7573742062652070656e64696e6720666f722074686973206f7065726174696f6e436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a723158207ccaff9332db2ebb6d18580e34e546d7b7457bded5b900fd208f0e3d1dae89d464736f6c63430005110032"

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// ORACLEREVISION is a free data retrieval call binding the contract method 0x996f5144.
//
// Solidity: function ORACLE_REVISION() view returns(uint256)
func (_Oracle *OracleCaller) ORACLEREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "ORACLE_REVISION")
	return *ret0, err
}

// ORACLEREVISION is a free data retrieval call binding the contract method 0x996f5144.
//
// Solidity: function ORACLE_REVISION() view returns(uint256)
func (_Oracle *OracleSession) ORACLEREVISION() (*big.Int, error) {
	return _Oracle.Contract.ORACLEREVISION(&_Oracle.CallOpts)
}

// ORACLEREVISION is a free data retrieval call binding the contract method 0x996f5144.
//
// Solidity: function ORACLE_REVISION() view returns(uint256)
func (_Oracle *OracleCallerSession) ORACLEREVISION() (*big.Int, error) {
	return _Oracle.Contract.ORACLEREVISION(&_Oracle.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_Oracle *OracleCaller) BridgeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "bridgeRegistry")
	return *ret0, err
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_Oracle *OracleSession) BridgeRegistry() (common.Address, error) {
	return _Oracle.Contract.BridgeRegistry(&_Oracle.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_Oracle *OracleCallerSession) BridgeRegistry() (common.Address, error) {
	return _Oracle.Contract.BridgeRegistry(&_Oracle.CallOpts)
}

// CheckBridgeUnlock is a free data retrieval call binding the contract method 0x5ecab647.
//
// Solidity: function checkBridgeUnlock(uint256 _unlockID) view returns(bool, uint256, uint256)
func (_Oracle *OracleCaller) CheckBridgeUnlock(opts *bind.CallOpts, _unlockID *big.Int) (bool, *big.Int, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Oracle.contract.Call(opts, out, "checkBridgeUnlock", _unlockID)
	return *ret0, *ret1, *ret2, err
}

// CheckBridgeUnlock is a free data retrieval call binding the contract method 0x5ecab647.
//
// Solidity: function checkBridgeUnlock(uint256 _unlockID) view returns(bool, uint256, uint256)
func (_Oracle *OracleSession) CheckBridgeUnlock(_unlockID *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.CheckBridgeUnlock(&_Oracle.CallOpts, _unlockID)
}

// CheckBridgeUnlock is a free data retrieval call binding the contract method 0x5ecab647.
//
// Solidity: function checkBridgeUnlock(uint256 _unlockID) view returns(bool, uint256, uint256)
func (_Oracle *OracleCallerSession) CheckBridgeUnlock(_unlockID *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.CheckBridgeUnlock(&_Oracle.CallOpts, _unlockID)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_Oracle *OracleCaller) ConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "consensusThreshold")
	return *ret0, err
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_Oracle *OracleSession) ConsensusThreshold() (*big.Int, error) {
	return _Oracle.Contract.ConsensusThreshold(&_Oracle.CallOpts)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_Oracle *OracleCallerSession) ConsensusThreshold() (*big.Int, error) {
	return _Oracle.Contract.ConsensusThreshold(&_Oracle.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_Oracle *OracleCaller) HarmonyBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "harmonyBridge")
	return *ret0, err
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_Oracle *OracleSession) HarmonyBridge() (common.Address, error) {
	return _Oracle.Contract.HarmonyBridge(&_Oracle.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_Oracle *OracleCallerSession) HarmonyBridge() (common.Address, error) {
	return _Oracle.Contract.HarmonyBridge(&_Oracle.CallOpts)
}

// HasMadeClaim is a free data retrieval call binding the contract method 0xa219763e.
//
// Solidity: function hasMadeClaim(uint256 , address ) view returns(bool)
func (_Oracle *OracleCaller) HasMadeClaim(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "hasMadeClaim", arg0, arg1)
	return *ret0, err
}

// HasMadeClaim is a free data retrieval call binding the contract method 0xa219763e.
//
// Solidity: function hasMadeClaim(uint256 , address ) view returns(bool)
func (_Oracle *OracleSession) HasMadeClaim(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Oracle.Contract.HasMadeClaim(&_Oracle.CallOpts, arg0, arg1)
}

// HasMadeClaim is a free data retrieval call binding the contract method 0xa219763e.
//
// Solidity: function hasMadeClaim(uint256 , address ) view returns(bool)
func (_Oracle *OracleCallerSession) HasMadeClaim(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Oracle.Contract.HasMadeClaim(&_Oracle.CallOpts, arg0, arg1)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleSession) Operator() (common.Address, error) {
	return _Oracle.Contract.Operator(&_Oracle.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleCallerSession) Operator() (common.Address, error) {
	return _Oracle.Contract.Operator(&_Oracle.CallOpts)
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x36e41341.
//
// Solidity: function oracleClaimValidators(uint256 , uint256 ) view returns(address)
func (_Oracle *OracleCaller) OracleClaimValidators(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "oracleClaimValidators", arg0, arg1)
	return *ret0, err
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x36e41341.
//
// Solidity: function oracleClaimValidators(uint256 , uint256 ) view returns(address)
func (_Oracle *OracleSession) OracleClaimValidators(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Oracle.Contract.OracleClaimValidators(&_Oracle.CallOpts, arg0, arg1)
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x36e41341.
//
// Solidity: function oracleClaimValidators(uint256 , uint256 ) view returns(address)
func (_Oracle *OracleCallerSession) OracleClaimValidators(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Oracle.Contract.OracleClaimValidators(&_Oracle.CallOpts, arg0, arg1)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleSession) Valset() (common.Address, error) {
	return _Oracle.Contract.Valset(&_Oracle.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleCallerSession) Valset() (common.Address, error) {
	return _Oracle.Contract.Valset(&_Oracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _bridgeRegistry, uint256 _consensusThreshold) returns()
func (_Oracle *OracleTransactor) Initialize(opts *bind.TransactOpts, _bridgeRegistry common.Address, _consensusThreshold *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "initialize", _bridgeRegistry, _consensusThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _bridgeRegistry, uint256 _consensusThreshold) returns()
func (_Oracle *OracleSession) Initialize(_bridgeRegistry common.Address, _consensusThreshold *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, _bridgeRegistry, _consensusThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _bridgeRegistry, uint256 _consensusThreshold) returns()
func (_Oracle *OracleTransactorSession) Initialize(_bridgeRegistry common.Address, _consensusThreshold *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, _bridgeRegistry, _consensusThreshold)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x568b3c4f.
//
// Solidity: function newOracleClaim(uint256 _unlockID, bytes32 _message, bytes _signature) returns()
func (_Oracle *OracleTransactor) NewOracleClaim(opts *bind.TransactOpts, _unlockID *big.Int, _message [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "newOracleClaim", _unlockID, _message, _signature)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x568b3c4f.
//
// Solidity: function newOracleClaim(uint256 _unlockID, bytes32 _message, bytes _signature) returns()
func (_Oracle *OracleSession) NewOracleClaim(_unlockID *big.Int, _message [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleClaim(&_Oracle.TransactOpts, _unlockID, _message, _signature)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x568b3c4f.
//
// Solidity: function newOracleClaim(uint256 _unlockID, bytes32 _message, bytes _signature) returns()
func (_Oracle *OracleTransactorSession) NewOracleClaim(_unlockID *big.Int, _message [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleClaim(&_Oracle.TransactOpts, _unlockID, _message, _signature)
}

// ProcessBridgeUnlock is a paid mutator transaction binding the contract method 0x49f7c797.
//
// Solidity: function processBridgeUnlock(uint256 _unlockID) returns()
func (_Oracle *OracleTransactor) ProcessBridgeUnlock(opts *bind.TransactOpts, _unlockID *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "processBridgeUnlock", _unlockID)
}

// ProcessBridgeUnlock is a paid mutator transaction binding the contract method 0x49f7c797.
//
// Solidity: function processBridgeUnlock(uint256 _unlockID) returns()
func (_Oracle *OracleSession) ProcessBridgeUnlock(_unlockID *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ProcessBridgeUnlock(&_Oracle.TransactOpts, _unlockID)
}

// ProcessBridgeUnlock is a paid mutator transaction binding the contract method 0x49f7c797.
//
// Solidity: function processBridgeUnlock(uint256 _unlockID) returns()
func (_Oracle *OracleTransactorSession) ProcessBridgeUnlock(_unlockID *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ProcessBridgeUnlock(&_Oracle.TransactOpts, _unlockID)
}

// OracleEthLogNewOracleClaimIterator is returned from FilterEthLogNewOracleClaim and is used to iterate over the raw logs and unpacked data for EthLogNewOracleClaim events raised by the Oracle contract.
type OracleEthLogNewOracleClaimIterator struct {
	Event *OracleEthLogNewOracleClaim // Event containing the contract specifics and raw log

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
func (it *OracleEthLogNewOracleClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleEthLogNewOracleClaim)
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
		it.Event = new(OracleEthLogNewOracleClaim)
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
func (it *OracleEthLogNewOracleClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleEthLogNewOracleClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleEthLogNewOracleClaim represents a EthLogNewOracleClaim event raised by the Oracle contract.
type OracleEthLogNewOracleClaim struct {
	UnlockID         *big.Int
	Message          [32]byte
	ValidatorAddress common.Address
	Signature        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEthLogNewOracleClaim is a free log retrieval operation binding the contract event 0x2044e3f2f4a0e17625f1a884d866d6cad0cc709957cdac54c4218bd39af180c4.
//
// Solidity: event EthLogNewOracleClaim(uint256 _unlockID, bytes32 _message, address _validatorAddress, bytes _signature)
func (_Oracle *OracleFilterer) FilterEthLogNewOracleClaim(opts *bind.FilterOpts) (*OracleEthLogNewOracleClaimIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "EthLogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return &OracleEthLogNewOracleClaimIterator{contract: _Oracle.contract, event: "EthLogNewOracleClaim", logs: logs, sub: sub}, nil
}

// WatchEthLogNewOracleClaim is a free log subscription operation binding the contract event 0x2044e3f2f4a0e17625f1a884d866d6cad0cc709957cdac54c4218bd39af180c4.
//
// Solidity: event EthLogNewOracleClaim(uint256 _unlockID, bytes32 _message, address _validatorAddress, bytes _signature)
func (_Oracle *OracleFilterer) WatchEthLogNewOracleClaim(opts *bind.WatchOpts, sink chan<- *OracleEthLogNewOracleClaim) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "EthLogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleEthLogNewOracleClaim)
				if err := _Oracle.contract.UnpackLog(event, "EthLogNewOracleClaim", log); err != nil {
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

// ParseEthLogNewOracleClaim is a log parse operation binding the contract event 0x2044e3f2f4a0e17625f1a884d866d6cad0cc709957cdac54c4218bd39af180c4.
//
// Solidity: event EthLogNewOracleClaim(uint256 _unlockID, bytes32 _message, address _validatorAddress, bytes _signature)
func (_Oracle *OracleFilterer) ParseEthLogNewOracleClaim(log types.Log) (*OracleEthLogNewOracleClaim, error) {
	event := new(OracleEthLogNewOracleClaim)
	if err := _Oracle.contract.UnpackLog(event, "EthLogNewOracleClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OracleEthLogUnlockProcessedIterator is returned from FilterEthLogUnlockProcessed and is used to iterate over the raw logs and unpacked data for EthLogUnlockProcessed events raised by the Oracle contract.
type OracleEthLogUnlockProcessedIterator struct {
	Event *OracleEthLogUnlockProcessed // Event containing the contract specifics and raw log

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
func (it *OracleEthLogUnlockProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleEthLogUnlockProcessed)
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
		it.Event = new(OracleEthLogUnlockProcessed)
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
func (it *OracleEthLogUnlockProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleEthLogUnlockProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleEthLogUnlockProcessed represents a EthLogUnlockProcessed event raised by the Oracle contract.
type OracleEthLogUnlockProcessed struct {
	UnlockID               *big.Int
	ProphecyPowerCurrent   *big.Int
	ProphecyPowerThreshold *big.Int
	Submitter              common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterEthLogUnlockProcessed is a free log retrieval operation binding the contract event 0x49d4d330587ca24c6c2c18500b1e3d931dea6d8011a7fd0b3404f4efa3ce91bb.
//
// Solidity: event EthLogUnlockProcessed(uint256 _unlockID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_Oracle *OracleFilterer) FilterEthLogUnlockProcessed(opts *bind.FilterOpts) (*OracleEthLogUnlockProcessedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "EthLogUnlockProcessed")
	if err != nil {
		return nil, err
	}
	return &OracleEthLogUnlockProcessedIterator{contract: _Oracle.contract, event: "EthLogUnlockProcessed", logs: logs, sub: sub}, nil
}

// WatchEthLogUnlockProcessed is a free log subscription operation binding the contract event 0x49d4d330587ca24c6c2c18500b1e3d931dea6d8011a7fd0b3404f4efa3ce91bb.
//
// Solidity: event EthLogUnlockProcessed(uint256 _unlockID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_Oracle *OracleFilterer) WatchEthLogUnlockProcessed(opts *bind.WatchOpts, sink chan<- *OracleEthLogUnlockProcessed) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "EthLogUnlockProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleEthLogUnlockProcessed)
				if err := _Oracle.contract.UnpackLog(event, "EthLogUnlockProcessed", log); err != nil {
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

// ParseEthLogUnlockProcessed is a log parse operation binding the contract event 0x49d4d330587ca24c6c2c18500b1e3d931dea6d8011a7fd0b3404f4efa3ce91bb.
//
// Solidity: event EthLogUnlockProcessed(uint256 _unlockID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_Oracle *OracleFilterer) ParseEthLogUnlockProcessed(log types.Log) (*OracleEthLogUnlockProcessed, error) {
	event := new(OracleEthLogUnlockProcessed)
	if err := _Oracle.contract.UnpackLog(event, "EthLogUnlockProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}
