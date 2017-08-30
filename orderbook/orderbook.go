package orderbook

import (
	"sync"
	"github.com/Loopring/ringminer/types"
)

type OrderBook struct {
	Filters []Filter
	Orders  []types.Order//保存未匹配的订单列表
	mtx     sync.RWMutex
}

var orderBook OrderBook

func AddFilter(filter Filter) {
	orderBook.Filters = append(orderBook.Filters, filter)
}

/**
接收到新订单后，进行保存、发送到match
 */
func (ob *OrderBook) InitializeOrderBook() {
	// TODO(fukun): add filter
	//for _, filter := range orderBook.Filters {
	//	filter.filter(ob)
	//}
	//orderBook.mtx.Lock()
	//defer orderBook.mtx.Unlock()
	//orderBook.Orders = append(orderBook.Orders, ob)
}

func NewOrder(ord *types.Order) {

}

func SetOrder() {

}

func GetOrder() {

}

//根据查询条件以及排序返回订单列表
func GetOrders() {

}
