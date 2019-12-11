package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type FeedBack struct {
	gorm.Model
	Name    string
	Contact string
	Content string
}

func (this *FeedBack) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}

func (this *FeedBack) List() (err error) {
	return
}
