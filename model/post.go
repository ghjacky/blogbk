package model

import (
	"blogbk/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type SPost struct {
	BaseModel
	Title    string `json:"title" gorm:"varchar(64)"`
	Author   string `json:"author" gorm:"varchar(32)"`
	Content  string `json:"content" gorm:"type:blob"`
	Summary  string `json:"summary"`
	Category string `json:"category"`
	Tags     STags  `json:"tags" gorm:"many2many:posts_tags"`
}

type SPosts []*SPost

func (p *SPost) TableName() string {
	return "posts"
}

func (p *SPost) AfterSave(db *gorm.DB) error {
	return db.Model(p).Association("Tags").Replace(p.Tags).Error
}

func (p *SPost) BeforeDelete(db *gorm.DB) error {
	return db.Model(p).Association("Tags").Clear().Error
}

func (p *SPost) Add() error {
	return common.Mysql.Debug().Create(p).Error
}

func (p *SPost) Update() error {
	return common.Mysql.Debug().Save(p).Error
}

func (p *SPost) Delete() error {
	return common.Mysql.Debug().Delete(p).Error
}

func (p *SPost) Get() error {
	return common.Mysql.Where("id = ?", p.ID).Preload("Tags").First(p).Error
}

func (ps *SPosts) Fetch(query SQuery) error {
	var where []string
	var order = "created_at desc"
	if len(query.Field) != 0 && len(query.Text) != 0 {
		where = append(where, fmt.Sprintf("%s like '%%%s%%'", query.Field, query.Text))
	}
	if len(query.Order) != 0 && strings.HasPrefix(query.Order, "-") {
		order = strings.TrimPrefix(query.Order, "-") + " desc"
	} else if len(query.Order) != 0 && strings.HasPrefix(query.Order, "+") {
		order = strings.TrimPrefix(query.Order, "+") + " asc"
	}
	return common.Mysql.Model(&SPost{}).Where("deleted_at is NULL").Where(strings.Join(where, " and ")).Order(order).Offset((query.Page - 1) * query.Limit).Limit(query.Limit).Select("id, created_at, title, summary, author, category").Scan(ps).Error
}
