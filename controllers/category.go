package controllers

import (
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"

	"jiazhen/config"
	"jiazhen/models"
)

func CategoryList(c *gin.Context) {
	list, err := models.CategoryList()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	for _, item := range list {
		item.Img = config.ServerConf.FilePath + item.Img
	}

	HandleOk(c, list)
}
