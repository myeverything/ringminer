package types_test

import (
	"testing"
	"github.com/Loopring/ringminer/types"
	"math/big"
	"time"
)

func TestOrder_MarshalJson(t *testing.T) {
	var (
		ord types.Order
		odw types.OrderWrap
	)

	ord.Protocol = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.Owner = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.OutToken = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.InToken = types.StringToAddress("0xb794f5ea0ba39494ce839613fffba74279579268")
	ord.OutAmount = big.NewInt(20000)
	ord.InAmount = big.NewInt(800)
	ord.Expiration = uint64(time.Now().Unix())
	ord.Fee = big.NewInt(30)
	ord.SavingShare = big.NewInt(51)
	ord.V = 8
	ord.R = types.StringToSign("hhhhhhhh")
	ord.S = types.StringToSign("fjalskdf")

	odw.RawOrder = &ord
	odw.InAmount = big.NewInt(400)
	odw.OutAmount = big.NewInt(10000)
	odw.Fee = big.NewInt(15)
	odw.PeerId = "Qme85LtECPhvx4Px5i7s2Ht2dXdHrgXYpqkDsKvxdpFQP4"

	if data, err := odw.MarshalJson(); err != nil {
		t.Log(err.Error())
	} else {
		t.Log(string(data))
	}
}

func TestOrder_UnMarshalJson(t *testing.T) {
	input := `{"rawOrder":{"Id":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"Protocol":[48,98,97,51,57,52,57,52,99,101,56,51,57,54,49,51,102,102,102,98,97,55,52,50,55,57,53,55,57,50,54,56],"Owner":[48,98,97,51,57,52,57,52,99,101,56,51,57,54,49,51,102,102,102,98,97,55,52,50,55,57,53,55,57,50,54,56],"OutToken":[48,98,97,51,57,52,57,52,99,101,56,51,57,54,49,51,102,102,102,98,97,55,52,50,55,57,53,55,57,50,54,56],"InToken":[48,98,97,51,57,52,57,52,99,101,56,51,57,54,49,51,102,102,102,98,97,55,52,50,55,57,53,55,57,50,54,56],"OutAmount":20000,"InAmount":800,"Expiration":1504085573,"Fee":30,"SavingShare":51,"V":8,"R":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,104,104,104,104,104,104,104,104],"S":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,102,106,97,108,115,107,100,102]},"peerId":"Qme85LtECPhvx4Px5i7s2Ht2dXdHrgXYpqkDsKvxdpFQP4","outAmount":10000,"inAmount":400,"fee":15}`
	var odw types.OrderWrap
	if err := odw.UnMarshalJson([]byte(input)); err != nil {
		t.Log(err.Error())
	} else {
		t.Log(odw.PeerId)
		t.Log(odw.Fee)
		t.Log(odw.InAmount)
		t.Log(odw.OutAmount)
		t.Log(odw.RawOrder.OutAmount)
		t.Log(odw.RawOrder.InAmount)
		t.Log(odw.RawOrder.InToken.Str())
		t.Log(odw.RawOrder.OutToken.Str())
		t.Log(odw.RawOrder.Owner.Str())
	}

}