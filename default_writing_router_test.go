package flog


import (
	"bytes"
	"errors"
	"strings"

	"testing"
)


func TestDefaultWritingRouterRoute(t *testing.T) {

	tests := []struct{
		Message string
		Context map[string]interface{}
		ExpectContains []string
	}{
		{
			Message: "Hello world!",
			Context: map[string]interface{}{
				"apple": "one",
				"banana": 2,
				"cherry": 3.3,
				"kiwi":   true,
				"~error": errors.New("test error"),
			},
			ExpectContains: []string{
				`"text"="Hello world!"`,
				` "ctx"."apple"="one" "ctx"."banana"="2" "ctx"."cherry"="3.300000" "ctx"."kiwi"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
				` "when"="`,
			},
		},



		{
			Message: "Apple\tBANANA\nCherry",
			Context: map[string]interface{}{
				"apple": "one",
				"banana": 2,
				"cherry": 3.3,
				"kiwi":   true,
				"~error": errors.New("test error"),
				"more": map[string]interface{}{
					"ONE":   "1",
					"TWO":   "2",
					"THREE": "3",
				},
			},
			ExpectContains: []string{
				`"text"="Apple\tBANANA\nCherry"`,
				` "ctx"."apple"="one" "ctx"."banana"="2" "ctx"."cherry"="3.300000" "ctx"."kiwi"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
				` "ctx"."more"."ONE"="1" "ctx"."more"."THREE"="3" "ctx"."more"."TWO"="2"`,
				` "when"="`,
			},
		},



		{
			Message: "Apple\tBANANA\nCherry",
			Context: map[string]interface{}{
				"apple": "one",
				"banana": 2,
				"cherry": 3.3,
				"kiwi":   true,
				"~error": errors.New("test error"),
				"more": map[string]interface{}{
					"ONE":   "1",
					"TWO":   "2",
					"THREE": "3",
					"FOUR":  map[string]interface{}{
						"a": "1st",
						"b": "2nd",
						"c": []string{
							"th",
							"i",
							"rd",
						},
					},
				},
			},
			ExpectContains: []string{
				`"text"="Apple\tBANANA\nCherry"`,
				` "ctx"."apple"="one" "ctx"."banana"="2" "ctx"."cherry"="3.300000" "ctx"."kiwi"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
				` "ctx"."more"."FOUR"."a"="1st" "ctx"."more"."FOUR"."b"="2nd" "ctx"."more"."FOUR"."c"=["th","i","rd"] "ctx"."more"."ONE"="1" "ctx"."more"."THREE"="3" "ctx"."more"."TWO"="2"`,
				` "when"="`,
			},
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var router Router = NewDefaultWritingRouter(&buffer)

		ctx := map[string]interface{}{}
		for n, v := range test.Context {
			ctx[n] = v
		}

		if err := router.Route(test.Message, ctx); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		{
			actual := buffer.String()
			for expectedNumber, expectContains := range test.ExpectContains {

				if !strings.Contains(actual, expectContains) {
					t.Errorf("For test #%d and expected #%d, expect to contain, actual:\n==)>%s<(==\n==)>%s<(==", testNumber, expectedNumber, expectContains, actual)
					continue
				}
			}
		}
	}
}


func TestDefaultWritingRouterWithPrefixRoute(t *testing.T) {

	tests := []struct{
		Message string
		Context map[string]interface{}
		Prefix map[string]interface{}
		ExpectContains []string
	}{
		{
			Message: "Hello world!",
			Context: map[string]interface{}{
				"apple": "one",
				"banana": 2,
				"cherry": 3.3,
				"kiwi":   true,
				"~error": errors.New("test error"),
			},
			Prefix: map[string]interface{}{
				"name": "backendapi",
				"number": "123",
			},
			ExpectContains: []string{
				`"name"="backendapi" "number"="123" "text"="Hello world!" "when"="`,
				` "ctx"."apple"="one" "ctx"."banana"="2" "ctx"."cherry"="3.300000" "ctx"."kiwi"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
			},
		},


		{
			Message: "Apple\tBANANA\nCherry",
			Context: map[string]interface{}{
				"apple": "one",
				"banana": 2,
				"cherry": 3.3,
				"kiwi":   true,
				"~error": errors.New("test error"),
				"more": map[string]interface{}{
					"ONE":   "1",
					"TWO":   "2",
					"THREE": "3",
					"FOUR":  map[string]interface{}{
						"a": "1st",
						"b": "2nd",
						"c": []string{
							"th",
							"i",
							"rd",
						},
					},
				},
			},
			Prefix: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "backendapi",
					"build": map[string]interface{}{
						"number": 123,
						"hash":  "4a844b2",
					},
				},
			},
			ExpectContains: []string{
				`"app"."build"."hash"="4a844b2" "app"."build"."number"="123" "app"."name"="backendapi" "text"="Apple\tBANANA\nCherry" "when"="`,
				` "ctx"."apple"="one" "ctx"."banana"="2" "ctx"."cherry"="3.300000" "ctx"."kiwi"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
				` "ctx"."more"."FOUR"."a"="1st" "ctx"."more"."FOUR"."b"="2nd" "ctx"."more"."FOUR"."c"=["th","i","rd"] "ctx"."more"."ONE"="1" "ctx"."more"."THREE"="3" "ctx"."more"."TWO"="2"`,
			},
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var router Router = NewDefaultWritingRouterWithPrefix(&buffer, test.Prefix)

		ctx := map[string]interface{}{}
		for n, v := range test.Context {
			ctx[n] = v
		}

		if err := router.Route(test.Message, ctx); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		{
			actual := buffer.String()
			for expectedNumber, expectContains := range test.ExpectContains {

				if !strings.Contains(actual, expectContains) {
					t.Errorf("For test #%d and expected #%d, expect to contain, actual:\n==)>%s<(==\n==)>%s<(==", testNumber, expectedNumber, expectContains, actual)
					continue
				}
			}
		}
	}
}
