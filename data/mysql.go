package data

import (
	"coupon/conf"
	"coupon/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

var tables = []interface{}{
	&model.User{},
	&model.Seller{},
	&model.Coupon{},
}

func initMysql() {
	userName := conf.AppConfig.Database.User
	password := conf.AppConfig.Database.Password
	dbName := conf.AppConfig.Database.DBName
	address := conf.AppConfig.Database.IPAddress
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, address, dbName)

	tablePrefix := ""
	if conf.AppConfig.Mode == "debug" {
		tablePrefix = "test_"
	}
	var err error
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
		},
	})
	if err != nil {
		panic(err)
	}

	innerDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	maxIdleCons := conf.AppConfig.Database.MaxIdle
	maxOpenCons := conf.AppConfig.Database.MaxOpen
	innerDB.SetMaxIdleConns(maxIdleCons)
	innerDB.SetMaxOpenConns(maxOpenCons)

	for _, table := range tables {
		if !DB.Migrator().HasTable(table) {
			DB.AutoMigrate(table)
		}
	}

}
