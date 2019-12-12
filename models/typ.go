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

func TypList() (r []*Typ, count int, err error) {
	r = make([]*Typ, 0)
	err = common.GetDB().Find(&r).Error
	if err != nil {
		return
	}
	err = common.GetDB().Model(&Typ{}).Count(&count).Error

	if err != nil {
		return
	}

	return
}

func (this *Typ) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}
