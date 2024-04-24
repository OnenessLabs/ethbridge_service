// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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

// MesonMetaData contains all meta data concerning the Meson contract.
var MesonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"premiumManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"poolIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"PoolAuthorizedAddrAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"poolIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"PoolAuthorizedAddrRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint48\",\"name\":\"poolTokenIndex\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PoolDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"poolIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"PoolOwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"poolIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint48\",\"name\":\"poolTokenIndex\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PoolWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevPremiumManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPremiumManager\",\"type\":\"address\"}],\"name\":\"PremiumManagerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapBonded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapPosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"SwapUnlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"indexes\",\"type\":\"uint8[]\"}],\"name\":\"addMultipleSupportedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"poolIndex\",\"type\":\"uint40\"}],\"name\":\"bondSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"cancelSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"}],\"name\":\"cancelSwapTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"poolTokenIndex\",\"type\":\"uint48\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"poolTokenIndex\",\"type\":\"uint48\"}],\"name\":\"depositAndRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"directExecuteSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"directRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"}],\"name\":\"directSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"depositToPool\",\"type\":\"bool\"}],\"name\":\"executeSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"}],\"name\":\"getLockedSwap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolOwner\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"until\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"getPostedSwap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getShortCoinType\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"indexes\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexOfToken\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"}],\"name\":\"lockSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"name\":\"ownerOfPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolOfAuthorizedAddr\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"poolTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"},{\"internalType\":\"uint200\",\"name\":\"postingValue\",\"type\":\"uint200\"}],\"name\":\"postSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint200\",\"name\":\"postingValue\",\"type\":\"uint200\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"postSwapFromContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint200\",\"name\":\"postingValue\",\"type\":\"uint200\"}],\"name\":\"postSwapFromInitiator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"yParityAndS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"serviceFeeCollected\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"}],\"name\":\"simpleExecuteSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"tokenForIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPremiumManager\",\"type\":\"address\"}],\"name\":\"transferPremiumManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedSwap\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"poolTokenIndex\",\"type\":\"uint48\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"toPoolIndex\",\"type\":\"uint40\"}],\"name\":\"withdrawServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MesonABI is the input ABI used to generate the binding from.
// Deprecated: Use MesonMetaData.ABI instead.
var MesonABI = MesonMetaData.ABI

// Meson is an auto generated Go binding around an Ethereum contract.
type Meson struct {
	MesonCaller     // Read-only binding to the contract
	MesonTransactor // Write-only binding to the contract
	MesonFilterer   // Log filterer for contract events
}

