package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/medivhzhan/weapp/v2"

	"jiazhen/common"
	"jiazhen/config"
	"jiazhen/models"

	"fmt"
	//"strings"
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

type Wxlogin struct {
	Code      string `json:"code"`
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
}

// @Summary 登录
// @Accept json
// @Produce  json
// @Param param body models.User true "{}"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /wx/login [post]
func WxLogin(c *gin.Context) {
	req := new(Wxlogin)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	fmt.Printf("req: %#v\n", req)
	fmt.Println("code:", req.Code)
	res, err := weapp.Login(config.WxConf.AppId, config.WxConf.AppSecret, req.Code)
	if err != nil {
		// 处理一般错误信息
		HandleErr(c, 1, err.Error())
		return
	}

	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		HandleErr(c, 1, err.Error())
		return
	}

	fmt.Printf("返回结果: %#v\n", res)
	wxUser := &models.WxUser{NickName: req.NickName, AvatarUrl: req.AvatarUrl, Gender: req.Gender, OpenId: res.OpenID}

	err = wxUser.GetByOpenId()
	if err != nil && err.Error() == "record not found" {
		wxUser.Add()
	}

	fmt.Printf("wxUser: %d \n", wxUser.ID)

	token, err := common.GenerateWxToken(wxUser.NickName, wxUser.OpenId, int(wxUser.ID))
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}
	mdata := make(map[string]string, 0)
	mdata["token"] = token

	HandleOk(c, mdata)
}
