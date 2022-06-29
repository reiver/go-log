package log


import (
	"fmt"
	"os"
)


var (
	fatalContext = map[string]interface{}{
		"~type":"fatal",
	}
)


func (flogger *internalFlogger) Fatal(v ...interface{}) {

	msg := fmt.Sprint(v...)

	flogger.route(msg, fatalContext)
	os.Exit(1)
}

func (flogger *internalFlogger) Fatalf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	flogger.route(msg, fatalContext)
	os.Exit(1)
}

func (flogger *internalFlogger) Fatalln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	flogger.route(msg, fatalContext)
	os.Exit(1)
}