// MesonCaller is an auto generated read-only Go binding around an Ethereum contract.
type MesonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MesonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MesonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MesonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MesonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MesonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MesonSession struct {
	Contract     *Meson            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MesonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MesonCallerSession struct {
	Contract *MesonCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MesonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MesonTransactorSession struct {
	Contract     *MesonTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MesonRaw is an auto generated low-level Go binding around an Ethereum contract.
type MesonRaw struct {
	Contract *Meson // Generic contract binding to access the raw methods on
}

// MesonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MesonCallerRaw struct {
	Contract *MesonCaller // Generic read-only contract binding to access the raw methods on
}

// MesonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MesonTransactorRaw struct {
	Contract *MesonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeson creates a new instance of Meson, bound to a specific deployed contract.
func NewMeson(address common.Address, backend bind.ContractBackend) (*Meson, error) {
	contract, err := bindMeson(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Meson{MesonCaller: MesonCaller{contract: contract}, MesonTransactor: MesonTransactor{contract: contract}, MesonFilterer: MesonFilterer{contract: contract}}, nil
}

// NewMesonCaller creates a new read-only instance of Meson, bound to a specific deployed contract.
func NewMesonCaller(address common.Address, caller bind.ContractCaller) (*MesonCaller, error) {
	contract, err := bindMeson(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MesonCaller{contract: contract}, nil
}

// NewMesonTransactor creates a new write-only instance of Meson, bound to a specific deployed contract.
func NewMesonTransactor(address common.Address, transactor bind.ContractTransactor) (*MesonTransactor, error) {
	contract, err := bindMeson(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MesonTransactor{contract: contract}, nil
}

// NewMesonFilterer creates a new log filterer instance of Meson, bound to a specific deployed contract.
func NewMesonFilterer(address common.Address, filterer bind.ContractFilterer) (*MesonFilterer, error) {
	contract, err := bindMeson(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MesonFilterer{contract: contract}, nil
}

// bindMeson binds a generic wrapper to an already deployed contract.
func bindMeson(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MesonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meson *MesonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meson.Contract.MesonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meson *MesonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meson.Contract.MesonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meson *MesonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meson.Contract.MesonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meson *MesonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meson.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meson *MesonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meson.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meson *MesonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meson.Contract.contract.Transact(opts, method, params...)
}

// GetLockedSwap is a free data retrieval call binding the contract method 0x60a2da98.
//
// Solidity: function getLockedSwap(uint256 encodedSwap, address initiator) view returns(address poolOwner, uint40 until)
func (_Meson *MesonCaller) GetLockedSwap(opts *bind.CallOpts, encodedSwap *big.Int, initiator common.Address) (struct {
	PoolOwner common.Address
	Until     *big.Int
}, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "getLockedSwap", encodedSwap, initiator)

	outstruct := new(struct {
		PoolOwner common.Address
		Until     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolOwner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Until = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLockedSwap is a free data retrieval call binding the contract method 0x60a2da98.
//
// Solidity: function getLockedSwap(uint256 encodedSwap, address initiator) view returns(address poolOwner, uint40 until)
func (_Meson *MesonSession) GetLockedSwap(encodedSwap *big.Int, initiator common.Address) (struct {
	PoolOwner common.Address
	Until     *big.Int
}, error) {
	return _Meson.Contract.GetLockedSwap(&_Meson.CallOpts, encodedSwap, initiator)
}

// GetLockedSwap is a free data retrieval call binding the contract method 0x60a2da98.
//
// Solidity: function getLockedSwap(uint256 encodedSwap, address initiator) view returns(address poolOwner, uint40 until)
func (_Meson *MesonCallerSession) GetLockedSwap(encodedSwap *big.Int, initiator common.Address) (struct {
	PoolOwner common.Address
	Until     *big.Int
}, error) {
	return _Meson.Contract.GetLockedSwap(&_Meson.CallOpts, encodedSwap, initiator)
}

// GetPostedSwap is a free data retrieval call binding the contract method 0x1e2a6075.
//
// Solidity: function getPostedSwap(uint256 encodedSwap) view returns(address initiator, address poolOwner, bool exist)
func (_Meson *MesonCaller) GetPostedSwap(opts *bind.CallOpts, encodedSwap *big.Int) (struct {
	Initiator common.Address
	PoolOwner common.Address
	Exist     bool
}, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "getPostedSwap", encodedSwap)

	outstruct := new(struct {
		Initiator common.Address
		PoolOwner common.Address
		Exist     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initiator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PoolOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Exist = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetPostedSwap is a free data retrieval call binding the contract method 0x1e2a6075.
//
// Solidity: function getPostedSwap(uint256 encodedSwap) view returns(address initiator, address poolOwner, bool exist)
func (_Meson *MesonSession) GetPostedSwap(encodedSwap *big.Int) (struct {
	Initiator common.Address
	PoolOwner common.Address
	Exist     bool
}, error) {
	return _Meson.Contract.GetPostedSwap(&_Meson.CallOpts, encodedSwap)
}

// GetPostedSwap is a free data retrieval call binding the contract method 0x1e2a6075.
//
// Solidity: function getPostedSwap(uint256 encodedSwap) view returns(address initiator, address poolOwner, bool exist)
func (_Meson *MesonCallerSession) GetPostedSwap(encodedSwap *big.Int) (struct {
	Initiator common.Address
	PoolOwner common.Address
	Exist     bool
}, error) {
	return _Meson.Contract.GetPostedSwap(&_Meson.CallOpts, encodedSwap)
}

// GetShortCoinType is a free data retrieval call binding the contract method 0xeba7fb77.
//
// Solidity: function getShortCoinType() pure returns(bytes2)
func (_Meson *MesonCaller) GetShortCoinType(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "getShortCoinType")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetShortCoinType is a free data retrieval call binding the contract method 0xeba7fb77.
//
// Solidity: function getShortCoinType() pure returns(bytes2)
func (_Meson *MesonSession) GetShortCoinType() ([2]byte, error) {
	return _Meson.Contract.GetShortCoinType(&_Meson.CallOpts)
}

// GetShortCoinType is a free data retrieval call binding the contract method 0xeba7fb77.
//
// Solidity: function getShortCoinType() pure returns(bytes2)
func (_Meson *MesonCallerSession) GetShortCoinType() ([2]byte, error) {
	return _Meson.Contract.GetShortCoinType(&_Meson.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens, uint8[] indexes)
func (_Meson *MesonCaller) GetSupportedTokens(opts *bind.CallOpts) (struct {
	Tokens  []common.Address
	Indexes []uint8
}, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "getSupportedTokens")

	outstruct := new(struct {
		Tokens  []common.Address
		Indexes []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Indexes = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens, uint8[] indexes)
func (_Meson *MesonSession) GetSupportedTokens() (struct {
	Tokens  []common.Address
	Indexes []uint8
}, error) {
	return _Meson.Contract.GetSupportedTokens(&_Meson.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens, uint8[] indexes)
func (_Meson *MesonCallerSession) GetSupportedTokens() (struct {
	Tokens  []common.Address
	Indexes []uint8
}, error) {
	return _Meson.Contract.GetSupportedTokens(&_Meson.CallOpts)
}

// IndexOfToken is a free data retrieval call binding the contract method 0x2335093c.
//
// Solidity: function indexOfToken(address ) view returns(uint8)
func (_Meson *MesonCaller) IndexOfToken(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "indexOfToken", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// IndexOfToken is a free data retrieval call binding the contract method 0x2335093c.
//
// Solidity: function indexOfToken(address ) view returns(uint8)
func (_Meson *MesonSession) IndexOfToken(arg0 common.Address) (uint8, error) {
	return _Meson.Contract.IndexOfToken(&_Meson.CallOpts, arg0)
}

// IndexOfToken is a free data retrieval call binding the contract method 0x2335093c.
//
// Solidity: function indexOfToken(address ) view returns(uint8)
func (_Meson *MesonCallerSession) IndexOfToken(arg0 common.Address) (uint8, error) {
	return _Meson.Contract.IndexOfToken(&_Meson.CallOpts, arg0)
}

// OwnerOfPool is a free data retrieval call binding the contract method 0x89a734c0.
//
// Solidity: function ownerOfPool(uint40 ) view returns(address)
func (_Meson *MesonCaller) OwnerOfPool(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "ownerOfPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOfPool is a free data retrieval call binding the contract method 0x89a734c0.
//
// Solidity: function ownerOfPool(uint40 ) view returns(address)
func (_Meson *MesonSession) OwnerOfPool(arg0 *big.Int) (common.Address, error) {
	return _Meson.Contract.OwnerOfPool(&_Meson.CallOpts, arg0)
}

// OwnerOfPool is a free data retrieval call binding the contract method 0x89a734c0.
//
// Solidity: function ownerOfPool(uint40 ) view returns(address)
func (_Meson *MesonCallerSession) OwnerOfPool(arg0 *big.Int) (common.Address, error) {
	return _Meson.Contract.OwnerOfPool(&_Meson.CallOpts, arg0)
}

// PoolOfAuthorizedAddr is a free data retrieval call binding the contract method 0x7fe0282b.
//
// Solidity: function poolOfAuthorizedAddr(address ) view returns(uint40)
func (_Meson *MesonCaller) PoolOfAuthorizedAddr(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "poolOfAuthorizedAddr", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolOfAuthorizedAddr is a free data retrieval call binding the contract method 0x7fe0282b.
//
// Solidity: function poolOfAuthorizedAddr(address ) view returns(uint40)
func (_Meson *MesonSession) PoolOfAuthorizedAddr(arg0 common.Address) (*big.Int, error) {
	return _Meson.Contract.PoolOfAuthorizedAddr(&_Meson.CallOpts, arg0)
}

// PoolOfAuthorizedAddr is a free data retrieval call binding the contract method 0x7fe0282b.
//
// Solidity: function poolOfAuthorizedAddr(address ) view returns(uint40)
func (_Meson *MesonCallerSession) PoolOfAuthorizedAddr(arg0 common.Address) (*big.Int, error) {
	return _Meson.Contract.PoolOfAuthorizedAddr(&_Meson.CallOpts, arg0)
}

// PoolTokenBalance is a free data retrieval call binding the contract method 0xd3e95ea4.
//
// Solidity: function poolTokenBalance(address token, address addr) view returns(uint256)
func (_Meson *MesonCaller) PoolTokenBalance(opts *bind.CallOpts, token common.Address, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "poolTokenBalance", token, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolTokenBalance is a free data retrieval call binding the contract method 0xd3e95ea4.
//
// Solidity: function poolTokenBalance(address token, address addr) view returns(uint256)
func (_Meson *MesonSession) PoolTokenBalance(token common.Address, addr common.Address) (*big.Int, error) {
	return _Meson.Contract.PoolTokenBalance(&_Meson.CallOpts, token, addr)
}

// PoolTokenBalance is a free data retrieval call binding the contract method 0xd3e95ea4.
//
// Solidity: function poolTokenBalance(address token, address addr) view returns(uint256)
func (_Meson *MesonCallerSession) PoolTokenBalance(token common.Address, addr common.Address) (*big.Int, error) {
	return _Meson.Contract.PoolTokenBalance(&_Meson.CallOpts, token, addr)
}

// ServiceFeeCollected is a free data retrieval call binding the contract method 0x8b0a7765.
//
// Solidity: function serviceFeeCollected(uint8 tokenIndex) view returns(uint256)
func (_Meson *MesonCaller) ServiceFeeCollected(opts *bind.CallOpts, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "serviceFeeCollected", tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeCollected is a free data retrieval call binding the contract method 0x8b0a7765.
//
// Solidity: function serviceFeeCollected(uint8 tokenIndex) view returns(uint256)
func (_Meson *MesonSession) ServiceFeeCollected(tokenIndex uint8) (*big.Int, error) {
	return _Meson.Contract.ServiceFeeCollected(&_Meson.CallOpts, tokenIndex)
}

// ServiceFeeCollected is a free data retrieval call binding the contract method 0x8b0a7765.
//
// Solidity: function serviceFeeCollected(uint8 tokenIndex) view returns(uint256)
func (_Meson *MesonCallerSession) ServiceFeeCollected(tokenIndex uint8) (*big.Int, error) {
	return _Meson.Contract.ServiceFeeCollected(&_Meson.CallOpts, tokenIndex)
}

// TokenForIndex is a free data retrieval call binding the contract method 0xff378719.
//
// Solidity: function tokenForIndex(uint8 ) view returns(address)
func (_Meson *MesonCaller) TokenForIndex(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _Meson.contract.Call(opts, &out, "tokenForIndex", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenForIndex is a free data retrieval call binding the contract method 0xff378719.
//
// Solidity: function tokenForIndex(uint8 ) view returns(address)
func (_Meson *MesonSession) TokenForIndex(arg0 uint8) (common.Address, error) {
	return _Meson.Contract.TokenForIndex(&_Meson.CallOpts, arg0)
}

// TokenForIndex is a free data retrieval call binding the contract method 0xff378719.
//
// Solidity: function tokenForIndex(uint8 ) view returns(address)
func (_Meson *MesonCallerSession) TokenForIndex(arg0 uint8) (common.Address, error) {
	return _Meson.Contract.TokenForIndex(&_Meson.CallOpts, arg0)
}

// AddAuthorizedAddr is a paid mutator transaction binding the contract method 0xff22f272.
//
// Solidity: function addAuthorizedAddr(address addr) returns()
func (_Meson *MesonTransactor) AddAuthorizedAddr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "addAuthorizedAddr", addr)
}

// AddAuthorizedAddr is a paid mutator transaction binding the contract method 0xff22f272.
//
// Solidity: function addAuthorizedAddr(address addr) returns()
func (_Meson *MesonSession) AddAuthorizedAddr(addr common.Address) (*types.Transaction, error) {
	return _Meson.Contract.AddAuthorizedAddr(&_Meson.TransactOpts, addr)
}

// AddAuthorizedAddr is a paid mutator transaction binding the contract method 0xff22f272.
//
// Solidity: function addAuthorizedAddr(address addr) returns()
func (_Meson *MesonTransactorSession) AddAuthorizedAddr(addr common.Address) (*types.Transaction, error) {
	return _Meson.Contract.AddAuthorizedAddr(&_Meson.TransactOpts, addr)
}

// AddMultipleSupportedTokens is a paid mutator transaction binding the contract method 0xc11d9ecb.
//
// Solidity: function addMultipleSupportedTokens(address[] tokens, uint8[] indexes) returns()
func (_Meson *MesonTransactor) AddMultipleSupportedTokens(opts *bind.TransactOpts, tokens []common.Address, indexes []uint8) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "addMultipleSupportedTokens", tokens, indexes)
}

// AddMultipleSupportedTokens is a paid mutator transaction binding the contract method 0xc11d9ecb.
//
// Solidity: function addMultipleSupportedTokens(address[] tokens, uint8[] indexes) returns()
func (_Meson *MesonSession) AddMultipleSupportedTokens(tokens []common.Address, indexes []uint8) (*types.Transaction, error) {
	return _Meson.Contract.AddMultipleSupportedTokens(&_Meson.TransactOpts, tokens, indexes)
}

// AddMultipleSupportedTokens is a paid mutator transaction binding the contract method 0xc11d9ecb.
//
// Solidity: function addMultipleSupportedTokens(address[] tokens, uint8[] indexes) returns()
func (_Meson *MesonTransactorSession) AddMultipleSupportedTokens(tokens []common.Address, indexes []uint8) (*types.Transaction, error) {
	return _Meson.Contract.AddMultipleSupportedTokens(&_Meson.TransactOpts, tokens, indexes)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0xd4f82322.
//
// Solidity: function addSupportToken(address token, uint8 index) returns()
func (_Meson *MesonTransactor) AddSupportToken(opts *bind.TransactOpts, token common.Address, index uint8) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "addSupportToken", token, index)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0xd4f82322.
//
// Solidity: function addSupportToken(address token, uint8 index) returns()
func (_Meson *MesonSession) AddSupportToken(token common.Address, index uint8) (*types.Transaction, error) {
	return _Meson.Contract.AddSupportToken(&_Meson.TransactOpts, token, index)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0xd4f82322.
//
// Solidity: function addSupportToken(address token, uint8 index) returns()
func (_Meson *MesonTransactorSession) AddSupportToken(token common.Address, index uint8) (*types.Transaction, error) {
	return _Meson.Contract.AddSupportToken(&_Meson.TransactOpts, token, index)
}

// BondSwap is a paid mutator transaction binding the contract method 0x35eff30f.
//
// Solidity: function bondSwap(uint256 encodedSwap, uint40 poolIndex) returns()
func (_Meson *MesonTransactor) BondSwap(opts *bind.TransactOpts, encodedSwap *big.Int, poolIndex *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "bondSwap", encodedSwap, poolIndex)
}

// BondSwap is a paid mutator transaction binding the contract method 0x35eff30f.
//
// Solidity: function bondSwap(uint256 encodedSwap, uint40 poolIndex) returns()
func (_Meson *MesonSession) BondSwap(encodedSwap *big.Int, poolIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.BondSwap(&_Meson.TransactOpts, encodedSwap, poolIndex)
}

// BondSwap is a paid mutator transaction binding the contract method 0x35eff30f.
//
// Solidity: function bondSwap(uint256 encodedSwap, uint40 poolIndex) returns()
func (_Meson *MesonTransactorSession) BondSwap(encodedSwap *big.Int, poolIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.BondSwap(&_Meson.TransactOpts, encodedSwap, poolIndex)
}

// CancelSwap is a paid mutator transaction binding the contract method 0x54d6a2b7.
//
// Solidity: function cancelSwap(uint256 encodedSwap) returns()
func (_Meson *MesonTransactor) CancelSwap(opts *bind.TransactOpts, encodedSwap *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "cancelSwap", encodedSwap)
}

// CancelSwap is a paid mutator transaction binding the contract method 0x54d6a2b7.
//
// Solidity: function cancelSwap(uint256 encodedSwap) returns()
func (_Meson *MesonSession) CancelSwap(encodedSwap *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.CancelSwap(&_Meson.TransactOpts, encodedSwap)
}

// CancelSwap is a paid mutator transaction binding the contract method 0x54d6a2b7.
//
// Solidity: function cancelSwap(uint256 encodedSwap) returns()
func (_Meson *MesonTransactorSession) CancelSwap(encodedSwap *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.CancelSwap(&_Meson.TransactOpts, encodedSwap)
}

// CancelSwapTo is a paid mutator transaction binding the contract method 0x9fa10279.
//
// Solidity: function cancelSwapTo(uint256 encodedSwap, address recipient, bytes32 r, bytes32 yParityAndS) returns()
func (_Meson *MesonTransactor) CancelSwapTo(opts *bind.TransactOpts, encodedSwap *big.Int, recipient common.Address, r [32]byte, yParityAndS [32]byte) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "cancelSwapTo", encodedSwap, recipient, r, yParityAndS)
}

// CancelSwapTo is a paid mutator transaction binding the contract method 0x9fa10279.
//
// Solidity: function cancelSwapTo(uint256 encodedSwap, address recipient, bytes32 r, bytes32 yParityAndS) returns()
func (_Meson *MesonSession) CancelSwapTo(encodedSwap *big.Int, recipient common.Address, r [32]byte, yParityAndS [32]byte) (*types.Transaction, error) {
	return _Meson.Contract.CancelSwapTo(&_Meson.TransactOpts, encodedSwap, recipient, r, yParityAndS)
}

// CancelSwapTo is a paid mutator transaction binding the contract method 0x9fa10279.
//
// Solidity: function cancelSwapTo(uint256 encodedSwap, address recipient, bytes32 r, bytes32 yParityAndS) returns()
func (_Meson *MesonTransactorSession) CancelSwapTo(encodedSwap *big.Int, recipient common.Address, r [32]byte, yParityAndS [32]byte) (*types.Transaction, error) {
	return _Meson.Contract.CancelSwapTo(&_Meson.TransactOpts, encodedSwap, recipient, r, yParityAndS)
}

// Deposit is a paid mutator transaction binding the contract method 0x37b90a4f.
//
// Solidity: function deposit(uint256 amount, uint48 poolTokenIndex) payable returns()
func (_Meson *MesonTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "deposit", amount, poolTokenIndex)
}

// Deposit is a paid mutator transaction binding the contract method 0x37b90a4f.
//
// Solidity: function deposit(uint256 amount, uint48 poolTokenIndex) payable returns()
func (_Meson *MesonSession) Deposit(amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.Deposit(&_Meson.TransactOpts, amount, poolTokenIndex)
}

// Deposit is a paid mutator transaction binding the contract method 0x37b90a4f.
//
// Solidity: function deposit(uint256 amount, uint48 poolTokenIndex) payable returns()
func (_Meson *MesonTransactorSession) Deposit(amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.Deposit(&_Meson.TransactOpts, amount, poolTokenIndex)
}

// DepositAndRegister is a paid mutator transaction binding the contract method 0x8f487dc9.
//
// Solidity: function depositAndRegister(uint256 amount, uint48 poolTokenIndex) payable returns()
func (_Meson *MesonTransactor) DepositAndRegister(opts *bind.TransactOpts, amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "depositAndRegister", amount, poolTokenIndex)
}

// DepositAndRegister is a paid mutator transaction binding the contract method 0x8f487dc9.
//
// Solidity: function depositAndRegister(uint256 amount, uint48 poolTokenIndex) payable returns()
func (_Meson *MesonSession) DepositAndRegister(amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.DepositAndRegister(&_Meson.TransactOpts, amount, poolTokenIndex)
}

// DepositAndRegister is a paid mutator transaction binding the contract method 0x8f487dc9.
//
// Solidity: function depositAndRegister(uint256 amount, uint48 poolTokenIndex) payable returns()
func (_Meson *MesonTransactorSession) DepositAndRegister(amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.DepositAndRegister(&_Meson.TransactOpts, amount, poolTokenIndex)
}

// DirectExecuteSwap is a paid mutator transaction binding the contract method 0xc8173c44.
//
// Solidity: function directExecuteSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) payable returns()
func (_Meson *MesonTransactor) DirectExecuteSwap(opts *bind.TransactOpts, encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "directExecuteSwap", encodedSwap, r, yParityAndS, initiator, recipient)
}

// DirectExecuteSwap is a paid mutator transaction binding the contract method 0xc8173c44.
//
// Solidity: function directExecuteSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) payable returns()
func (_Meson *MesonSession) DirectExecuteSwap(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.Contract.DirectExecuteSwap(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator, recipient)
}

// DirectExecuteSwap is a paid mutator transaction binding the contract method 0xc8173c44.
//
// Solidity: function directExecuteSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) payable returns()
func (_Meson *MesonTransactorSession) DirectExecuteSwap(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.Contract.DirectExecuteSwap(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator, recipient)
}

// DirectRelease is a paid mutator transaction binding the contract method 0xab115fd8.
//
// Solidity: function directRelease(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) returns()
func (_Meson *MesonTransactor) DirectRelease(opts *bind.TransactOpts, encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "directRelease", encodedSwap, r, yParityAndS, initiator, recipient)
}

// DirectRelease is a paid mutator transaction binding the contract method 0xab115fd8.
//
// Solidity: function directRelease(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) returns()
func (_Meson *MesonSession) DirectRelease(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.Contract.DirectRelease(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator, recipient)
}

// DirectRelease is a paid mutator transaction binding the contract method 0xab115fd8.
//
// Solidity: function directRelease(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) returns()
func (_Meson *MesonTransactorSession) DirectRelease(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.Contract.DirectRelease(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator, recipient)
}

// DirectSwap is a paid mutator transaction binding the contract method 0x741c8e2d.
//
// Solidity: function directSwap(uint256 encodedSwap, address recipient, bytes32 r, bytes32 yParityAndS) returns()
func (_Meson *MesonTransactor) DirectSwap(opts *bind.TransactOpts, encodedSwap *big.Int, recipient common.Address, r [32]byte, yParityAndS [32]byte) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "directSwap", encodedSwap, recipient, r, yParityAndS)
}

// DirectSwap is a paid mutator transaction binding the contract method 0x741c8e2d.
//
// Solidity: function directSwap(uint256 encodedSwap, address recipient, bytes32 r, bytes32 yParityAndS) returns()
func (_Meson *MesonSession) DirectSwap(encodedSwap *big.Int, recipient common.Address, r [32]byte, yParityAndS [32]byte) (*types.Transaction, error) {
	return _Meson.Contract.DirectSwap(&_Meson.TransactOpts, encodedSwap, recipient, r, yParityAndS)
}

// DirectSwap is a paid mutator transaction binding the contract method 0x741c8e2d.
//
// Solidity: function directSwap(uint256 encodedSwap, address recipient, bytes32 r, bytes32 yParityAndS) returns()
func (_Meson *MesonTransactorSession) DirectSwap(encodedSwap *big.Int, recipient common.Address, r [32]byte, yParityAndS [32]byte) (*types.Transaction, error) {
	return _Meson.Contract.DirectSwap(&_Meson.TransactOpts, encodedSwap, recipient, r, yParityAndS)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x827c87cc.
//
// Solidity: function executeSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address recipient, bool depositToPool) returns()
func (_Meson *MesonTransactor) ExecuteSwap(opts *bind.TransactOpts, encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, recipient common.Address, depositToPool bool) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "executeSwap", encodedSwap, r, yParityAndS, recipient, depositToPool)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x827c87cc.
//
// Solidity: function executeSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address recipient, bool depositToPool) returns()
func (_Meson *MesonSession) ExecuteSwap(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, recipient common.Address, depositToPool bool) (*types.Transaction, error) {
	return _Meson.Contract.ExecuteSwap(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, recipient, depositToPool)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x827c87cc.
//
// Solidity: function executeSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address recipient, bool depositToPool) returns()
func (_Meson *MesonTransactorSession) ExecuteSwap(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, recipient common.Address, depositToPool bool) (*types.Transaction, error) {
	return _Meson.Contract.ExecuteSwap(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, recipient, depositToPool)
}

// Lock is a paid mutator transaction binding the contract method 0x515147ab.
//
// Solidity: function lock(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator) returns()
func (_Meson *MesonTransactor) Lock(opts *bind.TransactOpts, encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "lock", encodedSwap, r, yParityAndS, initiator)
}

// Lock is a paid mutator transaction binding the contract method 0x515147ab.
//
// Solidity: function lock(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator) returns()
func (_Meson *MesonSession) Lock(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address) (*types.Transaction, error) {
	return _Meson.Contract.Lock(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator)
}

// Lock is a paid mutator transaction binding the contract method 0x515147ab.
//
// Solidity: function lock(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator) returns()
func (_Meson *MesonTransactorSession) Lock(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address) (*types.Transaction, error) {
	return _Meson.Contract.Lock(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator)
}

// LockSwap is a paid mutator transaction binding the contract method 0x58d9b4e1.
//
// Solidity: function lockSwap(uint256 encodedSwap, address initiator) returns()
func (_Meson *MesonTransactor) LockSwap(opts *bind.TransactOpts, encodedSwap *big.Int, initiator common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "lockSwap", encodedSwap, initiator)
}

// LockSwap is a paid mutator transaction binding the contract method 0x58d9b4e1.
//
// Solidity: function lockSwap(uint256 encodedSwap, address initiator) returns()
func (_Meson *MesonSession) LockSwap(encodedSwap *big.Int, initiator common.Address) (*types.Transaction, error) {
	return _Meson.Contract.LockSwap(&_Meson.TransactOpts, encodedSwap, initiator)
}

// LockSwap is a paid mutator transaction binding the contract method 0x58d9b4e1.
//
// Solidity: function lockSwap(uint256 encodedSwap, address initiator) returns()
func (_Meson *MesonTransactorSession) LockSwap(encodedSwap *big.Int, initiator common.Address) (*types.Transaction, error) {
	return _Meson.Contract.LockSwap(&_Meson.TransactOpts, encodedSwap, initiator)
}

// PostSwap is a paid mutator transaction binding the contract method 0xc5d7ca00.
//
// Solidity: function postSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, uint200 postingValue) returns()
func (_Meson *MesonTransactor) PostSwap(opts *bind.TransactOpts, encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, postingValue *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "postSwap", encodedSwap, r, yParityAndS, postingValue)
}

// PostSwap is a paid mutator transaction binding the contract method 0xc5d7ca00.
//
// Solidity: function postSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, uint200 postingValue) returns()
func (_Meson *MesonSession) PostSwap(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, postingValue *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.PostSwap(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, postingValue)
}

// PostSwap is a paid mutator transaction binding the contract method 0xc5d7ca00.
//
// Solidity: function postSwap(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, uint200 postingValue) returns()
func (_Meson *MesonTransactorSession) PostSwap(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, postingValue *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.PostSwap(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, postingValue)
}

// PostSwapFromContract is a paid mutator transaction binding the contract method 0x60b068be.
//
// Solidity: function postSwapFromContract(uint256 encodedSwap, uint200 postingValue, address contractAddress) payable returns()
func (_Meson *MesonTransactor) PostSwapFromContract(opts *bind.TransactOpts, encodedSwap *big.Int, postingValue *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "postSwapFromContract", encodedSwap, postingValue, contractAddress)
}

// PostSwapFromContract is a paid mutator transaction binding the contract method 0x60b068be.
//
// Solidity: function postSwapFromContract(uint256 encodedSwap, uint200 postingValue, address contractAddress) payable returns()
func (_Meson *MesonSession) PostSwapFromContract(encodedSwap *big.Int, postingValue *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _Meson.Contract.PostSwapFromContract(&_Meson.TransactOpts, encodedSwap, postingValue, contractAddress)
}

// PostSwapFromContract is a paid mutator transaction binding the contract method 0x60b068be.
//
// Solidity: function postSwapFromContract(uint256 encodedSwap, uint200 postingValue, address contractAddress) payable returns()
func (_Meson *MesonTransactorSession) PostSwapFromContract(encodedSwap *big.Int, postingValue *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _Meson.Contract.PostSwapFromContract(&_Meson.TransactOpts, encodedSwap, postingValue, contractAddress)
}

// PostSwapFromInitiator is a paid mutator transaction binding the contract method 0xdecf2a48.
//
// Solidity: function postSwapFromInitiator(uint256 encodedSwap, uint200 postingValue) payable returns()
func (_Meson *MesonTransactor) PostSwapFromInitiator(opts *bind.TransactOpts, encodedSwap *big.Int, postingValue *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "postSwapFromInitiator", encodedSwap, postingValue)
}

// PostSwapFromInitiator is a paid mutator transaction binding the contract method 0xdecf2a48.
//
// Solidity: function postSwapFromInitiator(uint256 encodedSwap, uint200 postingValue) payable returns()
func (_Meson *MesonSession) PostSwapFromInitiator(encodedSwap *big.Int, postingValue *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.PostSwapFromInitiator(&_Meson.TransactOpts, encodedSwap, postingValue)
}

// PostSwapFromInitiator is a paid mutator transaction binding the contract method 0xdecf2a48.
//
// Solidity: function postSwapFromInitiator(uint256 encodedSwap, uint200 postingValue) payable returns()
func (_Meson *MesonTransactorSession) PostSwapFromInitiator(encodedSwap *big.Int, postingValue *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.PostSwapFromInitiator(&_Meson.TransactOpts, encodedSwap, postingValue)
}

// Release is a paid mutator transaction binding the contract method 0xa5c9c66c.
//
// Solidity: function release(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) returns()
func (_Meson *MesonTransactor) Release(opts *bind.TransactOpts, encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "release", encodedSwap, r, yParityAndS, initiator, recipient)
}

// Release is a paid mutator transaction binding the contract method 0xa5c9c66c.
//
// Solidity: function release(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) returns()
func (_Meson *MesonSession) Release(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.Contract.Release(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator, recipient)
}

// Release is a paid mutator transaction binding the contract method 0xa5c9c66c.
//
// Solidity: function release(uint256 encodedSwap, bytes32 r, bytes32 yParityAndS, address initiator, address recipient) returns()
func (_Meson *MesonTransactorSession) Release(encodedSwap *big.Int, r [32]byte, yParityAndS [32]byte, initiator common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Meson.Contract.Release(&_Meson.TransactOpts, encodedSwap, r, yParityAndS, initiator, recipient)
}

// RemoveAuthorizedAddr is a paid mutator transaction binding the contract method 0x051119f5.
//
// Solidity: function removeAuthorizedAddr(address addr) returns()
func (_Meson *MesonTransactor) RemoveAuthorizedAddr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "removeAuthorizedAddr", addr)
}

