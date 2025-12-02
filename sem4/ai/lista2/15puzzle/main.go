package main

import "flag"

var runTests int

func init() {
	flag.IntVar(&runTests, "test", 0, "number of tests to run. Runs presentation on values lower than 2 and when omitted")
	flag.Parse()
}

func main() {
	if runTests < 2 {
		Presentation()
	} else {
		RunTests(4, runTests)
	}
}
