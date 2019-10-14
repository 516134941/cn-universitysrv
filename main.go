package main

import (
	"cn-universitysrv/config"
	"cn-universitysrv/handles"
	"cn-universitysrv/middlewares"
	"flag"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
)

var (
	tomlFile = flag.String("config", "docs/test.toml", "config file")
)

// init 初始化配置
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	flag.Parse()
	// 解析配置文件
	tomlConfig, err := config.UnmarshalConfig(*tomlFile)
	if err != nil {
		log.Errorf("UnmarshalConfig: err:%v\n", err)
		return
	}
	router := gin.New()
	router.Use(gin.Recovery())
	//设置跨域
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middlewares.Config(tomlConfig))
	router.Use(middlewares.Gorm("universitydb", tomlConfig))
	router.GET("/", handles.Index)                         // 首页
	router.POST("/store/schools", handles.StoreUniversity) // 存储大学信息
	router.GET("/schools", handles.GetUniversityList)      // 获取大学列表
	router.GET("/provinces", handles.GetProvinceList)      // 省份列表
	router.GET("/citys", handles.GetCityList)              // 获取城市列表
	// 启动服务
	log.Debugf("run cn-universitysrv at %v\n", tomlConfig.GetListenAddr())
	router.Run(tomlConfig.GetListenAddr())
}
