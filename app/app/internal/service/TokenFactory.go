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

// IFilecoinMinerControllerTemplateCheckData is an auto generated low-level Go binding around an user-defined struct.
type IFilecoinMinerControllerTemplateCheckData struct {
	CostRatePerToken   *big.Int
	CostBasePerToken   *big.Int
	ProfitRatePerToken *big.Int
	ProfitBasePerToken *big.Int
	StakeType          *big.Int
	StakeTypeRate      *big.Int
	StakeTypeBase      *big.Int
	RewardOwnerRate    *big.Int
	RewardRate         *big.Int
	RewardBase         *big.Int
	RewardBankRate     *big.Int
	TimeType           *big.Int
	KeyRatePerT        *big.Int
	KeyBasePerT        *big.Int
}

// IFilecoinMinerControllerTemplateCreateControllerData is an auto generated low-level Go binding around an user-defined struct.
type IFilecoinMinerControllerTemplateCreateControllerData struct {
	Owner           common.Address
	Actor           uint64
	Due             *big.Int
	Union           bool
	StakeType       *big.Int
	StakeTypeRate   *big.Int
	StakeTypeBase   *big.Int
	RewardOwnerRate *big.Int
	RewardRate      *big.Int
	RewardBankRate  *big.Int
	RewardBase      *big.Int
	ExtraTime       *big.Int
	KeyRatePerT     *big.Int
	KeyBasePerT     *big.Int
}

// TokenFactoryMetaData contains all meta data concerning the TokenFactory contract.
var TokenFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dfil_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filecoinMinerControllerTemplate_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filecoinMinerTemplate_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTemplate_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenUnionTemplate_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bankRewardDfil_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"top_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defaultAdminRoleLimit_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"ActorMinerControllerCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CheckOwnerCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"OwnerAcotrMinersCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"TopAcotrMinersCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnionTokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALL_PAIR_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CREATE_TOKEN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"acotrMinerController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bankRewardDfil\",\"outputs\":[{\"internalType\":\"contractITokenBankReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllerCheckData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"costRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardOwnerRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBankRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyRatePerT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyBasePerT\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"due\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"union\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardOwnerRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBankRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyRatePerT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyBasePerT\",\"type\":\"uint256\"}],\"internalType\":\"structIFilecoinMinerControllerTemplate.CreateControllerData\",\"name\":\"createControllerData\",\"type\":\"tuple\"}],\"name\":\"createActorMinerController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"filecoinControllerMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filecoinMiner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createCheckOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pledge\",\"type\":\"uint256\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createTokenOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"costRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeType\",\"type\":\"uint256\"}],\"name\":\"createTokenOwnerAcotrMiners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"costRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeType\",\"type\":\"uint256\"}],\"name\":\"createTopAcotrMiners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pledge\",\"type\":\"uint256\"}],\"name\":\"createUnionToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminRoleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dfil\",\"outputs\":[{\"internalType\":\"contractIDFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"existsCheckOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"existsOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"}],\"name\":\"existsOwnerActorMiner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"existsToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"existsTopUnionToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"existsUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"filecoinMinerControllerTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"filecoinMinerTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCheckOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getCheckOwnersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLowUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getOwnerAcotrMinersByIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"}],\"name\":\"getOwnerActorCurrentControllers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnerActorMiners\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnerUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getOwnersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"}],\"name\":\"getTokenOwnerActorMinersCheckData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"costRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardOwnerRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBankRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyRatePerT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyBasePerT\",\"type\":\"uint256\"}],\"internalType\":\"structIFilecoinMinerControllerTemplate.CheckData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getTokenOwnerTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getTokenOwnerTokensByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getTokenOwnerUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getTokensByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTop\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getTopAcotrMinersByIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopActorMiners\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"}],\"name\":\"getTopActorMinersCheckData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"costRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitRatePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitBasePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTypeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardOwnerRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBankRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyRatePerT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyBasePerT\",\"type\":\"uint256\"}],\"internalType\":\"structIFilecoinMinerControllerTemplate.CheckData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopUnionTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserOwnerByAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUserWithLowUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idoStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"internalType\":\"contractIKey\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"ownerActorCurrentControllers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeCheckOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"}],\"name\":\"removeTokenOwnerAcotrMiners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"}],\"name\":\"removeTopAcotrMiners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"returnPledge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"rewardController\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"actor\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"setAcotrMinerController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callPair_\",\"type\":\"address\"}],\"name\":\"setCallPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin_\",\"type\":\"address\"}],\"name\":\"setDefaultAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"}],\"name\":\"setF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"}],\"name\":\"setFM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"setIdoStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"setRewardStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapFactory_\",\"type\":\"address\"}],\"name\":\"setSwapFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenExchange_\",\"type\":\"address\"}],\"name\":\"setTokenExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setTu\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recommendAccount\",\"type\":\"address\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenExchange\",\"outputs\":[{\"internalType\":\"contractITokenExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenUnionTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"top\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnerFilecoinMinerController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userWithTokenOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userWithTopUsers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usersAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawFilecoinMinerController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenFactoryMetaData.ABI instead.
var TokenFactoryABI = TokenFactoryMetaData.ABI

// TokenFactory is an auto generated Go binding around an Ethereum contract.
type TokenFactory struct {
	TokenFactoryCaller     // Read-only binding to the contract
	TokenFactoryTransactor // Write-only binding to the contract
	TokenFactoryFilterer   // Log filterer for contract events
}

// TokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFactorySession struct {
	Contract     *TokenFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFactoryCallerSession struct {
	Contract *TokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFactoryTransactorSession struct {
	Contract     *TokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFactoryRaw struct {
	Contract *TokenFactory // Generic contract binding to access the raw methods on
}

// TokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFactoryCallerRaw struct {
	Contract *TokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFactoryTransactorRaw struct {
	Contract *TokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFactory creates a new instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactory(address common.Address, backend bind.ContractBackend) (*TokenFactory, error) {
	contract, err := bindTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// NewTokenFactoryCaller creates a new read-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenFactoryCaller, error) {
	contract, err := bindTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryCaller{contract: contract}, nil
}

// NewTokenFactoryTransactor creates a new write-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFactoryTransactor, error) {
	contract, err := bindTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTransactor{contract: contract}, nil
}

// NewTokenFactoryFilterer creates a new log filterer instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFactoryFilterer, error) {
	contract, err := bindTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryFilterer{contract: contract}, nil
}

// bindTokenFactory binds a generic wrapper to an already deployed contract.
func bindTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.TokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactorySession) ADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.ADMINROLE(&_TokenFactory.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) ADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.ADMINROLE(&_TokenFactory.CallOpts)
}

// CALLPAIRSETTERROLE is a free data retrieval call binding the contract method 0x1763bbe9.
//
// Solidity: function CALL_PAIR_SETTER_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) CALLPAIRSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "CALL_PAIR_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CALLPAIRSETTERROLE is a free data retrieval call binding the contract method 0x1763bbe9.
//
// Solidity: function CALL_PAIR_SETTER_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactorySession) CALLPAIRSETTERROLE() ([32]byte, error) {
	return _TokenFactory.Contract.CALLPAIRSETTERROLE(&_TokenFactory.CallOpts)
}

// CALLPAIRSETTERROLE is a free data retrieval call binding the contract method 0x1763bbe9.
//
// Solidity: function CALL_PAIR_SETTER_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) CALLPAIRSETTERROLE() ([32]byte, error) {
	return _TokenFactory.Contract.CALLPAIRSETTERROLE(&_TokenFactory.CallOpts)
}

// CREATETOKENROLE is a free data retrieval call binding the contract method 0xb4d8c2a4.
//
// Solidity: function CREATE_TOKEN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) CREATETOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "CREATE_TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATETOKENROLE is a free data retrieval call binding the contract method 0xb4d8c2a4.
//
// Solidity: function CREATE_TOKEN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactorySession) CREATETOKENROLE() ([32]byte, error) {
	return _TokenFactory.Contract.CREATETOKENROLE(&_TokenFactory.CallOpts)
}

// CREATETOKENROLE is a free data retrieval call binding the contract method 0xb4d8c2a4.
//
// Solidity: function CREATE_TOKEN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) CREATETOKENROLE() ([32]byte, error) {
	return _TokenFactory.Contract.CREATETOKENROLE(&_TokenFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.DEFAULTADMINROLE(&_TokenFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.DEFAULTADMINROLE(&_TokenFactory.CallOpts)
}

// REWARDADMINROLE is a free data retrieval call binding the contract method 0x65db6818.
//
// Solidity: function REWARD_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) REWARDADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "REWARD_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDADMINROLE is a free data retrieval call binding the contract method 0x65db6818.
//
// Solidity: function REWARD_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactorySession) REWARDADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.REWARDADMINROLE(&_TokenFactory.CallOpts)
}

// REWARDADMINROLE is a free data retrieval call binding the contract method 0x65db6818.
//
// Solidity: function REWARD_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) REWARDADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.REWARDADMINROLE(&_TokenFactory.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) SUPERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "SUPER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactorySession) SUPERADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.SUPERADMINROLE(&_TokenFactory.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) SUPERADMINROLE() ([32]byte, error) {
	return _TokenFactory.Contract.SUPERADMINROLE(&_TokenFactory.CallOpts)
}

// AcotrMinerController is a free data retrieval call binding the contract method 0xca01981a.
//
// Solidity: function acotrMinerController(uint64 ) view returns(address)
func (_TokenFactory *TokenFactoryCaller) AcotrMinerController(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "acotrMinerController", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcotrMinerController is a free data retrieval call binding the contract method 0xca01981a.
//
// Solidity: function acotrMinerController(uint64 ) view returns(address)
func (_TokenFactory *TokenFactorySession) AcotrMinerController(arg0 uint64) (common.Address, error) {
	return _TokenFactory.Contract.AcotrMinerController(&_TokenFactory.CallOpts, arg0)
}

// AcotrMinerController is a free data retrieval call binding the contract method 0xca01981a.
//
// Solidity: function acotrMinerController(uint64 ) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) AcotrMinerController(arg0 uint64) (common.Address, error) {
	return _TokenFactory.Contract.AcotrMinerController(&_TokenFactory.CallOpts, arg0)
}

// BankRewardDfil is a free data retrieval call binding the contract method 0xc3d2f25b.
//
// Solidity: function bankRewardDfil() view returns(address)
func (_TokenFactory *TokenFactoryCaller) BankRewardDfil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "bankRewardDfil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BankRewardDfil is a free data retrieval call binding the contract method 0xc3d2f25b.
//
// Solidity: function bankRewardDfil() view returns(address)
func (_TokenFactory *TokenFactorySession) BankRewardDfil() (common.Address, error) {
	return _TokenFactory.Contract.BankRewardDfil(&_TokenFactory.CallOpts)
}

// BankRewardDfil is a free data retrieval call binding the contract method 0xc3d2f25b.
//
// Solidity: function bankRewardDfil() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) BankRewardDfil() (common.Address, error) {
	return _TokenFactory.Contract.BankRewardDfil(&_TokenFactory.CallOpts)
}

// CallPair is a free data retrieval call binding the contract method 0xbcb6fe88.
//
// Solidity: function callPair() view returns(address)
func (_TokenFactory *TokenFactoryCaller) CallPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "callPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallPair is a free data retrieval call binding the contract method 0xbcb6fe88.
//
// Solidity: function callPair() view returns(address)
func (_TokenFactory *TokenFactorySession) CallPair() (common.Address, error) {
	return _TokenFactory.Contract.CallPair(&_TokenFactory.CallOpts)
}

// CallPair is a free data retrieval call binding the contract method 0xbcb6fe88.
//
// Solidity: function callPair() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) CallPair() (common.Address, error) {
	return _TokenFactory.Contract.CallPair(&_TokenFactory.CallOpts)
}

