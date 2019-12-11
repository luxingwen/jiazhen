package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/models"
)

type ResShifu struct {
	Sort     int
	Name     string
	Typ      int
	TypName  string
	Phone    string
	Img      string
	Price    string
	Location string // 地址
}

func ShifuList(c *gin.Context) {
	list, err := models.ShifuList()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	typList, err := models.TypList()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	mTyp := make(map[int64]string, 0)
	for _, item := range typList {
		mTyp[int64(item.ID)] = item.Name
	}

	resList := make([]*ResShifu, 0)
	for _, item := range list {
		itemv := &ResShifu{
			Sort:     item.Sort,
			Name:     item.Name,
			Typ:      item.Typ,
			Phone:    item.Phone,
			Img:      item.Img,
			Price:    item.Price,
			Location: item.Location,
		}
		if v, ok := mTyp[int64(item.Typ)]; ok {
			itemv.TypName = v
		}
		resList = append(resList, itemv)
	}

	HandleOk(c, resList)
}
