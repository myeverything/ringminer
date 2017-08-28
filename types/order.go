package types

import "math/big"

//订单原始信息
type Order struct {
	Id          Hash      // 订单id
	Protocol    Address   // 智能合约地址
	Owner       Address   // 订单发起者地址
	OutToken    Address   // 卖出erc20代币智能合约地址
	InToken     Address   // 买入erc20代币智能合约地址
	OutAmount   *big.Int  // 卖出erc20代币数量上限
	InAmount    *big.Int  // 买入erc20代币数量上限
	Expiration  uint64    // 订单过期时间
	Fee         *big.Int  // 交易总费用,部分成交的费用按该次撮合实际卖出代币额与比例计算
	SavingShare *big.Int  // 不为0时支付给交易所的分润比例，否则视为100%
	V           uint8
	R           Sign
	S           Sign
}

// TODO(fukun): 包含成交记录
type OrderWrap struct {
	Order             `json:"order"`
	PeerId   string   `json:"peerId"`
	RingList []Hash   `json:"ringList"`
}


type NewOrderEvent struct {
	Order
	PeerId   string   `json:"peerId"`
}

type OrderRingEvent struct {
	OrderRing
}

//todo:order、ring、event等重新整理定义
type BalanceChangeEvent struct {
	Address	Address
	Balance	*big.Int
	Change	*big.Int
}