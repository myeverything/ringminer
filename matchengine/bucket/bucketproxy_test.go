package bucket_test

import (
	"math/big"
	"strconv"
	"github.com/Loopring/ringminer/types"
	"testing"
	"github.com/Loopring/ringminer/matchengine/bucket"
	"sync"
	"time"
)

func newOrder(outToken string, inToken string, outAmount, inAmount int64, buyFirstEnough bool, idx int) *types.OrderState {
	orderState := &types.OrderState{}
	order := &types.Order{}

	outAddress := &types.Address{}
	outAddress.SetBytes([]byte(outToken))
	inAddress := &types.Address{}
	inAddress.SetBytes([]byte(inToken))

	order.TokenS = *outAddress
	order.TokenB = *inAddress
	order.AmountS = big.NewInt(outAmount)
	order.AmountB = big.NewInt(inAmount)
	order.BuyNoMoreThanAmountB = buyFirstEnough
	order.LrcFee = big.NewInt(1000)
	h := &types.Hash{}
	h.SetBytes([]byte(strconv.Itoa(idx)))
	orderState.RawOrder = order
	orderState.OrderHash = *h
	return orderState
}

type A struct {
	mtx *sync.RWMutex
}

func (a *A) t() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	println("dddddd")
	a.t1()
}
func (a *A) t1() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	println("eeeee")

}
func TestBucketProxy(t *testing.T) {
	//a := &A{mtx:&sync.RWMutex{}}
	//a.t()
	proxy := bucket.NewBucketProxy()
	//
	go proxy.Start()
	//
	order1 := newOrder("token1", "token2", 20000, 30000, true, 1)

	proxy.NewOrder(order1)

	order2 := newOrder("token2", "token3", 30000, 30000,true,  2)
	proxy.NewOrder(order2)
	////
	order3 := newOrder("token3", "token1", 40000, 20000,true,  3)
	proxy.NewOrder(order3)
	//
	//order4 := newOrder("token4", "token1", 2, 2, 3)
	//proxy.NewOrder(order4)

	//a := big.NewInt(101)
	//f := big.NewInt(10)
	//a.Mul(a, f)
	//println(a.Div(a, big.NewInt(100)).Int64())
	time.Sleep(100000)
	//for i:=0; i<10000; i++ {
	//	outToken := strconv.Itoa(rand.Intn(10))
	//	inToken := strconv.Itoa(rand.Intn(10))
	//	if (outToken == inToken) {
	//		inToken = strconv.Itoa(rand.Intn(10))
	//	}
	//	//log.Info("info", "info", " idx:" + strconv.Itoa(i))
	//	order := newOrder(outToken, inToken, 100, 10000, i)
	//	proxy.NewOrder(order)
	//}

}