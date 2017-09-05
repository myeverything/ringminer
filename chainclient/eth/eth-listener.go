package eth

import (
)
import "github.com/Loopring/ringminer/types"


/**
区块链的listener, 得到order，
 */

//监听内容有：环路订单，地址的token余额变动如transfer等
//todo:不同的channel，应当交给orderbook统一进行后续处理，可以将channel作为函数返回值、全局变量、参数等方式
type EthListener struct {
	NewOrderChan	chan *types.NewOrderEvent
	OrderChan	chan *types.OrderRingEvent
	BalanceChangeChan	chan *types.BalanceChangeEvent
}

//应当返回channel
func (listener *EthListener) StartNewOrder() {
	listener.NewOrderChan = make(chan *types.NewOrderEvent)
	Client.Subscribe(&listener.NewOrderChan)
	for {
		select {
		case ord := <- listener.NewOrderChan:
			println(ord.SavingSharePercentage)
			//println(ord.Id.Hex())
			//orderbook.NewOrder(ord.Order)
		}
	}
}

func (listener *EthListener) Stop() {
	close(listener.NewOrderChan)
}

func (listener *EthListener) Name() string {
	return "eth-chain-listener"
}




