package matchengine

import (
	"github.com/Loopring/ringminer/types"
	"math/big"
)

type LegalCurrency int

const (
	_ LegalCurrency = iota
	CNY
	USD
)

const (
	LRC_ADDRESS = "0x5132a8ce9a61b13b9cAEcd2261abF95323056423"
)

//todo:获取法币汇率
func GetLegalRate(currency LegalCurrency, tokenAddress types.Address) *types.EnlargedInt {
	return &types.EnlargedInt{Value:big.NewInt(100), Decimals:big.NewInt(100)}
}
