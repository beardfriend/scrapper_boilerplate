package infra

import (
	"boilerplate/config"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var onceDB sync.Once

func Database() *gorm.DB {
	if db != nil {
		return db
	}
	onceDB.Do(func() {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true", config.MYSQLUserName, config.MYSQLPassword, config.MYSQLHost, config.MYSQLPort, config.MYSQLDatabase)
		if config.MYSQLIsSSL {
			dsn += "&tls=true"
		}

		dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
		if err != nil {
			panic(err)
		}

		db = dbInstance
	})

	return db
}
