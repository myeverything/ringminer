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

//go:generate gencodec -type Order -field-override orderMarshaling -out gen_order_json.go
type Order struct {
	Protocol              Address  `json:"protocol" gencodec:"required"`// 智能合约地址
	TokenS                Address  `json:"tokenS" gencodec:"required"`// 卖出erc20代币智能合约地址
	TokenB                Address  `json:"tokenB" gencodec:"required"`// 买入erc20代币智能合约地址
	AmountS               *big.Int `json:"amountS" gencodec:"required"`// 卖出erc20代币数量上限
	AmountB               *big.Int `json:"amountB" gencodec:"required"`// 买入erc20代币数量上限
	Expiration            *big.Int `json:"expiration" gencodec:"required"`// 订单过期时间
	Rand                  *big.Int `json:"rand" gencodec:"required"`
	LrcFee                *big.Int `json:"lrcFee" `// 交易总费用,部分成交的费用按该次撮合实际卖出代币额与比例计算
	BuyNoMoreThanAmountB  bool `json:"buyNoMoreThanAmountB" gencodec:"required"`
	SavingSharePercentage int `json:"savingSharePercentage" gencodec:"required"`// 不为0时支付给交易所的分润比例，否则视为100%
	V                     uint8 `json:"v" gencodec:"required"`
	R                     Sign `json:"r" gencodec:"required"`
	S                     Sign `json:"s" gencodec:"required"`

	Hash                  Hash `json:"-"`
}

type orderMarshaling struct {
	AmountS *Big
	AmountB *Big
	Expiration *Big
	Rand *Big
	LrcFee *Big
}

func (o *Order) GenerateHash() Hash {
	h := &Hash{}
	hashBytes := crypto.CryptoInstance.GenerateHash(
		o.Protocol.Bytes(),
		o.TokenS.Bytes(),
		o.TokenB.Bytes(),
		o.AmountS.Bytes(),
		o.AmountB.Bytes(),
		o.Expiration.Bytes(),
		o.Rand.Bytes(),
		o.LrcFee.Bytes(),
		[]byte{byte(0)}, //todo:o.BuyNoMoreThanAmountB to byte, test with contract
		[]byte{byte(o.SavingSharePercentage)},
	)
	h.SetBytes(hashBytes)

	return *h
}

func (o *Order) ValidateSignatureValues() bool {
	return crypto.CryptoInstance.ValidateSignatureValues(byte(o.V), o.R.Bytes(), o.S.Bytes())
}

func (o *Order) SignerAddress() (Address, error) {
	address := &Address{}
	hash := o.Hash
	//todo:how to check hash is nil,this use big.Int
	if hash.Big().Cmp(big.NewInt(0)) == 0 {
		hash = o.GenerateHash()
	}

	sig := crypto.CryptoInstance.VRSToSig(o.V, o.R.Bytes(), o.S.Bytes())
	log.Debugf("orderstate.hash:%s", hash.Hex())

	if addressBytes, err := crypto.CryptoInstance.SigToAddress(hash.Bytes(), sig); nil != err {
		log.Errorf("error:%s", err.Error())
		return *address, err
	} else {
		address.SetBytes(addressBytes)
		return *address, nil
	}
}

//RateAmountS、FeeSelection 需要提交到contract
//go:generate gencodec -type FilledOrder -field-override filledOrderMarshaling -out gen_filledorder_json.go

type FilledOrder struct {
	OrderState       OrderState `json:"orderState" gencodec:"required"`
	FeeSelection     int      `json:"feeSelection"`//0 -> lrc
	RateAmountS      *big.Int `json:"rateAmountS"`//提交需要
	AvailableAmountS *big.Int `json:"availableAmountS"`//需要，也是用于计算fee
	//AvailableAmountB *big.Int	//需要，也是用于计算fee
	FillAmountS *EnlargedInt `json:"fillAmountS"`
	FillAmountB *EnlargedInt `json:"fillAmountB"`//计算需要
	LrcReward   *EnlargedInt `json:"lrcReward"`
	LrcFee      *EnlargedInt `json:"lrcFee"`
	FeeS        *EnlargedInt `json:"feeS"`
	//FeeB             *EnlargedInt
	LegalFee *EnlargedInt `json:"legalFee"`//法币计算的fee

	EnlargedSPrice *EnlargedInt `json:"enlargedSPrice"`
	EnlargedBPrice *EnlargedInt `json:"enlargedBPrice"`

	//FullFilled	bool	//this order is fullfilled
}

type filledOrderMarshaling struct {
	RateAmountS *Big
	AvailableAmountS *Big
}


//go:generate gencodec -type OrderState -field-override orderStateMarshaling -out gen_orderstate_json.go
type OrderState struct {
	RawOrder        Order `json:"rawOrder"`
	Owner           Address	`json:"owner" `
	Hash            Hash `json:"hash"`
	RemainedAmountS *big.Int `json:"remainedAmountS"`
	RemainedAmountB *big.Int `json:"remainedAmountB"`
	Status          OrderStatus `json:"status"`
}

type orderStateMarshaling struct {
	RemainedAmountS *Big
	RemainedAmountB *Big
}


type OrderMined struct {

}