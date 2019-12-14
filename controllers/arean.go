package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/models"
)

// @Summary 地区
// @Description 地区列表
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /wx/arean [get]
func AreanList(c *gin.Context) {
	list, err := models.AreaList()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, list)
}
