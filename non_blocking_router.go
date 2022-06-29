package log


// NewNonBlockingRouter returns an initialized NonBlockingRouter.
func NewNonBlockingRouter(subrouter Router) *NonBlockingRouter {
	router := NonBlockingRouter{
		subrouter:subrouter,
	}

	return &router
}


// NonBlockingRouter is a Router when its Route method is call its does not
// block and dealing with the routing in parallel.
//
// Note that this means that if the application could terminate before this
// completes.
type NonBlockingRouter struct {
	subrouter Router
}


func (router *NonBlockingRouter) Route(message string, context map[string]interface{}) error {
	if nil == router {
		return errNilReceiver
	}

	go func() {
		router.subrouter.Route(message, context)
	}()

	return nil
}
