// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethapi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// KolKeySimpleMetaData contains all meta data concerning the KolKeySimple contract.
var KolKeySimpleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"opType\",\"type\":\"bool\"}],\"name\":\"AdminOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdrawAllIncome\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdrawByOneKol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdrawByOneNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"KeyRebound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeyTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"KeyTransferAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"sourceID\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceConract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyNo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valPerKey\",\"type\":\"uint256\"}],\"name\":\"KolIncomeToPoolAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kolAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyNo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"KoltotalNo\",\"type\":\"uint256\"}],\"name\":\"KolKeyBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerKey\",\"type\":\"uint256\"}],\"name\":\"KolKeyOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSerficeFeeRate\",\"type\":\"uint256\"}],\"name\":\"ServiceFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"op\",\"type\":\"string\"}],\"name\":\"SystemSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"UpgradeToNewRule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"WithdrawService\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"AllIncomeOfAllKol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"AllIncomeOfOneKol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AllKolAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"IncomeOfOneNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncomeOfOneNonceByAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"InvestorAllKeysOfKol\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonce\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"InvestorOfKol\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"KeySettingsRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNo\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"KolOfOneInvestor\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"__admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__minValCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDelete\",\"type\":\"bool\"}],\"name\":\"adminOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminServiceFeeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetKeyFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetKolIncomeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyNo\",\"type\":\"uint256\"}],\"name\":\"adminSetMaxKolKeyNo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetWithdrawFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"adminUpgradeToNewRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allKolInSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kolAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keyNo\",\"type\":\"uint256\"}],\"name\":\"buyKolKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPluginInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"incomePerNoncePerKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"sourceID\",\"type\":\"int8\"},{\"internalType\":\"address\",\"name\":\"kolAddr\",\"type\":\"address\"}],\"name\":\"kolGotIncome\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolIncomeRatePerKeyBuy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceAddr\",\"type\":\"address\"}],\"name\":\"kolOpenKeyPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInFin\",\"type\":\"uint256\"}],\"name\":\"kolOpenKeySale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKeyNoForKol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"rebindKolKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeRatePerKeyBuy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"transferAllKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllIncome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"withdrawFromOneKol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"withdrawFromOneNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KolKeySimpleABI is the input ABI used to generate the binding from.
// Deprecated: Use KolKeySimpleMetaData.ABI instead.
var KolKeySimpleABI = KolKeySimpleMetaData.ABI

// KolKeySimple is an auto generated Go binding around an Ethereum contract.
type KolKeySimple struct {
	KolKeySimpleCaller     // Read-only binding to the contract
	KolKeySimpleTransactor // Write-only binding to the contract
	KolKeySimpleFilterer   // Log filterer for contract events
}

// KolKeySimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type KolKeySimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KolKeySimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KolKeySimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KolKeySimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KolKeySimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KolKeySimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KolKeySimpleSession struct {
	Contract     *KolKeySimple     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KolKeySimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KolKeySimpleCallerSession struct {
	Contract *KolKeySimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KolKeySimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KolKeySimpleTransactorSession struct {
	Contract     *KolKeySimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KolKeySimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type KolKeySimpleRaw struct {
	Contract *KolKeySimple // Generic contract binding to access the raw methods on
}

// KolKeySimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KolKeySimpleCallerRaw struct {
	Contract *KolKeySimpleCaller // Generic read-only contract binding to access the raw methods on
}

// KolKeySimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KolKeySimpleTransactorRaw struct {
	Contract *KolKeySimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKolKeySimple creates a new instance of KolKeySimple, bound to a specific deployed contract.
func NewKolKeySimple(address common.Address, backend bind.ContractBackend) (*KolKeySimple, error) {
	contract, err := bindKolKeySimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KolKeySimple{KolKeySimpleCaller: KolKeySimpleCaller{contract: contract}, KolKeySimpleTransactor: KolKeySimpleTransactor{contract: contract}, KolKeySimpleFilterer: KolKeySimpleFilterer{contract: contract}}, nil
}

// NewKolKeySimpleCaller creates a new read-only instance of KolKeySimple, bound to a specific deployed contract.
func NewKolKeySimpleCaller(address common.Address, caller bind.ContractCaller) (*KolKeySimpleCaller, error) {
	contract, err := bindKolKeySimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleCaller{contract: contract}, nil
}

// NewKolKeySimpleTransactor creates a new write-only instance of KolKeySimple, bound to a specific deployed contract.
func NewKolKeySimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*KolKeySimpleTransactor, error) {
	contract, err := bindKolKeySimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleTransactor{contract: contract}, nil
}

// NewKolKeySimpleFilterer creates a new log filterer instance of KolKeySimple, bound to a specific deployed contract.
func NewKolKeySimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*KolKeySimpleFilterer, error) {
	contract, err := bindKolKeySimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleFilterer{contract: contract}, nil
}

// bindKolKeySimple binds a generic wrapper to an already deployed contract.
func bindKolKeySimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KolKeySimpleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KolKeySimple *KolKeySimpleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KolKeySimple.Contract.KolKeySimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KolKeySimple *KolKeySimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeySimple.Contract.KolKeySimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KolKeySimple *KolKeySimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KolKeySimple.Contract.KolKeySimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KolKeySimple *KolKeySimpleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KolKeySimple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KolKeySimple *KolKeySimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeySimple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KolKeySimple *KolKeySimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KolKeySimple.Contract.contract.Transact(opts, method, params...)
}

// AllIncomeOfAllKol is a free data retrieval call binding the contract method 0xaf958a62.
//
// Solidity: function AllIncomeOfAllKol(address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) AllIncomeOfAllKol(opts *bind.CallOpts, investor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "AllIncomeOfAllKol", investor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllIncomeOfAllKol is a free data retrieval call binding the contract method 0xaf958a62.
//
// Solidity: function AllIncomeOfAllKol(address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) AllIncomeOfAllKol(investor common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.AllIncomeOfAllKol(&_KolKeySimple.CallOpts, investor)
}

// AllIncomeOfAllKol is a free data retrieval call binding the contract method 0xaf958a62.
//
// Solidity: function AllIncomeOfAllKol(address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) AllIncomeOfAllKol(investor common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.AllIncomeOfAllKol(&_KolKeySimple.CallOpts, investor)
}

// AllIncomeOfOneKol is a free data retrieval call binding the contract method 0x6d74bb5f.
//
// Solidity: function AllIncomeOfOneKol(address kol, address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) AllIncomeOfOneKol(opts *bind.CallOpts, kol common.Address, investor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "AllIncomeOfOneKol", kol, investor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllIncomeOfOneKol is a free data retrieval call binding the contract method 0x6d74bb5f.
//
// Solidity: function AllIncomeOfOneKol(address kol, address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) AllIncomeOfOneKol(kol common.Address, investor common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.AllIncomeOfOneKol(&_KolKeySimple.CallOpts, kol, investor)
}

// AllIncomeOfOneKol is a free data retrieval call binding the contract method 0x6d74bb5f.
//
// Solidity: function AllIncomeOfOneKol(address kol, address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) AllIncomeOfOneKol(kol common.Address, investor common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.AllIncomeOfOneKol(&_KolKeySimple.CallOpts, kol, investor)
}

// AllKolAddr is a free data retrieval call binding the contract method 0x399ccecf.
//
// Solidity: function AllKolAddr() view returns(address[])
func (_KolKeySimple *KolKeySimpleCaller) AllKolAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "AllKolAddr")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllKolAddr is a free data retrieval call binding the contract method 0x399ccecf.
//
// Solidity: function AllKolAddr() view returns(address[])
func (_KolKeySimple *KolKeySimpleSession) AllKolAddr() ([]common.Address, error) {
	return _KolKeySimple.Contract.AllKolAddr(&_KolKeySimple.CallOpts)
}

// AllKolAddr is a free data retrieval call binding the contract method 0x399ccecf.
//
// Solidity: function AllKolAddr() view returns(address[])
func (_KolKeySimple *KolKeySimpleCallerSession) AllKolAddr() ([]common.Address, error) {
	return _KolKeySimple.Contract.AllKolAddr(&_KolKeySimple.CallOpts)
}

// IncomeOfOneNonce is a free data retrieval call binding the contract method 0x3a99f6a0.
//
// Solidity: function IncomeOfOneNonce(address kol, uint256 nonce, address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) IncomeOfOneNonce(opts *bind.CallOpts, kol common.Address, nonce *big.Int, investor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "IncomeOfOneNonce", kol, nonce, investor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomeOfOneNonce is a free data retrieval call binding the contract method 0x3a99f6a0.
//
// Solidity: function IncomeOfOneNonce(address kol, uint256 nonce, address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) IncomeOfOneNonce(kol common.Address, nonce *big.Int, investor common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.IncomeOfOneNonce(&_KolKeySimple.CallOpts, kol, nonce, investor)
}

// IncomeOfOneNonce is a free data retrieval call binding the contract method 0x3a99f6a0.
//
// Solidity: function IncomeOfOneNonce(address kol, uint256 nonce, address investor) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) IncomeOfOneNonce(kol common.Address, nonce *big.Int, investor common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.IncomeOfOneNonce(&_KolKeySimple.CallOpts, kol, nonce, investor)
}

