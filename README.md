# BSC Staking Reward Demo

This is a demo project for BSC staking reward tracking. Please do not use it in your production environment.

## Config File
Configure the necessary parameters for the project.

```json
{
  "validator_operator_addr": "the validator operator address you want to track rewards for, in hex format",
  "bsc_rpc_addrs": [
    "bsc rpc address list"
  ],
  "bsc_blocks_for_finality": "parameter for fast finality, should be between 2 and 21",
  "bsc_start_height": "the start height for tracking rewards",
  "db_config": {
    "db_dialect": "sqlite3 or mysql",
    "db_path": "database file path or connection string"
  }
}
```

Please refer to the [example-config.json](example-config.json) file for example.

## Indexer
To track the delegation/undelegation/re-delegation events of the validator configured in the config file.
The database will store the events and do the simple calculations for the delegations.

Example command:

```shell
make build
./build/main --action index --config-path example-config.json 
```

## Querier
To get the latest staking rewards for a given delegator.

Example command:

```shell
make build
./build/main --action query --delegator 0x4B17e695512c6D1f25FD00334c88A662C56b80d1 --config-path example-config.json 
```

## More

* Please refer to [the document](https://docs.bnbchain.org/bc-fusion/developers/staking/) for more about the BSC staking.
* Please refer to [the document](https://docs.bnbchain.org/bc-fusion/developers/system-contracts/) for related ABI files.