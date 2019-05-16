package log

import (
	"fmt"
	"io"
	"os"
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

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Panic(msg string, args ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprintf(msg, args...))
	}
	panic(fmt.Sprintf(msg, args...))
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
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

func (l *Logger) Close() {
	l.Out.Close()
}
