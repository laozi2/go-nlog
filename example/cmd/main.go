package main

import (
	"fmt"
	"os"

	"prj_prod_cons/conf"

	"github.com/laozi2/go-nlog"
)

var (
	gconf conf.Config
	nlog  *log.Logger
)

func init() {
	confFile := "../conf/config.yaml"
	if err := conf.ParseConf(&gconf, confFile); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {
	nlog = log.NewLog(gconf.Nlog)
	if nlog == nil {
		os.Exit(1)
	}
	defer nlog.Close()

	nlog.Errorf("error = %s", "some error")
	nlog.Infof("info = %s", "some info")

	fmt.Println("main done")
}
