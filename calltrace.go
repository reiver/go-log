package log

import (
	"fmt"
	"runtime"
	"strings"
)

func calltrace() []string {

	const max = 10

	var frames *runtime.Frames
	{
		var buffer [max]uintptr

		var pc []uintptr = buffer[:]

		n := runtime.Callers(0, pc)
		if 0 >= n {
			return nil
		}
		pc = pc[:n]
		frames = runtime.CallersFrames(pc)
	}

	var trace []string
	{
		var buffer [max]string
		trace = buffer[:0]

		var more bool = true
		for more {
			var frame runtime.Frame

			frame, more = frames.Next()
			switch frame.Function {
			case "runtime.Callers":
				continue
			default:
				if strings.HasPrefix(frame.Function, "github.com/reiver/go-flog.") {
					continue
				}
			}

			var filename string = frame.File
			if index := strings.LastIndex(filename, "/"); 0 <= index {
				filename = filename[1+index:]
			}

			s := fmt.Sprintf("%s():%s:%d", frame.Function, filename, frame.Line)

			trace = append(trace, s)

			if !more {
				break
			}
		}
	}

	return trace
}
