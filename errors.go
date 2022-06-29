package log


import (
	"errors"
)


var (
	errNilReceiver = errors.New("Nil Receiver")
	errNilRouter   = errors.New("Nil Router")
)
