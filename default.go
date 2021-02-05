package flog

import (
	"os"
)

var Default Logger

func init() {
	Default = NewLogger(os.Stdout, "color")
}

func Debug(a ...interface{}) {
	Default.Debug(a...)
}
func Debugf(format string, a ...interface{}) {
	Default.Debugf(format, a...)
}

func Error(a ...interface{}) error {
	return Default.Error(a...)
}

func Errorf(format string, a ...interface{}) error {
	return Default.Errorf(format, a...)
}

func Fatal(a ...interface{}) {
	Default.Fatal(a...)
}

func Fatalf(format string, a ...interface{}) {
	Default.Fatalf(format, a...)
}

func Highlight(a ...interface{}) {
	Default.Highlight(a...)
}

func Highlightf(format string, a ...interface{}) {
	Default.Highlightf(format, a...)
}

func Inform(a ...interface{}) {
	Default.Inform(a...)
}

func Informf(format string, a ...interface{}) {
	Default.Informf(format, a...)
}

func Panic(a ...interface{}) {
	Default.Panic(a...)
}

func Panicf(format string, a ...interface{}) {
	Default.Panicf(format, a...)
}

func Trace(a ...interface{}) {
	Default.Trace(a...)
}

func Tracef(format string, a ...interface{}) {
	Default.Tracef(format, a...)
}

func Warn(a ...interface{}) {
	Default.Warn(a...)
}

func Warnf(format string, a ...interface{}) {
	Default.Warnf(format, a...)
}

func Prefix(a ...string) Logger {
	return Default.Prefix(a...)
}
