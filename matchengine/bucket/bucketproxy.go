package bucket

import (
	"sync"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/matchengine"
	"strconv"
)

/**
接受新订单并将新订单发送到各个bucket
每个bucket会将匹配成环的订单发送到该proxy进行汇总以及发送到区块链，并会将环发送到各个bucket，
 */

//todo:采用这种方式，需要计算出可能的所有的semiring的数量，对于增长数量，需要有理论支撑

/**
todo：设计合理的数据结构，需要满足，便于双向查找，便于添加、删除、修改等
todo：下周计划：完成以太坊链接、开始开发bucket、bucket的数据结构设计
 */

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
orderbook的所有的事件监听都需要实现，如：neworder、banlanceChange等
reactor模式
负责协调各个bucket，将ring发送到区块链，
该处负责接受neworder, updateorder等事件，并把事件广播给所有的bucket，同时调用client将已形成的环路发送至区块链，发送时需要再次查询订单的最新状态，保证无错，一旦出错需要更改ring的各种数据，如交易量、费用分成等
 */

type BucketProxy struct {
 	ringChan chan *types.RingState
	OrderChan chan *types.Order
	Buckets  map[types.Address]Bucket
	mtx  *sync.RWMutex
}

func NewBucketProxy() matchengine.Proxy {
	var proxy matchengine.Proxy
	bp := &BucketProxy{}

	ringChan := make(chan *types.RingState, 1000)
	bp.ringChan = ringChan

	orderChan := make(chan *types.Order)
	bp.OrderChan = orderChan

	bp.mtx = &sync.RWMutex{}
	bp.Buckets = make(map[types.Address]Bucket)

	proxy = bp
	return proxy
}

func (bp *BucketProxy) Start() {
	//proxy := bp
	bp.mtx.RLock()
	defer bp.mtx.RUnlock()

	for {
		select {
		case orderRing := <- bp.ringChan:
			s := ""
			for _,o := range orderRing.RawRing.Orders {
				s = s + " -> " + " {outtoken:" + string(o.RawOrder.TokenS.Bytes()) + " amount:" + o.RateAmountS.String() + ", intoken:" + string(o.RawOrder.TokenB.Bytes()) + "}"
			}
			println("ringChan receive:" + string(orderRing.Hash.Bytes()) + " ring is :" + s +
				" fee:" )
			println(orderRing.LegalFee.String())
			for _, b := range bp.Buckets {
				b.NewRing(orderRing)
			}
		//	newOrderRing(orderRing)
		////发送给每个bucket
		//	for _, bucket := range proxy.Buckets {
		//		go bucket.NewOrderRing(orderRing)
		//	}
		//case order := <- proxy.OrderChan:
		//	proxy.NewOrder(order)
		//	for _, bucket := range proxy.Buckets {
		//		go bucket.NewOrder(*order)
		//	}
		}
	}



}

func (bp *BucketProxy) Stop() {
	close(bp.ringChan)
	close(bp.OrderChan)
	for _,bucket := range bp.Buckets {
		bucket.Stop()
	}
}

func (bp *BucketProxy) NewOrder(order *types.OrderState) {
	//ring := &types.Ring{}
	//h := &types.Hash{}
	//h.SetBytes([]byte("1111"))
	//ring.Id = *h
	//bp.ringChan <- ring

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

	//println("new order:" + string(ord.OutToken.Bytes()) + " -> " + string(ord.InToken.Bytes()))
	//println("bucket count:" + strconv.Itoa(len(bp.Buckets)))
	//var c int
	for _, b := range bp.Buckets {
		b.NewOrder(*order)
		//
		//c1 := len(b.semiRings)
		//if (c <= c1) {
		//	c = c1
		//	println("bucket name:" + b.token.Str() +
		//		" semiRing count:" + strconv.Itoa(c) +
		//		" order count:" + strconv.Itoa(len(b.orders)) +
		//		" idx:" + string(ord.Id.Bytes()))
		//}
		println("bucket name:" + b.token.Str() +
			" semiRing count:" + strconv.Itoa(len(b.semiRings)) +
			" order count:" + strconv.Itoa(len(b.orders)) +
			" idx:" + string(order.OrderHash.Bytes()))
	}
}

func (bp *BucketProxy) UpdateOrder(order *types.OrderState) {
	for _, bucket := range bp.Buckets {
		bucket.UpdateOrder(*order)
	}
} //订单的更新

func (bp *BucketProxy) AddFilter() {

}

//todo:ring提交失败的处理