// IncomeOfOneNonceByAmount is a free data retrieval call binding the contract method 0x232d562f.
//
// Solidity: function IncomeOfOneNonceByAmount(address kol, uint256 nonce, uint256 amount) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) IncomeOfOneNonceByAmount(opts *bind.CallOpts, kol common.Address, nonce *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "IncomeOfOneNonceByAmount", kol, nonce, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomeOfOneNonceByAmount is a free data retrieval call binding the contract method 0x232d562f.
//
// Solidity: function IncomeOfOneNonceByAmount(address kol, uint256 nonce, uint256 amount) view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) IncomeOfOneNonceByAmount(kol common.Address, nonce *big.Int, amount *big.Int) (*big.Int, error) {
	return _KolKeySimple.Contract.IncomeOfOneNonceByAmount(&_KolKeySimple.CallOpts, kol, nonce, amount)
}

// IncomeOfOneNonceByAmount is a free data retrieval call binding the contract method 0x232d562f.
//
// Solidity: function IncomeOfOneNonceByAmount(address kol, uint256 nonce, uint256 amount) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) IncomeOfOneNonceByAmount(kol common.Address, nonce *big.Int, amount *big.Int) (*big.Int, error) {
	return _KolKeySimple.Contract.IncomeOfOneNonceByAmount(&_KolKeySimple.CallOpts, kol, nonce, amount)
}

