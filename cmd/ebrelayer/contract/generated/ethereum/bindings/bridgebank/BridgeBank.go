// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeBank

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

// BridgeBankABI is the input ABI used to generate the binding from.
const BridgeBankABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ethereumSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_harmonyTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"EthLogLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"EthLogUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeDenominator\",\"type\":\"uint256\"}],\"name\":\"EthUpdateFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newHarmonyBridge\",\"type\":\"address\"}],\"name\":\"EthUpdateHarmonyBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOracle\",\"type\":\"address\"}],\"name\":\"EthUpdateOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawETH\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"BRIDGEBANK_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETHAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ONEAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SAFE_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"activateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_harmonyToken\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bandOracleInterface\",\"outputs\":[{\"internalType\":\"contractBandOracleInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeRegistry\",\"outputs\":[{\"internalType\":\"contractBridgeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"checkUnlockable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"deactivateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"getLockedFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"getTokenMappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"getTotalERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"harmonyBridge\",\"outputs\":[{\"internalType\":\"contractHarmonyBridge\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bandOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lendingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_harmonyETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"isActiveToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lendingPool\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"swapETHForONE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"}],\"name\":\"swapETHForToken\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"swapETHForWrappedETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"}],\"name\":\"swapTokenForONE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"}],\"name\":\"swapTokenForToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"}],\"name\":\"swapTokenForWrappedETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"}],\"name\":\"swapToken_1_1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"harmonyMappedToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"}],\"name\":\"unlockERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"unlockETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeDenominator\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyBridge\",\"type\":\"address\"}],\"name\":\"updateHmyBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"updateOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethGateway\",\"outputs\":[{\"internalType\":\"contractIWETHGateway\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBankBin is the compiled bytecode used for deploying new contracts.
var BridgeBankBin = "0x6080604052600060015564e8d4a51000603755731111111111111111111111111111111111111111603b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550732222222222222222222222222222222222222222603c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016000806101000a81548160ff021916908315150217905550617e2f806100e56000396000f3fe60806040526004361061023b5760003560e01c8063972ded091161012e578063b5a9096e116100ab578063da228a9b1161006f578063da228a9b146108b4578063e86dea4a146108df578063ed581fd31461090a578063f126be2614610933578063fd930c7d1461095e5761023b565b8063b5a9096e146107f0578063ba1adba41461081b578063bb0a64db14610844578063c5e10eef1461086d578063cc2a9a5b146108985761023b565b8063a766a392116100f2578063a766a392146106e3578063ac4a875c1461070c578063ac66c6e014610749578063ad5c464814610788578063b4bfd9a7146107b35761023b565b8063972ded091461060e5780639b92910c146106395780639d05974b14610662578063a1c1d3e21461068d578063a59a9973146106b85761023b565b806349b1af93116101bc578063722b1d2311610180578063722b1d23146105435780637dc0d1d0146105805780637fabc432146105ab578063903e10ba146105c7578063964d042c146105e35761023b565b806349b1af931461046d578063570ca73514610496578063575c51c2146104c15780635e01297d146104fe57806368173bcf1461051a5761023b565b80632740c197116102035780632740c1971461039e578063316be171146103c75780633c298e78146103f257806344004cc11461041b5780634782f779146104445761023b565b80630d1ce2d2146102cd57806317877820146102f6578063180b0d7e1461031f5780631cb44dfc1461034a578063204897ff14610373575b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c290617984565b60405180910390fd5b005b3480156102d957600080fd5b506102f460048036036102ef91908101906166d0565b61099b565b005b34801561030257600080fd5b5061031d600480360361031891908101906166d0565b610b91565b005b34801561032b57600080fd5b50610334610c9c565b6040516103419190617a04565b60405180910390f35b34801561035657600080fd5b50610371600480360361036c91908101906166d0565b610ca2565b005b34801561037f57600080fd5b50610388610dad565b6040516103959190617a04565b60405180910390f35b3480156103aa57600080fd5b506103c560048036036103c09190810190616a59565b610db3565b005b3480156103d357600080fd5b506103dc610e8e565b6040516103e991906175b2565b60405180910390f35b3480156103fe57600080fd5b5061041960048036036104149190810190616722565b610eb4565b005b34801561042757600080fd5b50610442600480360361043d9190810190616836565b611415565b005b34801561045057600080fd5b5061046b60048036036104669190810190616771565b611861565b005b34801561047957600080fd5b50610494600480360361048f9190810190616836565b611b8c565b005b3480156104a257600080fd5b506104ab61220e565b6040516104b8919061733b565b60405180910390f35b3480156104cd57600080fd5b506104e860048036036104e391908101906166d0565b612234565b6040516104f5919061757c565b60405180910390f35b610518600480360361051391908101906168e8565b61228d565b005b34801561052657600080fd5b50610541600480360361053c91908101906166d0565b6127b5565b005b34801561054f57600080fd5b5061056a600480360361056591908101906166d0565b6129ab565b604051610577919061733b565b60405180910390f35b34801561058c57600080fd5b50610595612a17565b6040516105a2919061761e565b60405180910390f35b6105c560048036036105c091908101906168e8565b612a3d565b005b6105e160048036036105dc9190810190616924565b612ee3565b005b3480156105ef57600080fd5b506105f8613567565b6040516106059190617a04565b60405180910390f35b34801561061a57600080fd5b5061062361371b565b60405161063091906175cd565b60405180910390f35b34801561064557600080fd5b50610660600480360361065b9190810190616885565b613741565b005b34801561066e57600080fd5b50610677613eb7565b604051610684919061733b565b60405180910390f35b34801561069957600080fd5b506106a2613edd565b6040516106af9190617597565b60405180910390f35b3480156106c457600080fd5b506106cd613f03565b6040516106da91906175e8565b60405180910390f35b3480156106ef57600080fd5b5061070a60048036036107059190810190616924565b613f29565b005b34801561071857600080fd5b50610733600480360361072e91908101906166d0565b6142e1565b6040516107409190617a04565b60405180910390f35b34801561075557600080fd5b50610770600480360361076b91908101906166d0565b614585565b60405161077f93929190617a48565b60405180910390f35b34801561079457600080fd5b5061079d6145dc565b6040516107aa919061733b565b60405180910390f35b3480156107bf57600080fd5b506107da60048036036107d591908101906168e8565b614602565b6040516107e7919061757c565b60405180910390f35b3480156107fc57600080fd5b50610805614681565b6040516108129190617a04565b60405180910390f35b34801561082757600080fd5b50610842600480360361083d9190810190616836565b614687565b005b34801561085057600080fd5b5061086b60048036036108669190810190616836565b614ca1565b005b34801561087957600080fd5b50610882615226565b60405161088f9190617603565b60405180910390f35b6108b260048036036108ad91908101906167ad565b61524c565b005b3480156108c057600080fd5b506108c9615937565b6040516108d6919061733b565b60405180910390f35b3480156108eb57600080fd5b506108f461595d565b6040516109019190617a04565b60405180910390f35b34801561091657600080fd5b50610931600480360361092c9190810190616771565b615963565b005b34801561093f57600080fd5b50610948615cd3565b6040516109559190617a04565b60405180910390f35b34801561096a57600080fd5b50610985600480360361098091908101906166d0565b615cd8565b6040516109929190617a04565b60405180910390f35b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2290617984565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a93906177bc565b60405180910390fd5b60001515603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16151514610b32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b29906178e4565b60405180910390fd5b6001603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160146101000a81548160ff0219169083151502179055505050565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1890617984565b60405180910390fd5b80604060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2ee0580b26ac9d813eb02bff2dc08fbf08346dafcc883f3c21c44d014b8a2e8581604051610c91919061733b565b60405180910390a150565b60365481565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2990617984565b60405180910390fd5b80603f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe0dec806007e306f472d1bb9e8dcf714fbb3da18e34206ea34f2d2c4bba8965981604051610da2919061733b565b60405180910390a150565b60375481565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3a90617984565b60405180910390fd5b81603581905550806036819055507f87018147d45d73a403e53a1dae7a7e95ad1bf96326b915560c0fad22373a35218282604051610e82929190617a7f565b60405180910390a15050565b603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900460ff16610f02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef9906179c4565b60405180910390fd5b60008060006101000a81548160ff021916908315150217905550604060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa3906177dc565b60405180910390fd5b8060008111610ff0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe7906179e4565b60405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611061576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105890617964565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16603d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112a90617924565b60405180910390fd5b61113c846142e1565b83111561117e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611175906177fc565b60405180910390fd5b60008473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016111b99190617356565b60206040518083038186803b1580156111d157600080fd5b505afa1580156111e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506112099190810190616a30565b90508084116112a5578473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff1660e01b815260040161124d9291906173d1565b602060405180830381600087803b15801561126757600080fd5b505af115801561127b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061129f9190810190616973565b506113e8565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166369328dec86838703896040518463ffffffff1660e01b815260040161130693929190617500565b602060405180830381600087803b15801561132057600080fd5b505af1158015611334573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506113589190810190616a30565b508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87836040518363ffffffff1660e01b81526004016113949291906173d1565b602060405180830381600087803b1580156113ae57600080fd5b505af11580156113c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506113e69190810190616973565b505b6113f3858786615d24565b50505060016000806101000a81548160ff021916908315150217905550505050565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149c90617984565b60405180910390fd5b6000809054906101000a900460ff166114f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ea906179c4565b60405180910390fd5b60008060006101000a81548160ff021916908315150217905550603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154611559846142e1565b0381111561159c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611593906177fc565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016115d79190617356565b60206040518083038186803b1580156115ef57600080fd5b505afa158015611603573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506116279190810190616a30565b90508082116116c3578373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b815260040161166b9291906174a0565b602060405180830381600087803b15801561168557600080fd5b505af1158015611699573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506116bd9190810190616973565b50611806565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166369328dec85838503866040518463ffffffff1660e01b8152600401611724939291906174c9565b602060405180830381600087803b15801561173e57600080fd5b505af1158015611752573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506117769190810190616a30565b508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff1660e01b81526004016117b29291906174a0565b602060405180830381600087803b1580156117cc57600080fd5b505af11580156117e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506118049190810190616973565b505b7f378ae7e321aef491da8299462610eb48bb07e3b93619adfb2cdbe332b6aad7ce84848460405161183993929190617469565b60405180910390a15060016000806101000a81548160ff021916908315150217905550505050565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146118f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e890617984565b60405180910390fd5b6000809054906101000a900460ff1661193f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611936906179c4565b60405180910390fd5b60008060006101000a81548160ff021916908315150217905550603d6000603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001546119c6613567565b03811115611a09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a0090617904565b60405180910390fd5b478111611a5c578173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611a56573d6000803e3d6000fd5b50611b35565b604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166336118b52478303846040518363ffffffff1660e01b8152600401611abb929190617a1f565b600060405180830381600087803b158015611ad557600080fd5b505af1158015611ae9573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015611b33573d6000803e3d6000fd5b505b7fc7d71dfdba3c33b60061c39022c106fd785ed1edda60862045a0b684ca1f0e5d8282604051611b669291906173d1565b60405180910390a160016000806101000a81548160ff0219169083151502179055505050565b60385460016038540111611bd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bcc9061775c565b60405180910390fd5b6000809054906101000a900460ff16611c23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c1a906179c4565b60405180910390fd5b60008060006101000a81548160ff02191690831515021790555081603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16611ccd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cc4906179a4565b60405180910390fd5b8160008111611d11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d08906179e4565b60405180910390fd5b84600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611d82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d7990617964565b60405180910390fd5b6000611d8d85615e96565b90508573ffffffffffffffffffffffffffffffffffffffff166323b872dd33308489016040518463ffffffff1660e01b8152600401611dce93929190617371565b602060405180830381600087803b158015611de857600080fd5b505af1158015611dfc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611e209190810190616973565b611e5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e569061783c565b60405180910390fd5b611e676162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a75886040518263ffffffff1660e01b8152600401611ec2919061733b565b6101806040518083038186803b158015611edb57600080fd5b505afa158015611eef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611f139190810190616a06565b9050600073ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614611fe457604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8eda9df888489013060006040518563ffffffff1660e01b8152600401611fb19493929190617537565b600060405180830381600087803b158015611fcb57600080fd5b505af1158015611fdf573d6000803e3d6000fd5b505050505b611fec6163d5565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365555bcc8973ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561207057600080fd5b505afa158015612084573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506120ad919081019061699c565b6040518263ffffffff1660e01b81526004016120c991906176c7565b60606040518083038186803b1580156120e157600080fd5b505afa1580156120f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061211991908101906169dd565b9050600061214e670de0b6b3a764000061214084600001518b615f1190919063ffffffff16565b615f8190919063ffffffff16565b90506000603d6000603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506121e7338c8c848d87615fcb565b505050505050505060016000806101000a81548160ff021916908315150217905550505050565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000603d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff169050919050565b6000809054906101000a900460ff166122db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122d2906179c4565b60405180910390fd5b60008060006101000a81548160ff0219169083151502179055506038546001603854011161233e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123359061775c565b60405180910390fd5b603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff166123f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123e7906179a4565b60405180910390fd5b8160008111612434576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161242b906179e4565b60405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156124a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161249c90617964565b60405180910390fd5b60006124b085615e96565b905080850134146124f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124ed9061777c565b60405180910390fd5b6000604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16343060006040516024016125469291906173a8565b6040516020818303038152906040527f58c22be7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516125d09190617324565b60006040518083038185875af1925050503d806000811461260d576040519150601f19603f3d011682016040523d82523d6000602084013e612612565b606091505b5050905080612656576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161264d9061771c565b60405180910390fd5b61265e6163d5565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365555bcc6040518163ffffffff1660e01b81526004016126b790617891565b60606040518083038186803b1580156126cf57600080fd5b505afa1580156126e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061270791908101906169dd565b9050600061273c670de0b6b3a764000061272e84600001518b615f1190919063ffffffff16565b615f8190919063ffffffff16565b9050612790338a603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16603c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c86615fcb565b5050505050505060016000806101000a81548160ff0219169083151502179055505050565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161283c90617984565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156128b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128ad906177bc565b60405180910390fd5b60011515603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff1615151461294c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612943906176fc565b60405180910390fd5b6000603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160146101000a81548160ff0219169083151502179055505050565b6000603d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b603f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900460ff16612a8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a82906179c4565b60405180910390fd5b60008060006101000a81548160ff02191690831515021790555060385460016038540111612aee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ae59061775c565b60405180910390fd5b603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16612ba0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b97906179a4565b60405180910390fd5b8160008111612be4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bdb906179e4565b60405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612c55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c4c90617964565b60405180910390fd5b6000612c6085615e96565b90508085013414612ca6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c9d9061777c565b60405180910390fd5b6000604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1634306000604051602401612cf69291906173a8565b6040516020818303038152906040527f58c22be7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612d809190617324565b60006040518083038185875af1925050503d8060008114612dbd576040519150601f19603f3d011682016040523d82523d6000602084013e612dc2565b606091505b5050905080612e06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dfd9061771c565b60405180910390fd5b6000603d6000603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050612ebf3389603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848b8c615fcb565b50505050505060016000806101000a81548160ff0219169083151502179055505050565b6000809054906101000a900460ff16612f31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f28906179c4565b60405180910390fd5b60008060006101000a81548160ff02191690831515021790555060385460016038540111612f94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f8b9061775c565b60405180910390fd5b603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16613046576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161303d906179a4565b60405180910390fd5b81603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff166130d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130cd906179a4565b60405180910390fd5b836000811161311a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613111906179e4565b60405180910390fd5b85600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561318b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161318290617964565b60405180910390fd5b600061319687615e96565b905080870134146131dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131d39061777c565b60405180910390fd5b6000604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163430600060405160240161322c9291906173a8565b6040516020818303038152906040527f58c22be7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516132b69190617324565b60006040518083038185875af1925050503d80600081146132f3576040519150601f19603f3d011682016040523d82523d6000602084013e6132f8565b606091505b505090508061333c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133339061771c565b60405180910390fd5b6133446163d5565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365555bcc8973ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156133c857600080fd5b505afa1580156133dc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250613405919081019061699c565b6040518263ffffffff1660e01b8152600401613421919061785c565b60606040518083038186803b15801561343957600080fd5b505afa15801561344d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061347191908101906169dd565b905060006134a6670de0b6b3a764000061349884600001518d615f1190919063ffffffff16565b615f8190919063ffffffff16565b90506000603d60008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061353f338d603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848f87615fcb565b50505050505050505060016000806101000a81548160ff021916908315150217905550505050565b60006135716162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a75603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016135ee919061733b565b6101806040518083038186803b15801561360757600080fd5b505afa15801561361b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061363f9190810190616a06565b905060008160e001519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613712578073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016136b79190617356565b60206040518083038186803b1580156136cf57600080fd5b505afa1580156136e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506137079190810190616a30565b470192505050613718565b47925050505b90565b604060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6038546001603854011161378a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137819061775c565b60405180910390fd5b6000809054906101000a900460ff166137d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137cf906179c4565b60405180910390fd5b60008060006101000a81548160ff02191690831515021790555082603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16613882576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613879906179a4565b60405180910390fd5b81603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16613912576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613909906179a4565b60405180910390fd5b8360008111613956576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161394d906179e4565b60405180910390fd5b86600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156139c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139be90617964565b60405180910390fd5b60006139d287615e96565b90508773ffffffffffffffffffffffffffffffffffffffff166323b872dd3330848b016040518463ffffffff1660e01b8152600401613a1393929190617371565b602060405180830381600087803b158015613a2d57600080fd5b505af1158015613a41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613a659190810190616973565b613aa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a9b9061783c565b60405180910390fd5b613aac6162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a758a6040518263ffffffff1660e01b8152600401613b07919061733b565b6101806040518083038186803b158015613b2057600080fd5b505afa158015613b34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613b589190810190616a06565b9050600073ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614613c2957604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8eda9df8a848b013060006040518563ffffffff1660e01b8152600401613bf69493929190617537565b600060405180830381600087803b158015613c1057600080fd5b505af1158015613c24573d6000803e3d6000fd5b505050505b613c316163d5565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365555bcc8b73ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015613cb557600080fd5b505afa158015613cc9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250613cf2919081019061699c565b8a73ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015613d3857600080fd5b505afa158015613d4c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250613d75919081019061699c565b6040518363ffffffff1660e01b8152600401613d9292919061765b565b60606040518083038186803b158015613daa57600080fd5b505afa158015613dbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613de291908101906169dd565b90506000613e17670de0b6b3a7640000613e0984600001518d615f1190919063ffffffff16565b615f8190919063ffffffff16565b90506000603d60008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050613e8e338e8e848f87615fcb565b50505050505050505060016000806101000a81548160ff02191690831515021790555050505050565b603c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613fb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613fb090617984565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156140235750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b614062576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614059906177bc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16603d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161412a90617944565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161417093929190617371565b602060405180830381600087803b15801561418a57600080fd5b505af115801561419e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506141c29190810190616973565b5061420c603d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083836160cf565b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b3604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b81526004016142899291906174a0565b602060405180830381600087803b1580156142a357600080fd5b505af11580156142b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506142db9190810190616973565b50505050565b60006142eb6162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a75846040518263ffffffff1660e01b8152600401614346919061733b565b6101806040518083038186803b15801561435f57600080fd5b505afa158015614373573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506143979190810190616a06565b905060008160e001519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146144f2578073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161440f9190617356565b60206040518083038186803b15801561442757600080fd5b505afa15801561443b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061445f9190810190616a30565b8473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016144989190617356565b60206040518083038186803b1580156144b057600080fd5b505afa1580156144c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506144e89190810190616a30565b0192505050614580565b8373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161452b9190617356565b60206040518083038186803b15801561454357600080fd5b505afa158015614557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061457b9190810190616a30565b925050505b919050565b603d6020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16905083565b603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561466c5781614663613567565b1015905061467b565b81614676846142e1565b101590505b92915050565b60385481565b603854600160385401116146d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146c79061775c565b60405180910390fd5b6000809054906101000a900460ff1661471e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614715906179c4565b60405180910390fd5b60008060006101000a81548160ff02191690831515021790555081603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff166147c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016147bf906179a4565b60405180910390fd5b816000811161480c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614803906179e4565b60405180910390fd5b84600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561487d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161487490617964565b60405180910390fd5b600061488885615e96565b90508573ffffffffffffffffffffffffffffffffffffffff166323b872dd33308489016040518463ffffffff1660e01b81526004016148c993929190617371565b602060405180830381600087803b1580156148e357600080fd5b505af11580156148f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061491b9190810190616973565b61495a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149519061783c565b60405180910390fd5b6149626162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a75886040518263ffffffff1660e01b81526004016149bd919061733b565b6101806040518083038186803b1580156149d657600080fd5b505afa1580156149ea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614a0e9190810190616a06565b9050600073ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff1614614adf57604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8eda9df888489013060006040518563ffffffff1660e01b8152600401614aac9493929190617537565b600060405180830381600087803b158015614ac657600080fd5b505af1158015614ada573d6000803e3d6000fd5b505050505b614ae76163d5565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365555bcc8973ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015614b6b57600080fd5b505afa158015614b7f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250614ba8919081019061699c565b6040518263ffffffff1660e01b8152600401614bc49190617692565b60606040518083038186803b158015614bdc57600080fd5b505afa158015614bf0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614c1491908101906169dd565b90506000614c49670de0b6b3a7640000614c3b84600001518b615f1190919063ffffffff16565b615f8190919063ffffffff16565b9050614c7b338b8b603c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c86615fcb565b5050505050505060016000806101000a81548160ff021916908315150217905550505050565b6000809054906101000a900460ff16614cef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ce6906179c4565b60405180910390fd5b60008060006101000a81548160ff02191690831515021790555060385460016038540111614d52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d499061775c565b60405180910390fd5b81603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614de4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ddb906178c4565b60405180910390fd5b82603d60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900460ff16614e74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e6b906179a4565b60405180910390fd5b8260008111614eb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614eaf906179e4565b60405180910390fd5b85600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614f29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614f2090617964565b60405180910390fd5b6000614f3486615e96565b90508673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330848a016040518463ffffffff1660e01b8152600401614f7593929190617371565b602060405180830381600087803b158015614f8f57600080fd5b505af1158015614fa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614fc79190810190616973565b615006576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ffd9061783c565b60405180910390fd5b61500e6162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a75896040518263ffffffff1660e01b8152600401615069919061733b565b6101806040518083038186803b15801561508257600080fd5b505afa158015615096573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506150ba9190810190616a06565b9050600073ffffffffffffffffffffffffffffffffffffffff168160e0015173ffffffffffffffffffffffffffffffffffffffff161461518b57604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8eda9df89848a013060006040518563ffffffff1660e01b81526004016151589493929190617537565b600060405180830381600087803b15801561517257600080fd5b505af1158015615186573d6000803e3d6000fd5b505050505b6000603d60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050615200338b8b848c8d615fcb565b5050505050505060016000806101000a81548160ff021916908315150217905550505050565b604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061525661613d565b9050600260009054906101000a900460ff16806152775750615276616146565b5b80615283575060015481115b6152c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152b99061781c565b60405180910390fd5b6000600260009054906101000a900460ff161590508015615300576001600260006101000a81548160ff021916908315150217905550816001819055505b87603e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b1580156153a957600080fd5b505afa1580156153bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506153e191908101906166f9565b603960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663833b1fce6040518163ffffffff1660e01b815260040160206040518083038186803b15801561548957600080fd5b505afa15801561549d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506154c191908101906166f9565b603f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663476dc25c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561556957600080fd5b505afa15801561557d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506155a191908101906166f9565b604060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555086604160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085604260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084604360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083603a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000603881905550615758603d6000603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002034856160cf565b6157606162b2565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335ea6a75603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016157dd919061733b565b6101806040518083038186803b1580156157f657600080fd5b505afa15801561580a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061582e9190810190616a06565b905060008160e0015190508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b81526004016158b69291906174a0565b602060405180830381600087803b1580156158d057600080fd5b505af11580156158e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506159089190810190616973565b505050801561592d576000600260006101000a81548160ff0219169083151502179055505b5050505050505050565b603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60355481565b604060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146159f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016159ea906177dc565b60405180910390fd5b6000809054906101000a900460ff16615a41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615a38906179c4565b60405180910390fd5b60008060006101000a81548160ff0219169083151502179055508060008111615a9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615a96906179e4565b60405180910390fd5b82600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615b10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b0790617964565b60405180910390fd5b615b18613567565b831115615b5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b5190617904565b60405180910390fd5b478311615bad578373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015615ba7573d6000803e3d6000fd5b50615c86565b604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166336118b52478503866040518363ffffffff1660e01b8152600401615c0c929190617a1f565b600060405180830381600087803b158015615c2657600080fd5b505af1158015615c3a573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015615c84573d6000803e3d6000fd5b505b615cb3603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168585615d24565b505060016000806101000a81548160ff0219169083151502179055505050565b600181565b6000603d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b80603d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410615e0d57615dc281603d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015461615790919063ffffffff16565b603d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550615e56565b6000603d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055505b7f4445dabfef575721ee82c521f0725ef2cab18153d75cadda8115c2e31efc5508828483604051615e8993929190617469565b60405180910390a1505050565b600080600060355414158015615eaf5750600060365414155b15615f0857615f05603754615ef7603654615ee9603754615edb6035548a615f1190919063ffffffff16565b615f1190919063ffffffff16565b615f8190919063ffffffff16565b615f8190919063ffffffff16565b90505b80915050919050565b600080831415615f245760009050615f7b565b6000828402905082848281615f3557fe5b0414615f76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f6d9061779c565b60405180910390fd5b809150505b92915050565b6000615fc383836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506161a1565b905092915050565b615fe1600160385461620290919063ffffffff16565b60388190555061603c82603d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015461620290919063ffffffff16565b603d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055507f15bb5885b6550f0baf0cf174b9618b2d0c83f6ab83ea0de29b03cb6c9461afd58686868686866038546040516160bf97969594939291906173fa565b60405180910390a1505050505050565b818360000181905550808360010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018360010160146101000a81548160ff021916908315150217905550505050565b60006001905090565b600080303b90506000811491505090565b600061619983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250616257565b905092915050565b600080831182906161e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016161df9190617639565b60405180910390fd5b5060008385816161f457fe5b049050809150509392505050565b60008082840190508381101561624d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016162449061773c565b60405180910390fd5b8091505092915050565b600083831115829061629f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016162969190617639565b60405180910390fd5b5060008385039050809150509392505050565b6040518061018001604052806162c66163f6565b815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff168152602001600064ffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff1681525090565b60405180606001604052806000815260200160008152602001600081525090565b6040518060200160405280600081525090565b60008135905061641881617d4b565b92915050565b60008151905061642d81617d4b565b92915050565b60008135905061644281617d62565b92915050565b60008151905061645781617d79565b92915050565b600082601f83011261646e57600080fd5b815161648161647c82617ad5565b617aa8565b9150808252602083016020830185838301111561649d57600080fd5b6164a8838284617d07565b50505092915050565b6000606082840312156164c357600080fd5b6164cd6060617aa8565b905060006164dd84828501616691565b60008301525060206164f184828501616691565b602083015250604061650584828501616691565b60408301525092915050565b60006020828403121561652357600080fd5b61652d6020617aa8565b9050600061653d84828501616691565b60008301525092915050565b6000610180828403121561655c57600080fd5b616567610180617aa8565b9050600061657784828501616511565b600083015250602061658b84828501616667565b602083015250604061659f84828501616667565b60408301525060606165b384828501616667565b60608301525060806165c784828501616667565b60808301525060a06165db84828501616667565b60a08301525060c06165ef848285016166a6565b60c08301525060e06166038482850161641e565b60e0830152506101006166188482850161641e565b6101008301525061012061662e8482850161641e565b610120830152506101406166448482850161641e565b6101408301525061016061665a848285016166bb565b6101608301525092915050565b60008151905061667681617d90565b92915050565b60008135905061668b81617da7565b92915050565b6000815190506166a081617da7565b92915050565b6000815190506166b581617dbe565b92915050565b6000815190506166ca81617dd5565b92915050565b6000602082840312156166e257600080fd5b60006166f084828501616409565b91505092915050565b60006020828403121561670b57600080fd5b60006167198482850161641e565b91505092915050565b60008060006060848603121561673757600080fd5b600061674586828701616433565b935050602061675686828701616409565b92505060406167678682870161667c565b9150509250925092565b6000806040838503121561678457600080fd5b600061679285828601616433565b92505060206167a38582860161667c565b9150509250929050565b60008060008060008060c087890312156167c657600080fd5b60006167d489828a01616409565b96505060206167e589828a01616409565b95505060406167f689828a01616409565b945050606061680789828a01616409565b935050608061681889828a01616409565b92505060a061682989828a01616409565b9150509295509295509295565b60008060006060848603121561684b57600080fd5b600061685986828701616409565b935050602061686a86828701616409565b925050604061687b8682870161667c565b9150509250925092565b6000806000806080858703121561689b57600080fd5b60006168a987828801616409565b94505060206168ba87828801616409565b93505060406168cb8782880161667c565b92505060606168dc87828801616409565b91505092959194509250565b600080604083850312156168fb57600080fd5b600061690985828601616409565b925050602061691a8582860161667c565b9150509250929050565b60008060006060848603121561693957600080fd5b600061694786828701616409565b93505060206169588682870161667c565b925050604061696986828701616409565b9150509250925092565b60006020828403121561698557600080fd5b600061699384828501616448565b91505092915050565b6000602082840312156169ae57600080fd5b600082015167ffffffffffffffff8111156169c857600080fd5b6169d48482850161645d565b91505092915050565b6000606082840312156169ef57600080fd5b60006169fd848285016164b1565b91505092915050565b60006101808284031215616a1957600080fd5b6000616a2784828501616549565b91505092915050565b600060208284031215616a4257600080fd5b6000616a5084828501616691565b91505092915050565b60008060408385031215616a6c57600080fd5b6000616a7a8582860161667c565b9250506020616a8b8582860161667c565b9150509250929050565b616a9e81617bd5565b82525050565b616aad81617b45565b82525050565b616abc81617b33565b82525050565b616acb81617b57565b82525050565b6000616adc82617b01565b616ae68185617b17565b9350616af6818560208601617d07565b80840191505092915050565b616b0b81617be7565b82525050565b616b1a81617c0b565b82525050565b616b2981617c2f565b82525050565b616b3881617c53565b82525050565b616b4781617c77565b82525050565b616b5681617c9b565b82525050565b616b6581617cbf565b82525050565b616b7481617cd1565b82525050565b6000616b8582617b0c565b616b8f8185617b22565b9350616b9f818560208601617d07565b616ba881617d3a565b840191505092915050565b6000616bc0601a83617b22565b91507f546f6b656e20616c7265616479206465616374697661746564210000000000006000830152602082019050919050565b6000616c00602583617b22565b91507f41617665204c656e64696e67506f6f6c3a206465706f7369742045544820666160008301527f696c6564210000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616c66601b83617b22565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000616ca6601483617b22565b91507f4e6f20617661696c61626c65206e6f6e6365732e0000000000000000000000006000830152602082019050919050565b6000616ce6604283617b22565b91507f546865207472616e73616374696f6e732076616c7565206d757374206265206560008301527f7175616c207468652073706563696669656420616d6f756e742028696e20776560208301527f69290000000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000616d72600383617b22565b91507f4f4e4500000000000000000000000000000000000000000000000000000000006000830152602082019050919050565b6000616db2602183617b22565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616e18601b83617b22565b91507f546f6b656e2061646472657373206d7573742062652076616c696400000000006000830152602082019050919050565b6000616e58602783617b22565b91507f416363657373207265737472696374656420746f20746865206861726d6f6e7960008301527f20627269646765000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616ebe602c83617b22565b91507f457863656564656420616d6f756e74206f6620546f6b656e20616c6c6f77656460008301527f20746f20776974686472617700000000000000000000000000000000000000006020830152604082019050919050565b6000616f24602e83617b22565b91507f436f6e747261637420696e7374616e63652068617320616c726561647920626560008301527f656e20696e697469616c697a65640000000000000000000000000000000000006020830152604082019050919050565b6000616f8a604483617b22565b91507f436f6e747261637420746f6b656e20616c6c6f77616e63657320696e7375666660008301527f696369656e7420746f20636f6d706c6574652074686973206c6f636b2072657160208301527f75657374000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000617016600383617b22565b91507f45544800000000000000000000000000000000000000000000000000000000006000830152602082019050919050565b6000617056601583617b22565b91507f546f6b656e206d757374206e6f742062652045544800000000000000000000006000830152602082019050919050565b6000617096601883617b22565b91507f546f6b656e20616c7265616479206163746976617465642100000000000000006000830152602082019050919050565b60006170d6602a83617b22565b91507f457863656564656420616d6f756e74206f662045544820616c6c6f776564207460008301527f6f207769746864726177000000000000000000000000000000000000000000006020830152604082019050919050565b600061713c601583617b22565b91507f496e76616c696420746f6b656e206164647265737300000000000000000000006000830152602082019050919050565b600061717c601483617b22565b91507f546f6b656e20616c7265616479206164646564210000000000000000000000006000830152602082019050919050565b60006171bc601683617b22565b91507f5265636569766572206d7573742062652076616c6964000000000000000000006000830152602082019050919050565b60006171fc601c83617b22565b91507f4d7573742062652042726964676542616e6b206f70657261746f722e000000006000830152602082019050919050565b600061723c601383617b22565b91507f546f6b656e206973206e6f7420616374697665000000000000000000000000006000830152602082019050919050565b600061727c601f83617b22565b91507f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006000830152602082019050919050565b60006172bc602683617b22565b91507f416d6f756e7420746f6b656e206d75737420626520677265617465722074686160008301527f6e207a65726f00000000000000000000000000000000000000000000000000006020830152604082019050919050565b61731e81617bad565b82525050565b60006173308284616ad1565b915081905092915050565b60006020820190506173506000830184616ab3565b92915050565b600060208201905061736b6000830184616a95565b92915050565b60006060820190506173866000830186616a95565b6173936020830185616a95565b6173a06040830184617315565b949350505050565b60006040820190506173bd6000830185616aa4565b6173ca6020830184616b6b565b9392505050565b60006040820190506173e66000830185616a95565b6173f36020830184617315565b9392505050565b600060e08201905061740f600083018a616ab3565b61741c6020830189616ab3565b6174296040830188616ab3565b6174366060830187616ab3565b6174436080830186617315565b61745060a0830185617315565b61745d60c0830184617315565b98975050505050505050565b600060608201905061747e6000830186616ab3565b61748b6020830185616ab3565b6174986040830184617315565b949350505050565b60006040820190506174b56000830185616ab3565b6174c26020830184617315565b9392505050565b60006060820190506174de6000830186616ab3565b6174eb6020830185617315565b6174f86040830184616ab3565b949350505050565b60006060820190506175156000830186616ab3565b6175226020830185617315565b61752f6040830184616a95565b949350505050565b600060808201905061754c6000830187616ab3565b6175596020830186617315565b6175666040830185616a95565b6175736060830184616b5c565b95945050505050565b60006020820190506175916000830184616ac2565b92915050565b60006020820190506175ac6000830184616b02565b92915050565b60006020820190506175c76000830184616b11565b92915050565b60006020820190506175e26000830184616b20565b92915050565b60006020820190506175fd6000830184616b2f565b92915050565b60006020820190506176186000830184616b3e565b92915050565b60006020820190506176336000830184616b4d565b92915050565b600060208201905081810360008301526176538184616b7a565b905092915050565b600060408201905081810360008301526176758185616b7a565b905081810360208301526176898184616b7a565b90509392505050565b600060408201905081810360008301526176ac8184616b7a565b905081810360208301526176bf81616d65565b905092915050565b600060408201905081810360008301526176e18184616b7a565b905081810360208301526176f481617009565b905092915050565b6000602082019050818103600083015261771581616bb3565b9050919050565b6000602082019050818103600083015261773581616bf3565b9050919050565b6000602082019050818103600083015261775581616c59565b9050919050565b6000602082019050818103600083015261777581616c99565b9050919050565b6000602082019050818103600083015261779581616cd9565b9050919050565b600060208201905081810360008301526177b581616da5565b9050919050565b600060208201905081810360008301526177d581616e0b565b9050919050565b600060208201905081810360008301526177f581616e4b565b9050919050565b6000602082019050818103600083015261781581616eb1565b9050919050565b6000602082019050818103600083015261783581616f17565b9050919050565b6000602082019050818103600083015261785581616f7d565b9050919050565b6000604082019050818103600083015261787581617009565b905081810360208301526178898184616b7a565b905092915050565b600060408201905081810360008301526178aa81617009565b905081810360208301526178bd81616d65565b9050919050565b600060208201905081810360008301526178dd81617049565b9050919050565b600060208201905081810360008301526178fd81617089565b9050919050565b6000602082019050818103600083015261791d816170c9565b9050919050565b6000602082019050818103600083015261793d8161712f565b9050919050565b6000602082019050818103600083015261795d8161716f565b9050919050565b6000602082019050818103600083015261797d816171af565b9050919050565b6000602082019050818103600083015261799d816171ef565b9050919050565b600060208201905081810360008301526179bd8161722f565b9050919050565b600060208201905081810360008301526179dd8161726f565b9050919050565b600060208201905081810360008301526179fd816172af565b9050919050565b6000602082019050617a196000830184617315565b92915050565b6000604082019050617a346000830185617315565b617a416020830184616a95565b9392505050565b6000606082019050617a5d6000830186617315565b617a6a6020830185616ab3565b617a776040830184616ac2565b949350505050565b6000604082019050617a946000830185617315565b617aa16020830184617315565b9392505050565b6000604051905081810181811067ffffffffffffffff82111715617acb57600080fd5b8060405250919050565b600067ffffffffffffffff821115617aec57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000617b3e82617b8d565b9050919050565b6000617b5082617b8d565b9050919050565b60008115159050919050565b60006fffffffffffffffffffffffffffffffff82169050919050565b600061ffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600064ffffffffff82169050919050565b600060ff82169050919050565b6000617be082617ce3565b9050919050565b6000617bf282617bf9565b9050919050565b6000617c0482617b8d565b9050919050565b6000617c1682617c1d565b9050919050565b6000617c2882617b8d565b9050919050565b6000617c3a82617c41565b9050919050565b6000617c4c82617b8d565b9050919050565b6000617c5e82617c65565b9050919050565b6000617c7082617b8d565b9050919050565b6000617c8282617c89565b9050919050565b6000617c9482617b8d565b9050919050565b6000617ca682617cad565b9050919050565b6000617cb882617b8d565b9050919050565b6000617cca82617b7f565b9050919050565b6000617cdc82617bc8565b9050919050565b6000617cee82617cf5565b9050919050565b6000617d0082617b8d565b9050919050565b60005b83811015617d25578082015181840152602081019050617d0a565b83811115617d34576000848401525b50505050565b6000601f19601f8301169050919050565b617d5481617b33565b8114617d5f57600080fd5b50565b617d6b81617b45565b8114617d7657600080fd5b50565b617d8281617b57565b8114617d8d57600080fd5b50565b617d9981617b63565b8114617da457600080fd5b50565b617db081617bad565b8114617dbb57600080fd5b50565b617dc781617bb7565b8114617dd257600080fd5b50565b617dde81617bc8565b8114617de957600080fd5b5056fea365627a7a723158203b4cdc1fcb901e4dc2ea176d84836e95dd2eaffe417996d4c348b1874344f5776c6578706572696d656e74616cf564736f6c63430005110040"

