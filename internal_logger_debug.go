package flog

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) CanLogDebug() bool {
	return !receiver.canNotLogDebug
}

func (receiver internalLogger) Debug(a ...interface{}) {
	if !receiver.CanLogDebug() {
		return
	}
	if nil == receiver.writer {
		return
	}

	s := fmt.Sprint(a...)

	receiver.Debugf("%s", s)
}

func (receiver internalLogger) Debugf(format string, a ...interface{}) {
	if !receiver.CanLogDebug() {
		return
	}

	var writer io.Writer = receiver.writer
	if nil == writer {
		return
	}

	var newformat string
	{
		var buffer strings.Builder

		switch receiver.style{
		case"color":
			buffer.WriteString("\x1b[48;2;1;1;1m")
			buffer.WriteString("\x1b[38;2;44;181;233m")
		case "":
			buffer.WriteString("[debug] ")
		}

		buffer.WriteString(format)

		switch receiver.style {
		case "color":
			buffer.WriteString("\x1b[0m")
			buffer.WriteRune('\n')
		case "":
			buffer.WriteRune('\n')
		}

		newformat = buffer.String()
	}

	fmt.Fprintf(receiver.writer, newformat, a...)
}
