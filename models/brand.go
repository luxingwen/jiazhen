package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Brand struct {
	gorm.Model
	Url  string `gorm:"column:url" json:"url"`
	Sort int    `gorm:"column:sort" json:"sort"`
	Img  string `gorm:"column:img" json:"img"`
}

func BrandList() (r []*Brand, err error) {
	r = make([]*Brand, 0)
	err = common.GetDB().Find(&r).Error
	return
}

func (this *Brand) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}

func (this *Brand) Update() (err error) {
	err = common.GetDB().Save(&this).Error
	return
}
