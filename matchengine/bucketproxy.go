package matchengine

import (
	"sync"
	"github.com/Loopring/ringminer/types"
)

/**
接受新订单并将新订单发送到各个bucket
每个bucket会将匹配成环的订单发送到该proxy进行汇总以及发送到区块链，并会将环发送到各个bucket，
 */


type BucketProxy struct {
 	OrderRingChan chan Ring
	OrderChan chan types.Order
	Buckets  []Bucket
	mtx  sync.RWMutex
}

func (bp *BucketProxy) Start()  {
	proxy := bp
	for {
		select {
		case orderRing := <- proxy.OrderRingChan:
			newOrderRing(orderRing)
		//发送给每个bucket
			for _, bucket := range proxy.Buckets {
				go bucket.NewOrderRing(orderRing)
			}
		case order := <- proxy.OrderChan:
			newOrder(order)
			for _, bucket := range proxy.Buckets {
				go bucket.NewOrder(order)
			}
		}
	}
}

func (bp *BucketProxy) Stop() {

}

func (bp *BucketProxy) NewOrder() {

}

func (bp *BucketProxy) NewOrderRing(ring Ring) {

}

func (bp *BucketProxy) UpdateOrder()  //订单的更新

//环路检验
func (bp *BucketProxy) ringVerify() {

}

//ring提交失败的处理

func newOrderRing(ring Ring) {
	println(ring)
}

func newOrder(order Order) {
	println(order)
}

func init() {


}




