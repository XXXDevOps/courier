package main

import (
	"github.com/NioDevOps/courier/cfg"
	"github.com/NioDevOps/courier/media"
	log "github.com/Sirupsen/logrus"
	"github.com/NioDevOps/courier/webserver"
	"os"
)

func main(){
	mcfg := cfg.LoadCfg()

	InitLogger(mcfg.LogCfg)
	RegisterAndInitAllMedia(mcfg.MediaCfg)
	webserver.Start()
}


func InitLogger(c cfg.LogCfgStruct){
	log.SetLevel(cfg.LOGLEVELMAP[c.Level])
	log.SetOutput(os.Stdout)
}

func RegisterAndInitAllMedia(mediaCfg map[string]*cfg.MediaCfgStruct){
	center := media.GetMediaCenter()
	for k,v:=range mediaCfg{
		switch v.Type {
		case media.SMTPTYPE:
			m := media.NewSmtpMedia(v)
			e := center.Register(m)
			if e!=nil{
				log.Warn("register media  failed", k)
				continue
			}
			log.Info("register media ", k)
		}
	}
}