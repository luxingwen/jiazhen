package controllers

import (
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"

	"jiazhen/config"
	"jiazhen/models"

	"strconv"
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

func CategoryAdd(c *gin.Context) {
	category := new(models.Category)

	err := c.ShouldBindJSON(&category)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	err = category.Add()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}

func CategoryUpdate(c *gin.Context) {
	idstr := c.Param("id")

	id, _ := strconv.ParseUint(idstr, 10, 64)

	category := new(models.Category)

	category.ID = uint(id)
	err := c.ShouldBindJSON(&category)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	err = category.Add()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}