// DeployBridgeBank deploys a new Ethereum contract, binding an instance of BridgeBank to it.
func DeployBridgeBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeBank, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBankABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// BridgeBank is an auto generated Go binding around an Ethereum contract.
type BridgeBank struct {
	BridgeBankCaller     // Read-only binding to the contract
	BridgeBankTransactor // Write-only binding to the contract
	BridgeBankFilterer   // Log filterer for contract events
}

// BridgeBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBankSession struct {
	Contract     *BridgeBank       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBankCallerSession struct {
	Contract *BridgeBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBankTransactorSession struct {
	Contract     *BridgeBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBankRaw struct {
	Contract *BridgeBank // Generic contract binding to access the raw methods on
}

// BridgeBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBankCallerRaw struct {
	Contract *BridgeBankCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBankTransactorRaw struct {
	Contract *BridgeBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBank creates a new instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBank(address common.Address, backend bind.ContractBackend) (*BridgeBank, error) {
	contract, err := bindBridgeBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// NewBridgeBankCaller creates a new read-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankCaller(address common.Address, caller bind.ContractCaller) (*BridgeBankCaller, error) {
	contract, err := bindBridgeBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankCaller{contract: contract}, nil
}

// NewBridgeBankTransactor creates a new write-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBankTransactor, error) {
	contract, err := bindBridgeBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankTransactor{contract: contract}, nil
}

// NewBridgeBankFilterer creates a new log filterer instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBankFilterer, error) {
	contract, err := bindBridgeBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBankFilterer{contract: contract}, nil
}

// bindBridgeBank binds a generic wrapper to an already deployed contract.
func bindBridgeBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.BridgeBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEBANKREVISION is a free data retrieval call binding the contract method 0xf126be26.
//
// Solidity: function BRIDGEBANK_REVISION() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) BRIDGEBANKREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "BRIDGEBANK_REVISION")
	return *ret0, err
}

