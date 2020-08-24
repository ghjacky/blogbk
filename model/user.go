package model

import (
	"blogbk/common"
	"crypto/sha256"
	"encoding/base64"
	"github.com/jinzhu/gorm"
)

type SUser struct {
	Username string `json:"username" gorm:"varchar(32);primary_key"`
	Name     string `json:"name" gorm:"varchar(32)"`
	Avatar   []byte `json:"avatar" gorm:"type:blob"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Posts    SPosts `json:"posts" gorm:"foreignkey:Author;association_foreignkey:Username"`
}

type SUsers []*SUser

func (u *SUser) AfterFind(db *gorm.DB) {
	db.Model(u).Related(&u.Posts, "Author")
}

func (u *SUser) TableName() string {
	return "users"
}

func (u *SUser) Add() error {
	hs := sha256.New()
	hs.Write([]byte(u.Password))
	u.Password = base64.URLEncoding.EncodeToString(hs.Sum(nil))
	return common.Mysql.Debug().Create(u).Error
}

func (u *SUser) Update() error {
	return common.Mysql.Debug().Save(u).Error
}

func (u *SUser) GetByUsername() error {
	return common.Mysql.Where("username = ?", u.Username).First(u).Error
}

func (u *SUser) Delete() error {
	return common.Mysql.Debug().Delete(u).Error
}

func (us *SUsers) FetchList() error {
	return common.Mysql.Find(us).Error
}

func (u *SUser) Auth() (bool, error) {
	var _u = &SUser{Username: u.Username}
	if err := _u.GetByUsername(); err != nil {
		return false, err
	}
	hs := sha256.New()
	hs.Write([]byte(u.Password))
	if base64.URLEncoding.EncodeToString(hs.Sum(nil)) == _u.Password {
		return true, nil
	}
	return false, nil
}

func (u *SUser) Logout() (bool, error) {
	if v, ok := SessionPool.Load(u.Username); ok {
		if vv, okk := v.(*SSession); okk && vv != nil {
			SessionPool.Delete(u.Username)
		}
	}
	return true, nil
}
