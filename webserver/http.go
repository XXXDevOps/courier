package webserver

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	//r := gin.Default()
	r := gin.New()
	v1 := r.Group("/v1", V1Protocol())
	{
		v1.GET("/status", status)
		v1.POST("/send", send)
	}
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Run(":9091")
}

func status(c *gin.Context) {
	c.Set(DataField, true)
}
