package flog


import (
	"fmt"
)


var (
	debugContext = map[string]interface{}{
		"~type":"debug",
	}
)


func (flogger *internalFlogger) Debug(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, debugContext)
}

func (flogger *internalFlogger) Debugf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, debugContext)
}

func (flogger *internalFlogger) Debugln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, debugContext)
}
