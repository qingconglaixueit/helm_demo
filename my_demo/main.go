package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/query", queryInfo)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/query", queryInfo2)
	}

	router.Run(":8989")

}

func queryInfo(c *gin.Context){
	c.JSON(200,gin.H{
		"code":0,
		"msg":"success",
		"data":map[string]interface{}{
			"info": "hello world ,v1.queryInfo",
		},
	})
}

func queryInfo2(c *gin.Context){
	c.JSON(200,gin.H{
		"code":0,
		"msg":"success",
		"data":map[string]interface{}{
			"info": "hello world ,v2.queryInfo",
		},
	})
}

func getFIle(c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "readme.txt"))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./readme.txt")
}