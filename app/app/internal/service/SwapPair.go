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

// SwapPairMetaData contains all meta data concerning the SwapPair contract.
var SwapPairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amounts\",\"type\":\"uint256\"}],\"name\":\"Rewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"allRewardsOfUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"callMe\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHistoryRewardRate\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHistoryStakingTime\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingFinishTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historyRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historyStakingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callMe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardPool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"contractIRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardUNIPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardUNI\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardRate_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingTime_\",\"type\":\"uint256\"}],\"name\":\"setRewardRateAndStakingFinishTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingFinishTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardsPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapPairABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapPairMetaData.ABI instead.
var SwapPairABI = SwapPairMetaData.ABI

// SwapPair is an auto generated Go binding around an Ethereum contract.
type SwapPair struct {
	SwapPairCaller     // Read-only binding to the contract
	SwapPairTransactor // Write-only binding to the contract
	SwapPairFilterer   // Log filterer for contract events
}

// SwapPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapPairSession struct {
	Contract     *SwapPair         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapPairCallerSession struct {
	Contract *SwapPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SwapPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapPairTransactorSession struct {
	Contract     *SwapPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SwapPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapPairRaw struct {
	Contract *SwapPair // Generic contract binding to access the raw methods on
}

// SwapPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapPairCallerRaw struct {
	Contract *SwapPairCaller // Generic read-only contract binding to access the raw methods on
}

// SwapPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapPairTransactorRaw struct {
	Contract *SwapPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapPair creates a new instance of SwapPair, bound to a specific deployed contract.
func NewSwapPair(address common.Address, backend bind.ContractBackend) (*SwapPair, error) {
	contract, err := bindSwapPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapPair{SwapPairCaller: SwapPairCaller{contract: contract}, SwapPairTransactor: SwapPairTransactor{contract: contract}, SwapPairFilterer: SwapPairFilterer{contract: contract}}, nil
}

// NewSwapPairCaller creates a new read-only instance of SwapPair, bound to a specific deployed contract.
func NewSwapPairCaller(address common.Address, caller bind.ContractCaller) (*SwapPairCaller, error) {
	contract, err := bindSwapPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapPairCaller{contract: contract}, nil
}

// NewSwapPairTransactor creates a new write-only instance of SwapPair, bound to a specific deployed contract.
func NewSwapPairTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapPairTransactor, error) {
	contract, err := bindSwapPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapPairTransactor{contract: contract}, nil
}

// NewSwapPairFilterer creates a new log filterer instance of SwapPair, bound to a specific deployed contract.
func NewSwapPairFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapPairFilterer, error) {
	contract, err := bindSwapPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapPairFilterer{contract: contract}, nil
}

