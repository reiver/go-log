package flog


import (
	"fmt"
	"io"
	"time"
)


// PrettyWritingRouter returns an initialized PrettyWritingRouter
func NewPrettyWritingRouter(writer io.Writer) *PrettyWritingRouter {
	router := PrettyWritingRouter{
		writer:writer,
	}

	return &router
}



// PrettyWritingRouter is a router that writes a pretty version of
// the log it was give (including COLORS!) to the writer it was
// given when it was created.
//
// A PrettyWritingRouter is appropriate for a deployment (i.e., "DEV")
// deployment enviornment. (And  probably not appropriate a production
// (i.e., "PROD") deployment environment.)
type PrettyWritingRouter struct {
	writer io.Writer
}



func (router *PrettyWritingRouter) Route(message string, context map[string]interface{}) error {

	const STYLE_PANIC     = "\x1b[40;31;1m" // BG BLACK, FG RED, BOLD
	const STYLE_ERROR     = "\x1b[41;33;1m" // BG RED, FG YELLOW, BOLD
	const STYLE_WARNING   = "\x1b[43;37;1m" // BG YELLOW, FG WHITE, BOLD
	const STYLE_NOTICE    = "\x1b[42;33;1m" // BG GREEN, FG YELLOW, BOLD
	const STYLE_TIMESTAMP = "\x1b[2m"       // FAINT
	const STYLE_MESSAGE   = "\x1b[44;37;1m" // BG BLUE, FG WHITE, BOLD
	const STYLE_DEFAULT   = "\033[95m"      // HEADER
	const STYLE_RESET     = "\033[0m"       // RESET

	str := ""

	if nil != context {
		if _, ok := context["panic"]; ok {
			str = fmt.Sprintf("%s ☠ ☠ ☠ ☠ ☠ %s\t%s", STYLE_PANIC, STYLE_RESET, str)
		} else if _, ok := context["error"]; ok {
			str = fmt.Sprintf("%s 😨 😨 😨 😨 😨 %s\t%s", STYLE_ERROR, STYLE_RESET, str)
		} else if _, ok := context["warning"]; ok {
			str = fmt.Sprintf("%s 😟 😟 😟 😟 😟 %s\t%s", STYLE_WARNING, STYLE_RESET, str)
		} else if _, ok := context["notice"]; ok {
			str = fmt.Sprintf("%s 😮 😮 😮 😮 😮 %s\t%s", STYLE_NOTICE, STYLE_RESET, str)
		}
	}

	str = fmt.Sprintf("%s%s%s%s\t(%s%v%s)", str, STYLE_MESSAGE, message, STYLE_RESET, STYLE_TIMESTAMP, time.Now(), STYLE_RESET)
	for key, value := range context {
		style := STYLE_DEFAULT

		switch key {
		case "panic":
			style = STYLE_PANIC
		case "error":
			style = STYLE_ERROR
		case "warning":
			style = STYLE_WARNING
		case "notice":
			style = STYLE_NOTICE
		}

		str = fmt.Sprintf("%s\t%s%s%s=%s%#v%s", str, style, key, STYLE_RESET, style, value, STYLE_RESET)
	}

	fmt.Fprintln(router.writer, str)

//@TODO: Should this be checking for errors from fmt.Fprintln()?
	return nil
}