// RemoveAuthorizedAddr is a paid mutator transaction binding the contract method 0x051119f5.
//
// Solidity: function removeAuthorizedAddr(address addr) returns()
func (_Meson *MesonSession) RemoveAuthorizedAddr(addr common.Address) (*types.Transaction, error) {
	return _Meson.Contract.RemoveAuthorizedAddr(&_Meson.TransactOpts, addr)
}

// RemoveAuthorizedAddr is a paid mutator transaction binding the contract method 0x051119f5.
//
// Solidity: function removeAuthorizedAddr(address addr) returns()
func (_Meson *MesonTransactorSession) RemoveAuthorizedAddr(addr common.Address) (*types.Transaction, error) {
	return _Meson.Contract.RemoveAuthorizedAddr(&_Meson.TransactOpts, addr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xcb4f999b.
//
// Solidity: function removeSupportToken(uint8 index) returns()
func (_Meson *MesonTransactor) RemoveSupportToken(opts *bind.TransactOpts, index uint8) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "removeSupportToken", index)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xcb4f999b.
//
// Solidity: function removeSupportToken(uint8 index) returns()
func (_Meson *MesonSession) RemoveSupportToken(index uint8) (*types.Transaction, error) {
	return _Meson.Contract.RemoveSupportToken(&_Meson.TransactOpts, index)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xcb4f999b.
//
// Solidity: function removeSupportToken(uint8 index) returns()
func (_Meson *MesonTransactorSession) RemoveSupportToken(index uint8) (*types.Transaction, error) {
	return _Meson.Contract.RemoveSupportToken(&_Meson.TransactOpts, index)
}

// SimpleExecuteSwap is a paid mutator transaction binding the contract method 0x264849e7.
//
// Solidity: function simpleExecuteSwap(uint256 encodedSwap) payable returns()
func (_Meson *MesonTransactor) SimpleExecuteSwap(opts *bind.TransactOpts, encodedSwap *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "simpleExecuteSwap", encodedSwap)
}

// SimpleExecuteSwap is a paid mutator transaction binding the contract method 0x264849e7.
//
// Solidity: function simpleExecuteSwap(uint256 encodedSwap) payable returns()
func (_Meson *MesonSession) SimpleExecuteSwap(encodedSwap *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.SimpleExecuteSwap(&_Meson.TransactOpts, encodedSwap)
}

// SimpleExecuteSwap is a paid mutator transaction binding the contract method 0x264849e7.
//
// Solidity: function simpleExecuteSwap(uint256 encodedSwap) payable returns()
func (_Meson *MesonTransactorSession) SimpleExecuteSwap(encodedSwap *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.SimpleExecuteSwap(&_Meson.TransactOpts, encodedSwap)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meson *MesonTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meson *MesonSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Meson.Contract.TransferOwnership(&_Meson.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meson *MesonTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Meson.Contract.TransferOwnership(&_Meson.TransactOpts, newOwner)
}

// TransferPoolOwner is a paid mutator transaction binding the contract method 0x30f00f3a.
//
// Solidity: function transferPoolOwner(address addr) returns()
func (_Meson *MesonTransactor) TransferPoolOwner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "transferPoolOwner", addr)
}

