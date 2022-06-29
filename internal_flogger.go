package log


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
	if nil == flogger {
		return errNilReceiver
	}

	context := newContext(flogger.context, moreContext)

	router := flogger.router
	if nil == router {
		return errNilRouter
	}

	return router.Route(message, context)
}
