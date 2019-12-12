package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/models"
)

type ResShifu struct {
	Sort     int    `json:"sort"`
	Name     string `json:"name"`
	Typ      int    `json:"typ"`
	TypName  string `json:"typName"`
	Phone    string `json:"phone"`
	Img      string `json:"img"`
	Price    string `json:"price"`
	Location string `json:"location"` // 地址
	Desc     string `json:"desc"`
}

func ShifuList(c *gin.Context) {
	req := new(models.ReqShifu)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	list, count, err := models.ShifuList(req)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	m := make(map[string]interface{}, 0)

	m["count"] = count

	typList, _, err := models.TypList()
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
			Desc:     item.Desc,
		}
		if v, ok := mTyp[int64(item.Typ)]; ok {
			itemv.TypName = v
		}
		resList = append(resList, itemv)
	}

	m["list"] = resList

	HandleOk(c, m)
}

func ShifuAdd(c *gin.Context) {
	shifu := new(models.Shifu)

	err := c.ShouldBindJSON(&shifu)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	err = shifu.Add()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}
