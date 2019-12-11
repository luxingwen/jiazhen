package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Category struct {
	gorm.Model
	SortId int
	Name   string
	Img    string
}

func CategoryList() (r []*Category, err error) {
	r = make([]*Category, 0)
	err = common.GetDB().Find(&r).Error
	return
}
