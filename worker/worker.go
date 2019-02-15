package worker

import (
	"bytes"
	"encoding/json"
	"github.com/NioDevOps/courier/cfg"
	"github.com/NioDevOps/courier/models"
	"github.com/NioDevOps/courier/webserver"
	log "github.com/Sirupsen/logrus"
	"time"
)

func send_worker(jobs <-chan webserver.SendParams) {
	for true {
		j := <-jobs
		err := webserver.SendService(j.Media, j.Sender, j.Receivers, j.Message, j.Cc)
		if err != nil {
			log.Error(err)
		}
	}
}

func Init(c cfg.WorkerCfgStruct) {
	var p webserver.SendParams
	jobs := make(chan webserver.SendParams, c.ChannelSize)

	for w := 1; w <= c.PoolSize; w++ {
		go send_worker(jobs)
	}

	go func() {
		for true {
			redis_client := models.RedisPool.Get()
			reply, _ := redis_client.Do("LPOP", "courier_msg_queue")
			if reply == nil {
				time.Sleep(300 * time.Millisecond)
				continue
			}
			decoder := json.NewDecoder(bytes.NewReader(reply.([]byte)))
			if err := decoder.Decode(&p); err != nil {
				log.Error(err)
				continue
			}
			jobs <- p
		}
	}()
}
