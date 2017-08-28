package eth

import (
	"github.com/ethereum/go-ethereum/rpc"
)

//通用数据结构
type Block struct {
	Number	rpc.HexNumber
	Hash	string
	ParentHash	string
	Nonce	string
	Sha3Uncles	string
	LogsBloom	string
	TransactionsRoot	string
	ReceiptsRoot	string
	Miner	string
	Difficulty	rpc.HexNumber
	TotalDifficulty	rpc.HexNumber
	ExtraData	string
	Size	rpc.HexNumber
	GasLimit	rpc.HexNumber
	GasUsed	rpc.HexNumber
	Timestamp	rpc.HexNumber
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
	Nonce	rpc.HexNumber
	BlockHash	string
	BlockNumber	rpc.HexNumber
	TransactionIndex	rpc.HexNumber
	From	string
	To	string
	Value	rpc.HexNumber
	GasPrice	rpc.HexNumber
	Gas	rpc.HexNumber
	Input	string
}

type Log struct {
	LogIndex rpc.HexNumber
	BlockNumber rpc.HexNumber
	BlockHash	string
	TransactionHash	string
	Address	string
}

type LogParameter struct {
	Topics	[]string
}
