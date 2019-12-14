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

func CategoryList() (r []*Category, count int, err error) {
	r = make([]*Category, 0)
	err = common.GetDB().Find(&r).Error
	if err != nil {
		return
	}

	err = common.GetDB().Model(&Category{}).Count(&count).Error
	return
}

func (this *Category) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}

func (this *Category) Update() (err error) {
	err = common.GetDB().Save(&this).Error
	return
}
