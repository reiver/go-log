package flog


import (
	"fmt"
)


func (flogger *internalFlogger) Print(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.router.Route(msg, flogger.context)
}

func (flogger *internalFlogger) Printf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.router.Route(msg, flogger.context)
}

func (flogger *internalFlogger) Println(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.router.Route(msg, flogger.context)
}
