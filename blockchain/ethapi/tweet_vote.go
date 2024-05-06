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

// TweetVoteMetaData contains all meta data concerning the TweetVote contract.
var TweetVoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"opType\",\"type\":\"bool\"}],\"name\":\"AdminOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kolAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rightsNo\",\"type\":\"uint256\"}],\"name\":\"KolRightsBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KolWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typ\",\"type\":\"string\"}],\"name\":\"PluginChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSerficeFeeRate\",\"type\":\"uint256\"}],\"name\":\"ServiceFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rateName\",\"type\":\"string\"}],\"name\":\"SystemRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tweetHash\",\"type\":\"bytes32\"}],\"name\":\"TweetPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tweetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerVote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteNo\",\"type\":\"uint256\"}],\"name\":\"TweetVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"UpgradeToNewRule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"WithdrawService\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"__admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__minValCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminChangeKolKeyRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDelete\",\"type\":\"bool\"}],\"name\":\"adminOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminServiceFeeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGameAddr\",\"type\":\"address\"}],\"name\":\"adminSetGameContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetKolIncomePerTweetRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newKolContract\",\"type\":\"address\"}],\"name\":\"adminSetKolKeyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxVote\",\"type\":\"uint256\"}],\"name\":\"adminSetMaxVotePerTweet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetServiceFeeRateForPerTweetVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceInFinney\",\"type\":\"uint256\"}],\"name\":\"adminSetTweetPostPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceInFinney\",\"type\":\"uint256\"}],\"name\":\"adminSetTweetVotePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetWithdrawFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"adminStopKolKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"adminStopPlugin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"adminUpgradeToNewRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolIncomePerTweetVoteRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolKeyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolKeyIncomeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolKeyIncomeSourceID\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kolKeyStop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVotePerTweet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"tweetHashs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"migrateTweetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneFinney\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ownersOfAllTweets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"publishTweet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"prefixedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeePerTweetVoteRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tweetPostPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tweetVotePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tweetHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteNo\",\"type\":\"uint256\"}],\"name\":\"voteToTweets\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"all\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TweetVoteABI is the input ABI used to generate the binding from.
// Deprecated: Use TweetVoteMetaData.ABI instead.
var TweetVoteABI = TweetVoteMetaData.ABI

// TweetVote is an auto generated Go binding around an Ethereum contract.
type TweetVote struct {
	TweetVoteCaller     // Read-only binding to the contract
	TweetVoteTransactor // Write-only binding to the contract
	TweetVoteFilterer   // Log filterer for contract events
}

// TweetVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type TweetVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TweetVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TweetVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TweetVoteSession struct {
	Contract     *TweetVote        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TweetVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TweetVoteCallerSession struct {
	Contract *TweetVoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TweetVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TweetVoteTransactorSession struct {
	Contract     *TweetVoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TweetVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type TweetVoteRaw struct {
	Contract *TweetVote // Generic contract binding to access the raw methods on
}

// TweetVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TweetVoteCallerRaw struct {
	Contract *TweetVoteCaller // Generic read-only contract binding to access the raw methods on
}

// TweetVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TweetVoteTransactorRaw struct {
	Contract *TweetVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTweetVote creates a new instance of TweetVote, bound to a specific deployed contract.
func NewTweetVote(address common.Address, backend bind.ContractBackend) (*TweetVote, error) {
	contract, err := bindTweetVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TweetVote{TweetVoteCaller: TweetVoteCaller{contract: contract}, TweetVoteTransactor: TweetVoteTransactor{contract: contract}, TweetVoteFilterer: TweetVoteFilterer{contract: contract}}, nil
}

// NewTweetVoteCaller creates a new read-only instance of TweetVote, bound to a specific deployed contract.
func NewTweetVoteCaller(address common.Address, caller bind.ContractCaller) (*TweetVoteCaller, error) {
	contract, err := bindTweetVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TweetVoteCaller{contract: contract}, nil
}

// NewTweetVoteTransactor creates a new write-only instance of TweetVote, bound to a specific deployed contract.
func NewTweetVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*TweetVoteTransactor, error) {
	contract, err := bindTweetVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TweetVoteTransactor{contract: contract}, nil
}

// NewTweetVoteFilterer creates a new log filterer instance of TweetVote, bound to a specific deployed contract.
func NewTweetVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*TweetVoteFilterer, error) {
	contract, err := bindTweetVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TweetVoteFilterer{contract: contract}, nil
}

// bindTweetVote binds a generic wrapper to an already deployed contract.
func bindTweetVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TweetVoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TweetVote *TweetVoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TweetVote.Contract.TweetVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TweetVote *TweetVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetVote.Contract.TweetVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TweetVote *TweetVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TweetVote.Contract.TweetVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TweetVote *TweetVoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TweetVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TweetVote *TweetVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TweetVote *TweetVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TweetVote.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_TweetVote *TweetVoteCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "__admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_TweetVote *TweetVoteSession) Admins(arg0 common.Address) (bool, error) {
	return _TweetVote.Contract.Admins(&_TweetVote.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_TweetVote *TweetVoteCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _TweetVote.Contract.Admins(&_TweetVote.CallOpts, arg0)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_TweetVote *TweetVoteCaller) MinValCheck(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "__minValCheck")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_TweetVote *TweetVoteSession) MinValCheck() (*big.Int, error) {
	return _TweetVote.Contract.MinValCheck(&_TweetVote.CallOpts)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) MinValCheck() (*big.Int, error) {
	return _TweetVote.Contract.MinValCheck(&_TweetVote.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_TweetVote *TweetVoteCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_TweetVote *TweetVoteSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _TweetVote.Contract.Balance(&_TweetVote.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _TweetVote.Contract.Balance(&_TweetVote.CallOpts, arg0)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_TweetVote *TweetVoteCaller) ContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "contractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_TweetVote *TweetVoteSession) ContractBalance() (*big.Int, error) {
	return _TweetVote.Contract.ContractBalance(&_TweetVote.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) ContractBalance() (*big.Int, error) {
	return _TweetVote.Contract.ContractBalance(&_TweetVote.CallOpts)
}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_TweetVote *TweetVoteCaller) GameContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "gameContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_TweetVote *TweetVoteSession) GameContract() (common.Address, error) {
	return _TweetVote.Contract.GameContract(&_TweetVote.CallOpts)
}

// GameContract is a free data retrieval call binding the contract method 0xd3f33009.
//
// Solidity: function gameContract() view returns(address)
func (_TweetVote *TweetVoteCallerSession) GameContract() (common.Address, error) {
	return _TweetVote.Contract.GameContract(&_TweetVote.CallOpts)
}

// GameStop is a free data retrieval call binding the contract method 0xd70feef1.
//
// Solidity: function gameStop() view returns(bool)
func (_TweetVote *TweetVoteCaller) GameStop(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "gameStop")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameStop is a free data retrieval call binding the contract method 0xd70feef1.
//
// Solidity: function gameStop() view returns(bool)
func (_TweetVote *TweetVoteSession) GameStop() (bool, error) {
	return _TweetVote.Contract.GameStop(&_TweetVote.CallOpts)
}

