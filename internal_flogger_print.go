package log


import (
	"fmt"
)


var (
	printContext = map[string]interface{}{
		"~type":"print",
	}
)


func (flogger *internalFlogger) Print(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, printContext)
}

func (flogger *internalFlogger) Printf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, printContext)
}

func (flogger *internalFlogger) Println(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, printContext)
}
