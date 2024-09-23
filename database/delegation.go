package database

import (
	"github.com/shopspring/decimal"
)

type Delegation struct {
	Id            int64           `json:"id" gorm:"primaryKey"`
	Validator     string          `json:"validator"  gorm:"index:idx_delegation_operator_delegator;not null;size:42;"`
	Delegator     string          `json:"delegator"  gorm:"uniqueIndex:idx_delegation_operator_delegator;;not null;size:42;"`
	Amount        decimal.Decimal `json:"amount" gorm:"type:decimal(65,0);"`
	UpdatedHeight int64           `json:"updatedHeight"`
	CreatedAt     int64           `json:"createdAt" gorm:"NOT NULL;default:0"`
}