// BRIDGEBANKREVISION is a free data retrieval call binding the contract method 0xf126be26.
//
// Solidity: function BRIDGEBANK_REVISION() view returns(uint256)
func (_BridgeBank *BridgeBankSession) BRIDGEBANKREVISION() (*big.Int, error) {
	return _BridgeBank.Contract.BRIDGEBANKREVISION(&_BridgeBank.CallOpts)
}

// BRIDGEBANKREVISION is a free data retrieval call binding the contract method 0xf126be26.
//
// Solidity: function BRIDGEBANK_REVISION() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) BRIDGEBANKREVISION() (*big.Int, error) {
	return _BridgeBank.Contract.BRIDGEBANKREVISION(&_BridgeBank.CallOpts)
}

// ETHAddress is a free data retrieval call binding the contract method 0xda228a9b.
//
// Solidity: function ETHAddress() view returns(address)
func (_BridgeBank *BridgeBankCaller) ETHAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "ETHAddress")
	return *ret0, err
}

// ETHAddress is a free data retrieval call binding the contract method 0xda228a9b.
//
// Solidity: function ETHAddress() view returns(address)
func (_BridgeBank *BridgeBankSession) ETHAddress() (common.Address, error) {
	return _BridgeBank.Contract.ETHAddress(&_BridgeBank.CallOpts)
}

