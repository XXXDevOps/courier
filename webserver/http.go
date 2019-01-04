package webserver
import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Start(){
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/status",status)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Run(":9091")
}


func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
