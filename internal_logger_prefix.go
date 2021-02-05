package flog

import (
	"strings"
)

func (receiver internalLogger) Prefix(newprefix ...string) Logger {
	var buffer strings.Builder

	buffer.WriteString(receiver.prefix)

	for _, s := range newprefix {
		buffer.WriteString(s)
		buffer.WriteString(": ")
	}

	prefix := buffer.String()

	var logger internalLogger = receiver
	logger.prefix += prefix

	return &logger
}
