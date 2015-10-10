package flog


// NewFilteredRouter returns an initialized FilteredRouter.
//
// 'subrouter' is the sub-router that a FilteredRouter will
// re-Route a 'message' (and 'context') to, but only on the
// condition that 'filterFn' returns 'true' for the 'message'
// and 'context' passed to it.
//
// An example for 'filterFn' is the following.
//
//	func filterError(message string, context map[string]interface{})bool) bool {
//		if datum, ok := context["error"]; !ok {
//			return false
//		} else if _, ok := datum.(error); !ok {
//			return false
//		} else {
//			return true
//		}
//	}
//
// This func will make it so only re-route messages whose context #1 has the key "error"
// and #2 the value of the context at key "key" fits the builtin Go 'error' interface.
//
// Also, a rather useless example, but a 'filterFn' that would reject all messages (and
// contexts) is:
//
//	func filterRejectAll(message string, context map[string]interface{})bool) bool {
//		return false
//	}
//
//
// And also, another rather useless example, but a 'filterFn' that would allow all messages
// (and contexts) is:
//
//	func filterAcceptAll(message string, context map[string]interface{})bool) bool {
//		return true
//	}
//
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
