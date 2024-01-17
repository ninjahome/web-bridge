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

// TweetLotteryGameGameInfoOneRound is an auto generated low-level Go binding around an user-defined struct.
type TweetLotteryGameGameInfoOneRound struct {
	RandomHash   [32]byte
	DiscoverTime *big.Int
	Winner       common.Address
	WinTeam      [32]byte
	WinTicketID  *big.Int
	Bonus        *big.Int
	RandomVal    *big.Int
}

// TweetLotteryGameMetaData contains all meta data concerning the TweetLotteryGame contract.
var TweetLotteryGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeInMinutes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"opName\",\"type\":\"string\"}],\"name\":\"AdminOperated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"opType\",\"type\":\"bool\"}],\"name\":\"AdminOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"winnerTeam\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticketID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonusToTeam\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextRandomHash\",\"type\":\"bytes32\"}],\"name\":\"DiscoverWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSerficeFeeRate\",\"type\":\"uint256\"}],\"name\":\"ServiceFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"SkipToNewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"TicketSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"thash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"}],\"name\":\"TweetBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"UpgradeToNewRule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"WithdrawService\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"__admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__bonusRateToWinner\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__currentLotteryTicketID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__lotteryGameRoundTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__minValCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__openToOuterPlayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__serviceFeeRateForTicketBuy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__ticketPriceForOuter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminChangeBonusRateToWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeInMinutes\",\"type\":\"uint256\"}],\"name\":\"adminChangeRoundTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"adminOpenToOuterPlayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDelete\",\"type\":\"bool\"}],\"name\":\"adminOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminServiceFeeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetServiceFeeRateForTicketBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInFinney\",\"type\":\"uint256\"}],\"name\":\"adminSetTicketPriceForOuter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newRate\",\"type\":\"uint8\"}],\"name\":\"adminSetWithdrawFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"adminUpgradeToNewRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNo\",\"type\":\"uint256\"}],\"name\":\"allTeamInfo\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"tweets\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"memCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"voteCounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNo\",\"type\":\"uint256\"}],\"name\":\"allTeamInfoNo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tweetNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCountNo\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketNo\",\"type\":\"uint256\"}],\"name\":\"buyTicketFromOuter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buyerInfoIdxForTickets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"buyerInfoRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"team\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPluginInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRoundNo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nextRoundRandomHash\",\"type\":\"bytes32\"}],\"name\":\"discoverWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameInfoRecord\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"randomHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"discoverTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"winTeam\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"winTicketID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomVal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"historyRoundInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"randomHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"discoverTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"winTeam\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"winTicketID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomVal\",\"type\":\"uint256\"}],\"internalType\":\"structTweetLotteryGame.GameInfoOneRound[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"skipToNextRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"teamList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNo\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tweet\",\"type\":\"bytes32\"}],\"name\":\"teamMembers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"voteNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"voteNos\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tickList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketsOfBuyer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketsRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tweetHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tweetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voteNo\",\"type\":\"uint256\"}],\"name\":\"tweetBought\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNo\",\"type\":\"uint256\"}],\"name\":\"tweetList\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"tweets\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNo\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tweet\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"memAddr\",\"type\":\"address\"}],\"name\":\"voteNoOfTeammate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"all\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TweetLotteryGameABI is the input ABI used to generate the binding from.
// Deprecated: Use TweetLotteryGameMetaData.ABI instead.
var TweetLotteryGameABI = TweetLotteryGameMetaData.ABI

// TweetLotteryGame is an auto generated Go binding around an Ethereum contract.
type TweetLotteryGame struct {
	TweetLotteryGameCaller     // Read-only binding to the contract
	TweetLotteryGameTransactor // Write-only binding to the contract
	TweetLotteryGameFilterer   // Log filterer for contract events
}

// TweetLotteryGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type TweetLotteryGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetLotteryGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TweetLotteryGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetLotteryGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TweetLotteryGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetLotteryGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TweetLotteryGameSession struct {
	Contract     *TweetLotteryGame // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TweetLotteryGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TweetLotteryGameCallerSession struct {
	Contract *TweetLotteryGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TweetLotteryGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TweetLotteryGameTransactorSession struct {
	Contract     *TweetLotteryGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TweetLotteryGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type TweetLotteryGameRaw struct {
	Contract *TweetLotteryGame // Generic contract binding to access the raw methods on
}

// TweetLotteryGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TweetLotteryGameCallerRaw struct {
	Contract *TweetLotteryGameCaller // Generic read-only contract binding to access the raw methods on
}

// TweetLotteryGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TweetLotteryGameTransactorRaw struct {
	Contract *TweetLotteryGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTweetLotteryGame creates a new instance of TweetLotteryGame, bound to a specific deployed contract.
func NewTweetLotteryGame(address common.Address, backend bind.ContractBackend) (*TweetLotteryGame, error) {
	contract, err := bindTweetLotteryGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGame{TweetLotteryGameCaller: TweetLotteryGameCaller{contract: contract}, TweetLotteryGameTransactor: TweetLotteryGameTransactor{contract: contract}, TweetLotteryGameFilterer: TweetLotteryGameFilterer{contract: contract}}, nil
}

// NewTweetLotteryGameCaller creates a new read-only instance of TweetLotteryGame, bound to a specific deployed contract.
func NewTweetLotteryGameCaller(address common.Address, caller bind.ContractCaller) (*TweetLotteryGameCaller, error) {
	contract, err := bindTweetLotteryGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameCaller{contract: contract}, nil
}

// NewTweetLotteryGameTransactor creates a new write-only instance of TweetLotteryGame, bound to a specific deployed contract.
func NewTweetLotteryGameTransactor(address common.Address, transactor bind.ContractTransactor) (*TweetLotteryGameTransactor, error) {
	contract, err := bindTweetLotteryGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameTransactor{contract: contract}, nil
}

// NewTweetLotteryGameFilterer creates a new log filterer instance of TweetLotteryGame, bound to a specific deployed contract.
func NewTweetLotteryGameFilterer(address common.Address, filterer bind.ContractFilterer) (*TweetLotteryGameFilterer, error) {
	contract, err := bindTweetLotteryGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameFilterer{contract: contract}, nil
}

// bindTweetLotteryGame binds a generic wrapper to an already deployed contract.
func bindTweetLotteryGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TweetLotteryGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TweetLotteryGame *TweetLotteryGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TweetLotteryGame.Contract.TweetLotteryGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TweetLotteryGame *TweetLotteryGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.TweetLotteryGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TweetLotteryGame *TweetLotteryGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.TweetLotteryGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TweetLotteryGame *TweetLotteryGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TweetLotteryGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TweetLotteryGame *TweetLotteryGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TweetLotteryGame *TweetLotteryGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_TweetLotteryGame *TweetLotteryGameCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_TweetLotteryGame *TweetLotteryGameSession) Admins(arg0 common.Address) (bool, error) {
	return _TweetLotteryGame.Contract.Admins(&_TweetLotteryGame.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x22b429f4.
//
// Solidity: function __admins(address ) view returns(bool)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _TweetLotteryGame.Contract.Admins(&_TweetLotteryGame.CallOpts, arg0)
}

// BonusRateToWinner is a free data retrieval call binding the contract method 0xd8574327.
//
// Solidity: function __bonusRateToWinner() view returns(uint8)
func (_TweetLotteryGame *TweetLotteryGameCaller) BonusRateToWinner(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__bonusRateToWinner")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BonusRateToWinner is a free data retrieval call binding the contract method 0xd8574327.
//
// Solidity: function __bonusRateToWinner() view returns(uint8)
func (_TweetLotteryGame *TweetLotteryGameSession) BonusRateToWinner() (uint8, error) {
	return _TweetLotteryGame.Contract.BonusRateToWinner(&_TweetLotteryGame.CallOpts)
}

// BonusRateToWinner is a free data retrieval call binding the contract method 0xd8574327.
//
// Solidity: function __bonusRateToWinner() view returns(uint8)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) BonusRateToWinner() (uint8, error) {
	return _TweetLotteryGame.Contract.BonusRateToWinner(&_TweetLotteryGame.CallOpts)
}

// CurrentLotteryTicketID is a free data retrieval call binding the contract method 0x3c607d6f.
//
// Solidity: function __currentLotteryTicketID() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) CurrentLotteryTicketID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__currentLotteryTicketID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentLotteryTicketID is a free data retrieval call binding the contract method 0x3c607d6f.
//
// Solidity: function __currentLotteryTicketID() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) CurrentLotteryTicketID() (*big.Int, error) {
	return _TweetLotteryGame.Contract.CurrentLotteryTicketID(&_TweetLotteryGame.CallOpts)
}

