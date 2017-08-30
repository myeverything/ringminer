package orderbook

import (
	"sync"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/lrcdb"
)

const (
	FINISH_TABLE_NAME = "finished"
	PARTIAL_TABLE_NAME = "partial"
)

type OrderBook struct {
	db           lrcdb.Database
	finishTable  lrcdb.Database
	partialTable lrcdb.Database
	lock         sync.RWMutex
}

var orderBook OrderBook

type OrderBookConfig struct {
	DBName           string
	DBCacheCapcity   int
	DBBufferCapcity  int
}

// TODO(fukun): 通过智能合约查询未完成订单状态，完成后开始与matchengine交互
func InitializeOrderBook(c *OrderBookConfig) {
	orderBook.db = lrcdb.NewDB(c.DBName, c.DBCacheCapcity, c.DBBufferCapcity)
	orderBook.finishTable = lrcdb.NewTable(orderBook.db, FINISH_TABLE_NAME)
	orderBook.partialTable = lrcdb.NewTable(orderBook.db, PARTIAL_TABLE_NAME)
}

// 订单只会来源于p2p网络
// 1.判断订单是否合法
// 2.存储订单到db
// 3.发送订单到matchengine
func NewOrder(ord *types.OrderWrap) error {
	// TODO(fukun): 判断订单是否合法

	key := ord.RawOrder.Id.Bytes()
	value,err := ord.MarshalJson()
	if err != nil {
		return err
	}

	orderBook.partialTable.Put(key, value)

	// TODO(fukun): 发送订单到
	return nil
}

func SetOrder() {

}

func GetOrder() {

}

//根据查询条件以及排序返回订单列表
func GetOrders() {

}

func (ob *OrderBook) moveOrder(ord *types.Order) {

}