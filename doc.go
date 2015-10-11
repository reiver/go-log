/*
Package flog provides structured and formatted logging.

Basic Usage

Basic usage of the flogger looks like this:

	router := NewPrettyWritingRouter(os.Stdout)
	
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


Deployment Environment

Of course in a real application system you should (probably) create a different kind
of flogger for each deployment environment.

For example:

	var flogger flog.Flogger
	
	switch deploymentEnvironment {
	case "DEV":
		router := NewPrettyWritingRouter(os.Stdout)
		
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


*/
package flog
