// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// StakeCreditUnbondRequest is an auto generated low-level Go binding around an user-defined struct.
type StakeCreditUnbondRequest struct {
	Shares     *big.Int
	BnbAmount  *big.Int
	UnlockTime *big.Int
}

// StakeCreditMetaData contains all meta data concerning the StakeCredit contract.
var StakeCreditMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BC_FUSION_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimableUnbondRequest\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeReward\",\"inputs\":[{\"name\":\"commissionRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getPooledBNB\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPooledBNBByShares\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByPooledBNB\",\"inputs\":[{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_moniker\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"lockedBNBs\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingUnbondRequest\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardRecord\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"slashBnbAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalPooledBNB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalPooledBNBRecord\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbond\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondRequest\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakeCredit.UnbondRequest\",\"components\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unbondSequence\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bnbAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardReceived\",\"inputs\":[{\"name\":\"rewardToAll\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"commission\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApproveNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NoClaimableUnbondRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoUnbondRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestExisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"WrongInitContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroTotalPooledBNB\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroTotalShares\",\"inputs\":[]}]",
}

// StakeCreditABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeCreditMetaData.ABI instead.
var StakeCreditABI = StakeCreditMetaData.ABI

// StakeCredit is an auto generated Go binding around an Ethereum contract.
type StakeCredit struct {
	StakeCreditCaller     // Read-only binding to the contract
	StakeCreditTransactor // Write-only binding to the contract
	StakeCreditFilterer   // Log filterer for contract events
}

// StakeCreditCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeCreditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeCreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeCreditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeCreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeCreditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeCreditSession struct {
	Contract     *StakeCredit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCreditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCreditCallerSession struct {
	Contract *StakeCreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StakeCreditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeCreditTransactorSession struct {
	Contract     *StakeCreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StakeCreditRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeCreditRaw struct {
	Contract *StakeCredit // Generic contract binding to access the raw methods on
}

// StakeCreditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCreditCallerRaw struct {
	Contract *StakeCreditCaller // Generic read-only contract binding to access the raw methods on
}

// StakeCreditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeCreditTransactorRaw struct {
	Contract *StakeCreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeCredit creates a new instance of StakeCredit, bound to a specific deployed contract.
func NewStakeCredit(address common.Address, backend bind.ContractBackend) (*StakeCredit, error) {
	contract, err := bindStakeCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeCredit{StakeCreditCaller: StakeCreditCaller{contract: contract}, StakeCreditTransactor: StakeCreditTransactor{contract: contract}, StakeCreditFilterer: StakeCreditFilterer{contract: contract}}, nil
}

// NewStakeCreditCaller creates a new read-only instance of StakeCredit, bound to a specific deployed contract.
func NewStakeCreditCaller(address common.Address, caller bind.ContractCaller) (*StakeCreditCaller, error) {
	contract, err := bindStakeCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCreditCaller{contract: contract}, nil
}

// NewStakeCreditTransactor creates a new write-only instance of StakeCredit, bound to a specific deployed contract.
func NewStakeCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeCreditTransactor, error) {
	contract, err := bindStakeCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCreditTransactor{contract: contract}, nil
}

// NewStakeCreditFilterer creates a new log filterer instance of StakeCredit, bound to a specific deployed contract.
func NewStakeCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeCreditFilterer, error) {
	contract, err := bindStakeCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeCreditFilterer{contract: contract}, nil
}

