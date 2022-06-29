package log


// NewMappingRouter returns an initialized MappingRouter.
func NewMappingRouter(subrouter Router, fn func(string, map[string]interface{})(string, map[string]interface{})) *MappingRouter {
	router := MappingRouter{
		subrouter:subrouter,
		fn:fn,
	}

	return &router
}


// MappingRouter is a Router that can modify the message and context before
// re-routing it to its sub-router.
//
// Conceptually this is somewhat similar to "map" functions in functional
// programming.
type MappingRouter struct {
	subrouter Router
	fn func(string, map[string]interface{})(string, map[string]interface{})
}


func (router *MappingRouter) Route(message string, context map[string]interface{}) error {
	if nil == router {
		return errNilReceiver
	}

	return router.subrouter.Route(router.fn(message, context))
}