// TransferPoolOwner is a paid mutator transaction binding the contract method 0x30f00f3a.
//
// Solidity: function transferPoolOwner(address addr) returns()
func (_Meson *MesonSession) TransferPoolOwner(addr common.Address) (*types.Transaction, error) {
	return _Meson.Contract.TransferPoolOwner(&_Meson.TransactOpts, addr)
}

// TransferPoolOwner is a paid mutator transaction binding the contract method 0x30f00f3a.
//
// Solidity: function transferPoolOwner(address addr) returns()
func (_Meson *MesonTransactorSession) TransferPoolOwner(addr common.Address) (*types.Transaction, error) {
	return _Meson.Contract.TransferPoolOwner(&_Meson.TransactOpts, addr)
}

// TransferPremiumManager is a paid mutator transaction binding the contract method 0xb805f321.
//
// Solidity: function transferPremiumManager(address newPremiumManager) returns()
func (_Meson *MesonTransactor) TransferPremiumManager(opts *bind.TransactOpts, newPremiumManager common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "transferPremiumManager", newPremiumManager)
}

// TransferPremiumManager is a paid mutator transaction binding the contract method 0xb805f321.
//
// Solidity: function transferPremiumManager(address newPremiumManager) returns()
func (_Meson *MesonSession) TransferPremiumManager(newPremiumManager common.Address) (*types.Transaction, error) {
	return _Meson.Contract.TransferPremiumManager(&_Meson.TransactOpts, newPremiumManager)
}

