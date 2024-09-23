package database

import (
	"fmt"
	"github.com/bnb-chain/staking-reward-example/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config.DBConfig) (*gorm.DB, error) {
	if config.DBDialect == "sqlite3" {
		db, err := gorm.Open(sqlite.Open(config.DBPath), &gorm.Config{})

		if err = db.AutoMigrate(&Block{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Operation{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Delegation{}); err != nil {
			panic(err)
		}

		return db.Debug(), err
	} else if config.DBDialect == "mysql" {
		dbPath := fmt.Sprintf("%s:%s@%s", config.Username, config.Password, config.DBPath)
		db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		dbConfig, err := db.DB()
		if err != nil {
			panic(err)
		}
		dbConfig.SetMaxIdleConns(config.MaxIdleConns)
		dbConfig.SetMaxOpenConns(config.MaxOpenConns)

		if err = db.AutoMigrate(&Block{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Operation{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Delegation{}); err != nil {
			panic(err)
		}
		return db.Debug(), nil
	} else {
		return nil, fmt.Errorf("dialect %s not supported", config.DBDialect)
	}
}
