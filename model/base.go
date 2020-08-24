package model

import (
	"time"
)

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

type SQuery struct {
	Field string `json:"field" form:"field"`
	Text  string `json:"text" form:"text"`
	Page  uint   `json:"page" form:"page"`
	Limit uint   `json:"limit" form:"limit"`
	Order string `json:"order" form:"order"`
}
