package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/log_source"

	//_ "shop/data_source"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK,"success")
}

func Log(c *gin.Context)  {
	log_source.Logger.Info("这是测试日志")
	c.JSON(http.StatusOK,"success")
}


func main() {
	r := gin.Default()
	r.GET("/",Index)
	r.GET("/log",Log)
	_ = r.Run(":9000")
}
