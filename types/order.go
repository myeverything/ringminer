package types

import (
	"math/big"
)

//订单原始信息
/**
1、是否整体成交
2、指定成交对象，对方单的hash
3、分润比例 是否需要设置
4、成交方向 待定
5、过期时间，使用块数
 */
type Order struct {
	Protocol              Address  // 智能合约地址
	TokenS                Address  // 卖出erc20代币智能合约地址
	TokenB                Address  // 买入erc20代币智能合约地址
	AmountS               *big.Int // 卖出erc20代币数量上限
	AmountB               *big.Int // 买入erc20代币数量上限
	Rand                  *big.Int
	Expiration            uint64   // 订单过期时间
	LrcFee                *big.Int // 交易总费用,部分成交的费用按该次撮合实际卖出代币额与比例计算
	SavingSharePercentage int      // 不为0时支付给交易所的分润比例，否则视为100%
	IsCompleteFillMeasuredByTokenSDepleted bool
	V                     uint8
	R                     Sign
	S                     Sign
}

//RateAmountS、RateAmountB、FeeSelection 需要提交到contract
type FilledOrder struct {
	Order *Order
	FeeSelection int
	RateAmountS *big.Int
	RateAmountB *big.Int
	AvailableAmountS *big.Int
	FillAmountS *big.Int
	LrcReward *big.Int
	LrcFee *big.Int
	FeeSForThisOrder *big.Int
	FeeSForNextOrder *big.Int
}

type OrderState struct {
	Order *Order
	Owner Address
	OrderHash Hash
	RemainedAmountS int
	RemainedAmountB int
	//....
}

/**ring
1、撮合者费用的收益的地址
2、
 */
// TODO(fukun): 包含成交记录
type OrderWrap struct {
	RawOrder *Order             `json:"rawOrder"`
	PeerId   string   `json:"peerId"`
	OutAmount *big.Int	`json:"outAmount"` // 剩余量
	InAmount  *big.Int	`json:"inAmount"`  // 剩余量
	Fee *big.Int	`json:"fee"`
	ShareFee *big.Int ``
}

type NewOrderEvent struct {
	Order
	//PeerId   string   `json:"peerId"`
	//OrderHash Hash
}

type OrderRingEvent struct {
	RingState
}

//todo:order、ring、event等重新整理定义
type BalanceChangeEvent struct {
	Address	Address
	Balance	*big.Int
	Change	*big.Int
}