// ControllerCheckData is a free data retrieval call binding the contract method 0x937c16a0.
//
// Solidity: function controllerCheckData(address ) view returns(uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 stakeType, uint256 stakeTypeRate, uint256 stakeTypeBase, uint256 rewardOwnerRate, uint256 rewardRate, uint256 rewardBase, uint256 rewardBankRate, uint256 timeType, uint256 keyRatePerT, uint256 keyBasePerT)
func (_TokenFactory *TokenFactoryCaller) ControllerCheckData(opts *bind.CallOpts, arg0 common.Address) (struct {
	CostRatePerToken   *big.Int
	CostBasePerToken   *big.Int
	ProfitRatePerToken *big.Int
	ProfitBasePerToken *big.Int
	StakeType          *big.Int
	StakeTypeRate      *big.Int
	StakeTypeBase      *big.Int
	RewardOwnerRate    *big.Int
	RewardRate         *big.Int
	RewardBase         *big.Int
	RewardBankRate     *big.Int
	TimeType           *big.Int
	KeyRatePerT        *big.Int
	KeyBasePerT        *big.Int
}, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "controllerCheckData", arg0)

	outstruct := new(struct {
		CostRatePerToken   *big.Int
		CostBasePerToken   *big.Int
		ProfitRatePerToken *big.Int
		ProfitBasePerToken *big.Int
		StakeType          *big.Int
		StakeTypeRate      *big.Int
		StakeTypeBase      *big.Int
		RewardOwnerRate    *big.Int
		RewardRate         *big.Int
		RewardBase         *big.Int
		RewardBankRate     *big.Int
		TimeType           *big.Int
		KeyRatePerT        *big.Int
		KeyBasePerT        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CostRatePerToken = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CostBasePerToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProfitRatePerToken = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ProfitBasePerToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StakeType = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StakeTypeRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StakeTypeBase = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RewardOwnerRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.RewardBase = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.RewardBankRate = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TimeType = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.KeyRatePerT = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.KeyBasePerT = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ControllerCheckData is a free data retrieval call binding the contract method 0x937c16a0.
//
// Solidity: function controllerCheckData(address ) view returns(uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 stakeType, uint256 stakeTypeRate, uint256 stakeTypeBase, uint256 rewardOwnerRate, uint256 rewardRate, uint256 rewardBase, uint256 rewardBankRate, uint256 timeType, uint256 keyRatePerT, uint256 keyBasePerT)
func (_TokenFactory *TokenFactorySession) ControllerCheckData(arg0 common.Address) (struct {
	CostRatePerToken   *big.Int
	CostBasePerToken   *big.Int
	ProfitRatePerToken *big.Int
	ProfitBasePerToken *big.Int
	StakeType          *big.Int
	StakeTypeRate      *big.Int
	StakeTypeBase      *big.Int
	RewardOwnerRate    *big.Int
	RewardRate         *big.Int
	RewardBase         *big.Int
	RewardBankRate     *big.Int
	TimeType           *big.Int
	KeyRatePerT        *big.Int
	KeyBasePerT        *big.Int
}, error) {
	return _TokenFactory.Contract.ControllerCheckData(&_TokenFactory.CallOpts, arg0)
}

// ControllerCheckData is a free data retrieval call binding the contract method 0x937c16a0.
//
// Solidity: function controllerCheckData(address ) view returns(uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 stakeType, uint256 stakeTypeRate, uint256 stakeTypeBase, uint256 rewardOwnerRate, uint256 rewardRate, uint256 rewardBase, uint256 rewardBankRate, uint256 timeType, uint256 keyRatePerT, uint256 keyBasePerT)
func (_TokenFactory *TokenFactoryCallerSession) ControllerCheckData(arg0 common.Address) (struct {
	CostRatePerToken   *big.Int
	CostBasePerToken   *big.Int
	ProfitRatePerToken *big.Int
	ProfitBasePerToken *big.Int
	StakeType          *big.Int
	StakeTypeRate      *big.Int
	StakeTypeBase      *big.Int
	RewardOwnerRate    *big.Int
	RewardRate         *big.Int
	RewardBase         *big.Int
	RewardBankRate     *big.Int
	TimeType           *big.Int
	KeyRatePerT        *big.Int
	KeyBasePerT        *big.Int
}, error) {
	return _TokenFactory.Contract.ControllerCheckData(&_TokenFactory.CallOpts, arg0)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_TokenFactory *TokenFactoryCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_TokenFactory *TokenFactorySession) DefaultAdmin() (common.Address, error) {
	return _TokenFactory.Contract.DefaultAdmin(&_TokenFactory.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) DefaultAdmin() (common.Address, error) {
	return _TokenFactory.Contract.DefaultAdmin(&_TokenFactory.CallOpts)
}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_TokenFactory *TokenFactoryCaller) DefaultAdminRoleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "defaultAdminRoleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_TokenFactory *TokenFactorySession) DefaultAdminRoleLimit() (*big.Int, error) {
	return _TokenFactory.Contract.DefaultAdminRoleLimit(&_TokenFactory.CallOpts)
}

// DefaultAdminRoleLimit is a free data retrieval call binding the contract method 0x2cd8f377.
//
// Solidity: function defaultAdminRoleLimit() view returns(uint256)
func (_TokenFactory *TokenFactoryCallerSession) DefaultAdminRoleLimit() (*big.Int, error) {
	return _TokenFactory.Contract.DefaultAdminRoleLimit(&_TokenFactory.CallOpts)
}

// Dfil is a free data retrieval call binding the contract method 0x7baf0d5c.
//
// Solidity: function dfil() view returns(address)
func (_TokenFactory *TokenFactoryCaller) Dfil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "dfil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dfil is a free data retrieval call binding the contract method 0x7baf0d5c.
//
// Solidity: function dfil() view returns(address)
func (_TokenFactory *TokenFactorySession) Dfil() (common.Address, error) {
	return _TokenFactory.Contract.Dfil(&_TokenFactory.CallOpts)
}

// Dfil is a free data retrieval call binding the contract method 0x7baf0d5c.
//
// Solidity: function dfil() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) Dfil() (common.Address, error) {
	return _TokenFactory.Contract.Dfil(&_TokenFactory.CallOpts)
}

// ExistsCheckOwner is a free data retrieval call binding the contract method 0xd2a4ce17.
//
// Solidity: function existsCheckOwner(address owner) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) ExistsCheckOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "existsCheckOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsCheckOwner is a free data retrieval call binding the contract method 0xd2a4ce17.
//
// Solidity: function existsCheckOwner(address owner) view returns(bool)
func (_TokenFactory *TokenFactorySession) ExistsCheckOwner(owner common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsCheckOwner(&_TokenFactory.CallOpts, owner)
}

// ExistsCheckOwner is a free data retrieval call binding the contract method 0xd2a4ce17.
//
// Solidity: function existsCheckOwner(address owner) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) ExistsCheckOwner(owner common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsCheckOwner(&_TokenFactory.CallOpts, owner)
}

// ExistsOwner is a free data retrieval call binding the contract method 0x8f1b8293.
//
// Solidity: function existsOwner(address owner) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) ExistsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "existsOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsOwner is a free data retrieval call binding the contract method 0x8f1b8293.
//
// Solidity: function existsOwner(address owner) view returns(bool)
func (_TokenFactory *TokenFactorySession) ExistsOwner(owner common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsOwner(&_TokenFactory.CallOpts, owner)
}

// ExistsOwner is a free data retrieval call binding the contract method 0x8f1b8293.
//
// Solidity: function existsOwner(address owner) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) ExistsOwner(owner common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsOwner(&_TokenFactory.CallOpts, owner)
}

// ExistsOwnerActorMiner is a free data retrieval call binding the contract method 0x0081076b.
//
// Solidity: function existsOwnerActorMiner(address owner, uint64 actor) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) ExistsOwnerActorMiner(opts *bind.CallOpts, owner common.Address, actor uint64) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "existsOwnerActorMiner", owner, actor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsOwnerActorMiner is a free data retrieval call binding the contract method 0x0081076b.
//
// Solidity: function existsOwnerActorMiner(address owner, uint64 actor) view returns(bool)
func (_TokenFactory *TokenFactorySession) ExistsOwnerActorMiner(owner common.Address, actor uint64) (bool, error) {
	return _TokenFactory.Contract.ExistsOwnerActorMiner(&_TokenFactory.CallOpts, owner, actor)
}

// ExistsOwnerActorMiner is a free data retrieval call binding the contract method 0x0081076b.
//
// Solidity: function existsOwnerActorMiner(address owner, uint64 actor) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) ExistsOwnerActorMiner(owner common.Address, actor uint64) (bool, error) {
	return _TokenFactory.Contract.ExistsOwnerActorMiner(&_TokenFactory.CallOpts, owner, actor)
}

// ExistsToken is a free data retrieval call binding the contract method 0xcbfb3a37.
//
// Solidity: function existsToken(address token) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) ExistsToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "existsToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsToken is a free data retrieval call binding the contract method 0xcbfb3a37.
//
// Solidity: function existsToken(address token) view returns(bool)
func (_TokenFactory *TokenFactorySession) ExistsToken(token common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsToken(&_TokenFactory.CallOpts, token)
}

// ExistsToken is a free data retrieval call binding the contract method 0xcbfb3a37.
//
// Solidity: function existsToken(address token) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) ExistsToken(token common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsToken(&_TokenFactory.CallOpts, token)
}

// ExistsTopUnionToken is a free data retrieval call binding the contract method 0x82df4eb5.
//
// Solidity: function existsTopUnionToken(address token) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) ExistsTopUnionToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "existsTopUnionToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsTopUnionToken is a free data retrieval call binding the contract method 0x82df4eb5.
//
// Solidity: function existsTopUnionToken(address token) view returns(bool)
func (_TokenFactory *TokenFactorySession) ExistsTopUnionToken(token common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsTopUnionToken(&_TokenFactory.CallOpts, token)
}

// ExistsTopUnionToken is a free data retrieval call binding the contract method 0x82df4eb5.
//
// Solidity: function existsTopUnionToken(address token) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) ExistsTopUnionToken(token common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsTopUnionToken(&_TokenFactory.CallOpts, token)
}

// ExistsUser is a free data retrieval call binding the contract method 0xf409e87d.
//
// Solidity: function existsUser(address account) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) ExistsUser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "existsUser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsUser is a free data retrieval call binding the contract method 0xf409e87d.
//
// Solidity: function existsUser(address account) view returns(bool)
func (_TokenFactory *TokenFactorySession) ExistsUser(account common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsUser(&_TokenFactory.CallOpts, account)
}

// ExistsUser is a free data retrieval call binding the contract method 0xf409e87d.
//
// Solidity: function existsUser(address account) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) ExistsUser(account common.Address) (bool, error) {
	return _TokenFactory.Contract.ExistsUser(&_TokenFactory.CallOpts, account)
}

// FilecoinMinerControllerTemplate is a free data retrieval call binding the contract method 0xac9afa81.
//
// Solidity: function filecoinMinerControllerTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCaller) FilecoinMinerControllerTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "filecoinMinerControllerTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FilecoinMinerControllerTemplate is a free data retrieval call binding the contract method 0xac9afa81.
//
// Solidity: function filecoinMinerControllerTemplate() view returns(address)
func (_TokenFactory *TokenFactorySession) FilecoinMinerControllerTemplate() (common.Address, error) {
	return _TokenFactory.Contract.FilecoinMinerControllerTemplate(&_TokenFactory.CallOpts)
}

// FilecoinMinerControllerTemplate is a free data retrieval call binding the contract method 0xac9afa81.
//
// Solidity: function filecoinMinerControllerTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) FilecoinMinerControllerTemplate() (common.Address, error) {
	return _TokenFactory.Contract.FilecoinMinerControllerTemplate(&_TokenFactory.CallOpts)
}

// FilecoinMinerTemplate is a free data retrieval call binding the contract method 0xc728122b.
//
// Solidity: function filecoinMinerTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCaller) FilecoinMinerTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "filecoinMinerTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FilecoinMinerTemplate is a free data retrieval call binding the contract method 0xc728122b.
//
// Solidity: function filecoinMinerTemplate() view returns(address)
func (_TokenFactory *TokenFactorySession) FilecoinMinerTemplate() (common.Address, error) {
	return _TokenFactory.Contract.FilecoinMinerTemplate(&_TokenFactory.CallOpts)
}

// FilecoinMinerTemplate is a free data retrieval call binding the contract method 0xc728122b.
//
// Solidity: function filecoinMinerTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) FilecoinMinerTemplate() (common.Address, error) {
	return _TokenFactory.Contract.FilecoinMinerTemplate(&_TokenFactory.CallOpts)
}

// GetCheckOwners is a free data retrieval call binding the contract method 0xe4df80c6.
//
// Solidity: function getCheckOwners() view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetCheckOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getCheckOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCheckOwners is a free data retrieval call binding the contract method 0xe4df80c6.
//
// Solidity: function getCheckOwners() view returns(address[])
func (_TokenFactory *TokenFactorySession) GetCheckOwners() ([]common.Address, error) {
	return _TokenFactory.Contract.GetCheckOwners(&_TokenFactory.CallOpts)
}

// GetCheckOwners is a free data retrieval call binding the contract method 0xe4df80c6.
//
// Solidity: function getCheckOwners() view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetCheckOwners() ([]common.Address, error) {
	return _TokenFactory.Contract.GetCheckOwners(&_TokenFactory.CallOpts)
}

// GetCheckOwnersByIndex is a free data retrieval call binding the contract method 0xf164270f.
//
// Solidity: function getCheckOwnersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetCheckOwnersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getCheckOwnersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCheckOwnersByIndex is a free data retrieval call binding the contract method 0xf164270f.
//
// Solidity: function getCheckOwnersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetCheckOwnersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetCheckOwnersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetCheckOwnersByIndex is a free data retrieval call binding the contract method 0xf164270f.
//
// Solidity: function getCheckOwnersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetCheckOwnersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetCheckOwnersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetLowUsers is a free data retrieval call binding the contract method 0x78ce339e.
//
// Solidity: function getLowUsers(address account) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetLowUsers(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getLowUsers", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLowUsers is a free data retrieval call binding the contract method 0x78ce339e.
//
// Solidity: function getLowUsers(address account) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetLowUsers(account common.Address) ([]common.Address, error) {
	return _TokenFactory.Contract.GetLowUsers(&_TokenFactory.CallOpts, account)
}

