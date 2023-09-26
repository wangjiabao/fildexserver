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

// TokenExchangeMetaData contains all meta data concerning the TokenExchange contract.
var TokenExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dfil_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeTo_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defaultAdminRoleLimit_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AccountRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usePlatToken\",\"type\":\"uint256\"}],\"name\":\"DFIL2FILExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FIL2DFILExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FILINComplated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitRateAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"union\",\"type\":\"bool\"}],\"name\":\"FILLOCKComplated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ManagerWithdrawForFilComplated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"REWARDFIL2DFILExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usePlatToken\",\"type\":\"uint256\"}],\"name\":\"WITHDRAWFILComplated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"feeToWithdrawComplated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usePlatToken\",\"type\":\"uint256\"}],\"name\":\"DFIL2FIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FIL2DFIL\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FILIN\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitRateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"union\",\"type\":\"bool\"}],\"name\":\"FILLOCK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GRANT_TOKEN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"REWARDFIL2DFIL\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usePlatToken\",\"type\":\"uint256\"}],\"name\":\"UNIONWITHDRAWFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountUnionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowAccountUnionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"closeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminRoleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dfil\",\"outputs\":[{\"internalType\":\"contractIDFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeEnableNoLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"existsAccountUnion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractITokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountUnion\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowAccountUnion\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAllowAccountUnionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowAccountUnionAmountTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getAllowAccountUnionAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowAccountUnionLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeEnableNoLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getTokenOwnerDfilBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getTokenOwnerFilBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levelHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levelLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFilLimitAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerLimitAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownerFilBalanceChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platToken\",\"outputs\":[{\"internalType\":\"contractIPlatToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAccountUnion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userDefineRate\",\"type\":\"uint256\"}],\"name\":\"setAccountUnion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAllUserExchangeEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory_\",\"type\":\"address\"}],\"name\":\"setFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"feeTo_\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setManagerGetFilLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"levelLow_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"levelHigh_\",\"type\":\"uint256\"}],\"name\":\"setManagerLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setManagerLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"platToken_\",\"type\":\"address\"}],\"name\":\"setPlatToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setTokenRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unionLowRate_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unionRate_\",\"type\":\"uint256\"}],\"name\":\"setUnionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"useManager_\",\"type\":\"bool\"}],\"name\":\"setUseManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenOwnerDfilBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenOwnerFilBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenOwnerFilIn\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenOwnerFilOut\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unionBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unionLowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TokenExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenExchangeMetaData.ABI instead.
var TokenExchangeABI = TokenExchangeMetaData.ABI

// TokenExchange is an auto generated Go binding around an Ethereum contract.
type TokenExchange struct {
	TokenExchangeCaller     // Read-only binding to the contract
	TokenExchangeTransactor // Write-only binding to the contract
	TokenExchangeFilterer   // Log filterer for contract events
}

// TokenExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenExchangeSession struct {
	Contract     *TokenExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenExchangeCallerSession struct {
	Contract *TokenExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenExchangeTransactorSession struct {
	Contract     *TokenExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenExchangeRaw struct {
	Contract *TokenExchange // Generic contract binding to access the raw methods on
}

// TokenExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenExchangeCallerRaw struct {
	Contract *TokenExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// TokenExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenExchangeTransactorRaw struct {
	Contract *TokenExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenExchange creates a new instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchange(address common.Address, backend bind.ContractBackend) (*TokenExchange, error) {
	contract, err := bindTokenExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenExchange{TokenExchangeCaller: TokenExchangeCaller{contract: contract}, TokenExchangeTransactor: TokenExchangeTransactor{contract: contract}, TokenExchangeFilterer: TokenExchangeFilterer{contract: contract}}, nil
}

// NewTokenExchangeCaller creates a new read-only instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchangeCaller(address common.Address, caller bind.ContractCaller) (*TokenExchangeCaller, error) {
	contract, err := bindTokenExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeCaller{contract: contract}, nil
}

// NewTokenExchangeTransactor creates a new write-only instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenExchangeTransactor, error) {
	contract, err := bindTokenExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeTransactor{contract: contract}, nil
}

// NewTokenExchangeFilterer creates a new log filterer instance of TokenExchange, bound to a specific deployed contract.
func NewTokenExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenExchangeFilterer, error) {
	contract, err := bindTokenExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeFilterer{contract: contract}, nil
}

// bindTokenExchange binds a generic wrapper to an already deployed contract.
func bindTokenExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenExchange *TokenExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenExchange.Contract.TokenExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenExchange *TokenExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.Contract.TokenExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenExchange *TokenExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenExchange.Contract.TokenExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenExchange *TokenExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenExchange *TokenExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenExchange *TokenExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenExchange.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenExchange.Contract.DEFAULTADMINROLE(&_TokenExchange.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenExchange.Contract.DEFAULTADMINROLE(&_TokenExchange.CallOpts)
}

// GRANTTOKENROLE is a free data retrieval call binding the contract method 0x453b1816.
//
// Solidity: function GRANT_TOKEN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCaller) GRANTTOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "GRANT_TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GRANTTOKENROLE is a free data retrieval call binding the contract method 0x453b1816.
//
// Solidity: function GRANT_TOKEN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeSession) GRANTTOKENROLE() ([32]byte, error) {
	return _TokenExchange.Contract.GRANTTOKENROLE(&_TokenExchange.CallOpts)
}

// GRANTTOKENROLE is a free data retrieval call binding the contract method 0x453b1816.
//
// Solidity: function GRANT_TOKEN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCallerSession) GRANTTOKENROLE() ([32]byte, error) {
	return _TokenExchange.Contract.GRANTTOKENROLE(&_TokenExchange.CallOpts)
}

// SETMANAGERROLE is a free data retrieval call binding the contract method 0xbbadb957.
//
// Solidity: function SET_MANAGER_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCaller) SETMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "SET_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETMANAGERROLE is a free data retrieval call binding the contract method 0xbbadb957.
//
// Solidity: function SET_MANAGER_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeSession) SETMANAGERROLE() ([32]byte, error) {
	return _TokenExchange.Contract.SETMANAGERROLE(&_TokenExchange.CallOpts)
}

// SETMANAGERROLE is a free data retrieval call binding the contract method 0xbbadb957.
//
// Solidity: function SET_MANAGER_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCallerSession) SETMANAGERROLE() ([32]byte, error) {
	return _TokenExchange.Contract.SETMANAGERROLE(&_TokenExchange.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCaller) SUPERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "SUPER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeSession) SUPERADMINROLE() ([32]byte, error) {
	return _TokenExchange.Contract.SUPERADMINROLE(&_TokenExchange.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCallerSession) SUPERADMINROLE() ([32]byte, error) {
	return _TokenExchange.Contract.SUPERADMINROLE(&_TokenExchange.CallOpts)
}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCaller) TOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeSession) TOKENROLE() ([32]byte, error) {
	return _TokenExchange.Contract.TOKENROLE(&_TokenExchange.CallOpts)
}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_TokenExchange *TokenExchangeCallerSession) TOKENROLE() ([32]byte, error) {
	return _TokenExchange.Contract.TOKENROLE(&_TokenExchange.CallOpts)
}

// AccountUnionRate is a free data retrieval call binding the contract method 0x5f0fda9e.
//
// Solidity: function accountUnionRate(address ) view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) AccountUnionRate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "accountUnionRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountUnionRate is a free data retrieval call binding the contract method 0x5f0fda9e.
//
// Solidity: function accountUnionRate(address ) view returns(uint256)
func (_TokenExchange *TokenExchangeSession) AccountUnionRate(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.AccountUnionRate(&_TokenExchange.CallOpts, arg0)
}

// AccountUnionRate is a free data retrieval call binding the contract method 0x5f0fda9e.
//
// Solidity: function accountUnionRate(address ) view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) AccountUnionRate(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.AccountUnionRate(&_TokenExchange.CallOpts, arg0)
}

// AllowAccountUnionAmount is a free data retrieval call binding the contract method 0x86bd4ca9.
//
// Solidity: function allowAccountUnionAmount(address ) view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) AllowAccountUnionAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "allowAccountUnionAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowAccountUnionAmount is a free data retrieval call binding the contract method 0x86bd4ca9.
//
// Solidity: function allowAccountUnionAmount(address ) view returns(uint256)
func (_TokenExchange *TokenExchangeSession) AllowAccountUnionAmount(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.AllowAccountUnionAmount(&_TokenExchange.CallOpts, arg0)
}

// AllowAccountUnionAmount is a free data retrieval call binding the contract method 0x86bd4ca9.
//
// Solidity: function allowAccountUnionAmount(address ) view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) AllowAccountUnionAmount(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.AllowAccountUnionAmount(&_TokenExchange.CallOpts, arg0)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) Base() (*big.Int, error) {
	return _TokenExchange.Contract.Base(&_TokenExchange.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) Base() (*big.Int, error) {
	return _TokenExchange.Contract.Base(&_TokenExchange.CallOpts)
}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) DefaultAdminRoleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "defaultAdminRoleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) DefaultAdminRoleLimit() (*big.Int, error) {
	return _TokenExchange.Contract.DefaultAdminRoleLimit(&_TokenExchange.CallOpts)
}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) DefaultAdminRoleLimit() (*big.Int, error) {
	return _TokenExchange.Contract.DefaultAdminRoleLimit(&_TokenExchange.CallOpts)
}

// Dfil is a free data retrieval call binding the contract method 0x7baf0d5c.
//
// Solidity: function dfil() view returns(address)
func (_TokenExchange *TokenExchangeCaller) Dfil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "dfil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dfil is a free data retrieval call binding the contract method 0x7baf0d5c.
//
// Solidity: function dfil() view returns(address)
func (_TokenExchange *TokenExchangeSession) Dfil() (common.Address, error) {
	return _TokenExchange.Contract.Dfil(&_TokenExchange.CallOpts)
}

// Dfil is a free data retrieval call binding the contract method 0x7baf0d5c.
//
// Solidity: function dfil() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) Dfil() (common.Address, error) {
	return _TokenExchange.Contract.Dfil(&_TokenExchange.CallOpts)
}

// ExchangeEnableNoLimit is a free data retrieval call binding the contract method 0x93adb7a2.
//
// Solidity: function exchangeEnableNoLimit() view returns(bool)
func (_TokenExchange *TokenExchangeCaller) ExchangeEnableNoLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "exchangeEnableNoLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExchangeEnableNoLimit is a free data retrieval call binding the contract method 0x93adb7a2.
//
// Solidity: function exchangeEnableNoLimit() view returns(bool)
func (_TokenExchange *TokenExchangeSession) ExchangeEnableNoLimit() (bool, error) {
	return _TokenExchange.Contract.ExchangeEnableNoLimit(&_TokenExchange.CallOpts)
}

