package dao

import (
	"context"
	"github.com/bnb-chain/staking-reward-example/database"
	"gorm.io/gorm"
)

type DelegationDao interface {
	Create(context context.Context, delegation *database.Delegation) error
	Update(context context.Context, delegation *database.Delegation) error
	Get(context context.Context, validator, delegator string) (database.Delegation, error)
}

type dbDelegationDao struct {
	db *gorm.DB
}

func NewDbDelegationDao(db *gorm.DB) DelegationDao {
	return &dbDelegationDao{
		db: db,
	}
}

func (dao *dbDelegationDao) Create(context context.Context, delegation *database.Delegation) error {
	if err := dao.db.Create(delegation).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbDelegationDao) Update(context context.Context, delegation *database.Delegation) error {
	if err := dao.db.Save(delegation).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbDelegationDao) Get(context context.Context, validator, delegator string) (database.Delegation, error) {
	var Delegation = database.Delegation{}
	if err := dao.db.Where("validator = ? and delegator = ?", validator, delegator).Take(&Delegation).Error; err != nil {
		return Delegation, err
	}
	return Delegation, nil
}
