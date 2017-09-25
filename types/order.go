/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package types

import (
	"github.com/Loopring/ringminer/crypto"
	"github.com/Loopring/ringminer/log"
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
	Expiration            uint64   // 订单过期时间
	Rand                  *big.Int
	LrcFee                *big.Int // 交易总费用,部分成交的费用按该次撮合实际卖出代币额与比例计算
	BuyNoMoreThanAmountB  bool
	SavingSharePercentage int // 不为0时支付给交易所的分润比例，否则视为100%
	V                     uint8
	R                     *big.Int
	S                     *big.Int
}

/**
address(this),
            order.tokenS,
            order.tokenB,
            order.amountS,
            order.amountB,
            order.expiration,
            order.rand,
            order.lrcFee,
            order.buyNoMoreThanAmountB,
            order.savingSharePercentage
*/

func (o *OrderState) GenHash() Hash {
	h := &Hash{}
	hashBytes := crypto.CryptoInstance.GenerateHash(
		o.RawOrder.Protocol.Bytes(),
		o.RawOrder.TokenS.Bytes(),
		o.RawOrder.TokenB.Bytes(),
		o.RawOrder.AmountS.Bytes(),
		o.RawOrder.AmountB.Bytes(),
		[]byte{byte(o.RawOrder.Expiration)},
		o.RawOrder.Rand.Bytes(),
		o.RawOrder.LrcFee.Bytes(),
		[]byte{byte(0)},
		[]byte{byte(o.RawOrder.SavingSharePercentage)},
	)
	h.SetBytes(hashBytes)

	o.OrderHash = *h

	return *h
}

func (o *OrderState) ValidateSignatureValues() bool {
	return crypto.CryptoInstance.ValidateSignatureValues(byte(o.RawOrder.V), o.RawOrder.R, o.RawOrder.S)
}

func (o *OrderState) SignerAddress() (Address, error) {
	//todo:
	hash := o.GenHash()
	sig := crypto.CryptoInstance.VRSToSig(o.RawOrder.V, o.RawOrder.R, o.RawOrder.S)
	println("hash:", hash.Hex())
	address := &Address{}
	if addressBytes, err := crypto.CryptoInstance.SigToAddress(hash.Bytes(), sig); nil != err {
		log.Errorf("error:%s", err.Error())
		return *address, err
	} else {
		address.SetBytes(addressBytes)
		o.Owner = *address
		return *address, nil
	}
	//address := &Address{}
	//address.SetBytes(common.HexToAddress("0x1b9f8da77f2288eb4ce1096b1cfdc30cc4d0a961").Bytes())

	//o.Owner = *address
	//return *address,nil
}

//RateAmountS、FeeSelection 需要提交到contract
type FilledOrder struct {
	OrderState       OrderState
	FeeSelection     int      //0 -> lrc
	RateAmountS      *big.Int //提交需要
	AvailableAmountS *big.Int //需要，也是用于计算fee
	//AvailableAmountB *big.Int	//需要，也是用于计算fee
	FillAmountS *EnlargedInt
	FillAmountB *EnlargedInt //计算需要
	LrcReward   *EnlargedInt
	LrcFee      *EnlargedInt
	FeeS        *EnlargedInt
	//FeeB             *EnlargedInt
	LegalFee *EnlargedInt //法币计算的fee

	EnlargedSPrice *EnlargedInt
	EnlargedBPrice *EnlargedInt

	//FullFilled	bool	//this order is fullfilled
}

type OrderState struct {
	RawOrder        Order
	Owner           Address
	OrderHash       Hash
	RemainedAmountS *big.Int
	RemainedAmountB *big.Int
	Status          OrderStatus
}

// TODO(fukun): 来自以太坊的订单
type OrderMined struct {
}

// convert order to ordersate
func (ord *Order) Convert() *OrderState {
	var s OrderState
	s.RawOrder = *ord

	// TODO(fukun): 计算owner，hash等
	//s.Owner = StringToAddress("")
	//s.OrderHash = StringToHash("")
	s.RemainedAmountS = s.RawOrder.AmountS
	s.RemainedAmountB = s.RawOrder.AmountB
	s.Status = ORDER_NEW

	return &s
}