// ExchangeEnableNoLimit is a free data retrieval call binding the contract method 0x93adb7a2.
//
// Solidity: function exchangeEnableNoLimit() view returns(bool)
func (_TokenExchange *TokenExchangeCallerSession) ExchangeEnableNoLimit() (bool, error) {
	return _TokenExchange.Contract.ExchangeEnableNoLimit(&_TokenExchange.CallOpts)
}

// ExistsAccountUnion is a free data retrieval call binding the contract method 0xc7ca4660.
//
// Solidity: function existsAccountUnion() view returns(bool)
func (_TokenExchange *TokenExchangeCaller) ExistsAccountUnion(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "existsAccountUnion")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsAccountUnion is a free data retrieval call binding the contract method 0xc7ca4660.
//
// Solidity: function existsAccountUnion() view returns(bool)
func (_TokenExchange *TokenExchangeSession) ExistsAccountUnion() (bool, error) {
	return _TokenExchange.Contract.ExistsAccountUnion(&_TokenExchange.CallOpts)
}

// ExistsAccountUnion is a free data retrieval call binding the contract method 0xc7ca4660.
//
// Solidity: function existsAccountUnion() view returns(bool)
func (_TokenExchange *TokenExchangeCallerSession) ExistsAccountUnion() (bool, error) {
	return _TokenExchange.Contract.ExistsAccountUnion(&_TokenExchange.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_TokenExchange *TokenExchangeCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_TokenExchange *TokenExchangeSession) Factory() (common.Address, error) {
	return _TokenExchange.Contract.Factory(&_TokenExchange.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) Factory() (common.Address, error) {
	return _TokenExchange.Contract.Factory(&_TokenExchange.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TokenExchange *TokenExchangeCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TokenExchange *TokenExchangeSession) FeeTo() (common.Address, error) {
	return _TokenExchange.Contract.FeeTo(&_TokenExchange.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) FeeTo() (common.Address, error) {
	return _TokenExchange.Contract.FeeTo(&_TokenExchange.CallOpts)
}

// GetAccountUnion is a free data retrieval call binding the contract method 0xd9b21c91.
//
// Solidity: function getAccountUnion() view returns(address[])
func (_TokenExchange *TokenExchangeCaller) GetAccountUnion(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getAccountUnion")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAccountUnion is a free data retrieval call binding the contract method 0xd9b21c91.
//
// Solidity: function getAccountUnion() view returns(address[])
func (_TokenExchange *TokenExchangeSession) GetAccountUnion() ([]common.Address, error) {
	return _TokenExchange.Contract.GetAccountUnion(&_TokenExchange.CallOpts)
}

// GetAccountUnion is a free data retrieval call binding the contract method 0xd9b21c91.
//
// Solidity: function getAccountUnion() view returns(address[])
func (_TokenExchange *TokenExchangeCallerSession) GetAccountUnion() ([]common.Address, error) {
	return _TokenExchange.Contract.GetAccountUnion(&_TokenExchange.CallOpts)
}

// GetAllowAccountUnion is a free data retrieval call binding the contract method 0x9970dedd.
//
// Solidity: function getAllowAccountUnion() view returns(address[])
func (_TokenExchange *TokenExchangeCaller) GetAllowAccountUnion(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getAllowAccountUnion")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowAccountUnion is a free data retrieval call binding the contract method 0x9970dedd.
//
// Solidity: function getAllowAccountUnion() view returns(address[])
func (_TokenExchange *TokenExchangeSession) GetAllowAccountUnion() ([]common.Address, error) {
	return _TokenExchange.Contract.GetAllowAccountUnion(&_TokenExchange.CallOpts)
}

// GetAllowAccountUnion is a free data retrieval call binding the contract method 0x9970dedd.
//
// Solidity: function getAllowAccountUnion() view returns(address[])
func (_TokenExchange *TokenExchangeCallerSession) GetAllowAccountUnion() ([]common.Address, error) {
	return _TokenExchange.Contract.GetAllowAccountUnion(&_TokenExchange.CallOpts)
}

// GetAllowAccountUnionAmount is a free data retrieval call binding the contract method 0x2bbb9abe.
//
// Solidity: function getAllowAccountUnionAmount(address account) view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) GetAllowAccountUnionAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getAllowAccountUnionAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowAccountUnionAmount is a free data retrieval call binding the contract method 0x2bbb9abe.
//
// Solidity: function getAllowAccountUnionAmount(address account) view returns(uint256)
func (_TokenExchange *TokenExchangeSession) GetAllowAccountUnionAmount(account common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionAmount(&_TokenExchange.CallOpts, account)
}

// GetAllowAccountUnionAmount is a free data retrieval call binding the contract method 0x2bbb9abe.
//
// Solidity: function getAllowAccountUnionAmount(address account) view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) GetAllowAccountUnionAmount(account common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionAmount(&_TokenExchange.CallOpts, account)
}

// GetAllowAccountUnionAmountTotal is a free data retrieval call binding the contract method 0xfe4ae7bf.
//
// Solidity: function getAllowAccountUnionAmountTotal() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) GetAllowAccountUnionAmountTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getAllowAccountUnionAmountTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowAccountUnionAmountTotal is a free data retrieval call binding the contract method 0xfe4ae7bf.
//
// Solidity: function getAllowAccountUnionAmountTotal() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) GetAllowAccountUnionAmountTotal() (*big.Int, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionAmountTotal(&_TokenExchange.CallOpts)
}

// GetAllowAccountUnionAmountTotal is a free data retrieval call binding the contract method 0xfe4ae7bf.
//
// Solidity: function getAllowAccountUnionAmountTotal() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) GetAllowAccountUnionAmountTotal() (*big.Int, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionAmountTotal(&_TokenExchange.CallOpts)
}

// GetAllowAccountUnionAt is a free data retrieval call binding the contract method 0x8d1b9653.
//
// Solidity: function getAllowAccountUnionAt(uint256 i) view returns(address)
func (_TokenExchange *TokenExchangeCaller) GetAllowAccountUnionAt(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getAllowAccountUnionAt", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAllowAccountUnionAt is a free data retrieval call binding the contract method 0x8d1b9653.
//
// Solidity: function getAllowAccountUnionAt(uint256 i) view returns(address)
func (_TokenExchange *TokenExchangeSession) GetAllowAccountUnionAt(i *big.Int) (common.Address, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionAt(&_TokenExchange.CallOpts, i)
}

// GetAllowAccountUnionAt is a free data retrieval call binding the contract method 0x8d1b9653.
//
// Solidity: function getAllowAccountUnionAt(uint256 i) view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) GetAllowAccountUnionAt(i *big.Int) (common.Address, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionAt(&_TokenExchange.CallOpts, i)
}

// GetAllowAccountUnionLength is a free data retrieval call binding the contract method 0x45b8c1dd.
//
// Solidity: function getAllowAccountUnionLength() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) GetAllowAccountUnionLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getAllowAccountUnionLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowAccountUnionLength is a free data retrieval call binding the contract method 0x45b8c1dd.
//
// Solidity: function getAllowAccountUnionLength() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) GetAllowAccountUnionLength() (*big.Int, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionLength(&_TokenExchange.CallOpts)
}

// GetAllowAccountUnionLength is a free data retrieval call binding the contract method 0x45b8c1dd.
//
// Solidity: function getAllowAccountUnionLength() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) GetAllowAccountUnionLength() (*big.Int, error) {
	return _TokenExchange.Contract.GetAllowAccountUnionLength(&_TokenExchange.CallOpts)
}

// GetExchangeEnableNoLimit is a free data retrieval call binding the contract method 0xf3ff7294.
//
// Solidity: function getExchangeEnableNoLimit() view returns(bool)
func (_TokenExchange *TokenExchangeCaller) GetExchangeEnableNoLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getExchangeEnableNoLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetExchangeEnableNoLimit is a free data retrieval call binding the contract method 0xf3ff7294.
//
// Solidity: function getExchangeEnableNoLimit() view returns(bool)
func (_TokenExchange *TokenExchangeSession) GetExchangeEnableNoLimit() (bool, error) {
	return _TokenExchange.Contract.GetExchangeEnableNoLimit(&_TokenExchange.CallOpts)
}

// GetExchangeEnableNoLimit is a free data retrieval call binding the contract method 0xf3ff7294.
//
// Solidity: function getExchangeEnableNoLimit() view returns(bool)
func (_TokenExchange *TokenExchangeCallerSession) GetExchangeEnableNoLimit() (bool, error) {
	return _TokenExchange.Contract.GetExchangeEnableNoLimit(&_TokenExchange.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenExchange *TokenExchangeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenExchange *TokenExchangeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenExchange.Contract.GetRoleAdmin(&_TokenExchange.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenExchange *TokenExchangeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenExchange.Contract.GetRoleAdmin(&_TokenExchange.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenExchange *TokenExchangeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenExchange *TokenExchangeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenExchange.Contract.GetRoleMember(&_TokenExchange.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenExchange.Contract.GetRoleMember(&_TokenExchange.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenExchange *TokenExchangeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenExchange.Contract.GetRoleMemberCount(&_TokenExchange.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenExchange.Contract.GetRoleMemberCount(&_TokenExchange.CallOpts, role)
}

// GetTokenOwnerDfilBalance is a free data retrieval call binding the contract method 0x344c10e1.
//
// Solidity: function getTokenOwnerDfilBalance(address owner) view returns(int256)
func (_TokenExchange *TokenExchangeCaller) GetTokenOwnerDfilBalance(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getTokenOwnerDfilBalance", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenOwnerDfilBalance is a free data retrieval call binding the contract method 0x344c10e1.
//
// Solidity: function getTokenOwnerDfilBalance(address owner) view returns(int256)
func (_TokenExchange *TokenExchangeSession) GetTokenOwnerDfilBalance(owner common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.GetTokenOwnerDfilBalance(&_TokenExchange.CallOpts, owner)
}

// GetTokenOwnerDfilBalance is a free data retrieval call binding the contract method 0x344c10e1.
//
// Solidity: function getTokenOwnerDfilBalance(address owner) view returns(int256)
func (_TokenExchange *TokenExchangeCallerSession) GetTokenOwnerDfilBalance(owner common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.GetTokenOwnerDfilBalance(&_TokenExchange.CallOpts, owner)
}

// GetTokenOwnerFilBalance is a free data retrieval call binding the contract method 0x2a3b967c.
//
// Solidity: function getTokenOwnerFilBalance(address owner) view returns(int256)
func (_TokenExchange *TokenExchangeCaller) GetTokenOwnerFilBalance(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "getTokenOwnerFilBalance", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenOwnerFilBalance is a free data retrieval call binding the contract method 0x2a3b967c.
//
// Solidity: function getTokenOwnerFilBalance(address owner) view returns(int256)
func (_TokenExchange *TokenExchangeSession) GetTokenOwnerFilBalance(owner common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.GetTokenOwnerFilBalance(&_TokenExchange.CallOpts, owner)
}

// GetTokenOwnerFilBalance is a free data retrieval call binding the contract method 0x2a3b967c.
//
// Solidity: function getTokenOwnerFilBalance(address owner) view returns(int256)
func (_TokenExchange *TokenExchangeCallerSession) GetTokenOwnerFilBalance(owner common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.GetTokenOwnerFilBalance(&_TokenExchange.CallOpts, owner)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenExchange *TokenExchangeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenExchange *TokenExchangeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenExchange.Contract.HasRole(&_TokenExchange.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenExchange *TokenExchangeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenExchange.Contract.HasRole(&_TokenExchange.CallOpts, role, account)
}

// LevelHigh is a free data retrieval call binding the contract method 0x67cfd04f.
//
// Solidity: function levelHigh() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) LevelHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "levelHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LevelHigh is a free data retrieval call binding the contract method 0x67cfd04f.
//
// Solidity: function levelHigh() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) LevelHigh() (*big.Int, error) {
	return _TokenExchange.Contract.LevelHigh(&_TokenExchange.CallOpts)
}

// LevelHigh is a free data retrieval call binding the contract method 0x67cfd04f.
//
// Solidity: function levelHigh() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) LevelHigh() (*big.Int, error) {
	return _TokenExchange.Contract.LevelHigh(&_TokenExchange.CallOpts)
}

// LevelLow is a free data retrieval call binding the contract method 0x6ee1434f.
//
// Solidity: function levelLow() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) LevelLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "levelLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LevelLow is a free data retrieval call binding the contract method 0x6ee1434f.
//
// Solidity: function levelLow() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) LevelLow() (*big.Int, error) {
	return _TokenExchange.Contract.LevelLow(&_TokenExchange.CallOpts)
}

// LevelLow is a free data retrieval call binding the contract method 0x6ee1434f.
//
// Solidity: function levelLow() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) LevelLow() (*big.Int, error) {
	return _TokenExchange.Contract.LevelLow(&_TokenExchange.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_TokenExchange *TokenExchangeCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_TokenExchange *TokenExchangeSession) Manager() (common.Address, error) {
	return _TokenExchange.Contract.Manager(&_TokenExchange.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) Manager() (common.Address, error) {
	return _TokenExchange.Contract.Manager(&_TokenExchange.CallOpts)
}

// ManagerAmount is a free data retrieval call binding the contract method 0x3c096005.
//
// Solidity: function managerAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) ManagerAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "managerAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerAmount is a free data retrieval call binding the contract method 0x3c096005.
//
// Solidity: function managerAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) ManagerAmount() (*big.Int, error) {
	return _TokenExchange.Contract.ManagerAmount(&_TokenExchange.CallOpts)
}

// ManagerAmount is a free data retrieval call binding the contract method 0x3c096005.
//
// Solidity: function managerAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) ManagerAmount() (*big.Int, error) {
	return _TokenExchange.Contract.ManagerAmount(&_TokenExchange.CallOpts)
}

// ManagerFilLimitAmount is a free data retrieval call binding the contract method 0x1541ad80.
//
// Solidity: function managerFilLimitAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) ManagerFilLimitAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "managerFilLimitAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFilLimitAmount is a free data retrieval call binding the contract method 0x1541ad80.
//
// Solidity: function managerFilLimitAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) ManagerFilLimitAmount() (*big.Int, error) {
	return _TokenExchange.Contract.ManagerFilLimitAmount(&_TokenExchange.CallOpts)
}

// ManagerFilLimitAmount is a free data retrieval call binding the contract method 0x1541ad80.
//
// Solidity: function managerFilLimitAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) ManagerFilLimitAmount() (*big.Int, error) {
	return _TokenExchange.Contract.ManagerFilLimitAmount(&_TokenExchange.CallOpts)
}

// ManagerLimitAmount is a free data retrieval call binding the contract method 0xea4a9b53.
//
// Solidity: function managerLimitAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) ManagerLimitAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "managerLimitAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerLimitAmount is a free data retrieval call binding the contract method 0xea4a9b53.
//
// Solidity: function managerLimitAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) ManagerLimitAmount() (*big.Int, error) {
	return _TokenExchange.Contract.ManagerLimitAmount(&_TokenExchange.CallOpts)
}

// ManagerLimitAmount is a free data retrieval call binding the contract method 0xea4a9b53.
//
// Solidity: function managerLimitAmount() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) ManagerLimitAmount() (*big.Int, error) {
	return _TokenExchange.Contract.ManagerLimitAmount(&_TokenExchange.CallOpts)
}

// PlatToken is a free data retrieval call binding the contract method 0x42ad046d.
//
// Solidity: function platToken() view returns(address)
func (_TokenExchange *TokenExchangeCaller) PlatToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "platToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatToken is a free data retrieval call binding the contract method 0x42ad046d.
//
// Solidity: function platToken() view returns(address)
func (_TokenExchange *TokenExchangeSession) PlatToken() (common.Address, error) {
	return _TokenExchange.Contract.PlatToken(&_TokenExchange.CallOpts)
}

// PlatToken is a free data retrieval call binding the contract method 0x42ad046d.
//
// Solidity: function platToken() view returns(address)
func (_TokenExchange *TokenExchangeCallerSession) PlatToken() (common.Address, error) {
	return _TokenExchange.Contract.PlatToken(&_TokenExchange.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) Rate() (*big.Int, error) {
	return _TokenExchange.Contract.Rate(&_TokenExchange.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) Rate() (*big.Int, error) {
	return _TokenExchange.Contract.Rate(&_TokenExchange.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenExchange *TokenExchangeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenExchange *TokenExchangeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenExchange.Contract.SupportsInterface(&_TokenExchange.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenExchange *TokenExchangeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenExchange.Contract.SupportsInterface(&_TokenExchange.CallOpts, interfaceId)
}

// TokenOwnerDfilBalance is a free data retrieval call binding the contract method 0x259454bb.
//
// Solidity: function tokenOwnerDfilBalance(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCaller) TokenOwnerDfilBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "tokenOwnerDfilBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOwnerDfilBalance is a free data retrieval call binding the contract method 0x259454bb.
//
// Solidity: function tokenOwnerDfilBalance(address ) view returns(int256)
func (_TokenExchange *TokenExchangeSession) TokenOwnerDfilBalance(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerDfilBalance(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerDfilBalance is a free data retrieval call binding the contract method 0x259454bb.
//
// Solidity: function tokenOwnerDfilBalance(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCallerSession) TokenOwnerDfilBalance(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerDfilBalance(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerFilBalance is a free data retrieval call binding the contract method 0x2157090b.
//
// Solidity: function tokenOwnerFilBalance(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCaller) TokenOwnerFilBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "tokenOwnerFilBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOwnerFilBalance is a free data retrieval call binding the contract method 0x2157090b.
//
// Solidity: function tokenOwnerFilBalance(address ) view returns(int256)
func (_TokenExchange *TokenExchangeSession) TokenOwnerFilBalance(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerFilBalance(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerFilBalance is a free data retrieval call binding the contract method 0x2157090b.
//
// Solidity: function tokenOwnerFilBalance(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCallerSession) TokenOwnerFilBalance(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerFilBalance(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerFilIn is a free data retrieval call binding the contract method 0xd0524749.
//
// Solidity: function tokenOwnerFilIn(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCaller) TokenOwnerFilIn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "tokenOwnerFilIn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOwnerFilIn is a free data retrieval call binding the contract method 0xd0524749.
//
// Solidity: function tokenOwnerFilIn(address ) view returns(int256)
func (_TokenExchange *TokenExchangeSession) TokenOwnerFilIn(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerFilIn(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerFilIn is a free data retrieval call binding the contract method 0xd0524749.
//
// Solidity: function tokenOwnerFilIn(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCallerSession) TokenOwnerFilIn(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerFilIn(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerFilOut is a free data retrieval call binding the contract method 0xace1c7c0.
//
// Solidity: function tokenOwnerFilOut(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCaller) TokenOwnerFilOut(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "tokenOwnerFilOut", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOwnerFilOut is a free data retrieval call binding the contract method 0xace1c7c0.
//
// Solidity: function tokenOwnerFilOut(address ) view returns(int256)
func (_TokenExchange *TokenExchangeSession) TokenOwnerFilOut(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerFilOut(&_TokenExchange.CallOpts, arg0)
}

// TokenOwnerFilOut is a free data retrieval call binding the contract method 0xace1c7c0.
//
// Solidity: function tokenOwnerFilOut(address ) view returns(int256)
func (_TokenExchange *TokenExchangeCallerSession) TokenOwnerFilOut(arg0 common.Address) (*big.Int, error) {
	return _TokenExchange.Contract.TokenOwnerFilOut(&_TokenExchange.CallOpts, arg0)
}

// UnionBase is a free data retrieval call binding the contract method 0xa3916958.
//
// Solidity: function unionBase() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) UnionBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "unionBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnionBase is a free data retrieval call binding the contract method 0xa3916958.
//
// Solidity: function unionBase() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) UnionBase() (*big.Int, error) {
	return _TokenExchange.Contract.UnionBase(&_TokenExchange.CallOpts)
}

// UnionBase is a free data retrieval call binding the contract method 0xa3916958.
//
// Solidity: function unionBase() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) UnionBase() (*big.Int, error) {
	return _TokenExchange.Contract.UnionBase(&_TokenExchange.CallOpts)
}

// UnionLowRate is a free data retrieval call binding the contract method 0xc777939d.
//
// Solidity: function unionLowRate() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) UnionLowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "unionLowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnionLowRate is a free data retrieval call binding the contract method 0xc777939d.
//
// Solidity: function unionLowRate() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) UnionLowRate() (*big.Int, error) {
	return _TokenExchange.Contract.UnionLowRate(&_TokenExchange.CallOpts)
}

// UnionLowRate is a free data retrieval call binding the contract method 0xc777939d.
//
// Solidity: function unionLowRate() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) UnionLowRate() (*big.Int, error) {
	return _TokenExchange.Contract.UnionLowRate(&_TokenExchange.CallOpts)
}

// UnionRate is a free data retrieval call binding the contract method 0xc14cb606.
//
// Solidity: function unionRate() view returns(uint256)
func (_TokenExchange *TokenExchangeCaller) UnionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "unionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnionRate is a free data retrieval call binding the contract method 0xc14cb606.
//
// Solidity: function unionRate() view returns(uint256)
func (_TokenExchange *TokenExchangeSession) UnionRate() (*big.Int, error) {
	return _TokenExchange.Contract.UnionRate(&_TokenExchange.CallOpts)
}

// UnionRate is a free data retrieval call binding the contract method 0xc14cb606.
//
// Solidity: function unionRate() view returns(uint256)
func (_TokenExchange *TokenExchangeCallerSession) UnionRate() (*big.Int, error) {
	return _TokenExchange.Contract.UnionRate(&_TokenExchange.CallOpts)
}

// UseManager is a free data retrieval call binding the contract method 0x5934032d.
//
// Solidity: function useManager() view returns(bool)
func (_TokenExchange *TokenExchangeCaller) UseManager(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenExchange.contract.Call(opts, &out, "useManager")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseManager is a free data retrieval call binding the contract method 0x5934032d.
//
// Solidity: function useManager() view returns(bool)
func (_TokenExchange *TokenExchangeSession) UseManager() (bool, error) {
	return _TokenExchange.Contract.UseManager(&_TokenExchange.CallOpts)
}

// UseManager is a free data retrieval call binding the contract method 0x5934032d.
//
// Solidity: function useManager() view returns(bool)
func (_TokenExchange *TokenExchangeCallerSession) UseManager() (bool, error) {
	return _TokenExchange.Contract.UseManager(&_TokenExchange.CallOpts)
}

// DFIL2FIL is a paid mutator transaction binding the contract method 0x0aee2b54.
//
// Solidity: function DFIL2FIL(uint256 amount, uint256 usePlatToken) returns()
func (_TokenExchange *TokenExchangeTransactor) DFIL2FIL(opts *bind.TransactOpts, amount *big.Int, usePlatToken *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "DFIL2FIL", amount, usePlatToken)
}

// DFIL2FIL is a paid mutator transaction binding the contract method 0x0aee2b54.
//
// Solidity: function DFIL2FIL(uint256 amount, uint256 usePlatToken) returns()
func (_TokenExchange *TokenExchangeSession) DFIL2FIL(amount *big.Int, usePlatToken *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.DFIL2FIL(&_TokenExchange.TransactOpts, amount, usePlatToken)
}

// DFIL2FIL is a paid mutator transaction binding the contract method 0x0aee2b54.
//
// Solidity: function DFIL2FIL(uint256 amount, uint256 usePlatToken) returns()
func (_TokenExchange *TokenExchangeTransactorSession) DFIL2FIL(amount *big.Int, usePlatToken *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.DFIL2FIL(&_TokenExchange.TransactOpts, amount, usePlatToken)
}

// FIL2DFIL is a paid mutator transaction binding the contract method 0xedfd890c.
//
// Solidity: function FIL2DFIL() payable returns()
func (_TokenExchange *TokenExchangeTransactor) FIL2DFIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "FIL2DFIL")
}

// FIL2DFIL is a paid mutator transaction binding the contract method 0xedfd890c.
//
// Solidity: function FIL2DFIL() payable returns()
func (_TokenExchange *TokenExchangeSession) FIL2DFIL() (*types.Transaction, error) {
	return _TokenExchange.Contract.FIL2DFIL(&_TokenExchange.TransactOpts)
}

// FIL2DFIL is a paid mutator transaction binding the contract method 0xedfd890c.
//
// Solidity: function FIL2DFIL() payable returns()
func (_TokenExchange *TokenExchangeTransactorSession) FIL2DFIL() (*types.Transaction, error) {
	return _TokenExchange.Contract.FIL2DFIL(&_TokenExchange.TransactOpts)
}

// FILIN is a paid mutator transaction binding the contract method 0xd04ff329.
//
// Solidity: function FILIN(address owner) payable returns()
func (_TokenExchange *TokenExchangeTransactor) FILIN(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "FILIN", owner)
}

// FILIN is a paid mutator transaction binding the contract method 0xd04ff329.
//
// Solidity: function FILIN(address owner) payable returns()
func (_TokenExchange *TokenExchangeSession) FILIN(owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.FILIN(&_TokenExchange.TransactOpts, owner)
}

// FILIN is a paid mutator transaction binding the contract method 0xd04ff329.
//
// Solidity: function FILIN(address owner) payable returns()
func (_TokenExchange *TokenExchangeTransactorSession) FILIN(owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.FILIN(&_TokenExchange.TransactOpts, owner)
}

// FILLOCK is a paid mutator transaction binding the contract method 0xc0dc667e.
//
// Solidity: function FILLOCK(address owner, uint256 costAmount, uint256 profitRateAmount, uint256 depositAmount, bool union) returns()
func (_TokenExchange *TokenExchangeTransactor) FILLOCK(opts *bind.TransactOpts, owner common.Address, costAmount *big.Int, profitRateAmount *big.Int, depositAmount *big.Int, union bool) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "FILLOCK", owner, costAmount, profitRateAmount, depositAmount, union)
}

// FILLOCK is a paid mutator transaction binding the contract method 0xc0dc667e.
//
// Solidity: function FILLOCK(address owner, uint256 costAmount, uint256 profitRateAmount, uint256 depositAmount, bool union) returns()
func (_TokenExchange *TokenExchangeSession) FILLOCK(owner common.Address, costAmount *big.Int, profitRateAmount *big.Int, depositAmount *big.Int, union bool) (*types.Transaction, error) {
	return _TokenExchange.Contract.FILLOCK(&_TokenExchange.TransactOpts, owner, costAmount, profitRateAmount, depositAmount, union)
}

// FILLOCK is a paid mutator transaction binding the contract method 0xc0dc667e.
//
// Solidity: function FILLOCK(address owner, uint256 costAmount, uint256 profitRateAmount, uint256 depositAmount, bool union) returns()
func (_TokenExchange *TokenExchangeTransactorSession) FILLOCK(owner common.Address, costAmount *big.Int, profitRateAmount *big.Int, depositAmount *big.Int, union bool) (*types.Transaction, error) {
	return _TokenExchange.Contract.FILLOCK(&_TokenExchange.TransactOpts, owner, costAmount, profitRateAmount, depositAmount, union)
}

// REWARDFIL2DFIL is a paid mutator transaction binding the contract method 0x0fceee20.
//
// Solidity: function REWARDFIL2DFIL(address owner) payable returns()
func (_TokenExchange *TokenExchangeTransactor) REWARDFIL2DFIL(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "REWARDFIL2DFIL", owner)
}

// REWARDFIL2DFIL is a paid mutator transaction binding the contract method 0x0fceee20.
//
// Solidity: function REWARDFIL2DFIL(address owner) payable returns()
func (_TokenExchange *TokenExchangeSession) REWARDFIL2DFIL(owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.REWARDFIL2DFIL(&_TokenExchange.TransactOpts, owner)
}

// REWARDFIL2DFIL is a paid mutator transaction binding the contract method 0x0fceee20.
//
// Solidity: function REWARDFIL2DFIL(address owner) payable returns()
func (_TokenExchange *TokenExchangeTransactorSession) REWARDFIL2DFIL(owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.REWARDFIL2DFIL(&_TokenExchange.TransactOpts, owner)
}

// UNIONWITHDRAWFIL is a paid mutator transaction binding the contract method 0x0f9f6cf6.
//
// Solidity: function UNIONWITHDRAWFIL(address owner, uint256 costAmount, uint256 depositAmount, uint256 usePlatToken) returns()
func (_TokenExchange *TokenExchangeTransactor) UNIONWITHDRAWFIL(opts *bind.TransactOpts, owner common.Address, costAmount *big.Int, depositAmount *big.Int, usePlatToken *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "UNIONWITHDRAWFIL", owner, costAmount, depositAmount, usePlatToken)
}

// UNIONWITHDRAWFIL is a paid mutator transaction binding the contract method 0x0f9f6cf6.
//
// Solidity: function UNIONWITHDRAWFIL(address owner, uint256 costAmount, uint256 depositAmount, uint256 usePlatToken) returns()
func (_TokenExchange *TokenExchangeSession) UNIONWITHDRAWFIL(owner common.Address, costAmount *big.Int, depositAmount *big.Int, usePlatToken *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.UNIONWITHDRAWFIL(&_TokenExchange.TransactOpts, owner, costAmount, depositAmount, usePlatToken)
}

// UNIONWITHDRAWFIL is a paid mutator transaction binding the contract method 0x0f9f6cf6.
//
// Solidity: function UNIONWITHDRAWFIL(address owner, uint256 costAmount, uint256 depositAmount, uint256 usePlatToken) returns()
func (_TokenExchange *TokenExchangeTransactorSession) UNIONWITHDRAWFIL(owner common.Address, costAmount *big.Int, depositAmount *big.Int, usePlatToken *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.UNIONWITHDRAWFIL(&_TokenExchange.TransactOpts, owner, costAmount, depositAmount, usePlatToken)
}

// CloseManager is a paid mutator transaction binding the contract method 0x750082cb.
//
// Solidity: function closeManager(uint256 amount) returns()
func (_TokenExchange *TokenExchangeTransactor) CloseManager(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "closeManager", amount)
}

