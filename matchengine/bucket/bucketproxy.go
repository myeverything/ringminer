package matchengine

import (
	"sync"
	"github.com/Loopring/ringminer/types"
)

/**
接受新订单并将新订单发送到各个bucket
每个bucket会将匹配成环的订单发送到该proxy进行汇总以及发送到区块链，并会将环发送到各个bucket，
 */

/**
todo:设计合理的数据结构，需要满足，便于双向查找，便于添加、删除、修改等
todo：下周计划：完成以太坊链接、开始开发bucket、bucket的数据结构设计

 */


/**
思路：设计符合要求的数据格式，
orderbook的所有的事件监听都需要实现，如：neworder、banlanceChange等
reactor模式
负责协调各个bucket，将ring发送到区块链，
该处负责接受neworder, updateorder等事件，并把事件广播给所有的bucket，同时调用client将已形成的环路发送至区块链，发送时需要再次查询订单的最新状态，保证无错，一旦出错需要更改ring的各种数据，如交易量、费用分成等
 */

//todo：地址长度
var TokenAddressLength int64

type TokenAddress [TokenAddressLength]byte

type BucketProxy struct {
 	ringChan *chan *types.Ring
	OrderChan *chan *types.Order
	Buckets  map[TokenAddress]Bucket
	mtx  *sync.RWMutex
}

func (bp *BucketProxy) Start() {
	proxy := bp
	ringChan := make(chan *types.Ring)
	proxy.ringChan = &ringChan
	for {
		select {
		//case orderRing := <- proxy.OrderRingChan:
		//	newOrderRing(orderRing)
		////发送给每个bucket
		//	for _, bucket := range proxy.Buckets {
		//		go bucket.NewOrderRing(orderRing)
		//	}
		case order := <- proxy.OrderChan:
			proxy.newOrder(order)
			for _, bucket := range proxy.Buckets {
				go bucket.NewOrder(order)
			}
		}
	}
}

func (bp *BucketProxy) Stop() {

}

func (bp *BucketProxy) newOrder(order *types.Order) {
	//如果没有则，新建bucket
	if _,ok := bp.Buckets[order.OutToken] ; !ok {
		bucket := &Bucket{}
		bucket.ringChan = bp.ringChan
		bp.mtx = &sync.RWMutex{}
	}

	for _, bucket := range bp.Buckets {
		bucket.NewOrder(*order)
	}
}

func (bp *BucketProxy) updateOrder(order *types.Order) {
	for _, bucket := range bp.Buckets {
		bucket.UpdateOrder(*order)
	}
} //订单的更新

func (bp *BucketProxy) AddFilter() {

}

//todo:ring提交失败的处理

func init() {

}




