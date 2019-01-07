package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"runtime/debug"
	"fmt"
)

func V1Protocol()gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(){
			if e := recover(); e != nil {
				c.JSON(http.StatusInternalServerError, V1Error{Error:string(debug.Stack()), Code:-1})
			}
		}()
		st:=time.Now()
		c.Next()
		err, exist := c.Get(ErrorField)
		if exist{
			c.JSON(http.StatusBadRequest, V1Error{Error: fmt.Sprint(err), Code:1})
			return
		}
		if c.Writer.Written() {
			return
		}
		during:=time.Since(st)
		data, exist := c.Get(DataField)
		if !exist {
			data = nil
		}
		c.JSON(http.StatusOK, V1{Data: data, During: during.Seconds() * 1000})
		return
	}
}