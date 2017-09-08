package p2p

import (
	"github.com/ipfs/go-ipfs-api"
	"sync"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/config"
)

type IpfsConfig struct {
	topic string
}

type IPFSListener struct {
	conf IpfsConfig
	toml config.IpfsOptions
	sh *shell.Shell
	sub *shell.PubSubSubscription
	stop chan struct{}
	whisper *types.Whispers
	lock sync.RWMutex
}

func (l *IPFSListener) loadConfig() {
	l.conf.topic = l.toml.Topic
}

func NewListener(whisper *types.Whispers, options config.IpfsOptions) *IPFSListener {
	l := &IPFSListener{}

	l.toml = options
	l.loadConfig()

	l.sh = shell.NewLocalShell()
	sub, err := l.sh.PubSubSubscribe(l.conf.topic)
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
