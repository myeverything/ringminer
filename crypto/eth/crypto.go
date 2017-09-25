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
	"github.com/Loopring/ringminer/chainclient/eth"
	"github.com/Loopring/ringminer/log"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type EthCrypto struct {
	Homestead bool
}

//生成账号
func (c *EthCrypto) GenerateAccount(result interface{}) {
	privKey, err := crypto.GenerateKey()
	if nil == err {
		account := eth.Account{}
		account.PrivKey = privKey
		account.PubKey = &privKey.PublicKey
		account.Address = crypto.PubkeyToAddress(*account.PubKey)
		result = account
	}
}

//签名验证
func (c *EthCrypto) ValidateSignatureValues(v byte, r, s *big.Int) bool {
	return crypto.ValidateSignatureValues(v, r, s, c.Homestead)
}

//生成hash
func (c *EthCrypto) GenerateHash(data ...[]byte) []byte {
	return crypto.Keccak256(data...)
}

//签名回复到地址
func (c *EthCrypto) SigToAddress(hash, sig []byte) ([]byte, error) {
	pubKey, err := crypto.SigToPub(hash, sig)
	if nil != err {
		return nil, err
	} else {
		return crypto.PubkeyToAddress(*pubKey).Bytes(), nil
	}
}

func (c *EthCrypto) VRSToSig(v byte, r, s *big.Int) []byte {
	sig := make([]byte, 65)
	copy(sig[32-len(r.Bytes()):32], r.Bytes())
	copy(sig[64-len(s.Bytes()):64], s.Bytes())
	sig[64] = v
	return sig
}

func (c *EthCrypto) Sign(hash, pkBytes []byte) ([]byte, error) {
	if pk, err := crypto.ToECDSA(pkBytes); err != nil {
		log.Errorf("err:", err.Error())
		return nil, err
	} else {
		return crypto.Sign(hash, pk)
	}
}

func (c *EthCrypto) SigToVRS(sig []byte) (v byte, r *big.Int, s *big.Int) {
	r = big.NewInt(0)
	s = big.NewInt(0)
	v = sig[64]
	r.SetBytes(sig[0:32])
	s.SetBytes(sig[32:64])
	return v, r, s
}
