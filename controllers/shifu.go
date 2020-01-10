package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/config"
	"jiazhen/models"

	"strconv"
)

type ResShifu struct {
	ID       int    `json:"ID"`
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

// @Summary 师傅列表
// @Description 师傅列表
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param name query string true "name"
// @Param category query int true "category"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /wx/shifu [get]
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
			ID:       int(item.ID),
			Sort:     item.Sort,
			Name:     item.Name,
			Typ:      item.Typ,
			Phone:    item.Phone,
			Img:      config.ServerConf.ServerUrl + "/v1/download/image/origin/" + item.Img,
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

// @Summary 添加师傅
// @Description 添加师傅
// @Accept json
// @Produce json
// @Param param body models.Shifu true "{}"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /shifu [post]
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

// @Summary 师傅详细信息
// @Description 师傅详细信息
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /shifu/{id} [get]
func ShifuInfo(c *gin.Context) {
	idstr := c.Param("id")

	id, _ := strconv.ParseInt(idstr, 10, 64)

	shifu := new(models.Shifu)
	r, err := shifu.Get(id)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	r.Img = config.ServerConf.ServerUrl + "/v1/download/image/origin/" + r.Img
	HandleOk(c, r)
}
