package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type WxUser struct {
	gorm.Model
	NickName  string `gorm:"column:nickname"`
	AvatarUrl string `gorm:"column:avatar_url"`
	Gender    int    `gorm:"column:gender"`
	OpenId    string `gorm:"column:open_id;type:varchar(70);unique_index"`
}

func (this *WxUser) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}

func (this *WxUser) GetByOpenId() (err error) {
	err = common.GetDB().Where("open_id = ?", this.OpenId).First(&this).Error
	return
}
