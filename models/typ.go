package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Typ struct {
	gorm.Model
	CategoryId int64
	Name       string
}

func TypList() (r []*Typ, err error) {
	r = make([]*Typ, 0)
	err = common.GetDB().Find(&r).Error
	return
}
