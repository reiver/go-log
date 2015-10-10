package flog


import (
	"testing"

	"fmt"
	"math/rand"
	"time"
)


func TestFilteredRouterJustCreated(t *testing.T) {

	router := NewFilteredRouter()
	router.Register(NewDiscardRouter(), func(string, map[string]interface{}) bool {
		return false
	})
	if nil == router {
		t.Errorf("After trying to create a filtered router, expected it to be not nil, but was: %v", router)
	}

}


func TestFilteredRouterJustFilterParameters(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))


	var filterMessage string
	var filterContext map[string] interface{}
	var filterResult = false

	router := NewFilteredRouter()

	router.Register(NewDiscardRouter(), func(message string, context map[string]interface{}) bool {
		filterMessage = message
		filterContext = context

		filterResult = !filterResult

		return filterResult
	})

	const NUM_TESTS = 20
TLoop:	for testNumber:=0; testNumber<NUM_TESTS; testNumber++ {
		msg := fmt.Sprintf("string with random numbers: %d", randomness.Int63n(9999999999))

		ctx := make(map[string]interface{})
		lenCtx := randomness.Int63n(30)
		for i:=int64(0); i<lenCtx; i++ {
			ctx[ fmt.Sprintf("%x", randomness.Int63n(1000*lenCtx)) ] = fmt.Sprintf("%x", randomness.Int63n(999999999999999))
		}


		router.Route(msg, ctx)

		if expected, actual := msg, filterMessage; expected != actual {
			t.Errorf("For test #%d, expected message passed to filter func to be %q, but actually was %q.", testNumber, expected, actual)
			continue TLoop
		}

		if expected, actual := len(ctx), len(filterContext); expected != actual {
			t.Errorf("For test #%d, expected context passed to filter func to be len %d, but actually was %d.", testNumber, expected, actual)
			continue TLoop
		}

		for expectedKey, expectedValue := range ctx {
			if _, ok := filterContext[expectedKey]; !ok {
				t.Errorf("For test #%d, expected context passed to filter func to have key %q, but didn't.", testNumber, expectedKey)
				continue TLoop
			}

			if expected, actual := expectedValue.(string), filterContext[expectedKey].(string); expected != actual {
				t.Errorf("For test #%d, expected context passed to filter func for key %q to have value %q, but actuall had %q.", testNumber, expectedKey, expected, actual)
				continue TLoop
			}
		}
	}

}
