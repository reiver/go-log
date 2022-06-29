package log

import (
	"runtime"
	"time"
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
	switch casted := logger.(type) {
	case *internalLogger:
		casted.begin = time.Now()
	}

	a = append([]interface{}{"BEGIN "}, a...)

	logger.Debug(a...)

	return logger
}
