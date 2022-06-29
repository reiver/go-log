package log


import (
	"github.com/reiver/go-manyerrors"
)


// NewFanoutRouter returns an initialized FanoutRouter.
func NewFanoutRouter(subrouters ...Router) *FanoutRouter {
	router := FanoutRouter{
		subrouters:subrouters,
	}

	return &router
}


// FanoutRouter is a Router that re-routes any message (and its context) it
// receives to all of its sub-routers.
type FanoutRouter struct {
	subrouters []Router
}


func (router *FanoutRouter) Route(message string, context map[string]interface{}) error {
	if nil == router {
		return errNilReceiver
	}

	errors := []error{}

	for _, subrouter := range router.subrouters {
		if err := subrouter.Route(message, context); nil != err {
			errors = append(errors, err)
		}
	}

	if 0 < len(errors) {
		return manyerrors.New(errors...)
	}

	return nil
}
