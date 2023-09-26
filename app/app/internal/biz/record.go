package biz

import (
	"context"
	v1 "dhb/app/app/api"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type EthUserRecord struct {
	ID       int64
	UserId   int64
	Hash     string
	Status   string
	Type     string
	Amount   string
	CoinType string
}

type Location struct {
	ID           int64
	UserId       int64
	Status       string
	CurrentLevel int64
	Usdt         int64
	Current      int64
	CurrentMax   int64
	Row          int64
	Col          int64
	StopDate     time.Time
	CreatedAt    time.Time
}

type GlobalLock struct {
	ID     int64
	Status int64
}

type PairInfoPerSecond struct {
	ID        int64
	Pair      string
	Reserve0  string
	Reserve1  string
	Timestamp int64
}

type DFilInfoPerSecond struct {
	ID          int64
	TotalSupply string
	Timestamp   int64
}

type RecordUseCase struct {
	ethUserRecordRepo             EthUserRecordRepo
	userRecommendRepo             UserRecommendRepo
	configRepo                    ConfigRepo
	locationRepo                  LocationRepo
	userBalanceRepo               UserBalanceRepo
	userInfoRepo                  UserInfoRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	pairInfoRepo                  PairInfoRepo
	dfilInfoRepo                  DFilInfoRepo
	tx                            Transaction
	log                           *log.Helper
}

type EthUserRecordRepo interface {
	GetEthUserRecordListByHash(ctx context.Context, hash ...string) (map[string]*EthUserRecord, error)
	CreateEthUserRecordListByHash(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
}

type LocationRepo interface {
	CreateLocation(ctx context.Context, rel *Location) (*Location, error)
	GetLocationLast(ctx context.Context) (*Location, error)
	GetLocationDaily(ctx context.Context) ([]*Location, error)
	GetMyLocationLast(ctx context.Context, userId int64) (*Location, error)
	GetMyStopLocationLast(ctx context.Context, userId int64) (*Location, error)
	GetMyLocationRunningLast(ctx context.Context, userId int64) (*Location, error)
	GetLocationsByUserId(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetRewardLocationByRowOrCol(ctx context.Context, row int64, col int64, locationRowConfig int64) ([]*Location, error)
	GetRewardLocationByIds(ctx context.Context, ids ...int64) (map[int64]*Location, error)
	GetLocationMapByIds(ctx context.Context, userIds ...int64) (map[int64][]*Location, error)
	GetLocationByIds(ctx context.Context, userIds ...int64) ([]*Location, error)
	UpdateLocation(ctx context.Context, id int64, status string, current int64, stopDate time.Time) error
	GetLocations(ctx context.Context, b *Pagination, userId int64) ([]*Location, error, int64)
	UpdateLocationRowAndCol(ctx context.Context, id int64) error
	GetLocationsStopNotUpdate(ctx context.Context) ([]*Location, error)
	LockGlobalLocation(ctx context.Context) (bool, error)
	UnLockGlobalLocation(ctx context.Context) (bool, error)
	LockGlobalWithdraw(ctx context.Context) (bool, error)
	UnLockGlobalWithdraw(ctx context.Context) (bool, error)
	GetLockGlobalLocation(ctx context.Context) (*GlobalLock, error)

	GetMyStopLocationsLast(ctx context.Context, userId int64) ([]*LocationNew, error)
}

type PairInfoRepo interface {
	SetPairInfo(ctx context.Context, pair string, reserve0 string, reserve1 string, timestamp int64) error
	GetPairInfoByTime(pair string, start int64, end int64) ([]*PairInfoPerSecond, error)
}

type DFilInfoRepo interface {
	SetDFilInfo(ctx context.Context, totalSupply string, timestamp int64) error
	GetDFilInfoByTime(start int64, end int64) ([]*DFilInfoPerSecond, error)
}

func NewRecordUseCase(
	ethUserRecordRepo EthUserRecordRepo,
	locationRepo LocationRepo,
	userBalanceRepo UserBalanceRepo,
	userRecommendRepo UserRecommendRepo,
	userInfoRepo UserInfoRepo,
	configRepo ConfigRepo,
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo,
	pairInfoRepo PairInfoRepo,
	dfilInfoRepo DFilInfoRepo,
	tx Transaction,
	logger log.Logger) *RecordUseCase {
	return &RecordUseCase{
		ethUserRecordRepo:             ethUserRecordRepo,
		locationRepo:                  locationRepo,
		configRepo:                    configRepo,
		userRecommendRepo:             userRecommendRepo,
		userBalanceRepo:               userBalanceRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		userInfoRepo:                  userInfoRepo,
		pairInfoRepo:                  pairInfoRepo,
		dfilInfoRepo:                  dfilInfoRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger),
	}
}

func (ruc *RecordUseCase) GetEthUserRecordByTxHash(ctx context.Context, txHash ...string) (map[string]*EthUserRecord, error) {
	return ruc.ethUserRecordRepo.GetEthUserRecordListByHash(ctx, txHash...)
}

func (ruc *RecordUseCase) EthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	return true, nil
}

func (ruc *RecordUseCase) LockEthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	var (
		lock bool
	)
	// todo 全局锁
	for i := 0; i < 3; i++ {
		lock, _ = ruc.locationRepo.LockGlobalLocation(ctx)
		if lock {
			return true, nil
		}
		time.Sleep(5 * time.Second)
	}

	return false, nil
}

