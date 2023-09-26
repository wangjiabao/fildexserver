// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// ManagerFilMetaData contains all meta data concerning the ManagerFil contract.
var ManagerFilMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stfil_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stfilToken_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"exchange_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeTo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRewardAll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeWitdhrawRewardAll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"feeWithdrawRewardCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAll\",\"type\":\"uint256\"}],\"name\":\"payCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAll\",\"type\":\"uint256\"}],\"name\":\"rePayCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRewardAll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentReward\",\"type\":\"uint256\"}],\"name\":\"rewardCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"witdhrawRewardAll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRewardCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminRoleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRewardAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"feeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeWithdrawRewardAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardFeeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"feeTo_\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setRewardFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawStfil_\",\"type\":\"address\"}],\"name\":\"setWithdrawStfil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stfil\",\"outputs\":[{\"internalType\":\"contractISTFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stfilToken\",\"outputs\":[{\"internalType\":\"contractISTFILToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witdhrawRewardAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStfil\",\"outputs\":[{\"internalType\":\"contractIWithdrawStfil\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ManagerFilABI is the input ABI used to generate the binding from.
// Deprecated: Use ManagerFilMetaData.ABI instead.
var ManagerFilABI = ManagerFilMetaData.ABI

// ManagerFil is an auto generated Go binding around an Ethereum contract.
type ManagerFil struct {
	ManagerFilCaller     // Read-only binding to the contract
	ManagerFilTransactor // Write-only binding to the contract
	ManagerFilFilterer   // Log filterer for contract events
}

// ManagerFilCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerFilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerFilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerFilSession struct {
	Contract     *ManagerFil       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerFilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerFilCallerSession struct {
	Contract *ManagerFilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ManagerFilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerFilTransactorSession struct {
	Contract     *ManagerFilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ManagerFilRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerFilRaw struct {
	Contract *ManagerFil // Generic contract binding to access the raw methods on
}

// ManagerFilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerFilCallerRaw struct {
	Contract *ManagerFilCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerFilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerFilTransactorRaw struct {
	Contract *ManagerFilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManagerFil creates a new instance of ManagerFil, bound to a specific deployed contract.
func NewManagerFil(address common.Address, backend bind.ContractBackend) (*ManagerFil, error) {
	contract, err := bindManagerFil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ManagerFil{ManagerFilCaller: ManagerFilCaller{contract: contract}, ManagerFilTransactor: ManagerFilTransactor{contract: contract}, ManagerFilFilterer: ManagerFilFilterer{contract: contract}}, nil
}

// NewManagerFilCaller creates a new read-only instance of ManagerFil, bound to a specific deployed contract.
func NewManagerFilCaller(address common.Address, caller bind.ContractCaller) (*ManagerFilCaller, error) {
	contract, err := bindManagerFil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerFilCaller{contract: contract}, nil
}

// NewManagerFilTransactor creates a new write-only instance of ManagerFil, bound to a specific deployed contract.
func NewManagerFilTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerFilTransactor, error) {
	contract, err := bindManagerFil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerFilTransactor{contract: contract}, nil
}

// NewManagerFilFilterer creates a new log filterer instance of ManagerFil, bound to a specific deployed contract.
func NewManagerFilFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilFilterer, error) {
	contract, err := bindManagerFil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilFilterer{contract: contract}, nil
}

// bindManagerFil binds a generic wrapper to an already deployed contract.
func bindManagerFil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerFilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManagerFil *ManagerFilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManagerFil.Contract.ManagerFilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManagerFil *ManagerFilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagerFil.Contract.ManagerFilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManagerFil *ManagerFilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManagerFil.Contract.ManagerFilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManagerFil *ManagerFilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManagerFil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManagerFil *ManagerFilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagerFil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManagerFil *ManagerFilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManagerFil.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ManagerFil *ManagerFilCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ManagerFil *ManagerFilSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ManagerFil.Contract.DEFAULTADMINROLE(&_ManagerFil.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ManagerFil *ManagerFilCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ManagerFil.Contract.DEFAULTADMINROLE(&_ManagerFil.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_ManagerFil *ManagerFilCaller) SUPERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "SUPER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_ManagerFil *ManagerFilSession) SUPERADMINROLE() ([32]byte, error) {
	return _ManagerFil.Contract.SUPERADMINROLE(&_ManagerFil.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_ManagerFil *ManagerFilCallerSession) SUPERADMINROLE() ([32]byte, error) {
	return _ManagerFil.Contract.SUPERADMINROLE(&_ManagerFil.CallOpts)
}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) DefaultAdminRoleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "defaultAdminRoleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_ManagerFil *ManagerFilSession) DefaultAdminRoleLimit() (*big.Int, error) {
	return _ManagerFil.Contract.DefaultAdminRoleLimit(&_ManagerFil.CallOpts)
}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) DefaultAdminRoleLimit() (*big.Int, error) {
	return _ManagerFil.Contract.DefaultAdminRoleLimit(&_ManagerFil.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ManagerFil *ManagerFilCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ManagerFil *ManagerFilSession) Exchange() (common.Address, error) {
	return _ManagerFil.Contract.Exchange(&_ManagerFil.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ManagerFil *ManagerFilCallerSession) Exchange() (common.Address, error) {
	return _ManagerFil.Contract.Exchange(&_ManagerFil.CallOpts)
}

// FeeRewardAll is a free data retrieval call binding the contract method 0xb5b8aff1.
//
// Solidity: function feeRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) FeeRewardAll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "feeRewardAll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRewardAll is a free data retrieval call binding the contract method 0xb5b8aff1.
//
// Solidity: function feeRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilSession) FeeRewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.FeeRewardAll(&_ManagerFil.CallOpts)
}

// FeeRewardAll is a free data retrieval call binding the contract method 0xb5b8aff1.
//
// Solidity: function feeRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) FeeRewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.FeeRewardAll(&_ManagerFil.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ManagerFil *ManagerFilCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ManagerFil *ManagerFilSession) FeeTo() (common.Address, error) {
	return _ManagerFil.Contract.FeeTo(&_ManagerFil.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ManagerFil *ManagerFilCallerSession) FeeTo() (common.Address, error) {
	return _ManagerFil.Contract.FeeTo(&_ManagerFil.CallOpts)
}

// FeeWithdrawRewardAll is a free data retrieval call binding the contract method 0xe839d7d0.
//
// Solidity: function feeWithdrawRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) FeeWithdrawRewardAll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "feeWithdrawRewardAll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeWithdrawRewardAll is a free data retrieval call binding the contract method 0xe839d7d0.
//
// Solidity: function feeWithdrawRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilSession) FeeWithdrawRewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.FeeWithdrawRewardAll(&_ManagerFil.CallOpts)
}

