package flog

import (
	"bytes"

	"testing"
)

func TestInternalLogger_Errorf(t *testing.T) {

	tests := []struct{
		Format string
		Array []interface{}
		Expected string
	}{
		{
			Format: "",
			Array: []interface{}(nil),
			Expected: "[ERROR] \n",
		},
		{
			Format: "",
			Array: []interface{}{},
			Expected: "[ERROR] \n",
		},



		{
			Format: "hello world",
			Array: []interface{}(nil),
			Expected: "[ERROR] hello world\n",
		},
		{
			Format: "hello world",
			Array: []interface{}{},
			Expected: "[ERROR] hello world\n",
		},



		{
			Format: "hello %s",
			Array: []interface{}{"Joe"},
			Expected: "[ERROR] hello Joe\n",
		},
		{
			Format: "hello %s %s",
			Array: []interface{}{"Joe", "Blow"},
			Expected: "[ERROR] hello Joe Blow\n",
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		logger := NewLogger(&buffer)

		logger.Errorf(test.Format, test.Array...)

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For tst #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
