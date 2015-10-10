package flog


import (
	"testing"

	"github.com/reiver/go-manyerrors"

	"fmt"
	"math/rand"
	"time"
)


func TestNewFanoutRouter(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))


	router := NewFanoutRouter()
	if nil == router {
		t.Errorf("After trying to create a discard router, expected it to be not nil, but was: %v", router)
	}


	message := fmt.Sprint("%x", randomness.Int63n(9999999999))

	context := make(map[string]interface{})
	limit := randomness.Int63n(30)
	for i:=int64(0); i<limit; i++ {
		context[ fmt.Sprintf("%x", randomness.Int63n(1000*limit)) ] = fmt.Sprintf("%x", randomness.Int63n(999999999999999))
	}

	router.Route(message, context) // Just make sure it doesn't panic or deadlok, by calling this.
}


func TestFanoutRouterRoute(t *testing.T) {

	numApples   := 0
	numBananas  := 0
	numCherries := 0
	numFigs     := 0
	numKiwis    := 0
	numApplesOrFigs := 0

	appleRouter := NewFilteringRouter(NewDiscardingRouter(), func(message string, context map[string]interface{}) bool {
		if "apple" == message {
			numApples++
			return true
		}

		return false
	})
	bananaRouter := NewFilteringRouter(NewDiscardingRouter(), func(message string, context map[string]interface{}) bool {
		if "banana" == message {
			numBananas++
			return true
		}

		return false
	})
	cherryRouter := NewFilteringRouter(NewDiscardingRouter(), func(message string, context map[string]interface{}) bool {
		if "cherry" == message {
			numCherries++
			return true
		}

		return false
	})
	figRouter := NewFilteringRouter(NewDiscardingRouter(), func(message string, context map[string]interface{}) bool {
		if "fig" == message {
			numFigs++
			return true
		}

		return false
	})
	kiwiRouter := NewFilteringRouter(NewDiscardingRouter(), func(message string, context map[string]interface{}) bool {
		if "kiwi" == message {
			numKiwis++
			return true
		}

		return false
	})
	appleOrFigRouter := NewFilteringRouter(NewDiscardingRouter(), func(message string, context map[string]interface{}) bool {
		if "apple" == message || "fig" == message {
			numApplesOrFigs++
			return true
		}

		return false
	})


	router := NewFanoutRouter(appleRouter, bananaRouter, cherryRouter, figRouter, kiwiRouter, appleOrFigRouter)

	if expected, actual := 0, numApples; expected != actual {
		t.Errorf("Initially expected numApples to be %d, but actually was %d.", expected, actual)
		return
	}
	if expected, actual := 0, numBananas; expected != actual {
		t.Errorf("Initially expected numBananas to be %d, but actually was %d.", expected, actual)
		return
	}
	if expected, actual := 0, numCherries; expected != actual {
		t.Errorf("Initially expected numCherries to be %d, but actually was %d.", expected, actual)
		return
	}
	if expected, actual := 0, numFigs; expected != actual {
		t.Errorf("Initially expected numFigs to be %d, but actually was %d.", expected, actual)
		return
	}
	if expected, actual := 0, numKiwis; expected != actual {
		t.Errorf("Initially expected numKiwis to be %d, but actually was %d.", expected, actual)
		return
	}
	if expected, actual := 0, numApplesOrFigs; expected != actual {
		t.Errorf("Initially expected numApplesOrFigs to be %d, but actually was %d.", expected, actual)
		return
	}



	var message string
	var context map[string]interface{} = map[string]interface{}{}



	message = "apple"
	if err := router.Route(message, context); nil != err {
		switch errs := err.(type) {
		case manyerrors.Errors:
			t.Errorf("Received many error when trying to send message %q: %#v", message, errs.Errors())
		default:
			t.Errorf("Received error when trying to send message %q: %v", message, err)
		}
	}

	if expected, actual := 1, numApples; expected != actual {
		t.Errorf("After sending message %q expected numApples to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numBananas; expected != actual {
		t.Errorf("After sending message %q expected numBananas to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numCherries; expected != actual {
		t.Errorf("After sending message %q expected numCherries to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numFigs; expected != actual {
		t.Errorf("After sending message %q expected numFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numKiwis; expected != actual {
		t.Errorf("After sending message %q expected numKiwis to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numApplesOrFigs; expected != actual {
		t.Errorf("After sending message %q expected numApplesOrFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}



	message = "banana"
	if err := router.Route(message, context); nil != err {
		switch errs := err.(type) {
		case manyerrors.Errors:
			t.Errorf("Received many error when trying to send message %q: %#v", message, errs.Errors())
		default:
			t.Errorf("Received error when trying to send message %q: %v", message, err)
		}
	}

	if expected, actual := 1, numApples; expected != actual {
		t.Errorf("After sending message %q expected numApples to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numBananas; expected != actual {
		t.Errorf("After sending message %q expected numBananas to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numCherries; expected != actual {
		t.Errorf("After sending message %q expected numCherries to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numFigs; expected != actual {
		t.Errorf("After sending message %q expected numFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numKiwis; expected != actual {
		t.Errorf("After sending message %q expected numKiwis to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numApplesOrFigs; expected != actual {
		t.Errorf("After sending message %q expected numApplesOrFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}



	message = "cherry"
	if err := router.Route(message, context); nil != err {
		switch errs := err.(type) {
		case manyerrors.Errors:
			t.Errorf("Received many error when trying to send message %q: %#v", message, errs.Errors())
		default:
			t.Errorf("Received error when trying to send message %q: %v", message, err)
		}
	}

	if expected, actual := 1, numApples; expected != actual {
		t.Errorf("After sending message %q expected numApples to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numBananas; expected != actual {
		t.Errorf("After sending message %q expected numBananas to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numCherries; expected != actual {
		t.Errorf("After sending message %q expected numCherries to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numFigs; expected != actual {
		t.Errorf("After sending message %q expected numFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numKiwis; expected != actual {
		t.Errorf("After sending message %q expected numKiwis to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numApplesOrFigs; expected != actual {
		t.Errorf("After sending message %q expected numApplesOrFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}



	message = "fig"
	if err := router.Route(message, context); nil != err {
		switch errs := err.(type) {
		case manyerrors.Errors:
			t.Errorf("Received many error when trying to send message %q: %#v", message, errs.Errors())
		default:
			t.Errorf("Received error when trying to send message %q: %v", message, err)
		}
	}

	if expected, actual := 1, numApples; expected != actual {
		t.Errorf("After sending message %q expected numApples to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numBananas; expected != actual {
		t.Errorf("After sending message %q expected numBananas to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numCherries; expected != actual {
		t.Errorf("After sending message %q expected numCherries to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numFigs; expected != actual {
		t.Errorf("After sending message %q expected numFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 0, numKiwis; expected != actual {
		t.Errorf("After sending message %q expected numKiwis to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 2, numApplesOrFigs; expected != actual {
		t.Errorf("After sending message %q expected numApplesOrFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}



	message = "kiwi"
	if err := router.Route(message, context); nil != err {
		switch errs := err.(type) {
		case manyerrors.Errors:
			t.Errorf("Received many error when trying to send message %q: %#v", message, errs.Errors())
		default:
			t.Errorf("Received error when trying to send message %q: %v", message, err)
		}
	}

	if expected, actual := 1, numApples; expected != actual {
		t.Errorf("After sending message %q expected numApples to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numBananas; expected != actual {
		t.Errorf("After sending message %q expected numBananas to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numCherries; expected != actual {
		t.Errorf("After sending message %q expected numCherries to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numFigs; expected != actual {
		t.Errorf("After sending message %q expected numFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numKiwis; expected != actual {
		t.Errorf("After sending message %q expected numKiwis to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 2, numApplesOrFigs; expected != actual {
		t.Errorf("After sending message %q expected numApplesOrFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}



	message = "fig"
	if err := router.Route(message, context); nil != err {
		switch errs := err.(type) {
		case manyerrors.Errors:
			t.Errorf("Received many error when trying to send message %q: %#v", message, errs.Errors())
		default:
			t.Errorf("Received error when trying to send message %q: %v", message, err)
		}
	}

	if expected, actual := 1, numApples; expected != actual {
		t.Errorf("After sending message %q expected numApples to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numBananas; expected != actual {
		t.Errorf("After sending message %q expected numBananas to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numCherries; expected != actual {
		t.Errorf("After sending message %q expected numCherries to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 2, numFigs; expected != actual {
		t.Errorf("After sending message %q expected numFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 1, numKiwis; expected != actual {
		t.Errorf("After sending message %q expected numKiwis to be %d, but actually was %d.", message, expected, actual)
		return
	}
	if expected, actual := 3, numApplesOrFigs; expected != actual {
		t.Errorf("After sending message %q expected numApplesOrFigs to be %d, but actually was %d.", message, expected, actual)
		return
	}
}
