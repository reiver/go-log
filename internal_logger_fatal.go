package log

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func (receiver internalLogger) FatalMuted() bool {
	return receiver.mutedFatal
}

func (receiver internalLogger) Fatal(a ...interface{}) {
	s := fmt.Sprint(a...)

	receiver.Fatalf("%s", s)
}

func (receiver internalLogger) Fatalf(format string, a ...interface{}) {
	if receiver.FatalMuted() {
		os.Exit(1)
		return
	}

	var writer io.Writer = receiver.writer
	if nil == writer {
		os.Exit(1)
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
			buffer.WriteString("[PANIC] ")
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

	os.Exit(1)
}
