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
	"fmt"
	"math/big"
	"reflect"
)

//todo:test and fix bug (bug exists)
type EnlargedInt struct {
	Value    *big.Int
	Decimals *big.Int
}

func (ei *EnlargedInt) Div(x, y *EnlargedInt) *EnlargedInt {
	if ei.Value == nil {
		ei.Value = big.NewInt(0)
	}
	if ei.Decimals == nil {
		ei.Decimals = big.NewInt(0)
	}
	ei.Value.Div(x.Value, y.Value)
	ei.Decimals.Div(x.Decimals, y.Decimals)
	return ei
}

func (ei *EnlargedInt) DivBigInt(x *EnlargedInt, y *big.Int) *EnlargedInt {
	if ei.Value == nil {
		ei.Value = big.NewInt(1)
	}
	if ei.Decimals == nil {
		ei.Decimals = big.NewInt(1)
	}
	ei.Value.Div(x.Value, y)
	ei.Decimals = x.Decimals
	return ei
}

func (ei *EnlargedInt) Mul(x, y *EnlargedInt) {
	if ei.Value == nil {
		ei.Value = big.NewInt(1)
	}
	if ei.Decimals == nil {
		ei.Decimals = big.NewInt(1)
	}
	ei.Value.Mul(x.Value, y.Value)
	ei.Decimals.Mul(x.Decimals, y.Decimals)
}

func (ei *EnlargedInt) MulBigInt(x *EnlargedInt, y *big.Int) *EnlargedInt {
	if ei.Value == nil {
		ei.Value = big.NewInt(1)
	}
	if ei.Decimals == nil {
		ei.Decimals = big.NewInt(1)
	}
	ei.Value.Mul(x.Value, y)
	ei.Decimals = ei.Decimals.Mul(ei.Decimals, x.Decimals)
	return ei
}

func (ei *EnlargedInt) Sub(x, y *EnlargedInt) *EnlargedInt {

	if ei.Value == nil {
		ei.Value = big.NewInt(0)
	}
	if ei.Decimals == nil {
		ei.Decimals = big.NewInt(1)
	}
	if x.Decimals.Cmp(y.Decimals) == 0 {
		ei.Value.Sub(x.Value, y.Value)
		ei.Decimals = x.Decimals
	} else if x.Decimals.Cmp(y.Decimals) > 0 {
		decimals := big.NewInt(1)
		decimals.Div(x.Decimals, y.Decimals)
		value := big.NewInt(1)
		value.Mul(y.Value, decimals)

		ei.Value.Sub(x.Value, value)
		ei.Decimals = x.Decimals
	} else {
		decimals := big.NewInt(1)
		decimals.Div(y.Decimals, x.Decimals)
		value := big.NewInt(1)
		value.Mul(x.Value, decimals)

		ei.Value.Sub(value, y.Value)
		ei.Decimals = y.Decimals
	}

	return ei
}
func (ei *EnlargedInt) Add(x, y *EnlargedInt) *EnlargedInt {

	if ei.Value == nil {
		ei.Value = big.NewInt(0)
	}
	if ei.Decimals == nil {
		ei.Decimals = big.NewInt(1)
	}
	if x.Decimals.Cmp(y.Decimals) == 0 {
		ei.Value.Add(x.Value, y.Value)
		ei.Decimals = x.Decimals
	} else if x.Decimals.Cmp(y.Decimals) > 0 {
		decimals := big.NewInt(1)
		decimals.Div(x.Decimals, y.Decimals)
		value := big.NewInt(1)
		value.Mul(y.Value, decimals)

		ei.Value.Add(x.Value, value)
		ei.Decimals = x.Decimals
	} else {
		decimals := big.NewInt(1)
		decimals.Div(y.Decimals, x.Decimals)
		value := big.NewInt(1)
		value.Mul(x.Value, decimals)

		ei.Value.Add(value, y.Value)
		ei.Decimals = y.Decimals
	}

	return ei
}

func (ei *EnlargedInt) RealValue() *big.Int {
	realValue := big.NewInt(0)
	return realValue.Div(ei.Value, ei.Decimals)
}

func (ei *EnlargedInt) Cmp(x *EnlargedInt) int {
	return ei.RealValue().Cmp(x.RealValue())
}

func (ei *EnlargedInt) CmpBigInt(x *big.Int) int {
	return ei.RealValue().Cmp(x)
}

func NewEnlargedInt(value *big.Int) *EnlargedInt {
	return &EnlargedInt{Value: value, Decimals: big.NewInt(1)}
}

type HexNumber big.Int

// NewHexNumber creates a new hex number instance which will serialize the given val with `%#x` on marshal.
func NewHexNumber(val interface{}) *HexNumber {
	if val == nil {
		return nil // note, this doesn't catch nil pointers, only passing nil directly!
	}

	if v, ok := val.(*big.Int); ok {
		if v != nil {
			return (*HexNumber)(new(big.Int).Set(v))
		}
		return nil
	}

	rval := reflect.ValueOf(val)

	var unsigned uint64
	utype := reflect.TypeOf(unsigned)
	if t := rval.Type(); t.ConvertibleTo(utype) {
		hn := new(big.Int).SetUint64(rval.Convert(utype).Uint())
		return (*HexNumber)(hn)
	}

	var signed int64
	stype := reflect.TypeOf(signed)
	if t := rval.Type(); t.ConvertibleTo(stype) {
		hn := new(big.Int).SetInt64(rval.Convert(stype).Int())
		return (*HexNumber)(hn)
	}

	return nil
}

func (h *HexNumber) UnmarshalJSON(input []byte) error {
	length := len(input)
	if length >= 2 && input[0] == '"' && input[length-1] == '"' {
		input = input[1 : length-1]
	}

	hn := (*big.Int)(h)
	if _, ok := hn.SetString(string(input), 0); ok {
		return nil
	}

	return fmt.Errorf("Unable to parse number")
}

// MarshalJSON serialize the hex number instance to a hex representation.
func (h *HexNumber) MarshalJSON() ([]byte, error) {
	if h != nil {
		hn := (*big.Int)(h)
		if hn.BitLen() == 0 {
			return []byte(`"0x0"`), nil
		}
		return []byte(fmt.Sprintf(`"0x%x"`, hn)), nil
	}
	return nil, nil
}

func (h *HexNumber) Int() int {
	hn := (*big.Int)(h)
	return int(hn.Int64())
}

func (h *HexNumber) Int64() int64 {
	hn := (*big.Int)(h)
	return hn.Int64()
}

func (h *HexNumber) Uint() uint {
	hn := (*big.Int)(h)
	return uint(hn.Uint64())
}

func (h *HexNumber) Uint64() uint64 {
	hn := (*big.Int)(h)
	return hn.Uint64()
}

func (h *HexNumber) BigInt() *big.Int {
	return (*big.Int)(h)
}