// InvestorAllKeysOfKol is a free data retrieval call binding the contract method 0xb83bb4e6.
//
// Solidity: function InvestorAllKeysOfKol(address investor, address kol) view returns(uint256[] nonce, uint256[] amounts)
func (_KolKeySimple *KolKeySimpleCaller) InvestorAllKeysOfKol(opts *bind.CallOpts, investor common.Address, kol common.Address) (struct {
	Nonce   []*big.Int
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "InvestorAllKeysOfKol", investor, kol)

	outstruct := new(struct {
		Nonce   []*big.Int
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// InvestorAllKeysOfKol is a free data retrieval call binding the contract method 0xb83bb4e6.
//
// Solidity: function InvestorAllKeysOfKol(address investor, address kol) view returns(uint256[] nonce, uint256[] amounts)
func (_KolKeySimple *KolKeySimpleSession) InvestorAllKeysOfKol(investor common.Address, kol common.Address) (struct {
	Nonce   []*big.Int
	Amounts []*big.Int
}, error) {
	return _KolKeySimple.Contract.InvestorAllKeysOfKol(&_KolKeySimple.CallOpts, investor, kol)
}

// InvestorAllKeysOfKol is a free data retrieval call binding the contract method 0xb83bb4e6.
//
// Solidity: function InvestorAllKeysOfKol(address investor, address kol) view returns(uint256[] nonce, uint256[] amounts)
func (_KolKeySimple *KolKeySimpleCallerSession) InvestorAllKeysOfKol(investor common.Address, kol common.Address) (struct {
	Nonce   []*big.Int
	Amounts []*big.Int
}, error) {
	return _KolKeySimple.Contract.InvestorAllKeysOfKol(&_KolKeySimple.CallOpts, investor, kol)
}

// InvestorOfKol is a free data retrieval call binding the contract method 0x8b2b0b67.
//
// Solidity: function InvestorOfKol(address kol) view returns(address[])
func (_KolKeySimple *KolKeySimpleCaller) InvestorOfKol(opts *bind.CallOpts, kol common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "InvestorOfKol", kol)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// InvestorOfKol is a free data retrieval call binding the contract method 0x8b2b0b67.
//
// Solidity: function InvestorOfKol(address kol) view returns(address[])
func (_KolKeySimple *KolKeySimpleSession) InvestorOfKol(kol common.Address) ([]common.Address, error) {
	return _KolKeySimple.Contract.InvestorOfKol(&_KolKeySimple.CallOpts, kol)
}

// InvestorOfKol is a free data retrieval call binding the contract method 0x8b2b0b67.
//
// Solidity: function InvestorOfKol(address kol) view returns(address[])
func (_KolKeySimple *KolKeySimpleCallerSession) InvestorOfKol(kol common.Address) ([]common.Address, error) {
	return _KolKeySimple.Contract.InvestorOfKol(&_KolKeySimple.CallOpts, kol)
}

// KeySettingsRecord is a free data retrieval call binding the contract method 0x94447f45.
//
// Solidity: function KeySettingsRecord(address ) view returns(uint256 price, uint256 nonce, uint256 totalVal, uint256 totalNo)
func (_KolKeySimple *KolKeySimpleCaller) KeySettingsRecord(opts *bind.CallOpts, arg0 common.Address) (struct {
	Price    *big.Int
	Nonce    *big.Int
	TotalVal *big.Int
	TotalNo  *big.Int
}, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "KeySettingsRecord", arg0)

	outstruct := new(struct {
		Price    *big.Int
		Nonce    *big.Int
		TotalVal *big.Int
		TotalNo  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalVal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalNo = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeySettingsRecord is a free data retrieval call binding the contract method 0x94447f45.
//
// Solidity: function KeySettingsRecord(address ) view returns(uint256 price, uint256 nonce, uint256 totalVal, uint256 totalNo)
func (_KolKeySimple *KolKeySimpleSession) KeySettingsRecord(arg0 common.Address) (struct {
	Price    *big.Int
	Nonce    *big.Int
	TotalVal *big.Int
	TotalNo  *big.Int
}, error) {
	return _KolKeySimple.Contract.KeySettingsRecord(&_KolKeySimple.CallOpts, arg0)
}

// KeySettingsRecord is a free data retrieval call binding the contract method 0x94447f45.
//
// Solidity: function KeySettingsRecord(address ) view returns(uint256 price, uint256 nonce, uint256 totalVal, uint256 totalNo)
func (_KolKeySimple *KolKeySimpleCallerSession) KeySettingsRecord(arg0 common.Address) (struct {
	Price    *big.Int
	Nonce    *big.Int
	TotalVal *big.Int
	TotalNo  *big.Int
}, error) {
	return _KolKeySimple.Contract.KeySettingsRecord(&_KolKeySimple.CallOpts, arg0)
}

// KolOfOneInvestor is a free data retrieval call binding the contract method 0xce92faa2.
//
// Solidity: function KolOfOneInvestor(address investor) view returns(address[])
func (_KolKeySimple *KolKeySimpleCaller) KolOfOneInvestor(opts *bind.CallOpts, investor common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "KolOfOneInvestor", investor)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// KolOfOneInvestor is a free data retrieval call binding the contract method 0xce92faa2.
//
// Solidity: function KolOfOneInvestor(address investor) view returns(address[])
func (_KolKeySimple *KolKeySimpleSession) KolOfOneInvestor(investor common.Address) ([]common.Address, error) {
	return _KolKeySimple.Contract.KolOfOneInvestor(&_KolKeySimple.CallOpts, investor)
}

// KolOfOneInvestor is a free data retrieval call binding the contract method 0xce92faa2.
//
// Solidity: function KolOfOneInvestor(address investor) view returns(address[])
func (_KolKeySimple *KolKeySimpleCallerSession) KolOfOneInvestor(investor common.Address) ([]common.Address, error) {
	return _KolKeySimple.Contract.KolOfOneInvestor(&_KolKeySimple.CallOpts, investor)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_KolKeySimple *KolKeySimpleCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "__admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_KolKeySimple *KolKeySimpleSession) Admins(arg0 common.Address) (bool, error) {
	return _KolKeySimple.Contract.Admins(&_KolKeySimple.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_KolKeySimple *KolKeySimpleCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _KolKeySimple.Contract.Admins(&_KolKeySimple.CallOpts, arg0)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) MinValCheck(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "__minValCheck")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) MinValCheck() (*big.Int, error) {
	return _KolKeySimple.Contract.MinValCheck(&_KolKeySimple.CallOpts)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) MinValCheck() (*big.Int, error) {
	return _KolKeySimple.Contract.MinValCheck(&_KolKeySimple.CallOpts)
}

// AllKolInSystem is a free data retrieval call binding the contract method 0x2d5e031d.
//
// Solidity: function allKolInSystem(uint256 ) view returns(address)
func (_KolKeySimple *KolKeySimpleCaller) AllKolInSystem(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "allKolInSystem", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllKolInSystem is a free data retrieval call binding the contract method 0x2d5e031d.
//
// Solidity: function allKolInSystem(uint256 ) view returns(address)
func (_KolKeySimple *KolKeySimpleSession) AllKolInSystem(arg0 *big.Int) (common.Address, error) {
	return _KolKeySimple.Contract.AllKolInSystem(&_KolKeySimple.CallOpts, arg0)
}

// AllKolInSystem is a free data retrieval call binding the contract method 0x2d5e031d.
//
// Solidity: function allKolInSystem(uint256 ) view returns(address)
func (_KolKeySimple *KolKeySimpleCallerSession) AllKolInSystem(arg0 *big.Int) (common.Address, error) {
	return _KolKeySimple.Contract.AllKolInSystem(&_KolKeySimple.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.Balance(&_KolKeySimple.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _KolKeySimple.Contract.Balance(&_KolKeySimple.CallOpts, arg0)
}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_KolKeySimple *KolKeySimpleCaller) CheckPluginInterface(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "checkPluginInterface")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_KolKeySimple *KolKeySimpleSession) CheckPluginInterface() (bool, error) {
	return _KolKeySimple.Contract.CheckPluginInterface(&_KolKeySimple.CallOpts)
}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_KolKeySimple *KolKeySimpleCallerSession) CheckPluginInterface() (bool, error) {
	return _KolKeySimple.Contract.CheckPluginInterface(&_KolKeySimple.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KolKeySimple *KolKeySimpleCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KolKeySimple *KolKeySimpleSession) GetOwner() (common.Address, error) {
	return _KolKeySimple.Contract.GetOwner(&_KolKeySimple.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KolKeySimple *KolKeySimpleCallerSession) GetOwner() (common.Address, error) {
	return _KolKeySimple.Contract.GetOwner(&_KolKeySimple.CallOpts)
}

// IncomePerNoncePerKey is a free data retrieval call binding the contract method 0x53311742.
//
// Solidity: function incomePerNoncePerKey(address , uint256 ) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) IncomePerNoncePerKey(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "incomePerNoncePerKey", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomePerNoncePerKey is a free data retrieval call binding the contract method 0x53311742.
//
// Solidity: function incomePerNoncePerKey(address , uint256 ) view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) IncomePerNoncePerKey(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _KolKeySimple.Contract.IncomePerNoncePerKey(&_KolKeySimple.CallOpts, arg0, arg1)
}

// IncomePerNoncePerKey is a free data retrieval call binding the contract method 0x53311742.
//
// Solidity: function incomePerNoncePerKey(address , uint256 ) view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) IncomePerNoncePerKey(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _KolKeySimple.Contract.IncomePerNoncePerKey(&_KolKeySimple.CallOpts, arg0, arg1)
}

// KolIncomeRatePerKeyBuy is a free data retrieval call binding the contract method 0xafb337cf.
//
// Solidity: function kolIncomeRatePerKeyBuy() view returns(uint8)
func (_KolKeySimple *KolKeySimpleCaller) KolIncomeRatePerKeyBuy(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "kolIncomeRatePerKeyBuy")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KolIncomeRatePerKeyBuy is a free data retrieval call binding the contract method 0xafb337cf.
//
// Solidity: function kolIncomeRatePerKeyBuy() view returns(uint8)
func (_KolKeySimple *KolKeySimpleSession) KolIncomeRatePerKeyBuy() (uint8, error) {
	return _KolKeySimple.Contract.KolIncomeRatePerKeyBuy(&_KolKeySimple.CallOpts)
}

// KolIncomeRatePerKeyBuy is a free data retrieval call binding the contract method 0xafb337cf.
//
// Solidity: function kolIncomeRatePerKeyBuy() view returns(uint8)
func (_KolKeySimple *KolKeySimpleCallerSession) KolIncomeRatePerKeyBuy() (uint8, error) {
	return _KolKeySimple.Contract.KolIncomeRatePerKeyBuy(&_KolKeySimple.CallOpts)
}

// KolOpenKeyPool is a free data retrieval call binding the contract method 0xbb86a929.
//
// Solidity: function kolOpenKeyPool(address sourceAddr) view returns(bool)
func (_KolKeySimple *KolKeySimpleCaller) KolOpenKeyPool(opts *bind.CallOpts, sourceAddr common.Address) (bool, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "kolOpenKeyPool", sourceAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KolOpenKeyPool is a free data retrieval call binding the contract method 0xbb86a929.
//
// Solidity: function kolOpenKeyPool(address sourceAddr) view returns(bool)
func (_KolKeySimple *KolKeySimpleSession) KolOpenKeyPool(sourceAddr common.Address) (bool, error) {
	return _KolKeySimple.Contract.KolOpenKeyPool(&_KolKeySimple.CallOpts, sourceAddr)
}

// KolOpenKeyPool is a free data retrieval call binding the contract method 0xbb86a929.
//
// Solidity: function kolOpenKeyPool(address sourceAddr) view returns(bool)
func (_KolKeySimple *KolKeySimpleCallerSession) KolOpenKeyPool(sourceAddr common.Address) (bool, error) {
	return _KolKeySimple.Contract.KolOpenKeyPool(&_KolKeySimple.CallOpts, sourceAddr)
}

// MaxKeyNoForKol is a free data retrieval call binding the contract method 0xd5238a8c.
//
// Solidity: function maxKeyNoForKol() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) MaxKeyNoForKol(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "maxKeyNoForKol")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxKeyNoForKol is a free data retrieval call binding the contract method 0xd5238a8c.
//
// Solidity: function maxKeyNoForKol() view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) MaxKeyNoForKol() (*big.Int, error) {
	return _KolKeySimple.Contract.MaxKeyNoForKol(&_KolKeySimple.CallOpts)
}

// MaxKeyNoForKol is a free data retrieval call binding the contract method 0xd5238a8c.
//
// Solidity: function maxKeyNoForKol() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) MaxKeyNoForKol() (*big.Int, error) {
	return _KolKeySimple.Contract.MaxKeyNoForKol(&_KolKeySimple.CallOpts)
}

// ServiceFeeRatePerKeyBuy is a free data retrieval call binding the contract method 0xaf707457.
//
// Solidity: function serviceFeeRatePerKeyBuy() view returns(uint8)
func (_KolKeySimple *KolKeySimpleCaller) ServiceFeeRatePerKeyBuy(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "serviceFeeRatePerKeyBuy")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ServiceFeeRatePerKeyBuy is a free data retrieval call binding the contract method 0xaf707457.
//
// Solidity: function serviceFeeRatePerKeyBuy() view returns(uint8)
func (_KolKeySimple *KolKeySimpleSession) ServiceFeeRatePerKeyBuy() (uint8, error) {
	return _KolKeySimple.Contract.ServiceFeeRatePerKeyBuy(&_KolKeySimple.CallOpts)
}

// ServiceFeeRatePerKeyBuy is a free data retrieval call binding the contract method 0xaf707457.
//
// Solidity: function serviceFeeRatePerKeyBuy() view returns(uint8)
func (_KolKeySimple *KolKeySimpleCallerSession) ServiceFeeRatePerKeyBuy() (uint8, error) {
	return _KolKeySimple.Contract.ServiceFeeRatePerKeyBuy(&_KolKeySimple.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) ServiceFeeReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "serviceFeeReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) ServiceFeeReceived() (*big.Int, error) {
	return _KolKeySimple.Contract.ServiceFeeReceived(&_KolKeySimple.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) ServiceFeeReceived() (*big.Int, error) {
	return _KolKeySimple.Contract.ServiceFeeReceived(&_KolKeySimple.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCaller) WithdrawFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeySimple.contract.Call(opts, &out, "withdrawFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_KolKeySimple *KolKeySimpleSession) WithdrawFeeRate() (*big.Int, error) {
	return _KolKeySimple.Contract.WithdrawFeeRate(&_KolKeySimple.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_KolKeySimple *KolKeySimpleCallerSession) WithdrawFeeRate() (*big.Int, error) {
	return _KolKeySimple.Contract.WithdrawFeeRate(&_KolKeySimple.CallOpts)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminOperation(opts *bind.TransactOpts, admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminOperation", admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_KolKeySimple *KolKeySimpleSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminOperation(&_KolKeySimple.TransactOpts, admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminOperation(&_KolKeySimple.TransactOpts, admin, isDelete)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminServiceFeeWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminServiceFeeWithdraw")
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_KolKeySimple *KolKeySimpleSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminServiceFeeWithdraw(&_KolKeySimple.TransactOpts)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminServiceFeeWithdraw(&_KolKeySimple.TransactOpts)
}

// AdminSetKeyFeeRate is a paid mutator transaction binding the contract method 0x84854688.
//
// Solidity: function adminSetKeyFeeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminSetKeyFeeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminSetKeyFeeRate", newRate)
}

// AdminSetKeyFeeRate is a paid mutator transaction binding the contract method 0x84854688.
//
// Solidity: function adminSetKeyFeeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleSession) AdminSetKeyFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetKeyFeeRate(&_KolKeySimple.TransactOpts, newRate)
}

// AdminSetKeyFeeRate is a paid mutator transaction binding the contract method 0x84854688.
//
// Solidity: function adminSetKeyFeeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminSetKeyFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetKeyFeeRate(&_KolKeySimple.TransactOpts, newRate)
}

// AdminSetKolIncomeRate is a paid mutator transaction binding the contract method 0x34afeec0.
//
// Solidity: function adminSetKolIncomeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminSetKolIncomeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminSetKolIncomeRate", newRate)
}

// AdminSetKolIncomeRate is a paid mutator transaction binding the contract method 0x34afeec0.
//
// Solidity: function adminSetKolIncomeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleSession) AdminSetKolIncomeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetKolIncomeRate(&_KolKeySimple.TransactOpts, newRate)
}

// AdminSetKolIncomeRate is a paid mutator transaction binding the contract method 0x34afeec0.
//
// Solidity: function adminSetKolIncomeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminSetKolIncomeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetKolIncomeRate(&_KolKeySimple.TransactOpts, newRate)
}

// AdminSetMaxKolKeyNo is a paid mutator transaction binding the contract method 0x2d7cb2ee.
//
// Solidity: function adminSetMaxKolKeyNo(uint256 keyNo) returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminSetMaxKolKeyNo(opts *bind.TransactOpts, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminSetMaxKolKeyNo", keyNo)
}

