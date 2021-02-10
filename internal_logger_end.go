package flog

func (receiver internalLogger) End(a ...interface{}) {
	a = append([]interface{}{"END"}, a...)

	receiver.Debug(a...)
}
