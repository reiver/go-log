# go-log

A library that provides structured and formatted logging for the Go programming language.

(This Go package was originally named `flog`, but in version 2 was renamed to `log`.)

## Online Documention

Online documentation, which includes examples, can be found at: http://godoc.org/codeberg.org/reiver/go-log

[![GoDoc](https://godoc.org/codeberg.org/reiver/go-log?status.svg)](https://godoc.org/codeberg.org/reiver/go-log)

## Basic Usage

Basic usage of this logger looks like this:

```golang
router := log.NewPrettyWritingRouter(os.Stdout)

logger := log.New(router)
```
Once you have the logger, you can do things such as:

```golang
logger.Print("Hello world!")
logger.Println("Hello world!")
logger.Printf("Hello %s!", name)

logger.Panic("Uh oh!")
logger.Panicln("Uh oh!")
logger.Panicf("Uh oh, had a problem happen: %s.", problemDescription)

logger.Fatal("Something really bad happened!")
logger.Fatalln("Something really bad happened!")
logger.Fatalf("Something really bad happened: %s.", problemDescription)
```

BTW, if the PrettyWritingRouter was being used, then this:

```golang
logger.Print("Hello world!")
```

Would generate output like the following:

```
Hello world!	(2015-10-10 17:28:49.397356044 -0700 PDT)
```

(Although note that in actual usage this would have color.)

(Note that for for other routers the actual output would look very different!
What the output looks like is router dependent.)

## Structured Logging

But those method calls all generated unstructure data.

To include structured data the logger's With method needs to be used.
For example:

```golang
newLogger := logger.With(map[string]interface{}{
	"method":"Toil",
	"secret_note":"Hi there! How are you?",
})
```

Then if the PrettyWritingRouter was being used, then this:

```golang
newLogger.Print("Hello world!")
```

Would generate output like the following:

```
Hello world!	(2015-10-10 17:28:49.397356044 -0700 PDT)	method="Toil"	secret_note="Hi there! How are you?"
```

(Again, note that in actual usage this would have color.)

## Deployment Environment

Of course in a real application system you should (probably) create a different kind
of logger for each deployment environment.

Even though the PrettyWritingRouter is great for a development deployment environment
(i.e., "DEV") it is probably not appropriate for a production deployment environment
(i.e., "PROD").

For example:
```golang
var logger log.Logger

switch deploymentEnvironment {
case "DEV":
	router := log.NewPrettyWritingRouter(os.Stdout)
	
	logger = log.New(router)
case "PROD":
	verboseRouter = log.NewDiscardingRouter()
	if isVerboseMode {
		verboseRouter = NewCustomVerboseRouter()
	}
	
	panicDetectionRouter := log.NewFilteringRouter(NewCustomerPanicRecordingRouter(), filterOnlyPanicsFunc)
	
	errorDetectionRouter := log.NewFilteringRouter(NewCustomerPanicRecordingRouter(), filterOnlyErrorsFunc)
	
	router := NewFanoutRouter(verboseRouter, panicDetectionRouter, errorDetectionRouter)
	
	logger = log.New(router)
}
```

## Import

To import package **log** use `import` code like the following:

```
import "codeberg.org/reiver/go-log"
```

## Installation

To install package **log** do the following:

```
GOPROXY=direct go get codeberg.org/reiver/go-log
```

## Author

Package **log** was written by [Charles Iliya Krempeaux](http://reiver.link)
