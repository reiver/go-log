package log


import (
	"errors"
	"reflect"

	"testing"
)


func TestNewErrorContext(t *testing.T) {

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")
	err4 := errors.New("error 4")
	err5 := errors.New("error 5")


	tests := []struct{
		BaseContext map[string]interface{}
		V                    []interface{}
		Expected    map[string]interface{}
	}{
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				// Nothing here.
			},
			Expected: map[string]interface{}{
				"~type":"error",
			},
		},



		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				1,
			},
			Expected: map[string]interface{}{
				"~type":"error",
			},
		},
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				1, "two",
			},
			Expected: map[string]interface{}{
				"~type":"error",
			},
		},
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				1, "two", 3.0,
			},
			Expected: map[string]interface{}{
				"~type":"error",
			},
		},



		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				1, "two", 3.0,
			},
			Expected: map[string]interface{}{
				"~type":"error",
			},
		},



		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				err1,
			},
			Expected: map[string]interface{}{
				"~type":"error",
				"~errors": []error{
					err1,
				},
			},
		},
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				err1, err2,
			},
			Expected: map[string]interface{}{
				"~type":"error",
				"~errors": []error{
					err1, err2,
				},
			},
		},
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				err1, err2, err3,
			},
			Expected: map[string]interface{}{
				"~type":"error",
				"~errors": []error{
					err1, err2, err3,
				},
			},
		},
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				err1, err2, err3, err4,
			},
			Expected: map[string]interface{}{
				"~type":"error",
				"~errors": []error{
					err1, err2, err3, err4,
				},
			},
		},
		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				err1, err2, err3, err4, err5,
			},
			Expected: map[string]interface{}{
				"~type":"error",
				"~errors": []error{
					err1, err2, err3, err4, err5,
				},
			},
		},



		{
			BaseContext: map[string]interface{}{
				"~type":"error",
			},
			V: []interface{}{
				1, err1, "two", err2, 3.0, err3, '4',
			},
			Expected: map[string]interface{}{
				"~type":"error",
				"~errors": []error{
					err1, err2, err3,
				},
			},
		},
	}


	for testNumber, test := range tests {

		actualContext := newErrorContext(test.BaseContext, test.V...)
		if actual := actualContext; nil == actual {
			t.Errorf("For test #%d, did not expected nil, but actually got %v", testNumber, actual)
			continue
		}

		if expected, actual := len(test.Expected), len(actualContext); expected != actual  {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.Expected, actualContext; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, expected...\n%#v\nbut actually got...\n%#v", testNumber, expected, actual)
			continue
		}
	}
}
