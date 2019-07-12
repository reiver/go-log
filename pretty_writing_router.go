package flog

import (
	"fmt"
	"io"
	"sort"
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
// A PrettyWritingRouter is appropriate for a development (i.e., "DEV")
// deployment enviornment. (And  probably not appropriate a production
// (i.e., "PROD") deployment environment.)
type PrettyWritingRouter struct {
	writer io.Writer
}

func (router *PrettyWritingRouter) Route(message string, context map[string]interface{}) error {
	if nil == router {
		return errNilReceiver
	}

	const STYLE_FATAL     = "\x1b[40;33;1m" // BG BLACK,  FG YELLOW, BOLD
	const STYLE_PANIC     = "\x1b[40;31;1m" // BG BLACK,  FG RED,    BOLD
	const STYLE_ERROR     = "\x1b[41;37;1m" // BG RED,    FG WHITE,  BOLD
	const STYLE_WARNING   = "\x1b[43;37;1m" // BG YELLOW, FG WHITE,  BOLD
	const STYLE_NOTICE    = "\x1b[42;33;1m" // BG GREEN,  FG YELLOW, BOLD
	const STYLE_TIMESTAMP = "\x1b[2m"       // FAINT
	const STYLE_MESSAGE   = "\x1b[44;37;1m" // BG BLUE,   FG WHITE,  BOLD
	const STYLE_DEFAULT   = "\033[95m"      // HEADER
	const STYLE_RESET     = "\033[0m"       // RESET

	str := ""

	if nil != context {
		if _, ok := context["~fatal"]; ok {
			str = fmt.Sprintf("%s 💀 💀 💀 💀 💀 %s\t%s", STYLE_FATAL, STYLE_RESET, str)
		} else if _, ok := context["~panic"]; ok {
			str = fmt.Sprintf("%s ☠ ☠ ☠ ☠ ☠ %s\t%s", STYLE_PANIC, STYLE_RESET, str)
		} else if _, ok := context["~error"]; ok {
			str = fmt.Sprintf("%s 😨 😨 😨 😨 😨 %s\t%s", STYLE_ERROR, STYLE_RESET, str)
		} else if _, ok := context["~warn"]; ok {
			str = fmt.Sprintf("%s 😟 😟 😟 😟 😟 %s\t%s", STYLE_WARNING, STYLE_RESET, str)
		} else if _, ok := context["~print"]; ok {
			str = fmt.Sprintf("%s 😮 😮 😮 😮 😮 %s\t%s", STYLE_NOTICE, STYLE_RESET, str)
		}
	}

	str = fmt.Sprintf("%s%s%s%s\t(%s%v%s)", str, STYLE_MESSAGE, message, STYLE_RESET, STYLE_TIMESTAMP, time.Now(), STYLE_RESET)

	// If we have an error, then get the error.Error() into the log too.
	if errorFieldValue, ok := context["~error"]; ok {
		if err, ok := errorFieldValue.(error); ok {
			context["~~error"] = fmt.Sprintf("%s, {{%T}}", err.Error(), err)
		}
	}


//@TODO: This is a potential heavy operation. Is there a better way
//       to get the ultimate result this is trying to archive?
//
	sortedKeys := make([]string, len(context))
	i := 0
	for key, _ := range context {
		sortedKeys[i] = key
		i++
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {

		value := context[key]

		style := STYLE_DEFAULT

		switch key {
                case "~fatal", "~fatals":
                        style = STYLE_FATAL
		case "~panic", "~panics":
			style = STYLE_PANIC
		case "~error", "~errors", "~~error":
			style = STYLE_ERROR
		case "~warning", "~warnings":
			style = STYLE_WARNING
		case "~print", "~prints":
			style = STYLE_NOTICE
		}

		str = fmt.Sprintf("%s\t%s%s%s=%s%#v%s", str, style, key, STYLE_RESET, style, value, STYLE_RESET)
	}

	fmt.Fprintln(router.writer, str)

//@TODO: Should this be checking for errors from fmt.Fprintln()?
	return nil
}
