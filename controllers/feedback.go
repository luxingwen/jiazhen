package controllers

import (
	"github.com/gin-gonic/gin"

	//"jiazhen/config"
	"jiazhen/models"
)

func FeedBackAdd(c *gin.Context) {
	fb := new(models.FeedBack)

	err := c.ShouldBindJSON(&fb)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	err = fb.Add()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}