// FeeWithdrawRewardAll is a free data retrieval call binding the contract method 0xe839d7d0.
//
// Solidity: function feeWithdrawRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) FeeWithdrawRewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.FeeWithdrawRewardAll(&_ManagerFil.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ManagerFil *ManagerFilCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ManagerFil *ManagerFilSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ManagerFil.Contract.GetRoleAdmin(&_ManagerFil.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ManagerFil *ManagerFilCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ManagerFil.Contract.GetRoleAdmin(&_ManagerFil.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ManagerFil *ManagerFilCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ManagerFil *ManagerFilSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ManagerFil.Contract.GetRoleMember(&_ManagerFil.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ManagerFil *ManagerFilCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ManagerFil.Contract.GetRoleMember(&_ManagerFil.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ManagerFil *ManagerFilCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ManagerFil *ManagerFilSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ManagerFil.Contract.GetRoleMemberCount(&_ManagerFil.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ManagerFil.Contract.GetRoleMemberCount(&_ManagerFil.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ManagerFil *ManagerFilCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ManagerFil *ManagerFilSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ManagerFil.Contract.HasRole(&_ManagerFil.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ManagerFil *ManagerFilCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ManagerFil.Contract.HasRole(&_ManagerFil.CallOpts, role, account)
}

// LastReward is a free data retrieval call binding the contract method 0xc9b17149.
//
// Solidity: function lastReward() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) LastReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "lastReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReward is a free data retrieval call binding the contract method 0xc9b17149.
//
// Solidity: function lastReward() view returns(uint256)
func (_ManagerFil *ManagerFilSession) LastReward() (*big.Int, error) {
	return _ManagerFil.Contract.LastReward(&_ManagerFil.CallOpts)
}

// LastReward is a free data retrieval call binding the contract method 0xc9b17149.
//
// Solidity: function lastReward() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) LastReward() (*big.Int, error) {
	return _ManagerFil.Contract.LastReward(&_ManagerFil.CallOpts)
}

// PayAll is a free data retrieval call binding the contract method 0x5806beaf.
//
// Solidity: function payAll() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) PayAll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "payAll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayAll is a free data retrieval call binding the contract method 0x5806beaf.
//
// Solidity: function payAll() view returns(uint256)
func (_ManagerFil *ManagerFilSession) PayAll() (*big.Int, error) {
	return _ManagerFil.Contract.PayAll(&_ManagerFil.CallOpts)
}

// PayAll is a free data retrieval call binding the contract method 0x5806beaf.
//
// Solidity: function payAll() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) PayAll() (*big.Int, error) {
	return _ManagerFil.Contract.PayAll(&_ManagerFil.CallOpts)
}

// RewardAll is a free data retrieval call binding the contract method 0x137be1da.
//
// Solidity: function rewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) RewardAll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "rewardAll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardAll is a free data retrieval call binding the contract method 0x137be1da.
//
// Solidity: function rewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilSession) RewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.RewardAll(&_ManagerFil.CallOpts)
}

