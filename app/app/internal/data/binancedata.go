package data

import (
	"dhb/app/app/internal/biz"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type KLineMOne struct {
	ID                  int64     `gorm:"primarykey;type:int"`
	StartTime           int64     `gorm:"type:bigint;not null"`
	EndTime             int64     `gorm:"type:bigint;not null"`
	StartPrice          float64   `gorm:"type:decimal(65,20);not null"`
	TopPrice            float64   `gorm:"type:decimal(65,20);not null"`
	LowPrice            float64   `gorm:"type:decimal(65,20);not null"`
	EndPrice            float64   `gorm:"type:decimal(65,20);not null"`
	DealTotalAmount     float64   `gorm:"type:decimal(65,20);not null"`
	DealAmount          float64   `gorm:"type:decimal(65,20);not null"`
	DealTotal           int64     `gorm:"type:int;not null"`
	DealSelfTotalAmount float64   `gorm:"type:decimal(65,20);not null"`
	DealSelfAmount      float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt           time.Time `gorm:"type:datetime;not null"`
	UpdatedAt           time.Time `gorm:"type:datetime;not null"`
}

type KLineMOneRepo struct {
	data *Data
	log  *log.Helper
}

func NewKLineMOneRepo(data *Data, logger log.Logger) biz.KLineMOneRepo {
	return &KLineMOneRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (k *KLineMOneRepo) RequestBinanceMinuteKLinesData(symbol string, startTime string, endTime string, interval string, limit string) ([]*biz.KLineMOne, error) {
	apiUrl := "https://api.binance.us/api/v3/klines"
	// URL param
	data := url.Values{}
	data.Set("symbol", symbol)
	data.Set("interval", interval)
	data.Set("startTime", startTime)
	data.Set("endTime", endTime)
	data.Set("limit", limit)

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = data.Encode() // URL encode
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	//fmt.Println(u.String())
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var i [][]interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	res := make([]*biz.KLineMOne, 0)
	for _, v := range i {
		startTimeTmp, _ := strconv.ParseInt(strconv.FormatFloat(v[0].(float64), 'f', -1, 64), 10, 64)
		endTimeTmp, _ := strconv.ParseInt(strconv.FormatFloat(v[6].(float64), 'f', -1, 64), 10, 64)
		dealTotalTmp, _ := strconv.ParseInt(strconv.FormatFloat(v[8].(float64), 'f', -1, 64), 10, 64)
		startPriceTmp, _ := strconv.ParseFloat(v[1].(string), 64)
		endPriceTmp, _ := strconv.ParseFloat(v[4].(string), 64)
		topPriceTmp, _ := strconv.ParseFloat(v[2].(string), 64)
		lowPriceTmp, _ := strconv.ParseFloat(v[3].(string), 64)
		dealTotalAmountTmp, _ := strconv.ParseFloat(v[5].(string), 64)
		dealAmountTmp, _ := strconv.ParseFloat(v[7].(string), 64)
		dealSelfTotalAmountTmp, _ := strconv.ParseFloat(v[9].(string), 64)
		dealSelfAmountTmp, _ := strconv.ParseFloat(v[10].(string), 64)

		res = append(res, &biz.KLineMOne{
			StartTime:           startTimeTmp,
			StartPrice:          startPriceTmp,
			EndPrice:            endPriceTmp,
			TopPrice:            topPriceTmp,
			LowPrice:            lowPriceTmp,
			EndTime:             endTimeTmp,
			DealTotalAmount:     dealTotalAmountTmp,
			DealAmount:          dealAmountTmp,
			DealTotal:           dealTotalTmp,
			DealSelfTotalAmount: dealSelfTotalAmountTmp,
			DealSelfAmount:      dealSelfAmountTmp,
		})
	}

	return res, err
}