// bindStakeCredit binds a generic wrapper to an already deployed contract.
func bindStakeCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeCreditABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeCredit *StakeCreditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeCredit.Contract.StakeCreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeCredit *StakeCreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeCredit.Contract.StakeCreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeCredit *StakeCreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeCredit.Contract.StakeCreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeCredit *StakeCreditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeCredit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeCredit *StakeCreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeCredit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeCredit *StakeCreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeCredit.Contract.contract.Transact(opts, method, params...)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_StakeCredit *StakeCreditCaller) BCFUSIONCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "BC_FUSION_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_StakeCredit *StakeCreditSession) BCFUSIONCHANNELID() (uint8, error) {
	return _StakeCredit.Contract.BCFUSIONCHANNELID(&_StakeCredit.CallOpts)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_StakeCredit *StakeCreditCallerSession) BCFUSIONCHANNELID() (uint8, error) {
	return _StakeCredit.Contract.BCFUSIONCHANNELID(&_StakeCredit.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_StakeCredit *StakeCreditCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_StakeCredit *StakeCreditSession) STAKINGCHANNELID() (uint8, error) {
	return _StakeCredit.Contract.STAKINGCHANNELID(&_StakeCredit.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_StakeCredit *StakeCreditCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _StakeCredit.Contract.STAKINGCHANNELID(&_StakeCredit.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakeCredit *StakeCreditSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.Allowance(&_StakeCredit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.Allowance(&_StakeCredit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakeCredit *StakeCreditSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.BalanceOf(&_StakeCredit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.BalanceOf(&_StakeCredit.CallOpts, account)
}

// ClaimableUnbondRequest is a free data retrieval call binding the contract method 0x2f2d448a.
//
// Solidity: function claimableUnbondRequest(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) ClaimableUnbondRequest(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "claimableUnbondRequest", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableUnbondRequest is a free data retrieval call binding the contract method 0x2f2d448a.
//
// Solidity: function claimableUnbondRequest(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditSession) ClaimableUnbondRequest(delegator common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.ClaimableUnbondRequest(&_StakeCredit.CallOpts, delegator)
}

// ClaimableUnbondRequest is a free data retrieval call binding the contract method 0x2f2d448a.
//
// Solidity: function claimableUnbondRequest(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) ClaimableUnbondRequest(delegator common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.ClaimableUnbondRequest(&_StakeCredit.CallOpts, delegator)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakeCredit *StakeCreditCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakeCredit *StakeCreditSession) Decimals() (uint8, error) {
	return _StakeCredit.Contract.Decimals(&_StakeCredit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakeCredit *StakeCreditCallerSession) Decimals() (uint8, error) {
	return _StakeCredit.Contract.Decimals(&_StakeCredit.CallOpts)
}

// GetPooledBNB is a free data retrieval call binding the contract method 0x0913db47.
//
// Solidity: function getPooledBNB(address account) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) GetPooledBNB(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "getPooledBNB", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledBNB is a free data retrieval call binding the contract method 0x0913db47.
//
// Solidity: function getPooledBNB(address account) view returns(uint256)
func (_StakeCredit *StakeCreditSession) GetPooledBNB(account common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.GetPooledBNB(&_StakeCredit.CallOpts, account)
}

// GetPooledBNB is a free data retrieval call binding the contract method 0x0913db47.
//
// Solidity: function getPooledBNB(address account) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) GetPooledBNB(account common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.GetPooledBNB(&_StakeCredit.CallOpts, account)
}

// GetPooledBNBByShares is a free data retrieval call binding the contract method 0x91faf0b4.
//
// Solidity: function getPooledBNBByShares(uint256 shares) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) GetPooledBNBByShares(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "getPooledBNBByShares", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledBNBByShares is a free data retrieval call binding the contract method 0x91faf0b4.
//
// Solidity: function getPooledBNBByShares(uint256 shares) view returns(uint256)
func (_StakeCredit *StakeCreditSession) GetPooledBNBByShares(shares *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.GetPooledBNBByShares(&_StakeCredit.CallOpts, shares)
}

// GetPooledBNBByShares is a free data retrieval call binding the contract method 0x91faf0b4.
//
// Solidity: function getPooledBNBByShares(uint256 shares) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) GetPooledBNBByShares(shares *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.GetPooledBNBByShares(&_StakeCredit.CallOpts, shares)
}

// GetSharesByPooledBNB is a free data retrieval call binding the contract method 0x647df759.
//
// Solidity: function getSharesByPooledBNB(uint256 bnbAmount) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) GetSharesByPooledBNB(opts *bind.CallOpts, bnbAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "getSharesByPooledBNB", bnbAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledBNB is a free data retrieval call binding the contract method 0x647df759.
//
// Solidity: function getSharesByPooledBNB(uint256 bnbAmount) view returns(uint256)
func (_StakeCredit *StakeCreditSession) GetSharesByPooledBNB(bnbAmount *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.GetSharesByPooledBNB(&_StakeCredit.CallOpts, bnbAmount)
}

// GetSharesByPooledBNB is a free data retrieval call binding the contract method 0x647df759.
//
// Solidity: function getSharesByPooledBNB(uint256 bnbAmount) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) GetSharesByPooledBNB(bnbAmount *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.GetSharesByPooledBNB(&_StakeCredit.CallOpts, bnbAmount)
}

// LockedBNBs is a free data retrieval call binding the contract method 0xa9664feb.
//
// Solidity: function lockedBNBs(address delegator, uint256 number) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) LockedBNBs(opts *bind.CallOpts, delegator common.Address, number *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "lockedBNBs", delegator, number)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBNBs is a free data retrieval call binding the contract method 0xa9664feb.
//
// Solidity: function lockedBNBs(address delegator, uint256 number) view returns(uint256)
func (_StakeCredit *StakeCreditSession) LockedBNBs(delegator common.Address, number *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.LockedBNBs(&_StakeCredit.CallOpts, delegator, number)
}

// LockedBNBs is a free data retrieval call binding the contract method 0xa9664feb.
//
// Solidity: function lockedBNBs(address delegator, uint256 number) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) LockedBNBs(delegator common.Address, number *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.LockedBNBs(&_StakeCredit.CallOpts, delegator, number)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeCredit *StakeCreditCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeCredit *StakeCreditSession) Name() (string, error) {
	return _StakeCredit.Contract.Name(&_StakeCredit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeCredit *StakeCreditCallerSession) Name() (string, error) {
	return _StakeCredit.Contract.Name(&_StakeCredit.CallOpts)
}

// PendingUnbondRequest is a free data retrieval call binding the contract method 0x038c0023.
//
// Solidity: function pendingUnbondRequest(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) PendingUnbondRequest(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "pendingUnbondRequest", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingUnbondRequest is a free data retrieval call binding the contract method 0x038c0023.
//
// Solidity: function pendingUnbondRequest(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditSession) PendingUnbondRequest(delegator common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.PendingUnbondRequest(&_StakeCredit.CallOpts, delegator)
}

// PendingUnbondRequest is a free data retrieval call binding the contract method 0x038c0023.
//
// Solidity: function pendingUnbondRequest(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) PendingUnbondRequest(delegator common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.PendingUnbondRequest(&_StakeCredit.CallOpts, delegator)
}

