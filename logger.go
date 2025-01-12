package log

type Logger interface {
	Prefixer

	Alert(...interface{}) error
	Alertf(string, ...interface{}) error

	Begin(...interface{}) Logger

	Debug(...interface{})
	Debugf(string, ...interface{})

	End(...interface{})

	Error(...interface{}) error
	Errorf(string, ...interface{}) error

	Fatal(...interface{})
	Fatalf(string, ...interface{})

	Highlight(...interface{})
	Highlightf(string, ...interface{})

	Inform(...interface{})
	Informf(string, ...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
}
