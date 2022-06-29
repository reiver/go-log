package log


func (flogger *internalFlogger)  With(cascade ...interface{}) Flogger {

	var x    interface{} = flogger.context
	var xs []interface{} = []interface{}{x}

	newCascade := append(xs, cascade...)

	newFlogger := New(flogger.router, newCascade...)

	return newFlogger
}
