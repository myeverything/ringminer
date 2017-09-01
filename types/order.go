package types

import (
	"math/big"
)

type OrderStatus uint8

const (
	ORDER_NEW OrderStatus = iota
	ORDER_PENDING
	ORDER_PARTIAL
	ORDER_FINISHED
	ORDER_CANCEL
	ORDER_REJECT
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
	RawOrder			*Order
	FeeSelection 		int
	RateAmountS 		*big.Int
	RateAmountB 		*big.Int
	AvailableAmountS 	*big.Int
	FillAmountS 		*big.Int
	LrcReward 			*big.Int
	LrcFee 				*big.Int
	FeeSForThisOrder 	*big.Int
	FeeSForNextOrder 	*big.Int
}

type OrderState struct {
	RawOrder 		*Order
	Owner 			Address
	OrderHash 		Hash
	RemainedAmountS *big.Int
	RemainedAmountB *big.Int
	Status 			OrderStatus
}

// TODO(fukun): 来自以太坊的订单
type OrderMined struct {

}

type Whispers struct {
	PeerOrderChan			chan *Order
	ChainOrderChan			chan *OrderMined
	EngineOrderChan			chan *OrderState
}

// TODO(fukun):
func (ord *Order) GenHash() Hash {
	return StringToHash("")
}

// TODO(fukun)
func (ord *Order) VerifyHash() error {
	return nil
}