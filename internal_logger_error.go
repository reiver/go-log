package flog

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) CanLogError() bool {
	return !receiver.mutedError
}

func (receiver internalLogger) Error(a ...interface{}) error {
	s := fmt.Sprint(a...)

	return receiver.Errorf("%s", s)
}

func (receiver internalLogger) Errorf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)

	if !receiver.CanLogError() {
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
			buffer.WriteString("\x1b[48;2;1;1;1m")
			buffer.WriteString("\x1b[38;222;56;43;6m")
		case "":
			buffer.WriteString("[ERROR] ")
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

	return err
}
