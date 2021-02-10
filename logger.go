package flog

type Logger interface {
	Alert(...interface{}) error
	Alertf(string, ...interface{}) error
	MuteAlert()
	UnmuteAlert()
	AlertMuted() bool

	Begin(...interface{}) Logger

	Debug(...interface{})
	Debugf(string, ...interface{})
	MuteDebug()
	UnmuteDebug()
	DebugMuted() bool

	End(...interface{})

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
