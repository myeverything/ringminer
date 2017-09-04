package p2p

import (
	"github.com/ipfs/go-ipfs-api"
	"sync"
	"github.com/Loopring/ringminer/types"
)

/**
p2p Listener
 */

// TODO(fukun): 后面需要修改该topic
const TOPIC = "topic"

type IPFSListener struct {
	sh *shell.Shell
	sub *shell.PubSubSubscription
	stop chan struct{}
	whisper *types.Whispers
	lock sync.RWMutex
}

func NewListener(whisper *types.Whispers) *IPFSListener {
	l := &IPFSListener{}

	l.sh = shell.NewLocalShell()
	sub, err := l.sh.PubSubSubscribe(TOPIC)
	if err != nil {
		panic(err.Error())
	}
	l.sub = sub
	l.whisper = whisper

	return l
}

func (l *IPFSListener) Start() {
	l.stop = make(chan struct{})
	go func() {
		for {
			record, _ := l.sub.Next()
			data := record.Data()
			ord := GenOrder(data)
			l.whisper.PeerOrderChan <- ord
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
