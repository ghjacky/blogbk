package model

import "blogbk/common"

type STag struct {
	Name  string `json:"name" gorm:"primary_key"`
	Posts SPosts `json:"posts" gorm:"many2many:posts_tags"`
}

type STags []*STag

func (t *STag) TableName() string {
	return "tags"
}

func (ts *STags) FetchList() error {
	return common.Mysql.Find(ts).Error
}

func (t *STag) Get() error {
	return common.Mysql.Where("name = ?", t.Name).Preload("Posts").First(t).Error
}
