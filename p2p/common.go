package p2p

import (
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/log"
	"github.com/Loopring/ringminer/orderbook"
)

func Send(data []byte) {
	var ord types.Order
	err := ord.UnMarshalJson(data)
	if err != nil {
		log.Error(log.ERROR_P2P_LISTEN_ACCEPT,  err.Error())
	} else {
		log.Info(log.LOG_P2P_ACCEPT, string(data))
	}

	// 添加相关参数
	odw := &types.OrderWrap{}
	odw.RawOrder = &ord
	odw.PeerId = "0xdylenfu"
	odw.InAmount = types.IntToBig(0)
	odw.OutAmount = types.IntToBig(0)
	odw.Fee = types.IntToBig(0)

	orderbook.NewOrder(odw)
}