// CloseManager is a paid mutator transaction binding the contract method 0x750082cb.
//
// Solidity: function closeManager(uint256 amount) returns()
func (_TokenExchange *TokenExchangeSession) CloseManager(amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.CloseManager(&_TokenExchange.TransactOpts, amount)
}

// CloseManager is a paid mutator transaction binding the contract method 0x750082cb.
//
// Solidity: function closeManager(uint256 amount) returns()
func (_TokenExchange *TokenExchangeTransactorSession) CloseManager(amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.CloseManager(&_TokenExchange.TransactOpts, amount)
}

// FeeToWithdraw is a paid mutator transaction binding the contract method 0xb78f076a.
//
// Solidity: function feeToWithdraw() returns()
func (_TokenExchange *TokenExchangeTransactor) FeeToWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "feeToWithdraw")
}

// FeeToWithdraw is a paid mutator transaction binding the contract method 0xb78f076a.
//
// Solidity: function feeToWithdraw() returns()
func (_TokenExchange *TokenExchangeSession) FeeToWithdraw() (*types.Transaction, error) {
	return _TokenExchange.Contract.FeeToWithdraw(&_TokenExchange.TransactOpts)
}

// FeeToWithdraw is a paid mutator transaction binding the contract method 0xb78f076a.
//
// Solidity: function feeToWithdraw() returns()
func (_TokenExchange *TokenExchangeTransactorSession) FeeToWithdraw() (*types.Transaction, error) {
	return _TokenExchange.Contract.FeeToWithdraw(&_TokenExchange.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.GrantRole(&_TokenExchange.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.GrantRole(&_TokenExchange.TransactOpts, role, account)
}

// OwnerFilBalanceChange is a paid mutator transaction binding the contract method 0x728e6966.
//
// Solidity: function ownerFilBalanceChange(address owner) returns()
func (_TokenExchange *TokenExchangeTransactor) OwnerFilBalanceChange(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "ownerFilBalanceChange", owner)
}

// OwnerFilBalanceChange is a paid mutator transaction binding the contract method 0x728e6966.
//
// Solidity: function ownerFilBalanceChange(address owner) returns()
func (_TokenExchange *TokenExchangeSession) OwnerFilBalanceChange(owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.OwnerFilBalanceChange(&_TokenExchange.TransactOpts, owner)
}

// OwnerFilBalanceChange is a paid mutator transaction binding the contract method 0x728e6966.
//
// Solidity: function ownerFilBalanceChange(address owner) returns()
func (_TokenExchange *TokenExchangeTransactorSession) OwnerFilBalanceChange(owner common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.OwnerFilBalanceChange(&_TokenExchange.TransactOpts, owner)
}

// RemoveAccountUnion is a paid mutator transaction binding the contract method 0x1d9e23b0.
//
// Solidity: function removeAccountUnion() returns()
func (_TokenExchange *TokenExchangeTransactor) RemoveAccountUnion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "removeAccountUnion")
}

