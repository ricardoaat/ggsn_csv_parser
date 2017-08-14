package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ricardoaat/ggsn_csv_parser/config"
	"github.com/ricardoaat/ggsn_csv_parser/reporter"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	version   string
	buildDate string
)

func logInit() {

	now := time.Now()
	logfile := config.Conf.Path.LogPath + fmt.Sprintf("ggsn_reporter_%s.log", now.Format("20060102T150405"))
	fmt.Println("Loging to " + logfile)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(new(prefixed.TextFormatter))
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: logfile,
		log.InfoLevel:  logfile,
		log.ErrorLevel: logfile,
		log.WarnLevel:  logfile,
		log.PanicLevel: logfile,
	}))

}

func main() {

	err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal("Couldn't load config.toml ", err)
		fmt.Println(err)
	}
	logInit()
	log.Info("--------------Init program--------------")
	log.Info(fmt.Sprintf("Version: %s Build Date: %s", version, buildDate))
	log.Info("Loaded configuration " + fmt.Sprint(config.Conf))
	v := flag.Bool("v", false, "Returns the binary version and built date info")
	d := flag.Int("days", 20, "Number of days in the past to check for CDR files, default 20")
	flag.Parse()
	if !*v {
		reporter.Init(*d)
	}
}
