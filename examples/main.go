package main

import (
	"github.com/laozi2/go-log"
	"github.com/laozi2/go-log/writers"
)

func test_log() {
	udp_writer := new(writers.UdpWriter)
	udp_writer.Name = "udp_logger"
	udp_writer.LocalADDR = "0.0.0.0:5001"
	udp_writer.RemoteADDR = "192.168.17.213:5151"
	defer udp_writer.Close()

	nlog := &log.Logger{
		Level: log.INFO,
		//		Formatter: log.NewFormatter("$Time $Level $AppName $PID $FilePos $Msg", true),
		Formatter: log.NewFormatter(" sfsaffa[$Time]sdf$afdaf ", true),
		Out:       udp_writer,
	}

	nlog.Errorf("error = %s", "some error")
	nlog.Infof("info = %s", "some info")
}

func main() {
	test_log()
}