// RewardAll is a free data retrieval call binding the contract method 0x137be1da.
//
// Solidity: function rewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) RewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.RewardAll(&_ManagerFil.CallOpts)
}

// RewardFeeBase is a free data retrieval call binding the contract method 0xff3ea2d3.
//
// Solidity: function rewardFeeBase() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) RewardFeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "rewardFeeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardFeeBase is a free data retrieval call binding the contract method 0xff3ea2d3.
//
// Solidity: function rewardFeeBase() view returns(uint256)
func (_ManagerFil *ManagerFilSession) RewardFeeBase() (*big.Int, error) {
	return _ManagerFil.Contract.RewardFeeBase(&_ManagerFil.CallOpts)
}

// RewardFeeBase is a free data retrieval call binding the contract method 0xff3ea2d3.
//
// Solidity: function rewardFeeBase() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) RewardFeeBase() (*big.Int, error) {
	return _ManagerFil.Contract.RewardFeeBase(&_ManagerFil.CallOpts)
}

// RewardFeeRate is a free data retrieval call binding the contract method 0xf5de2d1f.
//
// Solidity: function rewardFeeRate() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) RewardFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "rewardFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardFeeRate is a free data retrieval call binding the contract method 0xf5de2d1f.
//
// Solidity: function rewardFeeRate() view returns(uint256)
func (_ManagerFil *ManagerFilSession) RewardFeeRate() (*big.Int, error) {
	return _ManagerFil.Contract.RewardFeeRate(&_ManagerFil.CallOpts)
}

// RewardFeeRate is a free data retrieval call binding the contract method 0xf5de2d1f.
//
// Solidity: function rewardFeeRate() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) RewardFeeRate() (*big.Int, error) {
	return _ManagerFil.Contract.RewardFeeRate(&_ManagerFil.CallOpts)
}

// Stfil is a free data retrieval call binding the contract method 0xb6cf693f.
//
// Solidity: function stfil() view returns(address)
func (_ManagerFil *ManagerFilCaller) Stfil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "stfil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stfil is a free data retrieval call binding the contract method 0xb6cf693f.
//
// Solidity: function stfil() view returns(address)
func (_ManagerFil *ManagerFilSession) Stfil() (common.Address, error) {
	return _ManagerFil.Contract.Stfil(&_ManagerFil.CallOpts)
}

// Stfil is a free data retrieval call binding the contract method 0xb6cf693f.
//
// Solidity: function stfil() view returns(address)
func (_ManagerFil *ManagerFilCallerSession) Stfil() (common.Address, error) {
	return _ManagerFil.Contract.Stfil(&_ManagerFil.CallOpts)
}

// StfilToken is a free data retrieval call binding the contract method 0xf9aa9819.
//
// Solidity: function stfilToken() view returns(address)
func (_ManagerFil *ManagerFilCaller) StfilToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "stfilToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StfilToken is a free data retrieval call binding the contract method 0xf9aa9819.
//
// Solidity: function stfilToken() view returns(address)
func (_ManagerFil *ManagerFilSession) StfilToken() (common.Address, error) {
	return _ManagerFil.Contract.StfilToken(&_ManagerFil.CallOpts)
}

// StfilToken is a free data retrieval call binding the contract method 0xf9aa9819.
//
// Solidity: function stfilToken() view returns(address)
func (_ManagerFil *ManagerFilCallerSession) StfilToken() (common.Address, error) {
	return _ManagerFil.Contract.StfilToken(&_ManagerFil.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ManagerFil *ManagerFilCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ManagerFil *ManagerFilSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ManagerFil.Contract.SupportsInterface(&_ManagerFil.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ManagerFil *ManagerFilCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ManagerFil.Contract.SupportsInterface(&_ManagerFil.CallOpts, interfaceId)
}

// WitdhrawRewardAll is a free data retrieval call binding the contract method 0x209abf61.
//
// Solidity: function witdhrawRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCaller) WitdhrawRewardAll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "witdhrawRewardAll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WitdhrawRewardAll is a free data retrieval call binding the contract method 0x209abf61.
//
// Solidity: function witdhrawRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilSession) WitdhrawRewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.WitdhrawRewardAll(&_ManagerFil.CallOpts)
}

// WitdhrawRewardAll is a free data retrieval call binding the contract method 0x209abf61.
//
// Solidity: function witdhrawRewardAll() view returns(uint256)
func (_ManagerFil *ManagerFilCallerSession) WitdhrawRewardAll() (*big.Int, error) {
	return _ManagerFil.Contract.WitdhrawRewardAll(&_ManagerFil.CallOpts)
}

// WithdrawStfil is a free data retrieval call binding the contract method 0x52a6549f.
//
// Solidity: function withdrawStfil() view returns(address)
func (_ManagerFil *ManagerFilCaller) WithdrawStfil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManagerFil.contract.Call(opts, &out, "withdrawStfil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawStfil is a free data retrieval call binding the contract method 0x52a6549f.
//
// Solidity: function withdrawStfil() view returns(address)
func (_ManagerFil *ManagerFilSession) WithdrawStfil() (common.Address, error) {
	return _ManagerFil.Contract.WithdrawStfil(&_ManagerFil.CallOpts)
}

// WithdrawStfil is a free data retrieval call binding the contract method 0x52a6549f.
//
// Solidity: function withdrawStfil() view returns(address)
func (_ManagerFil *ManagerFilCallerSession) WithdrawStfil() (common.Address, error) {
	return _ManagerFil.Contract.WithdrawStfil(&_ManagerFil.CallOpts)
}

// FeeWithdraw is a paid mutator transaction binding the contract method 0x7f442a9a.
//
// Solidity: function feeWithdraw(uint256 amount) returns()
func (_ManagerFil *ManagerFilTransactor) FeeWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "feeWithdraw", amount)
}

