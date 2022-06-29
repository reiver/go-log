package log

import (
	"bytes"

	"testing"
)

func TestInternalLogger_Informf(t *testing.T) {

	tests := []struct{
		Format string
		Array []interface{}
		Expected string
	}{
		{
			Format: "",
			Array: []interface{}(nil),
			Expected: "[inform] \n",
		},
		{
			Format: "",
			Array: []interface{}{},
			Expected: "[inform] \n",
		},



		{
			Format: "hello world",
			Array: []interface{}(nil),
			Expected: "[inform] hello world\n",
		},
		{
			Format: "hello world",
			Array: []interface{}{},
			Expected: "[inform] hello world\n",
		},



		{
			Format: "hello %s",
			Array: []interface{}{"Joe"},
			Expected: "[inform] hello Joe\n",
		},
		{
			Format: "hello %s %s",
			Array: []interface{}{"Joe", "Blow"},
			Expected: "[inform] hello Joe Blow\n",
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		logger := NewLogger(&buffer)

		logger.Informf(test.Format, test.Array...)

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For tst #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
