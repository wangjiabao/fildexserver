package data

import (
	"context"
	"dhb/app/app/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type EthUserRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Hash      string    `gorm:"type:varchar(100);not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	Amount    string    `gorm:"type:varchar(45);not null"`
	CoinType  string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PairInfoPerSecond struct {
	ID        int64  `gorm:"primarykey;type:int"`
	Pair      string `gorm:"type:varchar(200);not null"`
	Reserve0  string `gorm:"type:varchar(200);not null"`
	Reserve1  string `gorm:"type:varchar(200);not null"`
	Timestamp int64  `gorm:"type:int;not null"`
}

type DfilInfoPerSecond struct {
	ID          int64  `gorm:"primarykey;type:int"`
	TotalSupply string `gorm:"type:varchar(200);not null"`
	Timestamp   int64  `gorm:"type:int;not null"`
}

type EthUserRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewEthUserRecordRepo(data *Data, logger log.Logger) biz.EthUserRecordRepo {
	return &EthUserRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type PairInfoRepo struct {
	data *Data
	log  *log.Helper
}

func NewPairInfoRepo(data *Data, logger log.Logger) biz.PairInfoRepo {
	return &PairInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type DFilInfoRepo struct {
	data *Data
	log  *log.Helper
}

func NewDFilInfoRepo(data *Data, logger log.Logger) biz.DFilInfoRepo {
	return &DFilInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (e *EthUserRecordRepo) GetEthUserRecordListByHash(ctx context.Context, hash ...string) (map[string]*biz.EthUserRecord, error) {
	var ethUserRecord []*EthUserRecord
	if err := e.data.DB(ctx).Table("eth_user_record").Where("hash IN (?)", hash).Find(&ethUserRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	res := make(map[string]*biz.EthUserRecord, 0)
	for _, item := range ethUserRecord {
		res[item.Hash] = &biz.EthUserRecord{
			ID:       item.ID,
			UserId:   item.UserId,
			Hash:     item.Hash,
			Status:   item.Status,
			Type:     item.Type,
			Amount:   item.Amount,
			CoinType: item.CoinType,
		}
	}

	return res, nil
}

func (e *EthUserRecordRepo) CreateEthUserRecordListByHash(ctx context.Context, r *biz.EthUserRecord) (*biz.EthUserRecord, error) {
	var ethUserRecord EthUserRecord
	ethUserRecord.UserId = r.UserId
	ethUserRecord.Hash = r.Hash
	ethUserRecord.Type = r.Type
	ethUserRecord.Status = r.Status
	ethUserRecord.Amount = r.Amount
	ethUserRecord.CoinType = r.CoinType

	res := e.data.DB(ctx).Table("eth_user_record").Create(&ethUserRecord)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ETH_USER_RECORD_ERROR", "以太坊交易信息创建失败")
	}

	return &biz.EthUserRecord{
		ID:       ethUserRecord.ID,
		UserId:   ethUserRecord.UserId,
		Hash:     ethUserRecord.Hash,
		Status:   ethUserRecord.Status,
		Type:     ethUserRecord.Type,
		Amount:   ethUserRecord.Amount,
		CoinType: ethUserRecord.CoinType,
	}, nil
}

func (p *PairInfoRepo) SetPairInfo(ctx context.Context, pair string, reserve0 string, reserve1 string, timestamp int64) error {
	var pairInfoRecord PairInfoPerSecond
	pairInfoRecord.Pair = pair
	pairInfoRecord.Timestamp = timestamp
	pairInfoRecord.Reserve0 = reserve0
	pairInfoRecord.Reserve1 = reserve1

	res := p.data.DB(ctx).Table("pair_info_per_second").Create(&pairInfoRecord)
	if res.Error != nil {
		return errors.New(500, "CREATE_PAIR_INFO_RECORD_ERROR", "交易对每分钟记录创建失败")
	}

	return nil
}

// GetPairInfoByTime .
func (p *PairInfoRepo) GetPairInfoByTime(pair string, start int64, end int64) ([]*biz.PairInfoPerSecond, error) {
	var pairInfoPerSeconds []*PairInfoPerSecond
	if err := p.data.db.Table("pair_info_per_second").
		Where("pair = ? and timestamp >= ? and timestamp <= ?", pair, start, end).
		Find(&pairInfoPerSeconds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("INFO_NOT_FOUND", "location not found")
		}

		return nil, errors.New(500, "PAIR INFO ERROR", err.Error())
	}

	res := make([]*biz.PairInfoPerSecond, 0)
	for _, pairInfoPerSecond := range pairInfoPerSeconds {
		res = append(res, &biz.PairInfoPerSecond{
			ID:        pairInfoPerSecond.ID,
			Pair:      pairInfoPerSecond.Pair,
			Reserve0:  pairInfoPerSecond.Reserve0,
			Reserve1:  pairInfoPerSecond.Reserve1,
			Timestamp: pairInfoPerSecond.Timestamp,
		})
	}

	return res, nil
}

func (d *DFilInfoRepo) SetDFilInfo(ctx context.Context, totalSupply string, timestamp int64) error {
	var dfilInfoRecord DfilInfoPerSecond
	dfilInfoRecord.TotalSupply = totalSupply
	dfilInfoRecord.Timestamp = timestamp

	res := d.data.DB(ctx).Table("dfil_info_per_second").Create(&dfilInfoRecord)
	if res.Error != nil {
		return errors.New(500, "CREATE_DFIL_INFO_RECORD_ERROR", "dfil对每分钟记录创建失败")
	}

	return nil
}

// GetDFilInfoByTime .
func (d *DFilInfoRepo) GetDFilInfoByTime(start int64, end int64) ([]*biz.DFilInfoPerSecond, error) {
	var dfilInfoPerSeconds []*DfilInfoPerSecond
	if err := d.data.db.Table("dfil_info_per_second").
		Where("timestamp >= ? and timestamp <= ?", start, end).
		Find(&dfilInfoPerSeconds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("INFO_NOT_FOUND", "location not found")
		}

		return nil, errors.New(500, "DFIL INFO ERROR", err.Error())
	}

	res := make([]*biz.DFilInfoPerSecond, 0)
	for _, dfilInfoPerSecond := range dfilInfoPerSeconds {
		res = append(res, &biz.DFilInfoPerSecond{
			ID:          dfilInfoPerSecond.ID,
			TotalSupply: dfilInfoPerSecond.TotalSupply,
			Timestamp:   dfilInfoPerSecond.Timestamp,
		})
	}

	return res, nil
}
