package flog

type Logger interface {
	CanLogDebug() bool
	Debug(...interface{})
	Debugf(string, ...interface{})
	MuteDebug()

	CanLogError() bool
	Error(...interface{}) error
	Errorf(string, ...interface{}) error
	MuteError()

	CanLogFatal() bool
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	MuteFatal()

	CanLogHighlight() bool
	Highlight(...interface{})
	Highlightf(string, ...interface{})
	MuteHighlight()

	CanLogInform() bool
	Inform(...interface{})
	Informf(string, ...interface{})
	MuteInform()

	CanLogPanic() bool
	Panic(...interface{})
	Panicf(string, ...interface{})
	MutePanic()

	CanLogTrace() bool
	Trace(...interface{})
	Tracef(string, ...interface{})
	MuteTrace()

	CanLogWarn() bool
	Warn(...interface{})
	Warnf(string, ...interface{})
	MuteWarn()

	Prefix(...string) Logger
}
