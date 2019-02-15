package webserver

import (
	"encoding/json"
	"github.com/NioDevOps/courier/models"
	"github.com/gin-gonic/gin"
)

func send(c *gin.Context) {
	var p SendParams
	err := c.BindJSON(&p)
	if err != nil {
		c.Set(ErrorField, err.Error())
		return
	}

	err = SendService(p.Media, p.Sender, p.Receivers, p.Message, p.Cc)
	if err != nil {
		c.Set(ErrorField, err.Error())
		return
	}
	c.Set(DataField, p)
}

func send_async(c *gin.Context) {
	var p SendParams
	err := c.BindJSON(&p)
	if err != nil {
		c.Set(ErrorField, err.Error())
		return
	}

	b, err := json.Marshal(p)
	if err != nil {
		c.Set(ErrorField, err.Error())
		return
	}

	redis_client := models.RedisPool.Get()
	_, err = redis_client.Do("rpush", "courier_msg_queue", b)
	if err != nil {
		c.Set(ErrorField, err.Error())
		return
	}
}
