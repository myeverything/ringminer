package eth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//通用数据结构
type Block struct {
	Number	hexutil.Big
	Hash	string
	ParentHash	string
	Nonce	string
	Sha3Uncles	string
	LogsBloom	string
	TransactionsRoot	string
	ReceiptsRoot	string
	Miner	string
	Difficulty	hexutil.Big
	TotalDifficulty	hexutil.Big
	ExtraData	string
	Size	hexutil.Big
	GasLimit	hexutil.Big
	GasUsed	hexutil.Big
	Timestamp	hexutil.Big
	Uncles	[]string
}

type BlockWithTxObject struct {
	Block
	Transactions	[]Transaction
}

type BlockWithTxHash struct {
	Block
	Transactions	[]string
}

type Transaction struct {
	Hash	string
	Nonce	hexutil.Big
	BlockHash	string
	BlockNumber	hexutil.Big
	TransactionIndex	hexutil.Big
	From	string
	To	string
	Value	hexutil.Big
	GasPrice	hexutil.Big
	Gas	hexutil.Big
	Input	string
}

type Log struct {
	LogIndex hexutil.Big
	BlockNumber hexutil.Big
	BlockHash	string
	TransactionHash	string
	Address	string
}

type LogParameter struct {
	Topics	[]string
}
