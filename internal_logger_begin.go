package flog

import (
	"runtime"
)

func (receiver internalLogger) Begin(a ...interface{}) Logger {

	var funcName string = "<([-UNKNOWN-])>"
	{
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			fn := runtime.FuncForPC(pc)
			funcName = fn.Name()
		}
	}

	logger := receiver.Prefix(funcName)

	a = append([]interface{}{"BEGIN"}, a...)

	logger.Debug(a...)

	return logger
}
