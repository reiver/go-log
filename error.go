package flog


import (
	"fmt"
)


var (
	errorContext = map[string]interface{}{
		"~type":"error",
	}
)


func (flogger *internalFlogger) Error(v ...interface{}) {

	msg := fmt.Sprint(v...)

	errCtx := newErrorContext(errorContext, v...)

	flogger.route(msg, errCtx)
}

func (flogger *internalFlogger) Errorf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	errCtx := newErrorContext(errorContext, v...)

	flogger.route(msg, errCtx)
}

func (flogger *internalFlogger) Errorfe(err error, format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)

	context := map[string]interface{}{}
	for k,v := range errorContext {
		context[k] = v
	}
	context["~error"] = err

	flogger.route(msg, context)
}

func (flogger *internalFlogger) Errorln(v ...interface{}) {

	msg := fmt.Sprintln(v...)

	errCtx := newErrorContext(errorContext, v...)

	flogger.route(msg, errCtx)
}


func newErrorContext(baseContext map[string]interface{}, v ...interface{}) map[string]interface{} {


	// Collect any errors.
	errs := []error{}
	for _, datum := range v {
		if err, ok := datum.(error); ok {
			errs = append(errs, err)
		}
	}


	if 0 == len(errs) {
		return baseContext
	}


	// Copy the base context.
	context := map[string]interface{}{}
	for k,v := range baseContext {
		context[k] = v
	}


	// Put the collected errors in this new context.
	context["~errors"] = errs


	return context
}
