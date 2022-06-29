package log

import (
	"bytes"

	"testing"
)

func TestInternalLogger_Tracef(t *testing.T) {

	tests := []struct{
		Format string
		Array []interface{}
		Expected string
	}{
		{
			Format: "",
			Array: []interface{}(nil),
			Expected: "[trace] \n",
		},
		{
			Format: "",
			Array: []interface{}{},
			Expected: "[trace] \n",
		},



		{
			Format: "hello world",
			Array: []interface{}(nil),
			Expected: "[trace] hello world\n",
		},
		{
			Format: "hello world",
			Array: []interface{}{},
			Expected: "[trace] hello world\n",
		},



		{
			Format: "hello %s",
			Array: []interface{}{"Joe"},
			Expected: "[trace] hello Joe\n",
		},
		{
			Format: "hello %s %s",
			Array: []interface{}{"Joe", "Blow"},
			Expected: "[trace] hello Joe Blow\n",
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		logger := NewLogger(&buffer)

		logger.Tracef(test.Format, test.Array...)

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For tst #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
