package database

import (
	"github.com/shopspring/decimal"
)

type OperationType int

const (
	OperationDelegate   OperationType = 1
	OperationUndelegate OperationType = 2
	OperationRedelegate OperationType = 3
)

type Operation struct {
	Id            int64           `json:"id" gorm:"primaryKey"`
	Type          OperationType   `json:"type" gorm:"index:idx_operation_type;not null;"`
	Delegator     string          `json:"delegator"  gorm:"index:idx_operation_delegator;not null;size:42;"`
	SrcValidator  string          `json:"srcValidator"  gorm:"index:idx_operation_src_validator;not null;size:42;"`
	DstValidator  string          `json:"dstValidator"  gorm:"index:idx_operation_dst_validator;"`
	Amount        decimal.Decimal `json:"amount" gorm:"type:decimal(65,0);"`
	TxHash        string          `json:"txHash"`
	UpdatedHeight int64           `json:"updatedHeight"`
	CreatedAt     int64           `json:"createdAt" gorm:"NOT NULL;default:0"`
}
