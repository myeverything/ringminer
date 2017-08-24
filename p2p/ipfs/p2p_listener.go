package ipfs

import (
	"github.com/ipfs/go-ipfs-api"
	"log"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/orderbook"
)

/**
p2p Listener
 */

const TOPIC = "topic"

type IPFSListener struct {
	sh *shell.Shell
	sub *shell.PubSubSubscription
	quit chan bool
}

func NewListener() *IPFSListener{
	sh := shell.NewLocalShell()
	sub, err := sh.PubSubSubscribe(TOPIC)
	ch := make(chan bool, 1)
	if err != nil {
		log.Fatal("IPFS\t-", "listener start sub failed:", err.Error())
	}
	return &IPFSListener{sh, sub, ch}
}

// TODO(fukun): add go func external
func (listener *IPFSListener) Start() {
	listener.quit <- true
	for {
		record, _ := listener.sub.Next()
		data := record.Data()
		var ord types.Order
		err := ord.UnMarshalJson(data)
		if err != nil {
			log.Println("p2p listener\t-", "Listen data error:", err.Error())
		} else {
			log.Println("p2p listener\t-", "Listen data success:", string(data))
		}
		orderbook.NewOrder(ord)
	}
}

func (listener *IPFSListener) Stop() {
	<- listener.quit
	defer close(listener.quit)
}

func (listener *IPFSListener) Name() string {
	return "ipfs-listener"
}