// CurrentLotteryTicketID is a free data retrieval call binding the contract method 0x3c607d6f.
//
// Solidity: function __currentLotteryTicketID() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) CurrentLotteryTicketID() (*big.Int, error) {
	return _TweetLotteryGame.Contract.CurrentLotteryTicketID(&_TweetLotteryGame.CallOpts)
}

// LotteryGameRoundTime is a free data retrieval call binding the contract method 0x1550ddfb.
//
// Solidity: function __lotteryGameRoundTime() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) LotteryGameRoundTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__lotteryGameRoundTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LotteryGameRoundTime is a free data retrieval call binding the contract method 0x1550ddfb.
//
// Solidity: function __lotteryGameRoundTime() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) LotteryGameRoundTime() (*big.Int, error) {
	return _TweetLotteryGame.Contract.LotteryGameRoundTime(&_TweetLotteryGame.CallOpts)
}

// LotteryGameRoundTime is a free data retrieval call binding the contract method 0x1550ddfb.
//
// Solidity: function __lotteryGameRoundTime() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) LotteryGameRoundTime() (*big.Int, error) {
	return _TweetLotteryGame.Contract.LotteryGameRoundTime(&_TweetLotteryGame.CallOpts)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) MinValCheck(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__minValCheck")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) MinValCheck() (*big.Int, error) {
	return _TweetLotteryGame.Contract.MinValCheck(&_TweetLotteryGame.CallOpts)
}

// MinValCheck is a free data retrieval call binding the contract method 0x8509614b.
//
// Solidity: function __minValCheck() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) MinValCheck() (*big.Int, error) {
	return _TweetLotteryGame.Contract.MinValCheck(&_TweetLotteryGame.CallOpts)
}

// OpenToOuterPlayer is a free data retrieval call binding the contract method 0x853cf537.
//
// Solidity: function __openToOuterPlayer() view returns(bool)
func (_TweetLotteryGame *TweetLotteryGameCaller) OpenToOuterPlayer(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__openToOuterPlayer")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenToOuterPlayer is a free data retrieval call binding the contract method 0x853cf537.
//
// Solidity: function __openToOuterPlayer() view returns(bool)
func (_TweetLotteryGame *TweetLotteryGameSession) OpenToOuterPlayer() (bool, error) {
	return _TweetLotteryGame.Contract.OpenToOuterPlayer(&_TweetLotteryGame.CallOpts)
}

// OpenToOuterPlayer is a free data retrieval call binding the contract method 0x853cf537.
//
// Solidity: function __openToOuterPlayer() view returns(bool)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) OpenToOuterPlayer() (bool, error) {
	return _TweetLotteryGame.Contract.OpenToOuterPlayer(&_TweetLotteryGame.CallOpts)
}

// ServiceFeeRateForTicketBuy is a free data retrieval call binding the contract method 0xd54de261.
//
// Solidity: function __serviceFeeRateForTicketBuy() view returns(uint8)
func (_TweetLotteryGame *TweetLotteryGameCaller) ServiceFeeRateForTicketBuy(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__serviceFeeRateForTicketBuy")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ServiceFeeRateForTicketBuy is a free data retrieval call binding the contract method 0xd54de261.
//
// Solidity: function __serviceFeeRateForTicketBuy() view returns(uint8)
func (_TweetLotteryGame *TweetLotteryGameSession) ServiceFeeRateForTicketBuy() (uint8, error) {
	return _TweetLotteryGame.Contract.ServiceFeeRateForTicketBuy(&_TweetLotteryGame.CallOpts)
}

// ServiceFeeRateForTicketBuy is a free data retrieval call binding the contract method 0xd54de261.
//
// Solidity: function __serviceFeeRateForTicketBuy() view returns(uint8)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) ServiceFeeRateForTicketBuy() (uint8, error) {
	return _TweetLotteryGame.Contract.ServiceFeeRateForTicketBuy(&_TweetLotteryGame.CallOpts)
}

// TicketPriceForOuter is a free data retrieval call binding the contract method 0xf524142e.
//
// Solidity: function __ticketPriceForOuter() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) TicketPriceForOuter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "__ticketPriceForOuter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketPriceForOuter is a free data retrieval call binding the contract method 0xf524142e.
//
// Solidity: function __ticketPriceForOuter() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) TicketPriceForOuter() (*big.Int, error) {
	return _TweetLotteryGame.Contract.TicketPriceForOuter(&_TweetLotteryGame.CallOpts)
}

// TicketPriceForOuter is a free data retrieval call binding the contract method 0xf524142e.
//
// Solidity: function __ticketPriceForOuter() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TicketPriceForOuter() (*big.Int, error) {
	return _TweetLotteryGame.Contract.TicketPriceForOuter(&_TweetLotteryGame.CallOpts)
}

// AllTeamInfo is a free data retrieval call binding the contract method 0xf6ffdc0e.
//
// Solidity: function allTeamInfo(uint256 roundNo) view returns(bytes32[] tweets, uint256[] memCounts, uint256[] voteCounts)
func (_TweetLotteryGame *TweetLotteryGameCaller) AllTeamInfo(opts *bind.CallOpts, roundNo *big.Int) (struct {
	Tweets     [][32]byte
	MemCounts  []*big.Int
	VoteCounts []*big.Int
}, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "allTeamInfo", roundNo)

	outstruct := new(struct {
		Tweets     [][32]byte
		MemCounts  []*big.Int
		VoteCounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tweets = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.MemCounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VoteCounts = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// AllTeamInfo is a free data retrieval call binding the contract method 0xf6ffdc0e.
//
// Solidity: function allTeamInfo(uint256 roundNo) view returns(bytes32[] tweets, uint256[] memCounts, uint256[] voteCounts)
func (_TweetLotteryGame *TweetLotteryGameSession) AllTeamInfo(roundNo *big.Int) (struct {
	Tweets     [][32]byte
	MemCounts  []*big.Int
	VoteCounts []*big.Int
}, error) {
	return _TweetLotteryGame.Contract.AllTeamInfo(&_TweetLotteryGame.CallOpts, roundNo)
}

// AllTeamInfo is a free data retrieval call binding the contract method 0xf6ffdc0e.
//
// Solidity: function allTeamInfo(uint256 roundNo) view returns(bytes32[] tweets, uint256[] memCounts, uint256[] voteCounts)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) AllTeamInfo(roundNo *big.Int) (struct {
	Tweets     [][32]byte
	MemCounts  []*big.Int
	VoteCounts []*big.Int
}, error) {
	return _TweetLotteryGame.Contract.AllTeamInfo(&_TweetLotteryGame.CallOpts, roundNo)
}

