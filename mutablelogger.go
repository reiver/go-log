package log

type MutableLogger interface {
	Logger

	MuteAlert()
	UnmuteAlert()
	AlertMuted() bool

	MuteDebug()
	UnmuteDebug()
	DebugMuted() bool

	MuteError()
	UnmuteError()
	ErrorMuted() bool

	MuteFatal()
	UnmuteFatal()
	FatalMuted() bool

	MuteHighlight()
	UnmuteHighlight()
	HighlightMuted() bool

	MuteInform()
	UnmuteInform()
	InformMuted() bool

	MutePanic()
	UnmutePanic()
	PanicMuted() bool

	MuteTrace()
	UnmuteTrace()
	TraceMuted() bool

	MuteWarn()
	UnmuteWarn()
	WarnMuted() bool
}
