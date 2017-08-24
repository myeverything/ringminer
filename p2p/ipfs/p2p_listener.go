package ipfs

import (
	"github.com/ipfs/go-ipfs-api"
	"log"
)

/**
p2p Listener
 */

type IPFSListener struct {
	sh *shell.Shell
	sub *shell.PubSubSubscription
	quit chan bool
}

func NewListener(topic string) *IPFSListener{
	sh := shell.NewLocalShell()
	sub, err := sh.PubSubSubscribe(topic)
	ch := make(chan bool, 1)
	if err != nil {
		log.Fatal("IPFS\t-", "listener start sub failed:", err.Error())
	}
	return &IPFSListener{sh, sub, ch}
}

func (listener *IPFSListener) Start() {
	listener.quit <- true
	go func() {
		for {
			record, _ := listener.sub.Next()
			peerId := record.From().String()
			data := record.Data()

			if len(peerId) > 0 {
				log.Printf("p2p listener\t- Listen peerId %s,data %s", peerId, string(data))
			}

		}
	}()
}

func (listener *IPFSListener) Stop() {
	<- listener.quit
	defer close(listener.quit)
}