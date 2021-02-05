package flog

import (
	"io"
)

type internalLogger struct {
	prefix string
	style  string

	writer io.Writer

	canNotLogDebug      bool
	canNotLogError      bool
	canNotLogFatal      bool
	canNotLogHighlight  bool
	canNotLogInform     bool
	canNotLogPanic      bool
	canNotLogTrace      bool
	canNotLogWarn       bool
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

	return logger
}
