package log

import (
	"errors"
	"fmt"
	"os"
)

type DiscardLogger struct {}

var _ Logger = DiscardLogger{}

var Discard DiscardLogger

func (receiver DiscardLogger) Prefix(...string) Logger {
	return receiver
}

func (receiver DiscardLogger) Alert(a...any) error {
	return errors.New(fmt.Sprint(a...))
}
func (receiver DiscardLogger) Alertf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func (receiver DiscardLogger) Begin(...any) Logger {
	return receiver
}

func (receiver DiscardLogger) Debug(...any) {}
func (receiver DiscardLogger) Debugf(string, ...any) {}

func (receiver DiscardLogger) End(...any) {}

func (receiver DiscardLogger) Error(a ...any) error {
	return errors.New(fmt.Sprint(a...))
}
func (receiver DiscardLogger) Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func (receiver DiscardLogger) Fatal(...any) {
	os.Exit(1)
}
func (receiver DiscardLogger) Fatalf(string, ...any) {
	os.Exit(1)
}

func (receiver DiscardLogger) Highlight(...any) {}
func (receiver DiscardLogger) Highlightf(string, ...any) {}

func (receiver DiscardLogger) Inform(...any) {}
func (receiver DiscardLogger) Informf(string, ...any) {}

func (receiver DiscardLogger) Panic(a ...any) {
	panic(receiver.Error(a...))
}
func (receiver DiscardLogger) Panicf(format string, a ...any) {
	panic(receiver.Errorf(format, a...))
}

func (receiver DiscardLogger) Trace(...any) {}
func (receiver DiscardLogger) Tracef(string, ...any) {}

func (receiver DiscardLogger) Warn(...any) {}
func (receiver DiscardLogger) Warnf(string, ...any) {}
