package realtime

import "github.com/Loopring/ringminer/types"

//
type Pool interface {
	GetOrders() []*types.Order
	AddOrder() error
	UpdateOrder() error
}
