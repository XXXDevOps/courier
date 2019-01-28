package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"
)

type RequestLog struct {
	during float64
	body   string
	data   string
}

func V1Protocol() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				c.JSON(http.StatusInternalServerError, V1Error{Error: string(debug.Stack()), Code: -1})
			}
		}()

		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		var requestLog RequestLog
		requestLog.body = string(body)

		st := time.Now()
		c.Next()
		err, exist := c.Get(ErrorField)
		if exist {
			c.JSON(http.StatusBadRequest, V1Error{Error: fmt.Sprint(err), Code: 1})
			return
		}
		if c.Writer.Written() {
			return
		}
		during := time.Since(st)
		data, exist := c.Get(DataField)
		if !exist {
			data = nil
		}

		d, _ := json.Marshal(data)
		requestLog.data = string(d)
		requestLog.during = during.Seconds() * 1000
		log.Info(requestLog)
		c.JSON(http.StatusOK, V1{Data: data, During: during.Seconds() * 1000})
		return
	}
}
