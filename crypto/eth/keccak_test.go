package eth_test

import (
	"testing"
	"github.com/Loopring/ringminer/types"
	"time"
	"github.com/Loopring/ringminer/crypto/eth"
	"strconv"
)

func TestGenOrderHash(t *testing.T) {
	var ord types.Order

	ord.Protocol = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.TokenS = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.TokenB = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.AmountS = types.IntToBig(20000)
	ord.AmountB = types.IntToBig(800)
	ord.Expiration = uint64(time.Now().Unix())
	ord.Rand = types.IntToBig(int64(3))
	ord.LrcFee = types.IntToBig(30)
	ord.SavingSharePercentage = 51
	ord.V = 8
	ord.R = types.StringToSign("hhhhhhhh")
	ord.S = types.StringToSign("fjalskdf")

	hash := eth.GenOrderHash(ord)

	t.Log(string(hash))
	t.Log(types.BytesToHash(hash).Hex())
}

func TestGenOrderAddress(t *testing.T) {
	dst := []byte{'a', 'b'}
	data := strconv.AppendUint(dst, 16, 10)
	t.Log(string(data))
	t.Log(string(dst))

	t.Log(strconv.FormatUint(16, 10))
}