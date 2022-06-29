package log

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) AlertMuted() bool {
	return receiver.mutedAlert
}

func (receiver internalLogger) Alert(a ...interface{}) error {
	s := fmt.Sprint(a...)

	return receiver.Alertf("%s", s)
}

func (receiver internalLogger) Alertf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)

	if receiver.AlertMuted() {
		return err
	}

	var writer io.Writer = receiver.writer
	if nil == writer {
		return err
	}

	var newformat string
	{
		var buffer strings.Builder

		switch receiver.style{
		case"color":
			buffer.WriteString("☣️☣️☣️☣️☣️ ")
			buffer.WriteString("\x1b[48;2;0;43;54m")
			buffer.WriteString("\x1b[38;2;220;50;47m")
		case "":
			buffer.WriteString("[ALERT] ")
		}

		buffer.WriteString(receiver.prefix)
		buffer.WriteString(format)

		switch receiver.style {
		case "color":
			buffer.WriteString("\x1b[0m")
			buffer.WriteString(" ☣️☣️☣️☣️☣️")
			buffer.WriteRune('\n')
		case "":
			buffer.WriteRune('\n')
		}

		newformat = buffer.String()
	}

	fmt.Fprintf(receiver.writer, newformat, a...)

	return err
}