// TransferPremiumManager is a paid mutator transaction binding the contract method 0xb805f321.
//
// Solidity: function transferPremiumManager(address newPremiumManager) returns()
func (_Meson *MesonTransactorSession) TransferPremiumManager(newPremiumManager common.Address) (*types.Transaction, error) {
	return _Meson.Contract.TransferPremiumManager(&_Meson.TransactOpts, newPremiumManager)
}

// Unlock is a paid mutator transaction binding the contract method 0xf1d2ec1d.
//
// Solidity: function unlock(uint256 encodedSwap, address initiator) returns()
func (_Meson *MesonTransactor) Unlock(opts *bind.TransactOpts, encodedSwap *big.Int, initiator common.Address) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "unlock", encodedSwap, initiator)
}

// Unlock is a paid mutator transaction binding the contract method 0xf1d2ec1d.
//
// Solidity: function unlock(uint256 encodedSwap, address initiator) returns()
func (_Meson *MesonSession) Unlock(encodedSwap *big.Int, initiator common.Address) (*types.Transaction, error) {
	return _Meson.Contract.Unlock(&_Meson.TransactOpts, encodedSwap, initiator)
}

// Unlock is a paid mutator transaction binding the contract method 0xf1d2ec1d.
//
// Solidity: function unlock(uint256 encodedSwap, address initiator) returns()
func (_Meson *MesonTransactorSession) Unlock(encodedSwap *big.Int, initiator common.Address) (*types.Transaction, error) {
	return _Meson.Contract.Unlock(&_Meson.TransactOpts, encodedSwap, initiator)
}