// GetLowUsers is a free data retrieval call binding the contract method 0x78ce339e.
//
// Solidity: function getLowUsers(address account) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetLowUsers(account common.Address) ([]common.Address, error) {
	return _TokenFactory.Contract.GetLowUsers(&_TokenFactory.CallOpts, account)
}

// GetOwnerAcotrMinersByIndex is a free data retrieval call binding the contract method 0xb61ddd61.
//
// Solidity: function getOwnerAcotrMinersByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_TokenFactory *TokenFactoryCaller) GetOwnerAcotrMinersByIndex(opts *bind.CallOpts, owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getOwnerAcotrMinersByIndex", owner, startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOwnerAcotrMinersByIndex is a free data retrieval call binding the contract method 0xb61ddd61.
//
// Solidity: function getOwnerAcotrMinersByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_TokenFactory *TokenFactorySession) GetOwnerAcotrMinersByIndex(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _TokenFactory.Contract.GetOwnerAcotrMinersByIndex(&_TokenFactory.CallOpts, owner, startIndex, endIndex)
}

// GetOwnerAcotrMinersByIndex is a free data retrieval call binding the contract method 0xb61ddd61.
//
// Solidity: function getOwnerAcotrMinersByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_TokenFactory *TokenFactoryCallerSession) GetOwnerAcotrMinersByIndex(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _TokenFactory.Contract.GetOwnerAcotrMinersByIndex(&_TokenFactory.CallOpts, owner, startIndex, endIndex)
}

// GetOwnerActorCurrentControllers is a free data retrieval call binding the contract method 0x34bcb37b.
//
// Solidity: function getOwnerActorCurrentControllers(address account, uint64 actor) view returns(address)
func (_TokenFactory *TokenFactoryCaller) GetOwnerActorCurrentControllers(opts *bind.CallOpts, account common.Address, actor uint64) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getOwnerActorCurrentControllers", account, actor)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerActorCurrentControllers is a free data retrieval call binding the contract method 0x34bcb37b.
//
// Solidity: function getOwnerActorCurrentControllers(address account, uint64 actor) view returns(address)
func (_TokenFactory *TokenFactorySession) GetOwnerActorCurrentControllers(account common.Address, actor uint64) (common.Address, error) {
	return _TokenFactory.Contract.GetOwnerActorCurrentControllers(&_TokenFactory.CallOpts, account, actor)
}

// GetOwnerActorCurrentControllers is a free data retrieval call binding the contract method 0x34bcb37b.
//
// Solidity: function getOwnerActorCurrentControllers(address account, uint64 actor) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) GetOwnerActorCurrentControllers(account common.Address, actor uint64) (common.Address, error) {
	return _TokenFactory.Contract.GetOwnerActorCurrentControllers(&_TokenFactory.CallOpts, account, actor)
}

// GetOwnerActorMiners is a free data retrieval call binding the contract method 0xd83a075e.
//
// Solidity: function getOwnerActorMiners(address owner) view returns(uint256[])
func (_TokenFactory *TokenFactoryCaller) GetOwnerActorMiners(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getOwnerActorMiners", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOwnerActorMiners is a free data retrieval call binding the contract method 0xd83a075e.
//
// Solidity: function getOwnerActorMiners(address owner) view returns(uint256[])
func (_TokenFactory *TokenFactorySession) GetOwnerActorMiners(owner common.Address) ([]*big.Int, error) {
	return _TokenFactory.Contract.GetOwnerActorMiners(&_TokenFactory.CallOpts, owner)
}

// GetOwnerActorMiners is a free data retrieval call binding the contract method 0xd83a075e.
//
// Solidity: function getOwnerActorMiners(address owner) view returns(uint256[])
func (_TokenFactory *TokenFactoryCallerSession) GetOwnerActorMiners(owner common.Address) ([]*big.Int, error) {
	return _TokenFactory.Contract.GetOwnerActorMiners(&_TokenFactory.CallOpts, owner)
}

// GetOwnerUsers is a free data retrieval call binding the contract method 0x8ae541a0.
//
// Solidity: function getOwnerUsers(address owner) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetOwnerUsers(opts *bind.CallOpts, owner common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getOwnerUsers", owner)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwnerUsers is a free data retrieval call binding the contract method 0x8ae541a0.
//
// Solidity: function getOwnerUsers(address owner) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetOwnerUsers(owner common.Address) ([]common.Address, error) {
	return _TokenFactory.Contract.GetOwnerUsers(&_TokenFactory.CallOpts, owner)
}

// GetOwnerUsers is a free data retrieval call binding the contract method 0x8ae541a0.
//
// Solidity: function getOwnerUsers(address owner) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetOwnerUsers(owner common.Address) ([]common.Address, error) {
	return _TokenFactory.Contract.GetOwnerUsers(&_TokenFactory.CallOpts, owner)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_TokenFactory *TokenFactorySession) GetOwners() ([]common.Address, error) {
	return _TokenFactory.Contract.GetOwners(&_TokenFactory.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetOwners() ([]common.Address, error) {
	return _TokenFactory.Contract.GetOwners(&_TokenFactory.CallOpts)
}

// GetOwnersByIndex is a free data retrieval call binding the contract method 0x4485f182.
//
// Solidity: function getOwnersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetOwnersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getOwnersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwnersByIndex is a free data retrieval call binding the contract method 0x4485f182.
//
// Solidity: function getOwnersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetOwnersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetOwnersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetOwnersByIndex is a free data retrieval call binding the contract method 0x4485f182.
//
// Solidity: function getOwnersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetOwnersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetOwnersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenFactory *TokenFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenFactory *TokenFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenFactory.Contract.GetRoleAdmin(&_TokenFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenFactory *TokenFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenFactory.Contract.GetRoleAdmin(&_TokenFactory.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenFactory *TokenFactoryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenFactory *TokenFactorySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenFactory.Contract.GetRoleMember(&_TokenFactory.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenFactory.Contract.GetRoleMember(&_TokenFactory.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenFactory *TokenFactoryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenFactory *TokenFactorySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenFactory.Contract.GetRoleMemberCount(&_TokenFactory.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenFactory *TokenFactoryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenFactory.Contract.GetRoleMemberCount(&_TokenFactory.CallOpts, role)
}

// GetTokenOwnerActorMinersCheckData is a free data retrieval call binding the contract method 0xc0fd7c3c.
//
// Solidity: function getTokenOwnerActorMinersCheckData(address owner, uint64 actor) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_TokenFactory *TokenFactoryCaller) GetTokenOwnerActorMinersCheckData(opts *bind.CallOpts, owner common.Address, actor uint64) (IFilecoinMinerControllerTemplateCheckData, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokenOwnerActorMinersCheckData", owner, actor)

	if err != nil {
		return *new(IFilecoinMinerControllerTemplateCheckData), err
	}

	out0 := *abi.ConvertType(out[0], new(IFilecoinMinerControllerTemplateCheckData)).(*IFilecoinMinerControllerTemplateCheckData)

	return out0, err

}

// GetTokenOwnerActorMinersCheckData is a free data retrieval call binding the contract method 0xc0fd7c3c.
//
// Solidity: function getTokenOwnerActorMinersCheckData(address owner, uint64 actor) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_TokenFactory *TokenFactorySession) GetTokenOwnerActorMinersCheckData(owner common.Address, actor uint64) (IFilecoinMinerControllerTemplateCheckData, error) {
	return _TokenFactory.Contract.GetTokenOwnerActorMinersCheckData(&_TokenFactory.CallOpts, owner, actor)
}

// GetTokenOwnerActorMinersCheckData is a free data retrieval call binding the contract method 0xc0fd7c3c.
//
// Solidity: function getTokenOwnerActorMinersCheckData(address owner, uint64 actor) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_TokenFactory *TokenFactoryCallerSession) GetTokenOwnerActorMinersCheckData(owner common.Address, actor uint64) (IFilecoinMinerControllerTemplateCheckData, error) {
	return _TokenFactory.Contract.GetTokenOwnerActorMinersCheckData(&_TokenFactory.CallOpts, owner, actor)
}

// GetTokenOwnerTokens is a free data retrieval call binding the contract method 0xd18e0e00.
//
// Solidity: function getTokenOwnerTokens(address owner) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetTokenOwnerTokens(opts *bind.CallOpts, owner common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokenOwnerTokens", owner)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenOwnerTokens is a free data retrieval call binding the contract method 0xd18e0e00.
//
// Solidity: function getTokenOwnerTokens(address owner) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetTokenOwnerTokens(owner common.Address) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokenOwnerTokens(&_TokenFactory.CallOpts, owner)
}

// GetTokenOwnerTokens is a free data retrieval call binding the contract method 0xd18e0e00.
//
// Solidity: function getTokenOwnerTokens(address owner) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetTokenOwnerTokens(owner common.Address) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokenOwnerTokens(&_TokenFactory.CallOpts, owner)
}

// GetTokenOwnerTokensByIndex is a free data retrieval call binding the contract method 0x2f09e546.
//
// Solidity: function getTokenOwnerTokensByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetTokenOwnerTokensByIndex(opts *bind.CallOpts, owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokenOwnerTokensByIndex", owner, startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenOwnerTokensByIndex is a free data retrieval call binding the contract method 0x2f09e546.
//
// Solidity: function getTokenOwnerTokensByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetTokenOwnerTokensByIndex(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokenOwnerTokensByIndex(&_TokenFactory.CallOpts, owner, startIndex, endIndex)
}

// GetTokenOwnerTokensByIndex is a free data retrieval call binding the contract method 0x2f09e546.
//
// Solidity: function getTokenOwnerTokensByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetTokenOwnerTokensByIndex(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokenOwnerTokensByIndex(&_TokenFactory.CallOpts, owner, startIndex, endIndex)
}

// GetTokenOwnerUsersByIndex is a free data retrieval call binding the contract method 0x9eabf337.
//
// Solidity: function getTokenOwnerUsersByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetTokenOwnerUsersByIndex(opts *bind.CallOpts, owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokenOwnerUsersByIndex", owner, startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenOwnerUsersByIndex is a free data retrieval call binding the contract method 0x9eabf337.
//
// Solidity: function getTokenOwnerUsersByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetTokenOwnerUsersByIndex(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokenOwnerUsersByIndex(&_TokenFactory.CallOpts, owner, startIndex, endIndex)
}

// GetTokenOwnerUsersByIndex is a free data retrieval call binding the contract method 0x9eabf337.
//
// Solidity: function getTokenOwnerUsersByIndex(address owner, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetTokenOwnerUsersByIndex(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokenOwnerUsersByIndex(&_TokenFactory.CallOpts, owner, startIndex, endIndex)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_TokenFactory *TokenFactorySession) GetTokens() ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokens(&_TokenFactory.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetTokens() ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokens(&_TokenFactory.CallOpts)
}

// GetTokensByIndex is a free data retrieval call binding the contract method 0xf5ba16f5.
//
// Solidity: function getTokensByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetTokensByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokensByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokensByIndex is a free data retrieval call binding the contract method 0xf5ba16f5.
//
// Solidity: function getTokensByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetTokensByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokensByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetTokensByIndex is a free data retrieval call binding the contract method 0xf5ba16f5.
//
// Solidity: function getTokensByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetTokensByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetTokensByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetTop is a free data retrieval call binding the contract method 0x5c2b1119.
//
// Solidity: function getTop() view returns(address)
func (_TokenFactory *TokenFactoryCaller) GetTop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTop is a free data retrieval call binding the contract method 0x5c2b1119.
//
// Solidity: function getTop() view returns(address)
func (_TokenFactory *TokenFactorySession) GetTop() (common.Address, error) {
	return _TokenFactory.Contract.GetTop(&_TokenFactory.CallOpts)
}

// GetTop is a free data retrieval call binding the contract method 0x5c2b1119.
//
// Solidity: function getTop() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) GetTop() (common.Address, error) {
	return _TokenFactory.Contract.GetTop(&_TokenFactory.CallOpts)
}

