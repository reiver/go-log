package flog


// NewFilteredRouter returns an initialized FilteredRouter.
//
// 'subrouter' is the sub-router that a FilteredRouter will
// re-Route a 'message' (and 'context') to, but only on the
// condition that 'filterFn' returns 'true' for the 'message'
// and 'context' passed to it.
func NewFilteredRouter(subrouter Router, filterFn func(string, map[string]interface{})bool) *FilteredRouter {
	router := FilteredRouter{
		subrouter:subrouter,
		filterFn:filterFn,
	}

	return &router
}


// FilteredRouter is a Router that conditionally re-routes or discards a message (and its context).
type FilteredRouter struct {
	subrouter Router
	filterFn func(string, map[string]interface{})bool
}


func (router *FilteredRouter) Route(message string, context map[string]interface{}) error {
	if router.filterFn(message, context) {
		return router.subrouter.Route(message, context)
	}

	return nil
}
