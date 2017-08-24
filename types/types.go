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
	Id          Hash    `json:"id"`         // 订单id
	Protocol    Address `json:"protocol"`   // 智能合约地址
	Owner       Address `json:"owner"`      // 订单发起者地址
	OutToken    Address `json:"outToken"`   // 卖出erc20代币智能合约地址
	InToken     Address `json:"inToken"`    // 买入erc20代币智能合约地址
	OutAmount   uint64  `json:"outAmount"`  // 卖出erc20代币数量上限
	InAmount    uint64  `json:"inAmount"`   // 买入erc20代币数量上限
	Expiration  uint64  `json:"expiration"` // 订单过期时间
	Fee         uint64  `json:"fee"`        // 交易总费用,部分成交的费用按该次撮合实际卖出代币额与比例计算
	SavingShare uint64  `json:"savingShare"`// 不为0时支付给交易所的分润比例，否则视为100%
	V           uint8   `json:"v"`
	R           Sign    `json:"r"`
	S           Sign    `json:"s"`
}

// TODO(fukun): 包含成交记录
type OrderWrap struct {
	Order             `json:"order"`
	PeerId   string   `json:"peerId"`
	RingList []Hash   `json:"ringList"`
}

// 旷工在成本节约和fee上二选一，撮合者计算出:
// 1.fee(lrc)的市场价(法币交易价格)
// 2.成本节约(savingShare)的市场价(法币交易价格)
// 撮合者在fee和savingShare中二选一，作为自己的利润，
// 如果撮合者选择fee，则成本节约分给订单发起者，如果选择成本节约，则需要返还给用户一定的lrc
// 这样一来，撮合者的利润判断公式应该是max(fee, savingShare - fee * s),s为固定比例
// 此外，在选择最优环路的时候，撮合者会在确定了选择fee/savingShare后，选择某个具有最大利润的环路
// 但是，根据谷歌竞拍法则(A出价10,B出价20,最终成交价为10)，撮合者最终获得的利润只能是利润最小的环路利润
type Ring struct {
	Id                Hash    `json:"id"`                // 订单链id
	Orders            []Order `json:"orders"`            // 该次匹配的所有订单
	FeeRecipient      Address `json:"feeRecipient"`      // 费用收取地址
	AddtionalDiscount uint64  `json:"addtionalDiscount"` // 在费用基础上的再折扣价-eta
	Nonce             uint64  `json:"nonce"`             // 一个随机数
	V                 uint8   `json:"v"`
	R                 Sign    `json:"r"`
	S                 Sign    `json:"s"`
}

// TODO(fukun): 添加状态判断是否成环
type OrderRing struct {
	Ring         `json:"ring"`     // 订单链
	Closure bool `json:"closure"`  // 是否闭合
}
