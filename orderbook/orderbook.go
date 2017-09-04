package orderbook

import (
	"sync"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/lrcdb"
	"log"
	"os"
)

type ORDER_STATUS int

const (
	FINISH_TABLE_NAME = "finished"
	PARTIAL_TABLE_NAME = "partial"
)

type Config struct {
	DBName           string
	DBCacheCapcity   int
	DBBufferCapcity  int
}

type OrderBook struct {
	conf         Config
	db           lrcdb.Database
	finishTable  lrcdb.Database
	partialTable lrcdb.Database
	whisper      *types.Whispers
	lock         sync.RWMutex
}

func (ob *OrderBook) defaultConfig() {
	dir := os.Getenv("GOPATH") + "/github.com/Loopring/ringminer/"
	file := dir + "leveldb"
	ob.conf = Config{file, 8, 4}
}

// TODO(fukun): 通过智能合约查询未完成订单状态，完成后开始与matchengine交互
func NewOrderBook(whisper *types.Whispers) *OrderBook {
	s := &OrderBook{}
	s.defaultConfig()
	s.db = lrcdb.NewDB(s.conf.DBName, s.conf.DBCacheCapcity, s.conf.DBBufferCapcity)
	s.finishTable = lrcdb.NewTable(s.db, FINISH_TABLE_NAME)
	s.partialTable = lrcdb.NewTable(s.db, PARTIAL_TABLE_NAME)
	s.whisper = whisper

	return s
}

// 订单只会来源于p2p网络
// 1.判断订单是否合法
// 2.存储订单到db
// 3.发送订单到matchengine
func (s *OrderBook) Start() {
	go func() {
		for {
			select {
			case ord := <- s.whisper.PeerOrderChan:
				s.peerOrderHook(ord)
			case ord := <- s.whisper.ChainOrderChan:
				s.chainOrderHook(ord)
			}
		}
	}()
}

func (s *OrderBook) Stop() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.finishTable.Close()
	s.partialTable.Close()
	s.db.Close()
}

func (ob *OrderBook) peerOrderHook(ord *types.Order) error {
	// TODO(fukun): 判断订单是否合法
	ob.lock.Lock()
	defer ob.lock.Unlock()

	key := ord.GenHash().Bytes()
	value,err := ord.MarshalJson()
	if err != nil {
		return err
	}

	ob.partialTable.Put(key, value)

	// TODO(fukun): 发送订单到

	// TODO(fukun): delete after test
	if input, err := ob.partialTable.Get(key); err != nil {
		panic(err)
	} else {
		var ord types.Order
		ord.UnMarshalJson(input)
		log.Println(ord.TokenS.Str())
		log.Println(ord.TokenB.Str())
		log.Println(ord.AmountS.Uint64())
		log.Println(ord.AmountB.Uint64())
	}

	return nil
}

func (ob *OrderBook) chainOrderHook(ord *types.OrderMined) error {
	ob.lock.Lock()
	defer ob.lock.Unlock()

	return nil
}

func (ob *OrderBook) GetOrder(id types.Hash) (*types.OrderState, error) {
	var (
		value []byte
		err error
		ord types.OrderState
	)

	if value, err = ob.partialTable.Get(id.Bytes()); err != nil {
		value, err = ob.finishTable.Get(id.Bytes())
	}
	if err != nil {
		return nil, err
	}

	err = ord.UnMarshalJson(value)
	if err != nil {
		return nil, err
	}

	return &ord, nil
}

//根据查询条件以及排序返回订单列表
func (ob *OrderBook) GetOrders() {

}

// 只会从partial移动到finish
func (ob *OrderBook) moveOrder(odw *types.OrderState) error {
	key := odw.OrderHash.Bytes()
	value, err := odw.MarshalJson()
	if err != nil {
		return err
	}
	ob.partialTable.Delete(key)
	ob.finishTable.Put(key, value)
	return nil
}

// TODO(fukun): 从配置文件中读取不同合约地址对应代币的尘埃差值
func isFinished(odw *types.OrderState) bool {
	//if odw.RawOrder.

	return true
}