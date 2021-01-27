package router

import (
	"chatdemo/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() {

	//新建服务
	router := gin.Default()

	//设置模板分隔符,以免和elemtnUi分隔符冲突
	//这个设置要在加载模板设置前面
	router.Delims("{{{", "}}}")

	//设置静态目录
	router.Static("/static", "static")

	//设置模板位置
	router.LoadHTMLGlob("template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/sendMessage", app.SendMessageApi)

	// 启动服务
	router.Run(":9999")
}
