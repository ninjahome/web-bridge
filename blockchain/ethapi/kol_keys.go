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

// KolKeysKeyAuction is an auto generated low-level Go binding around an user-defined struct.
type KolKeysKeyAuction struct {
	Id     *big.Int
	Seller common.Address
	Bidder common.Address
	Kol    common.Address
	Nonce  *big.Int
	Amount *big.Int
	Price  *big.Int
	Status bool
}

// KolKeysKeySettings is an auto generated low-level Go binding around an user-defined struct.
type KolKeysKeySettings struct {
	Price   *big.Int
	MaxNo   *big.Int
	Nonce   *big.Int
	TotalNo *big.Int
}

// KolKeysMetaData contains all meta data concerning the KolKeys contract.
var KolKeysMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"opType\",\"type\":\"bool\"}],\"name\":\"AdminOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdrawAllIncome\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdrawByOneKol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdrawByOneNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typ\",\"type\":\"string\"}],\"name\":\"KeyAuctionAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeyTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"KeyTransferedAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kolAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyNo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"KoltotalNo\",\"type\":\"uint256\"}],\"name\":\"KolIpRightBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxKeyNo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"op\",\"type\":\"string\"}],\"name\":\"KolKeyOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSerficeFeeRate\",\"type\":\"uint256\"}],\"name\":\"ServiceFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"op\",\"type\":\"string\"}],\"name\":\"SystemSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteNoe\",\"type\":\"uint256\"}],\"name\":\"TweetBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"UpgradeToNewRule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"WithdrawService\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"AllIncomeOfAllKol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"AllIncomeOfOneKol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"}],\"name\":\"AuctionData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structKolKeys.KeyAuction[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncomeOfOneNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"InvestorAllKeysOfKol\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonce\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"InvestorOfKol\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"KeySettingsRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNo\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"KeyStatusOfKol\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNo\",\"type\":\"uint256\"}],\"internalType\":\"structKolKeys.KeySettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"KolOfOneInvestor\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"__admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__currentAuctionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__feeForKeyBiddig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__feeForKolKeyOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__minValCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDelete\",\"type\":\"bool\"}],\"name\":\"adminOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminServiceFeeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeInGwei\",\"type\":\"uint256\"}],\"name\":\"adminSetBddingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetKeyFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetKolIncomeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeInGwei\",\"type\":\"uint256\"}],\"name\":\"adminSetKolOperationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetWithdrawFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"adminUpgradeToNewRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionIdxMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidID\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kolAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keyNo\",\"type\":\"uint256\"}],\"name\":\"buyKolKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPluginInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"incomePerNoncePerKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"incomeToKolPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerKey\",\"type\":\"uint256\"}],\"name\":\"issueBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keyInAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"kolAddKeySupply\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolIncomeRatePerKeyBuy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerKey\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"maxKeyNo\",\"type\":\"int256\"}],\"name\":\"kolOpenKeySale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"kolSetKeyPrice\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidID\",\"type\":\"uint256\"}],\"name\":\"revokeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeRatePerKeyBuy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"transferAllKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tweetHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tweetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voteNo\",\"type\":\"uint256\"}],\"name\":\"tweetBought\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"all\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllIncome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"}],\"name\":\"withdrawFromOneKol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"withdrawFromOneNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KolKeysABI is the input ABI used to generate the binding from.
// Deprecated: Use KolKeysMetaData.ABI instead.
var KolKeysABI = KolKeysMetaData.ABI

// KolKeys is an auto generated Go binding around an Ethereum contract.
type KolKeys struct {
	KolKeysCaller     // Read-only binding to the contract
	KolKeysTransactor // Write-only binding to the contract
	KolKeysFilterer   // Log filterer for contract events
}

