package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Typ struct {
	gorm.Model
	CategoryId int64  `json:"categoryId"`
	Name       string `json:"name"`
}

func TypList() (r []*Typ, err error) {
	r = make([]*Typ, 0)
	err = common.GetDB().Find(&r).Error
	return
}

func (this *Typ) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}
