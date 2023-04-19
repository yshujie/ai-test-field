package main

import (
	"github.com/gin-gonic/gin"
	familiarisation "github.com/yshujie/family-doctor-gpt"
)

func main() {

	// 创建默认的路由引擎
	r := gin.Default()

	// 设置 GET 路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": familiarisation.Hello()})
	})

	// 启动 HTTP 服务
	r.Run()
}