// KolKeysCaller is an auto generated read-only Go binding around an Ethereum contract.
type KolKeysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KolKeysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KolKeysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KolKeysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KolKeysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KolKeysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KolKeysSession struct {
	Contract     *KolKeys          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KolKeysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KolKeysCallerSession struct {
	Contract *KolKeysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KolKeysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KolKeysTransactorSession struct {
	Contract     *KolKeysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KolKeysRaw is an auto generated low-level Go binding around an Ethereum contract.
type KolKeysRaw struct {
	Contract *KolKeys // Generic contract binding to access the raw methods on
}

// KolKeysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KolKeysCallerRaw struct {
	Contract *KolKeysCaller // Generic read-only contract binding to access the raw methods on
}

// KolKeysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KolKeysTransactorRaw struct {
	Contract *KolKeysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKolKeys creates a new instance of KolKeys, bound to a specific deployed contract.
func NewKolKeys(address common.Address, backend bind.ContractBackend) (*KolKeys, error) {
	contract, err := bindKolKeys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KolKeys{KolKeysCaller: KolKeysCaller{contract: contract}, KolKeysTransactor: KolKeysTransactor{contract: contract}, KolKeysFilterer: KolKeysFilterer{contract: contract}}, nil
}

// NewKolKeysCaller creates a new read-only instance of KolKeys, bound to a specific deployed contract.
func NewKolKeysCaller(address common.Address, caller bind.ContractCaller) (*KolKeysCaller, error) {
	contract, err := bindKolKeys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KolKeysCaller{contract: contract}, nil
}

// NewKolKeysTransactor creates a new write-only instance of KolKeys, bound to a specific deployed contract.
func NewKolKeysTransactor(address common.Address, transactor bind.ContractTransactor) (*KolKeysTransactor, error) {
	contract, err := bindKolKeys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KolKeysTransactor{contract: contract}, nil
}

// NewKolKeysFilterer creates a new log filterer instance of KolKeys, bound to a specific deployed contract.
func NewKolKeysFilterer(address common.Address, filterer bind.ContractFilterer) (*KolKeysFilterer, error) {
	contract, err := bindKolKeys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KolKeysFilterer{contract: contract}, nil
}

// bindKolKeys binds a generic wrapper to an already deployed contract.
func bindKolKeys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KolKeysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KolKeys *KolKeysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KolKeys.Contract.KolKeysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KolKeys *KolKeysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeys.Contract.KolKeysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KolKeys *KolKeysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KolKeys.Contract.KolKeysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KolKeys *KolKeysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KolKeys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KolKeys *KolKeysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KolKeys *KolKeysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KolKeys.Contract.contract.Transact(opts, method, params...)
}

// AllIncomeOfAllKol is a free data retrieval call binding the contract method 0xaf958a62.
//
// Solidity: function AllIncomeOfAllKol(address investor) view returns(uint256)
func (_KolKeys *KolKeysCaller) AllIncomeOfAllKol(opts *bind.CallOpts, investor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "AllIncomeOfAllKol", investor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllIncomeOfAllKol is a free data retrieval call binding the contract method 0xaf958a62.
//
// Solidity: function AllIncomeOfAllKol(address investor) view returns(uint256)
func (_KolKeys *KolKeysSession) AllIncomeOfAllKol(investor common.Address) (*big.Int, error) {
	return _KolKeys.Contract.AllIncomeOfAllKol(&_KolKeys.CallOpts, investor)
}

// AllIncomeOfAllKol is a free data retrieval call binding the contract method 0xaf958a62.
//
// Solidity: function AllIncomeOfAllKol(address investor) view returns(uint256)
func (_KolKeys *KolKeysCallerSession) AllIncomeOfAllKol(investor common.Address) (*big.Int, error) {
	return _KolKeys.Contract.AllIncomeOfAllKol(&_KolKeys.CallOpts, investor)
}

// AllIncomeOfOneKol is a free data retrieval call binding the contract method 0x6d74bb5f.
//
// Solidity: function AllIncomeOfOneKol(address kol, address investor) view returns(uint256)
func (_KolKeys *KolKeysCaller) AllIncomeOfOneKol(opts *bind.CallOpts, kol common.Address, investor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "AllIncomeOfOneKol", kol, investor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllIncomeOfOneKol is a free data retrieval call binding the contract method 0x6d74bb5f.
//
// Solidity: function AllIncomeOfOneKol(address kol, address investor) view returns(uint256)
func (_KolKeys *KolKeysSession) AllIncomeOfOneKol(kol common.Address, investor common.Address) (*big.Int, error) {
	return _KolKeys.Contract.AllIncomeOfOneKol(&_KolKeys.CallOpts, kol, investor)
}

// AllIncomeOfOneKol is a free data retrieval call binding the contract method 0x6d74bb5f.
//
// Solidity: function AllIncomeOfOneKol(address kol, address investor) view returns(uint256)
func (_KolKeys *KolKeysCallerSession) AllIncomeOfOneKol(kol common.Address, investor common.Address) (*big.Int, error) {
	return _KolKeys.Contract.AllIncomeOfOneKol(&_KolKeys.CallOpts, kol, investor)
}

// AuctionData is a free data retrieval call binding the contract method 0x86eaf599.
//
// Solidity: function AuctionData(uint256 start, uint256 no) view returns((uint256,address,address,address,uint256,uint256,uint256,bool)[] result)
func (_KolKeys *KolKeysCaller) AuctionData(opts *bind.CallOpts, start *big.Int, no *big.Int) ([]KolKeysKeyAuction, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "AuctionData", start, no)

	if err != nil {
		return *new([]KolKeysKeyAuction), err
	}

	out0 := *abi.ConvertType(out[0], new([]KolKeysKeyAuction)).(*[]KolKeysKeyAuction)

	return out0, err

}

// AuctionData is a free data retrieval call binding the contract method 0x86eaf599.
//
// Solidity: function AuctionData(uint256 start, uint256 no) view returns((uint256,address,address,address,uint256,uint256,uint256,bool)[] result)
func (_KolKeys *KolKeysSession) AuctionData(start *big.Int, no *big.Int) ([]KolKeysKeyAuction, error) {
	return _KolKeys.Contract.AuctionData(&_KolKeys.CallOpts, start, no)
}

// AuctionData is a free data retrieval call binding the contract method 0x86eaf599.
//
// Solidity: function AuctionData(uint256 start, uint256 no) view returns((uint256,address,address,address,uint256,uint256,uint256,bool)[] result)
func (_KolKeys *KolKeysCallerSession) AuctionData(start *big.Int, no *big.Int) ([]KolKeysKeyAuction, error) {
	return _KolKeys.Contract.AuctionData(&_KolKeys.CallOpts, start, no)
}

// IncomeOfOneNonce is a free data retrieval call binding the contract method 0x54ffe238.
//
// Solidity: function IncomeOfOneNonce(address kol, uint256 nonce, uint256 amount) view returns(uint256)
func (_KolKeys *KolKeysCaller) IncomeOfOneNonce(opts *bind.CallOpts, kol common.Address, nonce *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "IncomeOfOneNonce", kol, nonce, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomeOfOneNonce is a free data retrieval call binding the contract method 0x54ffe238.
//
// Solidity: function IncomeOfOneNonce(address kol, uint256 nonce, uint256 amount) view returns(uint256)
func (_KolKeys *KolKeysSession) IncomeOfOneNonce(kol common.Address, nonce *big.Int, amount *big.Int) (*big.Int, error) {
	return _KolKeys.Contract.IncomeOfOneNonce(&_KolKeys.CallOpts, kol, nonce, amount)
}

// IncomeOfOneNonce is a free data retrieval call binding the contract method 0x54ffe238.
//
// Solidity: function IncomeOfOneNonce(address kol, uint256 nonce, uint256 amount) view returns(uint256)
func (_KolKeys *KolKeysCallerSession) IncomeOfOneNonce(kol common.Address, nonce *big.Int, amount *big.Int) (*big.Int, error) {
	return _KolKeys.Contract.IncomeOfOneNonce(&_KolKeys.CallOpts, kol, nonce, amount)
}

// InvestorAllKeysOfKol is a free data retrieval call binding the contract method 0xb83bb4e6.
//
// Solidity: function InvestorAllKeysOfKol(address investor, address kol) view returns(uint256[] nonce, uint256[] amounts)
func (_KolKeys *KolKeysCaller) InvestorAllKeysOfKol(opts *bind.CallOpts, investor common.Address, kol common.Address) (struct {
	Nonce   []*big.Int
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "InvestorAllKeysOfKol", investor, kol)

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
func (_KolKeys *KolKeysSession) InvestorAllKeysOfKol(investor common.Address, kol common.Address) (struct {
	Nonce   []*big.Int
	Amounts []*big.Int
}, error) {
	return _KolKeys.Contract.InvestorAllKeysOfKol(&_KolKeys.CallOpts, investor, kol)
}

// InvestorAllKeysOfKol is a free data retrieval call binding the contract method 0xb83bb4e6.
//
// Solidity: function InvestorAllKeysOfKol(address investor, address kol) view returns(uint256[] nonce, uint256[] amounts)
func (_KolKeys *KolKeysCallerSession) InvestorAllKeysOfKol(investor common.Address, kol common.Address) (struct {
	Nonce   []*big.Int
	Amounts []*big.Int
}, error) {
	return _KolKeys.Contract.InvestorAllKeysOfKol(&_KolKeys.CallOpts, investor, kol)
}

// InvestorOfKol is a free data retrieval call binding the contract method 0x8b2b0b67.
//
// Solidity: function InvestorOfKol(address kol) view returns(address[])
func (_KolKeys *KolKeysCaller) InvestorOfKol(opts *bind.CallOpts, kol common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "InvestorOfKol", kol)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// InvestorOfKol is a free data retrieval call binding the contract method 0x8b2b0b67.
//
// Solidity: function InvestorOfKol(address kol) view returns(address[])
func (_KolKeys *KolKeysSession) InvestorOfKol(kol common.Address) ([]common.Address, error) {
	return _KolKeys.Contract.InvestorOfKol(&_KolKeys.CallOpts, kol)
}

// InvestorOfKol is a free data retrieval call binding the contract method 0x8b2b0b67.
//
// Solidity: function InvestorOfKol(address kol) view returns(address[])
func (_KolKeys *KolKeysCallerSession) InvestorOfKol(kol common.Address) ([]common.Address, error) {
	return _KolKeys.Contract.InvestorOfKol(&_KolKeys.CallOpts, kol)
}

// KeySettingsRecord is a free data retrieval call binding the contract method 0x94447f45.
//
// Solidity: function KeySettingsRecord(address ) view returns(uint256 price, uint256 maxNo, uint256 nonce, uint256 totalNo)
func (_KolKeys *KolKeysCaller) KeySettingsRecord(opts *bind.CallOpts, arg0 common.Address) (struct {
	Price   *big.Int
	MaxNo   *big.Int
	Nonce   *big.Int
	TotalNo *big.Int
}, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "KeySettingsRecord", arg0)

	outstruct := new(struct {
		Price   *big.Int
		MaxNo   *big.Int
		Nonce   *big.Int
		TotalNo *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxNo = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalNo = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeySettingsRecord is a free data retrieval call binding the contract method 0x94447f45.
//
// Solidity: function KeySettingsRecord(address ) view returns(uint256 price, uint256 maxNo, uint256 nonce, uint256 totalNo)
func (_KolKeys *KolKeysSession) KeySettingsRecord(arg0 common.Address) (struct {
	Price   *big.Int
	MaxNo   *big.Int
	Nonce   *big.Int
	TotalNo *big.Int
}, error) {
	return _KolKeys.Contract.KeySettingsRecord(&_KolKeys.CallOpts, arg0)
}

// KeySettingsRecord is a free data retrieval call binding the contract method 0x94447f45.
//
// Solidity: function KeySettingsRecord(address ) view returns(uint256 price, uint256 maxNo, uint256 nonce, uint256 totalNo)
func (_KolKeys *KolKeysCallerSession) KeySettingsRecord(arg0 common.Address) (struct {
	Price   *big.Int
	MaxNo   *big.Int
	Nonce   *big.Int
	TotalNo *big.Int
}, error) {
	return _KolKeys.Contract.KeySettingsRecord(&_KolKeys.CallOpts, arg0)
}

// KeyStatusOfKol is a free data retrieval call binding the contract method 0x9e4f5b5f.
//
// Solidity: function KeyStatusOfKol(address kol) view returns((uint256,uint256,uint256,uint256))
func (_KolKeys *KolKeysCaller) KeyStatusOfKol(opts *bind.CallOpts, kol common.Address) (KolKeysKeySettings, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "KeyStatusOfKol", kol)

	if err != nil {
		return *new(KolKeysKeySettings), err
	}

	out0 := *abi.ConvertType(out[0], new(KolKeysKeySettings)).(*KolKeysKeySettings)

	return out0, err

}

// KeyStatusOfKol is a free data retrieval call binding the contract method 0x9e4f5b5f.
//
// Solidity: function KeyStatusOfKol(address kol) view returns((uint256,uint256,uint256,uint256))
func (_KolKeys *KolKeysSession) KeyStatusOfKol(kol common.Address) (KolKeysKeySettings, error) {
	return _KolKeys.Contract.KeyStatusOfKol(&_KolKeys.CallOpts, kol)
}

// KeyStatusOfKol is a free data retrieval call binding the contract method 0x9e4f5b5f.
//
// Solidity: function KeyStatusOfKol(address kol) view returns((uint256,uint256,uint256,uint256))
func (_KolKeys *KolKeysCallerSession) KeyStatusOfKol(kol common.Address) (KolKeysKeySettings, error) {
	return _KolKeys.Contract.KeyStatusOfKol(&_KolKeys.CallOpts, kol)
}

// KolOfOneInvestor is a free data retrieval call binding the contract method 0xce92faa2.
//
// Solidity: function KolOfOneInvestor(address investor) view returns(address[])
func (_KolKeys *KolKeysCaller) KolOfOneInvestor(opts *bind.CallOpts, investor common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "KolOfOneInvestor", investor)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// KolOfOneInvestor is a free data retrieval call binding the contract method 0xce92faa2.
//
// Solidity: function KolOfOneInvestor(address investor) view returns(address[])
func (_KolKeys *KolKeysSession) KolOfOneInvestor(investor common.Address) ([]common.Address, error) {
	return _KolKeys.Contract.KolOfOneInvestor(&_KolKeys.CallOpts, investor)
}

// KolOfOneInvestor is a free data retrieval call binding the contract method 0xce92faa2.
//
// Solidity: function KolOfOneInvestor(address investor) view returns(address[])
func (_KolKeys *KolKeysCallerSession) KolOfOneInvestor(investor common.Address) ([]common.Address, error) {
	return _KolKeys.Contract.KolOfOneInvestor(&_KolKeys.CallOpts, investor)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_KolKeys *KolKeysCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "__admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_KolKeys *KolKeysSession) Admins(arg0 common.Address) (bool, error) {
	return _KolKeys.Contract.Admins(&_KolKeys.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_KolKeys *KolKeysCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _KolKeys.Contract.Admins(&_KolKeys.CallOpts, arg0)
}

// CurrentAuctionID is a free data retrieval call binding the contract method 0x1d6abf71.
//
// Solidity: function __currentAuctionID() view returns(uint256)
func (_KolKeys *KolKeysCaller) CurrentAuctionID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "__currentAuctionID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAuctionID is a free data retrieval call binding the contract method 0x1d6abf71.
//
// Solidity: function __currentAuctionID() view returns(uint256)
func (_KolKeys *KolKeysSession) CurrentAuctionID() (*big.Int, error) {
	return _KolKeys.Contract.CurrentAuctionID(&_KolKeys.CallOpts)
}

// CurrentAuctionID is a free data retrieval call binding the contract method 0x1d6abf71.
//
// Solidity: function __currentAuctionID() view returns(uint256)
func (_KolKeys *KolKeysCallerSession) CurrentAuctionID() (*big.Int, error) {
	return _KolKeys.Contract.CurrentAuctionID(&_KolKeys.CallOpts)
}

// FeeForKeyBiddig is a free data retrieval call binding the contract method 0xd4596578.
//
// Solidity: function __feeForKeyBiddig() view returns(uint256)
func (_KolKeys *KolKeysCaller) FeeForKeyBiddig(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "__feeForKeyBiddig")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeForKeyBiddig is a free data retrieval call binding the contract method 0xd4596578.
//
// Solidity: function __feeForKeyBiddig() view returns(uint256)
func (_KolKeys *KolKeysSession) FeeForKeyBiddig() (*big.Int, error) {
	return _KolKeys.Contract.FeeForKeyBiddig(&_KolKeys.CallOpts)
}

// FeeForKeyBiddig is a free data retrieval call binding the contract method 0xd4596578.
//
// Solidity: function __feeForKeyBiddig() view returns(uint256)
func (_KolKeys *KolKeysCallerSession) FeeForKeyBiddig() (*big.Int, error) {
	return _KolKeys.Contract.FeeForKeyBiddig(&_KolKeys.CallOpts)
}

// FeeForKolKeyOp is a free data retrieval call binding the contract method 0xcdd3e4da.
//
// Solidity: function __feeForKolKeyOp() view returns(uint256)
func (_KolKeys *KolKeysCaller) FeeForKolKeyOp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "__feeForKolKeyOp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeForKolKeyOp is a free data retrieval call binding the contract method 0xcdd3e4da.
//
// Solidity: function __feeForKolKeyOp() view returns(uint256)
func (_KolKeys *KolKeysSession) FeeForKolKeyOp() (*big.Int, error) {
	return _KolKeys.Contract.FeeForKolKeyOp(&_KolKeys.CallOpts)
}

// FeeForKolKeyOp is a free data retrieval call binding the contract method 0xcdd3e4da.
//
// Solidity: function __feeForKolKeyOp() view returns(uint256)
func (_KolKeys *KolKeysCallerSession) FeeForKolKeyOp() (*big.Int, error) {
	return _KolKeys.Contract.FeeForKolKeyOp(&_KolKeys.CallOpts)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_KolKeys *KolKeysCaller) MinValCheck(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "__minValCheck")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_KolKeys *KolKeysSession) MinValCheck() (*big.Int, error) {
	return _KolKeys.Contract.MinValCheck(&_KolKeys.CallOpts)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_KolKeys *KolKeysCallerSession) MinValCheck() (*big.Int, error) {
	return _KolKeys.Contract.MinValCheck(&_KolKeys.CallOpts)
}

// AuctionIdxMap is a free data retrieval call binding the contract method 0x794f3bf1.
//
// Solidity: function auctionIdxMap(uint256 ) view returns(uint256)
func (_KolKeys *KolKeysCaller) AuctionIdxMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "auctionIdxMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionIdxMap is a free data retrieval call binding the contract method 0x794f3bf1.
//
// Solidity: function auctionIdxMap(uint256 ) view returns(uint256)
func (_KolKeys *KolKeysSession) AuctionIdxMap(arg0 *big.Int) (*big.Int, error) {
	return _KolKeys.Contract.AuctionIdxMap(&_KolKeys.CallOpts, arg0)
}

// AuctionIdxMap is a free data retrieval call binding the contract method 0x794f3bf1.
//
// Solidity: function auctionIdxMap(uint256 ) view returns(uint256)
func (_KolKeys *KolKeysCallerSession) AuctionIdxMap(arg0 *big.Int) (*big.Int, error) {
	return _KolKeys.Contract.AuctionIdxMap(&_KolKeys.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_KolKeys *KolKeysCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_KolKeys *KolKeysSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _KolKeys.Contract.Balance(&_KolKeys.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_KolKeys *KolKeysCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _KolKeys.Contract.Balance(&_KolKeys.CallOpts, arg0)
}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_KolKeys *KolKeysCaller) CheckPluginInterface(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "checkPluginInterface")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_KolKeys *KolKeysSession) CheckPluginInterface() (bool, error) {
	return _KolKeys.Contract.CheckPluginInterface(&_KolKeys.CallOpts)
}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_KolKeys *KolKeysCallerSession) CheckPluginInterface() (bool, error) {
	return _KolKeys.Contract.CheckPluginInterface(&_KolKeys.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KolKeys *KolKeysCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KolKeys *KolKeysSession) GetOwner() (common.Address, error) {
	return _KolKeys.Contract.GetOwner(&_KolKeys.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KolKeys *KolKeysCallerSession) GetOwner() (common.Address, error) {
	return _KolKeys.Contract.GetOwner(&_KolKeys.CallOpts)
}

// IncomePerNoncePerKey is a free data retrieval call binding the contract method 0x53311742.
//
// Solidity: function incomePerNoncePerKey(address , uint256 ) view returns(uint256)
func (_KolKeys *KolKeysCaller) IncomePerNoncePerKey(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "incomePerNoncePerKey", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomePerNoncePerKey is a free data retrieval call binding the contract method 0x53311742.
//
// Solidity: function incomePerNoncePerKey(address , uint256 ) view returns(uint256)
func (_KolKeys *KolKeysSession) IncomePerNoncePerKey(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _KolKeys.Contract.IncomePerNoncePerKey(&_KolKeys.CallOpts, arg0, arg1)
}

// IncomePerNoncePerKey is a free data retrieval call binding the contract method 0x53311742.
//
// Solidity: function incomePerNoncePerKey(address , uint256 ) view returns(uint256)
func (_KolKeys *KolKeysCallerSession) IncomePerNoncePerKey(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _KolKeys.Contract.IncomePerNoncePerKey(&_KolKeys.CallOpts, arg0, arg1)
}

// KeyInAuction is a free data retrieval call binding the contract method 0xab3da16f.
//
// Solidity: function keyInAuction(uint256 ) view returns(uint256 id, address seller, address bidder, address kol, uint256 nonce, uint256 amount, uint256 price, bool status)
func (_KolKeys *KolKeysCaller) KeyInAuction(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Seller common.Address
	Bidder common.Address
	Kol    common.Address
	Nonce  *big.Int
	Amount *big.Int
	Price  *big.Int
	Status bool
}, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "keyInAuction", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Seller common.Address
		Bidder common.Address
		Kol    common.Address
		Nonce  *big.Int
		Amount *big.Int
		Price  *big.Int
		Status bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Bidder = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Kol = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Nonce = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// KeyInAuction is a free data retrieval call binding the contract method 0xab3da16f.
//
// Solidity: function keyInAuction(uint256 ) view returns(uint256 id, address seller, address bidder, address kol, uint256 nonce, uint256 amount, uint256 price, bool status)
func (_KolKeys *KolKeysSession) KeyInAuction(arg0 *big.Int) (struct {
	Id     *big.Int
	Seller common.Address
	Bidder common.Address
	Kol    common.Address
	Nonce  *big.Int
	Amount *big.Int
	Price  *big.Int
	Status bool
}, error) {
	return _KolKeys.Contract.KeyInAuction(&_KolKeys.CallOpts, arg0)
}

// KeyInAuction is a free data retrieval call binding the contract method 0xab3da16f.
//
// Solidity: function keyInAuction(uint256 ) view returns(uint256 id, address seller, address bidder, address kol, uint256 nonce, uint256 amount, uint256 price, bool status)
func (_KolKeys *KolKeysCallerSession) KeyInAuction(arg0 *big.Int) (struct {
	Id     *big.Int
	Seller common.Address
	Bidder common.Address
	Kol    common.Address
	Nonce  *big.Int
	Amount *big.Int
	Price  *big.Int
	Status bool
}, error) {
	return _KolKeys.Contract.KeyInAuction(&_KolKeys.CallOpts, arg0)
}

// KolIncomeRatePerKeyBuy is a free data retrieval call binding the contract method 0xafb337cf.
//
// Solidity: function kolIncomeRatePerKeyBuy() view returns(uint8)
func (_KolKeys *KolKeysCaller) KolIncomeRatePerKeyBuy(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "kolIncomeRatePerKeyBuy")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KolIncomeRatePerKeyBuy is a free data retrieval call binding the contract method 0xafb337cf.
//
// Solidity: function kolIncomeRatePerKeyBuy() view returns(uint8)
func (_KolKeys *KolKeysSession) KolIncomeRatePerKeyBuy() (uint8, error) {
	return _KolKeys.Contract.KolIncomeRatePerKeyBuy(&_KolKeys.CallOpts)
}

// KolIncomeRatePerKeyBuy is a free data retrieval call binding the contract method 0xafb337cf.
//
// Solidity: function kolIncomeRatePerKeyBuy() view returns(uint8)
func (_KolKeys *KolKeysCallerSession) KolIncomeRatePerKeyBuy() (uint8, error) {
	return _KolKeys.Contract.KolIncomeRatePerKeyBuy(&_KolKeys.CallOpts)
}

// ServiceFeeRatePerKeyBuy is a free data retrieval call binding the contract method 0xaf707457.
//
// Solidity: function serviceFeeRatePerKeyBuy() view returns(uint8)
func (_KolKeys *KolKeysCaller) ServiceFeeRatePerKeyBuy(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "serviceFeeRatePerKeyBuy")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ServiceFeeRatePerKeyBuy is a free data retrieval call binding the contract method 0xaf707457.
//
// Solidity: function serviceFeeRatePerKeyBuy() view returns(uint8)
func (_KolKeys *KolKeysSession) ServiceFeeRatePerKeyBuy() (uint8, error) {
	return _KolKeys.Contract.ServiceFeeRatePerKeyBuy(&_KolKeys.CallOpts)
}

// ServiceFeeRatePerKeyBuy is a free data retrieval call binding the contract method 0xaf707457.
//
// Solidity: function serviceFeeRatePerKeyBuy() view returns(uint8)
func (_KolKeys *KolKeysCallerSession) ServiceFeeRatePerKeyBuy() (uint8, error) {
	return _KolKeys.Contract.ServiceFeeRatePerKeyBuy(&_KolKeys.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_KolKeys *KolKeysCaller) ServiceFeeReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "serviceFeeReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_KolKeys *KolKeysSession) ServiceFeeReceived() (*big.Int, error) {
	return _KolKeys.Contract.ServiceFeeReceived(&_KolKeys.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_KolKeys *KolKeysCallerSession) ServiceFeeReceived() (*big.Int, error) {
	return _KolKeys.Contract.ServiceFeeReceived(&_KolKeys.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_KolKeys *KolKeysCaller) WithdrawFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KolKeys.contract.Call(opts, &out, "withdrawFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_KolKeys *KolKeysSession) WithdrawFeeRate() (*big.Int, error) {
	return _KolKeys.Contract.WithdrawFeeRate(&_KolKeys.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_KolKeys *KolKeysCallerSession) WithdrawFeeRate() (*big.Int, error) {
	return _KolKeys.Contract.WithdrawFeeRate(&_KolKeys.CallOpts)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_KolKeys *KolKeysTransactor) AdminOperation(opts *bind.TransactOpts, admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminOperation", admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_KolKeys *KolKeysSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminOperation(&_KolKeys.TransactOpts, admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_KolKeys *KolKeysTransactorSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminOperation(&_KolKeys.TransactOpts, admin, isDelete)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_KolKeys *KolKeysTransactor) AdminServiceFeeWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminServiceFeeWithdraw")
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_KolKeys *KolKeysSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _KolKeys.Contract.AdminServiceFeeWithdraw(&_KolKeys.TransactOpts)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_KolKeys *KolKeysTransactorSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _KolKeys.Contract.AdminServiceFeeWithdraw(&_KolKeys.TransactOpts)
}

// AdminSetBddingFee is a paid mutator transaction binding the contract method 0x388a82cc.
//
// Solidity: function adminSetBddingFee(uint256 newFeeInGwei) returns()
func (_KolKeys *KolKeysTransactor) AdminSetBddingFee(opts *bind.TransactOpts, newFeeInGwei *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminSetBddingFee", newFeeInGwei)
}

// AdminSetBddingFee is a paid mutator transaction binding the contract method 0x388a82cc.
//
// Solidity: function adminSetBddingFee(uint256 newFeeInGwei) returns()
func (_KolKeys *KolKeysSession) AdminSetBddingFee(newFeeInGwei *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetBddingFee(&_KolKeys.TransactOpts, newFeeInGwei)
}

// AdminSetBddingFee is a paid mutator transaction binding the contract method 0x388a82cc.
//
// Solidity: function adminSetBddingFee(uint256 newFeeInGwei) returns()
func (_KolKeys *KolKeysTransactorSession) AdminSetBddingFee(newFeeInGwei *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetBddingFee(&_KolKeys.TransactOpts, newFeeInGwei)
}

// AdminSetKeyFeeRate is a paid mutator transaction binding the contract method 0x84854688.
//
// Solidity: function adminSetKeyFeeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysTransactor) AdminSetKeyFeeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminSetKeyFeeRate", newRate)
}

// AdminSetKeyFeeRate is a paid mutator transaction binding the contract method 0x84854688.
//
// Solidity: function adminSetKeyFeeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysSession) AdminSetKeyFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetKeyFeeRate(&_KolKeys.TransactOpts, newRate)
}

// AdminSetKeyFeeRate is a paid mutator transaction binding the contract method 0x84854688.
//
// Solidity: function adminSetKeyFeeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysTransactorSession) AdminSetKeyFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetKeyFeeRate(&_KolKeys.TransactOpts, newRate)
}

// AdminSetKolIncomeRate is a paid mutator transaction binding the contract method 0x34afeec0.
//
// Solidity: function adminSetKolIncomeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysTransactor) AdminSetKolIncomeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminSetKolIncomeRate", newRate)
}

// AdminSetKolIncomeRate is a paid mutator transaction binding the contract method 0x34afeec0.
//
// Solidity: function adminSetKolIncomeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysSession) AdminSetKolIncomeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetKolIncomeRate(&_KolKeys.TransactOpts, newRate)
}

// AdminSetKolIncomeRate is a paid mutator transaction binding the contract method 0x34afeec0.
//
// Solidity: function adminSetKolIncomeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysTransactorSession) AdminSetKolIncomeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetKolIncomeRate(&_KolKeys.TransactOpts, newRate)
}

// AdminSetKolOperationFee is a paid mutator transaction binding the contract method 0xe9e60559.
//
// Solidity: function adminSetKolOperationFee(uint256 newFeeInGwei) returns()
func (_KolKeys *KolKeysTransactor) AdminSetKolOperationFee(opts *bind.TransactOpts, newFeeInGwei *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminSetKolOperationFee", newFeeInGwei)
}

// AdminSetKolOperationFee is a paid mutator transaction binding the contract method 0xe9e60559.
//
// Solidity: function adminSetKolOperationFee(uint256 newFeeInGwei) returns()
func (_KolKeys *KolKeysSession) AdminSetKolOperationFee(newFeeInGwei *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetKolOperationFee(&_KolKeys.TransactOpts, newFeeInGwei)
}

// AdminSetKolOperationFee is a paid mutator transaction binding the contract method 0xe9e60559.
//
// Solidity: function adminSetKolOperationFee(uint256 newFeeInGwei) returns()
func (_KolKeys *KolKeysTransactorSession) AdminSetKolOperationFee(newFeeInGwei *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetKolOperationFee(&_KolKeys.TransactOpts, newFeeInGwei)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysTransactor) AdminSetWithdrawFeeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminSetWithdrawFeeRate", newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetWithdrawFeeRate(&_KolKeys.TransactOpts, newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_KolKeys *KolKeysTransactorSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminSetWithdrawFeeRate(&_KolKeys.TransactOpts, newRate)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_KolKeys *KolKeysTransactor) AdminUpgradeToNewRule(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "adminUpgradeToNewRule", recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_KolKeys *KolKeysSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminUpgradeToNewRule(&_KolKeys.TransactOpts, recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_KolKeys *KolKeysTransactorSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.AdminUpgradeToNewRule(&_KolKeys.TransactOpts, recipient)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 bidID) payable returns()
func (_KolKeys *KolKeysTransactor) Bid(opts *bind.TransactOpts, bidID *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "bid", bidID)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 bidID) payable returns()
func (_KolKeys *KolKeysSession) Bid(bidID *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.Bid(&_KolKeys.TransactOpts, bidID)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 bidID) payable returns()
func (_KolKeys *KolKeysTransactorSession) Bid(bidID *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.Bid(&_KolKeys.TransactOpts, bidID)
}

// BuyKolKey is a paid mutator transaction binding the contract method 0xc7cdf93c.
//
// Solidity: function buyKolKey(address kolAddr, uint256 keyNo) payable returns()
func (_KolKeys *KolKeysTransactor) BuyKolKey(opts *bind.TransactOpts, kolAddr common.Address, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "buyKolKey", kolAddr, keyNo)
}

// BuyKolKey is a paid mutator transaction binding the contract method 0xc7cdf93c.
//
// Solidity: function buyKolKey(address kolAddr, uint256 keyNo) payable returns()
func (_KolKeys *KolKeysSession) BuyKolKey(kolAddr common.Address, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.BuyKolKey(&_KolKeys.TransactOpts, kolAddr, keyNo)
}

// BuyKolKey is a paid mutator transaction binding the contract method 0xc7cdf93c.
//
// Solidity: function buyKolKey(address kolAddr, uint256 keyNo) payable returns()
func (_KolKeys *KolKeysTransactorSession) BuyKolKey(kolAddr common.Address, keyNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.BuyKolKey(&_KolKeys.TransactOpts, kolAddr, keyNo)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x64d13339.
//
// Solidity: function changeOwner(bool stop) returns()
func (_KolKeys *KolKeysTransactor) ChangeOwner(opts *bind.TransactOpts, stop bool) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "changeOwner", stop)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x64d13339.
//
// Solidity: function changeOwner(bool stop) returns()
func (_KolKeys *KolKeysSession) ChangeOwner(stop bool) (*types.Transaction, error) {
	return _KolKeys.Contract.ChangeOwner(&_KolKeys.TransactOpts, stop)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x64d13339.
//
// Solidity: function changeOwner(bool stop) returns()
func (_KolKeys *KolKeysTransactorSession) ChangeOwner(stop bool) (*types.Transaction, error) {
	return _KolKeys.Contract.ChangeOwner(&_KolKeys.TransactOpts, stop)
}

// ChangeOwner0 is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_KolKeys *KolKeysTransactor) ChangeOwner0(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "changeOwner0", newOwner)
}

// ChangeOwner0 is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_KolKeys *KolKeysSession) ChangeOwner0(newOwner common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.ChangeOwner0(&_KolKeys.TransactOpts, newOwner)
}

