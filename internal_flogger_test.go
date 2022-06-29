package log


import (
	"testing"
)


func TestNew(t *testing.T) {

	flogger := New(NewDiscardingRouter())
	if nil == flogger {
		t.Errorf("Expected created flogger to not be nil, but was: %v", flogger)
	}
}


func TestNewForContext(t *testing.T) {

	tests := []struct{
		Cascade     []interface{}
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
			Expected: map[string]interface{}{
			},
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
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
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

		flogger := New(NewDiscardingRouter(), test.Cascade...)

		context := flogger.(*internalFlogger).context

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


func TestInternalFloggerRouteNilReceiver(t *testing.T) {

	tests := []struct{
		Message string
		MoreContext map[string]interface{}
	}{
		{
			Message: "",
			MoreContext: nil,
		},
		{
			Message: "",
			MoreContext: map[string]interface{}{},
		},
		{
			Message: "",
			MoreContext: map[string]interface{}{
				"apple":  "one",
				"banana": 2,
				"cherry": '3',
				"kiwi":   4.0,
			},
		},



		{
			Message: "Hello world!",
			MoreContext: nil,
		},
		{
			Message: "Hello world!",
			MoreContext: map[string]interface{}{},
		},
		{
			Message: "Hello world!",
			MoreContext: map[string]interface{}{
				"apple":  "one",
				"banana": 2,
				"cherry": '3',
				"kiwi":   4.0,
			},
		},



		{
			Message: " ",
			MoreContext: nil,
		},
		{
			Message: " ",
			MoreContext: map[string]interface{}{},
		},
		{
			Message: " ",
			MoreContext: map[string]interface{}{
				"apple":  "one",
				"banana": 2,
				"cherry": '3',
				"kiwi":   4.0,
			},
		},



		{
			Message: "one\ntwo\tthree\r\n",
			MoreContext: nil,
		},
		{
			Message: "one\ntwo\tthree\r\n",
			MoreContext: map[string]interface{}{},
		},
		{
			Message: "one\ntwo\tthree\r\n",
			MoreContext: map[string]interface{}{
				"apple":  "one",
				"banana": 2,
				"cherry": '3',
				"kiwi":   4.0,
			},
		},
	}



	for testNumber, test := range tests {

		var flogger *internalFlogger = nil

		err := flogger.route(test.Message, test.MoreContext)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %v", testNumber, err)
			continue
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("For test #%d, expected an error (%T) %q, but actually got (%T) %q", testNumber, expected, expected, actual, actual)
			continue
		}

	}
}


func TestInternalFloggerRouteNilRouter(t *testing.T) {

	moreContexts := []map[string]interface{}{
		nil,
		map[string]interface{}{},
		map[string]interface{}{
			"apple":  "one",
			"banana": 2,
			"cherry": '3',
			"kiwi":   4.0,
		},
	}


	messages := []string{
		"",
		"Hello world!",
		" ",
		"one\ntwo\tthree\r\n",
	}


	tests := []struct{
		Context map[string]interface{}
	}{
		{
			Context: nil,
		},
		{
			Context: map[string]interface{}{},
		},



		{
			Context: map[string]interface{}{
				"apple":  "one",
				"banana": 2,
				"cherry": '3',
				"kiwi":   4.0,
			},
		},
	}


	for testNumber, test := range tests {
		var flogger internalFlogger

		flogger.context = test.Context
		flogger.router  = nil

		for messageNumber, message := range messages {
			for moreContextNumber, moreContext := range moreContexts {
				err := flogger.route(message, moreContext)
				if nil == err {
					t.Errorf("For test #%d and message #%d and more context #%d, expected an error, but did not actually get one: %v", testNumber, messageNumber, moreContextNumber, err)
					continue
				}
				if expected, actual := errNilRouter, err; expected != actual {
					t.Errorf("For test #%d and message #%d and more context #%d, expected an error (%T) %q, but actually got (%T) %q", testNumber, messageNumber, moreContextNumber, expected, expected, actual, actual)
					continue
				}
			}
		}
	}
}
