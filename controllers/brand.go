package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/config"
	"jiazhen/models"
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
		item.Img = config.ServerConf.FilePath + item.Img
	}

	HandleOk(c, list)
}
