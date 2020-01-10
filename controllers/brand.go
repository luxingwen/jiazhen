package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/config"
	"jiazhen/models"

	"path"
	"strconv"
	"strings"
)

// @Summary 轮播图
// @Description Exposes some information about itself
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /wx/brand [get]
func BrandList(c *gin.Context) {
	list, err := models.BrandList()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	for _, item := range list {
		item.Img = config.ServerConf.ServerUrl + "/v1/download/image/origin/" + item.Img
	}

	HandleOk(c, list)
}

// @Summary 轮播图增加
// @Description 轮播图增加
// @Accept json
// @Produce json
// @Param param body models.Brand true "{}"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /brand [post]
func BrandAdd(c *gin.Context) {
	brand := new(models.Brand)

	err := c.ShouldBindJSON(&brand)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	err = brand.Add()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}

// @Summary 轮播图编辑
// @Description 轮播图编辑
// @Accept json
// @Produce json
// @Param param body models.Brand true "{}"
// @Param id path int true "id"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /brand/{id} [put]
func BrandUpdate(c *gin.Context) {

	idstr := c.Param("id")

	id, _ := strconv.ParseUint(idstr, 10, 64)

	brand := new(models.Brand)

	brand.ID = uint(id)

	err := c.ShouldBindJSON(&brand)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	if strings.HasPrefix(brand.Img, "http") || strings.HasPrefix(brand.Img, "https") {
		brand.Img = path.Base(brand.Img)
	}

	err = brand.Update()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}