// ETHAddress is a free data retrieval call binding the contract method 0xda228a9b.
//
// Solidity: function ETHAddress() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) ETHAddress() (common.Address, error) {
	return _BridgeBank.Contract.ETHAddress(&_BridgeBank.CallOpts)
}

// ONEAddress is a free data retrieval call binding the contract method 0x9d05974b.
//
// Solidity: function ONEAddress() view returns(address)
func (_BridgeBank *BridgeBankCaller) ONEAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "ONEAddress")
	return *ret0, err
}

// ONEAddress is a free data retrieval call binding the contract method 0x9d05974b.
//
// Solidity: function ONEAddress() view returns(address)
func (_BridgeBank *BridgeBankSession) ONEAddress() (common.Address, error) {
	return _BridgeBank.Contract.ONEAddress(&_BridgeBank.CallOpts)
}

// ONEAddress is a free data retrieval call binding the contract method 0x9d05974b.
//
// Solidity: function ONEAddress() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) ONEAddress() (common.Address, error) {
	return _BridgeBank.Contract.ONEAddress(&_BridgeBank.CallOpts)
}

// SAFENUMBER is a free data retrieval call binding the contract method 0x204897ff.
//
// Solidity: function SAFE_NUMBER() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) SAFENUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "SAFE_NUMBER")
	return *ret0, err
}

