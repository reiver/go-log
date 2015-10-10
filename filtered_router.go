package flog


func NewFilteredRouter() *FilteredRouter {
	registry := make([]struct{FilterFn func(string, map[string]interface{})bool ; Subrouter Router}, 0, 2)

	router := FilteredRouter{
		registry:registry,
	}

	return &router
}


type FilteredRouter struct {
	registry []struct{FilterFn func(string, map[string]interface{})bool ; Subrouter Router}
}


func (router *FilteredRouter) Route(message string, context map[string]interface{}) error {
	for _, datum := range router.registry {
		if datum.FilterFn(message, context) {
			return datum.Subrouter.Route(message, context)
		}
	}

	return nil
}


func (router *FilteredRouter) Register(subrouter Router, filterFn func(string, map[string]interface{})bool) {
	datum := struct{FilterFn func(string, map[string]interface{})bool ; Subrouter Router}{
		FilterFn: filterFn,
		Subrouter: subrouter,
	}

	router.registry = append(router.registry, datum)
}
