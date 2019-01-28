package cfg

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/NioDevOps/courier/version"
	log "github.com/Sirupsen/logrus"
	"os"
)

var LOGLEVELMAP = map[string]log.Level{"debug": log.DebugLevel, "info": log.InfoLevel, "warning": log.WarnLevel, "error": log.ErrorLevel, "fatal": log.FatalLevel, "panic": log.PanicLevel}

type MainCfg struct {
	LogCfg   LogCfgStruct               `toml:"log"`
	MediaCfg map[string]*MediaCfgStruct `toml:"Media"`
	Daemon   bool
	Debug    bool
}

//log config struct for toml
type LogCfgStruct struct {
	Level   string    `toml:"level"`
	LevelId log.Level `toml:"levelId"`
	LogPath string    `toml:"path"`
}

//Load all config
func LoadCfg() *MainCfg {
	//debug:=flag.Bool("debug", false, "Start in debug mode.")
	//daemon:=flag.Bool("daemon", false, "Start in daemon mode.")
	cfgfile := flag.String("config", version.APPNAME+".toml", "Configuration file ")
	help := flag.Bool("help", false, "Show all the help infomation")
	sv := flag.Bool("version", false, "Show version")
	flag.Parse()
	if *help {
		fmt.Println("====================================")
		fmt.Println("==============" + version.APPNAME + "==============")
		fmt.Println("====================================")
		fmt.Println("Usage:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *sv {
		fmt.Printf("Version:%f \n", version.VERSION)
		os.Exit(0)
	}
	var mainCfgObj *MainCfg
	meta, err := toml.DecodeFile(*cfgfile, &mainCfgObj)
	if err != nil {
		fmt.Printf("Configuration Error:%s\n", err.Error())
		os.Exit(-1)
	}
	//config  media
	for k, v := range mainCfgObj.MediaCfg {
		v.Name = k
	}
	//config log
	if !meta.IsDefined("log", "level") {
		mainCfgObj.LogCfg.Level = "info"
	}
	levelId, err := log.ParseLevel(mainCfgObj.LogCfg.Level)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		os.Exit(-1)
	} else {
		mainCfgObj.LogCfg.LevelId = levelId
	}
	return mainCfgObj
}
