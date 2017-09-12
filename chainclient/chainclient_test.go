package chainclient_test

import (
	"testing"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/Loopring/ringminer/chainclient/eth"
	"github.com/Loopring/ringminer/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type RpcMethod string

func (m *RpcMethod) Run(result interface{}, args ...interface{}) error {
	println("aaaaa")
	println(*m)
	return nil
}

type A struct {
	C1 string
	F RpcMethod
	D func(name string) func(test string) error
}

type B struct {
	A
}

type AI int

func f1(a string) func(b string) error {
	f := func(b string) error {
		println("aaaa")
		return nil
	}
	return f
}

func TestMatchEngine(t *testing.T) {
	println("TestMatchEngine")
	//proxy := bucket.NewBucketProxy()
	////
	//go proxy.Start()
	//
	//order1 := newOrder("token1", "token2", 100, 10000)
	//
	//proxy.NewOrder(order1)
	//
	//order2 := newOrder("token2", "token3", 100, 10000)
	//proxy.NewOrder(order2)
	////
	//order3 := newOrder("token3", "token1", 100, 10000)
	//proxy.NewOrder(order3)

	//log.NewLogger()
	//for i:=0; i<1000; i++ {
	//	outToken := strconv.Itoa(rand.Intn(100))
	//	inToken := strconv.Itoa(rand.Intn(100))
	//	//log.Info("info", "info", " idx:" + strconv.Itoa(i))
	//	order := newOrder(outToken, inToken, 100, 10000, i)
	//	proxy.NewOrder(order)
	//}

	//time.Sleep(100000)

	//proxy.Stop()

	//
	//order4 := newOrder("token1", "token4", 100, 10000)
	//proxy.NewOrder(order4)

}

//func newOrder(outToken string, inToken string, inAmount, outAmount int64, idx int) *types.Order {
//	order := &types.Order{}
//
//	outAddress := &types.Address{}
//	outAddress.SetBytes([]byte(outToken))
//	inAddress := &types.Address{}
//	inAddress.SetBytes([]byte(inToken))
//
//	order.OutToken = *outAddress
//	order.InToken = *inAddress
//	order.OutAmount = big.NewInt(outAmount)
//	order.InAmount = big.NewInt(inAmount)
//	h := &types.Hash{}
//	h.SetBytes([]byte(strconv.Itoa(idx)))
//	order.Id = *h
//	return order
//}

func TestChainClient(t *testing.T) {

	client, _ := rpc.Dial("http://127.0.0.1:8545")
	eth.RPCClient = client
	var amount types.HexNumber
	contractAddress := "0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520"

	//eth.EthClient.GetBalance(&amount, "0xd86ee51b02c5ac295e59711f4335fed9805c0148", "pending")
	erc20 := eth.NewErc20Token(contractAddress)
	err := erc20.BalanceOf.Call(&amount, "pending", common.HexToAddress("0xd86ee51b02c5ac295e59711f4335fed9805c0148"))
	if err != nil {
		println(err.Error())
	}
	//println(amount.BigInt().String())

	if txHash, err := erc20.Transfer.SendTransaction("0x4ec94e1007605d70a86279370ec5e4b755295eda",
		nil,
		nil,
		common.HexToAddress("0xd86ee51b02c5ac295e59711f4335fed9805c0148"),
		big.NewInt(10));err != nil {
		println(err.Error())
	} else {
		println("txHash:", txHash)
	}
	//var accounts []string
	//eth.EthClient.Accounts(&accounts)
	//println(accounts[0])

}