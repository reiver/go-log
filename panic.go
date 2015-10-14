package flog


import (
	"fmt"
)


var (
	panicContext = map[string]interface{}{
		"~type":"panic",
	}
)


func (flogger *internalFlogger) Panic(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, panicContext)
	panic(msg)
}

func (flogger *internalFlogger) Panicf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, panicContext)
	panic(msg)
}

func (flogger *internalFlogger) Panicfv(panicValue interface{}, format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, panicContext)
	panic(panicValue)
}

func (flogger *internalFlogger) Panicln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, panicContext)
	panic(msg)
}