// GetTopAcotrMinersByIndex is a free data retrieval call binding the contract method 0x9fc5b3cd.
//
// Solidity: function getTopAcotrMinersByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_TokenFactory *TokenFactoryCaller) GetTopAcotrMinersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTopAcotrMinersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTopAcotrMinersByIndex is a free data retrieval call binding the contract method 0x9fc5b3cd.
//
// Solidity: function getTopAcotrMinersByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_TokenFactory *TokenFactorySession) GetTopAcotrMinersByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _TokenFactory.Contract.GetTopAcotrMinersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetTopAcotrMinersByIndex is a free data retrieval call binding the contract method 0x9fc5b3cd.
//
// Solidity: function getTopAcotrMinersByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_TokenFactory *TokenFactoryCallerSession) GetTopAcotrMinersByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _TokenFactory.Contract.GetTopAcotrMinersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetTopActorMiners is a free data retrieval call binding the contract method 0x86068595.
//
// Solidity: function getTopActorMiners() view returns(uint256[])
func (_TokenFactory *TokenFactoryCaller) GetTopActorMiners(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTopActorMiners")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTopActorMiners is a free data retrieval call binding the contract method 0x86068595.
//
// Solidity: function getTopActorMiners() view returns(uint256[])
func (_TokenFactory *TokenFactorySession) GetTopActorMiners() ([]*big.Int, error) {
	return _TokenFactory.Contract.GetTopActorMiners(&_TokenFactory.CallOpts)
}

// GetTopActorMiners is a free data retrieval call binding the contract method 0x86068595.
//
// Solidity: function getTopActorMiners() view returns(uint256[])
func (_TokenFactory *TokenFactoryCallerSession) GetTopActorMiners() ([]*big.Int, error) {
	return _TokenFactory.Contract.GetTopActorMiners(&_TokenFactory.CallOpts)
}

// GetTopActorMinersCheckData is a free data retrieval call binding the contract method 0xf2cf05ec.
//
// Solidity: function getTopActorMinersCheckData(uint64 actor) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_TokenFactory *TokenFactoryCaller) GetTopActorMinersCheckData(opts *bind.CallOpts, actor uint64) (IFilecoinMinerControllerTemplateCheckData, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTopActorMinersCheckData", actor)

	if err != nil {
		return *new(IFilecoinMinerControllerTemplateCheckData), err
	}

	out0 := *abi.ConvertType(out[0], new(IFilecoinMinerControllerTemplateCheckData)).(*IFilecoinMinerControllerTemplateCheckData)

	return out0, err

}

// GetTopActorMinersCheckData is a free data retrieval call binding the contract method 0xf2cf05ec.
//
// Solidity: function getTopActorMinersCheckData(uint64 actor) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_TokenFactory *TokenFactorySession) GetTopActorMinersCheckData(actor uint64) (IFilecoinMinerControllerTemplateCheckData, error) {
	return _TokenFactory.Contract.GetTopActorMinersCheckData(&_TokenFactory.CallOpts, actor)
}

// GetTopActorMinersCheckData is a free data retrieval call binding the contract method 0xf2cf05ec.
//
// Solidity: function getTopActorMinersCheckData(uint64 actor) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_TokenFactory *TokenFactoryCallerSession) GetTopActorMinersCheckData(actor uint64) (IFilecoinMinerControllerTemplateCheckData, error) {
	return _TokenFactory.Contract.GetTopActorMinersCheckData(&_TokenFactory.CallOpts, actor)
}

// GetTopUnionTokens is a free data retrieval call binding the contract method 0xe418daea.
//
// Solidity: function getTopUnionTokens() view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetTopUnionTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTopUnionTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopUnionTokens is a free data retrieval call binding the contract method 0xe418daea.
//
// Solidity: function getTopUnionTokens() view returns(address[])
func (_TokenFactory *TokenFactorySession) GetTopUnionTokens() ([]common.Address, error) {
	return _TokenFactory.Contract.GetTopUnionTokens(&_TokenFactory.CallOpts)
}

// GetTopUnionTokens is a free data retrieval call binding the contract method 0xe418daea.
//
// Solidity: function getTopUnionTokens() view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetTopUnionTokens() ([]common.Address, error) {
	return _TokenFactory.Contract.GetTopUnionTokens(&_TokenFactory.CallOpts)
}

// GetUserOwner is a free data retrieval call binding the contract method 0x710646c5.
//
// Solidity: function getUserOwner() view returns(address)
func (_TokenFactory *TokenFactoryCaller) GetUserOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getUserOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserOwner is a free data retrieval call binding the contract method 0x710646c5.
//
// Solidity: function getUserOwner() view returns(address)
func (_TokenFactory *TokenFactorySession) GetUserOwner() (common.Address, error) {
	return _TokenFactory.Contract.GetUserOwner(&_TokenFactory.CallOpts)
}

// GetUserOwner is a free data retrieval call binding the contract method 0x710646c5.
//
// Solidity: function getUserOwner() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) GetUserOwner() (common.Address, error) {
	return _TokenFactory.Contract.GetUserOwner(&_TokenFactory.CallOpts)
}

// GetUserOwnerByAccount is a free data retrieval call binding the contract method 0xcb747105.
//
// Solidity: function getUserOwnerByAccount(address account) view returns(address)
func (_TokenFactory *TokenFactoryCaller) GetUserOwnerByAccount(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getUserOwnerByAccount", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserOwnerByAccount is a free data retrieval call binding the contract method 0xcb747105.
//
// Solidity: function getUserOwnerByAccount(address account) view returns(address)
func (_TokenFactory *TokenFactorySession) GetUserOwnerByAccount(account common.Address) (common.Address, error) {
	return _TokenFactory.Contract.GetUserOwnerByAccount(&_TokenFactory.CallOpts, account)
}

// GetUserOwnerByAccount is a free data retrieval call binding the contract method 0xcb747105.
//
// Solidity: function getUserOwnerByAccount(address account) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) GetUserOwnerByAccount(account common.Address) (common.Address, error) {
	return _TokenFactory.Contract.GetUserOwnerByAccount(&_TokenFactory.CallOpts, account)
}

// GetUserWithLowUsersByIndex is a free data retrieval call binding the contract method 0x1da2e481.
//
// Solidity: function getUserWithLowUsersByIndex(address user, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetUserWithLowUsersByIndex(opts *bind.CallOpts, user common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getUserWithLowUsersByIndex", user, startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserWithLowUsersByIndex is a free data retrieval call binding the contract method 0x1da2e481.
//
// Solidity: function getUserWithLowUsersByIndex(address user, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetUserWithLowUsersByIndex(user common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetUserWithLowUsersByIndex(&_TokenFactory.CallOpts, user, startIndex, endIndex)
}

// GetUserWithLowUsersByIndex is a free data retrieval call binding the contract method 0x1da2e481.
//
// Solidity: function getUserWithLowUsersByIndex(address user, uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetUserWithLowUsersByIndex(user common.Address, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetUserWithLowUsersByIndex(&_TokenFactory.CallOpts, user, startIndex, endIndex)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_TokenFactory *TokenFactorySession) GetUsers() ([]common.Address, error) {
	return _TokenFactory.Contract.GetUsers(&_TokenFactory.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetUsers() ([]common.Address, error) {
	return _TokenFactory.Contract.GetUsers(&_TokenFactory.CallOpts)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCaller) GetUsersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getUsersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactorySession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetUsersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_TokenFactory *TokenFactoryCallerSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _TokenFactory.Contract.GetUsersByIndex(&_TokenFactory.CallOpts, startIndex, endIndex)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenFactory *TokenFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenFactory.Contract.HasRole(&_TokenFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenFactory.Contract.HasRole(&_TokenFactory.CallOpts, role, account)
}

// IdoStartTime is a free data retrieval call binding the contract method 0x70a6dea8.
//
// Solidity: function idoStartTime() view returns(uint256)
func (_TokenFactory *TokenFactoryCaller) IdoStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "idoStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdoStartTime is a free data retrieval call binding the contract method 0x70a6dea8.
//
// Solidity: function idoStartTime() view returns(uint256)
func (_TokenFactory *TokenFactorySession) IdoStartTime() (*big.Int, error) {
	return _TokenFactory.Contract.IdoStartTime(&_TokenFactory.CallOpts)
}

// IdoStartTime is a free data retrieval call binding the contract method 0x70a6dea8.
//
// Solidity: function idoStartTime() view returns(uint256)
func (_TokenFactory *TokenFactoryCallerSession) IdoStartTime() (*big.Int, error) {
	return _TokenFactory.Contract.IdoStartTime(&_TokenFactory.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() view returns(address)
func (_TokenFactory *TokenFactoryCaller) Key(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "key")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() view returns(address)
func (_TokenFactory *TokenFactorySession) Key() (common.Address, error) {
	return _TokenFactory.Contract.Key(&_TokenFactory.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) Key() (common.Address, error) {
	return _TokenFactory.Contract.Key(&_TokenFactory.CallOpts)
}

// OwnerActorCurrentControllers is a free data retrieval call binding the contract method 0xd1da15c0.
//
// Solidity: function ownerActorCurrentControllers(address , uint64 ) view returns(address)
func (_TokenFactory *TokenFactoryCaller) OwnerActorCurrentControllers(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "ownerActorCurrentControllers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerActorCurrentControllers is a free data retrieval call binding the contract method 0xd1da15c0.
//
// Solidity: function ownerActorCurrentControllers(address , uint64 ) view returns(address)
func (_TokenFactory *TokenFactorySession) OwnerActorCurrentControllers(arg0 common.Address, arg1 uint64) (common.Address, error) {
	return _TokenFactory.Contract.OwnerActorCurrentControllers(&_TokenFactory.CallOpts, arg0, arg1)
}

// OwnerActorCurrentControllers is a free data retrieval call binding the contract method 0xd1da15c0.
//
// Solidity: function ownerActorCurrentControllers(address , uint64 ) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) OwnerActorCurrentControllers(arg0 common.Address, arg1 uint64) (common.Address, error) {
	return _TokenFactory.Contract.OwnerActorCurrentControllers(&_TokenFactory.CallOpts, arg0, arg1)
}

// RewardStartTime is a free data retrieval call binding the contract method 0x2cc138be.
//
// Solidity: function rewardStartTime() view returns(uint256)
func (_TokenFactory *TokenFactoryCaller) RewardStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "rewardStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardStartTime is a free data retrieval call binding the contract method 0x2cc138be.
//
// Solidity: function rewardStartTime() view returns(uint256)
func (_TokenFactory *TokenFactorySession) RewardStartTime() (*big.Int, error) {
	return _TokenFactory.Contract.RewardStartTime(&_TokenFactory.CallOpts)
}

// RewardStartTime is a free data retrieval call binding the contract method 0x2cc138be.
//
// Solidity: function rewardStartTime() view returns(uint256)
func (_TokenFactory *TokenFactoryCallerSession) RewardStartTime() (*big.Int, error) {
	return _TokenFactory.Contract.RewardStartTime(&_TokenFactory.CallOpts)
}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_TokenFactory *TokenFactoryCaller) SuperAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "superAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_TokenFactory *TokenFactorySession) SuperAdmin() (common.Address, error) {
	return _TokenFactory.Contract.SuperAdmin(&_TokenFactory.CallOpts)
}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) SuperAdmin() (common.Address, error) {
	return _TokenFactory.Contract.SuperAdmin(&_TokenFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenFactory *TokenFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenFactory *TokenFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenFactory.Contract.SupportsInterface(&_TokenFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenFactory *TokenFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenFactory.Contract.SupportsInterface(&_TokenFactory.CallOpts, interfaceId)
}

// SwapFactory is a free data retrieval call binding the contract method 0x7944f944.
//
// Solidity: function swapFactory() view returns(address)
func (_TokenFactory *TokenFactoryCaller) SwapFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "swapFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapFactory is a free data retrieval call binding the contract method 0x7944f944.
//
// Solidity: function swapFactory() view returns(address)
func (_TokenFactory *TokenFactorySession) SwapFactory() (common.Address, error) {
	return _TokenFactory.Contract.SwapFactory(&_TokenFactory.CallOpts)
}

// SwapFactory is a free data retrieval call binding the contract method 0x7944f944.
//
// Solidity: function swapFactory() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) SwapFactory() (common.Address, error) {
	return _TokenFactory.Contract.SwapFactory(&_TokenFactory.CallOpts)
}

// TokenExchange is a free data retrieval call binding the contract method 0x636f6159.
//
// Solidity: function tokenExchange() view returns(address)
func (_TokenFactory *TokenFactoryCaller) TokenExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "tokenExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenExchange is a free data retrieval call binding the contract method 0x636f6159.
//
// Solidity: function tokenExchange() view returns(address)
func (_TokenFactory *TokenFactorySession) TokenExchange() (common.Address, error) {
	return _TokenFactory.Contract.TokenExchange(&_TokenFactory.CallOpts)
}

// TokenExchange is a free data retrieval call binding the contract method 0x636f6159.
//
// Solidity: function tokenExchange() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) TokenExchange() (common.Address, error) {
	return _TokenFactory.Contract.TokenExchange(&_TokenFactory.CallOpts)
}

// TokenTemplate is a free data retrieval call binding the contract method 0x0814d3dd.
//
// Solidity: function tokenTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCaller) TokenTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "tokenTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenTemplate is a free data retrieval call binding the contract method 0x0814d3dd.
//
// Solidity: function tokenTemplate() view returns(address)
func (_TokenFactory *TokenFactorySession) TokenTemplate() (common.Address, error) {
	return _TokenFactory.Contract.TokenTemplate(&_TokenFactory.CallOpts)
}

// TokenTemplate is a free data retrieval call binding the contract method 0x0814d3dd.
//
// Solidity: function tokenTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) TokenTemplate() (common.Address, error) {
	return _TokenFactory.Contract.TokenTemplate(&_TokenFactory.CallOpts)
}