// ChangeOwner0 is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_KolKeys *KolKeysTransactorSession) ChangeOwner0(newOwner common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.ChangeOwner0(&_KolKeys.TransactOpts, newOwner)
}

// IncomeToKolPool is a paid mutator transaction binding the contract method 0x2b915094.
//
// Solidity: function incomeToKolPool(address kol) payable returns()
func (_KolKeys *KolKeysTransactor) IncomeToKolPool(opts *bind.TransactOpts, kol common.Address) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "incomeToKolPool", kol)
}

// IncomeToKolPool is a paid mutator transaction binding the contract method 0x2b915094.
//
// Solidity: function incomeToKolPool(address kol) payable returns()
func (_KolKeys *KolKeysSession) IncomeToKolPool(kol common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.IncomeToKolPool(&_KolKeys.TransactOpts, kol)
}

// IncomeToKolPool is a paid mutator transaction binding the contract method 0x2b915094.
//
// Solidity: function incomeToKolPool(address kol) payable returns()
func (_KolKeys *KolKeysTransactorSession) IncomeToKolPool(kol common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.IncomeToKolPool(&_KolKeys.TransactOpts, kol)
}

// IssueBid is a paid mutator transaction binding the contract method 0x627d27b1.
//
// Solidity: function issueBid(address kol, uint256 nonce, uint256 amount, uint256 pricePerKey) payable returns()
func (_KolKeys *KolKeysTransactor) IssueBid(opts *bind.TransactOpts, kol common.Address, nonce *big.Int, amount *big.Int, pricePerKey *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "issueBid", kol, nonce, amount, pricePerKey)
}

