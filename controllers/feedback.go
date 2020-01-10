package controllers

import (
	"github.com/gin-gonic/gin"

	//"jiazhen/config"
	"jiazhen/models"
)

// @Summary 意见反馈
// @Accept json
// @Produce  json
// @Param param body models.FeedBack true "{}"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /wx/feedback [post]
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