// AllTeamInfoNo is a free data retrieval call binding the contract method 0x8beb82e7.
//
// Solidity: function allTeamInfoNo(uint256 roundNo) view returns(uint256 tweetNo, uint256 voteCountNo)
func (_TweetLotteryGame *TweetLotteryGameCaller) AllTeamInfoNo(opts *bind.CallOpts, roundNo *big.Int) (struct {
	TweetNo     *big.Int
	VoteCountNo *big.Int
}, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "allTeamInfoNo", roundNo)

	outstruct := new(struct {
		TweetNo     *big.Int
		VoteCountNo *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TweetNo = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VoteCountNo = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllTeamInfoNo is a free data retrieval call binding the contract method 0x8beb82e7.
//
// Solidity: function allTeamInfoNo(uint256 roundNo) view returns(uint256 tweetNo, uint256 voteCountNo)
func (_TweetLotteryGame *TweetLotteryGameSession) AllTeamInfoNo(roundNo *big.Int) (struct {
	TweetNo     *big.Int
	VoteCountNo *big.Int
}, error) {
	return _TweetLotteryGame.Contract.AllTeamInfoNo(&_TweetLotteryGame.CallOpts, roundNo)
}

// AllTeamInfoNo is a free data retrieval call binding the contract method 0x8beb82e7.
//
// Solidity: function allTeamInfoNo(uint256 roundNo) view returns(uint256 tweetNo, uint256 voteCountNo)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) AllTeamInfoNo(roundNo *big.Int) (struct {
	TweetNo     *big.Int
	VoteCountNo *big.Int
}, error) {
	return _TweetLotteryGame.Contract.AllTeamInfoNo(&_TweetLotteryGame.CallOpts, roundNo)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _TweetLotteryGame.Contract.Balance(&_TweetLotteryGame.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _TweetLotteryGame.Contract.Balance(&_TweetLotteryGame.CallOpts, arg0)
}

// BuyerInfoIdxForTickets is a free data retrieval call binding the contract method 0x88df3b19.
//
// Solidity: function buyerInfoIdxForTickets(uint256 ) view returns(bytes32)
func (_TweetLotteryGame *TweetLotteryGameCaller) BuyerInfoIdxForTickets(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "buyerInfoIdxForTickets", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BuyerInfoIdxForTickets is a free data retrieval call binding the contract method 0x88df3b19.
//
// Solidity: function buyerInfoIdxForTickets(uint256 ) view returns(bytes32)
func (_TweetLotteryGame *TweetLotteryGameSession) BuyerInfoIdxForTickets(arg0 *big.Int) ([32]byte, error) {
	return _TweetLotteryGame.Contract.BuyerInfoIdxForTickets(&_TweetLotteryGame.CallOpts, arg0)
}

// BuyerInfoIdxForTickets is a free data retrieval call binding the contract method 0x88df3b19.
//
// Solidity: function buyerInfoIdxForTickets(uint256 ) view returns(bytes32)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) BuyerInfoIdxForTickets(arg0 *big.Int) ([32]byte, error) {
	return _TweetLotteryGame.Contract.BuyerInfoIdxForTickets(&_TweetLotteryGame.CallOpts, arg0)
}

// BuyerInfoRecords is a free data retrieval call binding the contract method 0xa90b6c22.
//
// Solidity: function buyerInfoRecords(bytes32 ) view returns(address addr, bytes32 team)
func (_TweetLotteryGame *TweetLotteryGameCaller) BuyerInfoRecords(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Addr common.Address
	Team [32]byte
}, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "buyerInfoRecords", arg0)

	outstruct := new(struct {
		Addr common.Address
		Team [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Team = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// BuyerInfoRecords is a free data retrieval call binding the contract method 0xa90b6c22.
//
// Solidity: function buyerInfoRecords(bytes32 ) view returns(address addr, bytes32 team)
func (_TweetLotteryGame *TweetLotteryGameSession) BuyerInfoRecords(arg0 [32]byte) (struct {
	Addr common.Address
	Team [32]byte
}, error) {
	return _TweetLotteryGame.Contract.BuyerInfoRecords(&_TweetLotteryGame.CallOpts, arg0)
}

// BuyerInfoRecords is a free data retrieval call binding the contract method 0xa90b6c22.
//
// Solidity: function buyerInfoRecords(bytes32 ) view returns(address addr, bytes32 team)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) BuyerInfoRecords(arg0 [32]byte) (struct {
	Addr common.Address
	Team [32]byte
}, error) {
	return _TweetLotteryGame.Contract.BuyerInfoRecords(&_TweetLotteryGame.CallOpts, arg0)
}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_TweetLotteryGame *TweetLotteryGameCaller) CheckPluginInterface(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "checkPluginInterface")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_TweetLotteryGame *TweetLotteryGameSession) CheckPluginInterface() (bool, error) {
	return _TweetLotteryGame.Contract.CheckPluginInterface(&_TweetLotteryGame.CallOpts)
}

// CheckPluginInterface is a free data retrieval call binding the contract method 0x807c758d.
//
// Solidity: function checkPluginInterface() pure returns(bool)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) CheckPluginInterface() (bool, error) {
	return _TweetLotteryGame.Contract.CheckPluginInterface(&_TweetLotteryGame.CallOpts)
}

// CurrentRoundNo is a free data retrieval call binding the contract method 0x1a8fb62e.
//
// Solidity: function currentRoundNo() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) CurrentRoundNo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "currentRoundNo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRoundNo is a free data retrieval call binding the contract method 0x1a8fb62e.
//
// Solidity: function currentRoundNo() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) CurrentRoundNo() (*big.Int, error) {
	return _TweetLotteryGame.Contract.CurrentRoundNo(&_TweetLotteryGame.CallOpts)
}

// CurrentRoundNo is a free data retrieval call binding the contract method 0x1a8fb62e.
//
// Solidity: function currentRoundNo() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) CurrentRoundNo() (*big.Int, error) {
	return _TweetLotteryGame.Contract.CurrentRoundNo(&_TweetLotteryGame.CallOpts)
}

// GameInfoRecord is a free data retrieval call binding the contract method 0x61373c31.
//
// Solidity: function gameInfoRecord(uint256 ) view returns(bytes32 randomHash, uint256 discoverTime, address winner, bytes32 winTeam, uint256 winTicketID, uint256 bonus, uint256 randomVal)
func (_TweetLotteryGame *TweetLotteryGameCaller) GameInfoRecord(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RandomHash   [32]byte
	DiscoverTime *big.Int
	Winner       common.Address
	WinTeam      [32]byte
	WinTicketID  *big.Int
	Bonus        *big.Int
	RandomVal    *big.Int
}, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "gameInfoRecord", arg0)

	outstruct := new(struct {
		RandomHash   [32]byte
		DiscoverTime *big.Int
		Winner       common.Address
		WinTeam      [32]byte
		WinTicketID  *big.Int
		Bonus        *big.Int
		RandomVal    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DiscoverTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.WinTeam = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.WinTicketID = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Bonus = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RandomVal = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GameInfoRecord is a free data retrieval call binding the contract method 0x61373c31.
//
// Solidity: function gameInfoRecord(uint256 ) view returns(bytes32 randomHash, uint256 discoverTime, address winner, bytes32 winTeam, uint256 winTicketID, uint256 bonus, uint256 randomVal)
func (_TweetLotteryGame *TweetLotteryGameSession) GameInfoRecord(arg0 *big.Int) (struct {
	RandomHash   [32]byte
	DiscoverTime *big.Int
	Winner       common.Address
	WinTeam      [32]byte
	WinTicketID  *big.Int
	Bonus        *big.Int
	RandomVal    *big.Int
}, error) {
	return _TweetLotteryGame.Contract.GameInfoRecord(&_TweetLotteryGame.CallOpts, arg0)
}

// GameInfoRecord is a free data retrieval call binding the contract method 0x61373c31.
//
// Solidity: function gameInfoRecord(uint256 ) view returns(bytes32 randomHash, uint256 discoverTime, address winner, bytes32 winTeam, uint256 winTicketID, uint256 bonus, uint256 randomVal)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) GameInfoRecord(arg0 *big.Int) (struct {
	RandomHash   [32]byte
	DiscoverTime *big.Int
	Winner       common.Address
	WinTeam      [32]byte
	WinTicketID  *big.Int
	Bonus        *big.Int
	RandomVal    *big.Int
}, error) {
	return _TweetLotteryGame.Contract.GameInfoRecord(&_TweetLotteryGame.CallOpts, arg0)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TweetLotteryGame *TweetLotteryGameCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TweetLotteryGame *TweetLotteryGameSession) GetOwner() (common.Address, error) {
	return _TweetLotteryGame.Contract.GetOwner(&_TweetLotteryGame.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) GetOwner() (common.Address, error) {
	return _TweetLotteryGame.Contract.GetOwner(&_TweetLotteryGame.CallOpts)
}

// HistoryRoundInfo is a free data retrieval call binding the contract method 0x355b6ec4.
//
// Solidity: function historyRoundInfo(uint256 from, uint256 to) view returns((bytes32,uint256,address,bytes32,uint256,uint256,uint256)[] infos)
func (_TweetLotteryGame *TweetLotteryGameCaller) HistoryRoundInfo(opts *bind.CallOpts, from *big.Int, to *big.Int) ([]TweetLotteryGameGameInfoOneRound, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "historyRoundInfo", from, to)

	if err != nil {
		return *new([]TweetLotteryGameGameInfoOneRound), err
	}

	out0 := *abi.ConvertType(out[0], new([]TweetLotteryGameGameInfoOneRound)).(*[]TweetLotteryGameGameInfoOneRound)

	return out0, err

}

// HistoryRoundInfo is a free data retrieval call binding the contract method 0x355b6ec4.
//
// Solidity: function historyRoundInfo(uint256 from, uint256 to) view returns((bytes32,uint256,address,bytes32,uint256,uint256,uint256)[] infos)
func (_TweetLotteryGame *TweetLotteryGameSession) HistoryRoundInfo(from *big.Int, to *big.Int) ([]TweetLotteryGameGameInfoOneRound, error) {
	return _TweetLotteryGame.Contract.HistoryRoundInfo(&_TweetLotteryGame.CallOpts, from, to)
}

