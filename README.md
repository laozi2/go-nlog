go-log
================

Logging package for UDP log



Installation
---------------

```bash
$ go get github.com/laozi2/go-nlog
```

Usage
---------------

```go
package main

import (
	"github.com/laozi2/go-nlog"
)

func main() {
	nlog := &log.Logger{
		Level:     log.INFO,
		Formatter: log.NewFormatter("$Time $Level $AppName $PID $FilePos $Msg", true),
		Out:       log.NewUdpWriter("0.0.0.0:5001", "192.168.17.213:5151"),
	}
	defer nlog.Close()

	nlog.Errorf("error = %s", "some error")
	nlog.Infof("info = %s", "some info")
}
```

* NewFormatter 第一个参数为自定义格式，支持变量

  ```
  $AppName:  程序名
  $Level:    日志级别
  $Time:     日期  2006-01-02T15:04:05.000
  $PID:      进程ID
  $FilePos:  代码位置
  $Msg:      日志消息内容
  ```

* NewFormatter 第二个参数可以指定是否在对日志级别显示颜色

* 日志示例

  ```
  2019-05-16T16:52:56.419 ERROR prj_udp.exe 9328 prj_udp/main.go:60 error = some error
  2019-05-16T16:52:56.432 INFO prj_udp.exe 9328 prj_udp/main.go:61 info = some info
  ```