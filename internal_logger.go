package log

import (
	"io"
	"time"
)

type internalLogger struct {
	prefix string
	style  string

	begin time.Time

	writer io.Writer

	mutedAlert      bool
	mutedDebug      bool
	mutedError      bool
	mutedFatal      bool
	mutedHighlight  bool
	mutedInform     bool
	mutedPanic      bool
	mutedTrace      bool
	mutedWarn       bool
}

func NewLogger(writer io.Writer, parameters ...string) Logger {
	logger := internalLogger{
		writer:writer,
	}
	if 1 <= len(parameters) {
		style := parameters[0]
		switch style {
		case "color","colour":
			logger.style = "color"
		default:
			logger.style = ""
		}
	}

	return &logger
}
