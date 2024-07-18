package gonetworking

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const RESET = "\x1b[0m"
const FAINT_WHITE = "\x1b[37;2m"
const BRIGHT_CYAN = "\x1b[96;22m"
const BRIGHT_YELLOW = "\x1b[93;22m"
const BRIGHT_RED = "\x1b[91;22m"

type Logger struct {
	Name string
}

func (l *Logger) log(color string, s string, args ...any) {
	now := time.Now()
	hour, min, sec := now.Clock()
	ms := now.Nanosecond() / 1_000_000
	var b strings.Builder
	b.Grow(128)
	fmt.Fprintf(&b, "%s%02d:%02d:%02d.%03d ", FAINT_WHITE, hour, min, sec, ms)
	if l.Name != "" {
		fmt.Fprintf(&b, "[%s] ", l.Name)
	}
	b.WriteString(color)
	fmt.Fprintf(&b, s, args...)
	b.WriteString(RESET)
	fmt.Println(b.String())
}
func (l *Logger) Debug(s string, args ...any) {
	l.log(RESET, s, args...)
}
func (l *Logger) Info(s string, args ...any) {
	l.log(BRIGHT_CYAN, s, args...)
}
func (l *Logger) Warn(s string, args ...any) {
	l.log(BRIGHT_YELLOW, s, args...)
}
func (l *Logger) Error(s string, args ...any) {
	l.log(BRIGHT_RED, s, args...)
}
func (l *Logger) CheckFatal(err error) {
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
}