// GameStop is a free data retrieval call binding the contract method 0xd70feef1.
//
// Solidity: function gameStop() view returns(bool)
func (_TweetVote *TweetVoteCallerSession) GameStop() (bool, error) {
	return _TweetVote.Contract.GameStop(&_TweetVote.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TweetVote *TweetVoteCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TweetVote *TweetVoteSession) GetOwner() (common.Address, error) {
	return _TweetVote.Contract.GetOwner(&_TweetVote.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TweetVote *TweetVoteCallerSession) GetOwner() (common.Address, error) {
	return _TweetVote.Contract.GetOwner(&_TweetVote.CallOpts)
}

// KolIncomePerTweetVoteRate is a free data retrieval call binding the contract method 0x7b3e47d1.
//
// Solidity: function kolIncomePerTweetVoteRate() view returns(uint8)
func (_TweetVote *TweetVoteCaller) KolIncomePerTweetVoteRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "kolIncomePerTweetVoteRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KolIncomePerTweetVoteRate is a free data retrieval call binding the contract method 0x7b3e47d1.
//
// Solidity: function kolIncomePerTweetVoteRate() view returns(uint8)
func (_TweetVote *TweetVoteSession) KolIncomePerTweetVoteRate() (uint8, error) {
	return _TweetVote.Contract.KolIncomePerTweetVoteRate(&_TweetVote.CallOpts)
}

// KolIncomePerTweetVoteRate is a free data retrieval call binding the contract method 0x7b3e47d1.
//
// Solidity: function kolIncomePerTweetVoteRate() view returns(uint8)
func (_TweetVote *TweetVoteCallerSession) KolIncomePerTweetVoteRate() (uint8, error) {
	return _TweetVote.Contract.KolIncomePerTweetVoteRate(&_TweetVote.CallOpts)
}

// KolKeyContract is a free data retrieval call binding the contract method 0x28293872.
//
// Solidity: function kolKeyContract() view returns(address)
func (_TweetVote *TweetVoteCaller) KolKeyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "kolKeyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KolKeyContract is a free data retrieval call binding the contract method 0x28293872.
//
// Solidity: function kolKeyContract() view returns(address)
func (_TweetVote *TweetVoteSession) KolKeyContract() (common.Address, error) {
	return _TweetVote.Contract.KolKeyContract(&_TweetVote.CallOpts)
}

// KolKeyContract is a free data retrieval call binding the contract method 0x28293872.
//
// Solidity: function kolKeyContract() view returns(address)
func (_TweetVote *TweetVoteCallerSession) KolKeyContract() (common.Address, error) {
	return _TweetVote.Contract.KolKeyContract(&_TweetVote.CallOpts)
}

// KolKeyIncomeRate is a free data retrieval call binding the contract method 0x5feab7cb.
//
// Solidity: function kolKeyIncomeRate() view returns(uint256)
func (_TweetVote *TweetVoteCaller) KolKeyIncomeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "kolKeyIncomeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KolKeyIncomeRate is a free data retrieval call binding the contract method 0x5feab7cb.
//
// Solidity: function kolKeyIncomeRate() view returns(uint256)
func (_TweetVote *TweetVoteSession) KolKeyIncomeRate() (*big.Int, error) {
	return _TweetVote.Contract.KolKeyIncomeRate(&_TweetVote.CallOpts)
}

// KolKeyIncomeRate is a free data retrieval call binding the contract method 0x5feab7cb.
//
// Solidity: function kolKeyIncomeRate() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) KolKeyIncomeRate() (*big.Int, error) {
	return _TweetVote.Contract.KolKeyIncomeRate(&_TweetVote.CallOpts)
}

// KolKeyIncomeSourceID is a free data retrieval call binding the contract method 0x98ac371b.
//
// Solidity: function kolKeyIncomeSourceID() view returns(int8)
func (_TweetVote *TweetVoteCaller) KolKeyIncomeSourceID(opts *bind.CallOpts) (int8, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "kolKeyIncomeSourceID")

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// KolKeyIncomeSourceID is a free data retrieval call binding the contract method 0x98ac371b.
//
// Solidity: function kolKeyIncomeSourceID() view returns(int8)
func (_TweetVote *TweetVoteSession) KolKeyIncomeSourceID() (int8, error) {
	return _TweetVote.Contract.KolKeyIncomeSourceID(&_TweetVote.CallOpts)
}

// KolKeyIncomeSourceID is a free data retrieval call binding the contract method 0x98ac371b.
//
// Solidity: function kolKeyIncomeSourceID() view returns(int8)
func (_TweetVote *TweetVoteCallerSession) KolKeyIncomeSourceID() (int8, error) {
	return _TweetVote.Contract.KolKeyIncomeSourceID(&_TweetVote.CallOpts)
}

// KolKeyStop is a free data retrieval call binding the contract method 0xf929837a.
//
// Solidity: function kolKeyStop() view returns(bool)
func (_TweetVote *TweetVoteCaller) KolKeyStop(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "kolKeyStop")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KolKeyStop is a free data retrieval call binding the contract method 0xf929837a.
//
// Solidity: function kolKeyStop() view returns(bool)
func (_TweetVote *TweetVoteSession) KolKeyStop() (bool, error) {
	return _TweetVote.Contract.KolKeyStop(&_TweetVote.CallOpts)
}

// KolKeyStop is a free data retrieval call binding the contract method 0xf929837a.
//
// Solidity: function kolKeyStop() view returns(bool)
func (_TweetVote *TweetVoteCallerSession) KolKeyStop() (bool, error) {
	return _TweetVote.Contract.KolKeyStop(&_TweetVote.CallOpts)
}

// MaxVotePerTweet is a free data retrieval call binding the contract method 0xdc8c071c.
//
// Solidity: function maxVotePerTweet() view returns(uint256)
func (_TweetVote *TweetVoteCaller) MaxVotePerTweet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "maxVotePerTweet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVotePerTweet is a free data retrieval call binding the contract method 0xdc8c071c.
//
// Solidity: function maxVotePerTweet() view returns(uint256)
func (_TweetVote *TweetVoteSession) MaxVotePerTweet() (*big.Int, error) {
	return _TweetVote.Contract.MaxVotePerTweet(&_TweetVote.CallOpts)
}

// MaxVotePerTweet is a free data retrieval call binding the contract method 0xdc8c071c.
//
// Solidity: function maxVotePerTweet() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) MaxVotePerTweet() (*big.Int, error) {
	return _TweetVote.Contract.MaxVotePerTweet(&_TweetVote.CallOpts)
}

// OneFinney is a free data retrieval call binding the contract method 0x080da5dc.
//
// Solidity: function oneFinney() view returns(uint256)
func (_TweetVote *TweetVoteCaller) OneFinney(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "oneFinney")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OneFinney is a free data retrieval call binding the contract method 0x080da5dc.
//
// Solidity: function oneFinney() view returns(uint256)
func (_TweetVote *TweetVoteSession) OneFinney() (*big.Int, error) {
	return _TweetVote.Contract.OneFinney(&_TweetVote.CallOpts)
}

// OneFinney is a free data retrieval call binding the contract method 0x080da5dc.
//
// Solidity: function oneFinney() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) OneFinney() (*big.Int, error) {
	return _TweetVote.Contract.OneFinney(&_TweetVote.CallOpts)
}

