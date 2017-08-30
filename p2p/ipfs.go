package p2p

import (
	"github.com/ipfs/go-ipfs-api"
	"sync"
)

/**
p2p Listener
 */

const TOPIC = "topic"

type IPFSListener struct {
	sh *shell.Shell
	sub *shell.PubSubSubscription
	stop chan struct{}
	lock sync.RWMutex
}

func NewListener() *IPFSListener {
	l := &IPFSListener{}

	l.sh = shell.NewLocalShell()
	sub, err := l.sh.PubSubSubscribe(TOPIC)
	if err != nil {
		panic(err.Error())
	}
	l.sub = sub

	return l
}

func (l *IPFSListener) Start() {
	l.stop = make(chan struct{})
	go func() {
		for {
			record, _ := l.sub.Next()
			data := record.Data()
			Send(data)
		}
	}()
}

func (listener *IPFSListener) Stop() {
	listener.lock.Lock()
	close(listener.stop)
	listener.lock.Unlock()
}

func (listener *IPFSListener) Name() string {
	return "ipfs-listener"
}