// SAFENUMBER is a free data retrieval call binding the contract method 0x204897ff.
//
// Solidity: function SAFE_NUMBER() view returns(uint256)
func (_BridgeBank *BridgeBankSession) SAFENUMBER() (*big.Int, error) {
	return _BridgeBank.Contract.SAFENUMBER(&_BridgeBank.CallOpts)
}

// SAFENUMBER is a free data retrieval call binding the contract method 0x204897ff.
//
// Solidity: function SAFE_NUMBER() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) SAFENUMBER() (*big.Int, error) {
	return _BridgeBank.Contract.SAFENUMBER(&_BridgeBank.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeBank *BridgeBankCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "WETH")
	return *ret0, err
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeBank *BridgeBankSession) WETH() (common.Address, error) {
	return _BridgeBank.Contract.WETH(&_BridgeBank.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) WETH() (common.Address, error) {
	return _BridgeBank.Contract.WETH(&_BridgeBank.CallOpts)
}

// BandOracleInterface is a free data retrieval call binding the contract method 0xa1c1d3e2.
//
// Solidity: function bandOracleInterface() view returns(address)
func (_BridgeBank *BridgeBankCaller) BandOracleInterface(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "bandOracleInterface")
	return *ret0, err
}

// BandOracleInterface is a free data retrieval call binding the contract method 0xa1c1d3e2.
//
// Solidity: function bandOracleInterface() view returns(address)
func (_BridgeBank *BridgeBankSession) BandOracleInterface() (common.Address, error) {
	return _BridgeBank.Contract.BandOracleInterface(&_BridgeBank.CallOpts)
}

// BandOracleInterface is a free data retrieval call binding the contract method 0xa1c1d3e2.
//
// Solidity: function bandOracleInterface() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) BandOracleInterface() (common.Address, error) {
	return _BridgeBank.Contract.BandOracleInterface(&_BridgeBank.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_BridgeBank *BridgeBankCaller) BridgeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "bridgeRegistry")
	return *ret0, err
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_BridgeBank *BridgeBankSession) BridgeRegistry() (common.Address, error) {
	return _BridgeBank.Contract.BridgeRegistry(&_BridgeBank.CallOpts)
}

// BridgeRegistry is a free data retrieval call binding the contract method 0x316be171.
//
// Solidity: function bridgeRegistry() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) BridgeRegistry() (common.Address, error) {
	return _BridgeBank.Contract.BridgeRegistry(&_BridgeBank.CallOpts)
}

// CheckUnlockable is a free data retrieval call binding the contract method 0xb4bfd9a7.
//
// Solidity: function checkUnlockable(address _token, uint256 _amount) view returns(bool)
func (_BridgeBank *BridgeBankCaller) CheckUnlockable(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "checkUnlockable", _token, _amount)
	return *ret0, err
}

// CheckUnlockable is a free data retrieval call binding the contract method 0xb4bfd9a7.
//
// Solidity: function checkUnlockable(address _token, uint256 _amount) view returns(bool)
func (_BridgeBank *BridgeBankSession) CheckUnlockable(_token common.Address, _amount *big.Int) (bool, error) {
	return _BridgeBank.Contract.CheckUnlockable(&_BridgeBank.CallOpts, _token, _amount)
}

// CheckUnlockable is a free data retrieval call binding the contract method 0xb4bfd9a7.
//
// Solidity: function checkUnlockable(address _token, uint256 _amount) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) CheckUnlockable(_token common.Address, _amount *big.Int) (bool, error) {
	return _BridgeBank.Contract.CheckUnlockable(&_BridgeBank.CallOpts, _token, _amount)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) FeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "feeDenominator")
	return *ret0, err
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_BridgeBank *BridgeBankSession) FeeDenominator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeDenominator(&_BridgeBank.CallOpts)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) FeeDenominator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeDenominator(&_BridgeBank.CallOpts)
}

// FeeNumerator is a free data retrieval call binding the contract method 0xe86dea4a.
//
// Solidity: function feeNumerator() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) FeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "feeNumerator")
	return *ret0, err
}

// FeeNumerator is a free data retrieval call binding the contract method 0xe86dea4a.
//
// Solidity: function feeNumerator() view returns(uint256)
func (_BridgeBank *BridgeBankSession) FeeNumerator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeNumerator(&_BridgeBank.CallOpts)
}

// FeeNumerator is a free data retrieval call binding the contract method 0xe86dea4a.
//
// Solidity: function feeNumerator() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) FeeNumerator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeNumerator(&_BridgeBank.CallOpts)
}

// GetLockedFund is a free data retrieval call binding the contract method 0xfd930c7d.
//
// Solidity: function getLockedFund(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCaller) GetLockedFund(opts *bind.CallOpts, _ethereumToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getLockedFund", _ethereumToken)
	return *ret0, err
}

// GetLockedFund is a free data retrieval call binding the contract method 0xfd930c7d.
//
// Solidity: function getLockedFund(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankSession) GetLockedFund(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetLockedFund(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetLockedFund is a free data retrieval call binding the contract method 0xfd930c7d.
//
// Solidity: function getLockedFund(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) GetLockedFund(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetLockedFund(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTokenMappedAddress is a free data retrieval call binding the contract method 0x722b1d23.
//
// Solidity: function getTokenMappedAddress(address _ethereumToken) view returns(address)
func (_BridgeBank *BridgeBankCaller) GetTokenMappedAddress(opts *bind.CallOpts, _ethereumToken common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getTokenMappedAddress", _ethereumToken)
	return *ret0, err
}

// GetTokenMappedAddress is a free data retrieval call binding the contract method 0x722b1d23.
//
// Solidity: function getTokenMappedAddress(address _ethereumToken) view returns(address)
func (_BridgeBank *BridgeBankSession) GetTokenMappedAddress(_ethereumToken common.Address) (common.Address, error) {
	return _BridgeBank.Contract.GetTokenMappedAddress(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTokenMappedAddress is a free data retrieval call binding the contract method 0x722b1d23.
//
// Solidity: function getTokenMappedAddress(address _ethereumToken) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) GetTokenMappedAddress(_ethereumToken common.Address) (common.Address, error) {
	return _BridgeBank.Contract.GetTokenMappedAddress(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTotalERC20Balance is a free data retrieval call binding the contract method 0xac4a875c.
//
// Solidity: function getTotalERC20Balance(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCaller) GetTotalERC20Balance(opts *bind.CallOpts, _ethereumToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getTotalERC20Balance", _ethereumToken)
	return *ret0, err
}

// GetTotalERC20Balance is a free data retrieval call binding the contract method 0xac4a875c.
//
// Solidity: function getTotalERC20Balance(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankSession) GetTotalERC20Balance(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalERC20Balance(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTotalERC20Balance is a free data retrieval call binding the contract method 0xac4a875c.
//
// Solidity: function getTotalERC20Balance(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) GetTotalERC20Balance(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalERC20Balance(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) GetTotalETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getTotalETHBalance")
	return *ret0, err
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_BridgeBank *BridgeBankSession) GetTotalETHBalance() (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalETHBalance(&_BridgeBank.CallOpts)
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) GetTotalETHBalance() (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalETHBalance(&_BridgeBank.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeBank *BridgeBankCaller) HarmonyBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "harmonyBridge")
	return *ret0, err
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeBank *BridgeBankSession) HarmonyBridge() (common.Address, error) {
	return _BridgeBank.Contract.HarmonyBridge(&_BridgeBank.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) HarmonyBridge() (common.Address, error) {
	return _BridgeBank.Contract.HarmonyBridge(&_BridgeBank.CallOpts)
}

// IsActiveToken is a free data retrieval call binding the contract method 0x575c51c2.
//
// Solidity: function isActiveToken(address _ethereumToken) view returns(bool)
func (_BridgeBank *BridgeBankCaller) IsActiveToken(opts *bind.CallOpts, _ethereumToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "isActiveToken", _ethereumToken)
	return *ret0, err
}

// IsActiveToken is a free data retrieval call binding the contract method 0x575c51c2.
//
// Solidity: function isActiveToken(address _ethereumToken) view returns(bool)
func (_BridgeBank *BridgeBankSession) IsActiveToken(_ethereumToken common.Address) (bool, error) {
	return _BridgeBank.Contract.IsActiveToken(&_BridgeBank.CallOpts, _ethereumToken)
}

