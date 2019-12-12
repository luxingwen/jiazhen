package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Shifu struct {
	gorm.Model
	Sort     int    `gorm:"column:sort" json:"sort"`
	Name     string `gorm:"column:name" json:"name"`
	Category int    `gorm:"column:category" json:"category"`
	Typ      int    `gorm:"column:typ" json:"typ"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Img      string `gorm:"column:img" json:"img"`
	Price    string `gorm:"column:price" json:"price"`
	Hot      int    `gorm:"column:hot" json:"hot"`
	Location string `gorm:"column:location" json:"location"` // 地址
	Desc     string `gorm:"colum:desc" json:"desc"`
}

type ReqShifu struct {
	Query
	Name string `form:"name"`
}

func ShifuList(q *ReqShifu) (r []*Shifu, err error) {
	limit := 10
	page := 0

	if q.Page != 0 {
		page = 0
	}

	if q.PageNum != 0 {
		limit = q.PageNum
	}
	offset := limit * page

	db := common.GetDB()
	if q.Name != "" {
		db = db.Where("name = ?", q.Name)
	}
	r = make([]*Shifu, 0)
	err = db.Offset(offset).Limit(limit).Find(&r).Error
	return
}

func (this *Shifu) Add() (err error) {
	err = common.GetDB().Create(&this).Error
	return
}
