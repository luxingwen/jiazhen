package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"jiazhen/common"
	"jiazhen/controllers"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")

	r.Use(Cors())

	wx := r.Group("/wx")
	{
		wx.GET("/category", controllers.CategoryList)
		wx.GET("/shifu", controllers.ShifuList)
		wx.GET("/brand", controllers.BrandList)
		wx.POST("/feedback", controllers.FeedBackAdd)
	}

	r.POST("/user/login", controllers.Login)

	r.Use(common.JWT())

	r.GET("/user/info", controllers.UserInfo)

	r.GET("/typ", controllers.TypList)

	// api:=r.Group("/api")

	// {
	// 	api.POST("/login")
	// }

	return r
}

////// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		fmt.Println("origin:", origin)
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "AccessToken, Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                   //	跨域请求是否需要带cookie信息 默认设置为true
			c.Header("content-type", "application/json")
			c.Set("content-type", "application/json") // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
			// c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next() //	处理请求
	}
}