// TokenUnionTemplate is a free data retrieval call binding the contract method 0x527f295f.
//
// Solidity: function tokenUnionTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCaller) TokenUnionTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "tokenUnionTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenUnionTemplate is a free data retrieval call binding the contract method 0x527f295f.
//
// Solidity: function tokenUnionTemplate() view returns(address)
func (_TokenFactory *TokenFactorySession) TokenUnionTemplate() (common.Address, error) {
	return _TokenFactory.Contract.TokenUnionTemplate(&_TokenFactory.CallOpts)
}

// TokenUnionTemplate is a free data retrieval call binding the contract method 0x527f295f.
//
// Solidity: function tokenUnionTemplate() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) TokenUnionTemplate() (common.Address, error) {
	return _TokenFactory.Contract.TokenUnionTemplate(&_TokenFactory.CallOpts)
}

// Top is a free data retrieval call binding the contract method 0xfe6dcdba.
//
// Solidity: function top() view returns(address)
func (_TokenFactory *TokenFactoryCaller) Top(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "top")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Top is a free data retrieval call binding the contract method 0xfe6dcdba.
//
// Solidity: function top() view returns(address)
func (_TokenFactory *TokenFactorySession) Top() (common.Address, error) {
	return _TokenFactory.Contract.Top(&_TokenFactory.CallOpts)
}

// Top is a free data retrieval call binding the contract method 0xfe6dcdba.
//
// Solidity: function top() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) Top() (common.Address, error) {
	return _TokenFactory.Contract.Top(&_TokenFactory.CallOpts)
}

// UserWithTokenOwner is a free data retrieval call binding the contract method 0xf1a77489.
//
// Solidity: function userWithTokenOwner(address ) view returns(address)
func (_TokenFactory *TokenFactoryCaller) UserWithTokenOwner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "userWithTokenOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserWithTokenOwner is a free data retrieval call binding the contract method 0xf1a77489.
//
// Solidity: function userWithTokenOwner(address ) view returns(address)
func (_TokenFactory *TokenFactorySession) UserWithTokenOwner(arg0 common.Address) (common.Address, error) {
	return _TokenFactory.Contract.UserWithTokenOwner(&_TokenFactory.CallOpts, arg0)
}

// UserWithTokenOwner is a free data retrieval call binding the contract method 0xf1a77489.
//
// Solidity: function userWithTokenOwner(address ) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) UserWithTokenOwner(arg0 common.Address) (common.Address, error) {
	return _TokenFactory.Contract.UserWithTokenOwner(&_TokenFactory.CallOpts, arg0)
}

// UserWithTopUsers is a free data retrieval call binding the contract method 0x09258779.
//
// Solidity: function userWithTopUsers(address ) view returns(address)
func (_TokenFactory *TokenFactoryCaller) UserWithTopUsers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "userWithTopUsers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserWithTopUsers is a free data retrieval call binding the contract method 0x09258779.
//
// Solidity: function userWithTopUsers(address ) view returns(address)
func (_TokenFactory *TokenFactorySession) UserWithTopUsers(arg0 common.Address) (common.Address, error) {
	return _TokenFactory.Contract.UserWithTopUsers(&_TokenFactory.CallOpts, arg0)
}

// UserWithTopUsers is a free data retrieval call binding the contract method 0x09258779.
//
// Solidity: function userWithTopUsers(address ) view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) UserWithTopUsers(arg0 common.Address) (common.Address, error) {
	return _TokenFactory.Contract.UserWithTopUsers(&_TokenFactory.CallOpts, arg0)
}

// UsersAt is a free data retrieval call binding the contract method 0x3bf87ffd.
//
// Solidity: function usersAt(address ) view returns(uint256)
func (_TokenFactory *TokenFactoryCaller) UsersAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "usersAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersAt is a free data retrieval call binding the contract method 0x3bf87ffd.
//
// Solidity: function usersAt(address ) view returns(uint256)
func (_TokenFactory *TokenFactorySession) UsersAt(arg0 common.Address) (*big.Int, error) {
	return _TokenFactory.Contract.UsersAt(&_TokenFactory.CallOpts, arg0)
}

// UsersAt is a free data retrieval call binding the contract method 0x3bf87ffd.
//
// Solidity: function usersAt(address ) view returns(uint256)
func (_TokenFactory *TokenFactoryCallerSession) UsersAt(arg0 common.Address) (*big.Int, error) {
	return _TokenFactory.Contract.UsersAt(&_TokenFactory.CallOpts, arg0)
}

// CreateActorMinerController is a paid mutator transaction binding the contract method 0x03ff719b.
//
// Solidity: function createActorMinerController((address,uint64,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) createControllerData) returns(address filecoinControllerMiner, address filecoinMiner)
func (_TokenFactory *TokenFactoryTransactor) CreateActorMinerController(opts *bind.TransactOpts, createControllerData IFilecoinMinerControllerTemplateCreateControllerData) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createActorMinerController", createControllerData)
}

// CreateActorMinerController is a paid mutator transaction binding the contract method 0x03ff719b.
//
// Solidity: function createActorMinerController((address,uint64,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) createControllerData) returns(address filecoinControllerMiner, address filecoinMiner)
func (_TokenFactory *TokenFactorySession) CreateActorMinerController(createControllerData IFilecoinMinerControllerTemplateCreateControllerData) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateActorMinerController(&_TokenFactory.TransactOpts, createControllerData)
}

// CreateActorMinerController is a paid mutator transaction binding the contract method 0x03ff719b.
//
// Solidity: function createActorMinerController((address,uint64,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) createControllerData) returns(address filecoinControllerMiner, address filecoinMiner)
func (_TokenFactory *TokenFactoryTransactorSession) CreateActorMinerController(createControllerData IFilecoinMinerControllerTemplateCreateControllerData) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateActorMinerController(&_TokenFactory.TransactOpts, createControllerData)
}

// CreateCheckOwner is a paid mutator transaction binding the contract method 0xacd3367d.
//
// Solidity: function createCheckOwner() returns()
func (_TokenFactory *TokenFactoryTransactor) CreateCheckOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createCheckOwner")
}

// CreateCheckOwner is a paid mutator transaction binding the contract method 0xacd3367d.
//
// Solidity: function createCheckOwner() returns()
func (_TokenFactory *TokenFactorySession) CreateCheckOwner() (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateCheckOwner(&_TokenFactory.TransactOpts)
}

// CreateCheckOwner is a paid mutator transaction binding the contract method 0xacd3367d.
//
// Solidity: function createCheckOwner() returns()
func (_TokenFactory *TokenFactoryTransactorSession) CreateCheckOwner() (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateCheckOwner(&_TokenFactory.TransactOpts)
}

// CreateToken is a paid mutator transaction binding the contract method 0x8c837315.
//
// Solidity: function createToken(uint256 cap, string name, string logo, address owner, uint64 actor, uint256 pledge) returns(address token)
func (_TokenFactory *TokenFactoryTransactor) CreateToken(opts *bind.TransactOpts, cap *big.Int, name string, logo string, owner common.Address, actor uint64, pledge *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createToken", cap, name, logo, owner, actor, pledge)
}

// CreateToken is a paid mutator transaction binding the contract method 0x8c837315.
//
// Solidity: function createToken(uint256 cap, string name, string logo, address owner, uint64 actor, uint256 pledge) returns(address token)
func (_TokenFactory *TokenFactorySession) CreateToken(cap *big.Int, name string, logo string, owner common.Address, actor uint64, pledge *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, cap, name, logo, owner, actor, pledge)
}

// CreateToken is a paid mutator transaction binding the contract method 0x8c837315.
//
// Solidity: function createToken(uint256 cap, string name, string logo, address owner, uint64 actor, uint256 pledge) returns(address token)
func (_TokenFactory *TokenFactoryTransactorSession) CreateToken(cap *big.Int, name string, logo string, owner common.Address, actor uint64, pledge *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, cap, name, logo, owner, actor, pledge)
}

// CreateTokenOwner is a paid mutator transaction binding the contract method 0x3cf0d286.
//
// Solidity: function createTokenOwner(address owner) returns()
func (_TokenFactory *TokenFactoryTransactor) CreateTokenOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createTokenOwner", owner)
}

// CreateTokenOwner is a paid mutator transaction binding the contract method 0x3cf0d286.
//
// Solidity: function createTokenOwner(address owner) returns()
func (_TokenFactory *TokenFactorySession) CreateTokenOwner(owner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateTokenOwner(&_TokenFactory.TransactOpts, owner)
}

// CreateTokenOwner is a paid mutator transaction binding the contract method 0x3cf0d286.
//
// Solidity: function createTokenOwner(address owner) returns()
func (_TokenFactory *TokenFactoryTransactorSession) CreateTokenOwner(owner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateTokenOwner(&_TokenFactory.TransactOpts, owner)
}

// CreateTokenOwnerAcotrMiners is a paid mutator transaction binding the contract method 0x014bbce1.
//
// Solidity: function createTokenOwnerAcotrMiners(uint64 actor, uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 timeType) returns()
func (_TokenFactory *TokenFactoryTransactor) CreateTokenOwnerAcotrMiners(opts *bind.TransactOpts, actor uint64, costRatePerToken *big.Int, costBasePerToken *big.Int, profitRatePerToken *big.Int, profitBasePerToken *big.Int, timeType *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createTokenOwnerAcotrMiners", actor, costRatePerToken, costBasePerToken, profitRatePerToken, profitBasePerToken, timeType)
}

// CreateTokenOwnerAcotrMiners is a paid mutator transaction binding the contract method 0x014bbce1.
//
// Solidity: function createTokenOwnerAcotrMiners(uint64 actor, uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 timeType) returns()
func (_TokenFactory *TokenFactorySession) CreateTokenOwnerAcotrMiners(actor uint64, costRatePerToken *big.Int, costBasePerToken *big.Int, profitRatePerToken *big.Int, profitBasePerToken *big.Int, timeType *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateTokenOwnerAcotrMiners(&_TokenFactory.TransactOpts, actor, costRatePerToken, costBasePerToken, profitRatePerToken, profitBasePerToken, timeType)
}

// CreateTokenOwnerAcotrMiners is a paid mutator transaction binding the contract method 0x014bbce1.
//
// Solidity: function createTokenOwnerAcotrMiners(uint64 actor, uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 timeType) returns()
func (_TokenFactory *TokenFactoryTransactorSession) CreateTokenOwnerAcotrMiners(actor uint64, costRatePerToken *big.Int, costBasePerToken *big.Int, profitRatePerToken *big.Int, profitBasePerToken *big.Int, timeType *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateTokenOwnerAcotrMiners(&_TokenFactory.TransactOpts, actor, costRatePerToken, costBasePerToken, profitRatePerToken, profitBasePerToken, timeType)
}

// CreateTopAcotrMiners is a paid mutator transaction binding the contract method 0x4e8b07f5.
//
// Solidity: function createTopAcotrMiners(uint64 actor, uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 timeType) returns()
func (_TokenFactory *TokenFactoryTransactor) CreateTopAcotrMiners(opts *bind.TransactOpts, actor uint64, costRatePerToken *big.Int, costBasePerToken *big.Int, profitRatePerToken *big.Int, profitBasePerToken *big.Int, timeType *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createTopAcotrMiners", actor, costRatePerToken, costBasePerToken, profitRatePerToken, profitBasePerToken, timeType)
}

// CreateTopAcotrMiners is a paid mutator transaction binding the contract method 0x4e8b07f5.
//
// Solidity: function createTopAcotrMiners(uint64 actor, uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 timeType) returns()
func (_TokenFactory *TokenFactorySession) CreateTopAcotrMiners(actor uint64, costRatePerToken *big.Int, costBasePerToken *big.Int, profitRatePerToken *big.Int, profitBasePerToken *big.Int, timeType *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateTopAcotrMiners(&_TokenFactory.TransactOpts, actor, costRatePerToken, costBasePerToken, profitRatePerToken, profitBasePerToken, timeType)
}

// CreateTopAcotrMiners is a paid mutator transaction binding the contract method 0x4e8b07f5.
//
// Solidity: function createTopAcotrMiners(uint64 actor, uint256 costRatePerToken, uint256 costBasePerToken, uint256 profitRatePerToken, uint256 profitBasePerToken, uint256 timeType) returns()
func (_TokenFactory *TokenFactoryTransactorSession) CreateTopAcotrMiners(actor uint64, costRatePerToken *big.Int, costBasePerToken *big.Int, profitRatePerToken *big.Int, profitBasePerToken *big.Int, timeType *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateTopAcotrMiners(&_TokenFactory.TransactOpts, actor, costRatePerToken, costBasePerToken, profitRatePerToken, profitBasePerToken, timeType)
}

