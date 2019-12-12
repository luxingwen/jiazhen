package controllers

import (
	"github.com/gin-gonic/gin"

	"jiazhen/common"
	"jiazhen/models"
)

// @Summary 登录
// @Accept json
// @Produce  json
// @Param param body models.User true "{}"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /user/login [post]
func Login(c *gin.Context) {

	user := new(models.User)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	token, err := common.GenerateToken(user.Username, user.Password)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	m := make(map[string]interface{}, 0)
	m["token"] = token
	HandleOk(c, m)
}

func UserInfo(c *gin.Context) {
	m := make(map[string]interface{}, 0)
	m["name"] = "admin"
	m["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"

	HandleOk(c, m)
}
