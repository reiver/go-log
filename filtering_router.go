package flog


// NewFilteringRouter returns an initialized FilteringRouter.
//
// 'subrouter' is the sub-router that a FilteringRouter will
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
// This func will cause the router it only re-route messages whose context #1 has the
// key "error" and #2 where the value of the context at key "key" fits the builtin Go
// 'error' interface.
//
// So, for example, for 'filterError' this would pass:
//
//	context := map[string]interface{}{
//		"apple":1,
//		"banana":2,
//		"cherry":3,
//		"error": errors.New("Something bad happened :-("),
//	}
//
// But, again for 'filterError', this would NOT pass:
//
//	context := map[string]interface{}{
//		"apple":1,
//		"banana":2,
//		"cherry":3,
//	}
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
func NewFilteringRouter(subrouter Router, filterFn func(string, map[string]interface{})bool) *FilteringRouter {
	router := FilteringRouter{
		subrouter:subrouter,
		filterFn:filterFn,
	}

	return &router
}


// FilteringRouter is a Router that conditionally re-routes or discards a message (and its context).
type FilteringRouter struct {
	subrouter Router
	filterFn func(string, map[string]interface{})bool
}


func (router *FilteringRouter) Route(message string, context map[string]interface{}) error {
	if router.filterFn(message, context) {
		return router.subrouter.Route(message, context)
	}

	return nil
}