// RewardRecord is a free data retrieval call binding the contract method 0xaa1966cd.
//
// Solidity: function rewardRecord(uint256 ) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) RewardRecord(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "rewardRecord", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRecord is a free data retrieval call binding the contract method 0xaa1966cd.
//
// Solidity: function rewardRecord(uint256 ) view returns(uint256)
func (_StakeCredit *StakeCreditSession) RewardRecord(arg0 *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.RewardRecord(&_StakeCredit.CallOpts, arg0)
}

// RewardRecord is a free data retrieval call binding the contract method 0xaa1966cd.
//
// Solidity: function rewardRecord(uint256 ) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) RewardRecord(arg0 *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.RewardRecord(&_StakeCredit.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeCredit *StakeCreditCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeCredit *StakeCreditSession) Symbol() (string, error) {
	return _StakeCredit.Contract.Symbol(&_StakeCredit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeCredit *StakeCreditCallerSession) Symbol() (string, error) {
	return _StakeCredit.Contract.Symbol(&_StakeCredit.CallOpts)
}

// TotalPooledBNB is a free data retrieval call binding the contract method 0x15d1f898.
//
// Solidity: function totalPooledBNB() view returns(uint256)
func (_StakeCredit *StakeCreditCaller) TotalPooledBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "totalPooledBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPooledBNB is a free data retrieval call binding the contract method 0x15d1f898.
//
// Solidity: function totalPooledBNB() view returns(uint256)
func (_StakeCredit *StakeCreditSession) TotalPooledBNB() (*big.Int, error) {
	return _StakeCredit.Contract.TotalPooledBNB(&_StakeCredit.CallOpts)
}

// TotalPooledBNB is a free data retrieval call binding the contract method 0x15d1f898.
//
// Solidity: function totalPooledBNB() view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) TotalPooledBNB() (*big.Int, error) {
	return _StakeCredit.Contract.TotalPooledBNB(&_StakeCredit.CallOpts)
}

