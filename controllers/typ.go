package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/models"
)

func TypList(c *gin.Context) {
	list, err := models.TypList()
	if err != nil {
		HandleErr(c, 0, err.Error())
		return
	}
	HandleOk(c, list)
}