// OwnersOfAllTweets is a free data retrieval call binding the contract method 0x60b7a81b.
//
// Solidity: function ownersOfAllTweets(bytes32 ) view returns(address)
func (_TweetVote *TweetVoteCaller) OwnersOfAllTweets(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "ownersOfAllTweets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnersOfAllTweets is a free data retrieval call binding the contract method 0x60b7a81b.
//
// Solidity: function ownersOfAllTweets(bytes32 ) view returns(address)
func (_TweetVote *TweetVoteSession) OwnersOfAllTweets(arg0 [32]byte) (common.Address, error) {
	return _TweetVote.Contract.OwnersOfAllTweets(&_TweetVote.CallOpts, arg0)
}

// OwnersOfAllTweets is a free data retrieval call binding the contract method 0x60b7a81b.
//
// Solidity: function ownersOfAllTweets(bytes32 ) view returns(address)
func (_TweetVote *TweetVoteCallerSession) OwnersOfAllTweets(arg0 [32]byte) (common.Address, error) {
	return _TweetVote.Contract.OwnersOfAllTweets(&_TweetVote.CallOpts, arg0)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 prefixedHash, bytes signature) pure returns(address)
func (_TweetVote *TweetVoteCaller) RecoverSigner(opts *bind.CallOpts, prefixedHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "recoverSigner", prefixedHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 prefixedHash, bytes signature) pure returns(address)
func (_TweetVote *TweetVoteSession) RecoverSigner(prefixedHash [32]byte, signature []byte) (common.Address, error) {
	return _TweetVote.Contract.RecoverSigner(&_TweetVote.CallOpts, prefixedHash, signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 prefixedHash, bytes signature) pure returns(address)
func (_TweetVote *TweetVoteCallerSession) RecoverSigner(prefixedHash [32]byte, signature []byte) (common.Address, error) {
	return _TweetVote.Contract.RecoverSigner(&_TweetVote.CallOpts, prefixedHash, signature)
}

// ServiceFeePerTweetVoteRate is a free data retrieval call binding the contract method 0x1028672f.
//
// Solidity: function serviceFeePerTweetVoteRate() view returns(uint8)
func (_TweetVote *TweetVoteCaller) ServiceFeePerTweetVoteRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "serviceFeePerTweetVoteRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ServiceFeePerTweetVoteRate is a free data retrieval call binding the contract method 0x1028672f.
//
// Solidity: function serviceFeePerTweetVoteRate() view returns(uint8)
func (_TweetVote *TweetVoteSession) ServiceFeePerTweetVoteRate() (uint8, error) {
	return _TweetVote.Contract.ServiceFeePerTweetVoteRate(&_TweetVote.CallOpts)
}

// ServiceFeePerTweetVoteRate is a free data retrieval call binding the contract method 0x1028672f.
//
// Solidity: function serviceFeePerTweetVoteRate() view returns(uint8)
func (_TweetVote *TweetVoteCallerSession) ServiceFeePerTweetVoteRate() (uint8, error) {
	return _TweetVote.Contract.ServiceFeePerTweetVoteRate(&_TweetVote.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_TweetVote *TweetVoteCaller) ServiceFeeReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "serviceFeeReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_TweetVote *TweetVoteSession) ServiceFeeReceived() (*big.Int, error) {
	return _TweetVote.Contract.ServiceFeeReceived(&_TweetVote.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) ServiceFeeReceived() (*big.Int, error) {
	return _TweetVote.Contract.ServiceFeeReceived(&_TweetVote.CallOpts)
}

// SystemSettings is a free data retrieval call binding the contract method 0x60b42f12.
//
// Solidity: function systemSettings() view returns(uint256, uint256, uint256, address, bool, uint8, uint8)
func (_TweetVote *TweetVoteCaller) SystemSettings(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, common.Address, bool, uint8, uint8, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "systemSettings")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(bool), *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)
	out5 := *abi.ConvertType(out[5], new(uint8)).(*uint8)
	out6 := *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// SystemSettings is a free data retrieval call binding the contract method 0x60b42f12.
//
// Solidity: function systemSettings() view returns(uint256, uint256, uint256, address, bool, uint8, uint8)
func (_TweetVote *TweetVoteSession) SystemSettings() (*big.Int, *big.Int, *big.Int, common.Address, bool, uint8, uint8, error) {
	return _TweetVote.Contract.SystemSettings(&_TweetVote.CallOpts)
}

// SystemSettings is a free data retrieval call binding the contract method 0x60b42f12.
//
// Solidity: function systemSettings() view returns(uint256, uint256, uint256, address, bool, uint8, uint8)
func (_TweetVote *TweetVoteCallerSession) SystemSettings() (*big.Int, *big.Int, *big.Int, common.Address, bool, uint8, uint8, error) {
	return _TweetVote.Contract.SystemSettings(&_TweetVote.CallOpts)
}

// TweetPostPrice is a free data retrieval call binding the contract method 0x04cea976.
//
// Solidity: function tweetPostPrice() view returns(uint256)
func (_TweetVote *TweetVoteCaller) TweetPostPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "tweetPostPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TweetPostPrice is a free data retrieval call binding the contract method 0x04cea976.
//
// Solidity: function tweetPostPrice() view returns(uint256)
func (_TweetVote *TweetVoteSession) TweetPostPrice() (*big.Int, error) {
	return _TweetVote.Contract.TweetPostPrice(&_TweetVote.CallOpts)
}

// TweetPostPrice is a free data retrieval call binding the contract method 0x04cea976.
//
// Solidity: function tweetPostPrice() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) TweetPostPrice() (*big.Int, error) {
	return _TweetVote.Contract.TweetPostPrice(&_TweetVote.CallOpts)
}

// TweetVotePrice is a free data retrieval call binding the contract method 0x33806cfe.
//
// Solidity: function tweetVotePrice() view returns(uint256)
func (_TweetVote *TweetVoteCaller) TweetVotePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "tweetVotePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TweetVotePrice is a free data retrieval call binding the contract method 0x33806cfe.
//
// Solidity: function tweetVotePrice() view returns(uint256)
func (_TweetVote *TweetVoteSession) TweetVotePrice() (*big.Int, error) {
	return _TweetVote.Contract.TweetVotePrice(&_TweetVote.CallOpts)
}

// TweetVotePrice is a free data retrieval call binding the contract method 0x33806cfe.
//
// Solidity: function tweetVotePrice() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) TweetVotePrice() (*big.Int, error) {
	return _TweetVote.Contract.TweetVotePrice(&_TweetVote.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_TweetVote *TweetVoteCaller) WithdrawFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetVote.contract.Call(opts, &out, "withdrawFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_TweetVote *TweetVoteSession) WithdrawFeeRate() (*big.Int, error) {
	return _TweetVote.Contract.WithdrawFeeRate(&_TweetVote.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_TweetVote *TweetVoteCallerSession) WithdrawFeeRate() (*big.Int, error) {
	return _TweetVote.Contract.WithdrawFeeRate(&_TweetVote.CallOpts)
}

// AdminChangeKolKeyRate is a paid mutator transaction binding the contract method 0x82580978.
//
// Solidity: function adminChangeKolKeyRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactor) AdminChangeKolKeyRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminChangeKolKeyRate", newRate)
}

// AdminChangeKolKeyRate is a paid mutator transaction binding the contract method 0x82580978.
//
// Solidity: function adminChangeKolKeyRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteSession) AdminChangeKolKeyRate(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminChangeKolKeyRate(&_TweetVote.TransactOpts, newRate)
}

// AdminChangeKolKeyRate is a paid mutator transaction binding the contract method 0x82580978.
//
// Solidity: function adminChangeKolKeyRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminChangeKolKeyRate(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminChangeKolKeyRate(&_TweetVote.TransactOpts, newRate)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_TweetVote *TweetVoteTransactor) AdminOperation(opts *bind.TransactOpts, admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminOperation", admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_TweetVote *TweetVoteSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminOperation(&_TweetVote.TransactOpts, admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminOperation(&_TweetVote.TransactOpts, admin, isDelete)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_TweetVote *TweetVoteTransactor) AdminServiceFeeWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminServiceFeeWithdraw")
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_TweetVote *TweetVoteSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _TweetVote.Contract.AdminServiceFeeWithdraw(&_TweetVote.TransactOpts)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_TweetVote *TweetVoteTransactorSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _TweetVote.Contract.AdminServiceFeeWithdraw(&_TweetVote.TransactOpts)
}

