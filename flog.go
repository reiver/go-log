package flog


type internalFlogger struct {
//@TODO: Do we really want this here?
//	logs       []string
	context map[string]interface{}
	router       Router
}


// New returns an initialized Flogger.
func New(router Router, cascade ...interface{}) Flogger {
//@TODO: Do we really want this here?
//	logs    := make([]string, 0, 8)

	context := newContext(cascade...)

	flogger := internalFlogger{
//@TODO: Do we really want this here?
//		logs:logs,
		context:context,
		router:router,
	}

	return &flogger
}
