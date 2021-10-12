package middlewares

import (
	"cn-universitysrv/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

// PgConnect 数据库连接
func PgConnect(dbName string, tomlConfig *config.Config){
	// 读取配置
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	db, err := gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}

	// 设置最大链接数
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)

	Db = db
}
