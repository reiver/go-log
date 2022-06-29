package log

import (
	"testing"
)

func TestInternalLoggerIsLogger(t *testing.T) {

	var x Logger = &internalLogger{} // THIS IS WHAT ACTUALLY MATTERS.

	if nil == x {
		t.Error("This should never happen.")
	}
}
