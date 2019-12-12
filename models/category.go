package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Category struct {
	gorm.Model
	SortId int    `json:"sortId"`
	Name   string `json:"name"`
	Img    string `json:"img"`
}

func CategoryList() (r []*Category, err error) {
	r = make([]*Category, 0)
	err = common.GetDB().Find(&r).Error
	return
}

func (this *Category) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}
