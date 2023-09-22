package biz

import (
	"context"
	v1 "dhb/app/app/api"
	"github.com/go-kratos/kratos/v2/log"
)

type KLineMOne struct {
	StartTime           int64
	EndTime             int64
	StartPrice          float64
	TopPrice            float64
	LowPrice            float64
	EndPrice            float64
	DealTotalAmount     float64
	DealAmount          float64
	DealTotal           int64
	DealSelfTotalAmount float64
	DealSelfAmount      float64
}

type KLineMOneRepo interface {
	RequestBinanceMinuteKLinesData(symbol string, startTime string, endTime string, interval string, limit string) ([]*KLineMOne, error)
}

// BinanceDataUsecase is a BinanceData usecase.
type BinanceDataUsecase struct {
	klineMOneRepo KLineMOneRepo
	tx            Transaction
	log           *log.Helper
}

// NewBinanceDataUsecase new a BinanceData usecase.
func NewBinanceDataUsecase(klineMOneRepo KLineMOneRepo, tx Transaction, logger log.Logger) *BinanceDataUsecase {
	return &BinanceDataUsecase{klineMOneRepo: klineMOneRepo, tx: tx, log: log.NewHelper(logger)}
}

func (bdu *BinanceDataUsecase) FilUsdt(ctx context.Context, req *v1.FilUsdtRequest) (*v1.FilUsdtReply, error) {
	var (
		klineMOne []*KLineMOne
		err       error
	)

	klineMOne, err = bdu.klineMOneRepo.RequestBinanceMinuteKLinesData(
		req.Symbol,
		req.StartTime,
		req.EndTime,
		req.Interval,
		req.Limit,
	)
	if nil != err {
		return nil, err
	}

	resKLineMOne := &v1.FilUsdtReply{DataListK: make([]*v1.FilUsdtReply_ListK, 0)}
	for _, v := range klineMOne {
		resKLineMOne.DataListK = append(resKLineMOne.DataListK, &v1.FilUsdtReply_ListK{
			StartTime:           v.StartTime,
			EndTime:             v.EndTime,
			StartPrice:          v.StartPrice,
			TopPrice:            v.TopPrice,
			LowPrice:            v.LowPrice,
			EndPrice:            v.EndPrice,
			DealTotalAmount:     v.DealTotalAmount,
			DealAmount:          v.DealAmount,
			DealTotal:           v.DealTotal,
			DealSelfTotalAmount: v.DealSelfTotalAmount,
			DealSelfAmount:      v.DealSelfAmount,
		})
	}

	return resKLineMOne, nil
}
