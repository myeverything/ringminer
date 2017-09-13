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
	"github.com/Loopring/ringminer/lrcdb"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/chainclient"
	"encoding/json"
	"os"
	"path/filepath"
	"github.com/Loopring/ringminer/chainclient/eth"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

//保存ring，并将ring发送到区块链，同样需要分为待完成和已完成
type RingClient struct {
	store lrcdb.Database

	submitedRingsStore lrcdb.Database

	unSubmitedRingsStore lrcdb.Database

	fingerprintChan chan *chainclient.FingerprintEvent

	//ring 的失败包括：提交失败，ring的合约执行时失败，执行时包括：gas不足，以及其他失败
	ringSubmitFailedChans []RingSubmitFailedChan

	stopChan chan bool

	mtx	*sync.RWMutex
}

func NewRingClient() *RingClient {
	ringClient := &RingClient{}
	ringClient.store = getdb()
	ringClient.unSubmitedRingsStore = lrcdb.NewTable(ringClient.store, "unsubmited")
	ringClient.submitedRingsStore = lrcdb.NewTable(ringClient.store, "submited")
	ringClient.mtx = &sync.RWMutex{}
	ringClient.ringSubmitFailedChans = make([]RingSubmitFailedChan,0)
	return ringClient
}

func (ringClient *RingClient) AddRingSubmitFailedChan(c RingSubmitFailedChan) {
	ringClient.mtx.Lock()
	defer ringClient.mtx.Unlock()
	ringClient.ringSubmitFailedChans = append(ringClient.ringSubmitFailedChans, c)
}

func (ringClient *RingClient) DeleteRingSubmitFailedChan(c RingSubmitFailedChan) {
	ringClient.mtx.Lock()
	defer ringClient.mtx.Unlock()

	chans := make([]RingSubmitFailedChan, 0)
	for _, v := range ringClient.ringSubmitFailedChans {
		if v != c {

			chans = append(chans, v)
		}
	}
	ringClient.ringSubmitFailedChans = chans
}

func (ringClient *RingClient) NewRing(ring *types.RingState) {
	ringClient.mtx.Lock()
	defer ringClient.mtx.Unlock()

	if (canSubmit(ring)) {
		//todo:save
		if ringBytes,err := json.Marshal(ring); err == nil {
			ringClient.unSubmitedRingsStore.Put(ring.Hash.Bytes(), ringBytes)
			println("ringHash:", ring.Hash.Hex())
			//todo:async send to block chain
			ringClient.sendRingFingerprint(ring)
		} else {
			println(err.Error())
		}
	}
}

func canSubmit(ring *types.RingState) bool {
	return true;
}

//send Fingerprint to block chain
func (ringClient *RingClient) sendRingFingerprint(ring *types.RingState) {
	//contractAddress := ring.RawRing.Orders[0].OrderState.RawOrder.Protocol
	//_, err := loopring.LoopringFingerprints[contractAddress].SubmitRingFingerprint.SendTransaction("",nil,nil,"")
	//if err != nil {
	//	println(err.Error())
	//}
}

//listen fingerprint  accept by chain and then send Ring to block chain
func (ringClient *RingClient) listenFingerprintSucessAndSendRing() {
	var filterId string
	addresses := []common.Address{common.HexToAddress("0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520")}
	filterReq := &eth.FilterQuery{}
	filterReq.Address = addresses
	filterReq.FromBlock = "latest"
	filterReq.ToBlock = "latest"
	if err := eth.EthClient.NewFilter(&filterId, filterReq); nil != err {
		println(err.Error())
	} else {
		println(filterId)
	}
	//todo：Uninstall this filterId when stop
	defer func() {
		var a string
		eth.EthClient.UninstallFilter(&a, filterId)
	}()

	logChan := make(chan []eth.Log)
	if err := eth.EthClient.Subscribe(&logChan, filterId);nil != err {
		println(err.Error())
	} else {
		for {
			select {
			case logs := <-logChan:
				for _, log := range logs {
					ringHash := []byte(log.TransactionHash)
					if _, err := ringClient.store.Get(ringHash); err == nil {
						ring := &types.RingState{}
						contractAddress := ring.RawRing.Orders[0].OrderState.RawOrder.Protocol
						//todo:发送到区块链
						_, err1 := Loopring.LoopringImpls[contractAddress].SubmitRing.SendTransactionWithSpecificGas("", nil, nil, "")
						if err1 != nil {
							println(err1.Error())
						} else {
							//标记为已删除,迁移到已完成的列表中
							ringClient.unSubmitedRingsStore.Delete(ringHash)
							//submitedRingsStore.Put(ringHash, ring.MarshalJSON())
						}
					} else {
						println(err.Error())
					}
				}
			case stop := <-ringClient.stopChan:
				if stop {
					break
				}
			}
		}
	}
}

//recover after restart
func (ringClient *RingClient) recover() {

	//iterator := unSubmitedRingsStore.NewIterator()
	//if (iterator.Next()) {
	//	keyBytes := iterator.Key()
	//	valueBytes := iterator.Value()
	//	println("key:",string(keyBytes)," value:", string(valueBytes))
	//}

	//todo: Traversal the uncompelete rings

	//hash := &types.Hash{}
	//hash.SetBytes([]byte("testtesthash"))
	//if ringBytes,err := unSubmitedRingsStore.Get(hash.Bytes());err == nil {
	//	ring := &types.RingState{}
	//	if err := json.Unmarshal(ringBytes, ring); err != nil {
	//		println(err.Error())
	//	} else {
	//		contractAddress := ring.RawRing.Orders[0].OrderState.RawOrder.Protocol
	//		var isSubmitFingerprint bool
	//		var isSubmitRing bool
	//		if err := loopring.LoopringFingerprints[contractAddress].FingerprintFound.Call(&isSubmitFingerprint, "", ""); err == nil {
	//			if (isSubmitFingerprint) {
	//				//todo:sendTransaction, check have ring been submited.
	//				if err := loopring.LoopringImpls[contractAddress].SettleRing.Call(&isSubmitRing, "", ""); err == nil {
	//					if (!isSubmitRing && canSubmit(ring)) {
	//						//loopring.LoopringImpls[contractAddress].SubmitRing.SendTransaction(contractAddress, "", "")
	//					}
	//				} else {
	//					println(err.Error())
	//				}
	//			} else {
	//				NewRing(ring)
	//			}
	//		} else {
	//			println(err.Error())
	//		}
	//	}
	//} else {
	//	println(err.Error())
	//}
}

func (ringClient *RingClient) Start() {

	recover()
	//go listenFingerprintSucessAndSendRing();

}

func file() string {
	gopath := os.Getenv("GOPATH")
	proj := "github.com/Loopring/ringminer"
	return gopath + string(filepath.Separator) + "src" + string(filepath.Separator) + proj + string(filepath.Separator) + "ldb"
}

func getdb() lrcdb.Database {
	return lrcdb.NewDB(file(), 12,12)
}