// HistoryRoundInfo is a free data retrieval call binding the contract method 0x355b6ec4.
//
// Solidity: function historyRoundInfo(uint256 from, uint256 to) view returns((bytes32,uint256,address,bytes32,uint256,uint256,uint256)[] infos)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) HistoryRoundInfo(from *big.Int, to *big.Int) ([]TweetLotteryGameGameInfoOneRound, error) {
	return _TweetLotteryGame.Contract.HistoryRoundInfo(&_TweetLotteryGame.CallOpts, from, to)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) ServiceFeeReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "serviceFeeReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) ServiceFeeReceived() (*big.Int, error) {
	return _TweetLotteryGame.Contract.ServiceFeeReceived(&_TweetLotteryGame.CallOpts)
}

// ServiceFeeReceived is a free data retrieval call binding the contract method 0xbe4479e8.
//
// Solidity: function serviceFeeReceived() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) ServiceFeeReceived() (*big.Int, error) {
	return _TweetLotteryGame.Contract.ServiceFeeReceived(&_TweetLotteryGame.CallOpts)
}

// SystemSettings is a free data retrieval call binding the contract method 0x60b42f12.
//
// Solidity: function systemSettings() view returns(uint256, uint256, uint256, bool)
func (_TweetLotteryGame *TweetLotteryGameCaller) SystemSettings(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "systemSettings")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// SystemSettings is a free data retrieval call binding the contract method 0x60b42f12.
//
// Solidity: function systemSettings() view returns(uint256, uint256, uint256, bool)
func (_TweetLotteryGame *TweetLotteryGameSession) SystemSettings() (*big.Int, *big.Int, *big.Int, bool, error) {
	return _TweetLotteryGame.Contract.SystemSettings(&_TweetLotteryGame.CallOpts)
}

// SystemSettings is a free data retrieval call binding the contract method 0x60b42f12.
//
// Solidity: function systemSettings() view returns(uint256, uint256, uint256, bool)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) SystemSettings() (*big.Int, *big.Int, *big.Int, bool, error) {
	return _TweetLotteryGame.Contract.SystemSettings(&_TweetLotteryGame.CallOpts)
}

// TeamList is a free data retrieval call binding the contract method 0x9030d81e.
//
// Solidity: function teamList(uint256 , uint256 ) view returns(bytes32)
func (_TweetLotteryGame *TweetLotteryGameCaller) TeamList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "teamList", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TeamList is a free data retrieval call binding the contract method 0x9030d81e.
//
// Solidity: function teamList(uint256 , uint256 ) view returns(bytes32)
func (_TweetLotteryGame *TweetLotteryGameSession) TeamList(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _TweetLotteryGame.Contract.TeamList(&_TweetLotteryGame.CallOpts, arg0, arg1)
}

// TeamList is a free data retrieval call binding the contract method 0x9030d81e.
//
// Solidity: function teamList(uint256 , uint256 ) view returns(bytes32)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TeamList(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _TweetLotteryGame.Contract.TeamList(&_TweetLotteryGame.CallOpts, arg0, arg1)
}

// TeamMembers is a free data retrieval call binding the contract method 0xd157740f.
//
// Solidity: function teamMembers(uint256 roundNo, bytes32 tweet) view returns(uint256 voteNo, uint256 memNo, uint256[] voteNos, address[] members)
func (_TweetLotteryGame *TweetLotteryGameCaller) TeamMembers(opts *bind.CallOpts, roundNo *big.Int, tweet [32]byte) (struct {
	VoteNo  *big.Int
	MemNo   *big.Int
	VoteNos []*big.Int
	Members []common.Address
}, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "teamMembers", roundNo, tweet)

	outstruct := new(struct {
		VoteNo  *big.Int
		MemNo   *big.Int
		VoteNos []*big.Int
		Members []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VoteNo = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MemNo = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteNos = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Members = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// TeamMembers is a free data retrieval call binding the contract method 0xd157740f.
//
// Solidity: function teamMembers(uint256 roundNo, bytes32 tweet) view returns(uint256 voteNo, uint256 memNo, uint256[] voteNos, address[] members)
func (_TweetLotteryGame *TweetLotteryGameSession) TeamMembers(roundNo *big.Int, tweet [32]byte) (struct {
	VoteNo  *big.Int
	MemNo   *big.Int
	VoteNos []*big.Int
	Members []common.Address
}, error) {
	return _TweetLotteryGame.Contract.TeamMembers(&_TweetLotteryGame.CallOpts, roundNo, tweet)
}

// TeamMembers is a free data retrieval call binding the contract method 0xd157740f.
//
// Solidity: function teamMembers(uint256 roundNo, bytes32 tweet) view returns(uint256 voteNo, uint256 memNo, uint256[] voteNos, address[] members)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TeamMembers(roundNo *big.Int, tweet [32]byte) (struct {
	VoteNo  *big.Int
	MemNo   *big.Int
	VoteNos []*big.Int
	Members []common.Address
}, error) {
	return _TweetLotteryGame.Contract.TeamMembers(&_TweetLotteryGame.CallOpts, roundNo, tweet)
}

// TickList is a free data retrieval call binding the contract method 0xdd57903b.
//
// Solidity: function tickList(uint256 round, address owner) view returns(uint256[], bytes32[])
func (_TweetLotteryGame *TweetLotteryGameCaller) TickList(opts *bind.CallOpts, round *big.Int, owner common.Address) ([]*big.Int, [][32]byte, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "tickList", round, owner)

	if err != nil {
		return *new([]*big.Int), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return out0, out1, err

}

// TickList is a free data retrieval call binding the contract method 0xdd57903b.
//
// Solidity: function tickList(uint256 round, address owner) view returns(uint256[], bytes32[])
func (_TweetLotteryGame *TweetLotteryGameSession) TickList(round *big.Int, owner common.Address) ([]*big.Int, [][32]byte, error) {
	return _TweetLotteryGame.Contract.TickList(&_TweetLotteryGame.CallOpts, round, owner)
}

// TickList is a free data retrieval call binding the contract method 0xdd57903b.
//
// Solidity: function tickList(uint256 round, address owner) view returns(uint256[], bytes32[])
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TickList(round *big.Int, owner common.Address) ([]*big.Int, [][32]byte, error) {
	return _TweetLotteryGame.Contract.TickList(&_TweetLotteryGame.CallOpts, round, owner)
}

// TicketsOfBuyer is a free data retrieval call binding the contract method 0x562beff7.
//
// Solidity: function ticketsOfBuyer(uint256 , address , uint256 ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) TicketsOfBuyer(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "ticketsOfBuyer", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketsOfBuyer is a free data retrieval call binding the contract method 0x562beff7.
//
// Solidity: function ticketsOfBuyer(uint256 , address , uint256 ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) TicketsOfBuyer(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _TweetLotteryGame.Contract.TicketsOfBuyer(&_TweetLotteryGame.CallOpts, arg0, arg1, arg2)
}

// TicketsOfBuyer is a free data retrieval call binding the contract method 0x562beff7.
//
// Solidity: function ticketsOfBuyer(uint256 , address , uint256 ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TicketsOfBuyer(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _TweetLotteryGame.Contract.TicketsOfBuyer(&_TweetLotteryGame.CallOpts, arg0, arg1, arg2)
}

// TicketsRecords is a free data retrieval call binding the contract method 0x11a7d78c.
//
// Solidity: function ticketsRecords(uint256 , uint256 ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) TicketsRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "ticketsRecords", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketsRecords is a free data retrieval call binding the contract method 0x11a7d78c.
//
// Solidity: function ticketsRecords(uint256 , uint256 ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) TicketsRecords(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TweetLotteryGame.Contract.TicketsRecords(&_TweetLotteryGame.CallOpts, arg0, arg1)
}

// TicketsRecords is a free data retrieval call binding the contract method 0x11a7d78c.
//
// Solidity: function ticketsRecords(uint256 , uint256 ) view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TicketsRecords(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TweetLotteryGame.Contract.TicketsRecords(&_TweetLotteryGame.CallOpts, arg0, arg1)
}

// TotalBonus is a free data retrieval call binding the contract method 0xa8dd07dc.
//
// Solidity: function totalBonus() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) TotalBonus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "totalBonus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBonus is a free data retrieval call binding the contract method 0xa8dd07dc.
//
// Solidity: function totalBonus() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) TotalBonus() (*big.Int, error) {
	return _TweetLotteryGame.Contract.TotalBonus(&_TweetLotteryGame.CallOpts)
}

// TotalBonus is a free data retrieval call binding the contract method 0xa8dd07dc.
//
// Solidity: function totalBonus() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TotalBonus() (*big.Int, error) {
	return _TweetLotteryGame.Contract.TotalBonus(&_TweetLotteryGame.CallOpts)
}

