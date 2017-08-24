package main

import (
	"github.com/Loopring/ringminer/p2p/ipfs"
	"time"
)

func DebugOrderBook(testcase string) {
	switch testcase{
	case "listen":
		ipfsListenTest()
	}
}

func ipfsListenTest() {
	listener := ipfs.NewListener("topic")
	listener.Start()
	time.Sleep(10 * time.Second)
	listener.Stop()
}