// RemoveAccountUnion is a paid mutator transaction binding the contract method 0x1d9e23b0.
//
// Solidity: function removeAccountUnion() returns()
func (_TokenExchange *TokenExchangeSession) RemoveAccountUnion() (*types.Transaction, error) {
	return _TokenExchange.Contract.RemoveAccountUnion(&_TokenExchange.TransactOpts)
}

// RemoveAccountUnion is a paid mutator transaction binding the contract method 0x1d9e23b0.
//
// Solidity: function removeAccountUnion() returns()
func (_TokenExchange *TokenExchangeTransactorSession) RemoveAccountUnion() (*types.Transaction, error) {
	return _TokenExchange.Contract.RemoveAccountUnion(&_TokenExchange.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.RenounceRole(&_TokenExchange.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.RenounceRole(&_TokenExchange.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.RevokeRole(&_TokenExchange.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenExchange *TokenExchangeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.RevokeRole(&_TokenExchange.TransactOpts, role, account)
}

// SetAccountUnion is a paid mutator transaction binding the contract method 0xac40c300.
//
// Solidity: function setAccountUnion(uint256 userDefineRate) returns()
func (_TokenExchange *TokenExchangeTransactor) SetAccountUnion(opts *bind.TransactOpts, userDefineRate *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setAccountUnion", userDefineRate)
}

// SetAccountUnion is a paid mutator transaction binding the contract method 0xac40c300.
//
// Solidity: function setAccountUnion(uint256 userDefineRate) returns()
func (_TokenExchange *TokenExchangeSession) SetAccountUnion(userDefineRate *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetAccountUnion(&_TokenExchange.TransactOpts, userDefineRate)
}

// SetAccountUnion is a paid mutator transaction binding the contract method 0xac40c300.
//
// Solidity: function setAccountUnion(uint256 userDefineRate) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetAccountUnion(userDefineRate *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetAccountUnion(&_TokenExchange.TransactOpts, userDefineRate)
}

// SetAllUserExchangeEnable is a paid mutator transaction binding the contract method 0xf2de9ca1.
//
// Solidity: function setAllUserExchangeEnable() returns()
func (_TokenExchange *TokenExchangeTransactor) SetAllUserExchangeEnable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setAllUserExchangeEnable")
}

// SetAllUserExchangeEnable is a paid mutator transaction binding the contract method 0xf2de9ca1.
//
// Solidity: function setAllUserExchangeEnable() returns()
func (_TokenExchange *TokenExchangeSession) SetAllUserExchangeEnable() (*types.Transaction, error) {
	return _TokenExchange.Contract.SetAllUserExchangeEnable(&_TokenExchange.TransactOpts)
}

// SetAllUserExchangeEnable is a paid mutator transaction binding the contract method 0xf2de9ca1.
//
// Solidity: function setAllUserExchangeEnable() returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetAllUserExchangeEnable() (*types.Transaction, error) {
	return _TokenExchange.Contract.SetAllUserExchangeEnable(&_TokenExchange.TransactOpts)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address factory_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetFactory(opts *bind.TransactOpts, factory_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setFactory", factory_)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address factory_) returns()
func (_TokenExchange *TokenExchangeSession) SetFactory(factory_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetFactory(&_TokenExchange.TransactOpts, factory_)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address factory_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetFactory(factory_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetFactory(&_TokenExchange.TransactOpts, factory_)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 rate_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetFee(opts *bind.TransactOpts, rate_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setFee", rate_)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 rate_) returns()
func (_TokenExchange *TokenExchangeSession) SetFee(rate_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetFee(&_TokenExchange.TransactOpts, rate_)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 rate_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetFee(rate_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetFee(&_TokenExchange.TransactOpts, rate_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetFeeTo(opts *bind.TransactOpts, feeTo_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setFeeTo", feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_TokenExchange *TokenExchangeSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetFeeTo(&_TokenExchange.TransactOpts, feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetFeeTo(&_TokenExchange.TransactOpts, feeTo_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetManager(opts *bind.TransactOpts, manager_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setManager", manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_TokenExchange *TokenExchangeSession) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManager(&_TokenExchange.TransactOpts, manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManager(&_TokenExchange.TransactOpts, manager_)
}

// SetManagerGetFilLevel is a paid mutator transaction binding the contract method 0x19f6c385.
//
// Solidity: function setManagerGetFilLevel(uint256 amount) returns()
func (_TokenExchange *TokenExchangeTransactor) SetManagerGetFilLevel(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setManagerGetFilLevel", amount)
}

// SetManagerGetFilLevel is a paid mutator transaction binding the contract method 0x19f6c385.
//
// Solidity: function setManagerGetFilLevel(uint256 amount) returns()
func (_TokenExchange *TokenExchangeSession) SetManagerGetFilLevel(amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManagerGetFilLevel(&_TokenExchange.TransactOpts, amount)
}

// SetManagerGetFilLevel is a paid mutator transaction binding the contract method 0x19f6c385.
//
// Solidity: function setManagerGetFilLevel(uint256 amount) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetManagerGetFilLevel(amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManagerGetFilLevel(&_TokenExchange.TransactOpts, amount)
}

// SetManagerLevel is a paid mutator transaction binding the contract method 0x7c6f14cc.
//
// Solidity: function setManagerLevel(uint256 levelLow_, uint256 levelHigh_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetManagerLevel(opts *bind.TransactOpts, levelLow_ *big.Int, levelHigh_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setManagerLevel", levelLow_, levelHigh_)
}

// SetManagerLevel is a paid mutator transaction binding the contract method 0x7c6f14cc.
//
// Solidity: function setManagerLevel(uint256 levelLow_, uint256 levelHigh_) returns()
func (_TokenExchange *TokenExchangeSession) SetManagerLevel(levelLow_ *big.Int, levelHigh_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManagerLevel(&_TokenExchange.TransactOpts, levelLow_, levelHigh_)
}

// SetManagerLevel is a paid mutator transaction binding the contract method 0x7c6f14cc.
//
// Solidity: function setManagerLevel(uint256 levelLow_, uint256 levelHigh_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetManagerLevel(levelLow_ *big.Int, levelHigh_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManagerLevel(&_TokenExchange.TransactOpts, levelLow_, levelHigh_)
}

// SetManagerLimitAmount is a paid mutator transaction binding the contract method 0xb1ac6385.
//
// Solidity: function setManagerLimitAmount(uint256 amount) returns()
func (_TokenExchange *TokenExchangeTransactor) SetManagerLimitAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setManagerLimitAmount", amount)
}

// SetManagerLimitAmount is a paid mutator transaction binding the contract method 0xb1ac6385.
//
// Solidity: function setManagerLimitAmount(uint256 amount) returns()
func (_TokenExchange *TokenExchangeSession) SetManagerLimitAmount(amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManagerLimitAmount(&_TokenExchange.TransactOpts, amount)
}

// SetManagerLimitAmount is a paid mutator transaction binding the contract method 0xb1ac6385.
//
// Solidity: function setManagerLimitAmount(uint256 amount) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetManagerLimitAmount(amount *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetManagerLimitAmount(&_TokenExchange.TransactOpts, amount)
}

// SetPlatToken is a paid mutator transaction binding the contract method 0x4fdc7efd.
//
// Solidity: function setPlatToken(address platToken_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetPlatToken(opts *bind.TransactOpts, platToken_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setPlatToken", platToken_)
}

// SetPlatToken is a paid mutator transaction binding the contract method 0x4fdc7efd.
//
// Solidity: function setPlatToken(address platToken_) returns()
func (_TokenExchange *TokenExchangeSession) SetPlatToken(platToken_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetPlatToken(&_TokenExchange.TransactOpts, platToken_)
}

// SetPlatToken is a paid mutator transaction binding the contract method 0x4fdc7efd.
//
// Solidity: function setPlatToken(address platToken_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetPlatToken(platToken_ common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetPlatToken(&_TokenExchange.TransactOpts, platToken_)
}

// SetTokenRole is a paid mutator transaction binding the contract method 0x13a834b7.
//
// Solidity: function setTokenRole(address token) returns()
func (_TokenExchange *TokenExchangeTransactor) SetTokenRole(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setTokenRole", token)
}

// SetTokenRole is a paid mutator transaction binding the contract method 0x13a834b7.
//
// Solidity: function setTokenRole(address token) returns()
func (_TokenExchange *TokenExchangeSession) SetTokenRole(token common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetTokenRole(&_TokenExchange.TransactOpts, token)
}

// SetTokenRole is a paid mutator transaction binding the contract method 0x13a834b7.
//
// Solidity: function setTokenRole(address token) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetTokenRole(token common.Address) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetTokenRole(&_TokenExchange.TransactOpts, token)
}

// SetUnionRate is a paid mutator transaction binding the contract method 0xbcd0551c.
//
// Solidity: function setUnionRate(uint256 unionLowRate_, uint256 unionRate_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetUnionRate(opts *bind.TransactOpts, unionLowRate_ *big.Int, unionRate_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setUnionRate", unionLowRate_, unionRate_)
}

