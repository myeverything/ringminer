package p2p

import (
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/log"
)

func GenOrder(data []byte) *types.Order {
	var ord types.Order
	err := ord.UnMarshalJson(data)
	if err != nil {
		log.Error(log.ERROR_P2P_LISTEN_ACCEPT,  err.Error())
	} else {
		log.Info(log.LOG_P2P_ACCEPT, string(data))
	}

	return &ord
}