// TotalPooledBNBRecord is a free data retrieval call binding the contract method 0x6bbf2249.
//
// Solidity: function totalPooledBNBRecord(uint256 ) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) TotalPooledBNBRecord(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "totalPooledBNBRecord", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPooledBNBRecord is a free data retrieval call binding the contract method 0x6bbf2249.
//
// Solidity: function totalPooledBNBRecord(uint256 ) view returns(uint256)
func (_StakeCredit *StakeCreditSession) TotalPooledBNBRecord(arg0 *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.TotalPooledBNBRecord(&_StakeCredit.CallOpts, arg0)
}

// TotalPooledBNBRecord is a free data retrieval call binding the contract method 0x6bbf2249.
//
// Solidity: function totalPooledBNBRecord(uint256 ) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) TotalPooledBNBRecord(arg0 *big.Int) (*big.Int, error) {
	return _StakeCredit.Contract.TotalPooledBNBRecord(&_StakeCredit.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeCredit *StakeCreditCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeCredit *StakeCreditSession) TotalSupply() (*big.Int, error) {
	return _StakeCredit.Contract.TotalSupply(&_StakeCredit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) TotalSupply() (*big.Int, error) {
	return _StakeCredit.Contract.TotalSupply(&_StakeCredit.CallOpts)
}

// UnbondRequest is a free data retrieval call binding the contract method 0xd241c1ea.
//
// Solidity: function unbondRequest(address delegator, uint256 _index) view returns((uint256,uint256,uint256))
func (_StakeCredit *StakeCreditCaller) UnbondRequest(opts *bind.CallOpts, delegator common.Address, _index *big.Int) (StakeCreditUnbondRequest, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "unbondRequest", delegator, _index)

	if err != nil {
		return *new(StakeCreditUnbondRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(StakeCreditUnbondRequest)).(*StakeCreditUnbondRequest)

	return out0, err

}

// UnbondRequest is a free data retrieval call binding the contract method 0xd241c1ea.
//
// Solidity: function unbondRequest(address delegator, uint256 _index) view returns((uint256,uint256,uint256))
func (_StakeCredit *StakeCreditSession) UnbondRequest(delegator common.Address, _index *big.Int) (StakeCreditUnbondRequest, error) {
	return _StakeCredit.Contract.UnbondRequest(&_StakeCredit.CallOpts, delegator, _index)
}

// UnbondRequest is a free data retrieval call binding the contract method 0xd241c1ea.
//
// Solidity: function unbondRequest(address delegator, uint256 _index) view returns((uint256,uint256,uint256))
func (_StakeCredit *StakeCreditCallerSession) UnbondRequest(delegator common.Address, _index *big.Int) (StakeCreditUnbondRequest, error) {
	return _StakeCredit.Contract.UnbondRequest(&_StakeCredit.CallOpts, delegator, _index)
}

// UnbondSequence is a free data retrieval call binding the contract method 0xc2cde2b2.
//
// Solidity: function unbondSequence(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditCaller) UnbondSequence(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "unbondSequence", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondSequence is a free data retrieval call binding the contract method 0xc2cde2b2.
//
// Solidity: function unbondSequence(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditSession) UnbondSequence(delegator common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.UnbondSequence(&_StakeCredit.CallOpts, delegator)
}

