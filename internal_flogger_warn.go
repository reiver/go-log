package log


import (
	"fmt"
)


var (
	warnContext = map[string]interface{}{
		"~type":"warn",
	}
)


func (flogger *internalFlogger) Warn(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, warnContext)
}

func (flogger *internalFlogger) Warnf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, warnContext)
}

func (flogger *internalFlogger) Warnln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, warnContext)
}
