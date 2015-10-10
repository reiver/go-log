package flog


import (
	"os"
)


func (flogger *internalFlogger) Fatal(v ...interface{}) {

	flogger.Print(v...)
	os.Exit(1)
}

func (flogger *internalFlogger) Fatalf(format string, v ...interface{}) {

	flogger.Printf(format, v...)
	os.Exit(1)
}

func (flogger *internalFlogger) Fatalln(v ...interface{}) {

	flogger.Println(v...)
	os.Exit(1)
}
