package flog


import (
	"fmt"
)


var (
	errorContext = map[string]interface{}{
		"~type":"error",
	}
)


func (flogger *internalFlogger) Error(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, errorContext)
}

func (flogger *internalFlogger) Errorf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, errorContext)
}

func (flogger *internalFlogger) Errorfe(err error, format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	context := map[string]interface{}{}
	for k,v := range errorContext {
		context[k] = v
	}
	context["~error"] = err

	flogger.route(msg, context)
}

func (flogger *internalFlogger) Errorln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, errorContext)
}
