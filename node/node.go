/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

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
	"github.com/Loopring/ringminer/matchengine/bucket"
)

// TODO(fk): add services
type Node struct {
	options 				*config.GlobalConfig
	p2pListener 			p2p.Listener
	ethListener 			eth.Listener
	orderbook 				*orderbook.OrderBook
	matchengine             matchengine.Proxy
	peerOrderChan			chan *types.Order
	chainOrderChan			chan *types.OrderMined
	engineOrderChan			chan *types.OrderState
	stop 					chan struct{}
	lock 					sync.RWMutex
	logger 					*zap.Logger
}

// TODO(fk): inject whisper and logger
func NewNode(logger *zap.Logger) *Node {
	n := &Node{}

	n.peerOrderChan = make(chan *types.Order)
	n.chainOrderChan = make(chan *types.OrderMined)
	n.engineOrderChan = make(chan *types.OrderState)
	n.logger = logger
	n.options = config.LoadConfig()

	n.registerP2PListener()
	n.registerOrderBook()

	return n
}

func (n *Node) Start() {
	n.p2pListener.Start()
	n.orderbook.Start()
	n.matchengine.Start()

	// TODO(fk): start eth client
	//n.ethListener.Start()
}

func (n *Node) Wait() {
	n.lock.RLock()

	// TODO(fk): states should be judged

	stop := n.stop
	n.lock.RUnlock()

	<-stop
}

func (n *Node) Stop() {
	n.lock.RLock()

	n.p2pListener.Stop()
	n.ethListener.Stop()
	n.orderbook.Stop()
	n.matchengine.Stop()

	close(n.stop)

	n.lock.RUnlock()
}

func (n *Node) registerP2PListener() {
	whisper := &p2p.Whisper{n.peerOrderChan}
	n.p2pListener = p2p.NewListener(n.options.Ipfs, whisper)
}

func (n *Node) registerOrderBook() {
	whisper := &orderbook.Whisper{n.peerOrderChan, n.engineOrderChan, n.chainOrderChan}
	n.orderbook = orderbook.NewOrderBook(n.options.Database, whisper)
}

func (n *Node) registerEthClient() {
	whisper := &eth.Whisper{n.chainOrderChan}
	n.ethListener = eth.NewListener(n.options.EthClient, whisper)
}

func (n *Node) registerProxy() {
	whisper := &bucket.Whisper{n.engineOrderChan}
	n.matchengine = bucket.NewBucketProxy(n.options.BucketProxy, whisper)
}