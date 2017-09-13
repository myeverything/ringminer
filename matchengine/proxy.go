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
	"github.com/Loopring/ringminer/chainclient/eth"
)

var loopring *chainclient.Loopring


//代理，控制整个match流程，其中会提供几种实现，如bucket、realtime，etc。

/**
orderbook 实现:
1、订单的最新状态
2、订单是否满足剩余交易量等条件
 */
//var orderMinAmount big.Int	//todo：订单的最小金额，可能需要用map，记录每种货币的最小金额，应该定义filter，过滤环的验证规则

var OrderStateChan types.EngineOrderChan

//ring 的失败包括：提交失败，ring的合约执行时失败，执行时包括：gas不足，以及其他失败
var RingSubmitFailedChan chan *types.RingState

type Proxy interface {
	Start()  //启动
	Stop() //停止
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
	//todo:change to inject
	loopring = &chainclient.Loopring{}
	loopring.LoopringImpls = make(map[types.Address]*chainclient.LoopringProtocolImpl)
	loopring.LoopringFingerprints = make(map[types.Address]*chainclient.LoopringFingerprintRegistry)
	loopring.Tokens = make(map[types.Address]*chainclient.Erc20Token)
	loopring.Client = eth.EthClient
	orderChan := make(chan *types.OrderState)
	OrderStateChan = types.EngineOrderChan(orderChan)

	RingSubmitFailedChan = make(chan *types.RingState)
}