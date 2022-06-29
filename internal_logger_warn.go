package log

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) WarnMuted() bool {
	return receiver.mutedWarn
}

func (receiver internalLogger) Warn(a ...interface{}) {
	if receiver.WarnMuted() {
		return
	}
	if nil == receiver.writer {
		return
	}

	s := fmt.Sprint(a...)

	receiver.Warnf("%s", s)
}

func (receiver internalLogger) Warnf(format string, a ...interface{}) {
	if receiver.WarnMuted() {
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
			buffer.WriteString("\x1b[38;222;56;43;6m")
		case "":
			buffer.WriteString("[warn] ")
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
