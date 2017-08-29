package ipfs

import (
	"github.com/ipfs/go-ipfs-api"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/orderbook"
	"sync"
	"github.com/Loopring/ringminer/log"
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
			var ord types.Order
			err := ord.UnMarshalJson(data)
			if err != nil {
				log.Error(log.ERROR_P2P_LISTEN_ACCEPT, "content", "")
			} else {
				log.Info(log.LOG_P2P_ACCEPT, "data", string(data))
			}
			orderbook.NewOrder(ord)
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