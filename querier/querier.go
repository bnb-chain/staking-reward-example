package querier

import (
	"context"
	"fmt"
	"github.com/bnb-chain/staking-reward-example/client"
	"github.com/bnb-chain/staking-reward-example/contracts"
	"github.com/bnb-chain/staking-reward-example/dao"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

var StakeHubAddr = ethcommon.HexToAddress("0x0000000000000000000000000000000000002002")

type Querier struct {
	validatorOperatorAddr string

	client         *client.BscCompositeClients
	stakeHubAbi    *abi.ABI
	stakeCreditAbi *abi.ABI

	delegationDao dao.DelegationDao
}

func NewQuerier(validatorOperatorAddr string,
	client *client.BscCompositeClients,
	delegationDao dao.DelegationDao,
) *Querier {
	stakeHubAbi, err := abi.JSON(strings.NewReader(contracts.StakeHubMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}

	stakeCreditAbi, err := abi.JSON(strings.NewReader(contracts.StakeCreditMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}

	return &Querier{
		validatorOperatorAddr: validatorOperatorAddr,

		client: client,

		stakeHubAbi:    &stakeHubAbi,
		stakeCreditAbi: &stakeCreditAbi,

		delegationDao: delegationDao,
	}
}

func (q *Querier) GetDelegatorReward(delegatorAddr string) {
	stakeHub, err := contracts.NewStakeHub(StakeHubAddr, q.client.GetEthClient())
	if err != nil {
		panic(err)
	}

	creditContractAddr, err := stakeHub.GetValidatorCreditContract(nil, ethcommon.HexToAddress(q.validatorOperatorAddr))
	if err != nil {
		panic(err)
	}
	fmt.Println("Credit Contract Addr:", creditContractAddr.String())

	stakeCredit, err := contracts.NewStakeCredit(creditContractAddr, q.client.GetEthClient())
	if err != nil {
		panic(err)
	}

	pooledBNB, err := stakeCredit.GetPooledBNB(nil, ethcommon.HexToAddress(delegatorAddr))
	if err != nil {
		panic(err)
	}
	fmt.Println("Pooled BNB:", pooledBNB.String())

	if pooledBNB.Cmp(ethcommon.Big0) == 0 {
		fmt.Println("No reward, for pooled BNB is zero")
		return
	}

	delegation, err := q.delegationDao.Get(context.Background(), q.validatorOperatorAddr, delegatorAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delegation:", delegation)

	reward := new(big.Int).Sub(pooledBNB, delegation.Amount.BigInt())
	fmt.Println("Reward:", reward)
}
