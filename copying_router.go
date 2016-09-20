package flog


// NewCopyingRouter returns an initialized CopyingRouter.
func NewCopyingRouter(subrouter Router) *CopyingRouter {
	copies  := make([]struct{Message string ; Context map[string]interface{}}, 0, 8)

	router := CopyingRouter {
		subrouter:subrouter,
		copies:copies,
	}

	return &router
}


// CopyingRouter is a Router that copies a message (and its context) to memory, and then
// re-routes a message (and its context) to a sub-router.
//
// This router is NOT designed to be used for an indefinite number of routings, and instead
// should only be used for a limited amount of routings, as it stores all the copies in
// memory.
type CopyingRouter struct {
	subrouter Router
	copies []struct{Message string ; Context map[string]interface{}}
}


func (router *CopyingRouter) Route(message string, context map[string]interface{}) error {
	if nil == router {
		return errNilReceiver
	}

	copy := struct{Message string ; Context map[string]interface{}}{
		Message: message,
		Context: context,
	}
	router.copies = append(router.copies, copy)

	return router.subrouter.Route(message, context)
}


func (router *CopyingRouter) Copies() <-chan struct{Message string ; Context map[string]interface{}} {
	ch := make(chan struct{Message string ; Context map[string]interface{}})

	go func() {
		for _, copy := range router.copies {
			ch <- copy
		}
		close(ch)
	}()

	return ch
}