// AdminSetGameContract is a paid mutator transaction binding the contract method 0xeb54be2a.
//
// Solidity: function adminSetGameContract(address newGameAddr) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetGameContract(opts *bind.TransactOpts, newGameAddr common.Address) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetGameContract", newGameAddr)
}

// AdminSetGameContract is a paid mutator transaction binding the contract method 0xeb54be2a.
//
// Solidity: function adminSetGameContract(address newGameAddr) returns()
func (_TweetVote *TweetVoteSession) AdminSetGameContract(newGameAddr common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetGameContract(&_TweetVote.TransactOpts, newGameAddr)
}

// AdminSetGameContract is a paid mutator transaction binding the contract method 0xeb54be2a.
//
// Solidity: function adminSetGameContract(address newGameAddr) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetGameContract(newGameAddr common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetGameContract(&_TweetVote.TransactOpts, newGameAddr)
}

// AdminSetKolIncomePerTweetRate is a paid mutator transaction binding the contract method 0xa1cd090a.
//
// Solidity: function adminSetKolIncomePerTweetRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetKolIncomePerTweetRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetKolIncomePerTweetRate", newRate)
}

// AdminSetKolIncomePerTweetRate is a paid mutator transaction binding the contract method 0xa1cd090a.
//
// Solidity: function adminSetKolIncomePerTweetRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteSession) AdminSetKolIncomePerTweetRate(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetKolIncomePerTweetRate(&_TweetVote.TransactOpts, newRate)
}

// AdminSetKolIncomePerTweetRate is a paid mutator transaction binding the contract method 0xa1cd090a.
//
// Solidity: function adminSetKolIncomePerTweetRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetKolIncomePerTweetRate(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetKolIncomePerTweetRate(&_TweetVote.TransactOpts, newRate)
}

// AdminSetKolKeyContract is a paid mutator transaction binding the contract method 0x23266c86.
//
// Solidity: function adminSetKolKeyContract(address newKolContract) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetKolKeyContract(opts *bind.TransactOpts, newKolContract common.Address) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetKolKeyContract", newKolContract)
}

// AdminSetKolKeyContract is a paid mutator transaction binding the contract method 0x23266c86.
//
// Solidity: function adminSetKolKeyContract(address newKolContract) returns()
func (_TweetVote *TweetVoteSession) AdminSetKolKeyContract(newKolContract common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetKolKeyContract(&_TweetVote.TransactOpts, newKolContract)
}

// AdminSetKolKeyContract is a paid mutator transaction binding the contract method 0x23266c86.
//
// Solidity: function adminSetKolKeyContract(address newKolContract) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetKolKeyContract(newKolContract common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetKolKeyContract(&_TweetVote.TransactOpts, newKolContract)
}

// AdminSetMaxVotePerTweet is a paid mutator transaction binding the contract method 0x489e2590.
//
// Solidity: function adminSetMaxVotePerTweet(uint256 newMaxVote) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetMaxVotePerTweet(opts *bind.TransactOpts, newMaxVote *big.Int) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetMaxVotePerTweet", newMaxVote)
}

// AdminSetMaxVotePerTweet is a paid mutator transaction binding the contract method 0x489e2590.
//
// Solidity: function adminSetMaxVotePerTweet(uint256 newMaxVote) returns()
func (_TweetVote *TweetVoteSession) AdminSetMaxVotePerTweet(newMaxVote *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetMaxVotePerTweet(&_TweetVote.TransactOpts, newMaxVote)
}

// AdminSetMaxVotePerTweet is a paid mutator transaction binding the contract method 0x489e2590.
//
// Solidity: function adminSetMaxVotePerTweet(uint256 newMaxVote) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetMaxVotePerTweet(newMaxVote *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetMaxVotePerTweet(&_TweetVote.TransactOpts, newMaxVote)
}

// AdminSetServiceFeeRateForPerTweetVote is a paid mutator transaction binding the contract method 0xcef8e395.
//
// Solidity: function adminSetServiceFeeRateForPerTweetVote(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetServiceFeeRateForPerTweetVote(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetServiceFeeRateForPerTweetVote", newRate)
}

// AdminSetServiceFeeRateForPerTweetVote is a paid mutator transaction binding the contract method 0xcef8e395.
//
// Solidity: function adminSetServiceFeeRateForPerTweetVote(uint8 newRate) returns()
func (_TweetVote *TweetVoteSession) AdminSetServiceFeeRateForPerTweetVote(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetServiceFeeRateForPerTweetVote(&_TweetVote.TransactOpts, newRate)
}

// AdminSetServiceFeeRateForPerTweetVote is a paid mutator transaction binding the contract method 0xcef8e395.
//
// Solidity: function adminSetServiceFeeRateForPerTweetVote(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetServiceFeeRateForPerTweetVote(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetServiceFeeRateForPerTweetVote(&_TweetVote.TransactOpts, newRate)
}

// AdminSetTweetPostPrice is a paid mutator transaction binding the contract method 0xf7d58241.
//
// Solidity: function adminSetTweetPostPrice(uint256 newPriceInFinney) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetTweetPostPrice(opts *bind.TransactOpts, newPriceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetTweetPostPrice", newPriceInFinney)
}

// AdminSetTweetPostPrice is a paid mutator transaction binding the contract method 0xf7d58241.
//
// Solidity: function adminSetTweetPostPrice(uint256 newPriceInFinney) returns()
func (_TweetVote *TweetVoteSession) AdminSetTweetPostPrice(newPriceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetTweetPostPrice(&_TweetVote.TransactOpts, newPriceInFinney)
}

// AdminSetTweetPostPrice is a paid mutator transaction binding the contract method 0xf7d58241.
//
// Solidity: function adminSetTweetPostPrice(uint256 newPriceInFinney) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetTweetPostPrice(newPriceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetTweetPostPrice(&_TweetVote.TransactOpts, newPriceInFinney)
}

// AdminSetTweetVotePrice is a paid mutator transaction binding the contract method 0x1a25d1c0.
//
// Solidity: function adminSetTweetVotePrice(uint256 newPriceInFinney) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetTweetVotePrice(opts *bind.TransactOpts, newPriceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetTweetVotePrice", newPriceInFinney)
}

// AdminSetTweetVotePrice is a paid mutator transaction binding the contract method 0x1a25d1c0.
//
// Solidity: function adminSetTweetVotePrice(uint256 newPriceInFinney) returns()
func (_TweetVote *TweetVoteSession) AdminSetTweetVotePrice(newPriceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetTweetVotePrice(&_TweetVote.TransactOpts, newPriceInFinney)
}

// AdminSetTweetVotePrice is a paid mutator transaction binding the contract method 0x1a25d1c0.
//
// Solidity: function adminSetTweetVotePrice(uint256 newPriceInFinney) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetTweetVotePrice(newPriceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetTweetVotePrice(&_TweetVote.TransactOpts, newPriceInFinney)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactor) AdminSetWithdrawFeeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminSetWithdrawFeeRate", newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetWithdrawFeeRate(&_TweetVote.TransactOpts, newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminSetWithdrawFeeRate(&_TweetVote.TransactOpts, newRate)
}

// AdminStopKolKey is a paid mutator transaction binding the contract method 0x6b6f28b7.
//
// Solidity: function adminStopKolKey(bool stop) returns()
func (_TweetVote *TweetVoteTransactor) AdminStopKolKey(opts *bind.TransactOpts, stop bool) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminStopKolKey", stop)
}

