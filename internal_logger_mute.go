package log

func (receiver *internalLogger) MuteAlert() {
	receiver.mutedDebug = true
}

func (receiver *internalLogger) MuteDebug() {
	receiver.mutedDebug = true
}

func (receiver *internalLogger) MuteError() {
	receiver.mutedError = true
}

func (receiver *internalLogger) MuteFatal() {
	receiver.mutedFatal = true
}

func (receiver *internalLogger) MuteHighlight() {
	receiver.mutedHighlight = true
}

func (receiver *internalLogger) MuteInform() {
	receiver.mutedInform = true
}

func (receiver *internalLogger) MutePanic() {
	receiver.mutedPanic = true
}

func (receiver *internalLogger) MuteTrace() {
	receiver.mutedTrace = true
}

func (receiver *internalLogger) MuteWarn() {
	receiver.mutedWarn = true
}