// SetUnionRate is a paid mutator transaction binding the contract method 0xbcd0551c.
//
// Solidity: function setUnionRate(uint256 unionLowRate_, uint256 unionRate_) returns()
func (_TokenExchange *TokenExchangeSession) SetUnionRate(unionLowRate_ *big.Int, unionRate_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetUnionRate(&_TokenExchange.TransactOpts, unionLowRate_, unionRate_)
}

// SetUnionRate is a paid mutator transaction binding the contract method 0xbcd0551c.
//
// Solidity: function setUnionRate(uint256 unionLowRate_, uint256 unionRate_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetUnionRate(unionLowRate_ *big.Int, unionRate_ *big.Int) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetUnionRate(&_TokenExchange.TransactOpts, unionLowRate_, unionRate_)
}

// SetUseManager is a paid mutator transaction binding the contract method 0x0829790a.
//
// Solidity: function setUseManager(bool useManager_) returns()
func (_TokenExchange *TokenExchangeTransactor) SetUseManager(opts *bind.TransactOpts, useManager_ bool) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "setUseManager", useManager_)
}

// SetUseManager is a paid mutator transaction binding the contract method 0x0829790a.
//
// Solidity: function setUseManager(bool useManager_) returns()
func (_TokenExchange *TokenExchangeSession) SetUseManager(useManager_ bool) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetUseManager(&_TokenExchange.TransactOpts, useManager_)
}

