package main

import (
	"fmt"
	"github.com/NioDevOps/courier/cfg"
	"github.com/NioDevOps/courier/media"
	"github.com/NioDevOps/courier/webserver"
	log "github.com/Sirupsen/logrus"
	"os"
)

func main() {
	mcfg := cfg.LoadCfg()

	InitLogger(mcfg.LogCfg)
	RegisterAndInitAllMedia(mcfg.MediaCfg)
	webserver.Start()
}

func InitLogger(c cfg.LogCfgStruct) {
	log.SetLevel(cfg.LOGLEVELMAP[c.Level])
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	logPath := c.LogPath

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default sederr")
	}
}

func RegisterAndInitAllMedia(mediaCfg map[string]*cfg.MediaCfgStruct) {
	center := media.GetMediaCenter()
	for k, v := range mediaCfg {
		switch v.Type {
		case media.SMTPTYPE:
			m := media.NewSmtpMedia(v)
			e := center.Register(m)
			if e != nil {
				log.Warn("register media  failed", k)
				continue
			}
			log.Info("register media ", k)
		case media.SMSTYPE:
			m := media.NewSmsMedia(v)
			e := center.Register(m)
			if e != nil {
				log.Warn("register media  failed", k)
				continue
			}
			log.Info("register media ", k)
		case media.WORKWEIXINTYPE:
			m := media.NewWorkWeixinMedia(v)
			e := center.Register(m)
			if e != nil {
				log.Warn("register media  failed", k)
				continue
			}
			log.Info("register media ", k)
		}
	}
}
