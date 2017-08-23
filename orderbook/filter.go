package orderbook

import "github.com/Loopring/ringminer/models"
/**
order过滤
 */

type Filter interface {
	filter(order models.Order)
}