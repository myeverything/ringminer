package orderbook

import (
	"sync"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/lrcdb"
	"errors"
)

type OrderBook struct {
	db      lrcdb.Database
	tables  map[string]lrcdb.Database
	lock    sync.RWMutex
}

var orderBook OrderBook

type OrderBookConfig struct {
	DBName           string
	DBCacheCapcity   int
	DBBufferCapcity  int
}

// TODO(fukun): 通过智能合约查询未完成订单状态，完成后开始与matchengine交互
func (ob *OrderBook) InitializeOrderBook(c *OrderBookConfig) {
	orderBook.db = lrcdb.NewDB(c.DBName, c.DBCacheCapcity, c.DBBufferCapcity)
	//orderBook.tables
}

// 订单只会来源于p2p网络
// 1.判断订单是否合法
// 2.存储订单到db
// 3.发送订单到matchengine
func NewOrder(ord *types.Order) {
	// TODO(fukun): 判断订单是否合法

	// 存储订单

}

func SetOrder() {

}

func GetOrder() {

}

//根据查询条件以及排序返回订单列表
func GetOrders() {

}

func (ob *OrderBook) checkTableExist(prefix string) bool {
	if _, ok := ob.tables[prefix]; !ok {
		return false
	}
	return true
}

func (ob *OrderBook) getTable(tn string) (lrcdb.Database, error) {
	if table,ok := ob.tables[tn]; ok {
		return table, nil
	}
	return nil, errors.New("table " + tn + " doesn't exist")
}

func (ob *OrderBook) moveOrder(src,dst string, ord *types.Order) {

}