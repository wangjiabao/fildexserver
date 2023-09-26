package service

import (
	"context"
	"crypto/ecdsa"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/biz"
	"dhb/app/app/internal/conf"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
)

// AppService service.
type AppService struct {
	v1.UnimplementedAppServer

	bduc *biz.BinanceDataUsecase
	ruc  *biz.RecordUseCase
	log  *log.Helper
	ca   *conf.Auth
}

// NewAppService new a service.
func NewAppService(bduc *biz.BinanceDataUsecase, ruc *biz.RecordUseCase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{bduc: bduc, ruc: ruc, log: log.NewHelper(logger), ca: ca}
}

func (a *AppService) FilUsdt(ctx context.Context, req *v1.FilUsdtRequest) (*v1.FilUsdtReply, error) {
	return a.bduc.FilUsdt(ctx, req)
}

func (a *AppService) SetPerSecondDFilTotal(ctx context.Context, req *v1.SetPerSecondDFilTotalRequest) (*v1.SetPerSecondDFilTotalReply, error) {
	var (
		totalSupply string
		err         error
	)

	totalSupply, err = GetDFilTotalSupply()
	if nil != err {
		return nil, err
	}

	return a.ruc.SetPerSecondDFilTotal(ctx, totalSupply)
}

func (a *AppService) GetPerSecondDFilTotal(ctx context.Context, req *v1.GetPerSecondDFilTotalRequest) (*v1.GetPerSecondDFilTotalReply, error) {
	return a.ruc.GetPerSecondDFilTotal(ctx, req)
}

func (a *AppService) SetPerSecondPairInfo(ctx context.Context, req *v1.SetPerSecondPairInfoRequest) (*v1.SetPerSecondPairInfoReply, error) {
	var (
		reserve0 string
		reserve1 string
		err      error
	)

	reserve0, reserve1, err = GetReserves("0xb160f026bA4B5da6dcC2e55aD4c9Cfb426084d93")
	if nil != err {
		return nil, err
	}

	return a.ruc.SetPerSecondPairInfo(ctx, "0xb160f026bA4B5da6dcC2e55aD4c9Cfb426084d93", reserve0, reserve1)
}

func (a *AppService) GetPerSecondPairInfo(ctx context.Context, req *v1.GetPerSecondPairInfoRequest) (*v1.GetPerSecondPairInfoReply, error) {
	return a.ruc.GetPerSecondPairInfo(ctx, req)
}

func (a *AppService) ReqContract(ctx context.Context, req *v1.ReqContractRequest) (*v1.ReqContractReply, error) {

	if "1" == req.ContractType {
		var (
			checkOwners     []string
			ownerFilBalance string
			err             error
		)
		checkOwners, err = GetFactoryCheckOwner()
		if nil != err {
			return nil, err
		}

		for _, vCheckOwners := range checkOwners {
			ownerFilBalance, err = GetExchangeOwnerFilBalance(vCheckOwners)
			if nil != err {
				fmt.Println("未获取到信息")
				continue
			}

			if ownerFilBalance < "1000000000000000000000" {
				fmt.Println(vCheckOwners, ownerFilBalance, "流动性不足")
				continue
			}

			_, err = createTokenOwner(vCheckOwners)
			if nil != err {
				fmt.Println("未获取到信息")
				continue
			}
		}
	} else if "2" == req.ContractType {
		var (
			err error
		)
		_, err = managerReward()
		if nil != err {
			//fmt.Println("未获取到信息", err)
		}
	}

	return nil, nil
}