// CreateUnionToken is a paid mutator transaction binding the contract method 0xf5005301.
//
// Solidity: function createUnionToken(uint256 cap, string name, string logo, uint64 actor, uint256 pledge) returns(address token)
func (_TokenFactory *TokenFactoryTransactor) CreateUnionToken(opts *bind.TransactOpts, cap *big.Int, name string, logo string, actor uint64, pledge *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createUnionToken", cap, name, logo, actor, pledge)
}

// CreateUnionToken is a paid mutator transaction binding the contract method 0xf5005301.
//
// Solidity: function createUnionToken(uint256 cap, string name, string logo, uint64 actor, uint256 pledge) returns(address token)
func (_TokenFactory *TokenFactorySession) CreateUnionToken(cap *big.Int, name string, logo string, actor uint64, pledge *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateUnionToken(&_TokenFactory.TransactOpts, cap, name, logo, actor, pledge)
}

// CreateUnionToken is a paid mutator transaction binding the contract method 0xf5005301.
//
// Solidity: function createUnionToken(uint256 cap, string name, string logo, uint64 actor, uint256 pledge) returns(address token)
func (_TokenFactory *TokenFactoryTransactorSession) CreateUnionToken(cap *big.Int, name string, logo string, actor uint64, pledge *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateUnionToken(&_TokenFactory.TransactOpts, cap, name, logo, actor, pledge)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.GrantRole(&_TokenFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.GrantRole(&_TokenFactory.TransactOpts, role, account)
}

// RemoveCheckOwner is a paid mutator transaction binding the contract method 0x073ea1b9.
//
// Solidity: function removeCheckOwner(address owner) returns()
func (_TokenFactory *TokenFactoryTransactor) RemoveCheckOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "removeCheckOwner", owner)
}

// RemoveCheckOwner is a paid mutator transaction binding the contract method 0x073ea1b9.
//
// Solidity: function removeCheckOwner(address owner) returns()
func (_TokenFactory *TokenFactorySession) RemoveCheckOwner(owner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RemoveCheckOwner(&_TokenFactory.TransactOpts, owner)
}

// RemoveCheckOwner is a paid mutator transaction binding the contract method 0x073ea1b9.
//
// Solidity: function removeCheckOwner(address owner) returns()
func (_TokenFactory *TokenFactoryTransactorSession) RemoveCheckOwner(owner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RemoveCheckOwner(&_TokenFactory.TransactOpts, owner)
}

// RemoveTokenOwnerAcotrMiners is a paid mutator transaction binding the contract method 0xee412443.
//
// Solidity: function removeTokenOwnerAcotrMiners(address owner, uint64 actor) returns()
func (_TokenFactory *TokenFactoryTransactor) RemoveTokenOwnerAcotrMiners(opts *bind.TransactOpts, owner common.Address, actor uint64) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "removeTokenOwnerAcotrMiners", owner, actor)
}

// RemoveTokenOwnerAcotrMiners is a paid mutator transaction binding the contract method 0xee412443.
//
// Solidity: function removeTokenOwnerAcotrMiners(address owner, uint64 actor) returns()
func (_TokenFactory *TokenFactorySession) RemoveTokenOwnerAcotrMiners(owner common.Address, actor uint64) (*types.Transaction, error) {
	return _TokenFactory.Contract.RemoveTokenOwnerAcotrMiners(&_TokenFactory.TransactOpts, owner, actor)
}

// RemoveTokenOwnerAcotrMiners is a paid mutator transaction binding the contract method 0xee412443.
//
// Solidity: function removeTokenOwnerAcotrMiners(address owner, uint64 actor) returns()
func (_TokenFactory *TokenFactoryTransactorSession) RemoveTokenOwnerAcotrMiners(owner common.Address, actor uint64) (*types.Transaction, error) {
	return _TokenFactory.Contract.RemoveTokenOwnerAcotrMiners(&_TokenFactory.TransactOpts, owner, actor)
}

// RemoveTopAcotrMiners is a paid mutator transaction binding the contract method 0x5bca526f.
//
// Solidity: function removeTopAcotrMiners(uint64 actor) returns()
func (_TokenFactory *TokenFactoryTransactor) RemoveTopAcotrMiners(opts *bind.TransactOpts, actor uint64) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "removeTopAcotrMiners", actor)
}

// RemoveTopAcotrMiners is a paid mutator transaction binding the contract method 0x5bca526f.
//
// Solidity: function removeTopAcotrMiners(uint64 actor) returns()
func (_TokenFactory *TokenFactorySession) RemoveTopAcotrMiners(actor uint64) (*types.Transaction, error) {
	return _TokenFactory.Contract.RemoveTopAcotrMiners(&_TokenFactory.TransactOpts, actor)
}

// RemoveTopAcotrMiners is a paid mutator transaction binding the contract method 0x5bca526f.
//
// Solidity: function removeTopAcotrMiners(uint64 actor) returns()
func (_TokenFactory *TokenFactoryTransactorSession) RemoveTopAcotrMiners(actor uint64) (*types.Transaction, error) {
	return _TokenFactory.Contract.RemoveTopAcotrMiners(&_TokenFactory.TransactOpts, actor)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RenounceRole(&_TokenFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RenounceRole(&_TokenFactory.TransactOpts, role, account)
}

// ReturnPledge is a paid mutator transaction binding the contract method 0x45c9b091.
//
// Solidity: function returnPledge(address controller) payable returns()
func (_TokenFactory *TokenFactoryTransactor) ReturnPledge(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "returnPledge", controller)
}

// ReturnPledge is a paid mutator transaction binding the contract method 0x45c9b091.
//
// Solidity: function returnPledge(address controller) payable returns()
func (_TokenFactory *TokenFactorySession) ReturnPledge(controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.ReturnPledge(&_TokenFactory.TransactOpts, controller)
}

// ReturnPledge is a paid mutator transaction binding the contract method 0x45c9b091.
//
// Solidity: function returnPledge(address controller) payable returns()
func (_TokenFactory *TokenFactoryTransactorSession) ReturnPledge(controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.ReturnPledge(&_TokenFactory.TransactOpts, controller)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RevokeRole(&_TokenFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenFactory *TokenFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RevokeRole(&_TokenFactory.TransactOpts, role, account)
}

// RewardController is a paid mutator transaction binding the contract method 0x2a29d109.
//
// Solidity: function rewardController(address controller) payable returns()
func (_TokenFactory *TokenFactoryTransactor) RewardController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "rewardController", controller)
}

// RewardController is a paid mutator transaction binding the contract method 0x2a29d109.
//
// Solidity: function rewardController(address controller) payable returns()
func (_TokenFactory *TokenFactorySession) RewardController(controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RewardController(&_TokenFactory.TransactOpts, controller)
}

// RewardController is a paid mutator transaction binding the contract method 0x2a29d109.
//
// Solidity: function rewardController(address controller) payable returns()
func (_TokenFactory *TokenFactoryTransactorSession) RewardController(controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.RewardController(&_TokenFactory.TransactOpts, controller)
}

// SetAcotrMinerController is a paid mutator transaction binding the contract method 0x771f7c4c.
//
// Solidity: function setAcotrMinerController(uint64 actor, address controller) returns()
func (_TokenFactory *TokenFactoryTransactor) SetAcotrMinerController(opts *bind.TransactOpts, actor uint64, controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setAcotrMinerController", actor, controller)
}

// SetAcotrMinerController is a paid mutator transaction binding the contract method 0x771f7c4c.
//
// Solidity: function setAcotrMinerController(uint64 actor, address controller) returns()
func (_TokenFactory *TokenFactorySession) SetAcotrMinerController(actor uint64, controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetAcotrMinerController(&_TokenFactory.TransactOpts, actor, controller)
}

// SetAcotrMinerController is a paid mutator transaction binding the contract method 0x771f7c4c.
//
// Solidity: function setAcotrMinerController(uint64 actor, address controller) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetAcotrMinerController(actor uint64, controller common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetAcotrMinerController(&_TokenFactory.TransactOpts, actor, controller)
}

// SetB is a paid mutator transaction binding the contract method 0x7004be72.
//
// Solidity: function setB(address b) returns()
func (_TokenFactory *TokenFactoryTransactor) SetB(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setB", b)
}

// SetB is a paid mutator transaction binding the contract method 0x7004be72.
//
// Solidity: function setB(address b) returns()
func (_TokenFactory *TokenFactorySession) SetB(b common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetB(&_TokenFactory.TransactOpts, b)
}

// SetB is a paid mutator transaction binding the contract method 0x7004be72.
//
// Solidity: function setB(address b) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetB(b common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetB(&_TokenFactory.TransactOpts, b)
}

// SetCallPair is a paid mutator transaction binding the contract method 0x1051d539.
//
// Solidity: function setCallPair(address callPair_) returns()
func (_TokenFactory *TokenFactoryTransactor) SetCallPair(opts *bind.TransactOpts, callPair_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setCallPair", callPair_)
}

// SetCallPair is a paid mutator transaction binding the contract method 0x1051d539.
//
// Solidity: function setCallPair(address callPair_) returns()
func (_TokenFactory *TokenFactorySession) SetCallPair(callPair_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetCallPair(&_TokenFactory.TransactOpts, callPair_)
}

// SetCallPair is a paid mutator transaction binding the contract method 0x1051d539.
//
// Solidity: function setCallPair(address callPair_) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetCallPair(callPair_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetCallPair(&_TokenFactory.TransactOpts, callPair_)
}

// SetDefaultAdmin is a paid mutator transaction binding the contract method 0x15a41150.
//
// Solidity: function setDefaultAdmin(address defaultAdmin_) returns()
func (_TokenFactory *TokenFactoryTransactor) SetDefaultAdmin(opts *bind.TransactOpts, defaultAdmin_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setDefaultAdmin", defaultAdmin_)
}

// SetDefaultAdmin is a paid mutator transaction binding the contract method 0x15a41150.
//
// Solidity: function setDefaultAdmin(address defaultAdmin_) returns()
func (_TokenFactory *TokenFactorySession) SetDefaultAdmin(defaultAdmin_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetDefaultAdmin(&_TokenFactory.TransactOpts, defaultAdmin_)
}

// SetDefaultAdmin is a paid mutator transaction binding the contract method 0x15a41150.
//
// Solidity: function setDefaultAdmin(address defaultAdmin_) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetDefaultAdmin(defaultAdmin_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetDefaultAdmin(&_TokenFactory.TransactOpts, defaultAdmin_)
}

// SetF is a paid mutator transaction binding the contract method 0xd478d047.
//
// Solidity: function setF(address f) returns()
func (_TokenFactory *TokenFactoryTransactor) SetF(opts *bind.TransactOpts, f common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setF", f)
}

// SetF is a paid mutator transaction binding the contract method 0xd478d047.
//
// Solidity: function setF(address f) returns()
func (_TokenFactory *TokenFactorySession) SetF(f common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetF(&_TokenFactory.TransactOpts, f)
}

// SetF is a paid mutator transaction binding the contract method 0xd478d047.
//
// Solidity: function setF(address f) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetF(f common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetF(&_TokenFactory.TransactOpts, f)
}

// SetFM is a paid mutator transaction binding the contract method 0x0021110b.
//
// Solidity: function setFM(address f) returns()
func (_TokenFactory *TokenFactoryTransactor) SetFM(opts *bind.TransactOpts, f common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setFM", f)
}

// SetFM is a paid mutator transaction binding the contract method 0x0021110b.
//
// Solidity: function setFM(address f) returns()
func (_TokenFactory *TokenFactorySession) SetFM(f common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetFM(&_TokenFactory.TransactOpts, f)
}

// SetFM is a paid mutator transaction binding the contract method 0x0021110b.
//
// Solidity: function setFM(address f) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetFM(f common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetFM(&_TokenFactory.TransactOpts, f)
}

// SetIdoStartTime is a paid mutator transaction binding the contract method 0xf3cafc9f.
//
// Solidity: function setIdoStartTime(uint256 timeStamp) returns()
func (_TokenFactory *TokenFactoryTransactor) SetIdoStartTime(opts *bind.TransactOpts, timeStamp *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setIdoStartTime", timeStamp)
}

// SetIdoStartTime is a paid mutator transaction binding the contract method 0xf3cafc9f.
//
// Solidity: function setIdoStartTime(uint256 timeStamp) returns()
func (_TokenFactory *TokenFactorySession) SetIdoStartTime(timeStamp *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetIdoStartTime(&_TokenFactory.TransactOpts, timeStamp)
}

