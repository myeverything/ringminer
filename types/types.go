package types

const (
	HashLength    = 20
	AddressLength = 32
	SignLength    = 32
)

// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

// Address represents the 20 byte address of an Ethereum account.
type Address [AddressLength]byte

// Sign represents the 32 byte of an ECDSA r/s
type Sign [SignLength]byte

//订单原始信息
type Order struct {
	Protocol    Address //智能合约地址
	Owner       Address //订单发起者地址
	OutToken    Address //卖出erc20代币智能合约地址
	InToken     Address //买入erc20代币智能合约地址
	OutAmount   uint64  //卖出erc20代币数量上限
	InAmount    uint64  //买入erc20代币数量上限
	Expiration  uint64  //订单过期时间
	Fee         uint64  //交易总费用,部分成交的费用按该次撮合实际卖出代币额与比例计算
	SavingShare uint64  //不为0时支付给交易所的分润比例，否则视为100%
	V           uint8
	R           Sign
	S           Sign
}

// TODO(fukun): 包含成交记录
type OrderWrap struct {
	Order
}

// TODO(fukun): 添加状态判断是否成环
type OrderRing struct {
}
