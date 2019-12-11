package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Brand struct {
	gorm.Model
	Url  string
	Sort int
	Img  string
}

func BrandList() (r []*Brand, err error) {
	r = make([]*Brand, 0)
	err = common.GetDB().Find(&r).Error
	return
}
