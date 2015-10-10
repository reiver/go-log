package flog


import (
	"fmt"
)


func (flogger *internalFlogger) Panic(v ...interface{}) {

	flogger.Print(v...)
	panic(fmt.Sprint(v...))
}

func (flogger *internalFlogger) Panicf(format string, v ...interface{}) {

	flogger.Printf(format, v...)
	panic(fmt.Sprintf(format, v...))
}

func (flogger *internalFlogger) Panicln(v ...interface{}) {

	flogger.Println(v...)
	panic(fmt.Sprintln(v...))
}
