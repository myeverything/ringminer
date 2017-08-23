package orderbook

import (
	"sync"
	"github.com/Loopring/ringminer/types"
)

type OrderWrapper struct {
	types.Order
	//被哪些环路匹配过了、剩余金额等
}

type OrderBook struct {
	Listeners []Listener
	Filters []Filter
	Store Store
	Orders []OrderWrapper //保存未匹配的订单列表
	mtx  sync.RWMutex

	OrderChan chan types.Order
}

func (ob *OrderBook) AddListener(listener Listener) {
	ob.Listeners = append(ob.Listeners, listener)
}

func (ob *OrderBook) AddFilter(filter Filter) {
	ob.Filters = append(ob.Filters, filter)
}

/**
接收到新订单后，进行保存、发送到match
 */
func (ob *OrderBook) NewOrder(order types.Order) {
	for _, filter := range ob.Filters {
		filter.filter(order)
	}
	var orderWrapper OrderWrapper
	orderWrapper.Order = order
	ob.Orders = append(ob.Orders, orderWrapper)
}

func (ob *OrderBook) GetOrder() {

}

//根据查询条件以及排序返回订单列表
func (ob *OrderBook) GetOrders() {

}

func (ob *OrderBook) Start() {
	for _, listener := range ob.Listeners {
		listener.Listen()
	}
}

func init() {

}