// TweetList is a free data retrieval call binding the contract method 0x532a9051.
//
// Solidity: function tweetList(uint256 roundNo) view returns(bytes32[] tweets)
func (_TweetLotteryGame *TweetLotteryGameCaller) TweetList(opts *bind.CallOpts, roundNo *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "tweetList", roundNo)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// TweetList is a free data retrieval call binding the contract method 0x532a9051.
//
// Solidity: function tweetList(uint256 roundNo) view returns(bytes32[] tweets)
func (_TweetLotteryGame *TweetLotteryGameSession) TweetList(roundNo *big.Int) ([][32]byte, error) {
	return _TweetLotteryGame.Contract.TweetList(&_TweetLotteryGame.CallOpts, roundNo)
}

// TweetList is a free data retrieval call binding the contract method 0x532a9051.
//
// Solidity: function tweetList(uint256 roundNo) view returns(bytes32[] tweets)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) TweetList(roundNo *big.Int) ([][32]byte, error) {
	return _TweetLotteryGame.Contract.TweetList(&_TweetLotteryGame.CallOpts, roundNo)
}

// VoteNoOfTeammate is a free data retrieval call binding the contract method 0x04b22aea.
//
// Solidity: function voteNoOfTeammate(uint256 roundNo, bytes32 tweet, address memAddr) view returns(uint256, uint256, uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) VoteNoOfTeammate(opts *bind.CallOpts, roundNo *big.Int, tweet [32]byte, memAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "voteNoOfTeammate", roundNo, tweet, memAddr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// VoteNoOfTeammate is a free data retrieval call binding the contract method 0x04b22aea.
//
// Solidity: function voteNoOfTeammate(uint256 roundNo, bytes32 tweet, address memAddr) view returns(uint256, uint256, uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) VoteNoOfTeammate(roundNo *big.Int, tweet [32]byte, memAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _TweetLotteryGame.Contract.VoteNoOfTeammate(&_TweetLotteryGame.CallOpts, roundNo, tweet, memAddr)
}

// VoteNoOfTeammate is a free data retrieval call binding the contract method 0x04b22aea.
//
// Solidity: function voteNoOfTeammate(uint256 roundNo, bytes32 tweet, address memAddr) view returns(uint256, uint256, uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) VoteNoOfTeammate(roundNo *big.Int, tweet [32]byte, memAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _TweetLotteryGame.Contract.VoteNoOfTeammate(&_TweetLotteryGame.CallOpts, roundNo, tweet, memAddr)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCaller) WithdrawFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "withdrawFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameSession) WithdrawFeeRate() (*big.Int, error) {
	return _TweetLotteryGame.Contract.WithdrawFeeRate(&_TweetLotteryGame.CallOpts)
}

// WithdrawFeeRate is a free data retrieval call binding the contract method 0xea99e689.
//
// Solidity: function withdrawFeeRate() view returns(uint256)
func (_TweetLotteryGame *TweetLotteryGameCallerSession) WithdrawFeeRate() (*big.Int, error) {
	return _TweetLotteryGame.Contract.WithdrawFeeRate(&_TweetLotteryGame.CallOpts)
}

// AdminChangeBonusRateToWinner is a paid mutator transaction binding the contract method 0x13957a7c.
//
// Solidity: function adminChangeBonusRateToWinner(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminChangeBonusRateToWinner(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminChangeBonusRateToWinner", newRate)
}

// AdminChangeBonusRateToWinner is a paid mutator transaction binding the contract method 0x13957a7c.
//
// Solidity: function adminChangeBonusRateToWinner(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminChangeBonusRateToWinner(newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminChangeBonusRateToWinner(&_TweetLotteryGame.TransactOpts, newRate)
}

// AdminChangeBonusRateToWinner is a paid mutator transaction binding the contract method 0x13957a7c.
//
// Solidity: function adminChangeBonusRateToWinner(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminChangeBonusRateToWinner(newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminChangeBonusRateToWinner(&_TweetLotteryGame.TransactOpts, newRate)
}

// AdminChangeRoundTime is a paid mutator transaction binding the contract method 0x3a80f9e6.
//
// Solidity: function adminChangeRoundTime(uint256 newTimeInMinutes) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminChangeRoundTime(opts *bind.TransactOpts, newTimeInMinutes *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminChangeRoundTime", newTimeInMinutes)
}

// AdminChangeRoundTime is a paid mutator transaction binding the contract method 0x3a80f9e6.
//
// Solidity: function adminChangeRoundTime(uint256 newTimeInMinutes) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminChangeRoundTime(newTimeInMinutes *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminChangeRoundTime(&_TweetLotteryGame.TransactOpts, newTimeInMinutes)
}

// AdminChangeRoundTime is a paid mutator transaction binding the contract method 0x3a80f9e6.
//
// Solidity: function adminChangeRoundTime(uint256 newTimeInMinutes) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminChangeRoundTime(newTimeInMinutes *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminChangeRoundTime(&_TweetLotteryGame.TransactOpts, newTimeInMinutes)
}

// AdminOpenToOuterPlayer is a paid mutator transaction binding the contract method 0x275a18da.
//
// Solidity: function adminOpenToOuterPlayer(bool isOpen) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminOpenToOuterPlayer(opts *bind.TransactOpts, isOpen bool) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminOpenToOuterPlayer", isOpen)
}

// AdminOpenToOuterPlayer is a paid mutator transaction binding the contract method 0x275a18da.
//
// Solidity: function adminOpenToOuterPlayer(bool isOpen) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminOpenToOuterPlayer(isOpen bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminOpenToOuterPlayer(&_TweetLotteryGame.TransactOpts, isOpen)
}

// AdminOpenToOuterPlayer is a paid mutator transaction binding the contract method 0x275a18da.
//
// Solidity: function adminOpenToOuterPlayer(bool isOpen) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminOpenToOuterPlayer(isOpen bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminOpenToOuterPlayer(&_TweetLotteryGame.TransactOpts, isOpen)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminOperation(opts *bind.TransactOpts, admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminOperation", admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminOperation(&_TweetLotteryGame.TransactOpts, admin, isDelete)
}

// AdminOperation is a paid mutator transaction binding the contract method 0x2b508fde.
//
// Solidity: function adminOperation(address admin, bool isDelete) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminOperation(admin common.Address, isDelete bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminOperation(&_TweetLotteryGame.TransactOpts, admin, isDelete)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminServiceFeeWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminServiceFeeWithdraw")
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminServiceFeeWithdraw(&_TweetLotteryGame.TransactOpts)
}

// AdminServiceFeeWithdraw is a paid mutator transaction binding the contract method 0x615f54a1.
//
// Solidity: function adminServiceFeeWithdraw() returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminServiceFeeWithdraw() (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminServiceFeeWithdraw(&_TweetLotteryGame.TransactOpts)
}

// AdminSetServiceFeeRateForTicketBuy is a paid mutator transaction binding the contract method 0xda5ddd43.
//
// Solidity: function adminSetServiceFeeRateForTicketBuy(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminSetServiceFeeRateForTicketBuy(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminSetServiceFeeRateForTicketBuy", newRate)
}

// AdminSetServiceFeeRateForTicketBuy is a paid mutator transaction binding the contract method 0xda5ddd43.
//
// Solidity: function adminSetServiceFeeRateForTicketBuy(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminSetServiceFeeRateForTicketBuy(newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminSetServiceFeeRateForTicketBuy(&_TweetLotteryGame.TransactOpts, newRate)
}

// AdminSetServiceFeeRateForTicketBuy is a paid mutator transaction binding the contract method 0xda5ddd43.
//
// Solidity: function adminSetServiceFeeRateForTicketBuy(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminSetServiceFeeRateForTicketBuy(newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminSetServiceFeeRateForTicketBuy(&_TweetLotteryGame.TransactOpts, newRate)
}

// AdminSetTicketPriceForOuter is a paid mutator transaction binding the contract method 0x2a2871c5.
//
// Solidity: function adminSetTicketPriceForOuter(uint256 priceInFinney) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminSetTicketPriceForOuter(opts *bind.TransactOpts, priceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminSetTicketPriceForOuter", priceInFinney)
}

// AdminSetTicketPriceForOuter is a paid mutator transaction binding the contract method 0x2a2871c5.
//
// Solidity: function adminSetTicketPriceForOuter(uint256 priceInFinney) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminSetTicketPriceForOuter(priceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminSetTicketPriceForOuter(&_TweetLotteryGame.TransactOpts, priceInFinney)
}

