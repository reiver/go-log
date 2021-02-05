package flog

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) CanLogPanic() bool {
	return !receiver.canNotLogPanic
}

func (receiver internalLogger) Panic(a ...interface{}) {
	s := fmt.Sprint(a...)

	receiver.Panicf("%s", s)
}

func (receiver internalLogger) Panicf(format string, a ...interface{}) {
	err := fmt.Errorf(format, a...)

	if !receiver.CanLogPanic() {
		panic(err)
	}

	var writer io.Writer = receiver.writer
	if nil == writer {
		panic(err)
	}

	var newformat string
	{
		var buffer strings.Builder

		switch receiver.style{
		case"color":
			buffer.WriteString("\x1b[48;2;1;1;1m")
			buffer.WriteString("\x1b[38;222;56;43;6m")
		case "":
			buffer.WriteString("[PANIC] ")
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

	panic(err)
}
