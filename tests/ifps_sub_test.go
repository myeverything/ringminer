package tests

import (
	"time"
	"github.com/Loopring/ringminer/p2p/ipfs"
	"testing"
)

func TestIpfsListener(t *testing.T) {
	listener := ipfs.NewListener()
	go listener.Start()
	time.Sleep(100 * time.Second)
	listener.Stop()
}