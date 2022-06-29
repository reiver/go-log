package log


import (
	"fmt"
)


var (
	traceContext = map[string]interface{}{
		"~type":"trace",
	}
)


func (flogger *internalFlogger) Trace(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, traceContext)
}

func (flogger *internalFlogger) Tracef(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, traceContext)
}

func (flogger *internalFlogger) Traceln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, traceContext)
}
