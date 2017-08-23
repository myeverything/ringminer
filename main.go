package main

import (
	"flag"
)

var (
	auth = flag.String("auth", "dylenfu", "chose developer")
)

func main() {
	switch *auth {
	case "dylenfu":
		DebugOrderBook()

	case "hongyu":
		DebugMatch()
	}
}