package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Exit is equals os.Exit
var Exit = os.Exit

// Logger is represents an active logging object
type Logger struct {
	m         sync.Mutex
	Level     Level
	Formatter Formatter
	Out       io.WriteCloser
}

var gLog = Logger{}

type NlogConf struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Stdout bool   `yaml:"stdout"`
	Laddr  string `yaml:"laddr"`
	Raddr  string `yaml:"raddr"`
	Color  bool   `yaml:"color"`
}

func InitLog(conf *NlogConf) bool {
	level, ok := StrLevelMap[strings.ToUpper(conf.Level)]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error log level %s\n", conf.Level)
		return false
	}
	if conf.Format == "" {
		fmt.Fprintln(os.Stderr, "Empty format")
		return false
	}
	if conf.Stdout {
		gLog = Logger{
			Level:     level,
			Formatter: NewFormatter(conf.Format, conf.Color),
			Out:       os.Stdout,
		}
		return true
	}
	if conf.Laddr == "" || conf.Raddr == "" {
		fmt.Fprintln(os.Stderr, "Empty laddr or raddr")
		return false
	}
	gLog = Logger{
		Level:     level,
		Formatter: NewFormatter(conf.Format, conf.Color),
		Out:       NewUdpWriter(conf.Laddr, conf.Raddr),
	}
	return true
}

func Debug(obj ...interface{}) {
	gLog.Debug(obj...)
}

func Info(obj ...interface{}) {
	gLog.Info(obj...)
}

func Print(obj ...interface{}) {
	gLog.Print(obj...)
}

func Warn(obj ...interface{}) {
	gLog.Warn(obj...)
}

func Error(obj ...interface{}) {
	gLog.Error(obj...)
}

func Panic(obj ...interface{}) {
	gLog.Panic(obj...)
}

func Fatal(obj ...interface{}) {
	gLog.Fatal(obj...)
}

func Debugf(msg string, args ...interface{}) {
	gLog.Debugf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	gLog.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	gLog.Warnf(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	gLog.Errorf(msg, args...)
}

func Panicf(msg string, args ...interface{}) {
	gLog.Panicf(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	gLog.Fatalf(msg, args...)
}

// Debug outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Debug(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprint(obj...))
	}
}

// Info outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Info(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

// Print outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Print(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

// Warn outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Warn(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprint(obj...))
	}
}

// Error outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Error(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprint(obj...))
	}
}

// Panic outputs message, and followed by a call to panic() Arguments are handled by fmt.Sprint
func (l *Logger) Panic(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprint(obj...))
	}
	panic(fmt.Sprint(obj...))
}

// Fatal outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprint
func (l *Logger) Fatal(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprint(obj...))
	}
	Exit(1)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Panicf(msg string, args ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprintf(msg, args...))
	}
	panic(fmt.Sprintf(msg, args...))
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprintf(msg, args...))
	}
	Exit(1)
}

func (l *Logger) log(level Level, msg string) {
	line := l.Formatter.Format(level, msg)

	l.m.Lock()
	defer l.m.Unlock()

	_, err := l.Out.Write(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write log, %v\n", err)
	}
}

func Close() {
	gLog.Close()
}

func (l *Logger) Close() {
	l.Out.Close()
}