// bindSwapPair binds a generic wrapper to an already deployed contract.
func bindSwapPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapPair *SwapPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapPair.Contract.SwapPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapPair *SwapPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.Contract.SwapPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapPair *SwapPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapPair.Contract.SwapPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapPair *SwapPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapPair *SwapPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapPair *SwapPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapPair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SwapPair *SwapPairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SwapPair *SwapPairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _SwapPair.Contract.DOMAINSEPARATOR(&_SwapPair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SwapPair *SwapPairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _SwapPair.Contract.DOMAINSEPARATOR(&_SwapPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_SwapPair *SwapPairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_SwapPair *SwapPairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _SwapPair.Contract.MINIMUMLIQUIDITY(&_SwapPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _SwapPair.Contract.MINIMUMLIQUIDITY(&_SwapPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_SwapPair *SwapPairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_SwapPair *SwapPairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _SwapPair.Contract.PERMITTYPEHASH(&_SwapPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_SwapPair *SwapPairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _SwapPair.Contract.PERMITTYPEHASH(&_SwapPair.CallOpts)
}

// AllRewardsOfUser is a free data retrieval call binding the contract method 0xe5b62296.
//
// Solidity: function allRewardsOfUser(address account) view returns(uint256)
func (_SwapPair *SwapPairCaller) AllRewardsOfUser(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "allRewardsOfUser", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllRewardsOfUser is a free data retrieval call binding the contract method 0xe5b62296.
//
// Solidity: function allRewardsOfUser(address account) view returns(uint256)
func (_SwapPair *SwapPairSession) AllRewardsOfUser(account common.Address) (*big.Int, error) {
	return _SwapPair.Contract.AllRewardsOfUser(&_SwapPair.CallOpts, account)
}

// AllRewardsOfUser is a free data retrieval call binding the contract method 0xe5b62296.
//
// Solidity: function allRewardsOfUser(address account) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) AllRewardsOfUser(account common.Address) (*big.Int, error) {
	return _SwapPair.Contract.AllRewardsOfUser(&_SwapPair.CallOpts, account)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SwapPair *SwapPairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SwapPair *SwapPairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Allowance(&_SwapPair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Allowance(&_SwapPair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SwapPair *SwapPairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SwapPair *SwapPairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.BalanceOf(&_SwapPair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.BalanceOf(&_SwapPair.CallOpts, arg0)
}

// CallMe is a free data retrieval call binding the contract method 0xb0bea725.
//
// Solidity: function callMe() view returns(address)
func (_SwapPair *SwapPairCaller) CallMe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "callMe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallMe is a free data retrieval call binding the contract method 0xb0bea725.
//
// Solidity: function callMe() view returns(address)
func (_SwapPair *SwapPairSession) CallMe() (common.Address, error) {
	return _SwapPair.Contract.CallMe(&_SwapPair.CallOpts)
}

// CallMe is a free data retrieval call binding the contract method 0xb0bea725.
//
// Solidity: function callMe() view returns(address)
func (_SwapPair *SwapPairCallerSession) CallMe() (common.Address, error) {
	return _SwapPair.Contract.CallMe(&_SwapPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SwapPair *SwapPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SwapPair *SwapPairSession) Decimals() (uint8, error) {
	return _SwapPair.Contract.Decimals(&_SwapPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SwapPair *SwapPairCallerSession) Decimals() (uint8, error) {
	return _SwapPair.Contract.Decimals(&_SwapPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapPair *SwapPairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapPair *SwapPairSession) Factory() (common.Address, error) {
	return _SwapPair.Contract.Factory(&_SwapPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapPair *SwapPairCallerSession) Factory() (common.Address, error) {
	return _SwapPair.Contract.Factory(&_SwapPair.CallOpts)
}

// GetHistoryRewardRate is a free data retrieval call binding the contract method 0xf54942c5.
//
// Solidity: function getHistoryRewardRate() view returns(uint256[])
func (_SwapPair *SwapPairCaller) GetHistoryRewardRate(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "getHistoryRewardRate")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetHistoryRewardRate is a free data retrieval call binding the contract method 0xf54942c5.
//
// Solidity: function getHistoryRewardRate() view returns(uint256[])
func (_SwapPair *SwapPairSession) GetHistoryRewardRate() ([]*big.Int, error) {
	return _SwapPair.Contract.GetHistoryRewardRate(&_SwapPair.CallOpts)
}

// GetHistoryRewardRate is a free data retrieval call binding the contract method 0xf54942c5.
//
// Solidity: function getHistoryRewardRate() view returns(uint256[])
func (_SwapPair *SwapPairCallerSession) GetHistoryRewardRate() ([]*big.Int, error) {
	return _SwapPair.Contract.GetHistoryRewardRate(&_SwapPair.CallOpts)
}

// GetHistoryStakingTime is a free data retrieval call binding the contract method 0xf392d818.
//
// Solidity: function getHistoryStakingTime() view returns(uint256[])
func (_SwapPair *SwapPairCaller) GetHistoryStakingTime(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "getHistoryStakingTime")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetHistoryStakingTime is a free data retrieval call binding the contract method 0xf392d818.
//
// Solidity: function getHistoryStakingTime() view returns(uint256[])
func (_SwapPair *SwapPairSession) GetHistoryStakingTime() ([]*big.Int, error) {
	return _SwapPair.Contract.GetHistoryStakingTime(&_SwapPair.CallOpts)
}

// GetHistoryStakingTime is a free data retrieval call binding the contract method 0xf392d818.
//
// Solidity: function getHistoryStakingTime() view returns(uint256[])
func (_SwapPair *SwapPairCallerSession) GetHistoryStakingTime() ([]*big.Int, error) {
	return _SwapPair.Contract.GetHistoryStakingTime(&_SwapPair.CallOpts)
}

// GetLastTime is a free data retrieval call binding the contract method 0x955d14cd.
//
// Solidity: function getLastTime() view returns(uint256)
func (_SwapPair *SwapPairCaller) GetLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "getLastTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastTime is a free data retrieval call binding the contract method 0x955d14cd.
//
// Solidity: function getLastTime() view returns(uint256)
func (_SwapPair *SwapPairSession) GetLastTime() (*big.Int, error) {
	return _SwapPair.Contract.GetLastTime(&_SwapPair.CallOpts)
}

// GetLastTime is a free data retrieval call binding the contract method 0x955d14cd.
//
// Solidity: function getLastTime() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) GetLastTime() (*big.Int, error) {
	return _SwapPair.Contract.GetLastTime(&_SwapPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_SwapPair *SwapPairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_SwapPair *SwapPairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _SwapPair.Contract.GetReserves(&_SwapPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_SwapPair *SwapPairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _SwapPair.Contract.GetReserves(&_SwapPair.CallOpts)
}

// GetStakingFinishTime is a free data retrieval call binding the contract method 0x51b89b8f.
//
// Solidity: function getStakingFinishTime() view returns(uint256)
func (_SwapPair *SwapPairCaller) GetStakingFinishTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "getStakingFinishTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingFinishTime is a free data retrieval call binding the contract method 0x51b89b8f.
//
// Solidity: function getStakingFinishTime() view returns(uint256)
func (_SwapPair *SwapPairSession) GetStakingFinishTime() (*big.Int, error) {
	return _SwapPair.Contract.GetStakingFinishTime(&_SwapPair.CallOpts)
}

// GetStakingFinishTime is a free data retrieval call binding the contract method 0x51b89b8f.
//
// Solidity: function getStakingFinishTime() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) GetStakingFinishTime() (*big.Int, error) {
	return _SwapPair.Contract.GetStakingFinishTime(&_SwapPair.CallOpts)
}

// HistoryRewardRate is a free data retrieval call binding the contract method 0xe60baa90.
//
// Solidity: function historyRewardRate(uint256 ) view returns(uint256)
func (_SwapPair *SwapPairCaller) HistoryRewardRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "historyRewardRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoryRewardRate is a free data retrieval call binding the contract method 0xe60baa90.
//
// Solidity: function historyRewardRate(uint256 ) view returns(uint256)
func (_SwapPair *SwapPairSession) HistoryRewardRate(arg0 *big.Int) (*big.Int, error) {
	return _SwapPair.Contract.HistoryRewardRate(&_SwapPair.CallOpts, arg0)
}

// HistoryRewardRate is a free data retrieval call binding the contract method 0xe60baa90.
//
// Solidity: function historyRewardRate(uint256 ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) HistoryRewardRate(arg0 *big.Int) (*big.Int, error) {
	return _SwapPair.Contract.HistoryRewardRate(&_SwapPair.CallOpts, arg0)
}

// HistoryStakingTime is a free data retrieval call binding the contract method 0x91b22645.
//
// Solidity: function historyStakingTime(uint256 ) view returns(uint256)
func (_SwapPair *SwapPairCaller) HistoryStakingTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "historyStakingTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoryStakingTime is a free data retrieval call binding the contract method 0x91b22645.
//
// Solidity: function historyStakingTime(uint256 ) view returns(uint256)
func (_SwapPair *SwapPairSession) HistoryStakingTime(arg0 *big.Int) (*big.Int, error) {
	return _SwapPair.Contract.HistoryStakingTime(&_SwapPair.CallOpts, arg0)
}

// HistoryStakingTime is a free data retrieval call binding the contract method 0x91b22645.
//
// Solidity: function historyStakingTime(uint256 ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) HistoryStakingTime(arg0 *big.Int) (*big.Int, error) {
	return _SwapPair.Contract.HistoryStakingTime(&_SwapPair.CallOpts, arg0)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_SwapPair *SwapPairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_SwapPair *SwapPairSession) KLast() (*big.Int, error) {
	return _SwapPair.Contract.KLast(&_SwapPair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) KLast() (*big.Int, error) {
	return _SwapPair.Contract.KLast(&_SwapPair.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_SwapPair *SwapPairCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_SwapPair *SwapPairSession) LastUpdateTime() (*big.Int, error) {
	return _SwapPair.Contract.LastUpdateTime(&_SwapPair.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) LastUpdateTime() (*big.Int, error) {
	return _SwapPair.Contract.LastUpdateTime(&_SwapPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SwapPair *SwapPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SwapPair *SwapPairSession) Name() (string, error) {
	return _SwapPair.Contract.Name(&_SwapPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SwapPair *SwapPairCallerSession) Name() (string, error) {
	return _SwapPair.Contract.Name(&_SwapPair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SwapPair *SwapPairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SwapPair *SwapPairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Nonces(&_SwapPair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Nonces(&_SwapPair.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_SwapPair *SwapPairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_SwapPair *SwapPairSession) Price0CumulativeLast() (*big.Int, error) {
	return _SwapPair.Contract.Price0CumulativeLast(&_SwapPair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _SwapPair.Contract.Price0CumulativeLast(&_SwapPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_SwapPair *SwapPairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_SwapPair *SwapPairSession) Price1CumulativeLast() (*big.Int, error) {
	return _SwapPair.Contract.Price1CumulativeLast(&_SwapPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _SwapPair.Contract.Price1CumulativeLast(&_SwapPair.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_SwapPair *SwapPairCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_SwapPair *SwapPairSession) RewardPerTokenStored() (*big.Int, error) {
	return _SwapPair.Contract.RewardPerTokenStored(&_SwapPair.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _SwapPair.Contract.RewardPerTokenStored(&_SwapPair.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_SwapPair *SwapPairCaller) RewardPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_SwapPair *SwapPairSession) RewardPool() (common.Address, error) {
	return _SwapPair.Contract.RewardPool(&_SwapPair.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_SwapPair *SwapPairCallerSession) RewardPool() (common.Address, error) {
	return _SwapPair.Contract.RewardPool(&_SwapPair.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_SwapPair *SwapPairCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_SwapPair *SwapPairSession) RewardRate() (*big.Int, error) {
	return _SwapPair.Contract.RewardRate(&_SwapPair.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) RewardRate() (*big.Int, error) {
	return _SwapPair.Contract.RewardRate(&_SwapPair.CallOpts)
}

// RewardUNIPerToken is a free data retrieval call binding the contract method 0x544ec62d.
//
// Solidity: function rewardUNIPerToken() view returns(uint256 rewardUNI)
func (_SwapPair *SwapPairCaller) RewardUNIPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "rewardUNIPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardUNIPerToken is a free data retrieval call binding the contract method 0x544ec62d.
//
// Solidity: function rewardUNIPerToken() view returns(uint256 rewardUNI)
func (_SwapPair *SwapPairSession) RewardUNIPerToken() (*big.Int, error) {
	return _SwapPair.Contract.RewardUNIPerToken(&_SwapPair.CallOpts)
}

// RewardUNIPerToken is a free data retrieval call binding the contract method 0x544ec62d.
//
// Solidity: function rewardUNIPerToken() view returns(uint256 rewardUNI)
func (_SwapPair *SwapPairCallerSession) RewardUNIPerToken() (*big.Int, error) {
	return _SwapPair.Contract.RewardUNIPerToken(&_SwapPair.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_SwapPair *SwapPairCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_SwapPair *SwapPairSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Rewards(&_SwapPair.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Rewards(&_SwapPair.CallOpts, arg0)
}

// StakingFinishTime is a free data retrieval call binding the contract method 0x6d3c09d8.
//
// Solidity: function stakingFinishTime() view returns(uint256)
func (_SwapPair *SwapPairCaller) StakingFinishTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "stakingFinishTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingFinishTime is a free data retrieval call binding the contract method 0x6d3c09d8.
//
// Solidity: function stakingFinishTime() view returns(uint256)
func (_SwapPair *SwapPairSession) StakingFinishTime() (*big.Int, error) {
	return _SwapPair.Contract.StakingFinishTime(&_SwapPair.CallOpts)
}

// StakingFinishTime is a free data retrieval call binding the contract method 0x6d3c09d8.
//
// Solidity: function stakingFinishTime() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) StakingFinishTime() (*big.Int, error) {
	return _SwapPair.Contract.StakingFinishTime(&_SwapPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SwapPair *SwapPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SwapPair *SwapPairSession) Symbol() (string, error) {
	return _SwapPair.Contract.Symbol(&_SwapPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SwapPair *SwapPairCallerSession) Symbol() (string, error) {
	return _SwapPair.Contract.Symbol(&_SwapPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SwapPair *SwapPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SwapPair *SwapPairSession) Token0() (common.Address, error) {
	return _SwapPair.Contract.Token0(&_SwapPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SwapPair *SwapPairCallerSession) Token0() (common.Address, error) {
	return _SwapPair.Contract.Token0(&_SwapPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SwapPair *SwapPairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SwapPair *SwapPairSession) Token1() (common.Address, error) {
	return _SwapPair.Contract.Token1(&_SwapPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SwapPair *SwapPairCallerSession) Token1() (common.Address, error) {
	return _SwapPair.Contract.Token1(&_SwapPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SwapPair *SwapPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SwapPair *SwapPairSession) TotalSupply() (*big.Int, error) {
	return _SwapPair.Contract.TotalSupply(&_SwapPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SwapPair *SwapPairCallerSession) TotalSupply() (*big.Int, error) {
	return _SwapPair.Contract.TotalSupply(&_SwapPair.CallOpts)
}

// UserRewardsPerToken is a free data retrieval call binding the contract method 0x09ed4db7.
//
// Solidity: function userRewardsPerToken(address ) view returns(uint256)
func (_SwapPair *SwapPairCaller) UserRewardsPerToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapPair.contract.Call(opts, &out, "userRewardsPerToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardsPerToken is a free data retrieval call binding the contract method 0x09ed4db7.
//
// Solidity: function userRewardsPerToken(address ) view returns(uint256)
func (_SwapPair *SwapPairSession) UserRewardsPerToken(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.UserRewardsPerToken(&_SwapPair.CallOpts, arg0)
}

// UserRewardsPerToken is a free data retrieval call binding the contract method 0x09ed4db7.
//
// Solidity: function userRewardsPerToken(address ) view returns(uint256)
func (_SwapPair *SwapPairCallerSession) UserRewardsPerToken(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.UserRewardsPerToken(&_SwapPair.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Approve(&_SwapPair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Approve(&_SwapPair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_SwapPair *SwapPairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_SwapPair *SwapPairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Burn(&_SwapPair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_SwapPair *SwapPairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Burn(&_SwapPair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _token0, address _token1, address _callMe, address _rewardPool) returns()
func (_SwapPair *SwapPairTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address, _callMe common.Address, _rewardPool common.Address) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "initialize", _token0, _token1, _callMe, _rewardPool)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _token0, address _token1, address _callMe, address _rewardPool) returns()
func (_SwapPair *SwapPairSession) Initialize(_token0 common.Address, _token1 common.Address, _callMe common.Address, _rewardPool common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Initialize(&_SwapPair.TransactOpts, _token0, _token1, _callMe, _rewardPool)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _token0, address _token1, address _callMe, address _rewardPool) returns()
func (_SwapPair *SwapPairTransactorSession) Initialize(_token0 common.Address, _token1 common.Address, _callMe common.Address, _rewardPool common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Initialize(&_SwapPair.TransactOpts, _token0, _token1, _callMe, _rewardPool)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_SwapPair *SwapPairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_SwapPair *SwapPairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Mint(&_SwapPair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_SwapPair *SwapPairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Mint(&_SwapPair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SwapPair *SwapPairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SwapPair *SwapPairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapPair.Contract.Permit(&_SwapPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SwapPair *SwapPairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapPair.Contract.Permit(&_SwapPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_SwapPair *SwapPairTransactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_SwapPair *SwapPairSession) Reward() (*types.Transaction, error) {
	return _SwapPair.Contract.Reward(&_SwapPair.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_SwapPair *SwapPairTransactorSession) Reward() (*types.Transaction, error) {
	return _SwapPair.Contract.Reward(&_SwapPair.TransactOpts)
}

// SetRewardRateAndStakingFinishTime is a paid mutator transaction binding the contract method 0x2c8d6c10.
//
// Solidity: function setRewardRateAndStakingFinishTime(uint256 rewardRate_, uint256 stakingTime_) returns()
func (_SwapPair *SwapPairTransactor) SetRewardRateAndStakingFinishTime(opts *bind.TransactOpts, rewardRate_ *big.Int, stakingTime_ *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "setRewardRateAndStakingFinishTime", rewardRate_, stakingTime_)
}

// SetRewardRateAndStakingFinishTime is a paid mutator transaction binding the contract method 0x2c8d6c10.
//
// Solidity: function setRewardRateAndStakingFinishTime(uint256 rewardRate_, uint256 stakingTime_) returns()
func (_SwapPair *SwapPairSession) SetRewardRateAndStakingFinishTime(rewardRate_ *big.Int, stakingTime_ *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.SetRewardRateAndStakingFinishTime(&_SwapPair.TransactOpts, rewardRate_, stakingTime_)
}

// SetRewardRateAndStakingFinishTime is a paid mutator transaction binding the contract method 0x2c8d6c10.
//
// Solidity: function setRewardRateAndStakingFinishTime(uint256 rewardRate_, uint256 stakingTime_) returns()
func (_SwapPair *SwapPairTransactorSession) SetRewardRateAndStakingFinishTime(rewardRate_ *big.Int, stakingTime_ *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.SetRewardRateAndStakingFinishTime(&_SwapPair.TransactOpts, rewardRate_, stakingTime_)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_SwapPair *SwapPairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_SwapPair *SwapPairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Skim(&_SwapPair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_SwapPair *SwapPairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.Skim(&_SwapPair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_SwapPair *SwapPairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_SwapPair *SwapPairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _SwapPair.Contract.Swap(&_SwapPair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_SwapPair *SwapPairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _SwapPair.Contract.Swap(&_SwapPair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_SwapPair *SwapPairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_SwapPair *SwapPairSession) Sync() (*types.Transaction, error) {
	return _SwapPair.Contract.Sync(&_SwapPair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_SwapPair *SwapPairTransactorSession) Sync() (*types.Transaction, error) {
	return _SwapPair.Contract.Sync(&_SwapPair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Transfer(&_SwapPair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Transfer(&_SwapPair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.TransferFrom(&_SwapPair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.TransferFrom(&_SwapPair.TransactOpts, from, to, value)
}

// SwapPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SwapPair contract.
type SwapPairApprovalIterator struct {
	Event *SwapPairApproval // Event containing the contract specifics and raw log

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
func (it *SwapPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairApproval)
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
		it.Event = new(SwapPairApproval)
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
func (it *SwapPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairApproval represents a Approval event raised by the SwapPair contract.
type SwapPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SwapPair *SwapPairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SwapPairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairApprovalIterator{contract: _SwapPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SwapPair *SwapPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SwapPairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairApproval)
				if err := _SwapPair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SwapPair *SwapPairFilterer) ParseApproval(log types.Log) (*SwapPairApproval, error) {
	event := new(SwapPairApproval)
	if err := _SwapPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the SwapPair contract.
type SwapPairBurnIterator struct {
	Event *SwapPairBurn // Event containing the contract specifics and raw log

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
func (it *SwapPairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairBurn)
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
		it.Event = new(SwapPairBurn)
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
func (it *SwapPairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairBurn represents a Burn event raised by the SwapPair contract.
type SwapPairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_SwapPair *SwapPairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SwapPairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairBurnIterator{contract: _SwapPair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_SwapPair *SwapPairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *SwapPairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairBurn)
				if err := _SwapPair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_SwapPair *SwapPairFilterer) ParseBurn(log types.Log) (*SwapPairBurn, error) {
	event := new(SwapPairBurn)
	if err := _SwapPair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the SwapPair contract.
type SwapPairMintIterator struct {
	Event *SwapPairMint // Event containing the contract specifics and raw log

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
func (it *SwapPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairMint)
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
		it.Event = new(SwapPairMint)
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
func (it *SwapPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairMint represents a Mint event raised by the SwapPair contract.
type SwapPairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_SwapPair *SwapPairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*SwapPairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairMintIterator{contract: _SwapPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_SwapPair *SwapPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SwapPairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairMint)
				if err := _SwapPair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_SwapPair *SwapPairFilterer) ParseMint(log types.Log) (*SwapPairMint, error) {
	event := new(SwapPairMint)
	if err := _SwapPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPairRewardsIterator is returned from FilterRewards and is used to iterate over the raw logs and unpacked data for Rewards events raised by the SwapPair contract.
type SwapPairRewardsIterator struct {
	Event *SwapPairRewards // Event containing the contract specifics and raw log

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
func (it *SwapPairRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairRewards)
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
		it.Event = new(SwapPairRewards)
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
func (it *SwapPairRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairRewards represents a Rewards event raised by the SwapPair contract.
type SwapPairRewards struct {
	Recipient common.Address
	Amounts   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewards is a free log retrieval operation binding the contract event 0xc083a1647e3ee591bf42b82564ffb4d16fdbb26068f0080da911c8d8300fd84a.
//
// Solidity: event Rewards(address recipient, uint256 amounts)
func (_SwapPair *SwapPairFilterer) FilterRewards(opts *bind.FilterOpts) (*SwapPairRewardsIterator, error) {

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Rewards")
	if err != nil {
		return nil, err
	}
	return &SwapPairRewardsIterator{contract: _SwapPair.contract, event: "Rewards", logs: logs, sub: sub}, nil
}

// WatchRewards is a free log subscription operation binding the contract event 0xc083a1647e3ee591bf42b82564ffb4d16fdbb26068f0080da911c8d8300fd84a.
//
// Solidity: event Rewards(address recipient, uint256 amounts)
func (_SwapPair *SwapPairFilterer) WatchRewards(opts *bind.WatchOpts, sink chan<- *SwapPairRewards) (event.Subscription, error) {

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Rewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairRewards)
				if err := _SwapPair.contract.UnpackLog(event, "Rewards", log); err != nil {
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

// ParseRewards is a log parse operation binding the contract event 0xc083a1647e3ee591bf42b82564ffb4d16fdbb26068f0080da911c8d8300fd84a.
//
// Solidity: event Rewards(address recipient, uint256 amounts)
func (_SwapPair *SwapPairFilterer) ParseRewards(log types.Log) (*SwapPairRewards, error) {
	event := new(SwapPairRewards)
	if err := _SwapPair.contract.UnpackLog(event, "Rewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapPair contract.
type SwapPairSwapIterator struct {
	Event *SwapPairSwap // Event containing the contract specifics and raw log

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
func (it *SwapPairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairSwap)
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
		it.Event = new(SwapPairSwap)
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
func (it *SwapPairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairSwap represents a Swap event raised by the SwapPair contract.
type SwapPairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SwapPair *SwapPairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SwapPairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairSwapIterator{contract: _SwapPair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SwapPair *SwapPairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapPairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairSwap)
				if err := _SwapPair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SwapPair *SwapPairFilterer) ParseSwap(log types.Log) (*SwapPairSwap, error) {
	event := new(SwapPairSwap)
	if err := _SwapPair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the SwapPair contract.
type SwapPairSyncIterator struct {
	Event *SwapPairSync // Event containing the contract specifics and raw log

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
func (it *SwapPairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairSync)
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
		it.Event = new(SwapPairSync)
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
func (it *SwapPairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairSync represents a Sync event raised by the SwapPair contract.
type SwapPairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_SwapPair *SwapPairFilterer) FilterSync(opts *bind.FilterOpts) (*SwapPairSyncIterator, error) {

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &SwapPairSyncIterator{contract: _SwapPair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_SwapPair *SwapPairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *SwapPairSync) (event.Subscription, error) {

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairSync)
				if err := _SwapPair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_SwapPair *SwapPairFilterer) ParseSync(log types.Log) (*SwapPairSync, error) {
	event := new(SwapPairSync)
	if err := _SwapPair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SwapPair contract.
type SwapPairTransferIterator struct {
	Event *SwapPairTransfer // Event containing the contract specifics and raw log

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
func (it *SwapPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairTransfer)
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
		it.Event = new(SwapPairTransfer)
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
func (it *SwapPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairTransfer represents a Transfer event raised by the SwapPair contract.
type SwapPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SwapPair *SwapPairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SwapPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairTransferIterator{contract: _SwapPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SwapPair *SwapPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SwapPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairTransfer)
				if err := _SwapPair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SwapPair *SwapPairFilterer) ParseTransfer(log types.Log) (*SwapPairTransfer, error) {
	event := new(SwapPairTransfer)
	if err := _SwapPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