// AdminSetMaxKolKeyNo is a paid mutator transaction binding the contract method 0x2d7cb2ee.
//
// Solidity: function adminSetMaxKolKeyNo(uint256 keyNo) returns()
func (_KolKeySimple *KolKeySimpleSession) AdminSetMaxKolKeyNo(keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetMaxKolKeyNo(&_KolKeySimple.TransactOpts, keyNo)
}

// AdminSetMaxKolKeyNo is a paid mutator transaction binding the contract method 0x2d7cb2ee.
//
// Solidity: function adminSetMaxKolKeyNo(uint256 keyNo) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminSetMaxKolKeyNo(keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetMaxKolKeyNo(&_KolKeySimple.TransactOpts, keyNo)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminSetWithdrawFeeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminSetWithdrawFeeRate", newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetWithdrawFeeRate(&_KolKeySimple.TransactOpts, newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminSetWithdrawFeeRate(&_KolKeySimple.TransactOpts, newRate)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_KolKeySimple *KolKeySimpleTransactor) AdminUpgradeToNewRule(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "adminUpgradeToNewRule", recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_KolKeySimple *KolKeySimpleSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminUpgradeToNewRule(&_KolKeySimple.TransactOpts, recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.AdminUpgradeToNewRule(&_KolKeySimple.TransactOpts, recipient)
}

// BuyKolKey is a paid mutator transaction binding the contract method 0xc7cdf93c.
//
// Solidity: function buyKolKey(address kolAddr, uint256 keyNo) payable returns()
func (_KolKeySimple *KolKeySimpleTransactor) BuyKolKey(opts *bind.TransactOpts, kolAddr common.Address, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "buyKolKey", kolAddr, keyNo)
}

// BuyKolKey is a paid mutator transaction binding the contract method 0xc7cdf93c.
//
// Solidity: function buyKolKey(address kolAddr, uint256 keyNo) payable returns()
func (_KolKeySimple *KolKeySimpleSession) BuyKolKey(kolAddr common.Address, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.BuyKolKey(&_KolKeySimple.TransactOpts, kolAddr, keyNo)
}

// BuyKolKey is a paid mutator transaction binding the contract method 0xc7cdf93c.
//
// Solidity: function buyKolKey(address kolAddr, uint256 keyNo) payable returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) BuyKolKey(kolAddr common.Address, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.BuyKolKey(&_KolKeySimple.TransactOpts, kolAddr, keyNo)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_KolKeySimple *KolKeySimpleTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_KolKeySimple *KolKeySimpleSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.ChangeOwner(&_KolKeySimple.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.ChangeOwner(&_KolKeySimple.TransactOpts, newOwner)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_KolKeySimple *KolKeySimpleTransactor) ChangeStatus(opts *bind.TransactOpts, stop bool) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "changeStatus", stop)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_KolKeySimple *KolKeySimpleSession) ChangeStatus(stop bool) (*types.Transaction, error) {
	return _KolKeySimple.Contract.ChangeStatus(&_KolKeySimple.TransactOpts, stop)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) ChangeStatus(stop bool) (*types.Transaction, error) {
	return _KolKeySimple.Contract.ChangeStatus(&_KolKeySimple.TransactOpts, stop)
}

// KolGotIncome is a paid mutator transaction binding the contract method 0xd842f385.
//
// Solidity: function kolGotIncome(int8 sourceID, address kolAddr) payable returns()
func (_KolKeySimple *KolKeySimpleTransactor) KolGotIncome(opts *bind.TransactOpts, sourceID int8, kolAddr common.Address) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "kolGotIncome", sourceID, kolAddr)
}

// KolGotIncome is a paid mutator transaction binding the contract method 0xd842f385.
//
// Solidity: function kolGotIncome(int8 sourceID, address kolAddr) payable returns()
func (_KolKeySimple *KolKeySimpleSession) KolGotIncome(sourceID int8, kolAddr common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.KolGotIncome(&_KolKeySimple.TransactOpts, sourceID, kolAddr)
}

// KolGotIncome is a paid mutator transaction binding the contract method 0xd842f385.
//
// Solidity: function kolGotIncome(int8 sourceID, address kolAddr) payable returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) KolGotIncome(sourceID int8, kolAddr common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.KolGotIncome(&_KolKeySimple.TransactOpts, sourceID, kolAddr)
}

// KolOpenKeySale is a paid mutator transaction binding the contract method 0x5c1ec100.
//
// Solidity: function kolOpenKeySale(uint256 priceInFin) returns()
func (_KolKeySimple *KolKeySimpleTransactor) KolOpenKeySale(opts *bind.TransactOpts, priceInFin *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "kolOpenKeySale", priceInFin)
}

// KolOpenKeySale is a paid mutator transaction binding the contract method 0x5c1ec100.
//
// Solidity: function kolOpenKeySale(uint256 priceInFin) returns()
func (_KolKeySimple *KolKeySimpleSession) KolOpenKeySale(priceInFin *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.KolOpenKeySale(&_KolKeySimple.TransactOpts, priceInFin)
}

// KolOpenKeySale is a paid mutator transaction binding the contract method 0x5c1ec100.
//
// Solidity: function kolOpenKeySale(uint256 priceInFin) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) KolOpenKeySale(priceInFin *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.KolOpenKeySale(&_KolKeySimple.TransactOpts, priceInFin)
}

// RebindKolKey is a paid mutator transaction binding the contract method 0x6390f644.
//
// Solidity: function rebindKolKey(address to) returns()
func (_KolKeySimple *KolKeySimpleTransactor) RebindKolKey(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "rebindKolKey", to)
}

// RebindKolKey is a paid mutator transaction binding the contract method 0x6390f644.
//
// Solidity: function rebindKolKey(address to) returns()
func (_KolKeySimple *KolKeySimpleSession) RebindKolKey(to common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.RebindKolKey(&_KolKeySimple.TransactOpts, to)
}

// RebindKolKey is a paid mutator transaction binding the contract method 0x6390f644.
//
// Solidity: function rebindKolKey(address to) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) RebindKolKey(to common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.RebindKolKey(&_KolKeySimple.TransactOpts, to)
}

// TransferAllKey is a paid mutator transaction binding the contract method 0x01c6c8da.
//
// Solidity: function transferAllKey(address to, address kol) returns()
func (_KolKeySimple *KolKeySimpleTransactor) TransferAllKey(opts *bind.TransactOpts, to common.Address, kol common.Address) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "transferAllKey", to, kol)
}

// TransferAllKey is a paid mutator transaction binding the contract method 0x01c6c8da.
//
// Solidity: function transferAllKey(address to, address kol) returns()
func (_KolKeySimple *KolKeySimpleSession) TransferAllKey(to common.Address, kol common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.TransferAllKey(&_KolKeySimple.TransactOpts, to, kol)
}

// TransferAllKey is a paid mutator transaction binding the contract method 0x01c6c8da.
//
// Solidity: function transferAllKey(address to, address kol) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) TransferAllKey(to common.Address, kol common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.TransferAllKey(&_KolKeySimple.TransactOpts, to, kol)
}

// TransferKey is a paid mutator transaction binding the contract method 0x7b46bfe8.
//
// Solidity: function transferKey(address to, address kol, uint256 nonce, uint256 amount) returns()
func (_KolKeySimple *KolKeySimpleTransactor) TransferKey(opts *bind.TransactOpts, to common.Address, kol common.Address, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "transferKey", to, kol, nonce, amount)
}

// TransferKey is a paid mutator transaction binding the contract method 0x7b46bfe8.
//
// Solidity: function transferKey(address to, address kol, uint256 nonce, uint256 amount) returns()
func (_KolKeySimple *KolKeySimpleSession) TransferKey(to common.Address, kol common.Address, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.TransferKey(&_KolKeySimple.TransactOpts, to, kol, nonce, amount)
}

// TransferKey is a paid mutator transaction binding the contract method 0x7b46bfe8.
//
// Solidity: function transferKey(address to, address kol, uint256 nonce, uint256 amount) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) TransferKey(to common.Address, kol common.Address, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.TransferKey(&_KolKeySimple.TransactOpts, to, kol, nonce, amount)
}

// WithdrawAllIncome is a paid mutator transaction binding the contract method 0x0c3e0452.
//
// Solidity: function withdrawAllIncome() returns()
func (_KolKeySimple *KolKeySimpleTransactor) WithdrawAllIncome(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "withdrawAllIncome")
}

// WithdrawAllIncome is a paid mutator transaction binding the contract method 0x0c3e0452.
//
// Solidity: function withdrawAllIncome() returns()
func (_KolKeySimple *KolKeySimpleSession) WithdrawAllIncome() (*types.Transaction, error) {
	return _KolKeySimple.Contract.WithdrawAllIncome(&_KolKeySimple.TransactOpts)
}

