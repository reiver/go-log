package flog


import (
	"testing"
)


func TestNewDiscardRouter(t *testing.T) {

	router := NewDiscardRouter()
	if nil == router {
		t.Errorf("After trying to create a discard router, expected it to be not nil, but was: %v", router)
	}
}
