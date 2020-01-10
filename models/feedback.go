package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type FeedBack struct {
	gorm.Model
	Name    string // 昵称
	Contact string // 联系方式
	Content string // 反馈内容
}

func (this *FeedBack) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}

func (this *FeedBack) List() (err error) {
	return
}
