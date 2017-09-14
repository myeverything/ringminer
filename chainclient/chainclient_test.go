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

package chainclient_test

import (
	"testing"
	"github.com/Loopring/ringminer/chainclient/eth"
	"github.com/Loopring/ringminer/types"
	"github.com/ethereum/go-ethereum/common"
	"encoding/json"
	"github.com/Loopring/ringminer/chainclient"
	"reflect"
)

func TestChainClient(t *testing.T) {

	var amount types.HexNumber
	eth.EthClient.GetBalance(&amount, "0xd86ee51b02c5ac295e59711f4335fed9805c0148", "pending")
	t.Log(amount.BigInt().String())

	var accounts []string
	eth.EthClient.Accounts(&accounts)
	t.Log(accounts[0])

}

func TestSubscribeNewBlock(t *testing.T) {
	var filterId string
	if err := eth.EthClient.NewBlockFilter(&filterId); nil != err {
		t.Error(err.Error())
	} else {
		t.Log(filterId)
	}
	hashChan := make(chan []string)

	if err := eth.EthClient.Subscribe(&hashChan, filterId);nil != err {
		t.Error(err.Error())
	} else {

		for {
			select {
			case blockHashes := <- hashChan:
				if len(blockHashes) > 0 {
					t.Log("len:", len(blockHashes), "first:",blockHashes[0])
				} else {
					t.Log("len:", len(blockHashes))
				}
			}
		}
	}

}

func TestErc20Transfer(t *testing.T) {
	//contractAddress := "0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520"
	//erc20 := eth.NewErc20Token(contractAddress)
	//
	//if txHash, err := erc20.Transfer.SendTransactionWithSpecificGas("0x4ec94e1007605d70a86279370ec5e4b755295eda",
	//	nil,
	//	nil,
	//	common.HexToAddress("0xd86ee51b02c5ac295e59711f4335fed9805c0148"),
	//	big.NewInt(10));err != nil {
	//	t.Error(err.Error())
	//} else {
	//	t.Log("txHash:", txHash)
	//}
}

func TestSubscribeErc20Event(t *testing.T) {
	var filterId string
	addresses := []common.Address{common.HexToAddress("0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520")}
	filterReq := &eth.FilterQuery{}
	filterReq.Address = addresses
	filterReq.FromBlock = "latest"
	filterReq.ToBlock = "latest"
	if err := eth.EthClient.NewFilter(&filterId, filterReq); nil != err {
		t.Log(err.Error())
	} else {
		t.Log(filterId)
	}

	//defer func() {
	//	eth.EthClient.UninstallFilter()
	//}()
	logChan := make(chan []eth.Log)
	if err := eth.EthClient.Subscribe(&logChan, filterId);nil != err {
		t.Error(err.Error())
	} else {
		for {
			select {
			case logs := <- logChan:
				if len(logs) > 0 {
					//println("len:", len(logs), "first:",logs[0])
					for _,log := range logs {
						logBytes,_ := json.Marshal(log)
						println(string(logBytes))
						println("blockNumber:", log.BlockNumber.Int64()," blockHash:", log.BlockHash)
					}
				} else {
					//println("len:", len(logs))
				}
			}
		}
	}
}


func TestNewContract(t *testing.T) {

	//var c *chainclient.Contract
	c := &chainclient.Erc20Token{}

	erc20TokenAbiStr := `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"totalSupply","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`
	eth.NewContract(c, "0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520", erc20TokenAbiStr)
	var amount types.HexNumber
	err := c.BalanceOf.Call(&amount, "pending", common.HexToAddress("0xd86ee51b02c5ac295e59711f4335fed9805c0148"))
	if err != nil {
		println(err.Error())
	}
	println(amount.Int64())
}