// AdminSetTicketPriceForOuter is a paid mutator transaction binding the contract method 0x2a2871c5.
//
// Solidity: function adminSetTicketPriceForOuter(uint256 priceInFinney) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminSetTicketPriceForOuter(priceInFinney *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminSetTicketPriceForOuter(&_TweetLotteryGame.TransactOpts, priceInFinney)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminSetWithdrawFeeRate(opts *bind.TransactOpts, newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminSetWithdrawFeeRate", newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminSetWithdrawFeeRate(&_TweetLotteryGame.TransactOpts, newRate)
}

// AdminSetWithdrawFeeRate is a paid mutator transaction binding the contract method 0x94c681d6.
//
// Solidity: function adminSetWithdrawFeeRate(uint8 newRate) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminSetWithdrawFeeRate(newRate uint8) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminSetWithdrawFeeRate(&_TweetLotteryGame.TransactOpts, newRate)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) AdminUpgradeToNewRule(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "adminUpgradeToNewRule", recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminUpgradeToNewRule(&_TweetLotteryGame.TransactOpts, recipient)
}

// AdminUpgradeToNewRule is a paid mutator transaction binding the contract method 0xc75539d1.
//
// Solidity: function adminUpgradeToNewRule(address recipient) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) AdminUpgradeToNewRule(recipient common.Address) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.AdminUpgradeToNewRule(&_TweetLotteryGame.TransactOpts, recipient)
}

// BuyTicketFromOuter is a paid mutator transaction binding the contract method 0x2ce9df84.
//
// Solidity: function buyTicketFromOuter(uint256 ticketNo) payable returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) BuyTicketFromOuter(opts *bind.TransactOpts, ticketNo *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "buyTicketFromOuter", ticketNo)
}

// BuyTicketFromOuter is a paid mutator transaction binding the contract method 0x2ce9df84.
//
// Solidity: function buyTicketFromOuter(uint256 ticketNo) payable returns()
func (_TweetLotteryGame *TweetLotteryGameSession) BuyTicketFromOuter(ticketNo *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.BuyTicketFromOuter(&_TweetLotteryGame.TransactOpts, ticketNo)
}

// BuyTicketFromOuter is a paid mutator transaction binding the contract method 0x2ce9df84.
//
// Solidity: function buyTicketFromOuter(uint256 ticketNo) payable returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) BuyTicketFromOuter(ticketNo *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.BuyTicketFromOuter(&_TweetLotteryGame.TransactOpts, ticketNo)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.ChangeOwner(&_TweetLotteryGame.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.ChangeOwner(&_TweetLotteryGame.TransactOpts, newOwner)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) ChangeStatus(opts *bind.TransactOpts, stop bool) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "changeStatus", stop)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) ChangeStatus(stop bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.ChangeStatus(&_TweetLotteryGame.TransactOpts, stop)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x9b3de49b.
//
// Solidity: function changeStatus(bool stop) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) ChangeStatus(stop bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.ChangeStatus(&_TweetLotteryGame.TransactOpts, stop)
}

// DiscoverWinner is a paid mutator transaction binding the contract method 0xf427f4f9.
//
// Solidity: function discoverWinner(uint256 random, bytes32 nextRoundRandomHash) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) DiscoverWinner(opts *bind.TransactOpts, random *big.Int, nextRoundRandomHash [32]byte) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "discoverWinner", random, nextRoundRandomHash)
}

// DiscoverWinner is a paid mutator transaction binding the contract method 0xf427f4f9.
//
// Solidity: function discoverWinner(uint256 random, bytes32 nextRoundRandomHash) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) DiscoverWinner(random *big.Int, nextRoundRandomHash [32]byte) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.DiscoverWinner(&_TweetLotteryGame.TransactOpts, random, nextRoundRandomHash)
}

// DiscoverWinner is a paid mutator transaction binding the contract method 0xf427f4f9.
//
// Solidity: function discoverWinner(uint256 random, bytes32 nextRoundRandomHash) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) DiscoverWinner(random *big.Int, nextRoundRandomHash [32]byte) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.DiscoverWinner(&_TweetLotteryGame.TransactOpts, random, nextRoundRandomHash)
}

// SkipToNextRound is a paid mutator transaction binding the contract method 0xedf9938d.
//
// Solidity: function skipToNextRound(bytes32 hash) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) SkipToNextRound(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "skipToNextRound", hash)
}

// SkipToNextRound is a paid mutator transaction binding the contract method 0xedf9938d.
//
// Solidity: function skipToNextRound(bytes32 hash) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) SkipToNextRound(hash [32]byte) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.SkipToNextRound(&_TweetLotteryGame.TransactOpts, hash)
}

// SkipToNextRound is a paid mutator transaction binding the contract method 0xedf9938d.
//
// Solidity: function skipToNextRound(bytes32 hash) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) SkipToNextRound(hash [32]byte) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.SkipToNextRound(&_TweetLotteryGame.TransactOpts, hash)
}

// TweetBought is a paid mutator transaction binding the contract method 0xd2f2eeb7.
//
// Solidity: function tweetBought(bytes32 tweetHash, address tweetOwner, address buyer, uint256 voteNo) payable returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) TweetBought(opts *bind.TransactOpts, tweetHash [32]byte, tweetOwner common.Address, buyer common.Address, voteNo *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "tweetBought", tweetHash, tweetOwner, buyer, voteNo)
}

// TweetBought is a paid mutator transaction binding the contract method 0xd2f2eeb7.
//
// Solidity: function tweetBought(bytes32 tweetHash, address tweetOwner, address buyer, uint256 voteNo) payable returns()
func (_TweetLotteryGame *TweetLotteryGameSession) TweetBought(tweetHash [32]byte, tweetOwner common.Address, buyer common.Address, voteNo *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.TweetBought(&_TweetLotteryGame.TransactOpts, tweetHash, tweetOwner, buyer, voteNo)
}

// TweetBought is a paid mutator transaction binding the contract method 0xd2f2eeb7.
//
// Solidity: function tweetBought(bytes32 tweetHash, address tweetOwner, address buyer, uint256 voteNo) payable returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) TweetBought(tweetHash [32]byte, tweetOwner common.Address, buyer common.Address, voteNo *big.Int) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.TweetBought(&_TweetLotteryGame.TransactOpts, tweetHash, tweetOwner, buyer, voteNo)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, all bool) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.Transact(opts, "withdraw", amount, all)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_TweetLotteryGame *TweetLotteryGameSession) Withdraw(amount *big.Int, all bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.Withdraw(&_TweetLotteryGame.TransactOpts, amount, all)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool all) returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) Withdraw(amount *big.Int, all bool) (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.Withdraw(&_TweetLotteryGame.TransactOpts, amount, all)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TweetLotteryGame *TweetLotteryGameTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetLotteryGame.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TweetLotteryGame *TweetLotteryGameSession) Receive() (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.Receive(&_TweetLotteryGame.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TweetLotteryGame *TweetLotteryGameTransactorSession) Receive() (*types.Transaction, error) {
	return _TweetLotteryGame.Contract.Receive(&_TweetLotteryGame.TransactOpts)
}

