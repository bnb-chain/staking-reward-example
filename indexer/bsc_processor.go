package indexer

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/bnb-chain/staking-reward-example/client"
	"github.com/bnb-chain/staking-reward-example/contracts"
	"github.com/bnb-chain/staking-reward-example/dao"
	"github.com/bnb-chain/staking-reward-example/database"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var StakeHubAddr = ethcommon.HexToAddress("0x0000000000000000000000000000000000002002")

type BscProcessor struct {
	validatorOperatorAddr string

	client      *client.BscCompositeClients
	stakeHubAbi *abi.ABI

	blockDao      dao.BlockDao
	operationDao  dao.OperationDao
	delegationDao dao.DelegationDao
}

func NewBscBlockProcessor(
	validatorOperatorAddr string,
	client *client.BscCompositeClients,
	blockDao dao.BlockDao,
	operationDao dao.OperationDao,
	delegationDao dao.DelegationDao,
) *BscProcessor {
	stakeHubAbi, err := abi.JSON(strings.NewReader(contracts.StakeHubMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}

	return &BscProcessor{
		validatorOperatorAddr: validatorOperatorAddr,

		client:      client,
		stakeHubAbi: &stakeHubAbi,

		blockDao:      blockDao,
		operationDao:  operationDao,
		delegationDao: delegationDao,
	}
}

func (p *BscProcessor) GetDatabaseBlockHeight() (uint64, error) {
	maxBlock, err := p.blockDao.Max(context.Background())
	if err != nil {
		return 0, err
	}
	return maxBlock.Height, nil
}

func (p *BscProcessor) GetBlockchainBlockHeight() (uint64, error) {
	return p.client.GetLatestFinalizedBlockHeightWithRetry() // get finalized block height
}

// Process processes the block at the given height and stores the data in the database.
// This is for demo purpose only, no database transaction is used for simplicity.
func (p *BscProcessor) Process(blockHeight uint64) error {
	err := p.processStake(blockHeight)
	if err != nil {
		return err
	}

	err = p.blockDao.Create(context.Background(), &database.Block{
		Height: blockHeight,
	})
	return err
}

func (p *BscProcessor) processStake(blockHeight uint64) error {
	topics := []ethcommon.Hash{
		ethcommon.HexToHash(eventDelegatedTopic),
		ethcommon.HexToHash(eventUndelegatedTopic),
		ethcommon.HexToHash(eventRedelegatedTopic),
	}

	header, err := p.client.GetBlockHeader(blockHeight)
	if err != nil {
		return err
	}

	logs, err := p.client.QueryChainLogs(blockHeight, blockHeight, topics, StakeHubAddr)
	if err != nil {
		return err
	}

	// to process logs in order
	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i].Index < logs[j].Index
	})

	for _, l := range logs {
		switch l.Topics[0].String() {
		case eventDelegatedTopic:
			err = p.handleDelegated(blockHeight, header, l)
		case eventUndelegatedTopic:
			err = p.handleUndelegated(blockHeight, header, l)
		case eventRedelegatedTopic:
			err = p.handleRedelegated(blockHeight, header, l)
		default:
			fmt.Println("unknown topic", l.Topics[0].String())
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *BscProcessor) handleDelegated(blockHeight uint64, header *types.Header, l types.Log) error {
	event, txHash, err := parseDelegatedEvent(l, p.stakeHubAbi)
	if err != nil {
		return err
	}
	if event == nil {
		return nil
	}

	if event.OperatorAddress.Hex() != p.validatorOperatorAddr {
		return nil
	}

	delegation, err := p.delegationDao.Get(context.Background(), event.OperatorAddress.Hex(), event.Delegator.Hex())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		delegation.Amount = delegation.Amount.Add(decimal.NewFromBigInt(event.BnbAmount, 1))
		delegation.UpdatedHeight = int64(blockHeight)
		err = p.delegationDao.Update(context.Background(), &delegation)
		if err != nil {
			return err
		}
	} else {
		err = p.delegationDao.Create(context.Background(), &database.Delegation{
			Validator:     event.OperatorAddress.Hex(),
			Delegator:     event.Delegator.Hex(),
			Amount:        decimal.NewFromBigInt(event.BnbAmount, 1),
			UpdatedHeight: int64(blockHeight),
			CreatedAt:     int64(header.Time),
		})
		if err != nil {
			return err
		}
	}

	err = p.operationDao.Create(context.Background(), &database.Operation{
		Type:          database.OperationDelegate,
		SrcValidator:  event.OperatorAddress.Hex(),
		DstValidator:  "",
		Delegator:     event.Delegator.Hex(),
		Amount:        decimal.NewFromBigInt(event.BnbAmount, 1),
		UpdatedHeight: int64(blockHeight),
		TxHash:        txHash,
		CreatedAt:     int64(header.Time),
	})
	return err
}