// AdminStopKolKey is a paid mutator transaction binding the contract method 0x6b6f28b7.
//
// Solidity: function adminStopKolKey(bool stop) returns()
func (_TweetVote *TweetVoteSession) AdminStopKolKey(stop bool) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminStopKolKey(&_TweetVote.TransactOpts, stop)
}

// AdminStopKolKey is a paid mutator transaction binding the contract method 0x6b6f28b7.
//
// Solidity: function adminStopKolKey(bool stop) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminStopKolKey(stop bool) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminStopKolKey(&_TweetVote.TransactOpts, stop)
}

// AdminStopPlugin is a paid mutator transaction binding the contract method 0x24a5834e.
//
// Solidity: function adminStopPlugin(bool stop) returns()
func (_TweetVote *TweetVoteTransactor) AdminStopPlugin(opts *bind.TransactOpts, stop bool) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminStopPlugin", stop)
}

// AdminStopPlugin is a paid mutator transaction binding the contract method 0x24a5834e.
//
// Solidity: function adminStopPlugin(bool stop) returns()
func (_TweetVote *TweetVoteSession) AdminStopPlugin(stop bool) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminStopPlugin(&_TweetVote.TransactOpts, stop)
}

// AdminStopPlugin is a paid mutator transaction binding the contract method 0x24a5834e.
//
// Solidity: function adminStopPlugin(bool stop) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminStopPlugin(stop bool) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminStopPlugin(&_TweetVote.TransactOpts, stop)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_TweetVote *TweetVoteTransactor) AdminUpgradeToNewRule(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "adminUpgradeToNewRule", recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_TweetVote *TweetVoteSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminUpgradeToNewRule(&_TweetVote.TransactOpts, recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_TweetVote *TweetVoteTransactorSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.AdminUpgradeToNewRule(&_TweetVote.TransactOpts, recipient)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_TweetVote *TweetVoteTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_TweetVote *TweetVoteSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.ChangeOwner(&_TweetVote.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_TweetVote *TweetVoteTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.ChangeOwner(&_TweetVote.TransactOpts, newOwner)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_TweetVote *TweetVoteTransactor) ChangeStatus(opts *bind.TransactOpts, stop bool) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "changeStatus", stop)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_TweetVote *TweetVoteSession) ChangeStatus(stop bool) (*types.Transaction, error) {
	return _TweetVote.Contract.ChangeStatus(&_TweetVote.TransactOpts, stop)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_TweetVote *TweetVoteTransactorSession) ChangeStatus(stop bool) (*types.Transaction, error) {
	return _TweetVote.Contract.ChangeStatus(&_TweetVote.TransactOpts, stop)
}

// MigrateTweetOwner is a paid mutator transaction binding the contract method 0xf8cc9bac.
//
// Solidity: function migrateTweetOwner(bytes32[] tweetHashs, address newOwner) returns()
func (_TweetVote *TweetVoteTransactor) MigrateTweetOwner(opts *bind.TransactOpts, tweetHashs [][32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "migrateTweetOwner", tweetHashs, newOwner)
}

// MigrateTweetOwner is a paid mutator transaction binding the contract method 0xf8cc9bac.
//
// Solidity: function migrateTweetOwner(bytes32[] tweetHashs, address newOwner) returns()
func (_TweetVote *TweetVoteSession) MigrateTweetOwner(tweetHashs [][32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.MigrateTweetOwner(&_TweetVote.TransactOpts, tweetHashs, newOwner)
}

// MigrateTweetOwner is a paid mutator transaction binding the contract method 0xf8cc9bac.
//
// Solidity: function migrateTweetOwner(bytes32[] tweetHashs, address newOwner) returns()
func (_TweetVote *TweetVoteTransactorSession) MigrateTweetOwner(tweetHashs [][32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _TweetVote.Contract.MigrateTweetOwner(&_TweetVote.TransactOpts, tweetHashs, newOwner)
}

// PublishTweet is a paid mutator transaction binding the contract method 0x4bf8d46f.
//
// Solidity: function publishTweet(bytes32 hash, bytes signature) payable returns()
func (_TweetVote *TweetVoteTransactor) PublishTweet(opts *bind.TransactOpts, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "publishTweet", hash, signature)
}

// PublishTweet is a paid mutator transaction binding the contract method 0x4bf8d46f.
//
// Solidity: function publishTweet(bytes32 hash, bytes signature) payable returns()
func (_TweetVote *TweetVoteSession) PublishTweet(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TweetVote.Contract.PublishTweet(&_TweetVote.TransactOpts, hash, signature)
}

// PublishTweet is a paid mutator transaction binding the contract method 0x4bf8d46f.
//
// Solidity: function publishTweet(bytes32 hash, bytes signature) payable returns()
func (_TweetVote *TweetVoteTransactorSession) PublishTweet(hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TweetVote.Contract.PublishTweet(&_TweetVote.TransactOpts, hash, signature)
}

// VoteToTweets is a paid mutator transaction binding the contract method 0x4e7da0e8.
//
// Solidity: function voteToTweets(bytes32 tweetHash, uint256 voteNo) payable returns()
func (_TweetVote *TweetVoteTransactor) VoteToTweets(opts *bind.TransactOpts, tweetHash [32]byte, voteNo *big.Int) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "voteToTweets", tweetHash, voteNo)
}

// VoteToTweets is a paid mutator transaction binding the contract method 0x4e7da0e8.
//
// Solidity: function voteToTweets(bytes32 tweetHash, uint256 voteNo) payable returns()
func (_TweetVote *TweetVoteSession) VoteToTweets(tweetHash [32]byte, voteNo *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.VoteToTweets(&_TweetVote.TransactOpts, tweetHash, voteNo)
}