//func (a *AppService) GetTrade(ctx context.Context, req *v1.GetTradeRequest) (*v1.GetTradeReply, error) {
//	// 在上下文 context 中取出 claims 对象
//	var (
//		amountB        int64
//		tmpValue       int64
//		hbs            float64
//		amountFloatHbs float64
//		amountFloatCsd float64
//		csd            string
//		err            error
//	)
//
//	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
//	amountFloatCsd = amountFloat * 10000000000
//	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloatCsd, 'f', -1, 64), 10, 64)
//	if 10000000000 > amount {
//		return nil, errors.New(500, "ERROR_TOKEN", "输入错误")
//	}
//
//	csd, err = GetAmountOut(req.SendBody.Amount + "000000000000000000")
//	if nil != err {
//		fmt.Println(2)
//		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
//	}
//	lenValue := len(csd)
//	if 10 > lenValue {
//		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
//	}
//	tmpValue, _ = strconv.ParseInt(csd[0:lenValue-8], 10, 64)
//	if 0 == tmpValue {
//		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
//	}
//
//	hbs, err = requestHbsResult()
//	if nil != err {
//		fmt.Println(1)
//		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
//	}
//	amountFloatHbs = amountFloat * 10
//	amountB = int64(amountFloatHbs / hbs * 10000000000)
//	if 0 >= amountB {
//		return nil, errors.New(500, "ERROR_TOKEN", "币价错误")
//	}
//
//	return &v1.GetTradeReply{
//		AmountCsd: fmt.Sprintf("%.4f", float64(tmpValue)/float64(10000000000)),
//		AmountHbs: fmt.Sprintf("%.4f", float64(amountB)/float64(10000000000)),
//	}, nil
//}
//
//func (a *AppService) TokenWithdraw(ctx context.Context, req *v1.TokenWithdrawRequest) (*v1.TokenWithdrawReply, error) {
//
//	var (
//		err error
//	)
//	for i := 0; i <= 5; i++ {
//		tmpUrl1 := "https://bnb-bscnews.rpc.blxrbdn.com/"
//		_, err = tokenWithdraw(tmpUrl1, 56)
//		if err == nil {
//			break
//		} else if "insufficient funds for gas * price + value" == err.Error() {
//			fmt.Println(5555, err)
//		} else if "execution reverted: ERC20: transfer amount exceeds balance" == err.Error() {
//			fmt.Println(4444, err)
//			break
//		} else if "execution reverted: time limit" == err.Error() {
//			fmt.Println(4441, err)
//			break
//		} else {
//			if 0 == i {
//				tmpUrl1 = "https://bsc-dataseed4.binance.org/"
//			} else if 1 == i {
//				tmpUrl1 = "https://bsc-dataseed3.binance.org"
//			} else if 2 == i {
//				tmpUrl1 = "https://bsc-dataseed2.binance.org"
//			} else if 3 == i {
//				tmpUrl1 = "https://bsc-dataseed1.binance.org"
//			} else if 4 == i {
//				tmpUrl1 = "https://bsc-dataseed.binance.org"
//			}
//			fmt.Println(3333, err)
//			time.Sleep(3 * time.Second)
//		}
//	}
//
//	return &v1.TokenWithdrawReply{}, nil
//}

func tokenWithdraw(requestUrl string, chainId int64) (bool, error) {

	client, err := ethclient.Dial(requestUrl)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return false, err
	}

	tokenAddress := common.HexToAddress("0xFC13153Bb4D285939FD23c7899eAdD785fBf6aA2")
	instance, err := NewTokenWithdraw(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA("")
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainId))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	//var res *types.Transaction
	_, err = instance.WithdrawSx(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasPrice: gasPrice,
		GasLimit: 0,
	})
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	//fmt.Println(res.Hash())

	return true, nil
}

func GetFactoryCheckOwner() ([]string, error) {
	var (
		balsString []string
	)

	url1 := "https://api.node.glif.io"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return balsString, err
		}

		tokenAddress := common.HexToAddress("0x8505616465ed0ba07Ead8eBA6295a9B16E5a325E")
		instance, err := NewTokenFactory(tokenAddress, client)
		if err != nil {
			return balsString, err
		}

		bals, err := instance.GetCheckOwners(&bind.CallOpts{})
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		balsString = make([]string, 0)
		for _, vBals := range bals {
			balsString = append(balsString, vBals.String())
		}

		break
	}

	return balsString, nil
}

