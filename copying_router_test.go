package flog


import (
	"testing"

	"fmt"
	"math/rand"
	"time"
)


func TestNewCopyingRouter(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))


	router := NewCopyingRouter(NewDiscardingRouter())
	if nil == router {
		t.Errorf("After trying to create a copying router, expected it to be not nil, but was: %v", router)
	}


	lenInitialCopies := 0
	for range router.Copies() {
		lenInitialCopies++
	}
	if expected, actual := 0, lenInitialCopies; expected != actual {
		t.Errorf("After creating a copying router, expected copies to be length %d, but was actually %d,", expected, actual)
	}


	message := fmt.Sprint("%x", randomness.Int63n(9999999999))

	context := make(map[string]interface{})
	limit := randomness.Int63n(30)
	for i:=int64(0); i<limit; i++ {
		context[ fmt.Sprintf("%x", randomness.Int63n(1000*limit)) ] = fmt.Sprintf("%x", randomness.Int63n(999999999999999))
	}

	router.Route(message, context) // Just make sure it doesn't panic or deadlok, by calling this.
}


func TestCopyingRouterRoute(t *testing.T) {

	tests := []struct{
		Data []struct{Message string ; Context map[string]interface{}}
	}{
		{
			Data: []struct{Message string ; Context map[string]interface{}}{
			},
		},


		{
			Data: []struct{Message string ; Context map[string]interface{}}{
				struct{Message string ; Context map[string]interface{}}{
					Message: "apple banana cherry",
					Context: map[string]interface{}{
						"one":1,
						"two":2,
						"three":3,
					},
				},
			},
		},


		{
			Data: []struct{Message string ; Context map[string]interface{}}{
				struct{Message string ; Context map[string]interface{}}{
					Message: "apple",
					Context: map[string]interface{}{
						"one":1,
					},
				},
				struct{Message string ; Context map[string]interface{}}{
					Message: "banana",
					Context: map[string]interface{}{
						"two":2,
					},
				},
				struct{Message string ; Context map[string]interface{}}{
					Message: "cherry",
					Context: map[string]interface{}{
						"cherry":3,
					},
				},
			},
		},
	}


TLoop:	for testNumber, test := range tests {

		router := NewCopyingRouter(NewDiscardingRouter())

		for _, datum := range test.Data {
			router.Route(datum.Message, datum.Context)
		}


		lenCopies := 0
		for range router.Copies() {
			lenCopies++
		}
		if expected, actual := len(test.Data), lenCopies; expected != actual {
			t.Errorf("For test #%d, after creating a copying router and (potentially) doing some routing, expected copies to be length %d, but was actually %d.", testNumber, expected, actual)
			continue TLoop
		}


		datumNumber := 0
		for actualDatum := range router.Copies() {
			if expected, actual := test.Data[datumNumber].Message, actualDatum.Message; expected != actual {
				t.Errorf("For test #%d, after creating a copying router and (potentially) doing some routing, expected message for copies datum #%d to be %q, but was actually %q.", testNumber, datumNumber, expected, actual)
				continue TLoop
			}

			if expected, actual := len(test.Data[datumNumber].Context), len(actualDatum.Context); expected != actual {
				t.Errorf("For test #%d, after creating a copying router and (potentially) doing some routing, expected length of context for copies datum #%d to be %d, but was actually %d.", testNumber, datumNumber, expected, actual)
				continue TLoop
			}

			for expectedKey, expectedValue := range test.Data[datumNumber].Context {
				if _, ok := actualDatum.Context[expectedKey]; !ok {
					t.Errorf("For test #%d, after creating a copying router and (potentially) doing some routing, expected context for copies datum #%d to have key %q, but didn't.", testNumber, datumNumber, expectedKey)
					continue TLoop
				}

				if actualValue := actualDatum.Context[expectedKey]; expectedValue != actualValue {
					t.Errorf("For test #%d, after creating a copying router and (potentially) doing some routing, expected value for context for copies datum #%d at key %q to have value [%v], but actually had [%v].", testNumber, datumNumber, expectedKey, expectedValue, actualValue)
					continue TLoop
				}
			}

			datumNumber++
		}
	}
}
