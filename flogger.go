package flog


type Flogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
	Errorfe(error, string, ...interface{})
	Errorln(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicfv(interface{}, string, ...interface{})
	Panicln(...interface{})

	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})
	Traceln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	With(...interface{}) Flogger
}