// Withdraw is a paid mutator transaction binding the contract method 0xce7f79b9.
//
// Solidity: function withdraw(uint256 amount, uint48 poolTokenIndex) returns()
func (_Meson *MesonTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "withdraw", amount, poolTokenIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0xce7f79b9.
//
// Solidity: function withdraw(uint256 amount, uint48 poolTokenIndex) returns()
func (_Meson *MesonSession) Withdraw(amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.Withdraw(&_Meson.TransactOpts, amount, poolTokenIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0xce7f79b9.
//
// Solidity: function withdraw(uint256 amount, uint48 poolTokenIndex) returns()
func (_Meson *MesonTransactorSession) Withdraw(amount *big.Int, poolTokenIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.Withdraw(&_Meson.TransactOpts, amount, poolTokenIndex)
}

// WithdrawServiceFee is a paid mutator transaction binding the contract method 0x7234cd95.
//
// Solidity: function withdrawServiceFee(uint8 tokenIndex, uint256 amount, uint40 toPoolIndex) returns()
func (_Meson *MesonTransactor) WithdrawServiceFee(opts *bind.TransactOpts, tokenIndex uint8, amount *big.Int, toPoolIndex *big.Int) (*types.Transaction, error) {
	return _Meson.contract.Transact(opts, "withdrawServiceFee", tokenIndex, amount, toPoolIndex)
}

// WithdrawServiceFee is a paid mutator transaction binding the contract method 0x7234cd95.
//
// Solidity: function withdrawServiceFee(uint8 tokenIndex, uint256 amount, uint40 toPoolIndex) returns()
func (_Meson *MesonSession) WithdrawServiceFee(tokenIndex uint8, amount *big.Int, toPoolIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.WithdrawServiceFee(&_Meson.TransactOpts, tokenIndex, amount, toPoolIndex)
}

// WithdrawServiceFee is a paid mutator transaction binding the contract method 0x7234cd95.
//
// Solidity: function withdrawServiceFee(uint8 tokenIndex, uint256 amount, uint40 toPoolIndex) returns()
func (_Meson *MesonTransactorSession) WithdrawServiceFee(tokenIndex uint8, amount *big.Int, toPoolIndex *big.Int) (*types.Transaction, error) {
	return _Meson.Contract.WithdrawServiceFee(&_Meson.TransactOpts, tokenIndex, amount, toPoolIndex)
}

// MesonOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Meson contract.
type MesonOwnerTransferredIterator struct {
	Event *MesonOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *MesonOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonOwnerTransferred)
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
		it.Event = new(MesonOwnerTransferred)
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
func (it *MesonOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonOwnerTransferred represents a OwnerTransferred event raised by the Meson contract.
type MesonOwnerTransferred struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: event OwnerTransferred(address indexed prevOwner, address indexed newOwner)
func (_Meson *MesonFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, prevOwner []common.Address, newOwner []common.Address) (*MesonOwnerTransferredIterator, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "OwnerTransferred", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MesonOwnerTransferredIterator{contract: _Meson.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: event OwnerTransferred(address indexed prevOwner, address indexed newOwner)
func (_Meson *MesonFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *MesonOwnerTransferred, prevOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "OwnerTransferred", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonOwnerTransferred)
				if err := _Meson.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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

// ParseOwnerTransferred is a log parse operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: event OwnerTransferred(address indexed prevOwner, address indexed newOwner)
func (_Meson *MesonFilterer) ParseOwnerTransferred(log types.Log) (*MesonOwnerTransferred, error) {
	event := new(MesonOwnerTransferred)
	if err := _Meson.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPoolAuthorizedAddrAddedIterator is returned from FilterPoolAuthorizedAddrAdded and is used to iterate over the raw logs and unpacked data for PoolAuthorizedAddrAdded events raised by the Meson contract.
type MesonPoolAuthorizedAddrAddedIterator struct {
	Event *MesonPoolAuthorizedAddrAdded // Event containing the contract specifics and raw log

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
func (it *MesonPoolAuthorizedAddrAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPoolAuthorizedAddrAdded)
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
		it.Event = new(MesonPoolAuthorizedAddrAdded)
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
func (it *MesonPoolAuthorizedAddrAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPoolAuthorizedAddrAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPoolAuthorizedAddrAdded represents a PoolAuthorizedAddrAdded event raised by the Meson contract.
type MesonPoolAuthorizedAddrAdded struct {
	PoolIndex *big.Int
	Addr      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolAuthorizedAddrAdded is a free log retrieval operation binding the contract event 0xd49cde4f679ccef3d23ff07aae4f6845e1c661e23e9fe6a54da26f0723fb695f.
//
// Solidity: event PoolAuthorizedAddrAdded(uint40 indexed poolIndex, address addr)
func (_Meson *MesonFilterer) FilterPoolAuthorizedAddrAdded(opts *bind.FilterOpts, poolIndex []*big.Int) (*MesonPoolAuthorizedAddrAddedIterator, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PoolAuthorizedAddrAdded", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return &MesonPoolAuthorizedAddrAddedIterator{contract: _Meson.contract, event: "PoolAuthorizedAddrAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAuthorizedAddrAdded is a free log subscription operation binding the contract event 0xd49cde4f679ccef3d23ff07aae4f6845e1c661e23e9fe6a54da26f0723fb695f.
//
// Solidity: event PoolAuthorizedAddrAdded(uint40 indexed poolIndex, address addr)
func (_Meson *MesonFilterer) WatchPoolAuthorizedAddrAdded(opts *bind.WatchOpts, sink chan<- *MesonPoolAuthorizedAddrAdded, poolIndex []*big.Int) (event.Subscription, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PoolAuthorizedAddrAdded", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPoolAuthorizedAddrAdded)
				if err := _Meson.contract.UnpackLog(event, "PoolAuthorizedAddrAdded", log); err != nil {
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

// ParsePoolAuthorizedAddrAdded is a log parse operation binding the contract event 0xd49cde4f679ccef3d23ff07aae4f6845e1c661e23e9fe6a54da26f0723fb695f.
//
// Solidity: event PoolAuthorizedAddrAdded(uint40 indexed poolIndex, address addr)
func (_Meson *MesonFilterer) ParsePoolAuthorizedAddrAdded(log types.Log) (*MesonPoolAuthorizedAddrAdded, error) {
	event := new(MesonPoolAuthorizedAddrAdded)
	if err := _Meson.contract.UnpackLog(event, "PoolAuthorizedAddrAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPoolAuthorizedAddrRemovedIterator is returned from FilterPoolAuthorizedAddrRemoved and is used to iterate over the raw logs and unpacked data for PoolAuthorizedAddrRemoved events raised by the Meson contract.
type MesonPoolAuthorizedAddrRemovedIterator struct {
	Event *MesonPoolAuthorizedAddrRemoved // Event containing the contract specifics and raw log

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
func (it *MesonPoolAuthorizedAddrRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPoolAuthorizedAddrRemoved)
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
		it.Event = new(MesonPoolAuthorizedAddrRemoved)
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
func (it *MesonPoolAuthorizedAddrRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPoolAuthorizedAddrRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPoolAuthorizedAddrRemoved represents a PoolAuthorizedAddrRemoved event raised by the Meson contract.
type MesonPoolAuthorizedAddrRemoved struct {
	PoolIndex *big.Int
	Addr      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolAuthorizedAddrRemoved is a free log retrieval operation binding the contract event 0x475b83c893df40ee19fd0783cf26478cdb58478dff65bb62560e1e7c36e0f22f.
//
// Solidity: event PoolAuthorizedAddrRemoved(uint40 indexed poolIndex, address addr)
func (_Meson *MesonFilterer) FilterPoolAuthorizedAddrRemoved(opts *bind.FilterOpts, poolIndex []*big.Int) (*MesonPoolAuthorizedAddrRemovedIterator, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PoolAuthorizedAddrRemoved", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return &MesonPoolAuthorizedAddrRemovedIterator{contract: _Meson.contract, event: "PoolAuthorizedAddrRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolAuthorizedAddrRemoved is a free log subscription operation binding the contract event 0x475b83c893df40ee19fd0783cf26478cdb58478dff65bb62560e1e7c36e0f22f.
//
// Solidity: event PoolAuthorizedAddrRemoved(uint40 indexed poolIndex, address addr)
func (_Meson *MesonFilterer) WatchPoolAuthorizedAddrRemoved(opts *bind.WatchOpts, sink chan<- *MesonPoolAuthorizedAddrRemoved, poolIndex []*big.Int) (event.Subscription, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PoolAuthorizedAddrRemoved", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPoolAuthorizedAddrRemoved)
				if err := _Meson.contract.UnpackLog(event, "PoolAuthorizedAddrRemoved", log); err != nil {
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

// ParsePoolAuthorizedAddrRemoved is a log parse operation binding the contract event 0x475b83c893df40ee19fd0783cf26478cdb58478dff65bb62560e1e7c36e0f22f.
//
// Solidity: event PoolAuthorizedAddrRemoved(uint40 indexed poolIndex, address addr)
func (_Meson *MesonFilterer) ParsePoolAuthorizedAddrRemoved(log types.Log) (*MesonPoolAuthorizedAddrRemoved, error) {
	event := new(MesonPoolAuthorizedAddrRemoved)
	if err := _Meson.contract.UnpackLog(event, "PoolAuthorizedAddrRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPoolDepositedIterator is returned from FilterPoolDeposited and is used to iterate over the raw logs and unpacked data for PoolDeposited events raised by the Meson contract.
type MesonPoolDepositedIterator struct {
	Event *MesonPoolDeposited // Event containing the contract specifics and raw log

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
func (it *MesonPoolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPoolDeposited)
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
		it.Event = new(MesonPoolDeposited)
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
func (it *MesonPoolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPoolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPoolDeposited represents a PoolDeposited event raised by the Meson contract.
type MesonPoolDeposited struct {
	PoolTokenIndex *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolDeposited is a free log retrieval operation binding the contract event 0x7d7d1df74ef3a6434d8d63dc0a25d13d5fa94dbe738c38a3cce26e6f892e2a76.
//
// Solidity: event PoolDeposited(uint48 indexed poolTokenIndex, uint256 amount)
func (_Meson *MesonFilterer) FilterPoolDeposited(opts *bind.FilterOpts, poolTokenIndex []*big.Int) (*MesonPoolDepositedIterator, error) {

	var poolTokenIndexRule []interface{}
	for _, poolTokenIndexItem := range poolTokenIndex {
		poolTokenIndexRule = append(poolTokenIndexRule, poolTokenIndexItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PoolDeposited", poolTokenIndexRule)
	if err != nil {
		return nil, err
	}
	return &MesonPoolDepositedIterator{contract: _Meson.contract, event: "PoolDeposited", logs: logs, sub: sub}, nil
}

// WatchPoolDeposited is a free log subscription operation binding the contract event 0x7d7d1df74ef3a6434d8d63dc0a25d13d5fa94dbe738c38a3cce26e6f892e2a76.
//
// Solidity: event PoolDeposited(uint48 indexed poolTokenIndex, uint256 amount)
func (_Meson *MesonFilterer) WatchPoolDeposited(opts *bind.WatchOpts, sink chan<- *MesonPoolDeposited, poolTokenIndex []*big.Int) (event.Subscription, error) {

	var poolTokenIndexRule []interface{}
	for _, poolTokenIndexItem := range poolTokenIndex {
		poolTokenIndexRule = append(poolTokenIndexRule, poolTokenIndexItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PoolDeposited", poolTokenIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPoolDeposited)
				if err := _Meson.contract.UnpackLog(event, "PoolDeposited", log); err != nil {
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

// ParsePoolDeposited is a log parse operation binding the contract event 0x7d7d1df74ef3a6434d8d63dc0a25d13d5fa94dbe738c38a3cce26e6f892e2a76.
//
// Solidity: event PoolDeposited(uint48 indexed poolTokenIndex, uint256 amount)
func (_Meson *MesonFilterer) ParsePoolDeposited(log types.Log) (*MesonPoolDeposited, error) {
	event := new(MesonPoolDeposited)
	if err := _Meson.contract.UnpackLog(event, "PoolDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPoolOwnerTransferredIterator is returned from FilterPoolOwnerTransferred and is used to iterate over the raw logs and unpacked data for PoolOwnerTransferred events raised by the Meson contract.
type MesonPoolOwnerTransferredIterator struct {
	Event *MesonPoolOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *MesonPoolOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPoolOwnerTransferred)
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
		it.Event = new(MesonPoolOwnerTransferred)
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
func (it *MesonPoolOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPoolOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPoolOwnerTransferred represents a PoolOwnerTransferred event raised by the Meson contract.
type MesonPoolOwnerTransferred struct {
	PoolIndex *big.Int
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolOwnerTransferred is a free log retrieval operation binding the contract event 0xc94089e0c0b1b79fdecc6e64fb759cdd390590a15c7e50d281e681ea8273261c.
//
// Solidity: event PoolOwnerTransferred(uint40 indexed poolIndex, address prevOwner, address newOwner)
func (_Meson *MesonFilterer) FilterPoolOwnerTransferred(opts *bind.FilterOpts, poolIndex []*big.Int) (*MesonPoolOwnerTransferredIterator, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PoolOwnerTransferred", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return &MesonPoolOwnerTransferredIterator{contract: _Meson.contract, event: "PoolOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchPoolOwnerTransferred is a free log subscription operation binding the contract event 0xc94089e0c0b1b79fdecc6e64fb759cdd390590a15c7e50d281e681ea8273261c.
//
// Solidity: event PoolOwnerTransferred(uint40 indexed poolIndex, address prevOwner, address newOwner)
func (_Meson *MesonFilterer) WatchPoolOwnerTransferred(opts *bind.WatchOpts, sink chan<- *MesonPoolOwnerTransferred, poolIndex []*big.Int) (event.Subscription, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PoolOwnerTransferred", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPoolOwnerTransferred)
				if err := _Meson.contract.UnpackLog(event, "PoolOwnerTransferred", log); err != nil {
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

// ParsePoolOwnerTransferred is a log parse operation binding the contract event 0xc94089e0c0b1b79fdecc6e64fb759cdd390590a15c7e50d281e681ea8273261c.
//
// Solidity: event PoolOwnerTransferred(uint40 indexed poolIndex, address prevOwner, address newOwner)
func (_Meson *MesonFilterer) ParsePoolOwnerTransferred(log types.Log) (*MesonPoolOwnerTransferred, error) {
	event := new(MesonPoolOwnerTransferred)
	if err := _Meson.contract.UnpackLog(event, "PoolOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Meson contract.
type MesonPoolRegisteredIterator struct {
	Event *MesonPoolRegistered // Event containing the contract specifics and raw log

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
func (it *MesonPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPoolRegistered)
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
		it.Event = new(MesonPoolRegistered)
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
func (it *MesonPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPoolRegistered represents a PoolRegistered event raised by the Meson contract.
type MesonPoolRegistered struct {
	PoolIndex *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolRegistered is a free log retrieval operation binding the contract event 0xb8d9c35a714d4e29eaf036b9bf8183a093c5573ac809453b4e8434e25c9126d2.
//
// Solidity: event PoolRegistered(uint40 indexed poolIndex, address owner)
func (_Meson *MesonFilterer) FilterPoolRegistered(opts *bind.FilterOpts, poolIndex []*big.Int) (*MesonPoolRegisteredIterator, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PoolRegistered", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return &MesonPoolRegisteredIterator{contract: _Meson.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0xb8d9c35a714d4e29eaf036b9bf8183a093c5573ac809453b4e8434e25c9126d2.
//
// Solidity: event PoolRegistered(uint40 indexed poolIndex, address owner)
func (_Meson *MesonFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *MesonPoolRegistered, poolIndex []*big.Int) (event.Subscription, error) {

	var poolIndexRule []interface{}
	for _, poolIndexItem := range poolIndex {
		poolIndexRule = append(poolIndexRule, poolIndexItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PoolRegistered", poolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPoolRegistered)
				if err := _Meson.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
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

// ParsePoolRegistered is a log parse operation binding the contract event 0xb8d9c35a714d4e29eaf036b9bf8183a093c5573ac809453b4e8434e25c9126d2.
//
// Solidity: event PoolRegistered(uint40 indexed poolIndex, address owner)
func (_Meson *MesonFilterer) ParsePoolRegistered(log types.Log) (*MesonPoolRegistered, error) {
	event := new(MesonPoolRegistered)
	if err := _Meson.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPoolWithdrawnIterator is returned from FilterPoolWithdrawn and is used to iterate over the raw logs and unpacked data for PoolWithdrawn events raised by the Meson contract.
type MesonPoolWithdrawnIterator struct {
	Event *MesonPoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *MesonPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPoolWithdrawn)
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
		it.Event = new(MesonPoolWithdrawn)
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
func (it *MesonPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPoolWithdrawn represents a PoolWithdrawn event raised by the Meson contract.
type MesonPoolWithdrawn struct {
	PoolTokenIndex *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolWithdrawn is a free log retrieval operation binding the contract event 0x34c3d1c46f89307d63d8818fcc5c2a9c07a5f7a01ea4319bfba1899f40c6f400.
//
// Solidity: event PoolWithdrawn(uint48 indexed poolTokenIndex, uint256 amount)
func (_Meson *MesonFilterer) FilterPoolWithdrawn(opts *bind.FilterOpts, poolTokenIndex []*big.Int) (*MesonPoolWithdrawnIterator, error) {

	var poolTokenIndexRule []interface{}
	for _, poolTokenIndexItem := range poolTokenIndex {
		poolTokenIndexRule = append(poolTokenIndexRule, poolTokenIndexItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PoolWithdrawn", poolTokenIndexRule)
	if err != nil {
		return nil, err
	}
	return &MesonPoolWithdrawnIterator{contract: _Meson.contract, event: "PoolWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPoolWithdrawn is a free log subscription operation binding the contract event 0x34c3d1c46f89307d63d8818fcc5c2a9c07a5f7a01ea4319bfba1899f40c6f400.
//
// Solidity: event PoolWithdrawn(uint48 indexed poolTokenIndex, uint256 amount)
func (_Meson *MesonFilterer) WatchPoolWithdrawn(opts *bind.WatchOpts, sink chan<- *MesonPoolWithdrawn, poolTokenIndex []*big.Int) (event.Subscription, error) {

	var poolTokenIndexRule []interface{}
	for _, poolTokenIndexItem := range poolTokenIndex {
		poolTokenIndexRule = append(poolTokenIndexRule, poolTokenIndexItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PoolWithdrawn", poolTokenIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPoolWithdrawn)
				if err := _Meson.contract.UnpackLog(event, "PoolWithdrawn", log); err != nil {
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

// ParsePoolWithdrawn is a log parse operation binding the contract event 0x34c3d1c46f89307d63d8818fcc5c2a9c07a5f7a01ea4319bfba1899f40c6f400.
//
// Solidity: event PoolWithdrawn(uint48 indexed poolTokenIndex, uint256 amount)
func (_Meson *MesonFilterer) ParsePoolWithdrawn(log types.Log) (*MesonPoolWithdrawn, error) {
	event := new(MesonPoolWithdrawn)
	if err := _Meson.contract.UnpackLog(event, "PoolWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonPremiumManagerTransferredIterator is returned from FilterPremiumManagerTransferred and is used to iterate over the raw logs and unpacked data for PremiumManagerTransferred events raised by the Meson contract.
type MesonPremiumManagerTransferredIterator struct {
	Event *MesonPremiumManagerTransferred // Event containing the contract specifics and raw log

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
func (it *MesonPremiumManagerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonPremiumManagerTransferred)
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
		it.Event = new(MesonPremiumManagerTransferred)
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
func (it *MesonPremiumManagerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonPremiumManagerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonPremiumManagerTransferred represents a PremiumManagerTransferred event raised by the Meson contract.
type MesonPremiumManagerTransferred struct {
	PrevPremiumManager common.Address
	NewPremiumManager  common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPremiumManagerTransferred is a free log retrieval operation binding the contract event 0x4798f31ad3d0ccde6359edf35fc39b882e4e1cff2968ca749b72074d373db27a.
//
// Solidity: event PremiumManagerTransferred(address indexed prevPremiumManager, address indexed newPremiumManager)
func (_Meson *MesonFilterer) FilterPremiumManagerTransferred(opts *bind.FilterOpts, prevPremiumManager []common.Address, newPremiumManager []common.Address) (*MesonPremiumManagerTransferredIterator, error) {

	var prevPremiumManagerRule []interface{}
	for _, prevPremiumManagerItem := range prevPremiumManager {
		prevPremiumManagerRule = append(prevPremiumManagerRule, prevPremiumManagerItem)
	}
	var newPremiumManagerRule []interface{}
	for _, newPremiumManagerItem := range newPremiumManager {
		newPremiumManagerRule = append(newPremiumManagerRule, newPremiumManagerItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "PremiumManagerTransferred", prevPremiumManagerRule, newPremiumManagerRule)
	if err != nil {
		return nil, err
	}
	return &MesonPremiumManagerTransferredIterator{contract: _Meson.contract, event: "PremiumManagerTransferred", logs: logs, sub: sub}, nil
}

// WatchPremiumManagerTransferred is a free log subscription operation binding the contract event 0x4798f31ad3d0ccde6359edf35fc39b882e4e1cff2968ca749b72074d373db27a.
//
// Solidity: event PremiumManagerTransferred(address indexed prevPremiumManager, address indexed newPremiumManager)
func (_Meson *MesonFilterer) WatchPremiumManagerTransferred(opts *bind.WatchOpts, sink chan<- *MesonPremiumManagerTransferred, prevPremiumManager []common.Address, newPremiumManager []common.Address) (event.Subscription, error) {

	var prevPremiumManagerRule []interface{}
	for _, prevPremiumManagerItem := range prevPremiumManager {
		prevPremiumManagerRule = append(prevPremiumManagerRule, prevPremiumManagerItem)
	}
	var newPremiumManagerRule []interface{}
	for _, newPremiumManagerItem := range newPremiumManager {
		newPremiumManagerRule = append(newPremiumManagerRule, newPremiumManagerItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "PremiumManagerTransferred", prevPremiumManagerRule, newPremiumManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonPremiumManagerTransferred)
				if err := _Meson.contract.UnpackLog(event, "PremiumManagerTransferred", log); err != nil {
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

// ParsePremiumManagerTransferred is a log parse operation binding the contract event 0x4798f31ad3d0ccde6359edf35fc39b882e4e1cff2968ca749b72074d373db27a.
//
// Solidity: event PremiumManagerTransferred(address indexed prevPremiumManager, address indexed newPremiumManager)
func (_Meson *MesonFilterer) ParsePremiumManagerTransferred(log types.Log) (*MesonPremiumManagerTransferred, error) {
	event := new(MesonPremiumManagerTransferred)
	if err := _Meson.contract.UnpackLog(event, "PremiumManagerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapBondedIterator is returned from FilterSwapBonded and is used to iterate over the raw logs and unpacked data for SwapBonded events raised by the Meson contract.
type MesonSwapBondedIterator struct {
	Event *MesonSwapBonded // Event containing the contract specifics and raw log

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
func (it *MesonSwapBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapBonded)
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
		it.Event = new(MesonSwapBonded)
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
func (it *MesonSwapBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapBonded represents a SwapBonded event raised by the Meson contract.
type MesonSwapBonded struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapBonded is a free log retrieval operation binding the contract event 0x60a99b51ae498c44acbbe11031aed2a06a32be66d2122e6e2a7a16c087865cc9.
//
// Solidity: event SwapBonded(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapBonded(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapBondedIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapBonded", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapBondedIterator{contract: _Meson.contract, event: "SwapBonded", logs: logs, sub: sub}, nil
}

// WatchSwapBonded is a free log subscription operation binding the contract event 0x60a99b51ae498c44acbbe11031aed2a06a32be66d2122e6e2a7a16c087865cc9.
//
// Solidity: event SwapBonded(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapBonded(opts *bind.WatchOpts, sink chan<- *MesonSwapBonded, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapBonded", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapBonded)
				if err := _Meson.contract.UnpackLog(event, "SwapBonded", log); err != nil {
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

// ParseSwapBonded is a log parse operation binding the contract event 0x60a99b51ae498c44acbbe11031aed2a06a32be66d2122e6e2a7a16c087865cc9.
//
// Solidity: event SwapBonded(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapBonded(log types.Log) (*MesonSwapBonded, error) {
	event := new(MesonSwapBonded)
	if err := _Meson.contract.UnpackLog(event, "SwapBonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapCancelledIterator is returned from FilterSwapCancelled and is used to iterate over the raw logs and unpacked data for SwapCancelled events raised by the Meson contract.
type MesonSwapCancelledIterator struct {
	Event *MesonSwapCancelled // Event containing the contract specifics and raw log

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
func (it *MesonSwapCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapCancelled)
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
		it.Event = new(MesonSwapCancelled)
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
func (it *MesonSwapCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapCancelled represents a SwapCancelled event raised by the Meson contract.
type MesonSwapCancelled struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapCancelled is a free log retrieval operation binding the contract event 0xf6b6b4f7a13f02512c1b3aa8dcc4a07d7775a6a4becbd439efcbd37c5408e67f.
//
// Solidity: event SwapCancelled(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapCancelled(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapCancelledIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapCancelled", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapCancelledIterator{contract: _Meson.contract, event: "SwapCancelled", logs: logs, sub: sub}, nil
}

// WatchSwapCancelled is a free log subscription operation binding the contract event 0xf6b6b4f7a13f02512c1b3aa8dcc4a07d7775a6a4becbd439efcbd37c5408e67f.
//
// Solidity: event SwapCancelled(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapCancelled(opts *bind.WatchOpts, sink chan<- *MesonSwapCancelled, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapCancelled", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapCancelled)
				if err := _Meson.contract.UnpackLog(event, "SwapCancelled", log); err != nil {
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

// ParseSwapCancelled is a log parse operation binding the contract event 0xf6b6b4f7a13f02512c1b3aa8dcc4a07d7775a6a4becbd439efcbd37c5408e67f.
//
// Solidity: event SwapCancelled(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapCancelled(log types.Log) (*MesonSwapCancelled, error) {
	event := new(MesonSwapCancelled)
	if err := _Meson.contract.UnpackLog(event, "SwapCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapExecutedIterator is returned from FilterSwapExecuted and is used to iterate over the raw logs and unpacked data for SwapExecuted events raised by the Meson contract.
type MesonSwapExecutedIterator struct {
	Event *MesonSwapExecuted // Event containing the contract specifics and raw log

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
func (it *MesonSwapExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapExecuted)
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
		it.Event = new(MesonSwapExecuted)
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
func (it *MesonSwapExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapExecuted represents a SwapExecuted event raised by the Meson contract.
type MesonSwapExecuted struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapExecuted is a free log retrieval operation binding the contract event 0x8d92c805c252261fcfff21ee60740eb8a38922469a7e6ee396976d57c22fc1c9.
//
// Solidity: event SwapExecuted(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapExecuted(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapExecutedIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapExecuted", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapExecutedIterator{contract: _Meson.contract, event: "SwapExecuted", logs: logs, sub: sub}, nil
}

// WatchSwapExecuted is a free log subscription operation binding the contract event 0x8d92c805c252261fcfff21ee60740eb8a38922469a7e6ee396976d57c22fc1c9.
//
// Solidity: event SwapExecuted(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapExecuted(opts *bind.WatchOpts, sink chan<- *MesonSwapExecuted, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapExecuted", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapExecuted)
				if err := _Meson.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
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

// ParseSwapExecuted is a log parse operation binding the contract event 0x8d92c805c252261fcfff21ee60740eb8a38922469a7e6ee396976d57c22fc1c9.
//
// Solidity: event SwapExecuted(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapExecuted(log types.Log) (*MesonSwapExecuted, error) {
	event := new(MesonSwapExecuted)
	if err := _Meson.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapLockedIterator is returned from FilterSwapLocked and is used to iterate over the raw logs and unpacked data for SwapLocked events raised by the Meson contract.
type MesonSwapLockedIterator struct {
	Event *MesonSwapLocked // Event containing the contract specifics and raw log

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
func (it *MesonSwapLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapLocked)
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
		it.Event = new(MesonSwapLocked)
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
func (it *MesonSwapLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapLocked represents a SwapLocked event raised by the Meson contract.
type MesonSwapLocked struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapLocked is a free log retrieval operation binding the contract event 0xbfb879c34323c5601fafe832c3a8a1e31e12c288695838726ddeada86034edb4.
//
// Solidity: event SwapLocked(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapLocked(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapLockedIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapLocked", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapLockedIterator{contract: _Meson.contract, event: "SwapLocked", logs: logs, sub: sub}, nil
}

// WatchSwapLocked is a free log subscription operation binding the contract event 0xbfb879c34323c5601fafe832c3a8a1e31e12c288695838726ddeada86034edb4.
//
// Solidity: event SwapLocked(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapLocked(opts *bind.WatchOpts, sink chan<- *MesonSwapLocked, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapLocked", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapLocked)
				if err := _Meson.contract.UnpackLog(event, "SwapLocked", log); err != nil {
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

// ParseSwapLocked is a log parse operation binding the contract event 0xbfb879c34323c5601fafe832c3a8a1e31e12c288695838726ddeada86034edb4.
//
// Solidity: event SwapLocked(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapLocked(log types.Log) (*MesonSwapLocked, error) {
	event := new(MesonSwapLocked)
	if err := _Meson.contract.UnpackLog(event, "SwapLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapPostedIterator is returned from FilterSwapPosted and is used to iterate over the raw logs and unpacked data for SwapPosted events raised by the Meson contract.
type MesonSwapPostedIterator struct {
	Event *MesonSwapPosted // Event containing the contract specifics and raw log

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
func (it *MesonSwapPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapPosted)
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
		it.Event = new(MesonSwapPosted)
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
func (it *MesonSwapPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapPosted represents a SwapPosted event raised by the Meson contract.
type MesonSwapPosted struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapPosted is a free log retrieval operation binding the contract event 0x5ce4019f772fda6cb703b26bce3ec3006eb36b73f1d3a0eb441213317d9f5e9d.
//
// Solidity: event SwapPosted(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapPosted(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapPostedIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapPosted", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapPostedIterator{contract: _Meson.contract, event: "SwapPosted", logs: logs, sub: sub}, nil
}

// WatchSwapPosted is a free log subscription operation binding the contract event 0x5ce4019f772fda6cb703b26bce3ec3006eb36b73f1d3a0eb441213317d9f5e9d.
//
// Solidity: event SwapPosted(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapPosted(opts *bind.WatchOpts, sink chan<- *MesonSwapPosted, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapPosted", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapPosted)
				if err := _Meson.contract.UnpackLog(event, "SwapPosted", log); err != nil {
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

// ParseSwapPosted is a log parse operation binding the contract event 0x5ce4019f772fda6cb703b26bce3ec3006eb36b73f1d3a0eb441213317d9f5e9d.
//
// Solidity: event SwapPosted(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapPosted(log types.Log) (*MesonSwapPosted, error) {
	event := new(MesonSwapPosted)
	if err := _Meson.contract.UnpackLog(event, "SwapPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapReleasedIterator is returned from FilterSwapReleased and is used to iterate over the raw logs and unpacked data for SwapReleased events raised by the Meson contract.
type MesonSwapReleasedIterator struct {
	Event *MesonSwapReleased // Event containing the contract specifics and raw log

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
func (it *MesonSwapReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapReleased)
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
		it.Event = new(MesonSwapReleased)
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
func (it *MesonSwapReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapReleased represents a SwapReleased event raised by the Meson contract.
type MesonSwapReleased struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapReleased is a free log retrieval operation binding the contract event 0xfa628b578e095243f0544bfad9255f49d79d03a5bbf6c85875d05a215e247ad2.
//
// Solidity: event SwapReleased(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapReleased(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapReleasedIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapReleased", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapReleasedIterator{contract: _Meson.contract, event: "SwapReleased", logs: logs, sub: sub}, nil
}

// WatchSwapReleased is a free log subscription operation binding the contract event 0xfa628b578e095243f0544bfad9255f49d79d03a5bbf6c85875d05a215e247ad2.
//
// Solidity: event SwapReleased(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapReleased(opts *bind.WatchOpts, sink chan<- *MesonSwapReleased, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapReleased", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapReleased)
				if err := _Meson.contract.UnpackLog(event, "SwapReleased", log); err != nil {
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

// ParseSwapReleased is a log parse operation binding the contract event 0xfa628b578e095243f0544bfad9255f49d79d03a5bbf6c85875d05a215e247ad2.
//
// Solidity: event SwapReleased(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapReleased(log types.Log) (*MesonSwapReleased, error) {
	event := new(MesonSwapReleased)
	if err := _Meson.contract.UnpackLog(event, "SwapReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MesonSwapUnlockedIterator is returned from FilterSwapUnlocked and is used to iterate over the raw logs and unpacked data for SwapUnlocked events raised by the Meson contract.
type MesonSwapUnlockedIterator struct {
	Event *MesonSwapUnlocked // Event containing the contract specifics and raw log

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
func (it *MesonSwapUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MesonSwapUnlocked)
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
		it.Event = new(MesonSwapUnlocked)
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
func (it *MesonSwapUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MesonSwapUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MesonSwapUnlocked represents a SwapUnlocked event raised by the Meson contract.
type MesonSwapUnlocked struct {
	EncodedSwap *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapUnlocked is a free log retrieval operation binding the contract event 0xac7d23c4f0137a4cc35b0e4b4bc8061ea6cb65805e87ceb0a77ca0c85814858c.
//
// Solidity: event SwapUnlocked(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) FilterSwapUnlocked(opts *bind.FilterOpts, encodedSwap []*big.Int) (*MesonSwapUnlockedIterator, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.FilterLogs(opts, "SwapUnlocked", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return &MesonSwapUnlockedIterator{contract: _Meson.contract, event: "SwapUnlocked", logs: logs, sub: sub}, nil
}

// WatchSwapUnlocked is a free log subscription operation binding the contract event 0xac7d23c4f0137a4cc35b0e4b4bc8061ea6cb65805e87ceb0a77ca0c85814858c.
//
// Solidity: event SwapUnlocked(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) WatchSwapUnlocked(opts *bind.WatchOpts, sink chan<- *MesonSwapUnlocked, encodedSwap []*big.Int) (event.Subscription, error) {

	var encodedSwapRule []interface{}
	for _, encodedSwapItem := range encodedSwap {
		encodedSwapRule = append(encodedSwapRule, encodedSwapItem)
	}

	logs, sub, err := _Meson.contract.WatchLogs(opts, "SwapUnlocked", encodedSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MesonSwapUnlocked)
				if err := _Meson.contract.UnpackLog(event, "SwapUnlocked", log); err != nil {
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

// ParseSwapUnlocked is a log parse operation binding the contract event 0xac7d23c4f0137a4cc35b0e4b4bc8061ea6cb65805e87ceb0a77ca0c85814858c.
//
// Solidity: event SwapUnlocked(uint256 indexed encodedSwap)
func (_Meson *MesonFilterer) ParseSwapUnlocked(log types.Log) (*MesonSwapUnlocked, error) {
	event := new(MesonSwapUnlocked)
	if err := _Meson.contract.UnpackLog(event, "SwapUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