// IssueBid is a paid mutator transaction binding the contract method 0x627d27b1.
//
// Solidity: function issueBid(address kol, uint256 nonce, uint256 amount, uint256 pricePerKey) payable returns()
func (_KolKeys *KolKeysSession) IssueBid(kol common.Address, nonce *big.Int, amount *big.Int, pricePerKey *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.IssueBid(&_KolKeys.TransactOpts, kol, nonce, amount, pricePerKey)
}

// IssueBid is a paid mutator transaction binding the contract method 0x627d27b1.
//
// Solidity: function issueBid(address kol, uint256 nonce, uint256 amount, uint256 pricePerKey) payable returns()
func (_KolKeys *KolKeysTransactorSession) IssueBid(kol common.Address, nonce *big.Int, amount *big.Int, pricePerKey *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.IssueBid(&_KolKeys.TransactOpts, kol, nonce, amount, pricePerKey)
}

// KolAddKeySupply is a paid mutator transaction binding the contract method 0xd7eca559.
//
// Solidity: function kolAddKeySupply(uint256 amount) payable returns()
func (_KolKeys *KolKeysTransactor) KolAddKeySupply(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "kolAddKeySupply", amount)
}

// KolAddKeySupply is a paid mutator transaction binding the contract method 0xd7eca559.
//
// Solidity: function kolAddKeySupply(uint256 amount) payable returns()
func (_KolKeys *KolKeysSession) KolAddKeySupply(amount *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.KolAddKeySupply(&_KolKeys.TransactOpts, amount)
}

// KolAddKeySupply is a paid mutator transaction binding the contract method 0xd7eca559.
//
// Solidity: function kolAddKeySupply(uint256 amount) payable returns()
func (_KolKeys *KolKeysTransactorSession) KolAddKeySupply(amount *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.KolAddKeySupply(&_KolKeys.TransactOpts, amount)
}

// KolOpenKeySale is a paid mutator transaction binding the contract method 0x4a26e8ea.
//
// Solidity: function kolOpenKeySale(uint256 pricePerKey, int256 maxKeyNo) payable returns()
func (_KolKeys *KolKeysTransactor) KolOpenKeySale(opts *bind.TransactOpts, pricePerKey *big.Int, maxKeyNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "kolOpenKeySale", pricePerKey, maxKeyNo)
}

