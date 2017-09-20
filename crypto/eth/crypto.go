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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/Loopring/ringminer/chainclient/eth"
	"math/big"
)

type EthCrypto struct {}

//生成账号
func (c *EthCrypto) GenerateAccount(result interface{}) {
	privKey,err := crypto.GenerateKey()
	if nil == err {
		account := eth.Account{}
		account.PrivKey = privKey
		account.PubKey = &privKey.PublicKey
		account.Address = crypto.PubkeyToAddress(*account.PubKey)
		result = account
	}
}

//签名验证
func (c *EthCrypto) ValidateSignatureValues(v byte, r, s *big.Int, homestead bool) bool {
	return crypto.ValidateSignatureValues(v, r, s, homestead)
}

//生成hash
func (c *EthCrypto) GenerateHash(data ...[]byte) []byte {
	return crypto.Keccak256(data...)
}

//签名回复到地址
func (c *EthCrypto) SigToAddress(hash, sig []byte) ([]byte, error) {
	pubKey,err := crypto.SigToPub(hash, sig)
	if nil != err {
		return nil, err
	} else {
		return crypto.PubkeyToAddress(*pubKey).Bytes(), nil
	}
}


//
//// TODO(fukun): 使用go-eth/crypto/keccak256生成hash，需要跟智能合约比对
//func GenOrderHash(ord types.Order) []byte {
//	return crypto.Keccak256(
//		ord.Protocol.Bytes(),
//		ord.TokenS.Bytes(),
//		ord.TokenB.Bytes(),
//		ord.AmountS.Bytes(),
//		ord.AmountB.Bytes(),
//		[]byte(strconv.FormatUint(ord.Expiration, 10)),
//		ord.Rand.Bytes(),
//		ord.LrcFee.Bytes(),
//		[]byte(strconv.FormatBool(ord.BuyNoMoreThanAmountB)),
//		[]byte(strconv.Itoa(ord.SavingSharePercentage)),
//	)
//}
//
//// TODO(fukun): 使用自实现方式生成address
//func GenOrderAddress(hash []byte, ord types.Order) ([]byte, error) {
//
//	//if len(hash) != 32 {
//	//	return nil, errors.New("GenOrderAddress error,hash length is incorrect")
//	//}
//	//
//	//data, err := crypto.Ecrecover(
//	//	crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), hash),
//	//	[]byte(strconv.FormatUint(uint64(ord.V), 10)),
//	//	ord.R.Bytes(),
//	//	ord.S.Bytes(),
//	//)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//return data, nil
//
//	return nil, nil
//}
//
//// TODO(fukun): 调用合约方式生成hash
//// TODO(fukun): 调用合约方式生成address