// WithdrawAllIncome is a paid mutator transaction binding the contract method 0x0c3e0452.
//
// Solidity: function withdrawAllIncome() returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) WithdrawAllIncome() (*types.Transaction, error) {
	return _KolKeySimple.Contract.WithdrawAllIncome(&_KolKeySimple.TransactOpts)
}

// WithdrawFromOneKol is a paid mutator transaction binding the contract method 0x0ef86b3f.
//
// Solidity: function withdrawFromOneKol(address kol) returns()
func (_KolKeySimple *KolKeySimpleTransactor) WithdrawFromOneKol(opts *bind.TransactOpts, kol common.Address) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "withdrawFromOneKol", kol)
}

// WithdrawFromOneKol is a paid mutator transaction binding the contract method 0x0ef86b3f.
//
// Solidity: function withdrawFromOneKol(address kol) returns()
func (_KolKeySimple *KolKeySimpleSession) WithdrawFromOneKol(kol common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.WithdrawFromOneKol(&_KolKeySimple.TransactOpts, kol)
}

// WithdrawFromOneKol is a paid mutator transaction binding the contract method 0x0ef86b3f.
//
// Solidity: function withdrawFromOneKol(address kol) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) WithdrawFromOneKol(kol common.Address) (*types.Transaction, error) {
	return _KolKeySimple.Contract.WithdrawFromOneKol(&_KolKeySimple.TransactOpts, kol)
}

// WithdrawFromOneNonce is a paid mutator transaction binding the contract method 0x3da6136e.
//
// Solidity: function withdrawFromOneNonce(address kol, uint256 nonce) returns()
func (_KolKeySimple *KolKeySimpleTransactor) WithdrawFromOneNonce(opts *bind.TransactOpts, kol common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.contract.Transact(opts, "withdrawFromOneNonce", kol, nonce)
}

// WithdrawFromOneNonce is a paid mutator transaction binding the contract method 0x3da6136e.
//
// Solidity: function withdrawFromOneNonce(address kol, uint256 nonce) returns()
func (_KolKeySimple *KolKeySimpleSession) WithdrawFromOneNonce(kol common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.WithdrawFromOneNonce(&_KolKeySimple.TransactOpts, kol, nonce)
}

// WithdrawFromOneNonce is a paid mutator transaction binding the contract method 0x3da6136e.
//
// Solidity: function withdrawFromOneNonce(address kol, uint256 nonce) returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) WithdrawFromOneNonce(kol common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _KolKeySimple.Contract.WithdrawFromOneNonce(&_KolKeySimple.TransactOpts, kol, nonce)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KolKeySimple *KolKeySimpleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeySimple.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KolKeySimple *KolKeySimpleSession) Receive() (*types.Transaction, error) {
	return _KolKeySimple.Contract.Receive(&_KolKeySimple.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KolKeySimple *KolKeySimpleTransactorSession) Receive() (*types.Transaction, error) {
	return _KolKeySimple.Contract.Receive(&_KolKeySimple.TransactOpts)
}

