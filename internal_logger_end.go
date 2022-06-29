package log

import (
	"fmt"
	"time"
)

func (receiver internalLogger) End(a ...interface{}) {
	diff := time.Now().Sub(receiver.begin)

	msg := fmt.Sprintf(" δt=%s", diff)

	a = append([]interface{}{"END "}, a...)
	a = append(a, msg)

	receiver.Debug(a...)
}
