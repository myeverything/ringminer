package node

import (
	"sync"
	"github.com/Loopring/ringminer/matchengine"
	"go.uber.org/zap"
	"github.com/Loopring/ringminer/orderbook"
	"github.com/Loopring/ringminer/p2p"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/config"
	"github.com/Loopring/ringminer/chainclient/eth"
)

// TODO(fukun): should add multi service
// TODO(fukun): 考虑使用微服务框架
type Node struct {
	options *config.GlobalConfig
	server *matchengine.Proxy
	p2pListener p2p.Listener
	ethListener eth.Listener
	orderbook *orderbook.OrderBook
	whisper *types.Whispers
	stop chan struct{}
	lock sync.RWMutex
	logger *zap.Logger
}

// TODO(fukun): modify server
func NewNode(logger *zap.Logger) *Node {
	n := &Node{}

	whisper := &types.Whispers{}
	whisper.PeerOrderChan = make(chan *types.Order)
	whisper.ChainOrderChan = make(chan *types.OrderMined)
	whisper.EngineOrderChan = make(chan *types.OrderState)

	n.whisper = whisper
	n.logger = logger
	n.options = config.LoadConfig()

	n.registerP2PListener()
	n.registerOrderBook()

	return n
}

func (n *Node) Start() {
	n.p2pListener.Start()
	n.orderbook.Start()

	// TODO(fukun): 放开eth监听
	//n.ethListener.Start()
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

	n.p2pListener.Stop()
	n.ethListener.Stop()
	close(n.stop)

	n.lock.RUnlock()
}

func (n *Node) registerP2PListener() {
	n.p2pListener = p2p.NewListener(n.whisper, n.options.Ipfs)
}

func (n *Node) registerOrderBook() {
	n.orderbook = orderbook.NewOrderBook(n.whisper, n.options.Database)
}

func (n *Node) registerEthClient() {
	n.ethListener = eth.NewListener(n.whisper, n.options.EthClient)
}