// FeeWithdraw is a paid mutator transaction binding the contract method 0x7f442a9a.
//
// Solidity: function feeWithdraw(uint256 amount) returns()
func (_ManagerFil *ManagerFilSession) FeeWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _ManagerFil.Contract.FeeWithdraw(&_ManagerFil.TransactOpts, amount)
}

// FeeWithdraw is a paid mutator transaction binding the contract method 0x7f442a9a.
//
// Solidity: function feeWithdraw(uint256 amount) returns()
func (_ManagerFil *ManagerFilTransactorSession) FeeWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _ManagerFil.Contract.FeeWithdraw(&_ManagerFil.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.GrantRole(&_ManagerFil.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.GrantRole(&_ManagerFil.TransactOpts, role, account)
}

// Pay is a paid mutator transaction binding the contract method 0x1b9265b8.
//
// Solidity: function pay() payable returns()
func (_ManagerFil *ManagerFilTransactor) Pay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "pay")
}

// Pay is a paid mutator transaction binding the contract method 0x1b9265b8.
//
// Solidity: function pay() payable returns()
func (_ManagerFil *ManagerFilSession) Pay() (*types.Transaction, error) {
	return _ManagerFil.Contract.Pay(&_ManagerFil.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0x1b9265b8.
//
// Solidity: function pay() payable returns()
func (_ManagerFil *ManagerFilTransactorSession) Pay() (*types.Transaction, error) {
	return _ManagerFil.Contract.Pay(&_ManagerFil.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.RenounceRole(&_ManagerFil.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.RenounceRole(&_ManagerFil.TransactOpts, role, account)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount) returns()
func (_ManagerFil *ManagerFilTransactor) Repay(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "repay", amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount) returns()
func (_ManagerFil *ManagerFilSession) Repay(amount *big.Int) (*types.Transaction, error) {
	return _ManagerFil.Contract.Repay(&_ManagerFil.TransactOpts, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount) returns()
func (_ManagerFil *ManagerFilTransactorSession) Repay(amount *big.Int) (*types.Transaction, error) {
	return _ManagerFil.Contract.Repay(&_ManagerFil.TransactOpts, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.RevokeRole(&_ManagerFil.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ManagerFil *ManagerFilTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.RevokeRole(&_ManagerFil.TransactOpts, role, account)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_ManagerFil *ManagerFilTransactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_ManagerFil *ManagerFilSession) Reward() (*types.Transaction, error) {
	return _ManagerFil.Contract.Reward(&_ManagerFil.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_ManagerFil *ManagerFilTransactorSession) Reward() (*types.Transaction, error) {
	return _ManagerFil.Contract.Reward(&_ManagerFil.TransactOpts)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_ManagerFil *ManagerFilTransactor) SetFeeTo(opts *bind.TransactOpts, feeTo_ common.Address) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "setFeeTo", feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_ManagerFil *ManagerFilSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.SetFeeTo(&_ManagerFil.TransactOpts, feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_ManagerFil *ManagerFilTransactorSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.SetFeeTo(&_ManagerFil.TransactOpts, feeTo_)
}

// SetRewardFeeRate is a paid mutator transaction binding the contract method 0xc6116df1.
//
// Solidity: function setRewardFeeRate(uint256 rate) returns()
func (_ManagerFil *ManagerFilTransactor) SetRewardFeeRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "setRewardFeeRate", rate)
}

// SetRewardFeeRate is a paid mutator transaction binding the contract method 0xc6116df1.
//
// Solidity: function setRewardFeeRate(uint256 rate) returns()
func (_ManagerFil *ManagerFilSession) SetRewardFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _ManagerFil.Contract.SetRewardFeeRate(&_ManagerFil.TransactOpts, rate)
}

// SetRewardFeeRate is a paid mutator transaction binding the contract method 0xc6116df1.
//
// Solidity: function setRewardFeeRate(uint256 rate) returns()
func (_ManagerFil *ManagerFilTransactorSession) SetRewardFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _ManagerFil.Contract.SetRewardFeeRate(&_ManagerFil.TransactOpts, rate)
}

// SetWithdrawStfil is a paid mutator transaction binding the contract method 0xb9bfcdca.
//
// Solidity: function setWithdrawStfil(address withdrawStfil_) returns()
func (_ManagerFil *ManagerFilTransactor) SetWithdrawStfil(opts *bind.TransactOpts, withdrawStfil_ common.Address) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "setWithdrawStfil", withdrawStfil_)
}

// SetWithdrawStfil is a paid mutator transaction binding the contract method 0xb9bfcdca.
//
// Solidity: function setWithdrawStfil(address withdrawStfil_) returns()
func (_ManagerFil *ManagerFilSession) SetWithdrawStfil(withdrawStfil_ common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.SetWithdrawStfil(&_ManagerFil.TransactOpts, withdrawStfil_)
}

// SetWithdrawStfil is a paid mutator transaction binding the contract method 0xb9bfcdca.
//
// Solidity: function setWithdrawStfil(address withdrawStfil_) returns()
func (_ManagerFil *ManagerFilTransactorSession) SetWithdrawStfil(withdrawStfil_ common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.SetWithdrawStfil(&_ManagerFil.TransactOpts, withdrawStfil_)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0x6a2c5352.
//
// Solidity: function withdrawReward(uint256 amount, address to) returns()
func (_ManagerFil *ManagerFilTransactor) WithdrawReward(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ManagerFil.contract.Transact(opts, "withdrawReward", amount, to)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0x6a2c5352.
//
// Solidity: function withdrawReward(uint256 amount, address to) returns()
func (_ManagerFil *ManagerFilSession) WithdrawReward(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.WithdrawReward(&_ManagerFil.TransactOpts, amount, to)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0x6a2c5352.
//
// Solidity: function withdrawReward(uint256 amount, address to) returns()
func (_ManagerFil *ManagerFilTransactorSession) WithdrawReward(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ManagerFil.Contract.WithdrawReward(&_ManagerFil.TransactOpts, amount, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ManagerFil *ManagerFilTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagerFil.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ManagerFil *ManagerFilSession) Receive() (*types.Transaction, error) {
	return _ManagerFil.Contract.Receive(&_ManagerFil.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ManagerFil *ManagerFilTransactorSession) Receive() (*types.Transaction, error) {
	return _ManagerFil.Contract.Receive(&_ManagerFil.TransactOpts)
}

// ManagerFilRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ManagerFil contract.
type ManagerFilRoleAdminChangedIterator struct {
	Event *ManagerFilRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ManagerFilRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilRoleAdminChanged)
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
		it.Event = new(ManagerFilRoleAdminChanged)
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
func (it *ManagerFilRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilRoleAdminChanged represents a RoleAdminChanged event raised by the ManagerFil contract.
type ManagerFilRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ManagerFil *ManagerFilFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ManagerFilRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ManagerFilRoleAdminChangedIterator{contract: _ManagerFil.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ManagerFil *ManagerFilFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ManagerFilRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilRoleAdminChanged)
				if err := _ManagerFil.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ManagerFil *ManagerFilFilterer) ParseRoleAdminChanged(log types.Log) (*ManagerFilRoleAdminChanged, error) {
	event := new(ManagerFilRoleAdminChanged)
	if err := _ManagerFil.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ManagerFil contract.
type ManagerFilRoleGrantedIterator struct {
	Event *ManagerFilRoleGranted // Event containing the contract specifics and raw log

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
func (it *ManagerFilRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilRoleGranted)
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
		it.Event = new(ManagerFilRoleGranted)
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
func (it *ManagerFilRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilRoleGranted represents a RoleGranted event raised by the ManagerFil contract.
type ManagerFilRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ManagerFil *ManagerFilFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ManagerFilRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ManagerFilRoleGrantedIterator{contract: _ManagerFil.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ManagerFil *ManagerFilFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ManagerFilRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilRoleGranted)
				if err := _ManagerFil.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ManagerFil *ManagerFilFilterer) ParseRoleGranted(log types.Log) (*ManagerFilRoleGranted, error) {
	event := new(ManagerFilRoleGranted)
	if err := _ManagerFil.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ManagerFil contract.
type ManagerFilRoleRevokedIterator struct {
	Event *ManagerFilRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ManagerFilRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilRoleRevoked)
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
		it.Event = new(ManagerFilRoleRevoked)
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
func (it *ManagerFilRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilRoleRevoked represents a RoleRevoked event raised by the ManagerFil contract.
type ManagerFilRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ManagerFil *ManagerFilFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ManagerFilRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ManagerFilRoleRevokedIterator{contract: _ManagerFil.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ManagerFil *ManagerFilFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ManagerFilRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilRoleRevoked)
				if err := _ManagerFil.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ManagerFil *ManagerFilFilterer) ParseRoleRevoked(log types.Log) (*ManagerFilRoleRevoked, error) {
	event := new(ManagerFilRoleRevoked)
	if err := _ManagerFil.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilFeeWithdrawRewardCompletedIterator is returned from FilterFeeWithdrawRewardCompleted and is used to iterate over the raw logs and unpacked data for FeeWithdrawRewardCompleted events raised by the ManagerFil contract.
type ManagerFilFeeWithdrawRewardCompletedIterator struct {
	Event *ManagerFilFeeWithdrawRewardCompleted // Event containing the contract specifics and raw log

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
func (it *ManagerFilFeeWithdrawRewardCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilFeeWithdrawRewardCompleted)
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
		it.Event = new(ManagerFilFeeWithdrawRewardCompleted)
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
func (it *ManagerFilFeeWithdrawRewardCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilFeeWithdrawRewardCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilFeeWithdrawRewardCompleted represents a FeeWithdrawRewardCompleted event raised by the ManagerFil contract.
type ManagerFilFeeWithdrawRewardCompleted struct {
	To                   common.Address
	FeeRewardAll         *big.Int
	FeeWitdhrawRewardAll *big.Int
	Amount               *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawRewardCompleted is a free log retrieval operation binding the contract event 0x812069824a914db65c2ec03d5d691d9b6599ca1092d266c81c28241dcc83da40.
//
// Solidity: event feeWithdrawRewardCompleted(address indexed to, uint256 feeRewardAll, uint256 feeWitdhrawRewardAll, uint256 amount)
func (_ManagerFil *ManagerFilFilterer) FilterFeeWithdrawRewardCompleted(opts *bind.FilterOpts, to []common.Address) (*ManagerFilFeeWithdrawRewardCompletedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "feeWithdrawRewardCompleted", toRule)
	if err != nil {
		return nil, err
	}
	return &ManagerFilFeeWithdrawRewardCompletedIterator{contract: _ManagerFil.contract, event: "feeWithdrawRewardCompleted", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawRewardCompleted is a free log subscription operation binding the contract event 0x812069824a914db65c2ec03d5d691d9b6599ca1092d266c81c28241dcc83da40.
//
// Solidity: event feeWithdrawRewardCompleted(address indexed to, uint256 feeRewardAll, uint256 feeWitdhrawRewardAll, uint256 amount)
func (_ManagerFil *ManagerFilFilterer) WatchFeeWithdrawRewardCompleted(opts *bind.WatchOpts, sink chan<- *ManagerFilFeeWithdrawRewardCompleted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "feeWithdrawRewardCompleted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilFeeWithdrawRewardCompleted)
				if err := _ManagerFil.contract.UnpackLog(event, "feeWithdrawRewardCompleted", log); err != nil {
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

// ParseFeeWithdrawRewardCompleted is a log parse operation binding the contract event 0x812069824a914db65c2ec03d5d691d9b6599ca1092d266c81c28241dcc83da40.
//
// Solidity: event feeWithdrawRewardCompleted(address indexed to, uint256 feeRewardAll, uint256 feeWitdhrawRewardAll, uint256 amount)
func (_ManagerFil *ManagerFilFilterer) ParseFeeWithdrawRewardCompleted(log types.Log) (*ManagerFilFeeWithdrawRewardCompleted, error) {
	event := new(ManagerFilFeeWithdrawRewardCompleted)
	if err := _ManagerFil.contract.UnpackLog(event, "feeWithdrawRewardCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilPayCompletedIterator is returned from FilterPayCompleted and is used to iterate over the raw logs and unpacked data for PayCompleted events raised by the ManagerFil contract.
type ManagerFilPayCompletedIterator struct {
	Event *ManagerFilPayCompleted // Event containing the contract specifics and raw log

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
func (it *ManagerFilPayCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilPayCompleted)
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
		it.Event = new(ManagerFilPayCompleted)
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
func (it *ManagerFilPayCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilPayCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilPayCompleted represents a PayCompleted event raised by the ManagerFil contract.
type ManagerFilPayCompleted struct {
	Amount *big.Int
	PayAll *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPayCompleted is a free log retrieval operation binding the contract event 0x7596196094f2132095222e5ee83c6b057a862cd7c973ce5dd6136d53d4060c78.
//
// Solidity: event payCompleted(uint256 amount, uint256 payAll)
func (_ManagerFil *ManagerFilFilterer) FilterPayCompleted(opts *bind.FilterOpts) (*ManagerFilPayCompletedIterator, error) {

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "payCompleted")
	if err != nil {
		return nil, err
	}
	return &ManagerFilPayCompletedIterator{contract: _ManagerFil.contract, event: "payCompleted", logs: logs, sub: sub}, nil
}

// WatchPayCompleted is a free log subscription operation binding the contract event 0x7596196094f2132095222e5ee83c6b057a862cd7c973ce5dd6136d53d4060c78.
//
// Solidity: event payCompleted(uint256 amount, uint256 payAll)
func (_ManagerFil *ManagerFilFilterer) WatchPayCompleted(opts *bind.WatchOpts, sink chan<- *ManagerFilPayCompleted) (event.Subscription, error) {

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "payCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilPayCompleted)
				if err := _ManagerFil.contract.UnpackLog(event, "payCompleted", log); err != nil {
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

// ParsePayCompleted is a log parse operation binding the contract event 0x7596196094f2132095222e5ee83c6b057a862cd7c973ce5dd6136d53d4060c78.
//
// Solidity: event payCompleted(uint256 amount, uint256 payAll)
func (_ManagerFil *ManagerFilFilterer) ParsePayCompleted(log types.Log) (*ManagerFilPayCompleted, error) {
	event := new(ManagerFilPayCompleted)
	if err := _ManagerFil.contract.UnpackLog(event, "payCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilRePayCompletedIterator is returned from FilterRePayCompleted and is used to iterate over the raw logs and unpacked data for RePayCompleted events raised by the ManagerFil contract.
type ManagerFilRePayCompletedIterator struct {
	Event *ManagerFilRePayCompleted // Event containing the contract specifics and raw log

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
func (it *ManagerFilRePayCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilRePayCompleted)
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
		it.Event = new(ManagerFilRePayCompleted)
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
func (it *ManagerFilRePayCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilRePayCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilRePayCompleted represents a RePayCompleted event raised by the ManagerFil contract.
type ManagerFilRePayCompleted struct {
	Amount *big.Int
	PayAll *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRePayCompleted is a free log retrieval operation binding the contract event 0xa4c8035d7e43f43cd335381f53cb6daceb9b22aa807e272cfb7019aee186b7e7.
//
// Solidity: event rePayCompleted(uint256 amount, uint256 payAll)
func (_ManagerFil *ManagerFilFilterer) FilterRePayCompleted(opts *bind.FilterOpts) (*ManagerFilRePayCompletedIterator, error) {

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "rePayCompleted")
	if err != nil {
		return nil, err
	}
	return &ManagerFilRePayCompletedIterator{contract: _ManagerFil.contract, event: "rePayCompleted", logs: logs, sub: sub}, nil
}

// WatchRePayCompleted is a free log subscription operation binding the contract event 0xa4c8035d7e43f43cd335381f53cb6daceb9b22aa807e272cfb7019aee186b7e7.
//
// Solidity: event rePayCompleted(uint256 amount, uint256 payAll)
func (_ManagerFil *ManagerFilFilterer) WatchRePayCompleted(opts *bind.WatchOpts, sink chan<- *ManagerFilRePayCompleted) (event.Subscription, error) {

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "rePayCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilRePayCompleted)
				if err := _ManagerFil.contract.UnpackLog(event, "rePayCompleted", log); err != nil {
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

// ParseRePayCompleted is a log parse operation binding the contract event 0xa4c8035d7e43f43cd335381f53cb6daceb9b22aa807e272cfb7019aee186b7e7.
//
// Solidity: event rePayCompleted(uint256 amount, uint256 payAll)
func (_ManagerFil *ManagerFilFilterer) ParseRePayCompleted(log types.Log) (*ManagerFilRePayCompleted, error) {
	event := new(ManagerFilRePayCompleted)
	if err := _ManagerFil.contract.UnpackLog(event, "rePayCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilRewardCompletedIterator is returned from FilterRewardCompleted and is used to iterate over the raw logs and unpacked data for RewardCompleted events raised by the ManagerFil contract.
type ManagerFilRewardCompletedIterator struct {
	Event *ManagerFilRewardCompleted // Event containing the contract specifics and raw log

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
func (it *ManagerFilRewardCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilRewardCompleted)
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
		it.Event = new(ManagerFilRewardCompleted)
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
func (it *ManagerFilRewardCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilRewardCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilRewardCompleted represents a RewardCompleted event raised by the ManagerFil contract.
type ManagerFilRewardCompleted struct {
	AllReward     *big.Int
	FeeRewardAll  *big.Int
	RewardAll     *big.Int
	CurrentReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardCompleted is a free log retrieval operation binding the contract event 0xb418edc1aa87ff5716cdb3f819b7844c28a26d6c67b08761937822989a379ace.
//
// Solidity: event rewardCompleted(uint256 allReward, uint256 feeRewardAll, uint256 rewardAll, uint256 currentReward)
func (_ManagerFil *ManagerFilFilterer) FilterRewardCompleted(opts *bind.FilterOpts) (*ManagerFilRewardCompletedIterator, error) {

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "rewardCompleted")
	if err != nil {
		return nil, err
	}
	return &ManagerFilRewardCompletedIterator{contract: _ManagerFil.contract, event: "rewardCompleted", logs: logs, sub: sub}, nil
}

// WatchRewardCompleted is a free log subscription operation binding the contract event 0xb418edc1aa87ff5716cdb3f819b7844c28a26d6c67b08761937822989a379ace.
//
// Solidity: event rewardCompleted(uint256 allReward, uint256 feeRewardAll, uint256 rewardAll, uint256 currentReward)
func (_ManagerFil *ManagerFilFilterer) WatchRewardCompleted(opts *bind.WatchOpts, sink chan<- *ManagerFilRewardCompleted) (event.Subscription, error) {

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "rewardCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilRewardCompleted)
				if err := _ManagerFil.contract.UnpackLog(event, "rewardCompleted", log); err != nil {
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

// ParseRewardCompleted is a log parse operation binding the contract event 0xb418edc1aa87ff5716cdb3f819b7844c28a26d6c67b08761937822989a379ace.
//
// Solidity: event rewardCompleted(uint256 allReward, uint256 feeRewardAll, uint256 rewardAll, uint256 currentReward)
func (_ManagerFil *ManagerFilFilterer) ParseRewardCompleted(log types.Log) (*ManagerFilRewardCompleted, error) {
	event := new(ManagerFilRewardCompleted)
	if err := _ManagerFil.contract.UnpackLog(event, "rewardCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerFilWithdrawRewardCompletedIterator is returned from FilterWithdrawRewardCompleted and is used to iterate over the raw logs and unpacked data for WithdrawRewardCompleted events raised by the ManagerFil contract.
type ManagerFilWithdrawRewardCompletedIterator struct {
	Event *ManagerFilWithdrawRewardCompleted // Event containing the contract specifics and raw log

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
func (it *ManagerFilWithdrawRewardCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerFilWithdrawRewardCompleted)
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
		it.Event = new(ManagerFilWithdrawRewardCompleted)
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
func (it *ManagerFilWithdrawRewardCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerFilWithdrawRewardCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerFilWithdrawRewardCompleted represents a WithdrawRewardCompleted event raised by the ManagerFil contract.
type ManagerFilWithdrawRewardCompleted struct {
	To                common.Address
	RewardAll         *big.Int
	WitdhrawRewardAll *big.Int
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRewardCompleted is a free log retrieval operation binding the contract event 0xa66eeac44f5bd3756ae069a9e27fe4dba993ff028d19f564448baa9b684a9644.
//
// Solidity: event withdrawRewardCompleted(address indexed to, uint256 rewardAll, uint256 witdhrawRewardAll, uint256 amount)
func (_ManagerFil *ManagerFilFilterer) FilterWithdrawRewardCompleted(opts *bind.FilterOpts, to []common.Address) (*ManagerFilWithdrawRewardCompletedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ManagerFil.contract.FilterLogs(opts, "withdrawRewardCompleted", toRule)
	if err != nil {
		return nil, err
	}
	return &ManagerFilWithdrawRewardCompletedIterator{contract: _ManagerFil.contract, event: "withdrawRewardCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawRewardCompleted is a free log subscription operation binding the contract event 0xa66eeac44f5bd3756ae069a9e27fe4dba993ff028d19f564448baa9b684a9644.
//
// Solidity: event withdrawRewardCompleted(address indexed to, uint256 rewardAll, uint256 witdhrawRewardAll, uint256 amount)
func (_ManagerFil *ManagerFilFilterer) WatchWithdrawRewardCompleted(opts *bind.WatchOpts, sink chan<- *ManagerFilWithdrawRewardCompleted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ManagerFil.contract.WatchLogs(opts, "withdrawRewardCompleted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerFilWithdrawRewardCompleted)
				if err := _ManagerFil.contract.UnpackLog(event, "withdrawRewardCompleted", log); err != nil {
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

// ParseWithdrawRewardCompleted is a log parse operation binding the contract event 0xa66eeac44f5bd3756ae069a9e27fe4dba993ff028d19f564448baa9b684a9644.
//
// Solidity: event withdrawRewardCompleted(address indexed to, uint256 rewardAll, uint256 witdhrawRewardAll, uint256 amount)
func (_ManagerFil *ManagerFilFilterer) ParseWithdrawRewardCompleted(log types.Log) (*ManagerFilWithdrawRewardCompleted, error) {
	event := new(ManagerFilWithdrawRewardCompleted)
	if err := _ManagerFil.contract.UnpackLog(event, "withdrawRewardCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
