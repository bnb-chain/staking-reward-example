package indexer

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Indexer struct {
	processor   *BscProcessor
	startHeight uint64
}

func NewIndexer(p *BscProcessor, startHeight uint64) *Indexer {
	return &Indexer{processor: p, startHeight: startHeight}
}

func (m *Indexer) Start() {
	ticker := time.NewTicker(200 * time.Millisecond)
	for range ticker.C {
		err := m.run()
		if err != nil {
			fmt.Printf("fail to run monitor with error: %s \n", err)
		}
	}
}

func (m *Indexer) run() error {
	blockchainHeight, err := m.processor.GetBlockchainBlockHeight()
	fmt.Printf("current blockchain height: %d, err: %v \n", blockchainHeight, err)
	if err != nil {
		return err
	}

	dbHeight, err := m.processor.GetDatabaseBlockHeight()
	fmt.Printf("current database height: %d, err: %v \n", dbHeight, err)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if dbHeight < m.startHeight {
		dbHeight = m.startHeight - 1 // to include start height
	}

	if dbHeight < blockchainHeight {
		fmt.Printf("processing height: %d \n", dbHeight+1)
		err = m.processor.Process(dbHeight + 1)
		if err != nil {
			return err
		}
	}

	return nil
}
