package handles

import (
	"cn-universitysrv/config"
	"cn-universitysrv/models"
	"cn-universitysrv/utils"
	"github.com/go-xweb/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Index 首页
func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

// StoreUniversity 存储大学json信息至数据库
func StoreUniversity(c *gin.Context) {
	defer utils.LogStat("StoreUniversity", c.Request, time.Now())

	var un []models.UniversityJSON
	config.LoadJSON("docs/ChinaUniversityList.json", &un)
	for _, v := range un {
		province := v.Province
		for _, sh := range v.Schools {
			city := sh.City
			name := sh.Name
			if err := models.StoreSchool(province, city, name); err != nil {
				log.Error("StoreUniversity err:%v\n", err)
				c.JSON(http.StatusOK, gin.H{"errno": "40", "errmsg": "存储出错"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"errno": "0"})
	return
}

// GetUniversityListRequest 获取大学列表请求
type GetUniversityListRequest struct {
	Province string `json:"province" form:"province"`
	City     string `json:"city" form:"city"`
}

// GetUniversityList 获取大学列表
func GetUniversityList(c *gin.Context) {
	defer utils.LogStat("GetUniversityList", c.Request, time.Now())

	// 前端数据捆绑
	var req GetUniversityListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "参数不匹配，请重试"})
		return
	}
	// 获取列表
	list, err := models.GetUniversityList(req.Province, req.City)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("GetUniversityList err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{"errno": "41", "errmsg": "获取列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errno": "0", "data": list})
}

// GetProvinceList 获取省份列表
func GetProvinceList(c *gin.Context) {
	defer utils.LogStat("GetProvinceList", c.Request, time.Now())

	// 获取列表
	list, err := models.GetProvinceList()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("GetUniversityList err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{"errno": "41", "errmsg": "获取列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errno": "0", "data": list})
}

// GetCityListRequest 获取城市列表请求
type GetCityListRequest struct {
	Province string `json:"province" form:"province" binding:"required"`
}

// GetCityList 获取城市列表
func GetCityList(c *gin.Context) {
	defer utils.LogStat("GetCityList", c.Request, time.Now())

	// 前端数据捆绑
	var req GetCityListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "参数不匹配，请重试"})
		return
	}
	// 获取列表
	list, err := models.GetCityList(req.Province)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("GetCityList err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{"errno": "41", "errmsg": "获取列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errno": "0", "data": list})
}
