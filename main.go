package main

import (
	"github.com/NioDevOps/courier/webserver"
	log "github.com/Sirupsen/logrus"
)

func main(){
	log.SetLevel(log.DebugLevel)
	webserver.Start()
}