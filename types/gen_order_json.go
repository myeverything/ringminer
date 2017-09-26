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

package types

import (
	"encoding/json"
	"reflect"
	"errors"
)

func (ord *Order) MarshalJson() ([]byte, error) {
	type order struct {
		Protocol              string `json:"protocol"`
		TokenS                string `json:"tokenS"`
		TokenB                string `json:"tokenB"`
		AmountS               string `json:"amountS"`
		AmountB               string `json:"amountB"`
		Rand                  string `json:"rand"`
		Expiration            string `json:"expiration"`
		LrcFee                string `json:"lrcFee"`
		SavingSharePercentage int    `json:"savingShareRate"`
		BuyNoMoreThanAmountB  bool   `json:"buyNoMoreThanAmountB"`
		V                     uint8  `json:"v"`
		R                     string `json:"r"`
		S                     string `json:"s"`
	}

	var enc order

	enc.Protocol = ord.Protocol.Hex()
	enc.TokenS = ord.TokenS.Hex()
	enc.TokenB = ord.TokenB.Hex()

	enc.AmountS = BigToHex(ord.AmountS)
	enc.AmountB = BigToHex(ord.AmountB)

	enc.Rand = BigToHex(ord.Rand)
	enc.Expiration = BigToHex(ord.Expiration)
	enc.LrcFee = BigToHex(ord.LrcFee)
	enc.SavingSharePercentage = ord.SavingSharePercentage
	enc.BuyNoMoreThanAmountB = ord.BuyNoMoreThanAmountB

	enc.V = ord.V
	enc.R = BigToHex(ord.R)
	enc.S = BigToHex(ord.S)

	return json.Marshal(enc)
}

func (ord *Order) UnMarshalJson(input []byte) error {
	type order struct {
		Protocol              string `json:"protocol"`
		TokenS                string `json:"tokenS"`
		TokenB                string `json:"tokenB"`
		AmountS               string `json:"amountS"`
		AmountB               string `json:"amountB"`
		Rand                  string `json:"rand"`
		Expiration            string `json:"expiration"`
		LrcFee                string `json:"lrcFee"`
		SavingSharePercentage int    `json:"savingShareRate"`
		BuyNoMoreThanAmountB  bool   `json:"buyNoMoreThanAmountB"`
		V                     uint8  `json:"v"`
		R                     string `json:"r"`
		S                     string `json:"s"`
	}

	var dec order
	err := json.Unmarshal(input, &dec)
	if err != nil {
		return err
	}

	if !reflect.ValueOf(dec.Protocol).IsValid() {
		return errors.New("missing required field 'Protocol' for order")
	}
	ord.Protocol = HexToAddress(dec.Protocol)

	if !reflect.ValueOf(dec.TokenS).IsValid() {
		return errors.New("missing required field 'tokenS' for order")
	}
	ord.TokenS = HexToAddress(dec.TokenS)

	if !reflect.ValueOf(dec.TokenB).IsValid() {
		return errors.New("missing required field 'tokenB' for order")
	}
	ord.TokenB = HexToAddress(dec.TokenB)

	if !reflect.ValueOf(dec.AmountS).IsValid() {
		return errors.New("missing required field 'amountS' for order")
	}
	ord.AmountS = HexToBig(dec.AmountS)

	if !reflect.ValueOf(dec.AmountB).IsValid() {
		return errors.New("missing required field 'amountB' for order")
	}
	ord.AmountB = HexToBig(dec.AmountB)

	if !reflect.ValueOf(dec.Rand).IsValid() {
		return errors.New("missing required field 'rand' for order")
	}
	ord.Rand = HexToBig(dec.Rand)

	if !reflect.ValueOf(dec.Expiration).IsValid() {
		return errors.New("missing required field 'expiration' for order")
	}
	ord.Expiration = HexToBig(dec.Expiration)

	if !reflect.ValueOf(dec.LrcFee).IsValid() {
		return errors.New("missing required field 'lrcFee' for order")
	}
	ord.LrcFee = HexToBig(dec.LrcFee)

	if !reflect.ValueOf(dec.SavingSharePercentage).IsValid() {
		return errors.New("missing required field 'savingSharePercentage' for order")
	}
	ord.SavingSharePercentage = dec.SavingSharePercentage

	if !reflect.ValueOf(dec.BuyNoMoreThanAmountB).IsValid() {
		return errors.New("missing required field 'fullyFilled' for order")
	}
	ord.BuyNoMoreThanAmountB = dec.BuyNoMoreThanAmountB

	if !reflect.ValueOf(dec.V).IsValid() {
		return errors.New("missing required field 'ECDSA.V' for order")
	}
	ord.V = dec.V

	if !reflect.ValueOf(dec.S).IsValid() {
		return errors.New("missing required field 'ECDSA.S' for order")
	}
	ord.S = HexToBig(dec.S)

	if !reflect.ValueOf(dec.R).IsValid() {
		return errors.New("missing required field 'ECSA.R' for order")
	}
	ord.R = HexToBig(dec.R)

	return nil
}

