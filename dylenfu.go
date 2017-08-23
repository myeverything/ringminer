package main

import (
	"github.com/Loopring/ringminer/p2p/ipfs"
)

func DebugOrderBook() {
	listener := ipfs.NewListener("topic")
	listener.Start()
}
