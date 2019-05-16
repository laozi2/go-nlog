package log

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	fmtBuffer = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

type Formatter struct {
	fmtString   string
	formatItems []string
	appName     string
	timeFormat  string
	colored     bool

	app []byte
	pid []byte
}

func NewFormatter(formatString string, colored bool) Formatter {
	var fmtr Formatter
	fmtr.fmtString = formatString
	fmtr.colored = colored

	re := regexp.MustCompile(`\$\w+`)
	ret := re.FindAllStringIndex(formatString, -1)
	if len(ret) == 0 {
		fmtr.formatItems = append(fmtr.formatItems, formatString)
		fmt.Printf("WARN: got no any args [%s]\n", formatString)
	} else {
		prepos := 0
		for _, group := range ret {
			preRawStr := string(formatString[prepos:group[0]])
			fmtr.formatItems = append(fmtr.formatItems, preRawStr)

			matchStr := string(formatString[group[0]:group[1]])
			fmtr.formatItems = append(fmtr.formatItems, matchStr)
			prepos = group[1]
		}
		fmtr.formatItems = append(fmtr.formatItems, string(formatString[prepos:]))
	}

	fmtr.appName = filepath.Base(os.Args[0])
	fmtr.app = []byte(fmtr.appName)
	fmtr.timeFormat = "2006-01-02T15:04:05.000"
	fmtr.pid = []byte(strconv.Itoa(os.Getpid()))

	return fmtr
}

// Format implements log.Formatter
func (f *Formatter) Format(level Level, msg string) []byte {
	buf := fmtBuffer.Get().(*bytes.Buffer)
	buf.Reset()
	defer fmtBuffer.Put(buf)

	for _, formatItem := range f.formatItems {
		switch formatItem {
		case "$AppName":
			formatterAppName(buf, f)
		case "$Level":
			formatterLevel(buf, level, f)
		case "$Time":
			formatterTime(buf, f)
		case "$PID":
			formatterPID(buf, f)
		case "$FilePos":
			formatterFilePos(buf)
		case "$Msg":
			formatterMsg(buf, msg)
		default:
			formatterMsg(buf, formatItem)
		}
	}
	buf.WriteByte('\n')

	return buf.Bytes()
}

func formatterAppName(buf *bytes.Buffer, f *Formatter) {
	buf.Write(f.app)
}

func formatterLevel(buf *bytes.Buffer, level Level, f *Formatter) {
	if f.colored {
		buf.WriteString(level.ColorString())
	} else {
		buf.WriteString(level.String())
	}
}

func formatterTime(buf *bytes.Buffer, f *Formatter) {
	timeStr := time.Now().Format(f.timeFormat)
	buf.WriteString(timeStr)
}

func formatterPID(buf *bytes.Buffer, f *Formatter) {
	buf.Write(f.pid)
}

func formatterFilePos(buf *bytes.Buffer) {
	file, line := FilelineCaller(5)
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
}

func formatterMsg(buf *bytes.Buffer, msg string) {
	buf.WriteString(msg)
}

// FilelineCaller returns file and line for caller
func FilelineCaller(skip int) (file string, line int) {
	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(skip + i)
		if !ok {
			return "???", 0
		}

		// file = pkg/file.go
		n := 0
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				n++
				if n >= 2 {
					file = file[i+1:]
					break
				}
			}
		}

		if !strings.HasPrefix(file, "go-log/") {
			return file, line
		}
	}

	return "???", 0
}
