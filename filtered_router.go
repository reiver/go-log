package flog


func NewFilteredRouter(subrouter Router, filterFn func(string, map[string]interface{})bool) *FilteredRouter {
	router := FilteredRouter{
		subrouter:subrouter,
		filterFn:filterFn,
	}

	return &router
}


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