// SetUseManager is a paid mutator transaction binding the contract method 0x0829790a.
//
// Solidity: function setUseManager(bool useManager_) returns()
func (_TokenExchange *TokenExchangeTransactorSession) SetUseManager(useManager_ bool) (*types.Transaction, error) {
	return _TokenExchange.Contract.SetUseManager(&_TokenExchange.TransactOpts, useManager_)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_TokenExchange *TokenExchangeTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_TokenExchange *TokenExchangeSession) WithdrawAll() (*types.Transaction, error) {
	return _TokenExchange.Contract.WithdrawAll(&_TokenExchange.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_TokenExchange *TokenExchangeTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _TokenExchange.Contract.WithdrawAll(&_TokenExchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenExchange *TokenExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenExchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenExchange *TokenExchangeSession) Receive() (*types.Transaction, error) {
	return _TokenExchange.Contract.Receive(&_TokenExchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TokenExchange *TokenExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _TokenExchange.Contract.Receive(&_TokenExchange.TransactOpts)
}

// TokenExchangeAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the TokenExchange contract.
type TokenExchangeAccountCreatedIterator struct {
	Event *TokenExchangeAccountCreated // Event containing the contract specifics and raw log

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
func (it *TokenExchangeAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeAccountCreated)
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
		it.Event = new(TokenExchangeAccountCreated)
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
func (it *TokenExchangeAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeAccountCreated represents a AccountCreated event raised by the TokenExchange contract.
type TokenExchangeAccountCreated struct {
	User common.Address
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x7325fb14450d3e5eb288c620b4ff1d6b43a41b33b4a4f143ae88cd4c12f99ea3.
//
// Solidity: event AccountCreated(address indexed user, uint256 rate)
func (_TokenExchange *TokenExchangeFilterer) FilterAccountCreated(opts *bind.FilterOpts, user []common.Address) (*TokenExchangeAccountCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "AccountCreated", userRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeAccountCreatedIterator{contract: _TokenExchange.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x7325fb14450d3e5eb288c620b4ff1d6b43a41b33b4a4f143ae88cd4c12f99ea3.
//
// Solidity: event AccountCreated(address indexed user, uint256 rate)
func (_TokenExchange *TokenExchangeFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *TokenExchangeAccountCreated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "AccountCreated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeAccountCreated)
				if err := _TokenExchange.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x7325fb14450d3e5eb288c620b4ff1d6b43a41b33b4a4f143ae88cd4c12f99ea3.
//
// Solidity: event AccountCreated(address indexed user, uint256 rate)
func (_TokenExchange *TokenExchangeFilterer) ParseAccountCreated(log types.Log) (*TokenExchangeAccountCreated, error) {
	event := new(TokenExchangeAccountCreated)
	if err := _TokenExchange.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeAccountRemovedIterator is returned from FilterAccountRemoved and is used to iterate over the raw logs and unpacked data for AccountRemoved events raised by the TokenExchange contract.
type TokenExchangeAccountRemovedIterator struct {
	Event *TokenExchangeAccountRemoved // Event containing the contract specifics and raw log

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
func (it *TokenExchangeAccountRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeAccountRemoved)
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
		it.Event = new(TokenExchangeAccountRemoved)
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
func (it *TokenExchangeAccountRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeAccountRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeAccountRemoved represents a AccountRemoved event raised by the TokenExchange contract.
type TokenExchangeAccountRemoved struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAccountRemoved is a free log retrieval operation binding the contract event 0xbf2e373b8263f701e10efcac80ea442afcb29c6852b3a42b0b46cc8edaaf54a7.
//
// Solidity: event AccountRemoved(address indexed user)
func (_TokenExchange *TokenExchangeFilterer) FilterAccountRemoved(opts *bind.FilterOpts, user []common.Address) (*TokenExchangeAccountRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "AccountRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeAccountRemovedIterator{contract: _TokenExchange.contract, event: "AccountRemoved", logs: logs, sub: sub}, nil
}

// WatchAccountRemoved is a free log subscription operation binding the contract event 0xbf2e373b8263f701e10efcac80ea442afcb29c6852b3a42b0b46cc8edaaf54a7.
//
// Solidity: event AccountRemoved(address indexed user)
func (_TokenExchange *TokenExchangeFilterer) WatchAccountRemoved(opts *bind.WatchOpts, sink chan<- *TokenExchangeAccountRemoved, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "AccountRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeAccountRemoved)
				if err := _TokenExchange.contract.UnpackLog(event, "AccountRemoved", log); err != nil {
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

// ParseAccountRemoved is a log parse operation binding the contract event 0xbf2e373b8263f701e10efcac80ea442afcb29c6852b3a42b0b46cc8edaaf54a7.
//
// Solidity: event AccountRemoved(address indexed user)
func (_TokenExchange *TokenExchangeFilterer) ParseAccountRemoved(log types.Log) (*TokenExchangeAccountRemoved, error) {
	event := new(TokenExchangeAccountRemoved)
	if err := _TokenExchange.contract.UnpackLog(event, "AccountRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeDFIL2FILExchangedIterator is returned from FilterDFIL2FILExchanged and is used to iterate over the raw logs and unpacked data for DFIL2FILExchanged events raised by the TokenExchange contract.
type TokenExchangeDFIL2FILExchangedIterator struct {
	Event *TokenExchangeDFIL2FILExchanged // Event containing the contract specifics and raw log

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
func (it *TokenExchangeDFIL2FILExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeDFIL2FILExchanged)
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
		it.Event = new(TokenExchangeDFIL2FILExchanged)
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
func (it *TokenExchangeDFIL2FILExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeDFIL2FILExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeDFIL2FILExchanged represents a DFIL2FILExchanged event raised by the TokenExchange contract.
type TokenExchangeDFIL2FILExchanged struct {
	User         common.Address
	Amount       *big.Int
	RelAmount    *big.Int
	Fee          *big.Int
	UsePlatToken *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDFIL2FILExchanged is a free log retrieval operation binding the contract event 0x3947fb2e0cb302073aabff3a963ea943dc25e21dc3bfd6365610178e16605925.
//
// Solidity: event DFIL2FILExchanged(address indexed user, uint256 amount, uint256 relAmount, uint256 fee, uint256 usePlatToken)
func (_TokenExchange *TokenExchangeFilterer) FilterDFIL2FILExchanged(opts *bind.FilterOpts, user []common.Address) (*TokenExchangeDFIL2FILExchangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "DFIL2FILExchanged", userRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeDFIL2FILExchangedIterator{contract: _TokenExchange.contract, event: "DFIL2FILExchanged", logs: logs, sub: sub}, nil
}

// WatchDFIL2FILExchanged is a free log subscription operation binding the contract event 0x3947fb2e0cb302073aabff3a963ea943dc25e21dc3bfd6365610178e16605925.
//
// Solidity: event DFIL2FILExchanged(address indexed user, uint256 amount, uint256 relAmount, uint256 fee, uint256 usePlatToken)
func (_TokenExchange *TokenExchangeFilterer) WatchDFIL2FILExchanged(opts *bind.WatchOpts, sink chan<- *TokenExchangeDFIL2FILExchanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "DFIL2FILExchanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeDFIL2FILExchanged)
				if err := _TokenExchange.contract.UnpackLog(event, "DFIL2FILExchanged", log); err != nil {
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

// ParseDFIL2FILExchanged is a log parse operation binding the contract event 0x3947fb2e0cb302073aabff3a963ea943dc25e21dc3bfd6365610178e16605925.
//
// Solidity: event DFIL2FILExchanged(address indexed user, uint256 amount, uint256 relAmount, uint256 fee, uint256 usePlatToken)
func (_TokenExchange *TokenExchangeFilterer) ParseDFIL2FILExchanged(log types.Log) (*TokenExchangeDFIL2FILExchanged, error) {
	event := new(TokenExchangeDFIL2FILExchanged)
	if err := _TokenExchange.contract.UnpackLog(event, "DFIL2FILExchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeFIL2DFILExchangedIterator is returned from FilterFIL2DFILExchanged and is used to iterate over the raw logs and unpacked data for FIL2DFILExchanged events raised by the TokenExchange contract.
type TokenExchangeFIL2DFILExchangedIterator struct {
	Event *TokenExchangeFIL2DFILExchanged // Event containing the contract specifics and raw log

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
func (it *TokenExchangeFIL2DFILExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeFIL2DFILExchanged)
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
		it.Event = new(TokenExchangeFIL2DFILExchanged)
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
func (it *TokenExchangeFIL2DFILExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeFIL2DFILExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeFIL2DFILExchanged represents a FIL2DFILExchanged event raised by the TokenExchange contract.
type TokenExchangeFIL2DFILExchanged struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFIL2DFILExchanged is a free log retrieval operation binding the contract event 0xa22ed7ce65493bf63e3fab859f0f6787d441ebc982c37f905c64e687ced72a11.
//
// Solidity: event FIL2DFILExchanged(address indexed user, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) FilterFIL2DFILExchanged(opts *bind.FilterOpts, user []common.Address) (*TokenExchangeFIL2DFILExchangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "FIL2DFILExchanged", userRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeFIL2DFILExchangedIterator{contract: _TokenExchange.contract, event: "FIL2DFILExchanged", logs: logs, sub: sub}, nil
}

// WatchFIL2DFILExchanged is a free log subscription operation binding the contract event 0xa22ed7ce65493bf63e3fab859f0f6787d441ebc982c37f905c64e687ced72a11.
//
// Solidity: event FIL2DFILExchanged(address indexed user, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) WatchFIL2DFILExchanged(opts *bind.WatchOpts, sink chan<- *TokenExchangeFIL2DFILExchanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "FIL2DFILExchanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeFIL2DFILExchanged)
				if err := _TokenExchange.contract.UnpackLog(event, "FIL2DFILExchanged", log); err != nil {
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

// ParseFIL2DFILExchanged is a log parse operation binding the contract event 0xa22ed7ce65493bf63e3fab859f0f6787d441ebc982c37f905c64e687ced72a11.
//
// Solidity: event FIL2DFILExchanged(address indexed user, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) ParseFIL2DFILExchanged(log types.Log) (*TokenExchangeFIL2DFILExchanged, error) {
	event := new(TokenExchangeFIL2DFILExchanged)
	if err := _TokenExchange.contract.UnpackLog(event, "FIL2DFILExchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeFILINComplatedIterator is returned from FilterFILINComplated and is used to iterate over the raw logs and unpacked data for FILINComplated events raised by the TokenExchange contract.
type TokenExchangeFILINComplatedIterator struct {
	Event *TokenExchangeFILINComplated // Event containing the contract specifics and raw log

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
func (it *TokenExchangeFILINComplatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeFILINComplated)
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
		it.Event = new(TokenExchangeFILINComplated)
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
func (it *TokenExchangeFILINComplatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeFILINComplatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeFILINComplated represents a FILINComplated event raised by the TokenExchange contract.
type TokenExchangeFILINComplated struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFILINComplated is a free log retrieval operation binding the contract event 0x2939fc3393cde7dea85a175005f92d665e3fb4aa56322da4b12a2d2170dbad7a.
//
// Solidity: event FILINComplated(address indexed token, address owner, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) FilterFILINComplated(opts *bind.FilterOpts, token []common.Address) (*TokenExchangeFILINComplatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "FILINComplated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeFILINComplatedIterator{contract: _TokenExchange.contract, event: "FILINComplated", logs: logs, sub: sub}, nil
}

// WatchFILINComplated is a free log subscription operation binding the contract event 0x2939fc3393cde7dea85a175005f92d665e3fb4aa56322da4b12a2d2170dbad7a.
//
// Solidity: event FILINComplated(address indexed token, address owner, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) WatchFILINComplated(opts *bind.WatchOpts, sink chan<- *TokenExchangeFILINComplated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "FILINComplated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeFILINComplated)
				if err := _TokenExchange.contract.UnpackLog(event, "FILINComplated", log); err != nil {
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

// ParseFILINComplated is a log parse operation binding the contract event 0x2939fc3393cde7dea85a175005f92d665e3fb4aa56322da4b12a2d2170dbad7a.
//
// Solidity: event FILINComplated(address indexed token, address owner, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) ParseFILINComplated(log types.Log) (*TokenExchangeFILINComplated, error) {
	event := new(TokenExchangeFILINComplated)
	if err := _TokenExchange.contract.UnpackLog(event, "FILINComplated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeFILLOCKComplatedIterator is returned from FilterFILLOCKComplated and is used to iterate over the raw logs and unpacked data for FILLOCKComplated events raised by the TokenExchange contract.
type TokenExchangeFILLOCKComplatedIterator struct {
	Event *TokenExchangeFILLOCKComplated // Event containing the contract specifics and raw log

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
func (it *TokenExchangeFILLOCKComplatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeFILLOCKComplated)
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
		it.Event = new(TokenExchangeFILLOCKComplated)
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
func (it *TokenExchangeFILLOCKComplatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeFILLOCKComplatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeFILLOCKComplated represents a FILLOCKComplated event raised by the TokenExchange contract.
type TokenExchangeFILLOCKComplated struct {
	Token            common.Address
	Owner            common.Address
	CostAmount       *big.Int
	ProfitRateAmount *big.Int
	DepositAmount    *big.Int
	Union            bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFILLOCKComplated is a free log retrieval operation binding the contract event 0x2fdb24e32119e887d39699cecff249ef427819e17c0c35b0edb40cb08ac8b1cd.
//
// Solidity: event FILLOCKComplated(address indexed token, address owner, uint256 costAmount, uint256 profitRateAmount, uint256 depositAmount, bool union)
func (_TokenExchange *TokenExchangeFilterer) FilterFILLOCKComplated(opts *bind.FilterOpts, token []common.Address) (*TokenExchangeFILLOCKComplatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "FILLOCKComplated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeFILLOCKComplatedIterator{contract: _TokenExchange.contract, event: "FILLOCKComplated", logs: logs, sub: sub}, nil
}

// WatchFILLOCKComplated is a free log subscription operation binding the contract event 0x2fdb24e32119e887d39699cecff249ef427819e17c0c35b0edb40cb08ac8b1cd.
//
// Solidity: event FILLOCKComplated(address indexed token, address owner, uint256 costAmount, uint256 profitRateAmount, uint256 depositAmount, bool union)
func (_TokenExchange *TokenExchangeFilterer) WatchFILLOCKComplated(opts *bind.WatchOpts, sink chan<- *TokenExchangeFILLOCKComplated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "FILLOCKComplated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeFILLOCKComplated)
				if err := _TokenExchange.contract.UnpackLog(event, "FILLOCKComplated", log); err != nil {
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

// ParseFILLOCKComplated is a log parse operation binding the contract event 0x2fdb24e32119e887d39699cecff249ef427819e17c0c35b0edb40cb08ac8b1cd.
//
// Solidity: event FILLOCKComplated(address indexed token, address owner, uint256 costAmount, uint256 profitRateAmount, uint256 depositAmount, bool union)
func (_TokenExchange *TokenExchangeFilterer) ParseFILLOCKComplated(log types.Log) (*TokenExchangeFILLOCKComplated, error) {
	event := new(TokenExchangeFILLOCKComplated)
	if err := _TokenExchange.contract.UnpackLog(event, "FILLOCKComplated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeManagerWithdrawForFilComplatedIterator is returned from FilterManagerWithdrawForFilComplated and is used to iterate over the raw logs and unpacked data for ManagerWithdrawForFilComplated events raised by the TokenExchange contract.
type TokenExchangeManagerWithdrawForFilComplatedIterator struct {
	Event *TokenExchangeManagerWithdrawForFilComplated // Event containing the contract specifics and raw log

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
func (it *TokenExchangeManagerWithdrawForFilComplatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeManagerWithdrawForFilComplated)
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
		it.Event = new(TokenExchangeManagerWithdrawForFilComplated)
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
func (it *TokenExchangeManagerWithdrawForFilComplatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeManagerWithdrawForFilComplatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeManagerWithdrawForFilComplated represents a ManagerWithdrawForFilComplated event raised by the TokenExchange contract.
type TokenExchangeManagerWithdrawForFilComplated struct {
	Manager common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerWithdrawForFilComplated is a free log retrieval operation binding the contract event 0x2a53db73488d8ca9be11f0b704293b801cf928142fdba92aaedf7bc749b9eea7.
//
// Solidity: event ManagerWithdrawForFilComplated(address indexed manager, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) FilterManagerWithdrawForFilComplated(opts *bind.FilterOpts, manager []common.Address) (*TokenExchangeManagerWithdrawForFilComplatedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "ManagerWithdrawForFilComplated", managerRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeManagerWithdrawForFilComplatedIterator{contract: _TokenExchange.contract, event: "ManagerWithdrawForFilComplated", logs: logs, sub: sub}, nil
}

// WatchManagerWithdrawForFilComplated is a free log subscription operation binding the contract event 0x2a53db73488d8ca9be11f0b704293b801cf928142fdba92aaedf7bc749b9eea7.
//
// Solidity: event ManagerWithdrawForFilComplated(address indexed manager, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) WatchManagerWithdrawForFilComplated(opts *bind.WatchOpts, sink chan<- *TokenExchangeManagerWithdrawForFilComplated, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "ManagerWithdrawForFilComplated", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeManagerWithdrawForFilComplated)
				if err := _TokenExchange.contract.UnpackLog(event, "ManagerWithdrawForFilComplated", log); err != nil {
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

// ParseManagerWithdrawForFilComplated is a log parse operation binding the contract event 0x2a53db73488d8ca9be11f0b704293b801cf928142fdba92aaedf7bc749b9eea7.
//
// Solidity: event ManagerWithdrawForFilComplated(address indexed manager, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) ParseManagerWithdrawForFilComplated(log types.Log) (*TokenExchangeManagerWithdrawForFilComplated, error) {
	event := new(TokenExchangeManagerWithdrawForFilComplated)
	if err := _TokenExchange.contract.UnpackLog(event, "ManagerWithdrawForFilComplated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeREWARDFIL2DFILExchangedIterator is returned from FilterREWARDFIL2DFILExchanged and is used to iterate over the raw logs and unpacked data for REWARDFIL2DFILExchanged events raised by the TokenExchange contract.
type TokenExchangeREWARDFIL2DFILExchangedIterator struct {
	Event *TokenExchangeREWARDFIL2DFILExchanged // Event containing the contract specifics and raw log

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
func (it *TokenExchangeREWARDFIL2DFILExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeREWARDFIL2DFILExchanged)
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
		it.Event = new(TokenExchangeREWARDFIL2DFILExchanged)
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
func (it *TokenExchangeREWARDFIL2DFILExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeREWARDFIL2DFILExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeREWARDFIL2DFILExchanged represents a REWARDFIL2DFILExchanged event raised by the TokenExchange contract.
type TokenExchangeREWARDFIL2DFILExchanged struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterREWARDFIL2DFILExchanged is a free log retrieval operation binding the contract event 0x4edafc631ecd5454728b58c5303ead1d5995cb0680959777dfeb5844701d13a0.
//
// Solidity: event REWARDFIL2DFILExchanged(address indexed token, address owner, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) FilterREWARDFIL2DFILExchanged(opts *bind.FilterOpts, token []common.Address) (*TokenExchangeREWARDFIL2DFILExchangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "REWARDFIL2DFILExchanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeREWARDFIL2DFILExchangedIterator{contract: _TokenExchange.contract, event: "REWARDFIL2DFILExchanged", logs: logs, sub: sub}, nil
}

// WatchREWARDFIL2DFILExchanged is a free log subscription operation binding the contract event 0x4edafc631ecd5454728b58c5303ead1d5995cb0680959777dfeb5844701d13a0.
//
// Solidity: event REWARDFIL2DFILExchanged(address indexed token, address owner, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) WatchREWARDFIL2DFILExchanged(opts *bind.WatchOpts, sink chan<- *TokenExchangeREWARDFIL2DFILExchanged, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "REWARDFIL2DFILExchanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeREWARDFIL2DFILExchanged)
				if err := _TokenExchange.contract.UnpackLog(event, "REWARDFIL2DFILExchanged", log); err != nil {
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

// ParseREWARDFIL2DFILExchanged is a log parse operation binding the contract event 0x4edafc631ecd5454728b58c5303ead1d5995cb0680959777dfeb5844701d13a0.
//
// Solidity: event REWARDFIL2DFILExchanged(address indexed token, address owner, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) ParseREWARDFIL2DFILExchanged(log types.Log) (*TokenExchangeREWARDFIL2DFILExchanged, error) {
	event := new(TokenExchangeREWARDFIL2DFILExchanged)
	if err := _TokenExchange.contract.UnpackLog(event, "REWARDFIL2DFILExchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TokenExchange contract.
type TokenExchangeRoleAdminChangedIterator struct {
	Event *TokenExchangeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenExchangeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeRoleAdminChanged)
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
		it.Event = new(TokenExchangeRoleAdminChanged)
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
func (it *TokenExchangeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeRoleAdminChanged represents a RoleAdminChanged event raised by the TokenExchange contract.
type TokenExchangeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenExchange *TokenExchangeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenExchangeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeRoleAdminChangedIterator{contract: _TokenExchange.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenExchange *TokenExchangeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenExchangeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeRoleAdminChanged)
				if err := _TokenExchange.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TokenExchange *TokenExchangeFilterer) ParseRoleAdminChanged(log types.Log) (*TokenExchangeRoleAdminChanged, error) {
	event := new(TokenExchangeRoleAdminChanged)
	if err := _TokenExchange.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TokenExchange contract.
type TokenExchangeRoleGrantedIterator struct {
	Event *TokenExchangeRoleGranted // Event containing the contract specifics and raw log

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
func (it *TokenExchangeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeRoleGranted)
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
		it.Event = new(TokenExchangeRoleGranted)
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
func (it *TokenExchangeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeRoleGranted represents a RoleGranted event raised by the TokenExchange contract.
type TokenExchangeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenExchange *TokenExchangeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenExchangeRoleGrantedIterator, error) {

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

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeRoleGrantedIterator{contract: _TokenExchange.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenExchange *TokenExchangeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenExchangeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeRoleGranted)
				if err := _TokenExchange.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TokenExchange *TokenExchangeFilterer) ParseRoleGranted(log types.Log) (*TokenExchangeRoleGranted, error) {
	event := new(TokenExchangeRoleGranted)
	if err := _TokenExchange.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TokenExchange contract.
type TokenExchangeRoleRevokedIterator struct {
	Event *TokenExchangeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TokenExchangeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeRoleRevoked)
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
		it.Event = new(TokenExchangeRoleRevoked)
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
func (it *TokenExchangeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeRoleRevoked represents a RoleRevoked event raised by the TokenExchange contract.
type TokenExchangeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenExchange *TokenExchangeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenExchangeRoleRevokedIterator, error) {

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

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeRoleRevokedIterator{contract: _TokenExchange.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenExchange *TokenExchangeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenExchangeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeRoleRevoked)
				if err := _TokenExchange.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TokenExchange *TokenExchangeFilterer) ParseRoleRevoked(log types.Log) (*TokenExchangeRoleRevoked, error) {
	event := new(TokenExchangeRoleRevoked)
	if err := _TokenExchange.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeWITHDRAWFILComplatedIterator is returned from FilterWITHDRAWFILComplated and is used to iterate over the raw logs and unpacked data for WITHDRAWFILComplated events raised by the TokenExchange contract.
type TokenExchangeWITHDRAWFILComplatedIterator struct {
	Event *TokenExchangeWITHDRAWFILComplated // Event containing the contract specifics and raw log

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
func (it *TokenExchangeWITHDRAWFILComplatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeWITHDRAWFILComplated)
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
		it.Event = new(TokenExchangeWITHDRAWFILComplated)
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
func (it *TokenExchangeWITHDRAWFILComplatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeWITHDRAWFILComplatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeWITHDRAWFILComplated represents a WITHDRAWFILComplated event raised by the TokenExchange contract.
type TokenExchangeWITHDRAWFILComplated struct {
	Token        common.Address
	Owner        common.Address
	Amount       *big.Int
	UsePlatToken *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWITHDRAWFILComplated is a free log retrieval operation binding the contract event 0x91beb7d6e1ab6146da014b2298d01198caed9dfdc58df92ce700c91dbf2fd864.
//
// Solidity: event WITHDRAWFILComplated(address indexed token, address owner, uint256 amount, uint256 usePlatToken)
func (_TokenExchange *TokenExchangeFilterer) FilterWITHDRAWFILComplated(opts *bind.FilterOpts, token []common.Address) (*TokenExchangeWITHDRAWFILComplatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "WITHDRAWFILComplated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeWITHDRAWFILComplatedIterator{contract: _TokenExchange.contract, event: "WITHDRAWFILComplated", logs: logs, sub: sub}, nil
}

// WatchWITHDRAWFILComplated is a free log subscription operation binding the contract event 0x91beb7d6e1ab6146da014b2298d01198caed9dfdc58df92ce700c91dbf2fd864.
//
// Solidity: event WITHDRAWFILComplated(address indexed token, address owner, uint256 amount, uint256 usePlatToken)
func (_TokenExchange *TokenExchangeFilterer) WatchWITHDRAWFILComplated(opts *bind.WatchOpts, sink chan<- *TokenExchangeWITHDRAWFILComplated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "WITHDRAWFILComplated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeWITHDRAWFILComplated)
				if err := _TokenExchange.contract.UnpackLog(event, "WITHDRAWFILComplated", log); err != nil {
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

// ParseWITHDRAWFILComplated is a log parse operation binding the contract event 0x91beb7d6e1ab6146da014b2298d01198caed9dfdc58df92ce700c91dbf2fd864.
//
// Solidity: event WITHDRAWFILComplated(address indexed token, address owner, uint256 amount, uint256 usePlatToken)
func (_TokenExchange *TokenExchangeFilterer) ParseWITHDRAWFILComplated(log types.Log) (*TokenExchangeWITHDRAWFILComplated, error) {
	event := new(TokenExchangeWITHDRAWFILComplated)
	if err := _TokenExchange.contract.UnpackLog(event, "WITHDRAWFILComplated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangeFeeToWithdrawComplatedIterator is returned from FilterFeeToWithdrawComplated and is used to iterate over the raw logs and unpacked data for FeeToWithdrawComplated events raised by the TokenExchange contract.
type TokenExchangeFeeToWithdrawComplatedIterator struct {
	Event *TokenExchangeFeeToWithdrawComplated // Event containing the contract specifics and raw log

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
func (it *TokenExchangeFeeToWithdrawComplatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchangeFeeToWithdrawComplated)
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
		it.Event = new(TokenExchangeFeeToWithdrawComplated)
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
func (it *TokenExchangeFeeToWithdrawComplatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangeFeeToWithdrawComplatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchangeFeeToWithdrawComplated represents a FeeToWithdrawComplated event raised by the TokenExchange contract.
type TokenExchangeFeeToWithdrawComplated struct {
	FeeTo  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeToWithdrawComplated is a free log retrieval operation binding the contract event 0xb4d5c89cf78ccc37cae30709b65bc80f8ef3229d4820bf30bbda41e4583610f4.
//
// Solidity: event feeToWithdrawComplated(address indexed feeTo, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) FilterFeeToWithdrawComplated(opts *bind.FilterOpts, feeTo []common.Address) (*TokenExchangeFeeToWithdrawComplatedIterator, error) {

	var feeToRule []interface{}
	for _, feeToItem := range feeTo {
		feeToRule = append(feeToRule, feeToItem)
	}

	logs, sub, err := _TokenExchange.contract.FilterLogs(opts, "feeToWithdrawComplated", feeToRule)
	if err != nil {
		return nil, err
	}
	return &TokenExchangeFeeToWithdrawComplatedIterator{contract: _TokenExchange.contract, event: "feeToWithdrawComplated", logs: logs, sub: sub}, nil
}

// WatchFeeToWithdrawComplated is a free log subscription operation binding the contract event 0xb4d5c89cf78ccc37cae30709b65bc80f8ef3229d4820bf30bbda41e4583610f4.
//
// Solidity: event feeToWithdrawComplated(address indexed feeTo, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) WatchFeeToWithdrawComplated(opts *bind.WatchOpts, sink chan<- *TokenExchangeFeeToWithdrawComplated, feeTo []common.Address) (event.Subscription, error) {

	var feeToRule []interface{}
	for _, feeToItem := range feeTo {
		feeToRule = append(feeToRule, feeToItem)
	}

	logs, sub, err := _TokenExchange.contract.WatchLogs(opts, "feeToWithdrawComplated", feeToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchangeFeeToWithdrawComplated)
				if err := _TokenExchange.contract.UnpackLog(event, "feeToWithdrawComplated", log); err != nil {
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

// ParseFeeToWithdrawComplated is a log parse operation binding the contract event 0xb4d5c89cf78ccc37cae30709b65bc80f8ef3229d4820bf30bbda41e4583610f4.
//
// Solidity: event feeToWithdrawComplated(address indexed feeTo, uint256 amount)
func (_TokenExchange *TokenExchangeFilterer) ParseFeeToWithdrawComplated(log types.Log) (*TokenExchangeFeeToWithdrawComplated, error) {
	event := new(TokenExchangeFeeToWithdrawComplated)
	if err := _TokenExchange.contract.UnpackLog(event, "feeToWithdrawComplated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