// SetIdoStartTime is a paid mutator transaction binding the contract method 0xf3cafc9f.
//
// Solidity: function setIdoStartTime(uint256 timeStamp) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetIdoStartTime(timeStamp *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetIdoStartTime(&_TokenFactory.TransactOpts, timeStamp)
}

// SetRewardStartTime is a paid mutator transaction binding the contract method 0xb9ebd8b5.
//
// Solidity: function setRewardStartTime(uint256 timeStamp) returns()
func (_TokenFactory *TokenFactoryTransactor) SetRewardStartTime(opts *bind.TransactOpts, timeStamp *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setRewardStartTime", timeStamp)
}

// SetRewardStartTime is a paid mutator transaction binding the contract method 0xb9ebd8b5.
//
// Solidity: function setRewardStartTime(uint256 timeStamp) returns()
func (_TokenFactory *TokenFactorySession) SetRewardStartTime(timeStamp *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetRewardStartTime(&_TokenFactory.TransactOpts, timeStamp)
}

// SetRewardStartTime is a paid mutator transaction binding the contract method 0xb9ebd8b5.
//
// Solidity: function setRewardStartTime(uint256 timeStamp) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetRewardStartTime(timeStamp *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetRewardStartTime(&_TokenFactory.TransactOpts, timeStamp)
}

// SetSwapFactory is a paid mutator transaction binding the contract method 0xc34e596f.
//
// Solidity: function setSwapFactory(address swapFactory_) returns()
func (_TokenFactory *TokenFactoryTransactor) SetSwapFactory(opts *bind.TransactOpts, swapFactory_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setSwapFactory", swapFactory_)
}

// SetSwapFactory is a paid mutator transaction binding the contract method 0xc34e596f.
//
// Solidity: function setSwapFactory(address swapFactory_) returns()
func (_TokenFactory *TokenFactorySession) SetSwapFactory(swapFactory_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetSwapFactory(&_TokenFactory.TransactOpts, swapFactory_)
}

// SetSwapFactory is a paid mutator transaction binding the contract method 0xc34e596f.
//
// Solidity: function setSwapFactory(address swapFactory_) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetSwapFactory(swapFactory_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetSwapFactory(&_TokenFactory.TransactOpts, swapFactory_)
}

// SetT is a paid mutator transaction binding the contract method 0x11ad7b85.
//
// Solidity: function setT(address t) returns()
func (_TokenFactory *TokenFactoryTransactor) SetT(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setT", t)
}

// SetT is a paid mutator transaction binding the contract method 0x11ad7b85.
//
// Solidity: function setT(address t) returns()
func (_TokenFactory *TokenFactorySession) SetT(t common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetT(&_TokenFactory.TransactOpts, t)
}

// SetT is a paid mutator transaction binding the contract method 0x11ad7b85.
//
// Solidity: function setT(address t) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetT(t common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetT(&_TokenFactory.TransactOpts, t)
}

// SetTokenExchange is a paid mutator transaction binding the contract method 0xf0d97469.
//
// Solidity: function setTokenExchange(address tokenExchange_) returns()
func (_TokenFactory *TokenFactoryTransactor) SetTokenExchange(opts *bind.TransactOpts, tokenExchange_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setTokenExchange", tokenExchange_)
}

// SetTokenExchange is a paid mutator transaction binding the contract method 0xf0d97469.
//
// Solidity: function setTokenExchange(address tokenExchange_) returns()
func (_TokenFactory *TokenFactorySession) SetTokenExchange(tokenExchange_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetTokenExchange(&_TokenFactory.TransactOpts, tokenExchange_)
}

// SetTokenExchange is a paid mutator transaction binding the contract method 0xf0d97469.
//
// Solidity: function setTokenExchange(address tokenExchange_) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetTokenExchange(tokenExchange_ common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetTokenExchange(&_TokenFactory.TransactOpts, tokenExchange_)
}

// SetTu is a paid mutator transaction binding the contract method 0x02624f68.
//
// Solidity: function setTu(address t) returns()
func (_TokenFactory *TokenFactoryTransactor) SetTu(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setTu", t)
}

// SetTu is a paid mutator transaction binding the contract method 0x02624f68.
//
// Solidity: function setTu(address t) returns()
func (_TokenFactory *TokenFactorySession) SetTu(t common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetTu(&_TokenFactory.TransactOpts, t)
}

// SetTu is a paid mutator transaction binding the contract method 0x02624f68.
//
// Solidity: function setTu(address t) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetTu(t common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetTu(&_TokenFactory.TransactOpts, t)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address recommendAccount) returns()
func (_TokenFactory *TokenFactoryTransactor) SetUser(opts *bind.TransactOpts, recommendAccount common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setUser", recommendAccount)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address recommendAccount) returns()
func (_TokenFactory *TokenFactorySession) SetUser(recommendAccount common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetUser(&_TokenFactory.TransactOpts, recommendAccount)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address recommendAccount) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetUser(recommendAccount common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetUser(&_TokenFactory.TransactOpts, recommendAccount)
}

// TransferOwnerFilecoinMinerController is a paid mutator transaction binding the contract method 0x38e05d70.
//
// Solidity: function transferOwnerFilecoinMinerController(address controller, address newOwner) returns()
func (_TokenFactory *TokenFactoryTransactor) TransferOwnerFilecoinMinerController(opts *bind.TransactOpts, controller common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "transferOwnerFilecoinMinerController", controller, newOwner)
}

// TransferOwnerFilecoinMinerController is a paid mutator transaction binding the contract method 0x38e05d70.
//
// Solidity: function transferOwnerFilecoinMinerController(address controller, address newOwner) returns()
func (_TokenFactory *TokenFactorySession) TransferOwnerFilecoinMinerController(controller common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.TransferOwnerFilecoinMinerController(&_TokenFactory.TransactOpts, controller, newOwner)
}

// TransferOwnerFilecoinMinerController is a paid mutator transaction binding the contract method 0x38e05d70.
//
// Solidity: function transferOwnerFilecoinMinerController(address controller, address newOwner) returns()
func (_TokenFactory *TokenFactoryTransactorSession) TransferOwnerFilecoinMinerController(controller common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.TransferOwnerFilecoinMinerController(&_TokenFactory.TransactOpts, controller, newOwner)
}

// WithdrawFilecoinMinerController is a paid mutator transaction binding the contract method 0xc19b91b5.
//
// Solidity: function withdrawFilecoinMinerController(address controller, address account) returns()
func (_TokenFactory *TokenFactoryTransactor) WithdrawFilecoinMinerController(opts *bind.TransactOpts, controller common.Address, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "withdrawFilecoinMinerController", controller, account)
}

// WithdrawFilecoinMinerController is a paid mutator transaction binding the contract method 0xc19b91b5.
//
// Solidity: function withdrawFilecoinMinerController(address controller, address account) returns()
func (_TokenFactory *TokenFactorySession) WithdrawFilecoinMinerController(controller common.Address, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.WithdrawFilecoinMinerController(&_TokenFactory.TransactOpts, controller, account)
}

// WithdrawFilecoinMinerController is a paid mutator transaction binding the contract method 0xc19b91b5.
//
// Solidity: function withdrawFilecoinMinerController(address controller, address account) returns()
func (_TokenFactory *TokenFactoryTransactorSession) WithdrawFilecoinMinerController(controller common.Address, account common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.WithdrawFilecoinMinerController(&_TokenFactory.TransactOpts, controller, account)
}

