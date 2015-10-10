package flog


import (
	"testing"

	"fmt"
	"math/rand"
	"time"
)


func TestNewDiscardingRouter(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))


	router := NewDiscardingRouter()
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
