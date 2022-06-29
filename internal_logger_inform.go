package log

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) InformMuted() bool {
	return receiver.mutedInform
}

func (receiver internalLogger) Inform(a ...interface{}) {
	if receiver.InformMuted() {
		return
	}
	if nil == receiver.writer {
		return
	}

	s := fmt.Sprint(a...)

	receiver.Informf("%s", s)
}

func (receiver internalLogger) Informf(format string, a ...interface{}) {
	if receiver.InformMuted() {
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
			buffer.WriteString("[inform] ")
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
