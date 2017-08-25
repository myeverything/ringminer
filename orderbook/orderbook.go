package orderbook

import (
	"sync"
	. "github.com/Loopring/ringminer/types"
)


type OrderBook struct {
	Filters []Filter
	Store Store
	Orders []Order//保存未匹配的订单列表
	mtx  sync.RWMutex
	OrderChan chan Order
}

var orderBook OrderBook

func init() {
	orderBook.OrderChan = make(chan Order)
}

func AddFilter(filter Filter) {
	orderBook.Filters = append(orderBook.Filters, filter)
}

/**
接收到新订单后，进行保存、发送到match
 */
func NewOrder(order Order) {
	for _, filter := range orderBook.Filters {
		filter.filter(order)
	}
	orderBook.mtx.Lock()
	defer orderBook.mtx.Unlock()
	orderBook.Orders = append(orderBook.Orders, order)
}

func SetOrder() {

}

func GetOrder() {

}

//根据查询条件以及排序返回订单列表
func GetOrders() {

}
