package log

func (receiver *internalLogger) UnmuteAlert() {
	receiver.mutedAlert = false
}

func (receiver *internalLogger) UnmuteDebug() {
	receiver.mutedDebug = false
}

func (receiver *internalLogger) UnmuteError() {
	receiver.mutedError = false
}

func (receiver *internalLogger) UnmuteFatal() {
	receiver.mutedFatal = false
}

func (receiver *internalLogger) UnmuteHighlight() {
	receiver.mutedHighlight = false
}

func (receiver *internalLogger) UnmuteInform() {
	receiver.mutedInform = false
}

func (receiver *internalLogger) UnmutePanic() {
	receiver.mutedPanic = false
}

func (receiver *internalLogger) UnmuteTrace() {
	receiver.mutedTrace = false
}

func (receiver *internalLogger) UnmuteWarn() {
	receiver.mutedWarn = false
}