// TokenFactoryActorMinerControllerCreatedIterator is returned from FilterActorMinerControllerCreated and is used to iterate over the raw logs and unpacked data for ActorMinerControllerCreated events raised by the TokenFactory contract.
type TokenFactoryActorMinerControllerCreatedIterator struct {
	Event *TokenFactoryActorMinerControllerCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryActorMinerControllerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryActorMinerControllerCreated)
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
		it.Event = new(TokenFactoryActorMinerControllerCreated)
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
func (it *TokenFactoryActorMinerControllerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryActorMinerControllerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryActorMinerControllerCreated represents a ActorMinerControllerCreated event raised by the TokenFactory contract.
type TokenFactoryActorMinerControllerCreated struct {
	Owner      common.Address
	Controller common.Address
	Miner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActorMinerControllerCreated is a free log retrieval operation binding the contract event 0x99d6b81ff21cfdea80b83b8248c362b04b65eb3c8c8ee0cb2522858bf97216ef.
//
// Solidity: event ActorMinerControllerCreated(address indexed owner, address controller, address miner)
func (_TokenFactory *TokenFactoryFilterer) FilterActorMinerControllerCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryActorMinerControllerCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "ActorMinerControllerCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryActorMinerControllerCreatedIterator{contract: _TokenFactory.contract, event: "ActorMinerControllerCreated", logs: logs, sub: sub}, nil
}

// WatchActorMinerControllerCreated is a free log subscription operation binding the contract event 0x99d6b81ff21cfdea80b83b8248c362b04b65eb3c8c8ee0cb2522858bf97216ef.
//
// Solidity: event ActorMinerControllerCreated(address indexed owner, address controller, address miner)
func (_TokenFactory *TokenFactoryFilterer) WatchActorMinerControllerCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryActorMinerControllerCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "ActorMinerControllerCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryActorMinerControllerCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "ActorMinerControllerCreated", log); err != nil {
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

// ParseActorMinerControllerCreated is a log parse operation binding the contract event 0x99d6b81ff21cfdea80b83b8248c362b04b65eb3c8c8ee0cb2522858bf97216ef.
//
// Solidity: event ActorMinerControllerCreated(address indexed owner, address controller, address miner)
func (_TokenFactory *TokenFactoryFilterer) ParseActorMinerControllerCreated(log types.Log) (*TokenFactoryActorMinerControllerCreated, error) {
	event := new(TokenFactoryActorMinerControllerCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "ActorMinerControllerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryCheckOwnerCreatedIterator is returned from FilterCheckOwnerCreated and is used to iterate over the raw logs and unpacked data for CheckOwnerCreated events raised by the TokenFactory contract.
type TokenFactoryCheckOwnerCreatedIterator struct {
	Event *TokenFactoryCheckOwnerCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryCheckOwnerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryCheckOwnerCreated)
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
		it.Event = new(TokenFactoryCheckOwnerCreated)
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
func (it *TokenFactoryCheckOwnerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryCheckOwnerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryCheckOwnerCreated represents a CheckOwnerCreated event raised by the TokenFactory contract.
type TokenFactoryCheckOwnerCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCheckOwnerCreated is a free log retrieval operation binding the contract event 0xfd82e0f64d53e0e384ad8554e4bdabd5e9b0f166507a60f19d539442fe8f3930.
//
// Solidity: event CheckOwnerCreated(address indexed owner)
func (_TokenFactory *TokenFactoryFilterer) FilterCheckOwnerCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryCheckOwnerCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "CheckOwnerCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryCheckOwnerCreatedIterator{contract: _TokenFactory.contract, event: "CheckOwnerCreated", logs: logs, sub: sub}, nil
}

// WatchCheckOwnerCreated is a free log subscription operation binding the contract event 0xfd82e0f64d53e0e384ad8554e4bdabd5e9b0f166507a60f19d539442fe8f3930.
//
// Solidity: event CheckOwnerCreated(address indexed owner)
func (_TokenFactory *TokenFactoryFilterer) WatchCheckOwnerCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryCheckOwnerCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "CheckOwnerCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryCheckOwnerCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "CheckOwnerCreated", log); err != nil {
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

// ParseCheckOwnerCreated is a log parse operation binding the contract event 0xfd82e0f64d53e0e384ad8554e4bdabd5e9b0f166507a60f19d539442fe8f3930.
//
// Solidity: event CheckOwnerCreated(address indexed owner)
func (_TokenFactory *TokenFactoryFilterer) ParseCheckOwnerCreated(log types.Log) (*TokenFactoryCheckOwnerCreated, error) {
	event := new(TokenFactoryCheckOwnerCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "CheckOwnerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryOwnerAcotrMinersCreatedIterator is returned from FilterOwnerAcotrMinersCreated and is used to iterate over the raw logs and unpacked data for OwnerAcotrMinersCreated events raised by the TokenFactory contract.
type TokenFactoryOwnerAcotrMinersCreatedIterator struct {
	Event *TokenFactoryOwnerAcotrMinersCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryOwnerAcotrMinersCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryOwnerAcotrMinersCreated)
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
		it.Event = new(TokenFactoryOwnerAcotrMinersCreated)
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
func (it *TokenFactoryOwnerAcotrMinersCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryOwnerAcotrMinersCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryOwnerAcotrMinersCreated represents a OwnerAcotrMinersCreated event raised by the TokenFactory contract.
type TokenFactoryOwnerAcotrMinersCreated struct {
	Owner common.Address
	Miner uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAcotrMinersCreated is a free log retrieval operation binding the contract event 0x239fcb798b324dc59dc9298cdf37f865a80014c9a5c7fa1abc2951e16565b0a4.
//
// Solidity: event OwnerAcotrMinersCreated(address indexed owner, uint64 miner)
func (_TokenFactory *TokenFactoryFilterer) FilterOwnerAcotrMinersCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryOwnerAcotrMinersCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "OwnerAcotrMinersCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryOwnerAcotrMinersCreatedIterator{contract: _TokenFactory.contract, event: "OwnerAcotrMinersCreated", logs: logs, sub: sub}, nil
}

// WatchOwnerAcotrMinersCreated is a free log subscription operation binding the contract event 0x239fcb798b324dc59dc9298cdf37f865a80014c9a5c7fa1abc2951e16565b0a4.
//
// Solidity: event OwnerAcotrMinersCreated(address indexed owner, uint64 miner)
func (_TokenFactory *TokenFactoryFilterer) WatchOwnerAcotrMinersCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryOwnerAcotrMinersCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "OwnerAcotrMinersCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryOwnerAcotrMinersCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "OwnerAcotrMinersCreated", log); err != nil {
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

// ParseOwnerAcotrMinersCreated is a log parse operation binding the contract event 0x239fcb798b324dc59dc9298cdf37f865a80014c9a5c7fa1abc2951e16565b0a4.
//
// Solidity: event OwnerAcotrMinersCreated(address indexed owner, uint64 miner)
func (_TokenFactory *TokenFactoryFilterer) ParseOwnerAcotrMinersCreated(log types.Log) (*TokenFactoryOwnerAcotrMinersCreated, error) {
	event := new(TokenFactoryOwnerAcotrMinersCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "OwnerAcotrMinersCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryOwnerCreatedIterator is returned from FilterOwnerCreated and is used to iterate over the raw logs and unpacked data for OwnerCreated events raised by the TokenFactory contract.
type TokenFactoryOwnerCreatedIterator struct {
	Event *TokenFactoryOwnerCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryOwnerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryOwnerCreated)
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
		it.Event = new(TokenFactoryOwnerCreated)
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
func (it *TokenFactoryOwnerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryOwnerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryOwnerCreated represents a OwnerCreated event raised by the TokenFactory contract.
type TokenFactoryOwnerCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerCreated is a free log retrieval operation binding the contract event 0x6b91e48d218f29b663f7f44cd7c857e3f04242025cf53d2e2b95715611a06190.
//
// Solidity: event OwnerCreated(address indexed owner)
func (_TokenFactory *TokenFactoryFilterer) FilterOwnerCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryOwnerCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "OwnerCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryOwnerCreatedIterator{contract: _TokenFactory.contract, event: "OwnerCreated", logs: logs, sub: sub}, nil
}

// WatchOwnerCreated is a free log subscription operation binding the contract event 0x6b91e48d218f29b663f7f44cd7c857e3f04242025cf53d2e2b95715611a06190.
//
// Solidity: event OwnerCreated(address indexed owner)
func (_TokenFactory *TokenFactoryFilterer) WatchOwnerCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryOwnerCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "OwnerCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryOwnerCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "OwnerCreated", log); err != nil {
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

// ParseOwnerCreated is a log parse operation binding the contract event 0x6b91e48d218f29b663f7f44cd7c857e3f04242025cf53d2e2b95715611a06190.
//
// Solidity: event OwnerCreated(address indexed owner)
func (_TokenFactory *TokenFactoryFilterer) ParseOwnerCreated(log types.Log) (*TokenFactoryOwnerCreated, error) {
	event := new(TokenFactoryOwnerCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "OwnerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TokenFactory contract.
type TokenFactoryRoleAdminChangedIterator struct {
	Event *TokenFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryRoleAdminChanged)
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
		it.Event = new(TokenFactoryRoleAdminChanged)
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
func (it *TokenFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the TokenFactory contract.
type TokenFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenFactory *TokenFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenFactoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryRoleAdminChangedIterator{contract: _TokenFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenFactory *TokenFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryRoleAdminChanged)
				if err := _TokenFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TokenFactory *TokenFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*TokenFactoryRoleAdminChanged, error) {
	event := new(TokenFactoryRoleAdminChanged)
	if err := _TokenFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TokenFactory contract.
type TokenFactoryRoleGrantedIterator struct {
	Event *TokenFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *TokenFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryRoleGranted)
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
		it.Event = new(TokenFactoryRoleGranted)
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
func (it *TokenFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryRoleGranted represents a RoleGranted event raised by the TokenFactory contract.
type TokenFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenFactory *TokenFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenFactoryRoleGrantedIterator, error) {

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

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryRoleGrantedIterator{contract: _TokenFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenFactory *TokenFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryRoleGranted)
				if err := _TokenFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TokenFactory *TokenFactoryFilterer) ParseRoleGranted(log types.Log) (*TokenFactoryRoleGranted, error) {
	event := new(TokenFactoryRoleGranted)
	if err := _TokenFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TokenFactory contract.
type TokenFactoryRoleRevokedIterator struct {
	Event *TokenFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TokenFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryRoleRevoked)
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
		it.Event = new(TokenFactoryRoleRevoked)
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
func (it *TokenFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryRoleRevoked represents a RoleRevoked event raised by the TokenFactory contract.
type TokenFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenFactory *TokenFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenFactoryRoleRevokedIterator, error) {

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

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryRoleRevokedIterator{contract: _TokenFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenFactory *TokenFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryRoleRevoked)
				if err := _TokenFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TokenFactory *TokenFactoryFilterer) ParseRoleRevoked(log types.Log) (*TokenFactoryRoleRevoked, error) {
	event := new(TokenFactoryRoleRevoked)
	if err := _TokenFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the TokenFactory contract.
type TokenFactoryTokenCreatedIterator struct {
	Event *TokenFactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryTokenCreated)
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
		it.Event = new(TokenFactoryTokenCreated)
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
func (it *TokenFactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryTokenCreated represents a TokenCreated event raised by the TokenFactory contract.
type TokenFactoryTokenCreated struct {
	Owner common.Address
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address indexed owner, address token)
func (_TokenFactory *TokenFactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryTokenCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "TokenCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTokenCreatedIterator{contract: _TokenFactory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address indexed owner, address token)
func (_TokenFactory *TokenFactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryTokenCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "TokenCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryTokenCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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

// ParseTokenCreated is a log parse operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address indexed owner, address token)
func (_TokenFactory *TokenFactoryFilterer) ParseTokenCreated(log types.Log) (*TokenFactoryTokenCreated, error) {
	event := new(TokenFactoryTokenCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryTopAcotrMinersCreatedIterator is returned from FilterTopAcotrMinersCreated and is used to iterate over the raw logs and unpacked data for TopAcotrMinersCreated events raised by the TokenFactory contract.
type TokenFactoryTopAcotrMinersCreatedIterator struct {
	Event *TokenFactoryTopAcotrMinersCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryTopAcotrMinersCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryTopAcotrMinersCreated)
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
		it.Event = new(TokenFactoryTopAcotrMinersCreated)
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
func (it *TokenFactoryTopAcotrMinersCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryTopAcotrMinersCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryTopAcotrMinersCreated represents a TopAcotrMinersCreated event raised by the TokenFactory contract.
type TokenFactoryTopAcotrMinersCreated struct {
	Owner common.Address
	Miner uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTopAcotrMinersCreated is a free log retrieval operation binding the contract event 0x373c9c22870ef86b7809dae5c5d4d303ca09d5f1982e55fd0c042dd6a1181ad1.
//
// Solidity: event TopAcotrMinersCreated(address indexed owner, uint64 miner)
func (_TokenFactory *TokenFactoryFilterer) FilterTopAcotrMinersCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryTopAcotrMinersCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "TopAcotrMinersCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTopAcotrMinersCreatedIterator{contract: _TokenFactory.contract, event: "TopAcotrMinersCreated", logs: logs, sub: sub}, nil
}

// WatchTopAcotrMinersCreated is a free log subscription operation binding the contract event 0x373c9c22870ef86b7809dae5c5d4d303ca09d5f1982e55fd0c042dd6a1181ad1.
//
// Solidity: event TopAcotrMinersCreated(address indexed owner, uint64 miner)
func (_TokenFactory *TokenFactoryFilterer) WatchTopAcotrMinersCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryTopAcotrMinersCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "TopAcotrMinersCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryTopAcotrMinersCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "TopAcotrMinersCreated", log); err != nil {
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

// ParseTopAcotrMinersCreated is a log parse operation binding the contract event 0x373c9c22870ef86b7809dae5c5d4d303ca09d5f1982e55fd0c042dd6a1181ad1.
//
// Solidity: event TopAcotrMinersCreated(address indexed owner, uint64 miner)
func (_TokenFactory *TokenFactoryFilterer) ParseTopAcotrMinersCreated(log types.Log) (*TokenFactoryTopAcotrMinersCreated, error) {
	event := new(TokenFactoryTopAcotrMinersCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "TopAcotrMinersCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryUnionTokenCreatedIterator is returned from FilterUnionTokenCreated and is used to iterate over the raw logs and unpacked data for UnionTokenCreated events raised by the TokenFactory contract.
type TokenFactoryUnionTokenCreatedIterator struct {
	Event *TokenFactoryUnionTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryUnionTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryUnionTokenCreated)
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
		it.Event = new(TokenFactoryUnionTokenCreated)
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
func (it *TokenFactoryUnionTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryUnionTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryUnionTokenCreated represents a UnionTokenCreated event raised by the TokenFactory contract.
type TokenFactoryUnionTokenCreated struct {
	Owner common.Address
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnionTokenCreated is a free log retrieval operation binding the contract event 0xd27443b9893d79dfde526bef62b0884dfec1c1776207749c749f0f3618917dd7.
//
// Solidity: event UnionTokenCreated(address indexed owner, address token)
func (_TokenFactory *TokenFactoryFilterer) FilterUnionTokenCreated(opts *bind.FilterOpts, owner []common.Address) (*TokenFactoryUnionTokenCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "UnionTokenCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryUnionTokenCreatedIterator{contract: _TokenFactory.contract, event: "UnionTokenCreated", logs: logs, sub: sub}, nil
}

// WatchUnionTokenCreated is a free log subscription operation binding the contract event 0xd27443b9893d79dfde526bef62b0884dfec1c1776207749c749f0f3618917dd7.
//
// Solidity: event UnionTokenCreated(address indexed owner, address token)
func (_TokenFactory *TokenFactoryFilterer) WatchUnionTokenCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryUnionTokenCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "UnionTokenCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryUnionTokenCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "UnionTokenCreated", log); err != nil {
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

// ParseUnionTokenCreated is a log parse operation binding the contract event 0xd27443b9893d79dfde526bef62b0884dfec1c1776207749c749f0f3618917dd7.
//
// Solidity: event UnionTokenCreated(address indexed owner, address token)
func (_TokenFactory *TokenFactoryFilterer) ParseUnionTokenCreated(log types.Log) (*TokenFactoryUnionTokenCreated, error) {
	event := new(TokenFactoryUnionTokenCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "UnionTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryUserCreatedIterator is returned from FilterUserCreated and is used to iterate over the raw logs and unpacked data for UserCreated events raised by the TokenFactory contract.
type TokenFactoryUserCreatedIterator struct {
	Event *TokenFactoryUserCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryUserCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryUserCreated)
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
		it.Event = new(TokenFactoryUserCreated)
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
func (it *TokenFactoryUserCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryUserCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryUserCreated represents a UserCreated event raised by the TokenFactory contract.
type TokenFactoryUserCreated struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserCreated is a free log retrieval operation binding the contract event 0x0b0376a109cbb578b709f85f6a7befcdac3ac1ab251c99ede87f30c9572ac4d3.
//
// Solidity: event UserCreated(address indexed user)
func (_TokenFactory *TokenFactoryFilterer) FilterUserCreated(opts *bind.FilterOpts, user []common.Address) (*TokenFactoryUserCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "UserCreated", userRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryUserCreatedIterator{contract: _TokenFactory.contract, event: "UserCreated", logs: logs, sub: sub}, nil
}

// WatchUserCreated is a free log subscription operation binding the contract event 0x0b0376a109cbb578b709f85f6a7befcdac3ac1ab251c99ede87f30c9572ac4d3.
//
// Solidity: event UserCreated(address indexed user)
func (_TokenFactory *TokenFactoryFilterer) WatchUserCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryUserCreated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "UserCreated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryUserCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "UserCreated", log); err != nil {
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

// ParseUserCreated is a log parse operation binding the contract event 0x0b0376a109cbb578b709f85f6a7befcdac3ac1ab251c99ede87f30c9572ac4d3.
//
// Solidity: event UserCreated(address indexed user)
func (_TokenFactory *TokenFactoryFilterer) ParseUserCreated(log types.Log) (*TokenFactoryUserCreated, error) {
	event := new(TokenFactoryUserCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "UserCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
