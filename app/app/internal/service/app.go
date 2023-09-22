package service

import (
	"context"
	"crypto/ecdsa"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/biz"
	"dhb/app/app/internal/conf"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"
)

// AppService service.
type AppService struct {
	v1.UnimplementedAppServer

	bduc *biz.BinanceDataUsecase
	log  *log.Helper
	ca   *conf.Auth
}

// NewAppService new a service.
func NewAppService(bduc *biz.BinanceDataUsecase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{bduc: bduc, log: log.NewHelper(logger), ca: ca}
}

func (a *AppService) FilUsdt(ctx context.Context, req *v1.FilUsdtRequest) (*v1.FilUsdtReply, error) {
	return a.bduc.FilUsdt(ctx, req)
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
		log.Fatal(err)
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
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

type eth struct {
	CoinId string
	Usd    float64
}

func requestHbsResult() (float64, error) {
	//apiUrl := "https://api-testnet.bscscan.com/api"
	apiUrl := "https://be.api.hbsswap.com/market/coin/rates"
	// URL param
	data := url.Values{}

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return 0, err
	}
	u.RawQuery = data.Encode() // URL encode
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(u.String())
	if err != nil {
		return 0, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var i struct {
		Data []*eth `json:"Data"`
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return 0, err
	}

	var price float64
	for _, v := range i.Data {
		if "HBS(BEP20)" == v.CoinId { // 接收者
			price = v.Usd
		}
	}

	return price, err
}
