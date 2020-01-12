package controllers

import (
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"

	"jiazhen/config"
	"jiazhen/models"

	"path"
	"strings"

	"strconv"
)

// @Summary 大类列表
// @Description 大类
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /wx/category [get]
func CategoryList(c *gin.Context) {

	list, count, err := models.CategoryList()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	for _, item := range list {
		item.Img = config.ServerConf.ServerUrl + "/v1/download/image/origin/" + item.Img
	}

	m := make(map[string]interface{}, 0)
	m["count"] = count
	m["list"] = list
	HandleOk(c, m)
}

// @Summary 添加大类
// @Description 添加大类
// @Accept json
// @Produce json
// @Param param body models.Category true "{}"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /category [post]
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

// @Summary 更新大类
// @Description 更新大类
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param param body models.Category true "{}"
// @Success 200 {string} string "{"code":200,"data":"","msg":"ok"}"
// @Router /category/{id} [post]
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

	if strings.HasPrefix(category.Img, "http") || strings.HasPrefix(category.Img, "https") {
		category.Img = path.Base(category.Img)
	}

	err = category.Update()
	if err != nil {
		HandleErr(c, 1, err.Error())
		return
	}

	HandleOk(c, "success")
}
