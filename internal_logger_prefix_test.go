package log

import (
	"bytes"
	"fmt"

	"testing"
)

func TestLoggerPrefix(t *testing.T) {
	tests := []struct{
		NewPrefix []string
		Expected string
	}{

		{
			NewPrefix: []string(nil),
			Expected: "",
		},
		{
			NewPrefix: []string{},
			Expected: "",
		},



		{
			NewPrefix: []string{""},
			Expected: ": ",
		},
		{
			NewPrefix: []string{"", ""},
			Expected: ": : ",
		},
		{
			NewPrefix: []string{"", "", ""},
			Expected: ": : : ",
		},

		{
			NewPrefix: []string{"", "", "", "", "", "", "", "", "", ""},
			Expected: ": : : : : : : : : : ",
		},



		{
			NewPrefix: []string{"ONE"},
			Expected: "ONE: ",
		},
		{
			NewPrefix: []string{"ONE", "TWO"},
			Expected: "ONE: TWO: ",
		},
		{
			NewPrefix: []string{"ONE", "TWO", "THREE"},
			Expected: "ONE: TWO: THREE: ",
		},
	}

	for testNumber, test := range tests {

		var logger internalLogger

		newLogger := logger.Prefix(test.NewPrefix...)

		newInternalLogger, casted := newLogger.(*internalLogger)
		if !casted {
			t.Errorf("For test #%d, could not cast to log.internalLogger.", testNumber)
			t.Logf("TYPE: %T", newLogger)
			continue
		}

		if expected, actual := test.Expected, newInternalLogger.prefix; expected != actual {
			t.Errorf("For test #%d, the actual prefix is not what was expected.", testNumber)
			t.Log("EXPECTED:", expected)
			t.Log("ACTUAL:  ", actual)
			continue
		}
	}
}

func TestLoggerPrefix_inform(t *testing.T) {

	var buffer bytes.Buffer

	log := NewLogger(&buffer).Prefix("one","two","three")

	func(){
		defer func(){
			if r := recover(); nil != r {
				fmt.Fprintf(&buffer, "RECOVER:%v\n",r)
			}
		}()
		log.Panic("hello PANIC")
	}()
	log.Inform("hello INFORM")
	log.Highlight("hello HIGHLIGHT")
	log.Error("hello ERROR")
	log.Warn("hello WARN")
	log.Debug("hello DEBUG")
	log.Trace("hello TRACE")


	log = log.Prefix("four")
	log.Highlight("new prefix FOUR")


	const expected = "[PANIC] one: two: three: hello PANIC"              + "\n" +
	                                  "RECOVER:hello PANIC"              + "\n" +
	                "[inform] one: two: three: hello INFORM"             + "\n" +
	                "[HIGHLIGHT] one: two: three: hello HIGHLIGHT"       + "\n" +
	                "[ERROR] one: two: three: hello ERROR"               + "\n" +
	                "[warn] one: two: three: hello WARN"                 + "\n" +
	                "[debug] one: two: three: hello DEBUG"               + "\n" +
	                "[trace] one: two: three: hello TRACE"               + "\n" +
	                "[HIGHLIGHT] one: two: three: four: new prefix FOUR" + "\n"

	if actual := buffer.String(); expected != actual {
		t.Error("The actual logs is not what was expected.")
		t.Log("EXPECTED:\n", expected)
		t.Log("ACTUAL:\n", actual)
		return
	}
}