// IsActiveToken is a free data retrieval call binding the contract method 0x575c51c2.
//
// Solidity: function isActiveToken(address _ethereumToken) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) IsActiveToken(_ethereumToken common.Address) (bool, error) {
	return _BridgeBank.Contract.IsActiveToken(&_BridgeBank.CallOpts, _ethereumToken)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_BridgeBank *BridgeBankCaller) LendingPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "lendingPool")
	return *ret0, err
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_BridgeBank *BridgeBankSession) LendingPool() (common.Address, error) {
	return _BridgeBank.Contract.LendingPool(&_BridgeBank.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) LendingPool() (common.Address, error) {
	return _BridgeBank.Contract.LendingPool(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "lockNonce")
	return *ret0, err
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// TokensData is a free data retrieval call binding the contract method 0xac66c6e0.
//
// Solidity: function tokensData(address ) view returns(uint256 lockedFund, address harmonyMappedToken, bool isActive)
func (_BridgeBank *BridgeBankCaller) TokensData(opts *bind.CallOpts, arg0 common.Address) (struct {
	LockedFund         *big.Int
	HarmonyMappedToken common.Address
	IsActive           bool
}, error) {
	ret := new(struct {
		LockedFund         *big.Int
		HarmonyMappedToken common.Address
		IsActive           bool
	})
	out := ret
	err := _BridgeBank.contract.Call(opts, out, "tokensData", arg0)
	return *ret, err
}

// TokensData is a free data retrieval call binding the contract method 0xac66c6e0.
//
// Solidity: function tokensData(address ) view returns(uint256 lockedFund, address harmonyMappedToken, bool isActive)
func (_BridgeBank *BridgeBankSession) TokensData(arg0 common.Address) (struct {
	LockedFund         *big.Int
	HarmonyMappedToken common.Address
	IsActive           bool
}, error) {
	return _BridgeBank.Contract.TokensData(&_BridgeBank.CallOpts, arg0)
}

// TokensData is a free data retrieval call binding the contract method 0xac66c6e0.
//
// Solidity: function tokensData(address ) view returns(uint256 lockedFund, address harmonyMappedToken, bool isActive)
func (_BridgeBank *BridgeBankCallerSession) TokensData(arg0 common.Address) (struct {
	LockedFund         *big.Int
	HarmonyMappedToken common.Address
	IsActive           bool
}, error) {
	return _BridgeBank.Contract.TokensData(&_BridgeBank.CallOpts, arg0)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_BridgeBank *BridgeBankCaller) WethGateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "wethGateway")
	return *ret0, err
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_BridgeBank *BridgeBankSession) WethGateway() (common.Address, error) {
	return _BridgeBank.Contract.WethGateway(&_BridgeBank.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) WethGateway() (common.Address, error) {
	return _BridgeBank.Contract.WethGateway(&_BridgeBank.CallOpts)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) ActivateToken(opts *bind.TransactOpts, _ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "activateToken", _ethereumToken)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankSession) ActivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ActivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) ActivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ActivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _ethereumToken, uint256 _ethereumTokenAmount, address _harmonyToken) returns()
func (_BridgeBank *BridgeBankTransactor) AddToken(opts *bind.TransactOpts, _ethereumToken common.Address, _ethereumTokenAmount *big.Int, _harmonyToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "addToken", _ethereumToken, _ethereumTokenAmount, _harmonyToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _ethereumToken, uint256 _ethereumTokenAmount, address _harmonyToken) returns()
func (_BridgeBank *BridgeBankSession) AddToken(_ethereumToken common.Address, _ethereumTokenAmount *big.Int, _harmonyToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddToken(&_BridgeBank.TransactOpts, _ethereumToken, _ethereumTokenAmount, _harmonyToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _ethereumToken, uint256 _ethereumTokenAmount, address _harmonyToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) AddToken(_ethereumToken common.Address, _ethereumTokenAmount *big.Int, _harmonyToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddToken(&_BridgeBank.TransactOpts, _ethereumToken, _ethereumTokenAmount, _harmonyToken)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) DeactivateToken(opts *bind.TransactOpts, _ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "deactivateToken", _ethereumToken)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankSession) DeactivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.DeactivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) DeactivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.DeactivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _bridgeRegistry, address _bandOracleAddress, address _lendingPool, address _wethGateway, address _weth, address _harmonyETH) payable returns()
func (_BridgeBank *BridgeBankTransactor) Initialize(opts *bind.TransactOpts, _bridgeRegistry common.Address, _bandOracleAddress common.Address, _lendingPool common.Address, _wethGateway common.Address, _weth common.Address, _harmonyETH common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "initialize", _bridgeRegistry, _bandOracleAddress, _lendingPool, _wethGateway, _weth, _harmonyETH)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _bridgeRegistry, address _bandOracleAddress, address _lendingPool, address _wethGateway, address _weth, address _harmonyETH) payable returns()
func (_BridgeBank *BridgeBankSession) Initialize(_bridgeRegistry common.Address, _bandOracleAddress common.Address, _lendingPool common.Address, _wethGateway common.Address, _weth common.Address, _harmonyETH common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize(&_BridgeBank.TransactOpts, _bridgeRegistry, _bandOracleAddress, _lendingPool, _wethGateway, _weth, _harmonyETH)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _bridgeRegistry, address _bandOracleAddress, address _lendingPool, address _wethGateway, address _weth, address _harmonyETH) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Initialize(_bridgeRegistry common.Address, _bandOracleAddress common.Address, _lendingPool common.Address, _wethGateway common.Address, _weth common.Address, _harmonyETH common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize(&_BridgeBank.TransactOpts, _bridgeRegistry, _bandOracleAddress, _lendingPool, _wethGateway, _weth, _harmonyETH)
}

// SwapETHForONE is a paid mutator transaction binding the contract method 0x5e01297d.
//
// Solidity: function swapETHForONE(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactor) SwapETHForONE(opts *bind.TransactOpts, _harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapETHForONE", _harmonyReceiver, _amountETH)
}

// SwapETHForONE is a paid mutator transaction binding the contract method 0x5e01297d.
//
// Solidity: function swapETHForONE(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankSession) SwapETHForONE(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapETHForONE is a paid mutator transaction binding the contract method 0x5e01297d.
//
// Solidity: function swapETHForONE(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapETHForONE(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapETHForToken is a paid mutator transaction binding the contract method 0x903e10ba.
//
// Solidity: function swapETHForToken(address _harmonyReceiver, uint256 _amountETH, address _destToken) payable returns()
func (_BridgeBank *BridgeBankTransactor) SwapETHForToken(opts *bind.TransactOpts, _harmonyReceiver common.Address, _amountETH *big.Int, _destToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapETHForToken", _harmonyReceiver, _amountETH, _destToken)
}

// SwapETHForToken is a paid mutator transaction binding the contract method 0x903e10ba.
//
// Solidity: function swapETHForToken(address _harmonyReceiver, uint256 _amountETH, address _destToken) payable returns()
func (_BridgeBank *BridgeBankSession) SwapETHForToken(_harmonyReceiver common.Address, _amountETH *big.Int, _destToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH, _destToken)
}

// SwapETHForToken is a paid mutator transaction binding the contract method 0x903e10ba.
//
// Solidity: function swapETHForToken(address _harmonyReceiver, uint256 _amountETH, address _destToken) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapETHForToken(_harmonyReceiver common.Address, _amountETH *big.Int, _destToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH, _destToken)
}

// SwapETHForWrappedETH is a paid mutator transaction binding the contract method 0x7fabc432.
//
// Solidity: function swapETHForWrappedETH(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactor) SwapETHForWrappedETH(opts *bind.TransactOpts, _harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapETHForWrappedETH", _harmonyReceiver, _amountETH)
}

// SwapETHForWrappedETH is a paid mutator transaction binding the contract method 0x7fabc432.
//
// Solidity: function swapETHForWrappedETH(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankSession) SwapETHForWrappedETH(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForWrappedETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapETHForWrappedETH is a paid mutator transaction binding the contract method 0x7fabc432.
//
// Solidity: function swapETHForWrappedETH(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapETHForWrappedETH(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForWrappedETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapTokenForONE is a paid mutator transaction binding the contract method 0xba1adba4.
//
// Solidity: function swapTokenForONE(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactor) SwapTokenForONE(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapTokenForONE", _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapTokenForONE is a paid mutator transaction binding the contract method 0xba1adba4.
//
// Solidity: function swapTokenForONE(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankSession) SwapTokenForONE(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapTokenForONE is a paid mutator transaction binding the contract method 0xba1adba4.
//
// Solidity: function swapTokenForONE(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapTokenForONE(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapTokenForToken is a paid mutator transaction binding the contract method 0x9b92910c.
//
// Solidity: function swapTokenForToken(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount, address _destToken) returns()
func (_BridgeBank *BridgeBankTransactor) SwapTokenForToken(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int, _destToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapTokenForToken", _harmonyReceiver, _ethereumToken, _ethereumTokenAmount, _destToken)
}

// SwapTokenForToken is a paid mutator transaction binding the contract method 0x9b92910c.
//
// Solidity: function swapTokenForToken(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount, address _destToken) returns()
func (_BridgeBank *BridgeBankSession) SwapTokenForToken(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int, _destToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount, _destToken)
}

// SwapTokenForToken is a paid mutator transaction binding the contract method 0x9b92910c.
//
// Solidity: function swapTokenForToken(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount, address _destToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapTokenForToken(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int, _destToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount, _destToken)
}

// SwapTokenForWrappedETH is a paid mutator transaction binding the contract method 0x49b1af93.
//
// Solidity: function swapTokenForWrappedETH(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactor) SwapTokenForWrappedETH(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapTokenForWrappedETH", _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapTokenForWrappedETH is a paid mutator transaction binding the contract method 0x49b1af93.
//
// Solidity: function swapTokenForWrappedETH(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankSession) SwapTokenForWrappedETH(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForWrappedETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapTokenForWrappedETH is a paid mutator transaction binding the contract method 0x49b1af93.
//
// Solidity: function swapTokenForWrappedETH(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapTokenForWrappedETH(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForWrappedETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapToken11 is a paid mutator transaction binding the contract method 0xbb0a64db.
//
// Solidity: function swapToken_1_1(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactor) SwapToken11(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapToken_1_1", _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapToken11 is a paid mutator transaction binding the contract method 0xbb0a64db.
//
// Solidity: function swapToken_1_1(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankSession) SwapToken11(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapToken11(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// SwapToken11 is a paid mutator transaction binding the contract method 0xbb0a64db.
//
// Solidity: function swapToken_1_1(address _harmonyReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapToken11(_harmonyReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapToken11(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _ethereumReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactor) UnlockERC20(opts *bind.TransactOpts, _ethereumReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unlockERC20", _ethereumReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _ethereumReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankSession) UnlockERC20(_ethereumReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockERC20(&_BridgeBank.TransactOpts, _ethereumReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _ethereumReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactorSession) UnlockERC20(_ethereumReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockERC20(&_BridgeBank.TransactOpts, _ethereumReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockETH is a paid mutator transaction binding the contract method 0xed581fd3.
//
// Solidity: function unlockETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactor) UnlockETH(opts *bind.TransactOpts, _ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unlockETH", _ethereumReceiver, _amountETH)
}