// TweetLotteryGameAdminOperatedIterator is returned from FilterAdminOperated and is used to iterate over the raw logs and unpacked data for AdminOperated events raised by the TweetLotteryGame contract.
type TweetLotteryGameAdminOperatedIterator struct {
	Event *TweetLotteryGameAdminOperated // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameAdminOperatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameAdminOperated)
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
		it.Event = new(TweetLotteryGameAdminOperated)
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
func (it *TweetLotteryGameAdminOperatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameAdminOperatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameAdminOperated represents a AdminOperated event raised by the TweetLotteryGame contract.
type TweetLotteryGameAdminOperated struct {
	NewTimeInMinutes *big.Int
	OpName           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAdminOperated is a free log retrieval operation binding the contract event 0x39865e03587bbccdfe9981182110f07a7cedc3cd608edb6a4efaa14bcdb4469d.
//
// Solidity: event AdminOperated(uint256 newTimeInMinutes, string opName)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterAdminOperated(opts *bind.FilterOpts) (*TweetLotteryGameAdminOperatedIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "AdminOperated")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameAdminOperatedIterator{contract: _TweetLotteryGame.contract, event: "AdminOperated", logs: logs, sub: sub}, nil
}

// WatchAdminOperated is a free log subscription operation binding the contract event 0x39865e03587bbccdfe9981182110f07a7cedc3cd608edb6a4efaa14bcdb4469d.
//
// Solidity: event AdminOperated(uint256 newTimeInMinutes, string opName)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchAdminOperated(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameAdminOperated) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "AdminOperated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameAdminOperated)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "AdminOperated", log); err != nil {
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

// ParseAdminOperated is a log parse operation binding the contract event 0x39865e03587bbccdfe9981182110f07a7cedc3cd608edb6a4efaa14bcdb4469d.
//
// Solidity: event AdminOperated(uint256 newTimeInMinutes, string opName)
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseAdminOperated(log types.Log) (*TweetLotteryGameAdminOperated, error) {
	event := new(TweetLotteryGameAdminOperated)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "AdminOperated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameAdminOperationIterator is returned from FilterAdminOperation and is used to iterate over the raw logs and unpacked data for AdminOperation events raised by the TweetLotteryGame contract.
type TweetLotteryGameAdminOperationIterator struct {
	Event *TweetLotteryGameAdminOperation // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameAdminOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameAdminOperation)
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
		it.Event = new(TweetLotteryGameAdminOperation)
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
func (it *TweetLotteryGameAdminOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameAdminOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameAdminOperation represents a AdminOperation event raised by the TweetLotteryGame contract.
type TweetLotteryGameAdminOperation struct {
	Admin  common.Address
	OpType bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAdminOperation is a free log retrieval operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterAdminOperation(opts *bind.FilterOpts) (*TweetLotteryGameAdminOperationIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameAdminOperationIterator{contract: _TweetLotteryGame.contract, event: "AdminOperation", logs: logs, sub: sub}, nil
}

// WatchAdminOperation is a free log subscription operation binding the contract event 0xa97d45d2d22d38f81ac1eb6ab339a6e56470c4cef221796a485da6841b02769e.
//
// Solidity: event AdminOperation(address admin, bool opType)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchAdminOperation(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameAdminOperation) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "AdminOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameAdminOperation)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "AdminOperation", log); err != nil {
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
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseAdminOperation(log types.Log) (*TweetLotteryGameAdminOperation, error) {
	event := new(TweetLotteryGameAdminOperation)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "AdminOperation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameDiscoverWinnerIterator is returned from FilterDiscoverWinner and is used to iterate over the raw logs and unpacked data for DiscoverWinner events raised by the TweetLotteryGame contract.
type TweetLotteryGameDiscoverWinnerIterator struct {
	Event *TweetLotteryGameDiscoverWinner // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameDiscoverWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameDiscoverWinner)
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
		it.Event = new(TweetLotteryGameDiscoverWinner)
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
func (it *TweetLotteryGameDiscoverWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameDiscoverWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameDiscoverWinner represents a DiscoverWinner event raised by the TweetLotteryGame contract.
type TweetLotteryGameDiscoverWinner struct {
	Winner         common.Address
	WinnerTeam     [32]byte
	TicketID       *big.Int
	Bonus          *big.Int
	BonusToTeam    *big.Int
	Random         *big.Int
	NextRandomHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDiscoverWinner is a free log retrieval operation binding the contract event 0xb46eb63b055209e08f856513cdeef787d4ad696c7f36cec7d5dd16b8d4a04445.
//
// Solidity: event DiscoverWinner(address winner, bytes32 winnerTeam, uint256 ticketID, uint256 bonus, uint256 bonusToTeam, uint256 random, bytes32 nextRandomHash)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterDiscoverWinner(opts *bind.FilterOpts) (*TweetLotteryGameDiscoverWinnerIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "DiscoverWinner")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameDiscoverWinnerIterator{contract: _TweetLotteryGame.contract, event: "DiscoverWinner", logs: logs, sub: sub}, nil
}

// WatchDiscoverWinner is a free log subscription operation binding the contract event 0xb46eb63b055209e08f856513cdeef787d4ad696c7f36cec7d5dd16b8d4a04445.
//
// Solidity: event DiscoverWinner(address winner, bytes32 winnerTeam, uint256 ticketID, uint256 bonus, uint256 bonusToTeam, uint256 random, bytes32 nextRandomHash)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchDiscoverWinner(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameDiscoverWinner) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "DiscoverWinner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameDiscoverWinner)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "DiscoverWinner", log); err != nil {
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

// ParseDiscoverWinner is a log parse operation binding the contract event 0xb46eb63b055209e08f856513cdeef787d4ad696c7f36cec7d5dd16b8d4a04445.
//
// Solidity: event DiscoverWinner(address winner, bytes32 winnerTeam, uint256 ticketID, uint256 bonus, uint256 bonusToTeam, uint256 random, bytes32 nextRandomHash)
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseDiscoverWinner(log types.Log) (*TweetLotteryGameDiscoverWinner, error) {
	event := new(TweetLotteryGameDiscoverWinner)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "DiscoverWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the TweetLotteryGame contract.
type TweetLotteryGameOwnerSetIterator struct {
	Event *TweetLotteryGameOwnerSet // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameOwnerSet)
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
		it.Event = new(TweetLotteryGameOwnerSet)
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
func (it *TweetLotteryGameOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameOwnerSet represents a OwnerSet event raised by the TweetLotteryGame contract.
type TweetLotteryGameOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*TweetLotteryGameOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameOwnerSetIterator{contract: _TweetLotteryGame.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameOwnerSet)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseOwnerSet(log types.Log) (*TweetLotteryGameOwnerSet, error) {
	event := new(TweetLotteryGameOwnerSet)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameServiceFeeChangedIterator is returned from FilterServiceFeeChanged and is used to iterate over the raw logs and unpacked data for ServiceFeeChanged events raised by the TweetLotteryGame contract.
type TweetLotteryGameServiceFeeChangedIterator struct {
	Event *TweetLotteryGameServiceFeeChanged // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameServiceFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameServiceFeeChanged)
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
		it.Event = new(TweetLotteryGameServiceFeeChanged)
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
func (it *TweetLotteryGameServiceFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameServiceFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameServiceFeeChanged represents a ServiceFeeChanged event raised by the TweetLotteryGame contract.
type TweetLotteryGameServiceFeeChanged struct {
	NewSerficeFeeRate *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeChanged is a free log retrieval operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterServiceFeeChanged(opts *bind.FilterOpts) (*TweetLotteryGameServiceFeeChangedIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameServiceFeeChangedIterator{contract: _TweetLotteryGame.contract, event: "ServiceFeeChanged", logs: logs, sub: sub}, nil
}

// WatchServiceFeeChanged is a free log subscription operation binding the contract event 0x1c068decb3b5138b265d62b22c4c2d8191a2e0bd3745e97b5b0ff66fa852eca5.
//
// Solidity: event ServiceFeeChanged(uint256 newSerficeFeeRate)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchServiceFeeChanged(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameServiceFeeChanged) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "ServiceFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameServiceFeeChanged)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
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
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseServiceFeeChanged(log types.Log) (*TweetLotteryGameServiceFeeChanged, error) {
	event := new(TweetLotteryGameServiceFeeChanged)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "ServiceFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameSkipToNewRoundIterator is returned from FilterSkipToNewRound and is used to iterate over the raw logs and unpacked data for SkipToNewRound events raised by the TweetLotteryGame contract.
type TweetLotteryGameSkipToNewRoundIterator struct {
	Event *TweetLotteryGameSkipToNewRound // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameSkipToNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameSkipToNewRound)
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
		it.Event = new(TweetLotteryGameSkipToNewRound)
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
func (it *TweetLotteryGameSkipToNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameSkipToNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameSkipToNewRound represents a SkipToNewRound event raised by the TweetLotteryGame contract.
type TweetLotteryGameSkipToNewRound struct {
	Hash  [32]byte
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkipToNewRound is a free log retrieval operation binding the contract event 0xa48adfcb41b7cc4607437c02693423a81944370a5efef4b28e07f07d11652860.
//
// Solidity: event SkipToNewRound(bytes32 hash, uint256 round)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterSkipToNewRound(opts *bind.FilterOpts) (*TweetLotteryGameSkipToNewRoundIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "SkipToNewRound")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameSkipToNewRoundIterator{contract: _TweetLotteryGame.contract, event: "SkipToNewRound", logs: logs, sub: sub}, nil
}

// WatchSkipToNewRound is a free log subscription operation binding the contract event 0xa48adfcb41b7cc4607437c02693423a81944370a5efef4b28e07f07d11652860.
//
// Solidity: event SkipToNewRound(bytes32 hash, uint256 round)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchSkipToNewRound(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameSkipToNewRound) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "SkipToNewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameSkipToNewRound)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "SkipToNewRound", log); err != nil {
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

// ParseSkipToNewRound is a log parse operation binding the contract event 0xa48adfcb41b7cc4607437c02693423a81944370a5efef4b28e07f07d11652860.
//
// Solidity: event SkipToNewRound(bytes32 hash, uint256 round)
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseSkipToNewRound(log types.Log) (*TweetLotteryGameSkipToNewRound, error) {
	event := new(TweetLotteryGameSkipToNewRound)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "SkipToNewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameTicketSoldIterator is returned from FilterTicketSold and is used to iterate over the raw logs and unpacked data for TicketSold events raised by the TweetLotteryGame contract.
type TweetLotteryGameTicketSoldIterator struct {
	Event *TweetLotteryGameTicketSold // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameTicketSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameTicketSold)
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
		it.Event = new(TweetLotteryGameTicketSold)
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
func (it *TweetLotteryGameTicketSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameTicketSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameTicketSold represents a TicketSold event raised by the TweetLotteryGame contract.
type TweetLotteryGameTicketSold struct {
	Buyer      common.Address
	No         *big.Int
	ServiceFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTicketSold is a free log retrieval operation binding the contract event 0x7486e610a6d07e1419063632389e57b8c133ced5e70fea2d926ad3f99df98218.
//
// Solidity: event TicketSold(address buyer, uint256 no, uint256 serviceFee)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterTicketSold(opts *bind.FilterOpts) (*TweetLotteryGameTicketSoldIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "TicketSold")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameTicketSoldIterator{contract: _TweetLotteryGame.contract, event: "TicketSold", logs: logs, sub: sub}, nil
}

// WatchTicketSold is a free log subscription operation binding the contract event 0x7486e610a6d07e1419063632389e57b8c133ced5e70fea2d926ad3f99df98218.
//
// Solidity: event TicketSold(address buyer, uint256 no, uint256 serviceFee)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchTicketSold(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameTicketSold) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "TicketSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameTicketSold)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "TicketSold", log); err != nil {
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

// ParseTicketSold is a log parse operation binding the contract event 0x7486e610a6d07e1419063632389e57b8c133ced5e70fea2d926ad3f99df98218.
//
// Solidity: event TicketSold(address buyer, uint256 no, uint256 serviceFee)
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseTicketSold(log types.Log) (*TweetLotteryGameTicketSold, error) {
	event := new(TweetLotteryGameTicketSold)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "TicketSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameTweetBoughtIterator is returned from FilterTweetBought and is used to iterate over the raw logs and unpacked data for TweetBought events raised by the TweetLotteryGame contract.
type TweetLotteryGameTweetBoughtIterator struct {
	Event *TweetLotteryGameTweetBought // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameTweetBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameTweetBought)
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
		it.Event = new(TweetLotteryGameTweetBought)
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
func (it *TweetLotteryGameTweetBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameTweetBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameTweetBought represents a TweetBought event raised by the TweetLotteryGame contract.
type TweetLotteryGameTweetBought struct {
	Thash [32]byte
	Owner common.Address
	Buyer common.Address
	Val   *big.Int
	No    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTweetBought is a free log retrieval operation binding the contract event 0xd4752ba343214266339e13b616bfa296ee8c71658e4a89e060b8050bfd2030c0.
//
// Solidity: event TweetBought(bytes32 thash, address owner, address buyer, uint256 val, uint256 no)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterTweetBought(opts *bind.FilterOpts) (*TweetLotteryGameTweetBoughtIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "TweetBought")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameTweetBoughtIterator{contract: _TweetLotteryGame.contract, event: "TweetBought", logs: logs, sub: sub}, nil
}

// WatchTweetBought is a free log subscription operation binding the contract event 0xd4752ba343214266339e13b616bfa296ee8c71658e4a89e060b8050bfd2030c0.
//
// Solidity: event TweetBought(bytes32 thash, address owner, address buyer, uint256 val, uint256 no)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchTweetBought(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameTweetBought) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "TweetBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameTweetBought)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "TweetBought", log); err != nil {
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

// ParseTweetBought is a log parse operation binding the contract event 0xd4752ba343214266339e13b616bfa296ee8c71658e4a89e060b8050bfd2030c0.
//
// Solidity: event TweetBought(bytes32 thash, address owner, address buyer, uint256 val, uint256 no)
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseTweetBought(log types.Log) (*TweetLotteryGameTweetBought, error) {
	event := new(TweetLotteryGameTweetBought)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "TweetBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameUpgradeToNewRuleIterator is returned from FilterUpgradeToNewRule and is used to iterate over the raw logs and unpacked data for UpgradeToNewRule events raised by the TweetLotteryGame contract.
type TweetLotteryGameUpgradeToNewRuleIterator struct {
	Event *TweetLotteryGameUpgradeToNewRule // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameUpgradeToNewRuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameUpgradeToNewRule)
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
		it.Event = new(TweetLotteryGameUpgradeToNewRule)
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
func (it *TweetLotteryGameUpgradeToNewRuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameUpgradeToNewRuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameUpgradeToNewRule represents a UpgradeToNewRule event raised by the TweetLotteryGame contract.
type TweetLotteryGameUpgradeToNewRule struct {
	NewContract common.Address
	Balance     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpgradeToNewRule is a free log retrieval operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterUpgradeToNewRule(opts *bind.FilterOpts) (*TweetLotteryGameUpgradeToNewRuleIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameUpgradeToNewRuleIterator{contract: _TweetLotteryGame.contract, event: "UpgradeToNewRule", logs: logs, sub: sub}, nil
}

// WatchUpgradeToNewRule is a free log subscription operation binding the contract event 0x8ff69c992616f2fe31c491c4a2ba58d0da8ab5cea0e68ad7ea1713d6fecbcdaa.
//
// Solidity: event UpgradeToNewRule(address newContract, uint256 balance)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchUpgradeToNewRule(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameUpgradeToNewRule) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "UpgradeToNewRule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameUpgradeToNewRule)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
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
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseUpgradeToNewRule(log types.Log) (*TweetLotteryGameUpgradeToNewRule, error) {
	event := new(TweetLotteryGameUpgradeToNewRule)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "UpgradeToNewRule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TweetLotteryGameWithdrawServiceIterator is returned from FilterWithdrawService and is used to iterate over the raw logs and unpacked data for WithdrawService events raised by the TweetLotteryGame contract.
type TweetLotteryGameWithdrawServiceIterator struct {
	Event *TweetLotteryGameWithdrawService // Event containing the contract specifics and raw log

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
func (it *TweetLotteryGameWithdrawServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TweetLotteryGameWithdrawService)
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
		it.Event = new(TweetLotteryGameWithdrawService)
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
func (it *TweetLotteryGameWithdrawServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TweetLotteryGameWithdrawServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TweetLotteryGameWithdrawService represents a WithdrawService event raised by the TweetLotteryGame contract.
type TweetLotteryGameWithdrawService struct {
	Owner   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawService is a free log retrieval operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_TweetLotteryGame *TweetLotteryGameFilterer) FilterWithdrawService(opts *bind.FilterOpts) (*TweetLotteryGameWithdrawServiceIterator, error) {

	logs, sub, err := _TweetLotteryGame.contract.FilterLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return &TweetLotteryGameWithdrawServiceIterator{contract: _TweetLotteryGame.contract, event: "WithdrawService", logs: logs, sub: sub}, nil
}

// WatchWithdrawService is a free log subscription operation binding the contract event 0x7cb5cedbb78bc8cf00b5ffe051ef08d865e93d587654a591285add48befaa51d.
//
// Solidity: event WithdrawService(address owner, uint256 balance)
func (_TweetLotteryGame *TweetLotteryGameFilterer) WatchWithdrawService(opts *bind.WatchOpts, sink chan<- *TweetLotteryGameWithdrawService) (event.Subscription, error) {

	logs, sub, err := _TweetLotteryGame.contract.WatchLogs(opts, "WithdrawService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TweetLotteryGameWithdrawService)
				if err := _TweetLotteryGame.contract.UnpackLog(event, "WithdrawService", log); err != nil {
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
func (_TweetLotteryGame *TweetLotteryGameFilterer) ParseWithdrawService(log types.Log) (*TweetLotteryGameWithdrawService, error) {
	event := new(TweetLotteryGameWithdrawService)
	if err := _TweetLotteryGame.contract.UnpackLog(event, "WithdrawService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
