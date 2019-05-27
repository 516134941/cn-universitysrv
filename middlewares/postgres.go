package middlewares

import (
	"cn-universitysrv/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Postgres sqlx连接
func Postgres(dbName string, tomlConfig *config.Config) gin.HandlerFunc {
	// 读取配置
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}
	// 链接数据库
	db, err := sqlx.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("sqlx.Open: err:%v", err))
	}
	return func(c *gin.Context) {
		c.Set(dbName, db)
		c.Next()
	}
}

// Gorm Gorm
func Gorm(dbName string, tomlConfig *config.Config) gin.HandlerFunc {
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

	return func(c *gin.Context) {
		c.Set(dbName, db)
		c.Next()
	}
}
