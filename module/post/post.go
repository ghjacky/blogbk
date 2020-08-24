package post

import (
	"blogbk/model"
	"github.com/jinzhu/gorm"
)

func Fetch(ps *model.SPosts, query model.SQuery) error {
	return ps.Fetch(query)
}

func Add(p *model.SPost) error {
	return p.Add()
}

func Update(p *model.SPost) error {
	var _p = new(model.SPost)
	_p.ID = p.ID
	if err := _p.Get(); err != nil && gorm.IsRecordNotFoundError(err) {
		return err
	}
	return p.Update()
}

func Get(p *model.SPost) error {
	return p.Get()
}

func Delete(p *model.SPost) error {
	return p.Delete()
}
