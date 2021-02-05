package flog

type Logger interface {
	CanLogDebug() bool
	Debug(...interface{})
	Debugf(string, ...interface{})

	CanLogError() bool
	Error(...interface{}) error
	Errorf(string, ...interface{}) error

	CanLogFatal() bool
	Fatal(...interface{})
	Fatalf(string, ...interface{})

	CanLogHighlight() bool
	Highlight(...interface{})
	Highlightf(string, ...interface{})

	CanLogInform() bool
	Inform(...interface{})
	Informf(string, ...interface{})

	CanLogPanic() bool
	Panic(...interface{})
	Panicf(string, ...interface{})

	CanLogTrace() bool
	Trace(...interface{})
	Tracef(string, ...interface{})

	CanLogWarn() bool
	Warn(...interface{})
	Warnf(string, ...interface{})

	Prefix(...string) Logger
}