// UnbondSequence is a free data retrieval call binding the contract method 0xc2cde2b2.
//
// Solidity: function unbondSequence(address delegator) view returns(uint256)
func (_StakeCredit *StakeCreditCallerSession) UnbondSequence(delegator common.Address) (*big.Int, error) {
	return _StakeCredit.Contract.UnbondSequence(&_StakeCredit.CallOpts, delegator)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_StakeCredit *StakeCreditCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeCredit.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_StakeCredit *StakeCreditSession) Validator() (common.Address, error) {
	return _StakeCredit.Contract.Validator(&_StakeCredit.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_StakeCredit *StakeCreditCallerSession) Validator() (common.Address, error) {
	return _StakeCredit.Contract.Validator(&_StakeCredit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Approve(&_StakeCredit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Approve(&_StakeCredit.TransactOpts, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address delegator, uint256 number) returns(uint256)
func (_StakeCredit *StakeCreditTransactor) Claim(opts *bind.TransactOpts, delegator common.Address, number *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "claim", delegator, number)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address delegator, uint256 number) returns(uint256)
func (_StakeCredit *StakeCreditSession) Claim(delegator common.Address, number *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Claim(&_StakeCredit.TransactOpts, delegator, number)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address delegator, uint256 number) returns(uint256)
func (_StakeCredit *StakeCreditTransactorSession) Claim(delegator common.Address, number *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Claim(&_StakeCredit.TransactOpts, delegator, number)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakeCredit *StakeCreditTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakeCredit *StakeCreditSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.DecreaseAllowance(&_StakeCredit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakeCredit *StakeCreditTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.DecreaseAllowance(&_StakeCredit.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegator) payable returns(uint256 shares)
func (_StakeCredit *StakeCreditTransactor) Delegate(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "delegate", delegator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegator) payable returns(uint256 shares)
func (_StakeCredit *StakeCreditSession) Delegate(delegator common.Address) (*types.Transaction, error) {
	return _StakeCredit.Contract.Delegate(&_StakeCredit.TransactOpts, delegator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegator) payable returns(uint256 shares)
func (_StakeCredit *StakeCreditTransactorSession) Delegate(delegator common.Address) (*types.Transaction, error) {
	return _StakeCredit.Contract.Delegate(&_StakeCredit.TransactOpts, delegator)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x5e607d76.
//
// Solidity: function distributeReward(uint64 commissionRate) payable returns()
func (_StakeCredit *StakeCreditTransactor) DistributeReward(opts *bind.TransactOpts, commissionRate uint64) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "distributeReward", commissionRate)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x5e607d76.
//
// Solidity: function distributeReward(uint64 commissionRate) payable returns()
func (_StakeCredit *StakeCreditSession) DistributeReward(commissionRate uint64) (*types.Transaction, error) {
	return _StakeCredit.Contract.DistributeReward(&_StakeCredit.TransactOpts, commissionRate)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x5e607d76.
//
// Solidity: function distributeReward(uint64 commissionRate) payable returns()
func (_StakeCredit *StakeCreditTransactorSession) DistributeReward(commissionRate uint64) (*types.Transaction, error) {
	return _StakeCredit.Contract.DistributeReward(&_StakeCredit.TransactOpts, commissionRate)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakeCredit *StakeCreditTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakeCredit *StakeCreditSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.IncreaseAllowance(&_StakeCredit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakeCredit *StakeCreditTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.IncreaseAllowance(&_StakeCredit.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _validator, string _moniker) payable returns()
func (_StakeCredit *StakeCreditTransactor) Initialize(opts *bind.TransactOpts, _validator common.Address, _moniker string) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "initialize", _validator, _moniker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _validator, string _moniker) payable returns()
func (_StakeCredit *StakeCreditSession) Initialize(_validator common.Address, _moniker string) (*types.Transaction, error) {
	return _StakeCredit.Contract.Initialize(&_StakeCredit.TransactOpts, _validator, _moniker)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _validator, string _moniker) payable returns()
func (_StakeCredit *StakeCreditTransactorSession) Initialize(_validator common.Address, _moniker string) (*types.Transaction, error) {
	return _StakeCredit.Contract.Initialize(&_StakeCredit.TransactOpts, _validator, _moniker)
}

// Slash is a paid mutator transaction binding the contract method 0x45bc4d10.
//
// Solidity: function slash(uint256 slashBnbAmount) returns(uint256)
func (_StakeCredit *StakeCreditTransactor) Slash(opts *bind.TransactOpts, slashBnbAmount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "slash", slashBnbAmount)
}

// Slash is a paid mutator transaction binding the contract method 0x45bc4d10.
//
// Solidity: function slash(uint256 slashBnbAmount) returns(uint256)
func (_StakeCredit *StakeCreditSession) Slash(slashBnbAmount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Slash(&_StakeCredit.TransactOpts, slashBnbAmount)
}

// Slash is a paid mutator transaction binding the contract method 0x45bc4d10.
//
// Solidity: function slash(uint256 slashBnbAmount) returns(uint256)
func (_StakeCredit *StakeCreditTransactorSession) Slash(slashBnbAmount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Slash(&_StakeCredit.TransactOpts, slashBnbAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Transfer(&_StakeCredit.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Transfer(&_StakeCredit.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.TransferFrom(&_StakeCredit.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StakeCredit *StakeCreditTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.TransferFrom(&_StakeCredit.TransactOpts, from, to, amount)
}

// Unbond is a paid mutator transaction binding the contract method 0xa5d059ca.
//
// Solidity: function unbond(address delegator, uint256 shares) returns(uint256 bnbAmount)
func (_StakeCredit *StakeCreditTransactor) Unbond(opts *bind.TransactOpts, delegator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "unbond", delegator, shares)
}

// Unbond is a paid mutator transaction binding the contract method 0xa5d059ca.
//
// Solidity: function unbond(address delegator, uint256 shares) returns(uint256 bnbAmount)
func (_StakeCredit *StakeCreditSession) Unbond(delegator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Unbond(&_StakeCredit.TransactOpts, delegator, shares)
}

// Unbond is a paid mutator transaction binding the contract method 0xa5d059ca.
//
// Solidity: function unbond(address delegator, uint256 shares) returns(uint256 bnbAmount)
func (_StakeCredit *StakeCreditTransactorSession) Unbond(delegator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Unbond(&_StakeCredit.TransactOpts, delegator, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address delegator, uint256 shares) returns(uint256 bnbAmount)
func (_StakeCredit *StakeCreditTransactor) Undelegate(opts *bind.TransactOpts, delegator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StakeCredit.contract.Transact(opts, "undelegate", delegator, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address delegator, uint256 shares) returns(uint256 bnbAmount)
func (_StakeCredit *StakeCreditSession) Undelegate(delegator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Undelegate(&_StakeCredit.TransactOpts, delegator, shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address delegator, uint256 shares) returns(uint256 bnbAmount)
func (_StakeCredit *StakeCreditTransactorSession) Undelegate(delegator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StakeCredit.Contract.Undelegate(&_StakeCredit.TransactOpts, delegator, shares)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakeCredit *StakeCreditTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeCredit.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakeCredit *StakeCreditSession) Receive() (*types.Transaction, error) {
	return _StakeCredit.Contract.Receive(&_StakeCredit.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakeCredit *StakeCreditTransactorSession) Receive() (*types.Transaction, error) {
	return _StakeCredit.Contract.Receive(&_StakeCredit.TransactOpts)
}

// StakeCreditApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakeCredit contract.
type StakeCreditApprovalIterator struct {
	Event *StakeCreditApproval // Event containing the contract specifics and raw log

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
func (it *StakeCreditApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeCreditApproval)
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
		it.Event = new(StakeCreditApproval)
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
func (it *StakeCreditApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeCreditApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeCreditApproval represents a Approval event raised by the StakeCredit contract.
type StakeCreditApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakeCredit *StakeCreditFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakeCreditApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakeCredit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakeCreditApprovalIterator{contract: _StakeCredit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakeCredit *StakeCreditFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakeCreditApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakeCredit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeCreditApproval)
				if err := _StakeCredit.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakeCredit *StakeCreditFilterer) ParseApproval(log types.Log) (*StakeCreditApproval, error) {
	event := new(StakeCreditApproval)
	if err := _StakeCredit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeCreditInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakeCredit contract.
type StakeCreditInitializedIterator struct {
	Event *StakeCreditInitialized // Event containing the contract specifics and raw log

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
func (it *StakeCreditInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeCreditInitialized)
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
		it.Event = new(StakeCreditInitialized)
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
func (it *StakeCreditInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeCreditInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeCreditInitialized represents a Initialized event raised by the StakeCredit contract.
type StakeCreditInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeCredit *StakeCreditFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakeCreditInitializedIterator, error) {

	logs, sub, err := _StakeCredit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakeCreditInitializedIterator{contract: _StakeCredit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeCredit *StakeCreditFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakeCreditInitialized) (event.Subscription, error) {

	logs, sub, err := _StakeCredit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeCreditInitialized)
				if err := _StakeCredit.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeCredit *StakeCreditFilterer) ParseInitialized(log types.Log) (*StakeCreditInitialized, error) {
	event := new(StakeCreditInitialized)
	if err := _StakeCredit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeCreditParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the StakeCredit contract.
type StakeCreditParamChangeIterator struct {
	Event *StakeCreditParamChange // Event containing the contract specifics and raw log

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
func (it *StakeCreditParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeCreditParamChange)
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
		it.Event = new(StakeCreditParamChange)
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
func (it *StakeCreditParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeCreditParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeCreditParamChange represents a ParamChange event raised by the StakeCredit contract.
type StakeCreditParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_StakeCredit *StakeCreditFilterer) FilterParamChange(opts *bind.FilterOpts) (*StakeCreditParamChangeIterator, error) {

	logs, sub, err := _StakeCredit.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &StakeCreditParamChangeIterator{contract: _StakeCredit.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_StakeCredit *StakeCreditFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *StakeCreditParamChange) (event.Subscription, error) {

	logs, sub, err := _StakeCredit.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeCreditParamChange)
				if err := _StakeCredit.contract.UnpackLog(event, "ParamChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_StakeCredit *StakeCreditFilterer) ParseParamChange(log types.Log) (*StakeCreditParamChange, error) {
	event := new(StakeCreditParamChange)
	if err := _StakeCredit.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeCreditRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the StakeCredit contract.
type StakeCreditRewardReceivedIterator struct {
	Event *StakeCreditRewardReceived // Event containing the contract specifics and raw log

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
func (it *StakeCreditRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeCreditRewardReceived)
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
		it.Event = new(StakeCreditRewardReceived)
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
func (it *StakeCreditRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeCreditRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeCreditRewardReceived represents a RewardReceived event raised by the StakeCredit contract.
type StakeCreditRewardReceived struct {
	RewardToAll *big.Int
	Commission  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xfb0e1482d62102ab9594f69d4c6d693749e3e2bf1c21af272f5456b2d5a4f6b5.
//
// Solidity: event RewardReceived(uint256 rewardToAll, uint256 commission)
func (_StakeCredit *StakeCreditFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*StakeCreditRewardReceivedIterator, error) {

	logs, sub, err := _StakeCredit.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &StakeCreditRewardReceivedIterator{contract: _StakeCredit.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xfb0e1482d62102ab9594f69d4c6d693749e3e2bf1c21af272f5456b2d5a4f6b5.
//
// Solidity: event RewardReceived(uint256 rewardToAll, uint256 commission)
func (_StakeCredit *StakeCreditFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *StakeCreditRewardReceived) (event.Subscription, error) {

	logs, sub, err := _StakeCredit.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeCreditRewardReceived)
				if err := _StakeCredit.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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

// ParseRewardReceived is a log parse operation binding the contract event 0xfb0e1482d62102ab9594f69d4c6d693749e3e2bf1c21af272f5456b2d5a4f6b5.
//
// Solidity: event RewardReceived(uint256 rewardToAll, uint256 commission)
func (_StakeCredit *StakeCreditFilterer) ParseRewardReceived(log types.Log) (*StakeCreditRewardReceived, error) {
	event := new(StakeCreditRewardReceived)
	if err := _StakeCredit.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeCreditTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakeCredit contract.
type StakeCreditTransferIterator struct {
	Event *StakeCreditTransfer // Event containing the contract specifics and raw log

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
func (it *StakeCreditTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeCreditTransfer)
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
		it.Event = new(StakeCreditTransfer)
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
func (it *StakeCreditTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeCreditTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeCreditTransfer represents a Transfer event raised by the StakeCredit contract.
type StakeCreditTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakeCredit *StakeCreditFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakeCreditTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakeCredit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakeCreditTransferIterator{contract: _StakeCredit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakeCredit *StakeCreditFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakeCreditTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakeCredit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeCreditTransfer)
				if err := _StakeCredit.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakeCredit *StakeCreditFilterer) ParseTransfer(log types.Log) (*StakeCreditTransfer, error) {
	event := new(StakeCreditTransfer)
	if err := _StakeCredit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
