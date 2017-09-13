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

package matchengine

import (
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/chainclient"
)

//代理，控制整个match流程，其中会提供几种实现，如bucket、realtime，etc。

/**
orderbook 实现:
1、订单的最新状态
2、订单是否满足剩余交易量等条件
 */
//var orderMinAmount big.Int	//todo：订单的最小金额，可能需要用map，记录每种货币的最小金额，应该定义filter，过滤环的验证规则

type RingSubmitFailedChan chan *types.RingState

var Loopring *chainclient.Loopring

type Proxy interface {
	Start()  //启动
	Stop() //停止
	GetOrderStateChan() chan *types.OrderState
	AddFilter()
}

/**
todo：仍需要工作：调试合约
1、保存匹配过的ring
2、ring发送前的检测
3、费用的整理
4、
 */

func init() {
	//RingSubmitFailedChan = make(chan *types.RingState)
}