// UnlockETH is a paid mutator transaction binding the contract method 0xed581fd3.
//
// Solidity: function unlockETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankSession) UnlockETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// UnlockETH is a paid mutator transaction binding the contract method 0xed581fd3.
//
// Solidity: function unlockETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactorSession) UnlockETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x2740c197.
//
// Solidity: function updateFee(uint256 _feeNumerator, uint256 _feeDenominator) returns()
func (_BridgeBank *BridgeBankTransactor) UpdateFee(opts *bind.TransactOpts, _feeNumerator *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "updateFee", _feeNumerator, _feeDenominator)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x2740c197.
//
// Solidity: function updateFee(uint256 _feeNumerator, uint256 _feeDenominator) returns()
func (_BridgeBank *BridgeBankSession) UpdateFee(_feeNumerator *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateFee(&_BridgeBank.TransactOpts, _feeNumerator, _feeDenominator)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x2740c197.
//
// Solidity: function updateFee(uint256 _feeNumerator, uint256 _feeDenominator) returns()
func (_BridgeBank *BridgeBankTransactorSession) UpdateFee(_feeNumerator *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateFee(&_BridgeBank.TransactOpts, _feeNumerator, _feeDenominator)
}

// UpdateHmyBridge is a paid mutator transaction binding the contract method 0x17877820.
//
// Solidity: function updateHmyBridge(address _harmonyBridge) returns()
func (_BridgeBank *BridgeBankTransactor) UpdateHmyBridge(opts *bind.TransactOpts, _harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "updateHmyBridge", _harmonyBridge)
}

// UpdateHmyBridge is a paid mutator transaction binding the contract method 0x17877820.
//
// Solidity: function updateHmyBridge(address _harmonyBridge) returns()
func (_BridgeBank *BridgeBankSession) UpdateHmyBridge(_harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateHmyBridge(&_BridgeBank.TransactOpts, _harmonyBridge)
}

