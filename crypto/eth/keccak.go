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

package eth

import (
	"github.com/Loopring/ringminer/types"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
)

/*
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

// TODO(fukun):
func GenHash(ord types.Order) {
	crypto.Keccak256(
		ord.Protocol.Bytes(),
		ord.TokenS.Bytes(),
		ord.TokenB.Bytes(),
		ord.AmountS.Bytes(),
		ord.AmountB.Bytes(),
		[]byte(ord.Expiration),
		ord.Rand.Bytes(),
		ord.LrcFee.Bytes(),
		[]byte(strconv.FormatBool(ord.BuyNoMoreThanAmountB)),
		[]byte(strconv.Itoa(ord.SavingSharePercentage)),
	)
}