// KolOpenKeySale is a paid mutator transaction binding the contract method 0x4a26e8ea.
//
// Solidity: function kolOpenKeySale(uint256 pricePerKey, int256 maxKeyNo) payable returns()
func (_KolKeys *KolKeysSession) KolOpenKeySale(pricePerKey *big.Int, maxKeyNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.KolOpenKeySale(&_KolKeys.TransactOpts, pricePerKey, maxKeyNo)
}

// KolOpenKeySale is a paid mutator transaction binding the contract method 0x4a26e8ea.
//
// Solidity: function kolOpenKeySale(uint256 pricePerKey, int256 maxKeyNo) payable returns()
func (_KolKeys *KolKeysTransactorSession) KolOpenKeySale(pricePerKey *big.Int, maxKeyNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.KolOpenKeySale(&_KolKeys.TransactOpts, pricePerKey, maxKeyNo)
}

// KolSetKeyPrice is a paid mutator transaction binding the contract method 0xdb003fc7.
//
// Solidity: function kolSetKeyPrice(uint256 newPrice) payable returns()
func (_KolKeys *KolKeysTransactor) KolSetKeyPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "kolSetKeyPrice", newPrice)
}

// KolSetKeyPrice is a paid mutator transaction binding the contract method 0xdb003fc7.
//
// Solidity: function kolSetKeyPrice(uint256 newPrice) payable returns()
func (_KolKeys *KolKeysSession) KolSetKeyPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.KolSetKeyPrice(&_KolKeys.TransactOpts, newPrice)
}

// KolSetKeyPrice is a paid mutator transaction binding the contract method 0xdb003fc7.
//
// Solidity: function kolSetKeyPrice(uint256 newPrice) payable returns()
func (_KolKeys *KolKeysTransactorSession) KolSetKeyPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.KolSetKeyPrice(&_KolKeys.TransactOpts, newPrice)
}

// RevokeBid is a paid mutator transaction binding the contract method 0xd1a1f09a.
//
// Solidity: function revokeBid(uint256 bidID) payable returns()
func (_KolKeys *KolKeysTransactor) RevokeBid(opts *bind.TransactOpts, bidID *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "revokeBid", bidID)
}

// RevokeBid is a paid mutator transaction binding the contract method 0xd1a1f09a.
//
// Solidity: function revokeBid(uint256 bidID) payable returns()
func (_KolKeys *KolKeysSession) RevokeBid(bidID *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.RevokeBid(&_KolKeys.TransactOpts, bidID)
}

// RevokeBid is a paid mutator transaction binding the contract method 0xd1a1f09a.
//
// Solidity: function revokeBid(uint256 bidID) payable returns()
func (_KolKeys *KolKeysTransactorSession) RevokeBid(bidID *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.RevokeBid(&_KolKeys.TransactOpts, bidID)
}

// TransferAllKey is a paid mutator transaction binding the contract method 0x01c6c8da.
//
// Solidity: function transferAllKey(address to, address kol) returns()
func (_KolKeys *KolKeysTransactor) TransferAllKey(opts *bind.TransactOpts, to common.Address, kol common.Address) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "transferAllKey", to, kol)
}

// TransferAllKey is a paid mutator transaction binding the contract method 0x01c6c8da.
//
// Solidity: function transferAllKey(address to, address kol) returns()
func (_KolKeys *KolKeysSession) TransferAllKey(to common.Address, kol common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.TransferAllKey(&_KolKeys.TransactOpts, to, kol)
}

// TransferAllKey is a paid mutator transaction binding the contract method 0x01c6c8da.
//
// Solidity: function transferAllKey(address to, address kol) returns()
func (_KolKeys *KolKeysTransactorSession) TransferAllKey(to common.Address, kol common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.TransferAllKey(&_KolKeys.TransactOpts, to, kol)
}

// TransferKey is a paid mutator transaction binding the contract method 0x7b46bfe8.
//
// Solidity: function transferKey(address to, address kol, uint256 nonce, uint256 amount) returns()
func (_KolKeys *KolKeysTransactor) TransferKey(opts *bind.TransactOpts, to common.Address, kol common.Address, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "transferKey", to, kol, nonce, amount)
}

// TransferKey is a paid mutator transaction binding the contract method 0x7b46bfe8.
//
// Solidity: function transferKey(address to, address kol, uint256 nonce, uint256 amount) returns()
func (_KolKeys *KolKeysSession) TransferKey(to common.Address, kol common.Address, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.TransferKey(&_KolKeys.TransactOpts, to, kol, nonce, amount)
}

// TransferKey is a paid mutator transaction binding the contract method 0x7b46bfe8.
//
// Solidity: function transferKey(address to, address kol, uint256 nonce, uint256 amount) returns()
func (_KolKeys *KolKeysTransactorSession) TransferKey(to common.Address, kol common.Address, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.TransferKey(&_KolKeys.TransactOpts, to, kol, nonce, amount)
}

// TweetBought is a paid mutator transaction binding the contract method 0xd2f2eeb7.
//
// Solidity: function tweetBought(bytes32 tweetHash, address tweetOwner, address buyer, uint256 voteNo) payable returns()
func (_KolKeys *KolKeysTransactor) TweetBought(opts *bind.TransactOpts, tweetHash [32]byte, tweetOwner common.Address, buyer common.Address, voteNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "tweetBought", tweetHash, tweetOwner, buyer, voteNo)
}

// TweetBought is a paid mutator transaction binding the contract method 0xd2f2eeb7.
//
// Solidity: function tweetBought(bytes32 tweetHash, address tweetOwner, address buyer, uint256 voteNo) payable returns()
func (_KolKeys *KolKeysSession) TweetBought(tweetHash [32]byte, tweetOwner common.Address, buyer common.Address, voteNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.TweetBought(&_KolKeys.TransactOpts, tweetHash, tweetOwner, buyer, voteNo)
}

// TweetBought is a paid mutator transaction binding the contract method 0xd2f2eeb7.
//
// Solidity: function tweetBought(bytes32 tweetHash, address tweetOwner, address buyer, uint256 voteNo) payable returns()
func (_KolKeys *KolKeysTransactorSession) TweetBought(tweetHash [32]byte, tweetOwner common.Address, buyer common.Address, voteNo *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.TweetBought(&_KolKeys.TransactOpts, tweetHash, tweetOwner, buyer, voteNo)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_KolKeys *KolKeysTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, all bool) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "withdraw", amount, all)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_KolKeys *KolKeysSession) Withdraw(amount *big.Int, all bool) (*types.Transaction, error) {
	return _KolKeys.Contract.Withdraw(&_KolKeys.TransactOpts, amount, all)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_KolKeys *KolKeysTransactorSession) Withdraw(amount *big.Int, all bool) (*types.Transaction, error) {
	return _KolKeys.Contract.Withdraw(&_KolKeys.TransactOpts, amount, all)
}

// WithdrawAllIncome is a paid mutator transaction binding the contract method 0x0c3e0452.
//
// Solidity: function withdrawAllIncome() returns()
func (_KolKeys *KolKeysTransactor) WithdrawAllIncome(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "withdrawAllIncome")
}

// WithdrawAllIncome is a paid mutator transaction binding the contract method 0x0c3e0452.
//
// Solidity: function withdrawAllIncome() returns()
func (_KolKeys *KolKeysSession) WithdrawAllIncome() (*types.Transaction, error) {
	return _KolKeys.Contract.WithdrawAllIncome(&_KolKeys.TransactOpts)
}

// WithdrawAllIncome is a paid mutator transaction binding the contract method 0x0c3e0452.
//
// Solidity: function withdrawAllIncome() returns()
func (_KolKeys *KolKeysTransactorSession) WithdrawAllIncome() (*types.Transaction, error) {
	return _KolKeys.Contract.WithdrawAllIncome(&_KolKeys.TransactOpts)
}

// WithdrawFromOneKol is a paid mutator transaction binding the contract method 0x0ef86b3f.
//
// Solidity: function withdrawFromOneKol(address kol) returns()
func (_KolKeys *KolKeysTransactor) WithdrawFromOneKol(opts *bind.TransactOpts, kol common.Address) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "withdrawFromOneKol", kol)
}

// WithdrawFromOneKol is a paid mutator transaction binding the contract method 0x0ef86b3f.
//
// Solidity: function withdrawFromOneKol(address kol) returns()
func (_KolKeys *KolKeysSession) WithdrawFromOneKol(kol common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.WithdrawFromOneKol(&_KolKeys.TransactOpts, kol)
}

// WithdrawFromOneKol is a paid mutator transaction binding the contract method 0x0ef86b3f.
//
// Solidity: function withdrawFromOneKol(address kol) returns()
func (_KolKeys *KolKeysTransactorSession) WithdrawFromOneKol(kol common.Address) (*types.Transaction, error) {
	return _KolKeys.Contract.WithdrawFromOneKol(&_KolKeys.TransactOpts, kol)
}

// WithdrawFromOneNonce is a paid mutator transaction binding the contract method 0x3da6136e.
//
// Solidity: function withdrawFromOneNonce(address kol, uint256 nonce) returns()
func (_KolKeys *KolKeysTransactor) WithdrawFromOneNonce(opts *bind.TransactOpts, kol common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _KolKeys.contract.Transact(opts, "withdrawFromOneNonce", kol, nonce)
}

// WithdrawFromOneNonce is a paid mutator transaction binding the contract method 0x3da6136e.
//
// Solidity: function withdrawFromOneNonce(address kol, uint256 nonce) returns()
func (_KolKeys *KolKeysSession) WithdrawFromOneNonce(kol common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.WithdrawFromOneNonce(&_KolKeys.TransactOpts, kol, nonce)
}