func (ruc *RecordUseCase) UnLockEthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	return ruc.locationRepo.UnLockGlobalLocation(ctx)
}

func (ruc *RecordUseCase) SetPerSecondDFilTotal(ctx context.Context, totalSupply string) (*v1.SetPerSecondDFilTotalReply, error) {
	var (
		err error
	)
	now := time.Now().UTC()
	timestamp := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC).Unix()

	err = ruc.dfilInfoRepo.SetDFilInfo(ctx, totalSupply, timestamp)
	if nil != err {
		return nil, err
	}

	return &v1.SetPerSecondDFilTotalReply{}, nil
}

func (ruc *RecordUseCase) SetPerSecondPairInfo(ctx context.Context, pair string, reserve0 string, reserve1 string) (*v1.SetPerSecondPairInfoReply, error) {
	var (
		err error
	)
	now := time.Now().UTC()
	timestamp := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC).Unix()

	err = ruc.pairInfoRepo.SetPairInfo(ctx, pair, reserve0, reserve1, timestamp)
	if nil != err {
		return nil, err
	}

	return &v1.SetPerSecondPairInfoReply{}, nil
}

func (ruc *RecordUseCase) GetPerSecondPairInfo(ctx context.Context, req *v1.GetPerSecondPairInfoRequest) (*v1.GetPerSecondPairInfoReply, error) {
	var (
		pairInfos []*PairInfoPerSecond
		endTime   int64
		startTime int64
		err       error
	)

	resPairInfos := &v1.GetPerSecondPairInfoReply{DataListPair: make([]*v1.GetPerSecondPairInfoReply_ListPair, 0)}

	if req.StartTime > req.EndTime {
		return resPairInfos, nil
	}

	startTime = req.StartTime
	endTime = req.EndTime
	if req.StartTime+2592000 < req.EndTime {
		endTime = req.StartTime + 2592000
	}

	pairInfos, err = ruc.pairInfoRepo.GetPairInfoByTime(req.Pair, startTime, endTime)
	if nil != err {
		return nil, err
	}

	for _, v := range pairInfos {
		resPairInfos.DataListPair = append(resPairInfos.DataListPair, &v1.GetPerSecondPairInfoReply_ListPair{
			Time:     v.Timestamp,
			Reserve0: v.Reserve0,
			Reserve1: v.Reserve1,
		})
	}

	return resPairInfos, nil
}

func (ruc *RecordUseCase) GetPerSecondDFilTotal(ctx context.Context, req *v1.GetPerSecondDFilTotalRequest) (*v1.GetPerSecondDFilTotalReply, error) {
	var (
		dfilInfos []*DFilInfoPerSecond
		endTime   int64
		startTime int64
		err       error
	)

	resDFilInfos := &v1.GetPerSecondDFilTotalReply{DataListDFil: make([]*v1.GetPerSecondDFilTotalReply_ListDfil, 0)}

	if req.StartTime > req.EndTime {
		return resDFilInfos, nil
	}

	startTime = req.StartTime
	endTime = req.EndTime
	if req.StartTime+2592000 < req.EndTime {
		endTime = req.StartTime + 2592000
	}

	dfilInfos, err = ruc.dfilInfoRepo.GetDFilInfoByTime(startTime, endTime)
	if nil != err {
		return nil, err
	}

	for _, v := range dfilInfos {
		resDFilInfos.DataListDFil = append(resDFilInfos.DataListDFil, &v1.GetPerSecondDFilTotalReply_ListDfil{
			Time:        v.Timestamp,
			TotalSupply: v.TotalSupply,
		})
	}

	return resDFilInfos, nil
}
