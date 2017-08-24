package main

import (
	"flag"
)

var (
	auth = flag.String("auth", "dylenfu", "chose developer")
	testcase = flag.String("testcase", "listen", "chose test case")
)

func main() {
	flag.Parse()

	switch *auth {
	case "dylenfu":
		DebugOrderBook(*testcase)

	case "hongyu":
		DebugMatch()
	}
}