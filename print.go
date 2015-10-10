package flog


import (
	"fmt"
)


func (flogger *internalFlogger) Print(v ...interface{}) {

	msg := fmt.Sprint(v...)
//@TODO: Do we really want this here?
//	flogger.logs = append(flogger.logs, msg)

	flogger.router.Route(msg, flogger.context)
}

func (flogger *internalFlogger) Printf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
//@TODO: Do we really want this here?
//	flogger.logs = append(flogger.logs, msg)

	flogger.router.Route(msg, flogger.context)
}

func (flogger *internalFlogger) Println(v ...interface{}) {

	msg := fmt.Sprintln(v...)
//@TODO: Do we really want this here?
//	flogger.logs = append(flogger.logs, msg)

	flogger.router.Route(msg, flogger.context)
}
