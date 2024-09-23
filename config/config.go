package config

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	DBDialect    string `json:"db_dialect"`
	DBPath       string `json:"db_path"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type Config struct {
	ValidatorOperatorAddr string `json:"validator_operator_addr"`

	BscRpcAddrs          []string `json:"bsc_rpc_addrs"`
	BscBlocksForFinality int      `json:"bsc_blocks_for_finality"`
	BscStartHeight       uint64   `json:"bsc_start_height"`

	DBConfig *DBConfig `json:"db_config"`
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}
