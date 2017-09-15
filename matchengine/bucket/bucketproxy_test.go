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

package bucket_test

import (
	"math/big"
	"strconv"
	"github.com/Loopring/ringminer/types"
	"testing"
	"github.com/Loopring/ringminer/matchengine/bucket"
	"time"
	"github.com/Loopring/ringminer/matchengine"
	"github.com/Loopring/ringminer/config"
	"math"
	"github.com/Loopring/ringminer/log"
)

func newOrder(outToken string, inToken string, outAmount, inAmount int64, buyFirstEnough bool, idx int) *types.OrderState {
	orderState := &types.OrderState{}
	order := &types.Order{}

	outAddress := &types.Address{}
	outAddress.SetBytes([]byte(outToken))
	inAddress := &types.Address{}
	inAddress.SetBytes([]byte(inToken))

	order.TokenS = *outAddress
	order.TokenB = *inAddress
	order.AmountS = big.NewInt(outAmount)
	order.AmountB = big.NewInt(inAmount)
	order.BuyNoMoreThanAmountB = buyFirstEnough
	order.LrcFee = big.NewInt(1000)
	order.SavingSharePercentage = 30
	h := &types.Hash{}
	h.SetBytes([]byte(strconv.Itoa(idx)))
	orderState.RawOrder = *order
	orderState.OrderHash = *h
	orderState.Status = types.ORDER_NEW
	return orderState
}

func TestBucketProxy(t *testing.T) {

	ringClient := matchengine.NewRingClient()
	ringClient.Start()
	c := &config.BucketProxyOptions{}
	proxy := bucket.NewBucketProxy(*c, ringClient)
	go proxy.Start()

	order1 := newOrder("token1", "token2", 20000, 10000, true, 1)

	proxy.GetOrderStateChan() <- order1

	order2 := newOrder("token2", "token3", 30000, 30000, true,  2)
	proxy.GetOrderStateChan() <- order2

	order3 := newOrder("token3", "token1", 40000, 20000, true,  3)
	proxy.GetOrderStateChan() <- order3

	time.Sleep(100000000)
}

func TestBaoxian(t *testing.T) {
	//baseAmount := 0.5
	//rate := 0.05

	log.NewLogger()
	//log.Error("llll", "dddd")
	//order1 := newOrder("token1", "token2", 20000, 10000, true, 1)
	//
	//log.Info("dddddddd%s",log.NewField("order", order1))
	//
	//println( time.Now().Unix())
	//for i:=0;i<=10;i++ {
	//	log.Fatal("dddddddd",log.NewField("order", "kkkk"))
	//
	//}

	//incomeAmount := 0.0
	//println(baseAmount*rate*2 + baseAmount*rate*rate)
	//incomeAmount = income(baseAmount, rate, 20, 20, incomeAmount)
	//println(int(incomeAmount*100))
}

func income(baseAmount,rate float64,number int,years int, incomeAmount float64) float64 {
	if (number > 0) {
		number--
		//println("number:", number, " years:", years)
		incomeAmount = incomeAmount + baseAmount * float64(years - number) *math.Pow(rate, float64(1+number))
		return income(baseAmount, rate, number,years, incomeAmount)
	} else {
		return incomeAmount
	}
}
