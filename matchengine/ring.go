package matchengine

import (
	"github.com/Loopring/ringminer/types"
	"strconv"
	"math/rand"
	"math/big"
)

//ring相关的，包含链的判定、费用等数据计算等, 分润费用是必须计算的，便于判断等

//type Ring types.Ring

func IsRing() {

}

//费用、收取费用方式、折扣率等一切计算，在此完成
//计算匹配比例
//todo:折扣
//周一：完成费用的计算，并进行完整的测试
//本周：完成撮合引擎、整体的思路设计
func ComputeRing(ring *types.RingState) {
	DECIMALS := big.NewInt(10000) //todo:最好采用10的18次方，或者对应token的DECIMALS

	ring.LegalFee = big.NewInt(1)

	//根据订单原始金额，计算成交量、成交价
	//todo:
	//reducedRate := 0.1
	//ring.ReducedRate = big.NewInt(int(DECIMALS * reducedRate))
	ring.ReducedRate = big.NewInt(1000)
	shareRate := 0
	//todo:计算fee，为了取出最大的环路
	//LRC等比例下降，首先需要计算fillAmountS
	//分润的fee，首先需要计算fillAmountS，fillAmountS取决于整个环路上的完全匹配的订单
	//如何计算最小成交量的订单，计算下一次订单的卖出或买入，然后根据比例替换
	minVolumeIdx := 0

	for idx, order := range ring.RawRing.Orders {
		//var nextOrder *types.OrderState
		//todo：计算price，为了计算交易量
		sPrice := big.NewInt(0)
		a := big.NewInt(0)
		a.Mul(order.RawOrder.AmountS, DECIMALS)
		sPrice.Div(a , order.RawOrder.AmountB)
		sPrice.Mul(sPrice, ring.ReducedRate)

		//指定feeSelection
		//todo:当以Sell为基准时，考虑账户余额、订单剩余金额的最小值
		order.AvailableAmountS = big.NewInt(10000)

		//根据用户设置，判断是以卖还是买为基准
		//买入不超过amountB
		if (order.RawOrder.BuyNoMoreThanAmountB) {
			//不在计算具体的成交量，因此，不需要价格
			//sPrice := big.NewInt(0)
			////价格扩大DECIMALS
			//enlargeAmountS := big.NewInt(1)
			//enlargeAmountS.Mul(order.RawOrder.AmountS, DECIMALS)
			//enlargeAmountS.Mul(enlargeAmountS, ring.ReducedRate) //折价比例
			//sPrice.Div(enlargeAmountS, order.RawOrder.AmountB)
			//
			//sAmount := big.NewInt(0)
			////计算出来的为卖出的折价部分
			//sAmount.Mul(order.RawOrder.AmountS, sPrice) //todo:小数 通过乘以系数转换为整数进行计算
			//
			////扩大了两次系数
			//sAmount.Div(sAmount, DECIMALS)
			//sAmount.Div(sAmount, DECIMALS)
			sAmount := big.NewInt(0)	//节省的金额
			sAmount.Mul(order.RawOrder.AmountS, ring.ReducedRate)
			order.FeeS = &big.Int{}
			order.FeeS.Mul(sAmount, big.NewInt(int64(order.RawOrder.SavingSharePercentage)))
			decimals := &big.Int{}
			decimals.Mul(DECIMALS, big.NewInt(100))
			//减去扩大系数，分润的系数
			order.FeeS.Div(order.FeeS, decimals)

			//减去扩大系数
			sAmount.Div(sAmount, DECIMALS)

			order.RateAmountS = &big.Int{}
			order.RateAmountS.Sub(order.RawOrder.AmountS, sAmount)


			//todo:计算availableAmountS
			remainAmountB := big.NewInt(100)
			availableAmountS := big.NewInt(1)
			rate := big.NewInt(10)
			rate.Div(remainAmountB, order.RawOrder.AmountB)
			availableAmountS.Mul(order.RateAmountS, rate)
			if (availableAmountS.Cmp(order.AvailableAmountS) < 0) {
				order.AvailableAmountS = availableAmountS
			}

		} else {
			order.RateAmountS = order.RawOrder.AmountS

			//todo:
			remainAmountS := big.NewInt(1000)
			if (remainAmountS.Cmp(order.AvailableAmountS) < 0) {
				order.AvailableAmountS = remainAmountS
			}
			//order.FeeS = &big.Int{}
			//order.FeeS.Mul(sAmount, big.NewInt(order.RawOrder.SavingSharePercentage))
			//decimals := &big.Int{}
			//decimals.Mul(DECIMALS, big.NewInt(100))
			////减去扩大系数，分润的系数
			//order.FeeS.Div(order.FeeS, decimals)
		}

		//与上一订单的买入进行比较
		var lastOrder *types.FilledOrder
		if (idx > 0) {
			lastOrder = ring.RawRing.Orders[idx - 1]
		}

		if (lastOrder != nil && lastOrder.FillAmountB.Cmp(order.AvailableAmountS) > 0) {
			//当前订单为最小订单
			order.FillAmountS = order.AvailableAmountS
			minVolumeIdx = idx
		} else {
			//上一订单为最小订单需要对remainAmountS进行折扣计算
			order.FillAmountS = order.AvailableAmountS
		}

		order.FillAmountB = big.NewInt(0)
		order.FillAmountB.Div(order.FillAmountS, sPrice);

	}

	//根据minVolumeIdx进行最小交易量的计算,两个方向进行

	for i := minVolumeIdx; i >= 0; i-- {
		//按照前面的，同步减少交易量
	}

	for i := minVolumeIdx; i < len(ring.RawRing.Orders); i++ {

	}


	//计算ring以及各个订单的费用，以及费用支付方式
	for _, order := range ring.RawRing.Orders {
		//todo:成本节约为：
		savingAmount := big.NewInt(0)
		savingAmount.Mul(order.AvailableAmountS, ring.ReducedRate)
		order.FeeS = savingAmount
		//lrcFee等比例
		order.LrcFee = big.NewInt(1)
		rate := big.NewInt(0)
		rate.Div(order.AvailableAmountS , order.RawOrder.AmountS)
		order.LrcFee.Mul(order.RawOrder.LrcFee , rate)

		legalAmountOfLrc := big.NewInt(1)
		lrcAddress := &types.Address{}
		lrcAddress.SetBytes([]byte(LRC_ADDRESS))
		legalAmountOfLrc.Mul(order.LrcFee, GetLegalRate(CNY, *lrcAddress))

		legalAmountOfS := big.NewInt(1)
		//todo:address of sell token
		legalAmountOfS.Mul(order.FeeS, GetLegalRate(CNY, *lrcAddress))

		//todo：金额转换为法币
		if (legalAmountOfLrc.Cmp(legalAmountOfS) > 0) {
			order.FeeSelection = 0
			order.LegalFee = legalAmountOfLrc
		} else {
			order.FeeSelection = 1
			order.LegalFee = legalAmountOfS
		}

		//todo:计算ReducedRate
		if (shareRate > order.RawOrder.SavingSharePercentage) {
			shareRate = order.RawOrder.SavingSharePercentage
		}

		ring.LegalFee.Add(ring.LegalFee, order.LegalFee)
	}


	//
	//shareLegalTenderAmount := big.NewFloat(0.0)
	//if (shareRate > 0) {
	//	for _,ord := range ring.RawRing.Orders {
	//
	//		//todo:math method
	//		//ord.ShareFee = ord.Volume * (ord.OutAmount/ord.InAmount) * reducedRate
	//		ord.ShareFee = big.NewInt(22)
	//
	//		//todo:数据计算
	//		//ord.LegalTenderAmount = ord.ShareFee * GetLegalRate(CNY, ord.RawOrder.OutToken)
	//		//shareLegalTenderAmount += ord.ShareFee * GetLegalRate(CNY, ord.RawOrder.OutToken)
	//		shareLegalTenderAmount.Add(shareLegalTenderAmount, big.NewFloat(2.0))
	//	}
	//}
	//
	//lrcLegalTenderAmount := big.NewFloat(0.0)
	//for _,ord := range ring.Orders {
	//	//ord.Fee = ord.Fee * (ord.Volume/ord.RawOrder.OutAmount)
	//	//todo: fee计算
	//	ord.Fee = big.NewInt(11)
	//
	//	lrcAddress := &types.Address{}
	//	lrcAddress.SetBytes([]byte(LRC_ADDRESS))
	//	//todo: 数据计算
	//	lrcLegalTenderAmount.Add(lrcLegalTenderAmount, big.NewFloat(2.0))
	//	//lrcLegalTenderAmount += ord.Fee * GetLegalRate(CNY, lrcAddress)
	//}
	//
	////todo:2.0为返还用户的lrc比例
	//lrcFee := lrcLegalTenderAmount.Mul(lrcLegalTenderAmount, big.NewFloat(2.0))
	//if (shareLegalTenderAmount.Cmp(lrcFee) >= 0 ) {
	//	ring.FeeMode = 1
	//	ring.LegalTenderAmount = shareLegalTenderAmount
	//} else {
	//	ring.FeeMode = 0
	//	ring.LegalTenderAmount = lrcLegalTenderAmount
	//}
}

func Hash(ring *types.RingState) types.Hash {
	h := &types.Hash{}
	 h.SetBytes([]byte(strconv.Itoa(rand.Int())))
	return *h
}


//成环之后才可计算能否成交，否则不需计算，判断是否能够成交，不能使用除法计算
func PriceValid(ring *types.RingState) bool {
	amountS := big.NewInt(1)
	amountB := big.NewInt(1)
	for _, order := range ring.RawRing.Orders {
		amountS.Mul(amountS, order.RawOrder.AmountS)
		amountB.Mul(amountB, order.RawOrder.AmountB)
	}
	return amountS.Cmp(amountB) >= 0
}
