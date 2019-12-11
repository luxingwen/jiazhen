package models

import (
	"github.com/jinzhu/gorm"

	"jiazhen/common"
)

type Shifu struct {
	gorm.Model
	Sort     int    `gorm:"column:sort"`
	Name     string `gorm:"column:name"`
	Typ      int    `gorm:"column:typ"`
	Phone    string `gorm:"column:phone"`
	Img      string `gorm:"column:img"`
	Price    string `gorm:"column:price"`
	Hot      int    `gorm:"column:hot"`
	Location string `gorm:"column:location"` // 地址
}

func ShifuList() (r []*Shifu, err error) {
	r = make([]*Shifu, 0)
	err = common.GetDB().Find(&r).Error
	return
}
