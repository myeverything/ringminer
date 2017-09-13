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

package bucket

import (
	"sync"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/matchengine"
	"github.com/Loopring/ringminer/chainclient"
	"github.com/Loopring/ringminer/config"
)

/**
暂时不处理以下情况
todo：此时环路的撮合驱动是由新订单的到来进行驱动，但是新订单并不是一直到达的，因此为了不浪费计算量以及增加匹配度，在没有新订单到达时，需要进行下一个长度的匹配
 bucket在解决限定长度的，新订单的快速匹配较好，但是在订单不频繁时，如何挖掘现有的订单进行处理？
 如何进行新的匹配
 首先是需要跨bucket的，进行整合的
 bucket中的更改如何反映到现有的,如何进行semiRing的遍历
 需要一个pool，对bucket进行抽象，由realtime调用pool接口，进行实时计算
 可能需要使用归约订单的结构
 */

/**
思路：设计符合要求的数据格式，
负责协调各个bucket，将ring发送到区块链，
该处负责接受neworder, cancleorder等事件，并把事件广播给所有的bucket，同时调用client将已形成的环路发送至区块链，发送时需要再次查询订单的最新状态，保证无错，一旦出错需要更改ring的各种数据，如交易量、费用分成等
 */

var loopring *chainclient.Loopring

// TODO(fukun): modify config
type BucketProxyConfig struct {
	Num int
}

type Whisper struct {
	EngineOrderChan chan *types.OrderState
}

type BucketProxy struct {
	ringChan       chan *types.RingState
	//orderStateChan chan *types.Order
	Buckets        map[types.Address]Bucket
	ringClient	*matchengine.RingClient
	mtx            *sync.RWMutex
	config   *BucketProxyConfig
	opts     config.BucketProxyOptions
 	ringChan chan *types.RingState
	OrderChan chan *types.Order
	whisper  *Whisper
	Buckets  map[types.Address]Bucket
	mtx  *sync.RWMutex
}

func NewBucketProxy(ringClient *matchengine.RingClient) matchengine.Proxy {
// TODO(fukun): add configs options
func (bp *BucketProxy) loadConfig() {

}

func NewBucketProxy(opts config.BucketProxyOptions, whisper *Whisper) matchengine.Proxy {
	var proxy matchengine.Proxy
	bp := &BucketProxy{}

	bp.opts = opts
	bp.loadConfig()

	ringChan := make(chan *types.RingState, 1000)
	bp.ringChan = ringChan

	//orderChan := make(chan *types.Order)
	//bp.OrderChan = orderChan

	bp.whisper = whisper

	bp.mtx = &sync.RWMutex{}
	bp.Buckets = make(map[types.Address]Bucket)
	bp.ringClient = ringClient
	proxy = bp
	return proxy
}

func (bp *BucketProxy) Start() {
	//orderstatechan and ringchan
	go bp.listenOrderState()
	for {
		select {
		case orderRing := <- bp.ringChan:
			s := ""
			for _,o := range orderRing.RawRing.Orders {
				s = s + " -> " + " {outtoken:" + string(o.OrderState.RawOrder.TokenS.Bytes()) + " fillamountS:" + o.FillAmountS.RealValue().String() + ", intoken:" + string(o.OrderState.RawOrder.TokenB.Bytes()) + "}"
			}
			println("ringChan receive:" + string(orderRing.Hash.Bytes()) + " ring is :" + s)

			bp.ringClient.NewRing(orderRing)
			for _, b := range bp.Buckets {
				b.NewRing(orderRing)
			}
		}
	}
}

func (bp *BucketProxy) Stop() {
	close(bp.ringChan)
	close(matchengine.OrderStateChan)
	for _,bucket := range bp.Buckets {
		bucket.Stop()
	}
}

func (bp *BucketProxy) listenOrderState() {
	for {
		select {
		case order := <- matchengine.OrderStateChan:
			if (types.ORDER_NEW == order.Status) {
				bp.newOrder(order)
			} else if (types.ORDER_CANCEL == order.Status || types.ORDER_FINISHED == order.Status) {
				bp.deleteOrder(order)
			}
		}
	}
}

func (bp *BucketProxy) newOrder(order *types.OrderState) {
	bp.mtx.RLock()
	defer bp.mtx.RUnlock()
	//如果没有则，新建bucket, todo:需要将其他bucket中的导入到当前bucket
	if _,ok := bp.Buckets[order.RawOrder.TokenS] ; !ok {
		bucket := NewBucket(order.RawOrder.TokenS, bp.ringChan)
		bp.Buckets[order.RawOrder.TokenS] = *bucket
	}
	if _,ok := bp.Buckets[order.RawOrder.TokenB] ; !ok {
		bucket := NewBucket(order.RawOrder.TokenB, bp.ringChan)
		bp.Buckets[order.RawOrder.TokenB] = *bucket
	}

	for _, b := range bp.Buckets {
		b.newOrder(*order)
	}
}

func (bp *BucketProxy) deleteOrder(order *types.OrderState) {
	for _, bucket := range bp.Buckets {
		bucket.deleteOrder(*order)
	}
} //订单的更新

func (bp *BucketProxy) AddFilter() {

}


//todo:提交ring的具体工作放在ringclient中
///**
//提交ring
////todo:用户的金额等是否需要缓存
//1、首先检查订单的状态, 重新计算成交量
//2、再提交hash
//3、hash打到块之后，再提交ring
// */
//func (bp *BucketProxy) submitRingFingerprint(ring *types.RingState) {
//	//根据最小容量，重新设置，重新计算费用
//	matchengine.ComputeRing(ring)
//	//todo:再次判断是否需要提交
//	if (!bp.canSubmit(ring)) {
//		bp.submitFailed(ring)
//	} else {
//		//todo:提交ring
//		//提交凭证，之后，等待凭证成功的event，然后提交ring，待提交的ring需要保存
//		//fingerContractAddress := &types.Address{}
//		//loopring.LoopringFingerprints[*fingerContractAddress].SubmitRingFingerprint.SendTransaction(fingerContractAddress.Hex())
//	}
//}
//
////凭证提交后，提交ring
//func (bp *BucketProxy) submitRing(ringHash string) error {
//
//	return nil
//}
//
////todo:imp it
//func (bp *BucketProxy) canSubmit(ring *types.RingState) bool {
//	return true;
//}

func (bp *BucketProxy) listenRingSubmit() {
	for {
		select {
		case ring := <-matchengine.RingSubmitFailedChan:
			bp.submitFailed(ring)
		}
	}
}

//todo:需要ringclient在提交失败后通知到该proxy，估计使用chan
func (bp *BucketProxy) submitFailed(ring *types.RingState) {
	for _,bucket := range bp.Buckets {
		bucket.SubmitFailed(ring)
	}
}



