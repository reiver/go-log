package log


import (
	"testing"
)


func TestWith(t *testing.T) {

	tests := []struct{
		Cascade1    []interface{}
		Cascade2    []interface{}
		Expected map[string]interface{}
	}{
		{
			Cascade1: []interface{}{},
			Cascade2: []interface{}{},
			Expected: map[string]interface{}{},
		},


		{
			Cascade1: []interface{}{
				"apple",
			},
			Cascade2: []interface{}{
			},
			Expected: map[string]interface{}{
				"text":"apple",
			},
		},
		{
			Cascade1: []interface{}{
			},
			Cascade2: []interface{}{
				"apple",
			},
			Expected: map[string]interface{}{
				"text":"apple",
			},
		},
		{
			Cascade1: []interface{}{
				"apple",
			},
			Cascade2: []interface{}{
				"apple",
			},
			Expected: map[string]interface{}{
				"text":"apple",
			},
		},


		{
			Cascade1: []interface{}{
				"apple",
			},
			Cascade2: []interface{}{
				"banana",
			},
			Expected: map[string]interface{}{
				"text":"banana",
			},
		},


		{
			Cascade1: []interface{}{
				"apple",
				"banana",
				"cherry",
			},
			Cascade2: []interface{}{
			},
			Expected: map[string]interface{}{
				"text":"cherry",
			},
		},
		{
			Cascade1: []interface{}{
				"apple",
				"banana",
			},
			Cascade2: []interface{}{
				"cherry",
			},
			Expected: map[string]interface{}{
				"text":"cherry",
			},
		},
		{
			Cascade1: []interface{}{
				"apple",
			},
			Cascade2: []interface{}{
				"banana",
				"cherry",
			},
			Expected: map[string]interface{}{
				"text":"cherry",
			},
		},
		{
			Cascade1: []interface{}{
			},
			Cascade2: []interface{}{
				"apple",
				"banana",
				"cherry",
			},
			Expected: map[string]interface{}{
				"text":"cherry",
			},
		},


		{
			Cascade1: []interface{}{
				map[string]string{
				},
			},
			Cascade2: []interface{}{
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
			},
		},


		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade1: []interface{}{
				map[string]string{
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},


		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},
		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"cherry":"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},
		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"banana":"two",
					"cherry":"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},
		{
			Cascade1: []interface{}{
				map[string]string{
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},


		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"kiwi":"four",
					"watermelon":"five",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
					"kiwi":"four",
					"watermelon":"five",
			},
		},

		{
			Cascade1: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
					"fig":"THIS SHOULD BE REPLACED",
				},
			},
			Cascade2: []interface{}{
				map[string]string{
					"fig":"THIS SHOULD REMAIN",
					"kiwi":"four",
					"watermelon":"five",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
					"fig":"THIS SHOULD REMAIN",
					"kiwi":"four",
					"watermelon":"five",
			},
		},
	}


TestLoop:
	for testNumber, test := range tests {

		flogger1 := New(NewDiscardingRouter(), test.Cascade1...)

		flogger2 := flogger1.With(test.Cascade2...)

		context := flogger2.(*internalFlogger).context

		if expected, actual := len(test.Expected), len(context); expected != actual {
			t.Errorf("For test #%d, expected length to be %d but actually was %d.", testNumber, expected, actual)
	               continue TestLoop
		}

		for expectedKey, expectedValue := range test.Expected {
			if _, ok := context[expectedKey]; !ok {
				t.Errorf("For test #%d, expected key %q to be in resulting context but wasn't.", testNumber, expectedKey)
				continue TestLoop
			}

			if expected, actual := expectedValue, context[expectedKey]; expected != actual {
                                t.Errorf("For test #%d, expected value for key %q to be %q in resulting context, but was actually %q.", testNumber, expectedKey, expected, actual)
				continue TestLoop
			}
		}
	}
}