func (ord *OrderState) MarshalJson() ([]byte, error) {
	type state struct {
		Protocol              string `json:"protocol"`
		TokenS                string `json:"tokenS"`
		TokenB                string `json:"tokenB"`
		AmountS               string `json:"amountS"`
		AmountB               string `json:"amountB"`
		Rand                  string `json:"rand"`
		Expiration            string `json:"expiration"`
		LrcFee                string `json:"lrcFee"`
		SavingSharePercentage int    `json:"savingShareRate"`
		BuyNoMoreThanAmountB  bool   `json:"buyNoMoreThanAmountB"`
		V                     uint8  `json:"v"`
		R                     string `json:"r"`
		S                     string `json:"s"`
		Owner                 string `json:"owner"`
		OrderHash             string `json:"hash"`
		RemainedAmountS       string `json:"remainedAmountS"`
		RemainedAmountB       string `json:"remainedAmountB"`
		Status                uint8  `json:"status"`
	}

	var enc state

	enc.Protocol = ord.RawOrder.Protocol.Hex()
	enc.TokenS = ord.RawOrder.TokenS.Hex()
	enc.TokenB = ord.RawOrder.TokenB.Hex()

	enc.AmountS = BigToHex(ord.RawOrder.AmountS)
	enc.AmountB = BigToHex(ord.RawOrder.AmountB)

	enc.Rand = BigToHex(ord.RawOrder.Rand)
	enc.Expiration = BigToHex(ord.RawOrder.Expiration)
	enc.LrcFee = BigToHex(ord.RawOrder.LrcFee)
	enc.SavingSharePercentage = ord.RawOrder.SavingSharePercentage
	enc.BuyNoMoreThanAmountB = ord.RawOrder.BuyNoMoreThanAmountB

	enc.V = ord.RawOrder.V
	enc.R = BigToHex(ord.RawOrder.R)
	enc.S = BigToHex(ord.RawOrder.S)

	enc.Owner = ord.Owner.Hex()
	enc.OrderHash = ord.OrderHash.Hex()
	enc.RemainedAmountS = BigToHex(ord.RemainedAmountS)
	enc.RemainedAmountB = BigToHex(ord.RemainedAmountB)
	enc.Status = uint8(ord.Status)

	return json.Marshal(enc)
}