func GetExchangeOwnerFilBalance(account string) (string, error) {
	var (
		balString string
	)

	url1 := "https://api.node.glif.io"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return balString, err
		}

		tokenAddress := common.HexToAddress("0x9Ec5b4Ce25dfAE269151f926dBAc7ff10A33a34a")
		instance, err := NewTokenExchange(tokenAddress, client)
		if err != nil {
			return balString, err
		}

		bals, err := instance.GetTokenOwnerFilBalance(&bind.CallOpts{}, common.HexToAddress(account))
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}
		balString = bals.String()
		break
	}

	return balString, nil
}

func createTokenOwner(account string) (string, error) {

	requestUrl := "https://api.node.glif.io"

	client, err := ethclient.Dial(requestUrl)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x8505616465ed0ba07Ead8eBA6295a9B16E5a325E")
	instance, err := NewTokenFactory(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA("")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(314))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	_, err = instance.CreateTokenOwner(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasLimit: 0,
	}, common.HexToAddress(account))
	if err != nil {
		fmt.Println(1, account, err)
		return "", err
	}

	return "", nil
}

func managerReward() (string, error) {

	requestUrl := "https://api.node.glif.io"

	client, err := ethclient.Dial(requestUrl)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x8b026AfaE06A0F86284427EBa48962248D22570b")
	instance, err := NewManagerFil(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA("")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(314))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	_, err = instance.Reward(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasLimit: 0,
	})
	if err != nil {
		return "", err
	}

	return "", nil
}

func GetReserves(address string) (string, string, error) {
	var (
		balAString string
		balBString string
	)

	url1 := "https://api.node.glif.io"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return balAString, balBString, err
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewSwapPair(tokenAddress, client)
		if err != nil {
			return balAString, balBString, err
		}

		bals, err := instance.GetReserves(&bind.CallOpts{})
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}
		balAString = bals.Reserve0.String()
		balBString = bals.Reserve1.String()
		break
	}

	return balAString, balBString, nil
}

func GetDFilTotalSupply() (string, error) {
	var (
		balString string
	)

	url1 := "https://api.node.glif.io"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return balString, err
		}

		tokenAddress := common.HexToAddress("0x8CE4501A420FcAE4B0fBAe73cC07f3A763B3aA7B")
		instance, err := NewDfil(tokenAddress, client)
		if err != nil {
			return balString, err
		}

		bals, err := instance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}
		balString = bals.String()
		break
	}

	return balString, nil
}

func GetAmountOut(strAmount string) (string, error) {

	var balString string
	url1 := "https://bnb-bscnews.rpc.blxrbdn.com/"

	for i := 4; i < 16; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return "", err
		}

		tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
		instance, err := NewPancakerouterv2(tokenAddress, client)
		if err != nil {
			return "", err
		}

		addresses := make([]common.Address, 0)
		addresses = append(addresses, common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), common.HexToAddress("0xfAd476cd33Ed9213ED0a2F4c20f6865A98bf0a8B"))
		amount, _ := new(big.Int).SetString(strAmount, 10)

		bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
		if err != nil {
			fmt.Println(err)
			if 0 == i%4 {
				url1 = "https://bsc-dataseed4.binance.org"
			} else if 1 == i%4 {
				url1 = "https://bsc-dataseed1.binance.org"
			} else if 2 == i%4 {
				url1 = "https://bsc-dataseed.binance.org"
			} else if 3 == i%4 {
				url1 = "https://bsc-dataseed3.binance.org"
			}
			continue
		}
		balString = bals[1].String()
		break
	}

	return balString, nil
}

func GetAmountOut1(strAmount string) (string, error) {

	var balString string
	url1 := "https://bnb-bscnews.rpc.blxrbdn.com/"

	client, err := ethclient.Dial(url1)
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
	instance, err := NewPancakerouterv2(tokenAddress, client)
	if err != nil {
		return "", err
	}

	addresses := make([]common.Address, 0)
	addresses = append(addresses, common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), common.HexToAddress("0x538ac017aa01ba9665052660ea5783ba91a48092"))
	amount, _ := new(big.Int).SetString(strAmount, 10)

	bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	balString = bals[1].String()
	return balString, nil
}