// KolKeySimpleAdminOperationIterator is returned from FilterAdminOperation and is used to iterate over the raw logs and unpacked data for AdminOperation events raised by the KolKeySimple contract.
type KolKeySimpleAdminOperationIterator struct {
	Event *KolKeySimpleAdminOperation // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleAdminOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleAdminOperation)
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
		it.Event = new(KolKeySimpleAdminOperation)
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
func (it *KolKeySimpleAdminOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleAdminOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleAdminOperation represents a AdminOperation event raised by the KolKeySimple contract.
type KolKeySimpleAdminOperation struct {
	Admin  common.Address
	OpType bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAdminOperation is a free log retrieval operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_KolKeySimple *KolKeySimpleFilterer) FilterAdminOperation(opts *bind.FilterOpts) (*KolKeySimpleAdminOperationIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleAdminOperationIterator{contract: _KolKeySimple.contract, event: "AdminOperation", logs: logs, sub: sub}, nil
}

// WatchAdminOperation is a free log subscription operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_KolKeySimple *KolKeySimpleFilterer) WatchAdminOperation(opts *bind.WatchOpts, sink chan<- *KolKeySimpleAdminOperation) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleAdminOperation)
				if err := _KolKeySimple.contract.UnpackLog(event, "AdminOperation", log); err != nil {
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

// ParseAdminOperation is a log parse operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_KolKeySimple *KolKeySimpleFilterer) ParseAdminOperation(log types.Log) (*KolKeySimpleAdminOperation, error) {
	event := new(KolKeySimpleAdminOperation)
	if err := _KolKeySimple.contract.UnpackLog(event, "AdminOperation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleInvestorWithdrawAllIncomeIterator is returned from FilterInvestorWithdrawAllIncome and is used to iterate over the raw logs and unpacked data for InvestorWithdrawAllIncome events raised by the KolKeySimple contract.
type KolKeySimpleInvestorWithdrawAllIncomeIterator struct {
	Event *KolKeySimpleInvestorWithdrawAllIncome // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleInvestorWithdrawAllIncomeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleInvestorWithdrawAllIncome)
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
		it.Event = new(KolKeySimpleInvestorWithdrawAllIncome)
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
func (it *KolKeySimpleInvestorWithdrawAllIncomeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleInvestorWithdrawAllIncomeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleInvestorWithdrawAllIncome represents a InvestorWithdrawAllIncome event raised by the KolKeySimple contract.
type KolKeySimpleInvestorWithdrawAllIncome struct {
	Investor common.Address
	Val      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdrawAllIncome is a free log retrieval operation binding the contract event 0x5dbe3bbfa87813b8be328cdb0b812e04de36cf36d1141967e1df795a7c95db0a.
//
// Solidity: event InvestorWithdrawAllIncome(address investor, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) FilterInvestorWithdrawAllIncome(opts *bind.FilterOpts) (*KolKeySimpleInvestorWithdrawAllIncomeIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "InvestorWithdrawAllIncome")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleInvestorWithdrawAllIncomeIterator{contract: _KolKeySimple.contract, event: "InvestorWithdrawAllIncome", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdrawAllIncome is a free log subscription operation binding the contract event 0x5dbe3bbfa87813b8be328cdb0b812e04de36cf36d1141967e1df795a7c95db0a.
//
// Solidity: event InvestorWithdrawAllIncome(address investor, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) WatchInvestorWithdrawAllIncome(opts *bind.WatchOpts, sink chan<- *KolKeySimpleInvestorWithdrawAllIncome) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "InvestorWithdrawAllIncome")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleInvestorWithdrawAllIncome)
				if err := _KolKeySimple.contract.UnpackLog(event, "InvestorWithdrawAllIncome", log); err != nil {
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

// ParseInvestorWithdrawAllIncome is a log parse operation binding the contract event 0x5dbe3bbfa87813b8be328cdb0b812e04de36cf36d1141967e1df795a7c95db0a.
//
// Solidity: event InvestorWithdrawAllIncome(address investor, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) ParseInvestorWithdrawAllIncome(log types.Log) (*KolKeySimpleInvestorWithdrawAllIncome, error) {
	event := new(KolKeySimpleInvestorWithdrawAllIncome)
	if err := _KolKeySimple.contract.UnpackLog(event, "InvestorWithdrawAllIncome", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleInvestorWithdrawByOneKolIterator is returned from FilterInvestorWithdrawByOneKol and is used to iterate over the raw logs and unpacked data for InvestorWithdrawByOneKol events raised by the KolKeySimple contract.
type KolKeySimpleInvestorWithdrawByOneKolIterator struct {
	Event *KolKeySimpleInvestorWithdrawByOneKol // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleInvestorWithdrawByOneKolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleInvestorWithdrawByOneKol)
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
		it.Event = new(KolKeySimpleInvestorWithdrawByOneKol)
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
func (it *KolKeySimpleInvestorWithdrawByOneKolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleInvestorWithdrawByOneKolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleInvestorWithdrawByOneKol represents a InvestorWithdrawByOneKol event raised by the KolKeySimple contract.
type KolKeySimpleInvestorWithdrawByOneKol struct {
	Investor common.Address
	Kol      common.Address
	Val      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdrawByOneKol is a free log retrieval operation binding the contract event 0x6ce7ea243972341c9e8ddfa93ba573fec9d104862f00b1bcd6aab059efbc9a6d.
//
// Solidity: event InvestorWithdrawByOneKol(address investor, address kol, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) FilterInvestorWithdrawByOneKol(opts *bind.FilterOpts) (*KolKeySimpleInvestorWithdrawByOneKolIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "InvestorWithdrawByOneKol")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleInvestorWithdrawByOneKolIterator{contract: _KolKeySimple.contract, event: "InvestorWithdrawByOneKol", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdrawByOneKol is a free log subscription operation binding the contract event 0x6ce7ea243972341c9e8ddfa93ba573fec9d104862f00b1bcd6aab059efbc9a6d.
//
// Solidity: event InvestorWithdrawByOneKol(address investor, address kol, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) WatchInvestorWithdrawByOneKol(opts *bind.WatchOpts, sink chan<- *KolKeySimpleInvestorWithdrawByOneKol) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "InvestorWithdrawByOneKol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleInvestorWithdrawByOneKol)
				if err := _KolKeySimple.contract.UnpackLog(event, "InvestorWithdrawByOneKol", log); err != nil {
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

// ParseInvestorWithdrawByOneKol is a log parse operation binding the contract event 0x6ce7ea243972341c9e8ddfa93ba573fec9d104862f00b1bcd6aab059efbc9a6d.
//
// Solidity: event InvestorWithdrawByOneKol(address investor, address kol, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) ParseInvestorWithdrawByOneKol(log types.Log) (*KolKeySimpleInvestorWithdrawByOneKol, error) {
	event := new(KolKeySimpleInvestorWithdrawByOneKol)
	if err := _KolKeySimple.contract.UnpackLog(event, "InvestorWithdrawByOneKol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleInvestorWithdrawByOneNonceIterator is returned from FilterInvestorWithdrawByOneNonce and is used to iterate over the raw logs and unpacked data for InvestorWithdrawByOneNonce events raised by the KolKeySimple contract.
type KolKeySimpleInvestorWithdrawByOneNonceIterator struct {
	Event *KolKeySimpleInvestorWithdrawByOneNonce // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleInvestorWithdrawByOneNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleInvestorWithdrawByOneNonce)
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
		it.Event = new(KolKeySimpleInvestorWithdrawByOneNonce)
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
func (it *KolKeySimpleInvestorWithdrawByOneNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleInvestorWithdrawByOneNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleInvestorWithdrawByOneNonce represents a InvestorWithdrawByOneNonce event raised by the KolKeySimple contract.
type KolKeySimpleInvestorWithdrawByOneNonce struct {
	Investor common.Address
	Kol      common.Address
	Nonce    *big.Int
	Val      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdrawByOneNonce is a free log retrieval operation binding the contract event 0xc85771d86e5aae3490e851850d1f22697f3a3fc112709b5d2bea1c3ea84fcefc.
//
// Solidity: event InvestorWithdrawByOneNonce(address investor, address kol, uint256 nonce, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) FilterInvestorWithdrawByOneNonce(opts *bind.FilterOpts) (*KolKeySimpleInvestorWithdrawByOneNonceIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "InvestorWithdrawByOneNonce")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleInvestorWithdrawByOneNonceIterator{contract: _KolKeySimple.contract, event: "InvestorWithdrawByOneNonce", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdrawByOneNonce is a free log subscription operation binding the contract event 0xc85771d86e5aae3490e851850d1f22697f3a3fc112709b5d2bea1c3ea84fcefc.
//
// Solidity: event InvestorWithdrawByOneNonce(address investor, address kol, uint256 nonce, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) WatchInvestorWithdrawByOneNonce(opts *bind.WatchOpts, sink chan<- *KolKeySimpleInvestorWithdrawByOneNonce) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "InvestorWithdrawByOneNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleInvestorWithdrawByOneNonce)
				if err := _KolKeySimple.contract.UnpackLog(event, "InvestorWithdrawByOneNonce", log); err != nil {
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

// ParseInvestorWithdrawByOneNonce is a log parse operation binding the contract event 0xc85771d86e5aae3490e851850d1f22697f3a3fc112709b5d2bea1c3ea84fcefc.
//
// Solidity: event InvestorWithdrawByOneNonce(address investor, address kol, uint256 nonce, uint256 val)
func (_KolKeySimple *KolKeySimpleFilterer) ParseInvestorWithdrawByOneNonce(log types.Log) (*KolKeySimpleInvestorWithdrawByOneNonce, error) {
	event := new(KolKeySimpleInvestorWithdrawByOneNonce)
	if err := _KolKeySimple.contract.UnpackLog(event, "InvestorWithdrawByOneNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleKeyReboundIterator is returned from FilterKeyRebound and is used to iterate over the raw logs and unpacked data for KeyRebound events raised by the KolKeySimple contract.
type KolKeySimpleKeyReboundIterator struct {
	Event *KolKeySimpleKeyRebound // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleKeyReboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleKeyRebound)
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
		it.Event = new(KolKeySimpleKeyRebound)
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
func (it *KolKeySimpleKeyReboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleKeyReboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleKeyRebound represents a KeyRebound event raised by the KolKeySimple contract.
type KolKeySimpleKeyRebound struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyRebound is a free log retrieval operation binding the contract event 0xc8818e5ea70e3894e3e342afce8d6f4510704239dc7de14dca1b055ae77bc8f5.
//
// Solidity: event KeyRebound(address from, address to)
func (_KolKeySimple *KolKeySimpleFilterer) FilterKeyRebound(opts *bind.FilterOpts) (*KolKeySimpleKeyReboundIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "KeyRebound")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleKeyReboundIterator{contract: _KolKeySimple.contract, event: "KeyRebound", logs: logs, sub: sub}, nil
}

// WatchKeyRebound is a free log subscription operation binding the contract event 0xc8818e5ea70e3894e3e342afce8d6f4510704239dc7de14dca1b055ae77bc8f5.
//
// Solidity: event KeyRebound(address from, address to)
func (_KolKeySimple *KolKeySimpleFilterer) WatchKeyRebound(opts *bind.WatchOpts, sink chan<- *KolKeySimpleKeyRebound) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "KeyRebound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleKeyRebound)
				if err := _KolKeySimple.contract.UnpackLog(event, "KeyRebound", log); err != nil {
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

// ParseKeyRebound is a log parse operation binding the contract event 0xc8818e5ea70e3894e3e342afce8d6f4510704239dc7de14dca1b055ae77bc8f5.
//
// Solidity: event KeyRebound(address from, address to)
func (_KolKeySimple *KolKeySimpleFilterer) ParseKeyRebound(log types.Log) (*KolKeySimpleKeyRebound, error) {
	event := new(KolKeySimpleKeyRebound)
	if err := _KolKeySimple.contract.UnpackLog(event, "KeyRebound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleKeyTransferIterator is returned from FilterKeyTransfer and is used to iterate over the raw logs and unpacked data for KeyTransfer events raised by the KolKeySimple contract.
type KolKeySimpleKeyTransferIterator struct {
	Event *KolKeySimpleKeyTransfer // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleKeyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleKeyTransfer)
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
		it.Event = new(KolKeySimpleKeyTransfer)
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
func (it *KolKeySimpleKeyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleKeyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleKeyTransfer represents a KeyTransfer event raised by the KolKeySimple contract.
type KolKeySimpleKeyTransfer struct {
	From   common.Address
	To     common.Address
	Kol    common.Address
	Nonce  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeyTransfer is a free log retrieval operation binding the contract event 0x31882570e44c86140b6ff4fe0574206363f217951c784c58da2ecbdf1c9a5133.
//
// Solidity: event KeyTransfer(address from, address to, address kol, uint256 nonce, uint256 amount)
func (_KolKeySimple *KolKeySimpleFilterer) FilterKeyTransfer(opts *bind.FilterOpts) (*KolKeySimpleKeyTransferIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "KeyTransfer")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleKeyTransferIterator{contract: _KolKeySimple.contract, event: "KeyTransfer", logs: logs, sub: sub}, nil
}

// WatchKeyTransfer is a free log subscription operation binding the contract event 0x31882570e44c86140b6ff4fe0574206363f217951c784c58da2ecbdf1c9a5133.
//
// Solidity: event KeyTransfer(address from, address to, address kol, uint256 nonce, uint256 amount)
func (_KolKeySimple *KolKeySimpleFilterer) WatchKeyTransfer(opts *bind.WatchOpts, sink chan<- *KolKeySimpleKeyTransfer) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "KeyTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleKeyTransfer)
				if err := _KolKeySimple.contract.UnpackLog(event, "KeyTransfer", log); err != nil {
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

// ParseKeyTransfer is a log parse operation binding the contract event 0x31882570e44c86140b6ff4fe0574206363f217951c784c58da2ecbdf1c9a5133.
//
// Solidity: event KeyTransfer(address from, address to, address kol, uint256 nonce, uint256 amount)
func (_KolKeySimple *KolKeySimpleFilterer) ParseKeyTransfer(log types.Log) (*KolKeySimpleKeyTransfer, error) {
	event := new(KolKeySimpleKeyTransfer)
	if err := _KolKeySimple.contract.UnpackLog(event, "KeyTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleKeyTransferAllIterator is returned from FilterKeyTransferAll and is used to iterate over the raw logs and unpacked data for KeyTransferAll events raised by the KolKeySimple contract.
type KolKeySimpleKeyTransferAllIterator struct {
	Event *KolKeySimpleKeyTransferAll // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleKeyTransferAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleKeyTransferAll)
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
		it.Event = new(KolKeySimpleKeyTransferAll)
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
func (it *KolKeySimpleKeyTransferAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleKeyTransferAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleKeyTransferAll represents a KeyTransferAll event raised by the KolKeySimple contract.
type KolKeySimpleKeyTransferAll struct {
	From common.Address
	To   common.Address
	Kol  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyTransferAll is a free log retrieval operation binding the contract event 0x35afe37362e9d3a07e354d5f0bedb6cea752a4d72d32228b667ecb7b31af728a.
//
// Solidity: event KeyTransferAll(address from, address to, address kol)
func (_KolKeySimple *KolKeySimpleFilterer) FilterKeyTransferAll(opts *bind.FilterOpts) (*KolKeySimpleKeyTransferAllIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "KeyTransferAll")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleKeyTransferAllIterator{contract: _KolKeySimple.contract, event: "KeyTransferAll", logs: logs, sub: sub}, nil
}

// WatchKeyTransferAll is a free log subscription operation binding the contract event 0x35afe37362e9d3a07e354d5f0bedb6cea752a4d72d32228b667ecb7b31af728a.
//
// Solidity: event KeyTransferAll(address from, address to, address kol)
func (_KolKeySimple *KolKeySimpleFilterer) WatchKeyTransferAll(opts *bind.WatchOpts, sink chan<- *KolKeySimpleKeyTransferAll) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "KeyTransferAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleKeyTransferAll)
				if err := _KolKeySimple.contract.UnpackLog(event, "KeyTransferAll", log); err != nil {
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

// ParseKeyTransferAll is a log parse operation binding the contract event 0x35afe37362e9d3a07e354d5f0bedb6cea752a4d72d32228b667ecb7b31af728a.
//
// Solidity: event KeyTransferAll(address from, address to, address kol)
func (_KolKeySimple *KolKeySimpleFilterer) ParseKeyTransferAll(log types.Log) (*KolKeySimpleKeyTransferAll, error) {
	event := new(KolKeySimpleKeyTransferAll)
	if err := _KolKeySimple.contract.UnpackLog(event, "KeyTransferAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleKolIncomeToPoolActionIterator is returned from FilterKolIncomeToPoolAction and is used to iterate over the raw logs and unpacked data for KolIncomeToPoolAction events raised by the KolKeySimple contract.
type KolKeySimpleKolIncomeToPoolActionIterator struct {
	Event *KolKeySimpleKolIncomeToPoolAction // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleKolIncomeToPoolActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleKolIncomeToPoolAction)
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
		it.Event = new(KolKeySimpleKolIncomeToPoolAction)
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
func (it *KolKeySimpleKolIncomeToPoolActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleKolIncomeToPoolActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleKolIncomeToPoolAction represents a KolIncomeToPoolAction event raised by the KolKeySimple contract.
type KolKeySimpleKolIncomeToPoolAction struct {
	SourceID      int8
	SourceConract common.Address
	Kol           common.Address
	KeyNo         *big.Int
	KeyNonce      *big.Int
	Amount        *big.Int
	ValPerKey     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterKolIncomeToPoolAction is a free log retrieval operation binding the contract event 0x4c20a8329633f42cfa3bee42045c6ee0360ac4ceeff4088d60d0bb24c624e330.
//
// Solidity: event KolIncomeToPoolAction(int8 sourceID, address sourceConract, address kol, uint256 keyNo, uint256 keyNonce, uint256 amount, uint256 valPerKey)
func (_KolKeySimple *KolKeySimpleFilterer) FilterKolIncomeToPoolAction(opts *bind.FilterOpts) (*KolKeySimpleKolIncomeToPoolActionIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "KolIncomeToPoolAction")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleKolIncomeToPoolActionIterator{contract: _KolKeySimple.contract, event: "KolIncomeToPoolAction", logs: logs, sub: sub}, nil
}

// WatchKolIncomeToPoolAction is a free log subscription operation binding the contract event 0x4c20a8329633f42cfa3bee42045c6ee0360ac4ceeff4088d60d0bb24c624e330.
//
// Solidity: event KolIncomeToPoolAction(int8 sourceID, address sourceConract, address kol, uint256 keyNo, uint256 keyNonce, uint256 amount, uint256 valPerKey)
func (_KolKeySimple *KolKeySimpleFilterer) WatchKolIncomeToPoolAction(opts *bind.WatchOpts, sink chan<- *KolKeySimpleKolIncomeToPoolAction) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "KolIncomeToPoolAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleKolIncomeToPoolAction)
				if err := _KolKeySimple.contract.UnpackLog(event, "KolIncomeToPoolAction", log); err != nil {
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

// ParseKolIncomeToPoolAction is a log parse operation binding the contract event 0x4c20a8329633f42cfa3bee42045c6ee0360ac4ceeff4088d60d0bb24c624e330.
//
// Solidity: event KolIncomeToPoolAction(int8 sourceID, address sourceConract, address kol, uint256 keyNo, uint256 keyNonce, uint256 amount, uint256 valPerKey)
func (_KolKeySimple *KolKeySimpleFilterer) ParseKolIncomeToPoolAction(log types.Log) (*KolKeySimpleKolIncomeToPoolAction, error) {
	event := new(KolKeySimpleKolIncomeToPoolAction)
	if err := _KolKeySimple.contract.UnpackLog(event, "KolIncomeToPoolAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleKolKeyBoughtIterator is returned from FilterKolKeyBought and is used to iterate over the raw logs and unpacked data for KolKeyBought events raised by the KolKeySimple contract.
type KolKeySimpleKolKeyBoughtIterator struct {
	Event *KolKeySimpleKolKeyBought // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleKolKeyBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleKolKeyBought)
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
		it.Event = new(KolKeySimpleKolKeyBought)
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
func (it *KolKeySimpleKolKeyBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleKolKeyBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleKolKeyBought represents a KolKeyBought event raised by the KolKeySimple contract.
type KolKeySimpleKolKeyBought struct {
	KolAddr    common.Address
	Buyer      common.Address
	KeyNo      *big.Int
	CurNonce   *big.Int
	KoltotalNo *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKolKeyBought is a free log retrieval operation binding the contract event 0xa43d1ee2d295f126339683f7a2c5bbb65a75bd3506059340c09ad67168ebbfa8.
//
// Solidity: event KolKeyBought(address kolAddr, address buyer, uint256 keyNo, uint256 curNonce, uint256 KoltotalNo)
func (_KolKeySimple *KolKeySimpleFilterer) FilterKolKeyBought(opts *bind.FilterOpts) (*KolKeySimpleKolKeyBoughtIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "KolKeyBought")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleKolKeyBoughtIterator{contract: _KolKeySimple.contract, event: "KolKeyBought", logs: logs, sub: sub}, nil
}

// WatchKolKeyBought is a free log subscription operation binding the contract event 0xa43d1ee2d295f126339683f7a2c5bbb65a75bd3506059340c09ad67168ebbfa8.
//
// Solidity: event KolKeyBought(address kolAddr, address buyer, uint256 keyNo, uint256 curNonce, uint256 KoltotalNo)
func (_KolKeySimple *KolKeySimpleFilterer) WatchKolKeyBought(opts *bind.WatchOpts, sink chan<- *KolKeySimpleKolKeyBought) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "KolKeyBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleKolKeyBought)
				if err := _KolKeySimple.contract.UnpackLog(event, "KolKeyBought", log); err != nil {
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

// ParseKolKeyBought is a log parse operation binding the contract event 0xa43d1ee2d295f126339683f7a2c5bbb65a75bd3506059340c09ad67168ebbfa8.
//
// Solidity: event KolKeyBought(address kolAddr, address buyer, uint256 keyNo, uint256 curNonce, uint256 KoltotalNo)
func (_KolKeySimple *KolKeySimpleFilterer) ParseKolKeyBought(log types.Log) (*KolKeySimpleKolKeyBought, error) {
	event := new(KolKeySimpleKolKeyBought)
	if err := _KolKeySimple.contract.UnpackLog(event, "KolKeyBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleKolKeyOpenedIterator is returned from FilterKolKeyOpened and is used to iterate over the raw logs and unpacked data for KolKeyOpened events raised by the KolKeySimple contract.
type KolKeySimpleKolKeyOpenedIterator struct {
	Event *KolKeySimpleKolKeyOpened // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleKolKeyOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleKolKeyOpened)
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
		it.Event = new(KolKeySimpleKolKeyOpened)
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
func (it *KolKeySimpleKolKeyOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleKolKeyOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleKolKeyOpened represents a KolKeyOpened event raised by the KolKeySimple contract.
type KolKeySimpleKolKeyOpened struct {
	Kol         common.Address
	PricePerKey *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKolKeyOpened is a free log retrieval operation binding the contract event 0x9752e725b9e262a3587203f1a8d8324ac9974062075f51c1c852d8422c27f943.
//
// Solidity: event KolKeyOpened(address kol, uint256 pricePerKey)
func (_KolKeySimple *KolKeySimpleFilterer) FilterKolKeyOpened(opts *bind.FilterOpts) (*KolKeySimpleKolKeyOpenedIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "KolKeyOpened")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleKolKeyOpenedIterator{contract: _KolKeySimple.contract, event: "KolKeyOpened", logs: logs, sub: sub}, nil
}

// WatchKolKeyOpened is a free log subscription operation binding the contract event 0x9752e725b9e262a3587203f1a8d8324ac9974062075f51c1c852d8422c27f943.
//
// Solidity: event KolKeyOpened(address kol, uint256 pricePerKey)
func (_KolKeySimple *KolKeySimpleFilterer) WatchKolKeyOpened(opts *bind.WatchOpts, sink chan<- *KolKeySimpleKolKeyOpened) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "KolKeyOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleKolKeyOpened)
				if err := _KolKeySimple.contract.UnpackLog(event, "KolKeyOpened", log); err != nil {
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

// ParseKolKeyOpened is a log parse operation binding the contract event 0x9752e725b9e262a3587203f1a8d8324ac9974062075f51c1c852d8422c27f943.
//
// Solidity: event KolKeyOpened(address kol, uint256 pricePerKey)
func (_KolKeySimple *KolKeySimpleFilterer) ParseKolKeyOpened(log types.Log) (*KolKeySimpleKolKeyOpened, error) {
	event := new(KolKeySimpleKolKeyOpened)
	if err := _KolKeySimple.contract.UnpackLog(event, "KolKeyOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the KolKeySimple contract.
type KolKeySimpleOwnerSetIterator struct {
	Event *KolKeySimpleOwnerSet // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleOwnerSet)
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
		it.Event = new(KolKeySimpleOwnerSet)
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
func (it *KolKeySimpleOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleOwnerSet represents a OwnerSet event raised by the KolKeySimple contract.
type KolKeySimpleOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_KolKeySimple *KolKeySimpleFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*KolKeySimpleOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleOwnerSetIterator{contract: _KolKeySimple.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_KolKeySimple *KolKeySimpleFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *KolKeySimpleOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleOwnerSet)
				if err := _KolKeySimple.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_KolKeySimple *KolKeySimpleFilterer) ParseOwnerSet(log types.Log) (*KolKeySimpleOwnerSet, error) {
	event := new(KolKeySimpleOwnerSet)
	if err := _KolKeySimple.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleServiceFeeChangedIterator is returned from FilterServiceFeeChanged and is used to iterate over the raw logs and unpacked data for ServiceFeeChanged events raised by the KolKeySimple contract.
type KolKeySimpleServiceFeeChangedIterator struct {
	Event *KolKeySimpleServiceFeeChanged // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleServiceFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleServiceFeeChanged)
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
		it.Event = new(KolKeySimpleServiceFeeChanged)
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
func (it *KolKeySimpleServiceFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleServiceFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleServiceFeeChanged represents a ServiceFeeChanged event raised by the KolKeySimple contract.
type KolKeySimpleServiceFeeChanged struct {
	NewSerficeFeeRate *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeChanged is a free log retrieval operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_KolKeySimple *KolKeySimpleFilterer) FilterServiceFeeChanged(opts *bind.FilterOpts) (*KolKeySimpleServiceFeeChangedIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleServiceFeeChangedIterator{contract: _KolKeySimple.contract, event: "ServiceFeeChanged", logs: logs, sub: sub}, nil
}

// WatchServiceFeeChanged is a free log subscription operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_KolKeySimple *KolKeySimpleFilterer) WatchServiceFeeChanged(opts *bind.WatchOpts, sink chan<- *KolKeySimpleServiceFeeChanged) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleServiceFeeChanged)
				if err := _KolKeySimple.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
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

// ParseServiceFeeChanged is a log parse operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_KolKeySimple *KolKeySimpleFilterer) ParseServiceFeeChanged(log types.Log) (*KolKeySimpleServiceFeeChanged, error) {
	event := new(KolKeySimpleServiceFeeChanged)
	if err := _KolKeySimple.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleSystemSetIterator is returned from FilterSystemSet and is used to iterate over the raw logs and unpacked data for SystemSet events raised by the KolKeySimple contract.
type KolKeySimpleSystemSetIterator struct {
	Event *KolKeySimpleSystemSet // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleSystemSet)
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
		it.Event = new(KolKeySimpleSystemSet)
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
func (it *KolKeySimpleSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleSystemSet represents a SystemSet event raised by the KolKeySimple contract.
type KolKeySimpleSystemSet struct {
	Num *big.Int
	Op  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSystemSet is a free log retrieval operation binding the contract event 0xcd714485b704afdf17516cfe0560106372ea651897ae76ae6879e628756fb2a0.
//
// Solidity: event SystemSet(uint256 num, string op)
func (_KolKeySimple *KolKeySimpleFilterer) FilterSystemSet(opts *bind.FilterOpts) (*KolKeySimpleSystemSetIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "SystemSet")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleSystemSetIterator{contract: _KolKeySimple.contract, event: "SystemSet", logs: logs, sub: sub}, nil
}

// WatchSystemSet is a free log subscription operation binding the contract event 0xcd714485b704afdf17516cfe0560106372ea651897ae76ae6879e628756fb2a0.
//
// Solidity: event SystemSet(uint256 num, string op)
func (_KolKeySimple *KolKeySimpleFilterer) WatchSystemSet(opts *bind.WatchOpts, sink chan<- *KolKeySimpleSystemSet) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "SystemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleSystemSet)
				if err := _KolKeySimple.contract.UnpackLog(event, "SystemSet", log); err != nil {
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

// ParseSystemSet is a log parse operation binding the contract event 0xcd714485b704afdf17516cfe0560106372ea651897ae76ae6879e628756fb2a0.
//
// Solidity: event SystemSet(uint256 num, string op)
func (_KolKeySimple *KolKeySimpleFilterer) ParseSystemSet(log types.Log) (*KolKeySimpleSystemSet, error) {
	event := new(KolKeySimpleSystemSet)
	if err := _KolKeySimple.contract.UnpackLog(event, "SystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleUpgradeToNewRuleIterator is returned from FilterUpgradeToNewRule and is used to iterate over the raw logs and unpacked data for UpgradeToNewRule events raised by the KolKeySimple contract.
type KolKeySimpleUpgradeToNewRuleIterator struct {
	Event *KolKeySimpleUpgradeToNewRule // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleUpgradeToNewRuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleUpgradeToNewRule)
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
		it.Event = new(KolKeySimpleUpgradeToNewRule)
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
func (it *KolKeySimpleUpgradeToNewRuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleUpgradeToNewRuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleUpgradeToNewRule represents a UpgradeToNewRule event raised by the KolKeySimple contract.
type KolKeySimpleUpgradeToNewRule struct {
	NewContract common.Address
	Balance     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpgradeToNewRule is a free log retrieval operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_KolKeySimple *KolKeySimpleFilterer) FilterUpgradeToNewRule(opts *bind.FilterOpts) (*KolKeySimpleUpgradeToNewRuleIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleUpgradeToNewRuleIterator{contract: _KolKeySimple.contract, event: "UpgradeToNewRule", logs: logs, sub: sub}, nil
}

// WatchUpgradeToNewRule is a free log subscription operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_KolKeySimple *KolKeySimpleFilterer) WatchUpgradeToNewRule(opts *bind.WatchOpts, sink chan<- *KolKeySimpleUpgradeToNewRule) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleUpgradeToNewRule)
				if err := _KolKeySimple.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
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

// ParseUpgradeToNewRule is a log parse operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_KolKeySimple *KolKeySimpleFilterer) ParseUpgradeToNewRule(log types.Log) (*KolKeySimpleUpgradeToNewRule, error) {
	event := new(KolKeySimpleUpgradeToNewRule)
	if err := _KolKeySimple.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeySimpleWithdrawServiceIterator is returned from FilterWithdrawService and is used to iterate over the raw logs and unpacked data for WithdrawService events raised by the KolKeySimple contract.
type KolKeySimpleWithdrawServiceIterator struct {
	Event *KolKeySimpleWithdrawService // Event containing the contract specifics and raw log

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
func (it *KolKeySimpleWithdrawServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeySimpleWithdrawService)
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
		it.Event = new(KolKeySimpleWithdrawService)
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
func (it *KolKeySimpleWithdrawServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeySimpleWithdrawServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeySimpleWithdrawService represents a WithdrawService event raised by the KolKeySimple contract.
type KolKeySimpleWithdrawService struct {
	Owner   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawService is a free log retrieval operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_KolKeySimple *KolKeySimpleFilterer) FilterWithdrawService(opts *bind.FilterOpts) (*KolKeySimpleWithdrawServiceIterator, error) {

	logs, sub, err := _KolKeySimple.contract.FilterLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return &KolKeySimpleWithdrawServiceIterator{contract: _KolKeySimple.contract, event: "WithdrawService", logs: logs, sub: sub}, nil
}

// WatchWithdrawService is a free log subscription operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_KolKeySimple *KolKeySimpleFilterer) WatchWithdrawService(opts *bind.WatchOpts, sink chan<- *KolKeySimpleWithdrawService) (event.Subscription, error) {

	logs, sub, err := _KolKeySimple.contract.WatchLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeySimpleWithdrawService)
				if err := _KolKeySimple.contract.UnpackLog(event, "WithdrawService", log); err != nil {
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

// ParseWithdrawService is a log parse operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_KolKeySimple *KolKeySimpleFilterer) ParseWithdrawService(log types.Log) (*KolKeySimpleWithdrawService, error) {
	event := new(KolKeySimpleWithdrawService)
	if err := _KolKeySimple.contract.UnpackLog(event, "WithdrawService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
