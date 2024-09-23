package main

import (
	"flag"
	"fmt"
	"github.com/bnb-chain/staking-reward-example/client"
	"github.com/bnb-chain/staking-reward-example/config"
	"github.com/bnb-chain/staking-reward-example/dao"
	"github.com/bnb-chain/staking-reward-example/database"
	"github.com/bnb-chain/staking-reward-example/indexer"
	"github.com/bnb-chain/staking-reward-example/querier"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

const (
	flagAction     = "action"
	flagConfigPath = "config-path"
	flagDelegator  = "delegator"
)

func initFlags() {
	flag.String(flagAction, "", "index or query")
	flag.String(flagConfigPath, "", "config path")
	flag.String(flagDelegator, "", "delegator address, only for query action")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flags error, err=%s", err))
	}
}

func printUsage() {
	fmt.Print("usage: ./main --action [index|query] --config-path config_file_path\n")
}

func main() {
	initFlags()

	configFilePath := viper.GetString(flagConfigPath)
	action := viper.GetString(flagAction)
	if configFilePath == "" || action == "" {
		printUsage()
		return
	}
	fileConfig := config.ParseConfigFromFile(configFilePath)

	db, err := database.ConnectDatabase(fileConfig.DBConfig)
	if err != nil {
		fmt.Printf("connect database error, err=%s \n", err.Error())
		return
	}

	bscClient := client.NewBscCompositeClients(fileConfig.BscRpcAddrs, fileConfig.BscBlocksForFinality)

	blockDao := dao.NewDbBlockDao(db)
	operationDao := dao.NewDbOperationDao(db)
	delegationDao := dao.NewDbDelegationDao(db)

	if strings.ToLower(action) == "index" {
		bscProcessor := indexer.NewBscBlockProcessor(fileConfig.ValidatorOperatorAddr,
			bscClient, blockDao, operationDao, delegationDao)
		bscIndexer := indexer.NewIndexer(bscProcessor, fileConfig.BscStartHeight)

		go bscIndexer.Start()

		select {}
	} else if strings.ToLower(action) == "query" {
		delegator := viper.GetString(flagDelegator)
		bscQuerier := querier.NewQuerier(fileConfig.ValidatorOperatorAddr, bscClient, delegationDao)
		bscQuerier.GetDelegatorReward(delegator)
	} else {
		fmt.Printf("action=%s is invalid, quit now\n", action)
	}
}
