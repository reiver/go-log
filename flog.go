package flog


type internalFlogger struct {
	context map[string]interface{}
	router       Router
}


// New returns an initialized Flogger.
func New(router Router, cascade ...interface{}) Flogger {
	context := newContext(cascade...)

	flogger := internalFlogger{
		context:context,
		router:router,
	}

	return &flogger
}



func (flogger *internalFlogger) route(message string, moreContext map[string]interface{}) error {

	context := newContext(flogger.context, moreContext)

	return flogger.router.Route(message, context)
}
