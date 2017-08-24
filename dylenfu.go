package main

import (
	"github.com/Loopring/ringminer/p2p/ipfs"
	"time"
	"encoding/json"
	"log"
	"math/big"
	"math"
	"reflect"
)

func DebugOrderBook(testcase string) {
	switch testcase{
	case "listen":
		ipfsListenTest()
		break

	case "json":
		jsonTest()
		break

	case "bigint":
		bigIntTest()
		break

	default:
		break
	}
}

func ipfsListenTest() {
	listener := ipfs.NewListener()
	go listener.Start()
	time.Sleep(100 * time.Second)
	listener.Stop()
}

// 构造的时候不能缺少字段，解析的时候可以
func jsonTest() {
	type JT struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Ted  string `json:ted`
	}

	//
	//t := &JT{"dylenfu", 30}
	//m,_ := json.Marshal(t)
	//log.Println("test\t-", "json marshal", string(m))

	m := `{"name":"dylenfu","age":30}`
	u := &JT{}
	json.Unmarshal([]byte(m), u)
	log.Println("test\t-", "json unmarshal", u.Name, u.Age)
}

func bigIntTest() {
	b := big.NewInt(math.MaxInt64)
	log.Println("test\t-", "initial b", b)
	n := big.NewInt(100)
	log.Println("test\t", "b multiple", b.Mul(n,b))

	log.Println(big.NewInt(46877).Uint64())

	log.Println(reflect.ValueOf(nil).IsValid())
}