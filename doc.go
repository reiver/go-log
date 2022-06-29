/*
Package flog provides structured and formatted logging.

Basic Usage

Basic usage of the flogger looks like this:

	router := flog.NewPrettyWritingRouter(os.Stdout)
	
	flogger := flog.New(router)

Once you have the flogger, you can do things such as:

	flogger.Print("Hello world!")
	flogger.Println("Hello world!")
	flogger.Printf("Hello %s!", name)
	
	flogger.Panic("Uh oh!")
	flogger.Panicln("Uh oh!")
	flogger.Panicf("Uh oh, had a problem happen: %s.", problemDescription)
	
	flogger.Fatal("Something really bad happened!")
	flogger.Fatalln("Something really bad happened!")
	flogger.Fatalf("Something really bad happened: %s.", problemDescription)

BTW, if the PrettyWritingRouter was being used, then this:

	flogger.Print("Hello world!")

Would generate output like the following:

	Hello world!	(2015-10-10 17:28:49.397356044 -0700 PDT)

(Although note that in actual usage this would have color.)

(Note that for for other routers the actual output would look very different!
What the output looks like is router dependent.)

Structured Logging

But those method calls all generated unstructure data.

To include structured data the flogger's With method needs to be used.
For example:

	newFlogger := flogger.With(map[string]interface{}{
		"method":"Toil",
		"secret_note":"Hi there! How are you?",
	})

Then if the PrettyWritingRouter was being used, then this:

	newFlogger.Print("Hello world!")

Would generate output like the following:

	Hello world!	(2015-10-10 17:28:49.397356044 -0700 PDT)	method="Toil"	secret_note="Hi there! How are you?"

(Again, note that in actual usage this would have color.)

Deployment Environment

Of course in a real application system you should (probably) create a different kind
of flogger for each deployment environment.

Even though the PrettyWritingRouter is great for a development deployment environment
(i.e., "DEV") it is probably not appropriate for a production deployment environment
(i.e., "PROD").

For example:

	var flogger flog.Flogger
	
	switch deploymentEnvironment {
	case "DEV":
		router := flog.NewPrettyWritingRouter(os.Stdout)
		
		flogger = flog.New(router)
	case "PROD":
		verboseRouter = flog.NewDiscardingRouter()
		if isVerboseMode {
			verboseRouter = NewCustomVerboseRouter()
		}

		panicDetectionRouter := flog.NewFilteringRouter(NewCustomerPanicRecordingRouter(), filterOnlyPanicsFunc)

		errorDetectionRouter := flog.NewFilteringRouter(NewCustomerPanicRecordingRouter(), filterOnlyErrorsFunc)

		router := NewFanoutRouter(verboseRouter, panicDetectionRouter, errorDetectionRouter)
		
		flogger = flog.New(router)
	}

More Routers

In addition to the builtin routers that go-flog comes with, other external
routers are also available. These include:

go-slackchannelrouter:
Makes it so log messages get posted to a Slack channel.
https://github.com/reiver/go-slackchannelrouter

*/
package log
