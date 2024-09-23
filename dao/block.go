package dao

import (
	"context"
	"github.com/bnb-chain/staking-reward-example/database"
	"gorm.io/gorm"
)

type BlockDao interface {
	Create(context context.Context, block *database.Block) error
	Max(context context.Context) (database.Block, error)
}

type dbBlockDao struct {
	db *gorm.DB
}

func NewDbBlockDao(db *gorm.DB) BlockDao {
	return &dbBlockDao{
		db: db,
	}
}

func (dao *dbBlockDao) Create(context context.Context, block *database.Block) error {
	if err := dao.db.Create(block).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbBlockDao) Max(context context.Context) (database.Block, error) {
	var block database.Block
	if err := dao.db.Raw("select * from blocks order by height desc limit 1").First(&block).Error; err != nil {
		return block, err
	}
	return block, nil
}
