package main

import (
	"fmt"
	"os"

	"prj_prod_cons/conf"

	"github.com/laozi2/go-nlog"
)

var (
	gconf conf.Config
)

func init() {
	confFile := "../conf/config.yaml"
	if err := conf.ParseConf(&gconf, confFile); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {
	ok := log.InitLog(gconf.Nlog)
	if !ok {
		os.Exit(1)
	}
	defer log.Close()

	log.Errorf("error = %s", "some error")
	log.Infof("info = %s", "some info")

	fmt.Println("main done")
}
