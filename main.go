package main

import (
	"flag"
)

var (
	auth = flag.String("auth", "dylenfu", "chose developer")
	testcase = flag.String("testcase", "listen", "chose test case")
)

func main() {
	switch *auth {
	case "dylenfu":
		DebugOrderBook(*testcase)

	case "hongyu":
		DebugMatch()
	}
}