package webserver

import (
	"github.com/gin-gonic/gin"
)


func send(c *gin.Context) {
	var p sendParams
	err := c.BindJSON(&p)
	if err!=nil{
		c.Set(ErrorField, err.Error())
		return
	}
	err = SendService(p.Media, p.Sender, p.Receivers, p.Message, p.Cc)
	if err!=nil{
		c.Set(ErrorField, err.Error())
		return
	}
	c.Set(DataField, p)
}