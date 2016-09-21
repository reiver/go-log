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
				` "ctx"."apple"."type"="string" "ctx"."apple"."value"="one" "ctx"."banana"."type"="int" "ctx"."banana"."value"="2" "ctx"."cherry"."type"="float64" "ctx"."cherry"."value"="3.3" "ctx"."kiwi"."type"="bool" "ctx"."kiwi"."value"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
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
				`"name"="backendapi" "number"="123"`,
				`"text"="Hello world!"`,
				` "ctx"."apple"."type"="string" "ctx"."apple"."value"="one" "ctx"."banana"."type"="int" "ctx"."banana"."value"="2" "ctx"."cherry"."type"="float64" "ctx"."cherry"."value"="3.3" "ctx"."kiwi"."type"="bool" "ctx"."kiwi"."value"="true"`,
				` "error"."type"="*errors.errorString" "error"."text"="test error" `,
				` "when"="`,
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
