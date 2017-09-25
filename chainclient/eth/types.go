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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Block struct {
	Number           hexutil.Big
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	ReceiptsRoot     string
	Miner            string
	Difficulty       hexutil.Big
	TotalDifficulty  hexutil.Big
	ExtraData        string
	Size             hexutil.Big
	GasLimit         hexutil.Big
	GasUsed          hexutil.Big
	Timestamp        hexutil.Big
	Uncles           []string
}

type BlockWithTxObject struct {
	Block
	Transactions []Transaction
}

type BlockWithTxHash struct {
	Block
	Transactions []string
}

type Transaction struct {
	Hash             string
	Nonce            hexutil.Big
	BlockHash        string
	BlockNumber      hexutil.Big
	TransactionIndex hexutil.Big
	From             string
	To               string
	Value            hexutil.Big
	GasPrice         hexutil.Big
	Gas              hexutil.Big
	Input            string
}

type Log struct {
	LogIndex         types.HexNumber `json:"logIndex"`
	BlockNumber      types.HexNumber `json:"blockNumber"`
	BlockHash        string          `json:"blockHash"`
	TransactionHash  string          `json:"transactionHash"`
	TransactionIndex types.HexNumber `json:"transactionIndex"`
	Address          string          `json:"address"`
	Data             string          `json:"data"`
	Topics           []string        `json:"topics"`
}

type FilterQuery struct {
	FromBlock string           `json:"fromBlock"`
	ToBlock   string           `json:"toBlock"`
	Address   []common.Address `json:"addresse"`
	Topics    [][]common.Hash  `json:"topics"`
}

type LogParameter struct {
	Topics []string
}