func (ord *OrderState) UnMarshalJson(input []byte) error {
	type state struct {
		Protocol              string `json:"protocol"`
		TokenS                string `json:"tokenS"`
		TokenB                string `json:"tokenB"`
		AmountS               string `json:"amountS"`
		AmountB               string `json:"amountB"`
		Rand                  string `json:"rand"`
		Expiration            string `json:"expiration"`
		LrcFee                string `json:"lrcFee"`
		SavingSharePercentage int    `json:"savingShareRate"`
		BuyNoMoreThanAmountB  bool   `json:"buyNoMoreThanAmountB"`
		V                     uint8  `json:"v"`
		R                     string `json:"r"`
		S                     string `json:"s"`
		Owner                 string `json:"owner"`
		OrderHash             string `json:"hash"`
		RemainedAmountS       string `json:"remainedAmountS"`
		RemainedAmountB       string `json:"remainedAmountB"`
		Status                uint8  `json:"status"`
	}

	var dec state
	err := json.Unmarshal(input, &dec)
	if err != nil {
		return err
	}
	//
	if !reflect.ValueOf(dec.Protocol).IsValid() {
		return errors.New("missing required field 'Protocol' for orderState")
	}
	ord.RawOrder.Protocol = HexToAddress(dec.Protocol)

	if !reflect.ValueOf(dec.TokenS).IsValid() {
		return errors.New("missing required field 'tokenS' for orderState")
	}
	ord.RawOrder.TokenS = HexToAddress(dec.TokenS)

	if !reflect.ValueOf(dec.TokenB).IsValid() {
		return errors.New("missing required field 'tokenB' for orderState")
	}
	ord.RawOrder.TokenB = HexToAddress(dec.TokenB)

	if !reflect.ValueOf(dec.AmountS).IsValid() {
		return errors.New("missing required field 'amountS' for orderState")
	}
	ord.RawOrder.AmountS = HexToBig(dec.AmountS)

	if !reflect.ValueOf(dec.AmountB).IsValid() {
		return errors.New("missing required field 'amountB' for orderState")
	}
	ord.RawOrder.AmountB = HexToBig(dec.AmountB)

	if !reflect.ValueOf(dec.Rand).IsValid() {
		return errors.New("missing required field 'rand' for orderState")
	}
	ord.RawOrder.Rand = HexToBig(dec.Rand)

	if !reflect.ValueOf(dec.Expiration).IsValid() {
		return errors.New("missing required field 'expiration' for orderState")
	}
	ord.RawOrder.Expiration = HexToBig(dec.Expiration)

	if !reflect.ValueOf(dec.LrcFee).IsValid() {
		return errors.New("missing required field 'lrcFee' for orderState")
	}
	ord.RawOrder.LrcFee = HexToBig(dec.LrcFee)

	if !reflect.ValueOf(dec.SavingSharePercentage).IsValid() {
		return errors.New("missing required field 'savingSharePercentage' for orderState")
	}
	ord.RawOrder.SavingSharePercentage = dec.SavingSharePercentage

	if !reflect.ValueOf(dec.BuyNoMoreThanAmountB).IsValid() {
		return errors.New("missing required field 'fullyFilled' for orderState")
	}
	ord.RawOrder.BuyNoMoreThanAmountB = dec.BuyNoMoreThanAmountB

	if !reflect.ValueOf(dec.V).IsValid() {
		return errors.New("missing required field 'ECDSA.V' for orderState")
	}
	ord.RawOrder.V = dec.V

	if !reflect.ValueOf(dec.S).IsValid() {
		return errors.New("missing required field 'ECDSA.S' for orderState")
	}
	ord.RawOrder.S = HexToBig(dec.S)

	if !reflect.ValueOf(dec.R).IsValid() {
		return errors.New("missing required field 'ECSA.R' for orderState")
	}
	ord.RawOrder.R = HexToBig(dec.R)

	if !reflect.ValueOf(dec.Owner).IsValid() {
		return errors.New("missing required field 'owner' for orderState")
	}
	ord.Owner = HexToAddress(dec.Owner)

	if !reflect.ValueOf(dec.OrderHash).IsValid() {
		return errors.New("missing required field 'orderHash' for orderState")
	}
	ord.OrderHash = StringToHash(dec.OrderHash)

	if !reflect.ValueOf(dec.RemainedAmountS).IsValid() {
		return errors.New("missing required field 'remainedAmountS' for orderState")
	}
	ord.RemainedAmountS = HexToBig(dec.RemainedAmountS)

	if !reflect.ValueOf(dec.RemainedAmountB).IsValid() {
		return errors.New("missing required field 'remainedAmountB' for orderState")
	}
	ord.RemainedAmountB = HexToBig(dec.RemainedAmountB)

	if !reflect.ValueOf(dec.Status).IsValid() {
		return errors.New("missing required field 'status' for orderState")
	}
	ord.Status = OrderStatus(dec.Status)

	return nil
}
