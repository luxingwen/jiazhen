package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/config"
	"jiazhen/models"
)

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