// UpdateHmyBridge is a paid mutator transaction binding the contract method 0x17877820.
//
// Solidity: function updateHmyBridge(address _harmonyBridge) returns()
func (_BridgeBank *BridgeBankTransactorSession) UpdateHmyBridge(_harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateHmyBridge(&_BridgeBank.TransactOpts, _harmonyBridge)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x1cb44dfc.
//
// Solidity: function updateOracle(address _oracleAddress) returns()
func (_BridgeBank *BridgeBankTransactor) UpdateOracle(opts *bind.TransactOpts, _oracleAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "updateOracle", _oracleAddress)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x1cb44dfc.
//
// Solidity: function updateOracle(address _oracleAddress) returns()
func (_BridgeBank *BridgeBankSession) UpdateOracle(_oracleAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateOracle(&_BridgeBank.TransactOpts, _oracleAddress)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x1cb44dfc.
//
// Solidity: function updateOracle(address _oracleAddress) returns()
func (_BridgeBank *BridgeBankTransactorSession) UpdateOracle(_oracleAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateOracle(&_BridgeBank.TransactOpts, _oracleAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _ethereumToken, address _ethereumReceiver, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactor) WithdrawERC20(opts *bind.TransactOpts, _ethereumToken common.Address, _ethereumReceiver common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "withdrawERC20", _ethereumToken, _ethereumReceiver, _ethereumTokenAmount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _ethereumToken, address _ethereumReceiver, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankSession) WithdrawERC20(_ethereumToken common.Address, _ethereumReceiver common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawERC20(&_BridgeBank.TransactOpts, _ethereumToken, _ethereumReceiver, _ethereumTokenAmount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _ethereumToken, address _ethereumReceiver, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactorSession) WithdrawERC20(_ethereumToken common.Address, _ethereumReceiver common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawERC20(&_BridgeBank.TransactOpts, _ethereumToken, _ethereumReceiver, _ethereumTokenAmount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactor) WithdrawETH(opts *bind.TransactOpts, _ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "withdrawETH", _ethereumReceiver, _amountETH)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankSession) WithdrawETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactorSession) WithdrawETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.Contract.Fallback(&_BridgeBank.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.Contract.Fallback(&_BridgeBank.TransactOpts, calldata)
}

// BridgeBankEthLogLockIterator is returned from FilterEthLogLock and is used to iterate over the raw logs and unpacked data for EthLogLock events raised by the BridgeBank contract.
type BridgeBankEthLogLockIterator struct {
	Event *BridgeBankEthLogLock // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthLogLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthLogLock)
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
		it.Event = new(BridgeBankEthLogLock)
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
func (it *BridgeBankEthLogLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthLogLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthLogLock represents a EthLogLock event raised by the BridgeBank contract.
type BridgeBankEthLogLock struct {
	EthereumSender      common.Address
	HarmonyReceiver     common.Address
	EthereumToken       common.Address
	HarmonyToken        common.Address
	EthereumTokenAmount *big.Int
	HarmonyTokenAmount  *big.Int
	Nonce               *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEthLogLock is a free log retrieval operation binding the contract event 0x15bb5885b6550f0baf0cf174b9618b2d0c83f6ab83ea0de29b03cb6c9461afd5.
//
// Solidity: event EthLogLock(address _ethereumSender, address _harmonyReceiver, address _ethereumToken, address _harmonyToken, uint256 _ethereumTokenAmount, uint256 _harmonyTokenAmount, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) FilterEthLogLock(opts *bind.FilterOpts) (*BridgeBankEthLogLockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthLogLock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthLogLockIterator{contract: _BridgeBank.contract, event: "EthLogLock", logs: logs, sub: sub}, nil
}

// WatchEthLogLock is a free log subscription operation binding the contract event 0x15bb5885b6550f0baf0cf174b9618b2d0c83f6ab83ea0de29b03cb6c9461afd5.
//
// Solidity: event EthLogLock(address _ethereumSender, address _harmonyReceiver, address _ethereumToken, address _harmonyToken, uint256 _ethereumTokenAmount, uint256 _harmonyTokenAmount, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) WatchEthLogLock(opts *bind.WatchOpts, sink chan<- *BridgeBankEthLogLock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthLogLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthLogLock)
				if err := _BridgeBank.contract.UnpackLog(event, "EthLogLock", log); err != nil {
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

// ParseEthLogLock is a log parse operation binding the contract event 0x15bb5885b6550f0baf0cf174b9618b2d0c83f6ab83ea0de29b03cb6c9461afd5.
//
// Solidity: event EthLogLock(address _ethereumSender, address _harmonyReceiver, address _ethereumToken, address _harmonyToken, uint256 _ethereumTokenAmount, uint256 _harmonyTokenAmount, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) ParseEthLogLock(log types.Log) (*BridgeBankEthLogLock, error) {
	event := new(BridgeBankEthLogLock)
	if err := _BridgeBank.contract.UnpackLog(event, "EthLogLock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankEthLogUnlockIterator is returned from FilterEthLogUnlock and is used to iterate over the raw logs and unpacked data for EthLogUnlock events raised by the BridgeBank contract.
type BridgeBankEthLogUnlockIterator struct {
	Event *BridgeBankEthLogUnlock // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthLogUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthLogUnlock)
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
		it.Event = new(BridgeBankEthLogUnlock)
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
func (it *BridgeBankEthLogUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthLogUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthLogUnlock represents a EthLogUnlock event raised by the BridgeBank contract.
type BridgeBankEthLogUnlock struct {
	To    common.Address
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthLogUnlock is a free log retrieval operation binding the contract event 0x4445dabfef575721ee82c521f0725ef2cab18153d75cadda8115c2e31efc5508.
//
// Solidity: event EthLogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) FilterEthLogUnlock(opts *bind.FilterOpts) (*BridgeBankEthLogUnlockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthLogUnlock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthLogUnlockIterator{contract: _BridgeBank.contract, event: "EthLogUnlock", logs: logs, sub: sub}, nil
}

// WatchEthLogUnlock is a free log subscription operation binding the contract event 0x4445dabfef575721ee82c521f0725ef2cab18153d75cadda8115c2e31efc5508.
//
// Solidity: event EthLogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) WatchEthLogUnlock(opts *bind.WatchOpts, sink chan<- *BridgeBankEthLogUnlock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthLogUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthLogUnlock)
				if err := _BridgeBank.contract.UnpackLog(event, "EthLogUnlock", log); err != nil {
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

// ParseEthLogUnlock is a log parse operation binding the contract event 0x4445dabfef575721ee82c521f0725ef2cab18153d75cadda8115c2e31efc5508.
//
// Solidity: event EthLogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) ParseEthLogUnlock(log types.Log) (*BridgeBankEthLogUnlock, error) {
	event := new(BridgeBankEthLogUnlock)
	if err := _BridgeBank.contract.UnpackLog(event, "EthLogUnlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankEthUpdateFeeIterator is returned from FilterEthUpdateFee and is used to iterate over the raw logs and unpacked data for EthUpdateFee events raised by the BridgeBank contract.
type BridgeBankEthUpdateFeeIterator struct {
	Event *BridgeBankEthUpdateFee // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthUpdateFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthUpdateFee)
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
		it.Event = new(BridgeBankEthUpdateFee)
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
func (it *BridgeBankEthUpdateFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthUpdateFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthUpdateFee represents a EthUpdateFee event raised by the BridgeBank contract.
type BridgeBankEthUpdateFee struct {
	FeeNumerator   *big.Int
	FeeDenominator *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEthUpdateFee is a free log retrieval operation binding the contract event 0x87018147d45d73a403e53a1dae7a7e95ad1bf96326b915560c0fad22373a3521.
//
// Solidity: event EthUpdateFee(uint256 _feeNumerator, uint256 _feeDenominator)
func (_BridgeBank *BridgeBankFilterer) FilterEthUpdateFee(opts *bind.FilterOpts) (*BridgeBankEthUpdateFeeIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthUpdateFee")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthUpdateFeeIterator{contract: _BridgeBank.contract, event: "EthUpdateFee", logs: logs, sub: sub}, nil
}

// WatchEthUpdateFee is a free log subscription operation binding the contract event 0x87018147d45d73a403e53a1dae7a7e95ad1bf96326b915560c0fad22373a3521.
//
// Solidity: event EthUpdateFee(uint256 _feeNumerator, uint256 _feeDenominator)
func (_BridgeBank *BridgeBankFilterer) WatchEthUpdateFee(opts *bind.WatchOpts, sink chan<- *BridgeBankEthUpdateFee) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthUpdateFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthUpdateFee)
				if err := _BridgeBank.contract.UnpackLog(event, "EthUpdateFee", log); err != nil {
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

// ParseEthUpdateFee is a log parse operation binding the contract event 0x87018147d45d73a403e53a1dae7a7e95ad1bf96326b915560c0fad22373a3521.
//
// Solidity: event EthUpdateFee(uint256 _feeNumerator, uint256 _feeDenominator)
func (_BridgeBank *BridgeBankFilterer) ParseEthUpdateFee(log types.Log) (*BridgeBankEthUpdateFee, error) {
	event := new(BridgeBankEthUpdateFee)
	if err := _BridgeBank.contract.UnpackLog(event, "EthUpdateFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankEthUpdateHarmonyBridgeIterator is returned from FilterEthUpdateHarmonyBridge and is used to iterate over the raw logs and unpacked data for EthUpdateHarmonyBridge events raised by the BridgeBank contract.
type BridgeBankEthUpdateHarmonyBridgeIterator struct {
	Event *BridgeBankEthUpdateHarmonyBridge // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthUpdateHarmonyBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthUpdateHarmonyBridge)
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
		it.Event = new(BridgeBankEthUpdateHarmonyBridge)
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
func (it *BridgeBankEthUpdateHarmonyBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthUpdateHarmonyBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthUpdateHarmonyBridge represents a EthUpdateHarmonyBridge event raised by the BridgeBank contract.
type BridgeBankEthUpdateHarmonyBridge struct {
	NewHarmonyBridge common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEthUpdateHarmonyBridge is a free log retrieval operation binding the contract event 0x2ee0580b26ac9d813eb02bff2dc08fbf08346dafcc883f3c21c44d014b8a2e85.
//
// Solidity: event EthUpdateHarmonyBridge(address _newHarmonyBridge)
func (_BridgeBank *BridgeBankFilterer) FilterEthUpdateHarmonyBridge(opts *bind.FilterOpts) (*BridgeBankEthUpdateHarmonyBridgeIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthUpdateHarmonyBridge")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthUpdateHarmonyBridgeIterator{contract: _BridgeBank.contract, event: "EthUpdateHarmonyBridge", logs: logs, sub: sub}, nil
}

// WatchEthUpdateHarmonyBridge is a free log subscription operation binding the contract event 0x2ee0580b26ac9d813eb02bff2dc08fbf08346dafcc883f3c21c44d014b8a2e85.
//
// Solidity: event EthUpdateHarmonyBridge(address _newHarmonyBridge)
func (_BridgeBank *BridgeBankFilterer) WatchEthUpdateHarmonyBridge(opts *bind.WatchOpts, sink chan<- *BridgeBankEthUpdateHarmonyBridge) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthUpdateHarmonyBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthUpdateHarmonyBridge)
				if err := _BridgeBank.contract.UnpackLog(event, "EthUpdateHarmonyBridge", log); err != nil {
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

// ParseEthUpdateHarmonyBridge is a log parse operation binding the contract event 0x2ee0580b26ac9d813eb02bff2dc08fbf08346dafcc883f3c21c44d014b8a2e85.
//
// Solidity: event EthUpdateHarmonyBridge(address _newHarmonyBridge)
func (_BridgeBank *BridgeBankFilterer) ParseEthUpdateHarmonyBridge(log types.Log) (*BridgeBankEthUpdateHarmonyBridge, error) {
	event := new(BridgeBankEthUpdateHarmonyBridge)
	if err := _BridgeBank.contract.UnpackLog(event, "EthUpdateHarmonyBridge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankEthUpdateOracleIterator is returned from FilterEthUpdateOracle and is used to iterate over the raw logs and unpacked data for EthUpdateOracle events raised by the BridgeBank contract.
type BridgeBankEthUpdateOracleIterator struct {
	Event *BridgeBankEthUpdateOracle // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthUpdateOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthUpdateOracle)
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
		it.Event = new(BridgeBankEthUpdateOracle)
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
func (it *BridgeBankEthUpdateOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthUpdateOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthUpdateOracle represents a EthUpdateOracle event raised by the BridgeBank contract.
type BridgeBankEthUpdateOracle struct {
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthUpdateOracle is a free log retrieval operation binding the contract event 0xe0dec806007e306f472d1bb9e8dcf714fbb3da18e34206ea34f2d2c4bba89659.
//
// Solidity: event EthUpdateOracle(address _newOracle)
func (_BridgeBank *BridgeBankFilterer) FilterEthUpdateOracle(opts *bind.FilterOpts) (*BridgeBankEthUpdateOracleIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthUpdateOracle")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthUpdateOracleIterator{contract: _BridgeBank.contract, event: "EthUpdateOracle", logs: logs, sub: sub}, nil
}

// WatchEthUpdateOracle is a free log subscription operation binding the contract event 0xe0dec806007e306f472d1bb9e8dcf714fbb3da18e34206ea34f2d2c4bba89659.
//
// Solidity: event EthUpdateOracle(address _newOracle)
func (_BridgeBank *BridgeBankFilterer) WatchEthUpdateOracle(opts *bind.WatchOpts, sink chan<- *BridgeBankEthUpdateOracle) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthUpdateOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthUpdateOracle)
				if err := _BridgeBank.contract.UnpackLog(event, "EthUpdateOracle", log); err != nil {
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

// ParseEthUpdateOracle is a log parse operation binding the contract event 0xe0dec806007e306f472d1bb9e8dcf714fbb3da18e34206ea34f2d2c4bba89659.
//
// Solidity: event EthUpdateOracle(address _newOracle)
func (_BridgeBank *BridgeBankFilterer) ParseEthUpdateOracle(log types.Log) (*BridgeBankEthUpdateOracle, error) {
	event := new(BridgeBankEthUpdateOracle)
	if err := _BridgeBank.contract.UnpackLog(event, "EthUpdateOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankEthWithdrawERC20Iterator is returned from FilterEthWithdrawERC20 and is used to iterate over the raw logs and unpacked data for EthWithdrawERC20 events raised by the BridgeBank contract.
type BridgeBankEthWithdrawERC20Iterator struct {
	Event *BridgeBankEthWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthWithdrawERC20)
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
		it.Event = new(BridgeBankEthWithdrawERC20)
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
func (it *BridgeBankEthWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthWithdrawERC20 represents a EthWithdrawERC20 event raised by the BridgeBank contract.
type BridgeBankEthWithdrawERC20 struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawERC20 is a free log retrieval operation binding the contract event 0x378ae7e321aef491da8299462610eb48bb07e3b93619adfb2cdbe332b6aad7ce.
//
// Solidity: event EthWithdrawERC20(address _token, address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) FilterEthWithdrawERC20(opts *bind.FilterOpts) (*BridgeBankEthWithdrawERC20Iterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthWithdrawERC20")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthWithdrawERC20Iterator{contract: _BridgeBank.contract, event: "EthWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawERC20 is a free log subscription operation binding the contract event 0x378ae7e321aef491da8299462610eb48bb07e3b93619adfb2cdbe332b6aad7ce.
//
// Solidity: event EthWithdrawERC20(address _token, address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) WatchEthWithdrawERC20(opts *bind.WatchOpts, sink chan<- *BridgeBankEthWithdrawERC20) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthWithdrawERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthWithdrawERC20)
				if err := _BridgeBank.contract.UnpackLog(event, "EthWithdrawERC20", log); err != nil {
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

// ParseEthWithdrawERC20 is a log parse operation binding the contract event 0x378ae7e321aef491da8299462610eb48bb07e3b93619adfb2cdbe332b6aad7ce.
//
// Solidity: event EthWithdrawERC20(address _token, address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) ParseEthWithdrawERC20(log types.Log) (*BridgeBankEthWithdrawERC20, error) {
	event := new(BridgeBankEthWithdrawERC20)
	if err := _BridgeBank.contract.UnpackLog(event, "EthWithdrawERC20", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankEthWithdrawETHIterator is returned from FilterEthWithdrawETH and is used to iterate over the raw logs and unpacked data for EthWithdrawETH events raised by the BridgeBank contract.
type BridgeBankEthWithdrawETHIterator struct {
	Event *BridgeBankEthWithdrawETH // Event containing the contract specifics and raw log

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
func (it *BridgeBankEthWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankEthWithdrawETH)
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
		it.Event = new(BridgeBankEthWithdrawETH)
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
func (it *BridgeBankEthWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankEthWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankEthWithdrawETH represents a EthWithdrawETH event raised by the BridgeBank contract.
type BridgeBankEthWithdrawETH struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawETH is a free log retrieval operation binding the contract event 0xc7d71dfdba3c33b60061c39022c106fd785ed1edda60862045a0b684ca1f0e5d.
//
// Solidity: event EthWithdrawETH(address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) FilterEthWithdrawETH(opts *bind.FilterOpts) (*BridgeBankEthWithdrawETHIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "EthWithdrawETH")
	if err != nil {
		return nil, err
	}
	return &BridgeBankEthWithdrawETHIterator{contract: _BridgeBank.contract, event: "EthWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawETH is a free log subscription operation binding the contract event 0xc7d71dfdba3c33b60061c39022c106fd785ed1edda60862045a0b684ca1f0e5d.
//
// Solidity: event EthWithdrawETH(address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) WatchEthWithdrawETH(opts *bind.WatchOpts, sink chan<- *BridgeBankEthWithdrawETH) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "EthWithdrawETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankEthWithdrawETH)
				if err := _BridgeBank.contract.UnpackLog(event, "EthWithdrawETH", log); err != nil {
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

// ParseEthWithdrawETH is a log parse operation binding the contract event 0xc7d71dfdba3c33b60061c39022c106fd785ed1edda60862045a0b684ca1f0e5d.
//
// Solidity: event EthWithdrawETH(address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) ParseEthWithdrawETH(log types.Log) (*BridgeBankEthWithdrawETH, error) {
	event := new(BridgeBankEthWithdrawETH)
	if err := _BridgeBank.contract.UnpackLog(event, "EthWithdrawETH", log); err != nil {
		return nil, err
	}
	return event, nil
}
