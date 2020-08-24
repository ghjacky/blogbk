package model

import (
	"blogbk/common"
	"github.com/jinzhu/gorm"
)

type SCategory struct {
	BaseModel
	Name     string      `json:"name" gorm:"unique"`
	Parent   uint        `json:"parent"`
	Children SCategories `json:"children" gorm:"foreignkey:Parent;association_foreignkey:ID"`
	Posts    SPosts      `json:"posts" gorm:"foreignkey:Category;association_foreignkey:Name"`
}

type SCategories []*SCategory

func (c *SCategory) TableName() string {
	return "categories"
}

func (c *SCategory) Add() error {
	return common.Mysql.Debug().Create(c).Error
}

func (c *SCategory) Delete() error {
	return common.Mysql.Debug().Delete(c).Error
}

func (c *SCategory) Update() error {
	return common.Mysql.Debug().Save(c).Error
}

func (c *SCategory) Get() error {
	return common.Mysql.Where("id = ?", c.ID).Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}).First(c).Error
}

func (cs *SCategories) FetchTree() error {
	return common.Mysql.Where("parent = 0").Preload("Children").Find(cs).Error
}

func (cs *SCategories) FetchList() error {
	return common.Mysql.Raw("select * from categories where deleted_at is NULL").Scan(cs).Error
}