// VoteToTweets is a paid mutator transaction binding the contract method 0x4e7da0e8.
//
// Solidity: function voteToTweets(bytes32 tweetHash, uint256 voteNo) payable returns()
func (_TweetVote *TweetVoteTransactorSession) VoteToTweets(tweetHash [32]byte, voteNo *big.Int) (*types.Transaction, error) {
	return _TweetVote.Contract.VoteToTweets(&_TweetVote.TransactOpts, tweetHash, voteNo)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_TweetVote *TweetVoteTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, all bool) (*types.Transaction, error) {
	return _TweetVote.contract.Transact(opts, "withdraw", amount, all)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_TweetVote *TweetVoteSession) Withdraw(amount *big.Int, all bool) (*types.Transaction, error) {
	return _TweetVote.Contract.Withdraw(&_TweetVote.TransactOpts, amount, all)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_TweetVote *TweetVoteTransactorSession) Withdraw(amount *big.Int, all bool) (*types.Transaction, error) {
	return _TweetVote.Contract.Withdraw(&_TweetVote.TransactOpts, amount, all)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TweetVote *TweetVoteTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetVote.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TweetVote *TweetVoteSession) Receive() (*types.Transaction, error) {
	return _TweetVote.Contract.Receive(&_TweetVote.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TweetVote *TweetVoteTransactorSession) Receive() (*types.Transaction, error) {
	return _TweetVote.Contract.Receive(&_TweetVote.TransactOpts)
}

// TweetVoteAdminOperationIterator is returned from FilterAdminOperation and is used to iterate over the raw logs and unpacked data for AdminOperation events raised by the TweetVote contract.
type TweetVoteAdminOperationIterator struct {
	Event *TweetVoteAdminOperation // Event containing the contract specifics and raw log

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
func (it *TweetVoteAdminOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteAdminOperation)
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
		it.Event = new(TweetVoteAdminOperation)
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
func (it *TweetVoteAdminOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteAdminOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteAdminOperation represents a AdminOperation event raised by the TweetVote contract.
type TweetVoteAdminOperation struct {
	Admin  common.Address
	OpType bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAdminOperation is a free log retrieval operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_TweetVote *TweetVoteFilterer) FilterAdminOperation(opts *bind.FilterOpts) (*TweetVoteAdminOperationIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return &TweetVoteAdminOperationIterator{contract: _TweetVote.contract, event: "AdminOperation", logs: logs, sub: sub}, nil
}

// WatchAdminOperation is a free log subscription operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_TweetVote *TweetVoteFilterer) WatchAdminOperation(opts *bind.WatchOpts, sink chan<- *TweetVoteAdminOperation) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteAdminOperation)
				if err := _TweetVote.contract.UnpackLog(event, "AdminOperation", log); err != nil {
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
func (_TweetVote *TweetVoteFilterer) ParseAdminOperation(log types.Log) (*TweetVoteAdminOperation, error) {
	event := new(TweetVoteAdminOperation)
	if err := _TweetVote.contract.UnpackLog(event, "AdminOperation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteKolRightsBoughtIterator is returned from FilterKolRightsBought and is used to iterate over the raw logs and unpacked data for KolRightsBought events raised by the TweetVote contract.
type TweetVoteKolRightsBoughtIterator struct {
	Event *TweetVoteKolRightsBought // Event containing the contract specifics and raw log

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
func (it *TweetVoteKolRightsBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteKolRightsBought)
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
		it.Event = new(TweetVoteKolRightsBought)
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
func (it *TweetVoteKolRightsBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteKolRightsBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteKolRightsBought represents a KolRightsBought event raised by the TweetVote contract.
type TweetVoteKolRightsBought struct {
	KolAddr  common.Address
	Buyer    common.Address
	RightsNo *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKolRightsBought is a free log retrieval operation binding the contract event 0x965468a49d826fce59d9e77b6be02c36a3de0e2bda03b078403f3139b0b8e29e.
//
// Solidity: event KolRightsBought(address kolAddr, address buyer, uint256 rightsNo)
func (_TweetVote *TweetVoteFilterer) FilterKolRightsBought(opts *bind.FilterOpts) (*TweetVoteKolRightsBoughtIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "KolRightsBought")
	if err != nil {
		return nil, err
	}
	return &TweetVoteKolRightsBoughtIterator{contract: _TweetVote.contract, event: "KolRightsBought", logs: logs, sub: sub}, nil
}

// WatchKolRightsBought is a free log subscription operation binding the contract event 0x965468a49d826fce59d9e77b6be02c36a3de0e2bda03b078403f3139b0b8e29e.
//
// Solidity: event KolRightsBought(address kolAddr, address buyer, uint256 rightsNo)
func (_TweetVote *TweetVoteFilterer) WatchKolRightsBought(opts *bind.WatchOpts, sink chan<- *TweetVoteKolRightsBought) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "KolRightsBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteKolRightsBought)
				if err := _TweetVote.contract.UnpackLog(event, "KolRightsBought", log); err != nil {
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

// ParseKolRightsBought is a log parse operation binding the contract event 0x965468a49d826fce59d9e77b6be02c36a3de0e2bda03b078403f3139b0b8e29e.
//
// Solidity: event KolRightsBought(address kolAddr, address buyer, uint256 rightsNo)
func (_TweetVote *TweetVoteFilterer) ParseKolRightsBought(log types.Log) (*TweetVoteKolRightsBought, error) {
	event := new(TweetVoteKolRightsBought)
	if err := _TweetVote.contract.UnpackLog(event, "KolRightsBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteKolWithdrawIterator is returned from FilterKolWithdraw and is used to iterate over the raw logs and unpacked data for KolWithdraw events raised by the TweetVote contract.
type TweetVoteKolWithdrawIterator struct {
	Event *TweetVoteKolWithdraw // Event containing the contract specifics and raw log

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
func (it *TweetVoteKolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteKolWithdraw)
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
		it.Event = new(TweetVoteKolWithdraw)
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
func (it *TweetVoteKolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteKolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteKolWithdraw represents a KolWithdraw event raised by the TweetVote contract.
type TweetVoteKolWithdraw struct {
	Kol    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKolWithdraw is a free log retrieval operation binding the contract event 0xcb69b6e91aa9504b27fc0545fcda1c4013efec7e04245bf33130a90a7e8277a8.
//
// Solidity: event KolWithdraw(address indexed kol, uint256 amount)
func (_TweetVote *TweetVoteFilterer) FilterKolWithdraw(opts *bind.FilterOpts, kol []common.Address) (*TweetVoteKolWithdrawIterator, error) {

	var kolRule []interface{}
	for _, kolItem := range kol {
		kolRule = append(kolRule, kolItem)
	}

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "KolWithdraw", kolRule)
	if err != nil {
		return nil, err
	}
	return &TweetVoteKolWithdrawIterator{contract: _TweetVote.contract, event: "KolWithdraw", logs: logs, sub: sub}, nil
}

// WatchKolWithdraw is a free log subscription operation binding the contract event 0xcb69b6e91aa9504b27fc0545fcda1c4013efec7e04245bf33130a90a7e8277a8.
//
// Solidity: event KolWithdraw(address indexed kol, uint256 amount)
func (_TweetVote *TweetVoteFilterer) WatchKolWithdraw(opts *bind.WatchOpts, sink chan<- *TweetVoteKolWithdraw, kol []common.Address) (event.Subscription, error) {

	var kolRule []interface{}
	for _, kolItem := range kol {
		kolRule = append(kolRule, kolItem)
	}

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "KolWithdraw", kolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteKolWithdraw)
				if err := _TweetVote.contract.UnpackLog(event, "KolWithdraw", log); err != nil {
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

// ParseKolWithdraw is a log parse operation binding the contract event 0xcb69b6e91aa9504b27fc0545fcda1c4013efec7e04245bf33130a90a7e8277a8.
//
// Solidity: event KolWithdraw(address indexed kol, uint256 amount)
func (_TweetVote *TweetVoteFilterer) ParseKolWithdraw(log types.Log) (*TweetVoteKolWithdraw, error) {
	event := new(TweetVoteKolWithdraw)
	if err := _TweetVote.contract.UnpackLog(event, "KolWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the TweetVote contract.
type TweetVoteOwnerSetIterator struct {
	Event *TweetVoteOwnerSet // Event containing the contract specifics and raw log

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
func (it *TweetVoteOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteOwnerSet)
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
		it.Event = new(TweetVoteOwnerSet)
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
func (it *TweetVoteOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteOwnerSet represents a OwnerSet event raised by the TweetVote contract.
type TweetVoteOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_TweetVote *TweetVoteFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*TweetVoteOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TweetVoteOwnerSetIterator{contract: _TweetVote.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_TweetVote *TweetVoteFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *TweetVoteOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteOwnerSet)
				if err := _TweetVote.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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
func (_TweetVote *TweetVoteFilterer) ParseOwnerSet(log types.Log) (*TweetVoteOwnerSet, error) {
	event := new(TweetVoteOwnerSet)
	if err := _TweetVote.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVotePluginChangedIterator is returned from FilterPluginChanged and is used to iterate over the raw logs and unpacked data for PluginChanged events raised by the TweetVote contract.
type TweetVotePluginChangedIterator struct {
	Event *TweetVotePluginChanged // Event containing the contract specifics and raw log

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
func (it *TweetVotePluginChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVotePluginChanged)
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
		it.Event = new(TweetVotePluginChanged)
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
func (it *TweetVotePluginChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVotePluginChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVotePluginChanged represents a PluginChanged event raised by the TweetVote contract.
type TweetVotePluginChanged struct {
	PAddr common.Address
	Stop  bool
	Typ   string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPluginChanged is a free log retrieval operation binding the contract event 0x1d236a1b586b708c1de53379e0b333ec2e36d883b3e162e8506e5351971174ac.
//
// Solidity: event PluginChanged(address pAddr, bool stop, string typ)
func (_TweetVote *TweetVoteFilterer) FilterPluginChanged(opts *bind.FilterOpts) (*TweetVotePluginChangedIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "PluginChanged")
	if err != nil {
		return nil, err
	}
	return &TweetVotePluginChangedIterator{contract: _TweetVote.contract, event: "PluginChanged", logs: logs, sub: sub}, nil
}

// WatchPluginChanged is a free log subscription operation binding the contract event 0x1d236a1b586b708c1de53379e0b333ec2e36d883b3e162e8506e5351971174ac.
//
// Solidity: event PluginChanged(address pAddr, bool stop, string typ)
func (_TweetVote *TweetVoteFilterer) WatchPluginChanged(opts *bind.WatchOpts, sink chan<- *TweetVotePluginChanged) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "PluginChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVotePluginChanged)
				if err := _TweetVote.contract.UnpackLog(event, "PluginChanged", log); err != nil {
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

// ParsePluginChanged is a log parse operation binding the contract event 0x1d236a1b586b708c1de53379e0b333ec2e36d883b3e162e8506e5351971174ac.
//
// Solidity: event PluginChanged(address pAddr, bool stop, string typ)
func (_TweetVote *TweetVoteFilterer) ParsePluginChanged(log types.Log) (*TweetVotePluginChanged, error) {
	event := new(TweetVotePluginChanged)
	if err := _TweetVote.contract.UnpackLog(event, "PluginChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the TweetVote contract.
type TweetVoteReceivedIterator struct {
	Event *TweetVoteReceived // Event containing the contract specifics and raw log

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
func (it *TweetVoteReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteReceived)
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
		it.Event = new(TweetVoteReceived)
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
func (it *TweetVoteReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteReceived represents a Received event raised by the TweetVote contract.
type TweetVoteReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 amount)
func (_TweetVote *TweetVoteFilterer) FilterReceived(opts *bind.FilterOpts, sender []common.Address) (*TweetVoteReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return &TweetVoteReceivedIterator{contract: _TweetVote.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 amount)
func (_TweetVote *TweetVoteFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *TweetVoteReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteReceived)
				if err := _TweetVote.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 amount)
func (_TweetVote *TweetVoteFilterer) ParseReceived(log types.Log) (*TweetVoteReceived, error) {
	event := new(TweetVoteReceived)
	if err := _TweetVote.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteServiceFeeChangedIterator is returned from FilterServiceFeeChanged and is used to iterate over the raw logs and unpacked data for ServiceFeeChanged events raised by the TweetVote contract.
type TweetVoteServiceFeeChangedIterator struct {
	Event *TweetVoteServiceFeeChanged // Event containing the contract specifics and raw log

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
func (it *TweetVoteServiceFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteServiceFeeChanged)
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
		it.Event = new(TweetVoteServiceFeeChanged)
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
func (it *TweetVoteServiceFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteServiceFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteServiceFeeChanged represents a ServiceFeeChanged event raised by the TweetVote contract.
type TweetVoteServiceFeeChanged struct {
	NewSerficeFeeRate *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeChanged is a free log retrieval operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_TweetVote *TweetVoteFilterer) FilterServiceFeeChanged(opts *bind.FilterOpts) (*TweetVoteServiceFeeChangedIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return &TweetVoteServiceFeeChangedIterator{contract: _TweetVote.contract, event: "ServiceFeeChanged", logs: logs, sub: sub}, nil
}

// WatchServiceFeeChanged is a free log subscription operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_TweetVote *TweetVoteFilterer) WatchServiceFeeChanged(opts *bind.WatchOpts, sink chan<- *TweetVoteServiceFeeChanged) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteServiceFeeChanged)
				if err := _TweetVote.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
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
func (_TweetVote *TweetVoteFilterer) ParseServiceFeeChanged(log types.Log) (*TweetVoteServiceFeeChanged, error) {
	event := new(TweetVoteServiceFeeChanged)
	if err := _TweetVote.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteSystemRateChangedIterator is returned from FilterSystemRateChanged and is used to iterate over the raw logs and unpacked data for SystemRateChanged events raised by the TweetVote contract.
type TweetVoteSystemRateChangedIterator struct {
	Event *TweetVoteSystemRateChanged // Event containing the contract specifics and raw log

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
func (it *TweetVoteSystemRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteSystemRateChanged)
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
		it.Event = new(TweetVoteSystemRateChanged)
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
func (it *TweetVoteSystemRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteSystemRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteSystemRateChanged represents a SystemRateChanged event raised by the TweetVote contract.
type TweetVoteSystemRateChanged struct {
	PricePost *big.Int
	RateName  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSystemRateChanged is a free log retrieval operation binding the contract event 0xbd22687be7d42e6c7dfcbfb3fe1d0288fd2d292251d6457912b9e914152ddcef.
//
// Solidity: event SystemRateChanged(uint256 pricePost, string rateName)
func (_TweetVote *TweetVoteFilterer) FilterSystemRateChanged(opts *bind.FilterOpts) (*TweetVoteSystemRateChangedIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "SystemRateChanged")
	if err != nil {
		return nil, err
	}
	return &TweetVoteSystemRateChangedIterator{contract: _TweetVote.contract, event: "SystemRateChanged", logs: logs, sub: sub}, nil
}

// WatchSystemRateChanged is a free log subscription operation binding the contract event 0xbd22687be7d42e6c7dfcbfb3fe1d0288fd2d292251d6457912b9e914152ddcef.
//
// Solidity: event SystemRateChanged(uint256 pricePost, string rateName)
func (_TweetVote *TweetVoteFilterer) WatchSystemRateChanged(opts *bind.WatchOpts, sink chan<- *TweetVoteSystemRateChanged) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "SystemRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteSystemRateChanged)
				if err := _TweetVote.contract.UnpackLog(event, "SystemRateChanged", log); err != nil {
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

// ParseSystemRateChanged is a log parse operation binding the contract event 0xbd22687be7d42e6c7dfcbfb3fe1d0288fd2d292251d6457912b9e914152ddcef.
//
// Solidity: event SystemRateChanged(uint256 pricePost, string rateName)
func (_TweetVote *TweetVoteFilterer) ParseSystemRateChanged(log types.Log) (*TweetVoteSystemRateChanged, error) {
	event := new(TweetVoteSystemRateChanged)
	if err := _TweetVote.contract.UnpackLog(event, "SystemRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteTweetPublishedIterator is returned from FilterTweetPublished and is used to iterate over the raw logs and unpacked data for TweetPublished events raised by the TweetVote contract.
type TweetVoteTweetPublishedIterator struct {
	Event *TweetVoteTweetPublished // Event containing the contract specifics and raw log

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
func (it *TweetVoteTweetPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteTweetPublished)
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
		it.Event = new(TweetVoteTweetPublished)
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
func (it *TweetVoteTweetPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteTweetPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteTweetPublished represents a TweetPublished event raised by the TweetVote contract.
type TweetVoteTweetPublished struct {
	From      common.Address
	TweetHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTweetPublished is a free log retrieval operation binding the contract event 0xc6d6ed5e39f5688dfcbb4af05f438a6eb7064336069d430bfd67c227eb04f55f.
//
// Solidity: event TweetPublished(address indexed from, bytes32 tweetHash)
func (_TweetVote *TweetVoteFilterer) FilterTweetPublished(opts *bind.FilterOpts, from []common.Address) (*TweetVoteTweetPublishedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "TweetPublished", fromRule)
	if err != nil {
		return nil, err
	}
	return &TweetVoteTweetPublishedIterator{contract: _TweetVote.contract, event: "TweetPublished", logs: logs, sub: sub}, nil
}

// WatchTweetPublished is a free log subscription operation binding the contract event 0xc6d6ed5e39f5688dfcbb4af05f438a6eb7064336069d430bfd67c227eb04f55f.
//
// Solidity: event TweetPublished(address indexed from, bytes32 tweetHash)
func (_TweetVote *TweetVoteFilterer) WatchTweetPublished(opts *bind.WatchOpts, sink chan<- *TweetVoteTweetPublished, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "TweetPublished", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteTweetPublished)
				if err := _TweetVote.contract.UnpackLog(event, "TweetPublished", log); err != nil {
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

// ParseTweetPublished is a log parse operation binding the contract event 0xc6d6ed5e39f5688dfcbb4af05f438a6eb7064336069d430bfd67c227eb04f55f.
//
// Solidity: event TweetPublished(address indexed from, bytes32 tweetHash)
func (_TweetVote *TweetVoteFilterer) ParseTweetPublished(log types.Log) (*TweetVoteTweetPublished, error) {
	event := new(TweetVoteTweetPublished)
	if err := _TweetVote.contract.UnpackLog(event, "TweetPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteTweetVotedIterator is returned from FilterTweetVoted and is used to iterate over the raw logs and unpacked data for TweetVoted events raised by the TweetVote contract.
type TweetVoteTweetVotedIterator struct {
	Event *TweetVoteTweetVoted // Event containing the contract specifics and raw log

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
func (it *TweetVoteTweetVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteTweetVoted)
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
		it.Event = new(TweetVoteTweetVoted)
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
func (it *TweetVoteTweetVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteTweetVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteTweetVoted represents a TweetVoted event raised by the TweetVote contract.
type TweetVoteTweetVoted struct {
	TweetHash    [32]byte
	Voter        common.Address
	PricePerVote *big.Int
	VoteNo       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTweetVoted is a free log retrieval operation binding the contract event 0xb67ade271cf9f89231289b7312378de3c5219f0e45581d2d38273f949dd114f6.
//
// Solidity: event TweetVoted(bytes32 tweetHash, address voter, uint256 pricePerVote, uint256 voteNo)
func (_TweetVote *TweetVoteFilterer) FilterTweetVoted(opts *bind.FilterOpts) (*TweetVoteTweetVotedIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "TweetVoted")
	if err != nil {
		return nil, err
	}
	return &TweetVoteTweetVotedIterator{contract: _TweetVote.contract, event: "TweetVoted", logs: logs, sub: sub}, nil
}

// WatchTweetVoted is a free log subscription operation binding the contract event 0xb67ade271cf9f89231289b7312378de3c5219f0e45581d2d38273f949dd114f6.
//
// Solidity: event TweetVoted(bytes32 tweetHash, address voter, uint256 pricePerVote, uint256 voteNo)
func (_TweetVote *TweetVoteFilterer) WatchTweetVoted(opts *bind.WatchOpts, sink chan<- *TweetVoteTweetVoted) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "TweetVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteTweetVoted)
				if err := _TweetVote.contract.UnpackLog(event, "TweetVoted", log); err != nil {
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

// ParseTweetVoted is a log parse operation binding the contract event 0xb67ade271cf9f89231289b7312378de3c5219f0e45581d2d38273f949dd114f6.
//
// Solidity: event TweetVoted(bytes32 tweetHash, address voter, uint256 pricePerVote, uint256 voteNo)
func (_TweetVote *TweetVoteFilterer) ParseTweetVoted(log types.Log) (*TweetVoteTweetVoted, error) {
	event := new(TweetVoteTweetVoted)
	if err := _TweetVote.contract.UnpackLog(event, "TweetVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteUpgradeToNewRuleIterator is returned from FilterUpgradeToNewRule and is used to iterate over the raw logs and unpacked data for UpgradeToNewRule events raised by the TweetVote contract.
type TweetVoteUpgradeToNewRuleIterator struct {
	Event *TweetVoteUpgradeToNewRule // Event containing the contract specifics and raw log

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
func (it *TweetVoteUpgradeToNewRuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteUpgradeToNewRule)
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
		it.Event = new(TweetVoteUpgradeToNewRule)
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
func (it *TweetVoteUpgradeToNewRuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteUpgradeToNewRuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteUpgradeToNewRule represents a UpgradeToNewRule event raised by the TweetVote contract.
type TweetVoteUpgradeToNewRule struct {
	NewContract common.Address
	Balance     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpgradeToNewRule is a free log retrieval operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_TweetVote *TweetVoteFilterer) FilterUpgradeToNewRule(opts *bind.FilterOpts) (*TweetVoteUpgradeToNewRuleIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return &TweetVoteUpgradeToNewRuleIterator{contract: _TweetVote.contract, event: "UpgradeToNewRule", logs: logs, sub: sub}, nil
}

// WatchUpgradeToNewRule is a free log subscription operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_TweetVote *TweetVoteFilterer) WatchUpgradeToNewRule(opts *bind.WatchOpts, sink chan<- *TweetVoteUpgradeToNewRule) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteUpgradeToNewRule)
				if err := _TweetVote.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
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
func (_TweetVote *TweetVoteFilterer) ParseUpgradeToNewRule(log types.Log) (*TweetVoteUpgradeToNewRule, error) {
	event := new(TweetVoteUpgradeToNewRule)
	if err := _TweetVote.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetVoteWithdrawServiceIterator is returned from FilterWithdrawService and is used to iterate over the raw logs and unpacked data for WithdrawService events raised by the TweetVote contract.
type TweetVoteWithdrawServiceIterator struct {
	Event *TweetVoteWithdrawService // Event containing the contract specifics and raw log

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
func (it *TweetVoteWithdrawServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetVoteWithdrawService)
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
		it.Event = new(TweetVoteWithdrawService)
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
func (it *TweetVoteWithdrawServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetVoteWithdrawServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetVoteWithdrawService represents a WithdrawService event raised by the TweetVote contract.
type TweetVoteWithdrawService struct {
	Owner   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawService is a free log retrieval operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_TweetVote *TweetVoteFilterer) FilterWithdrawService(opts *bind.FilterOpts) (*TweetVoteWithdrawServiceIterator, error) {

	logs, sub, err := _TweetVote.contract.FilterLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return &TweetVoteWithdrawServiceIterator{contract: _TweetVote.contract, event: "WithdrawService", logs: logs, sub: sub}, nil
}

// WatchWithdrawService is a free log subscription operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_TweetVote *TweetVoteFilterer) WatchWithdrawService(opts *bind.WatchOpts, sink chan<- *TweetVoteWithdrawService) (event.Subscription, error) {

	logs, sub, err := _TweetVote.contract.WatchLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetVoteWithdrawService)
				if err := _TweetVote.contract.UnpackLog(event, "WithdrawService", log); err != nil {
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
func (_TweetVote *TweetVoteFilterer) ParseWithdrawService(log types.Log) (*TweetVoteWithdrawService, error) {
	event := new(TweetVoteWithdrawService)
	if err := _TweetVote.contract.UnpackLog(event, "WithdrawService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
