package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/models"
)

func TypList(c *gin.Context) {
	list, count, err := models.TypList()
	if err != nil {
		HandleErr(c, 0, err.Error())
		return
	}

	m := make(map[string]interface{}, 0)

	m["count"] = count
	m["list"] = list
	HandleOk(c, m)
}

func TypAdd(c *gin.Context) {
	typ := new(models.Typ)

	err := c.ShouldBindJSON(&typ)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	err = typ.Add()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}
