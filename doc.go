/*
Package flog provides structured and formatted logging.

Basic Usage

Basic usage of the flogger looks like this:

	router := NewPrettyWritingRouter(os.Stdout)
	
	flogger := flog.New(router)

Once you have this, you can do things such as:

	flogger.Print("Hello world!")
	flogger.Println("Hello world!")
	flogger.Printf("Hello %s!", name)
	
	flogger.Panic("Uh oh!")
	flogger.Panicln("Uh oh!")
	flogger.Panicf("Uh oh, had a problem happen: %s.", problemDescription)
	
	flogger.Fatal("Something really bad happened!")
	flogger.Fatalln("Something really bad happened!")
	flogger.Fatalf("Something really bad happened: %s.", problemDescription)


*/
package flog
