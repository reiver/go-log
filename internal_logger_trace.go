package log

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) TraceMuted() bool {
	return receiver.mutedTrace
}

func (receiver internalLogger) Trace(a ...interface{}) {
	if receiver.TraceMuted() {
		return
	}
	if nil == receiver.writer {
		return
	}

	s := fmt.Sprint(a...)

	receiver.Tracef("%s", s)
}

func (receiver internalLogger) Tracef(format string, a ...interface{}) {
	if receiver.TraceMuted() {
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
			buffer.WriteString("\x1b[38;2;255;199;6m")
		case "":
			buffer.WriteString("[trace] ")
		}

		buffer.WriteString(receiver.prefix)
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
