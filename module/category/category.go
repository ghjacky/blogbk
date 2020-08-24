package category

import (
	"blogbk/model"
	"github.com/jinzhu/gorm"
)

func Add(cat *model.SCategory) error {
	return cat.Add()
}

func Update(cat *model.SCategory) error {
	var c = new(model.SCategory)
	c.ID = cat.ID
	if err := c.Get(); gorm.IsRecordNotFoundError(err) {
		return err
	}
	return cat.Update()
}

func Fetch(cats *model.SCategories, tree int) error {
	if tree == 0 {
		return cats.FetchList()
	} else {
		return cats.FetchTree()
	}
}

func Delete(cat *model.SCategory) error {
	return cat.Delete()
}

func Get(cat *model.SCategory) error {
	return cat.Get()
}
