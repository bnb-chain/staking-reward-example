package dao

import (
	"context"
	"github.com/bnb-chain/staking-reward-example/database"
	"gorm.io/gorm"
)

type OperationDao interface {
	Create(context context.Context, Operation *database.Operation) error
}

type dbOperationDao struct {
	db *gorm.DB
}

func NewDbOperationDao(db *gorm.DB) OperationDao {
	return &dbOperationDao{
		db: db,
	}
}

func (dao *dbOperationDao) Create(context context.Context, Operation *database.Operation) error {
	if err := dao.db.Create(Operation).Error; err != nil {
		return err
	}
	return nil
}
