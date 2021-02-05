package flog

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	MuteDebug()
	DebugMuted() bool

	Error(...interface{}) error
	Errorf(string, ...interface{}) error
	MuteError()
	ErrorMuted() bool

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	MuteFatal()
	FatalMuted() bool

	Highlight(...interface{})
	Highlightf(string, ...interface{})
	MuteHighlight()
	HighlightMuted() bool

	Inform(...interface{})
	Informf(string, ...interface{})
	MuteInform()
	InformMuted() bool

	Panic(...interface{})
	Panicf(string, ...interface{})
	MutePanic()
	PanicMuted() bool

	Trace(...interface{})
	Tracef(string, ...interface{})
	MuteTrace()
	TraceMuted() bool

	Warn(...interface{})
	Warnf(string, ...interface{})
	MuteWarn()
	WarnMuted() bool

	Prefix(...string) Logger
}
