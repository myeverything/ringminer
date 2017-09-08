package types

type Whispers struct {
	PeerOrderChan			chan *Order
	ChainOrderChan			chan *OrderMined
	EngineOrderChan			chan *OrderState
}