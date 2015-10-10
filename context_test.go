package flog


import (
	"testing"
)


func TestNewContext(t *testing.T) {

	tests := []struct{
		Cascade []interface{}
		Expected map[string]interface{}
	}{
		{
			Cascade: []interface{}{},
			Expected: map[string]interface{}{},
		},


		{
			Cascade: []interface{}{
				"apple",
			},
			Expected: map[string]interface{}{
				"text":"apple",
			},
		},
		{
			Cascade: []interface{}{
				"apple",
				"banana",
			},
			Expected: map[string]interface{}{
				"text":"banana",
			},
		},
		{
			Cascade: []interface{}{
				"apple",
				"banana",
				"cherry",
			},
			Expected: map[string]interface{}{
				"text":"cherry",
			},
		},


		{
			Cascade: []interface{}{
				map[string]string{
				},
			},
			Expected: map[string]interface{}{},
		},


		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade: []interface{}{
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
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
				},
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
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
					"fig":"THIS SHOULD BE REPLACED",
				},
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

		context := newContext(test.Cascade...)

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
