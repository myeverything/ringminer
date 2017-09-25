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

package orderbook

import (
	"github.com/Loopring/ringminer/config"
	"github.com/Loopring/ringminer/db"
	"github.com/Loopring/ringminer/log"
	"github.com/Loopring/ringminer/types"
	"sync"
)

type ORDER_STATUS int

const (
	FINISH_TABLE_NAME  = "finished"
	PARTIAL_TABLE_NAME = "partial"
)

type Whisper struct {
	PeerOrderChan   chan *types.Order
	EngineOrderChan chan *types.OrderState
	ChainOrderChan  chan *types.OrderMined
}

type OrderBook struct {
	toml         config.DbOptions
	db           db.Database
	finishTable  db.Database
	partialTable db.Database
	whisper      *Whisper
	lock         sync.RWMutex
}

func NewOrderBook(database db.Database, whisper *Whisper) *OrderBook {
	s := &OrderBook{}

	s.finishTable = db.NewTable(database, FINISH_TABLE_NAME)
	s.partialTable = db.NewTable(database, PARTIAL_TABLE_NAME)
	s.whisper = whisper

	return s
}

// Start start orderbook as a service
func (s *OrderBook) Start() {
	go func() {
		for {
			select {
			case ord := <-s.whisper.PeerOrderChan:
				log.Debugf("accept data from peer:%s", ord.Protocol.Hex())
				s.peerOrderHook(ord)
			case ord := <-s.whisper.ChainOrderChan:
				s.chainOrderHook(ord)
			}
		}
	}()
}

func (s *OrderBook) Stop() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.finishTable.Close()
	s.partialTable.Close()
	s.db.Close()
}

func (ob *OrderBook) peerOrderHook(ord *types.Order) error {

	// TODO(fk): order filtering

	ob.lock.Lock()
	defer ob.lock.Unlock()

	//todo:apologize for it
	//key := ord.GenHash().Bytes()
	//value,err := ord.MarshalJson()
	//if err != nil {
	//	return err
	//}
	//
	//ob.partialTable.Put(key, value)
	//
	//// TODO(fk): delete after test
	//if input, err := ob.partialTable.Get(key); err != nil {
	//	panic(err)
	//} else {
	//	var ord types.Order
	//	ord.UnMarshalJson(input)
	//	log.Println(ord.TokenS.Str())
	//	log.Println(ord.TokenB.Str())
	//	log.Println(ord.AmountS.Uint64())
	//	log.Println(ord.AmountB.Uint64())
	//}
	//
	//// TODO(fk): send orderState to matchengine

	state := ord.Convert()
	state.GenHash()

	if addr, err := state.SignerAddress(); err != nil {
		log.Errorf("err:%s", err.Error())
	} else {
		log.Debugf("addrreeseresrs:%s", addr.Hex())
	}
	log.Debugf("state hash:%s", state.OrderHash.Hex())
	ob.whisper.EngineOrderChan <- state

	return nil
}

func (ob *OrderBook) chainOrderHook(ord *types.OrderMined) error {
	ob.lock.Lock()
	defer ob.lock.Unlock()

	return nil
}

// GetOrder get single order with hash
func (ob *OrderBook) GetOrder(id types.Hash) (*types.OrderState, error) {
	var (
		value []byte
		err   error
		ord   types.OrderState
	)

	if value, err = ob.partialTable.Get(id.Bytes()); err != nil {
		value, err = ob.finishTable.Get(id.Bytes())
	}
	if err != nil {
		return nil, err
	}

	err = ord.UnMarshalJson(value)
	if err != nil {
		return nil, err
	}

	return &ord, nil
}

// GetOrders get orders from persistence database
func (ob *OrderBook) GetOrders() {

}

// moveOrder move order when partial finished order fully exchanged
func (ob *OrderBook) moveOrder(odw *types.OrderState) error {
	key := odw.OrderHash.Bytes()
	value, err := odw.MarshalJson()
	if err != nil {
		return err
	}
	ob.partialTable.Delete(key)
	ob.finishTable.Put(key, value)
	return nil
}

// isFinished judge order state
func isFinished(odw *types.OrderState) bool {
	//if odw.RawOrder.

	return true
}
