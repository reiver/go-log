package flog

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	MuteDebug()
	UnmuteDebug()
	DebugMuted() bool

	Error(...interface{}) error
	Errorf(string, ...interface{}) error
	MuteError()
	UnmuteError()
	ErrorMuted() bool

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	MuteFatal()
	UnmuteFatal()
	FatalMuted() bool

	Highlight(...interface{})
	Highlightf(string, ...interface{})
	MuteHighlight()
	UnmuteHighlight()
	HighlightMuted() bool

	Inform(...interface{})
	Informf(string, ...interface{})
	MuteInform()
	UnmuteInform()
	InformMuted() bool

	Panic(...interface{})
	Panicf(string, ...interface{})
	MutePanic()
	UnmutePanic()
	PanicMuted() bool

	Trace(...interface{})
	Tracef(string, ...interface{})
	MuteTrace()
	UnmuteTrace()
	TraceMuted() bool

	Warn(...interface{})
	Warnf(string, ...interface{})
	MuteWarn()
	UnmuteWarn()
	WarnMuted() bool

	Prefix(...string) Logger
}
