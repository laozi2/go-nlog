package main

import (
	"github.com/laozi2/go-log"
)

func test_log() {
	nlog := &log.Logger{
		Level:     log.INFO,
		Formatter: log.NewFormatter("$Time $Level $AppName $PID $FilePos $Msg", true),
		Out:       log.NewUdpWriter("0.0.0.0:5001", "192.168.17.213:5151"),
	}
	defer nlog.Close()

	nlog.Errorf("error = %s", "some error")
	nlog.Infof("info = %s", "some info")
}

func main() {
	test_log()
}