// WithdrawFromOneNonce is a paid mutator transaction binding the contract method 0x3da6136e.
//
// Solidity: function withdrawFromOneNonce(address kol, uint256 nonce) returns()
func (_KolKeys *KolKeysTransactorSession) WithdrawFromOneNonce(kol common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _KolKeys.Contract.WithdrawFromOneNonce(&_KolKeys.TransactOpts, kol, nonce)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KolKeys *KolKeysTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KolKeys.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KolKeys *KolKeysSession) Receive() (*types.Transaction, error) {
	return _KolKeys.Contract.Receive(&_KolKeys.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KolKeys *KolKeysTransactorSession) Receive() (*types.Transaction, error) {
	return _KolKeys.Contract.Receive(&_KolKeys.TransactOpts)
}

// KolKeysAdminOperationIterator is returned from FilterAdminOperation and is used to iterate over the raw logs and unpacked data for AdminOperation events raised by the KolKeys contract.
type KolKeysAdminOperationIterator struct {
	Event *KolKeysAdminOperation // Event containing the contract specifics and raw log

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
func (it *KolKeysAdminOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysAdminOperation)
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
		it.Event = new(KolKeysAdminOperation)
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
func (it *KolKeysAdminOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysAdminOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysAdminOperation represents a AdminOperation event raised by the KolKeys contract.
type KolKeysAdminOperation struct {
	Admin  common.Address
	OpType bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAdminOperation is a free log retrieval operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_KolKeys *KolKeysFilterer) FilterAdminOperation(opts *bind.FilterOpts) (*KolKeysAdminOperationIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return &KolKeysAdminOperationIterator{contract: _KolKeys.contract, event: "AdminOperation", logs: logs, sub: sub}, nil
}

// WatchAdminOperation is a free log subscription operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_KolKeys *KolKeysFilterer) WatchAdminOperation(opts *bind.WatchOpts, sink chan<- *KolKeysAdminOperation) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysAdminOperation)
				if err := _KolKeys.contract.UnpackLog(event, "AdminOperation", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseAdminOperation(log types.Log) (*KolKeysAdminOperation, error) {
	event := new(KolKeysAdminOperation)
	if err := _KolKeys.contract.UnpackLog(event, "AdminOperation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysInvestorWithdrawAllIncomeIterator is returned from FilterInvestorWithdrawAllIncome and is used to iterate over the raw logs and unpacked data for InvestorWithdrawAllIncome events raised by the KolKeys contract.
type KolKeysInvestorWithdrawAllIncomeIterator struct {
	Event *KolKeysInvestorWithdrawAllIncome // Event containing the contract specifics and raw log

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
func (it *KolKeysInvestorWithdrawAllIncomeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysInvestorWithdrawAllIncome)
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
		it.Event = new(KolKeysInvestorWithdrawAllIncome)
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
func (it *KolKeysInvestorWithdrawAllIncomeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysInvestorWithdrawAllIncomeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysInvestorWithdrawAllIncome represents a InvestorWithdrawAllIncome event raised by the KolKeys contract.
type KolKeysInvestorWithdrawAllIncome struct {
	Investor common.Address
	Val      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdrawAllIncome is a free log retrieval operation binding the contract event 0x5dbe3bbfa87813b8be328cdb0b812e04de36cf36d1141967e1df795a7c95db0a.
//
// Solidity: event InvestorWithdrawAllIncome(address investor, uint256 val)
func (_KolKeys *KolKeysFilterer) FilterInvestorWithdrawAllIncome(opts *bind.FilterOpts) (*KolKeysInvestorWithdrawAllIncomeIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "InvestorWithdrawAllIncome")
	if err != nil {
		return nil, err
	}
	return &KolKeysInvestorWithdrawAllIncomeIterator{contract: _KolKeys.contract, event: "InvestorWithdrawAllIncome", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdrawAllIncome is a free log subscription operation binding the contract event 0x5dbe3bbfa87813b8be328cdb0b812e04de36cf36d1141967e1df795a7c95db0a.
//
// Solidity: event InvestorWithdrawAllIncome(address investor, uint256 val)
func (_KolKeys *KolKeysFilterer) WatchInvestorWithdrawAllIncome(opts *bind.WatchOpts, sink chan<- *KolKeysInvestorWithdrawAllIncome) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "InvestorWithdrawAllIncome")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysInvestorWithdrawAllIncome)
				if err := _KolKeys.contract.UnpackLog(event, "InvestorWithdrawAllIncome", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseInvestorWithdrawAllIncome(log types.Log) (*KolKeysInvestorWithdrawAllIncome, error) {
	event := new(KolKeysInvestorWithdrawAllIncome)
	if err := _KolKeys.contract.UnpackLog(event, "InvestorWithdrawAllIncome", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysInvestorWithdrawByOneKolIterator is returned from FilterInvestorWithdrawByOneKol and is used to iterate over the raw logs and unpacked data for InvestorWithdrawByOneKol events raised by the KolKeys contract.
type KolKeysInvestorWithdrawByOneKolIterator struct {
	Event *KolKeysInvestorWithdrawByOneKol // Event containing the contract specifics and raw log

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
func (it *KolKeysInvestorWithdrawByOneKolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysInvestorWithdrawByOneKol)
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
		it.Event = new(KolKeysInvestorWithdrawByOneKol)
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
func (it *KolKeysInvestorWithdrawByOneKolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysInvestorWithdrawByOneKolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysInvestorWithdrawByOneKol represents a InvestorWithdrawByOneKol event raised by the KolKeys contract.
type KolKeysInvestorWithdrawByOneKol struct {
	Investor common.Address
	Kol      common.Address
	Val      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdrawByOneKol is a free log retrieval operation binding the contract event 0x6ce7ea243972341c9e8ddfa93ba573fec9d104862f00b1bcd6aab059efbc9a6d.
//
// Solidity: event InvestorWithdrawByOneKol(address investor, address kol, uint256 val)
func (_KolKeys *KolKeysFilterer) FilterInvestorWithdrawByOneKol(opts *bind.FilterOpts) (*KolKeysInvestorWithdrawByOneKolIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "InvestorWithdrawByOneKol")
	if err != nil {
		return nil, err
	}
	return &KolKeysInvestorWithdrawByOneKolIterator{contract: _KolKeys.contract, event: "InvestorWithdrawByOneKol", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdrawByOneKol is a free log subscription operation binding the contract event 0x6ce7ea243972341c9e8ddfa93ba573fec9d104862f00b1bcd6aab059efbc9a6d.
//
// Solidity: event InvestorWithdrawByOneKol(address investor, address kol, uint256 val)
func (_KolKeys *KolKeysFilterer) WatchInvestorWithdrawByOneKol(opts *bind.WatchOpts, sink chan<- *KolKeysInvestorWithdrawByOneKol) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "InvestorWithdrawByOneKol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysInvestorWithdrawByOneKol)
				if err := _KolKeys.contract.UnpackLog(event, "InvestorWithdrawByOneKol", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseInvestorWithdrawByOneKol(log types.Log) (*KolKeysInvestorWithdrawByOneKol, error) {
	event := new(KolKeysInvestorWithdrawByOneKol)
	if err := _KolKeys.contract.UnpackLog(event, "InvestorWithdrawByOneKol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysInvestorWithdrawByOneNonceIterator is returned from FilterInvestorWithdrawByOneNonce and is used to iterate over the raw logs and unpacked data for InvestorWithdrawByOneNonce events raised by the KolKeys contract.
type KolKeysInvestorWithdrawByOneNonceIterator struct {
	Event *KolKeysInvestorWithdrawByOneNonce // Event containing the contract specifics and raw log

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
func (it *KolKeysInvestorWithdrawByOneNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysInvestorWithdrawByOneNonce)
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
		it.Event = new(KolKeysInvestorWithdrawByOneNonce)
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
func (it *KolKeysInvestorWithdrawByOneNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysInvestorWithdrawByOneNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysInvestorWithdrawByOneNonce represents a InvestorWithdrawByOneNonce event raised by the KolKeys contract.
type KolKeysInvestorWithdrawByOneNonce struct {
	Investor common.Address
	Kol      common.Address
	Nonce    *big.Int
	Val      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdrawByOneNonce is a free log retrieval operation binding the contract event 0xc85771d86e5aae3490e851850d1f22697f3a3fc112709b5d2bea1c3ea84fcefc.
//
// Solidity: event InvestorWithdrawByOneNonce(address investor, address kol, uint256 nonce, uint256 val)
func (_KolKeys *KolKeysFilterer) FilterInvestorWithdrawByOneNonce(opts *bind.FilterOpts) (*KolKeysInvestorWithdrawByOneNonceIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "InvestorWithdrawByOneNonce")
	if err != nil {
		return nil, err
	}
	return &KolKeysInvestorWithdrawByOneNonceIterator{contract: _KolKeys.contract, event: "InvestorWithdrawByOneNonce", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdrawByOneNonce is a free log subscription operation binding the contract event 0xc85771d86e5aae3490e851850d1f22697f3a3fc112709b5d2bea1c3ea84fcefc.
//
// Solidity: event InvestorWithdrawByOneNonce(address investor, address kol, uint256 nonce, uint256 val)
func (_KolKeys *KolKeysFilterer) WatchInvestorWithdrawByOneNonce(opts *bind.WatchOpts, sink chan<- *KolKeysInvestorWithdrawByOneNonce) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "InvestorWithdrawByOneNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysInvestorWithdrawByOneNonce)
				if err := _KolKeys.contract.UnpackLog(event, "InvestorWithdrawByOneNonce", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseInvestorWithdrawByOneNonce(log types.Log) (*KolKeysInvestorWithdrawByOneNonce, error) {
	event := new(KolKeysInvestorWithdrawByOneNonce)
	if err := _KolKeys.contract.UnpackLog(event, "InvestorWithdrawByOneNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysKeyAuctionActionIterator is returned from FilterKeyAuctionAction and is used to iterate over the raw logs and unpacked data for KeyAuctionAction events raised by the KolKeys contract.
type KolKeysKeyAuctionActionIterator struct {
	Event *KolKeysKeyAuctionAction // Event containing the contract specifics and raw log

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
func (it *KolKeysKeyAuctionActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysKeyAuctionAction)
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
		it.Event = new(KolKeysKeyAuctionAction)
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
func (it *KolKeysKeyAuctionActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysKeyAuctionActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysKeyAuctionAction represents a KeyAuctionAction event raised by the KolKeys contract.
type KolKeysKeyAuctionAction struct {
	Kol         common.Address
	Seller      common.Address
	Buyer       common.Address
	Nonce       *big.Int
	Amount      *big.Int
	PricePerKey *big.Int
	Typ         string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKeyAuctionAction is a free log retrieval operation binding the contract event 0xa6b5fa844f538ba008c65f296b54934aff4c2f11f1a9d5ef89c5f535b95681a2.
//
// Solidity: event KeyAuctionAction(address kol, address seller, address buyer, uint256 nonce, uint256 amount, uint256 pricePerKey, string typ)
func (_KolKeys *KolKeysFilterer) FilterKeyAuctionAction(opts *bind.FilterOpts) (*KolKeysKeyAuctionActionIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "KeyAuctionAction")
	if err != nil {
		return nil, err
	}
	return &KolKeysKeyAuctionActionIterator{contract: _KolKeys.contract, event: "KeyAuctionAction", logs: logs, sub: sub}, nil
}

// WatchKeyAuctionAction is a free log subscription operation binding the contract event 0xa6b5fa844f538ba008c65f296b54934aff4c2f11f1a9d5ef89c5f535b95681a2.
//
// Solidity: event KeyAuctionAction(address kol, address seller, address buyer, uint256 nonce, uint256 amount, uint256 pricePerKey, string typ)
func (_KolKeys *KolKeysFilterer) WatchKeyAuctionAction(opts *bind.WatchOpts, sink chan<- *KolKeysKeyAuctionAction) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "KeyAuctionAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysKeyAuctionAction)
				if err := _KolKeys.contract.UnpackLog(event, "KeyAuctionAction", log); err != nil {
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

// ParseKeyAuctionAction is a log parse operation binding the contract event 0xa6b5fa844f538ba008c65f296b54934aff4c2f11f1a9d5ef89c5f535b95681a2.
//
// Solidity: event KeyAuctionAction(address kol, address seller, address buyer, uint256 nonce, uint256 amount, uint256 pricePerKey, string typ)
func (_KolKeys *KolKeysFilterer) ParseKeyAuctionAction(log types.Log) (*KolKeysKeyAuctionAction, error) {
	event := new(KolKeysKeyAuctionAction)
	if err := _KolKeys.contract.UnpackLog(event, "KeyAuctionAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysKeyTransferedIterator is returned from FilterKeyTransfered and is used to iterate over the raw logs and unpacked data for KeyTransfered events raised by the KolKeys contract.
type KolKeysKeyTransferedIterator struct {
	Event *KolKeysKeyTransfered // Event containing the contract specifics and raw log

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
func (it *KolKeysKeyTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysKeyTransfered)
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
		it.Event = new(KolKeysKeyTransfered)
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
func (it *KolKeysKeyTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysKeyTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysKeyTransfered represents a KeyTransfered event raised by the KolKeys contract.
type KolKeysKeyTransfered struct {
	From   common.Address
	To     common.Address
	Kol    common.Address
	Nonce  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeyTransfered is a free log retrieval operation binding the contract event 0xab9bb64eb94ba07b3652f3aac538e6365596995aabed6d8de6e33b2bcdf7a6ed.
//
// Solidity: event KeyTransfered(address from, address to, address kol, uint256 nonce, uint256 amount)
func (_KolKeys *KolKeysFilterer) FilterKeyTransfered(opts *bind.FilterOpts) (*KolKeysKeyTransferedIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "KeyTransfered")
	if err != nil {
		return nil, err
	}
	return &KolKeysKeyTransferedIterator{contract: _KolKeys.contract, event: "KeyTransfered", logs: logs, sub: sub}, nil
}

// WatchKeyTransfered is a free log subscription operation binding the contract event 0xab9bb64eb94ba07b3652f3aac538e6365596995aabed6d8de6e33b2bcdf7a6ed.
//
// Solidity: event KeyTransfered(address from, address to, address kol, uint256 nonce, uint256 amount)
func (_KolKeys *KolKeysFilterer) WatchKeyTransfered(opts *bind.WatchOpts, sink chan<- *KolKeysKeyTransfered) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "KeyTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysKeyTransfered)
				if err := _KolKeys.contract.UnpackLog(event, "KeyTransfered", log); err != nil {
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

// ParseKeyTransfered is a log parse operation binding the contract event 0xab9bb64eb94ba07b3652f3aac538e6365596995aabed6d8de6e33b2bcdf7a6ed.
//
// Solidity: event KeyTransfered(address from, address to, address kol, uint256 nonce, uint256 amount)
func (_KolKeys *KolKeysFilterer) ParseKeyTransfered(log types.Log) (*KolKeysKeyTransfered, error) {
	event := new(KolKeysKeyTransfered)
	if err := _KolKeys.contract.UnpackLog(event, "KeyTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysKeyTransferedAllIterator is returned from FilterKeyTransferedAll and is used to iterate over the raw logs and unpacked data for KeyTransferedAll events raised by the KolKeys contract.
type KolKeysKeyTransferedAllIterator struct {
	Event *KolKeysKeyTransferedAll // Event containing the contract specifics and raw log

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
func (it *KolKeysKeyTransferedAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysKeyTransferedAll)
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
		it.Event = new(KolKeysKeyTransferedAll)
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
func (it *KolKeysKeyTransferedAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysKeyTransferedAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysKeyTransferedAll represents a KeyTransferedAll event raised by the KolKeys contract.
type KolKeysKeyTransferedAll struct {
	From common.Address
	To   common.Address
	Kol  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyTransferedAll is a free log retrieval operation binding the contract event 0x67e33abe0f9a96e0749d2036016e153284d65989cd4875f4c5a42853b35e8bf1.
//
// Solidity: event KeyTransferedAll(address from, address to, address kol)
func (_KolKeys *KolKeysFilterer) FilterKeyTransferedAll(opts *bind.FilterOpts) (*KolKeysKeyTransferedAllIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "KeyTransferedAll")
	if err != nil {
		return nil, err
	}
	return &KolKeysKeyTransferedAllIterator{contract: _KolKeys.contract, event: "KeyTransferedAll", logs: logs, sub: sub}, nil
}

// WatchKeyTransferedAll is a free log subscription operation binding the contract event 0x67e33abe0f9a96e0749d2036016e153284d65989cd4875f4c5a42853b35e8bf1.
//
// Solidity: event KeyTransferedAll(address from, address to, address kol)
func (_KolKeys *KolKeysFilterer) WatchKeyTransferedAll(opts *bind.WatchOpts, sink chan<- *KolKeysKeyTransferedAll) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "KeyTransferedAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysKeyTransferedAll)
				if err := _KolKeys.contract.UnpackLog(event, "KeyTransferedAll", log); err != nil {
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

// ParseKeyTransferedAll is a log parse operation binding the contract event 0x67e33abe0f9a96e0749d2036016e153284d65989cd4875f4c5a42853b35e8bf1.
//
// Solidity: event KeyTransferedAll(address from, address to, address kol)
func (_KolKeys *KolKeysFilterer) ParseKeyTransferedAll(log types.Log) (*KolKeysKeyTransferedAll, error) {
	event := new(KolKeysKeyTransferedAll)
	if err := _KolKeys.contract.UnpackLog(event, "KeyTransferedAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysKolIpRightBoughtIterator is returned from FilterKolIpRightBought and is used to iterate over the raw logs and unpacked data for KolIpRightBought events raised by the KolKeys contract.
type KolKeysKolIpRightBoughtIterator struct {
	Event *KolKeysKolIpRightBought // Event containing the contract specifics and raw log

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
func (it *KolKeysKolIpRightBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysKolIpRightBought)
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
		it.Event = new(KolKeysKolIpRightBought)
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
func (it *KolKeysKolIpRightBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysKolIpRightBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysKolIpRightBought represents a KolIpRightBought event raised by the KolKeys contract.
type KolKeysKolIpRightBought struct {
	KolAddr    common.Address
	Buyer      common.Address
	KeyNo      *big.Int
	CurNonce   *big.Int
	KoltotalNo *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKolIpRightBought is a free log retrieval operation binding the contract event 0x804b19ef8fca4826d0401b0cc0bf4318c4ca656488c9ee92220dc3a6e2a251ac.
//
// Solidity: event KolIpRightBought(address kolAddr, address buyer, uint256 keyNo, uint256 curNonce, uint256 KoltotalNo)
func (_KolKeys *KolKeysFilterer) FilterKolIpRightBought(opts *bind.FilterOpts) (*KolKeysKolIpRightBoughtIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "KolIpRightBought")
	if err != nil {
		return nil, err
	}
	return &KolKeysKolIpRightBoughtIterator{contract: _KolKeys.contract, event: "KolIpRightBought", logs: logs, sub: sub}, nil
}

// WatchKolIpRightBought is a free log subscription operation binding the contract event 0x804b19ef8fca4826d0401b0cc0bf4318c4ca656488c9ee92220dc3a6e2a251ac.
//
// Solidity: event KolIpRightBought(address kolAddr, address buyer, uint256 keyNo, uint256 curNonce, uint256 KoltotalNo)
func (_KolKeys *KolKeysFilterer) WatchKolIpRightBought(opts *bind.WatchOpts, sink chan<- *KolKeysKolIpRightBought) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "KolIpRightBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysKolIpRightBought)
				if err := _KolKeys.contract.UnpackLog(event, "KolIpRightBought", log); err != nil {
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

// ParseKolIpRightBought is a log parse operation binding the contract event 0x804b19ef8fca4826d0401b0cc0bf4318c4ca656488c9ee92220dc3a6e2a251ac.
//
// Solidity: event KolIpRightBought(address kolAddr, address buyer, uint256 keyNo, uint256 curNonce, uint256 KoltotalNo)
func (_KolKeys *KolKeysFilterer) ParseKolIpRightBought(log types.Log) (*KolKeysKolIpRightBought, error) {
	event := new(KolKeysKolIpRightBought)
	if err := _KolKeys.contract.UnpackLog(event, "KolIpRightBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysKolKeyOperationIterator is returned from FilterKolKeyOperation and is used to iterate over the raw logs and unpacked data for KolKeyOperation events raised by the KolKeys contract.
type KolKeysKolKeyOperationIterator struct {
	Event *KolKeysKolKeyOperation // Event containing the contract specifics and raw log

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
func (it *KolKeysKolKeyOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysKolKeyOperation)
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
		it.Event = new(KolKeysKolKeyOperation)
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
func (it *KolKeysKolKeyOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysKolKeyOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysKolKeyOperation represents a KolKeyOperation event raised by the KolKeys contract.
type KolKeysKolKeyOperation struct {
	Kol      common.Address
	Price    *big.Int
	MaxKeyNo *big.Int
	Op       string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKolKeyOperation is a free log retrieval operation binding the contract event 0xbc15fd271700fa06e159fed0c580b92322ca63f6dee0e074e3b5b1cb69846df9.
//
// Solidity: event KolKeyOperation(address kol, uint256 price, uint256 maxKeyNo, string op)
func (_KolKeys *KolKeysFilterer) FilterKolKeyOperation(opts *bind.FilterOpts) (*KolKeysKolKeyOperationIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "KolKeyOperation")
	if err != nil {
		return nil, err
	}
	return &KolKeysKolKeyOperationIterator{contract: _KolKeys.contract, event: "KolKeyOperation", logs: logs, sub: sub}, nil
}

// WatchKolKeyOperation is a free log subscription operation binding the contract event 0xbc15fd271700fa06e159fed0c580b92322ca63f6dee0e074e3b5b1cb69846df9.
//
// Solidity: event KolKeyOperation(address kol, uint256 price, uint256 maxKeyNo, string op)
func (_KolKeys *KolKeysFilterer) WatchKolKeyOperation(opts *bind.WatchOpts, sink chan<- *KolKeysKolKeyOperation) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "KolKeyOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysKolKeyOperation)
				if err := _KolKeys.contract.UnpackLog(event, "KolKeyOperation", log); err != nil {
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

// ParseKolKeyOperation is a log parse operation binding the contract event 0xbc15fd271700fa06e159fed0c580b92322ca63f6dee0e074e3b5b1cb69846df9.
//
// Solidity: event KolKeyOperation(address kol, uint256 price, uint256 maxKeyNo, string op)
func (_KolKeys *KolKeysFilterer) ParseKolKeyOperation(log types.Log) (*KolKeysKolKeyOperation, error) {
	event := new(KolKeysKolKeyOperation)
	if err := _KolKeys.contract.UnpackLog(event, "KolKeyOperation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the KolKeys contract.
type KolKeysOwnerSetIterator struct {
	Event *KolKeysOwnerSet // Event containing the contract specifics and raw log

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
func (it *KolKeysOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysOwnerSet)
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
		it.Event = new(KolKeysOwnerSet)
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
func (it *KolKeysOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysOwnerSet represents a OwnerSet event raised by the KolKeys contract.
type KolKeysOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_KolKeys *KolKeysFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*KolKeysOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KolKeysOwnerSetIterator{contract: _KolKeys.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_KolKeys *KolKeysFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *KolKeysOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysOwnerSet)
				if err := _KolKeys.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseOwnerSet(log types.Log) (*KolKeysOwnerSet, error) {
	event := new(KolKeysOwnerSet)
	if err := _KolKeys.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysServiceFeeChangedIterator is returned from FilterServiceFeeChanged and is used to iterate over the raw logs and unpacked data for ServiceFeeChanged events raised by the KolKeys contract.
type KolKeysServiceFeeChangedIterator struct {
	Event *KolKeysServiceFeeChanged // Event containing the contract specifics and raw log

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
func (it *KolKeysServiceFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysServiceFeeChanged)
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
		it.Event = new(KolKeysServiceFeeChanged)
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
func (it *KolKeysServiceFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysServiceFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysServiceFeeChanged represents a ServiceFeeChanged event raised by the KolKeys contract.
type KolKeysServiceFeeChanged struct {
	NewSerficeFeeRate *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeChanged is a free log retrieval operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_KolKeys *KolKeysFilterer) FilterServiceFeeChanged(opts *bind.FilterOpts) (*KolKeysServiceFeeChangedIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return &KolKeysServiceFeeChangedIterator{contract: _KolKeys.contract, event: "ServiceFeeChanged", logs: logs, sub: sub}, nil
}

// WatchServiceFeeChanged is a free log subscription operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_KolKeys *KolKeysFilterer) WatchServiceFeeChanged(opts *bind.WatchOpts, sink chan<- *KolKeysServiceFeeChanged) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysServiceFeeChanged)
				if err := _KolKeys.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseServiceFeeChanged(log types.Log) (*KolKeysServiceFeeChanged, error) {
	event := new(KolKeysServiceFeeChanged)
	if err := _KolKeys.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysSystemSetIterator is returned from FilterSystemSet and is used to iterate over the raw logs and unpacked data for SystemSet events raised by the KolKeys contract.
type KolKeysSystemSetIterator struct {
	Event *KolKeysSystemSet // Event containing the contract specifics and raw log

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
func (it *KolKeysSystemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysSystemSet)
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
		it.Event = new(KolKeysSystemSet)
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
func (it *KolKeysSystemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysSystemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysSystemSet represents a SystemSet event raised by the KolKeys contract.
type KolKeysSystemSet struct {
	Num *big.Int
	Op  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSystemSet is a free log retrieval operation binding the contract event 0xcd714485b704afdf17516cfe0560106372ea651897ae76ae6879e628756fb2a0.
//
// Solidity: event SystemSet(uint256 num, string op)
func (_KolKeys *KolKeysFilterer) FilterSystemSet(opts *bind.FilterOpts) (*KolKeysSystemSetIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "SystemSet")
	if err != nil {
		return nil, err
	}
	return &KolKeysSystemSetIterator{contract: _KolKeys.contract, event: "SystemSet", logs: logs, sub: sub}, nil
}

// WatchSystemSet is a free log subscription operation binding the contract event 0xcd714485b704afdf17516cfe0560106372ea651897ae76ae6879e628756fb2a0.
//
// Solidity: event SystemSet(uint256 num, string op)
func (_KolKeys *KolKeysFilterer) WatchSystemSet(opts *bind.WatchOpts, sink chan<- *KolKeysSystemSet) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "SystemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysSystemSet)
				if err := _KolKeys.contract.UnpackLog(event, "SystemSet", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseSystemSet(log types.Log) (*KolKeysSystemSet, error) {
	event := new(KolKeysSystemSet)
	if err := _KolKeys.contract.UnpackLog(event, "SystemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysTweetBoughtIterator is returned from FilterTweetBought and is used to iterate over the raw logs and unpacked data for TweetBought events raised by the KolKeys contract.
type KolKeysTweetBoughtIterator struct {
	Event *KolKeysTweetBought // Event containing the contract specifics and raw log

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
func (it *KolKeysTweetBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysTweetBought)
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
		it.Event = new(KolKeysTweetBought)
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
func (it *KolKeysTweetBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysTweetBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysTweetBought represents a TweetBought event raised by the KolKeys contract.
type KolKeysTweetBought struct {
	THash   [32]byte
	Owner   common.Address
	Buyer   common.Address
	VoteNoe *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTweetBought is a free log retrieval operation binding the contract event 0xf2f9674ebd4d561babe7a6525f2ffd0990fb102b9d29ec89a42e5b08b5455ebf.
//
// Solidity: event TweetBought(bytes32 tHash, address owner, address buyer, uint256 voteNoe)
func (_KolKeys *KolKeysFilterer) FilterTweetBought(opts *bind.FilterOpts) (*KolKeysTweetBoughtIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "TweetBought")
	if err != nil {
		return nil, err
	}
	return &KolKeysTweetBoughtIterator{contract: _KolKeys.contract, event: "TweetBought", logs: logs, sub: sub}, nil
}

// WatchTweetBought is a free log subscription operation binding the contract event 0xf2f9674ebd4d561babe7a6525f2ffd0990fb102b9d29ec89a42e5b08b5455ebf.
//
// Solidity: event TweetBought(bytes32 tHash, address owner, address buyer, uint256 voteNoe)
func (_KolKeys *KolKeysFilterer) WatchTweetBought(opts *bind.WatchOpts, sink chan<- *KolKeysTweetBought) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "TweetBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysTweetBought)
				if err := _KolKeys.contract.UnpackLog(event, "TweetBought", log); err != nil {
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

// ParseTweetBought is a log parse operation binding the contract event 0xf2f9674ebd4d561babe7a6525f2ffd0990fb102b9d29ec89a42e5b08b5455ebf.
//
// Solidity: event TweetBought(bytes32 tHash, address owner, address buyer, uint256 voteNoe)
func (_KolKeys *KolKeysFilterer) ParseTweetBought(log types.Log) (*KolKeysTweetBought, error) {
	event := new(KolKeysTweetBought)
	if err := _KolKeys.contract.UnpackLog(event, "TweetBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysUpgradeToNewRuleIterator is returned from FilterUpgradeToNewRule and is used to iterate over the raw logs and unpacked data for UpgradeToNewRule events raised by the KolKeys contract.
type KolKeysUpgradeToNewRuleIterator struct {
	Event *KolKeysUpgradeToNewRule // Event containing the contract specifics and raw log

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
func (it *KolKeysUpgradeToNewRuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysUpgradeToNewRule)
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
		it.Event = new(KolKeysUpgradeToNewRule)
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
func (it *KolKeysUpgradeToNewRuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysUpgradeToNewRuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysUpgradeToNewRule represents a UpgradeToNewRule event raised by the KolKeys contract.
type KolKeysUpgradeToNewRule struct {
	NewContract common.Address
	Balance     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpgradeToNewRule is a free log retrieval operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_KolKeys *KolKeysFilterer) FilterUpgradeToNewRule(opts *bind.FilterOpts) (*KolKeysUpgradeToNewRuleIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return &KolKeysUpgradeToNewRuleIterator{contract: _KolKeys.contract, event: "UpgradeToNewRule", logs: logs, sub: sub}, nil
}

// WatchUpgradeToNewRule is a free log subscription operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_KolKeys *KolKeysFilterer) WatchUpgradeToNewRule(opts *bind.WatchOpts, sink chan<- *KolKeysUpgradeToNewRule) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysUpgradeToNewRule)
				if err := _KolKeys.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseUpgradeToNewRule(log types.Log) (*KolKeysUpgradeToNewRule, error) {
	event := new(KolKeysUpgradeToNewRule)
	if err := _KolKeys.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KolKeysWithdrawServiceIterator is returned from FilterWithdrawService and is used to iterate over the raw logs and unpacked data for WithdrawService events raised by the KolKeys contract.
type KolKeysWithdrawServiceIterator struct {
	Event *KolKeysWithdrawService // Event containing the contract specifics and raw log

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
func (it *KolKeysWithdrawServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KolKeysWithdrawService)
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
		it.Event = new(KolKeysWithdrawService)
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
func (it *KolKeysWithdrawServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KolKeysWithdrawServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KolKeysWithdrawService represents a WithdrawService event raised by the KolKeys contract.
type KolKeysWithdrawService struct {
	Owner   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawService is a free log retrieval operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_KolKeys *KolKeysFilterer) FilterWithdrawService(opts *bind.FilterOpts) (*KolKeysWithdrawServiceIterator, error) {

	logs, sub, err := _KolKeys.contract.FilterLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return &KolKeysWithdrawServiceIterator{contract: _KolKeys.contract, event: "WithdrawService", logs: logs, sub: sub}, nil
}

// WatchWithdrawService is a free log subscription operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_KolKeys *KolKeysFilterer) WatchWithdrawService(opts *bind.WatchOpts, sink chan<- *KolKeysWithdrawService) (event.Subscription, error) {

	logs, sub, err := _KolKeys.contract.WatchLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KolKeysWithdrawService)
				if err := _KolKeys.contract.UnpackLog(event, "WithdrawService", log); err != nil {
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
func (_KolKeys *KolKeysFilterer) ParseWithdrawService(log types.Log) (*KolKeysWithdrawService, error) {
	event := new(KolKeysWithdrawService)
	if err := _KolKeys.contract.UnpackLog(event, "WithdrawService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
