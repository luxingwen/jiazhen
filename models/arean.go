package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Arean struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
}

func AreaList() (r []*Arean, err error) {
	r = make([]*Arean, 0)
	err = common.GetDB().Find(&r).Error
	return
}
