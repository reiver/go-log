package log

import (
	"fmt"
	"io"
	"strings"
)

func (receiver internalLogger) HighlightMuted() bool {
	return receiver.mutedHighlight
}

func (receiver internalLogger) Highlight(a ...interface{}) {
	if receiver.HighlightMuted() {
		return
	}
	if nil == receiver.writer {
		return
	}

	s := fmt.Sprint(a...)

	receiver.Highlightf("%s", s)
}

func (receiver internalLogger) Highlightf(format string, a ...interface{}) {
	if receiver.HighlightMuted() {
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
			buffer.WriteString("\x1b[48;2;153;0;17m")
			buffer.WriteString("\x1b[38;2;252;246;245m")
		case "":
			buffer.WriteString("[HIGHLIGHT] ")
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
