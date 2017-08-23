package orderbook

import "github.com/Loopring/ringminer/types"
/**
order过滤
 */

type Filter interface {
	filter(order types.Order)
}