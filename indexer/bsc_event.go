package indexer

import (
	"github.com/bnb-chain/staking-reward-example/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const eventDelegatedTopic = "0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04"
const eventUndelegatedTopic = "0x3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802"
const eventRedelegatedTopic = "0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4"

func isTargetEvent(targetTopic ethcommon.Hash, l types.Log) bool {
	return targetTopic.String() == l.Topics[0].String()
}

func parseDelegatedEvent(l types.Log, abi *abi.ABI) (*contracts.StakeHubDelegated, string, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventDelegatedTopic), l) {
		return nil, "", nil
	}

	delegate := &contracts.StakeHubDelegated{}
	err := abi.UnpackIntoInterface(delegate, "Delegated", l.Data)
	if err != nil {
		return nil, "", err
	}
	delegate.OperatorAddress = ethcommon.BytesToAddress(l.Topics[1].Bytes())
	delegate.Delegator = ethcommon.BytesToAddress(l.Topics[2].Bytes())
	return delegate, l.TxHash.Hex(), nil
}

func parseUndelegatedEvent(l types.Log, abi *abi.ABI) (*contracts.StakeHubUndelegated, string, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventUndelegatedTopic), l) {
		return nil, "", nil
	}

	undelegate := &contracts.StakeHubUndelegated{}
	err := abi.UnpackIntoInterface(undelegate, "Undelegated", l.Data)
	if err != nil {
		return nil, "", err
	}
	undelegate.OperatorAddress = ethcommon.BytesToAddress(l.Topics[1].Bytes())
	undelegate.Delegator = ethcommon.BytesToAddress(l.Topics[2].Bytes())
	return undelegate, l.TxHash.Hex(), nil
}

func parseRedelegatedEvent(l types.Log, abi *abi.ABI) (*contracts.StakeHubRedelegated, string, error) {
	if !isTargetEvent(ethcommon.HexToHash(eventRedelegatedTopic), l) {
		return nil, "", nil
	}

	redelegate := &contracts.StakeHubRedelegated{}
	err := abi.UnpackIntoInterface(redelegate, "Redelegated", l.Data)
	if err != nil {
		return nil, "", err
	}
	redelegate.SrcValidator = ethcommon.BytesToAddress(l.Topics[1].Bytes())
	redelegate.DstValidator = ethcommon.BytesToAddress(l.Topics[2].Bytes())
	redelegate.Delegator = ethcommon.BytesToAddress(l.Topics[3].Bytes())
	return redelegate, l.TxHash.Hex(), nil
}
