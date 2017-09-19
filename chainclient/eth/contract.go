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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"reflect"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	types "github.com/Loopring/ringminer/types"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"qiniupkg.com/x/errors.v7"
	"github.com/Loopring/ringminer/log"
)

type AbiMethod struct {
	abi.Method
	Abi *abi.ABI
	Address string
}

func (m *AbiMethod) Call(result interface{}, blockParameter string, args ...interface{}) error {
	dataBytes, err := m.Abi.Pack(m.Name, args...)

	if (nil != err) {
		return err
	}
	data := common.ToHex(dataBytes)
	//when call a contract method，gas,gasPrice and value are not needed.
	arg := &CallArg{}
	arg.From = m.Address
	arg.To = m.Address	//when call a contract method this arg is unnecessary.
	arg.Data = data
	//todo:m.Abi.Unpack
	return EthClient.Call(result, arg, blockParameter)
}

//contract transaction
func (m *AbiMethod) SendTransaction(from string, args ...interface{}) (string,error) {
	var gas, gasPrice *types.HexNumber
	dataBytes, err := m.Abi.Pack(m.Name, args...)

	if (nil != err) {
		return "", err
	}

	if err = EthClient.GasPrice(&gasPrice); nil != err {
		return "", err
	}

	dataHex := common.ToHex(dataBytes)
	callArg := &CallArg{}
	callArg.From = from
	callArg.To = m.Address
	callArg.Data = dataHex
	callArg.GasPrice = hexutil.Big(*gasPrice)
	if err = EthClient.EstimateGas(&gas, callArg); nil != err {
		return "", err
	}

	//todo: m.Abi.Pack is double used
	return m.SendTransactionWithSpecificGas(from, gas.BigInt(), gasPrice.BigInt(), args...)
}

func (m *AbiMethod) SendTransactionWithSpecificGas(from string, gas, gasPrice *big.Int, args ...interface{}) (string,error) {
	dataBytes, err := m.Abi.Pack(m.Name, args...)

	if (nil != err) {
		return "", err
	}

	if nil == gasPrice || gasPrice.Cmp(big.NewInt(0)) <= 0 {
		return "", errors.New("gasPrice must be setted.")
	}

	if nil == gas || gas.Cmp(big.NewInt(0)) <= 0 {
		return "", errors.New("gas must be setted.")
	}

	var nonce types.HexNumber
	if err = EthClient.GetTransactionCount(&nonce, from, "pending"); nil != err {
		return "", err
	}

	transaction := ethTypes.NewTransaction(nonce.Uint64(),
		common.HexToAddress(m.Address),
		big.NewInt(0),
		gas,
		gasPrice,
		dataBytes)
	var txHash string

	err = EthClient.SignAndSendTransaction(&txHash, from, transaction)
	return txHash, err
}

func applyAbiMethod(e reflect.Value, cabi *abi.ABI, address string) {
	for _, method := range cabi.Methods {
		methodName := strings.ToUpper(method.Name[0:1]) + method.Name[1:]
		abiMethod := &AbiMethod{}
		abiMethod.Name = method.Name
		abiMethod.Abi = cabi
		abiMethod.Address = address
		e.FieldByName(methodName).Set(reflect.ValueOf(abiMethod))
	}
}

func NewContract(contract interface{}, address, abiStr string  ) error {

	cabi := &abi.ABI{}
	if err := cabi.UnmarshalJSON([]byte(abiStr)); err != nil {
		log.Fatalf("error:%s", err.Error())
	}

	e := reflect.ValueOf(contract).Elem()

	e.FieldByName("Abi").Set(reflect.ValueOf(cabi))
	e.FieldByName("Address").Set(reflect.ValueOf(address))

	applyAbiMethod(e, cabi, address)

	return nil
}