func (p *BscProcessor) handleUndelegated(blockHeight uint64, header *types.Header, l types.Log) error {
	event, txHash, err := parseUndelegatedEvent(l, p.stakeHubAbi)
	if err != nil {
		return err
	}
	if event == nil {
		return nil
	}

	if event.OperatorAddress.Hex() != p.validatorOperatorAddr {
		return nil
	}

	delegation, err := p.delegationDao.Get(context.Background(), event.OperatorAddress.Hex(), event.Delegator.Hex())
	if err != nil {
		return err
	}

	delegation.Amount = delegation.Amount.Sub(decimal.NewFromBigInt(event.BnbAmount, 1))
	delegation.UpdatedHeight = int64(blockHeight)
	err = p.delegationDao.Update(context.Background(), &delegation)
	if err != nil {
		return err
	}

	err = p.operationDao.Create(context.Background(), &database.Operation{
		Type:          database.OperationUndelegate,
		SrcValidator:  event.OperatorAddress.Hex(),
		DstValidator:  "",
		Delegator:     event.Delegator.Hex(),
		Amount:        decimal.NewFromBigInt(event.BnbAmount, 1),
		UpdatedHeight: int64(blockHeight),
		TxHash:        txHash,
		CreatedAt:     int64(header.Time),
	})
	return err
}

func (p *BscProcessor) handleRedelegated(blockHeight uint64, header *types.Header, l types.Log) error {
	event, txHash, err := parseRedelegatedEvent(l, p.stakeHubAbi)
	if err != nil {
		return err
	}
	if event == nil {
		return nil
	}

	if event.SrcValidator.Hex() != p.validatorOperatorAddr && event.DstValidator.Hex() != p.validatorOperatorAddr {
		return nil
	}

	srcUndelegation, err := p.delegationDao.Get(context.Background(), event.SrcValidator.Hex(), event.Delegator.Hex())
	if err != nil {
		return err
	}

	srcUndelegation.Amount = srcUndelegation.Amount.Sub(decimal.NewFromBigInt(event.BnbAmount, 1))
	srcUndelegation.UpdatedHeight = int64(blockHeight)
	err = p.delegationDao.Update(context.Background(), &srcUndelegation)
	if err != nil {
		return err
	}

	dstDelegation, err := p.delegationDao.Get(context.Background(), event.DstValidator.Hex(), event.Delegator.Hex())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		dstDelegation.Amount = dstDelegation.Amount.Add(decimal.NewFromBigInt(event.BnbAmount, 1))
		dstDelegation.UpdatedHeight = int64(blockHeight)
		err = p.delegationDao.Update(context.Background(), &dstDelegation)
		if err != nil {
			return err
		}
	} else {
		err = p.delegationDao.Create(context.Background(), &database.Delegation{
			Validator:     event.DstValidator.Hex(),
			Delegator:     event.Delegator.Hex(),
			Amount:        decimal.NewFromBigInt(event.BnbAmount, 1),
			UpdatedHeight: int64(blockHeight),
			CreatedAt:     int64(header.Time),
		})
		if err != nil {
			return err
		}
	}

	err = p.operationDao.Create(context.Background(), &database.Operation{
		Type:          database.OperationRedelegate,
		SrcValidator:  event.SrcValidator.Hex(),
		DstValidator:  event.DstValidator.Hex(),
		Delegator:     event.Delegator.Hex(),
		Amount:        decimal.NewFromBigInt(event.BnbAmount, 1),
		UpdatedHeight: int64(blockHeight),
		TxHash:        txHash,
		CreatedAt:     int64(header.Time),
	})
	return err
}
