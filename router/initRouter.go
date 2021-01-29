package router

import (
	"github.com/gin-gonic/gin"
)


func CreateApp(port string) {
	r := gin.Default() // 使用默认中间件（logger和recovery）
	r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{ // 返回一个JSON，状态码是200，gin.H是map[string]interface{}的简写
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:3005") // 启动服务，并默认监听8080端口
}