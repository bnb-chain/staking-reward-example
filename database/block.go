package database

import "time"

type Block struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Height    uint64    `json:"height" gorm:"NOT NULL;index:idx_block_height"`
	CreatedAt time.Time `json:"createdAt" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
}
