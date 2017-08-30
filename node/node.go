package node

import (
	"sync"
	"github.com/Loopring/ringminer/matchengine"
	"github.com/Loopring/ringminer/chainclient"
	"go.uber.org/zap"
	"github.com/Loopring/ringminer/orderbook"
	"github.com/Loopring/ringminer/p2p"
)

// TODO(fukun): should add multi service
// TODO(fukun): 考虑使用微服务框架
type Node struct {
	server *matchengine.Proxy
	p2pListener orderbook.Listener
	ethListener *chainclient.Client
	stop chan struct{}
	lock sync.RWMutex
	logger *zap.Logger
}

// TODO(fukun): modify server
func NewNode(logger *zap.Logger) *Node {
	n := &Node{}
	n.logger = logger

	n.registerP2PListener()

	return n
}

func (n *Node) Start() {
	n.p2pListener.Start()
}

func (n *Node) Wait() {
	n.lock.RLock()

	//if n.server == nil {
	//	n.lock.RUnlock()
	//	n.logger.Error("matchengine should not be empty")
	//	return
	//}

	stop := n.stop
	n.lock.RUnlock()

	<-stop
}

func (n *Node) Stop() {
	n.lock.RLock()

	close(n.stop)

	n.p2pListener.Stop()

	n.lock.RUnlock()
}

func (n *Node) registerP2PListener() {
	n.p2pListener = p2p.NewListener()
}