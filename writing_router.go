package flog


import (
	"fmt"
	"io"
	"time"
)


func NewWritingRouter(writer io.Writer) *WritingRouter {
	router := WritingRouter{
		writer:writer,
	}

	return &router
}



type WritingRouter struct {
	writer io.Writer
}



func (router *WritingRouter) Route(message string, context map[string]interface{}) error {
	const BOLD      = "\033[1m"
	const HEADER    = "\033[95m"
	const UNDERLINE = "\033[4m"
	const ENDC      = "\033[0m"

	str := fmt.Sprintf("%s%s%s\t(%s%v%s)", UNDERLINE, message, ENDC, HEADER, time.Now(), ENDC)
	for key, value := range context {
		str = fmt.Sprintf("%s\t%s%s%s=%q", str, HEADER, key, ENDC, value)
	}

	fmt.Fprintln(router.writer, str)

//@TODO: Should this be checking for errors from fmt.Fprintln()?
	return nil
}
