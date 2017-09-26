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

package types_test

import (
	"github.com/Loopring/ringminer/types"
	"testing"
)

func TestStringToAddress(t *testing.T) {
	str := "0xb794f5ea0ba39494ce839613fffba74279579268"
	add := types.StringToAddress(str)
	t.Log(add.Str())
	t.Log(len("0x08935625ce172eb3c6561404c09f130268808d08ba59dda70aefa0016619acbc"))
}

func TestHash(t *testing.T) {
	s := "0x093e56de3901764da17fef7e89f016cfdd1a88b98b1f8e3d2ebda4aff2343380"
	h := types.HexToHash(s)
	t.Log(h.Hex())
}

func TestAddress(t *testing.T) {
	s := "0xc184dd351f215f689f481c329916bb33d8df8ced"
	addr := types.HexToAddress(s)
	//addr := &types.Address{}
	//addr.SetBytes(types.Hex2Bytes(s))
	t.Log(addr.Hex())
}