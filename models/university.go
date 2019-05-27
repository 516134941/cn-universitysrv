package models

import "github.com/jinzhu/gorm"

// UniversityJSON json数据
type UniversityJSON struct {
	Province string         `json:"province"`
	Schools  []SchoolStruct `json:"schools"`
}

// SchoolStruct 学校结构体
type SchoolStruct struct {
	Name string `json:"name"`
	City string `json:"city"`
}

// StoreSchool 存储学校列表
func StoreSchool(db *gorm.DB, province, city, name string) (err error) {
	tpl := "insert into university(province,city,name)values(?,?,?)"
	if err = db.Exec(tpl, province, city, name).Error; err != nil {
		return
	}
	return
}

// University 大学表
type University struct {
	//ID       int    `json:"id" gorm:"column:id"`
	Province string `json:"province" gorm:"column:province"`
	City     string `json:"city" gorm:"column:city"`
	Name     string `json:"name" gorm:"column:name"`
}

// GetUniversityList 获取大学列表
func GetUniversityList(db *gorm.DB, province, city string) (res []University, err error) {
	db = db.Table("university")
	// 条件查询
	if province != "" {
		db = db.Where("province=?", province)
	}
	if city != "" {
		db = db.Where("city=?", city)
	}
	err = db.Order("id").Find(&res).Error
	return
}

// GetProvinceList 获取省份列表
func GetProvinceList(db *gorm.DB) (res []string, err error) {
	err = db.Table("university").Pluck("distinct(province)", &res).Error
	return
}

// GetCityList 获取城市列表
func GetCityList(db *gorm.DB, province string) (res []string, err error) {
	err = db.Table("university").Where("province=?", province).Pluck("distinct(city)", &res).Error
	return
}
