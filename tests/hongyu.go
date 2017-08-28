package tests


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

func DebugMatch() {
	//	abiStr := `
	//[
	//	{ "type" : "function", "name" : "balance", "constant" : true ,"outputs": [ { "type": "uint256" } ]},
	//	{ "type" : "function", "name" : "send", "constant" : false, "inputs" : [ { "name" : "amount", "type" : "uint256" } ] },
	//	{ "type" : "function", "name" : "test", "constant" : false, "inputs" : [ { "name" : "number", "type" : "uint32" } ] },
	//	{ "type" : "function", "name" : "string", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "string" } ] },
	//	{ "type" : "function", "name" : "bool", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "bool" } ] },
	//	{ "type" : "function", "name" : "address", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "address" } ] },
	//	{ "type" : "function", "name" : "uint64[2]", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "uint64[2]" } ] },
	//	{ "type" : "function", "name" : "uint64[]", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "uint64[]" } ] },
	//	{ "type" : "function", "name" : "foo", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "uint32" } ] },
	//	{ "type" : "function", "name" : "bar", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "uint32" }, { "name" : "string", "type" : "uint16" } ] },
	//	{ "type" : "function", "name" : "slice", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "uint32[2]" } ] },
	//	{ "type" : "function", "name" : "slice256", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "uint256[2]" } ] },
	//	{ "type" : "function", "name" : "sliceAddress", "constant" : false, "inputs" : [ { "name" : "inputs", "type" : "address[]" } ] },
	//	{ "type" : "function", "name" : "sliceMultiAddress", "constant" : false, "inputs" : [ { "name" : "a", "type" : "address[]" }, { "name" : "b", "type" : "address[]" } ] }
	//]`
	//	a := &abi.ABI{}
	//	a.UnmarshalJSON([]byte(abiStr))
	//
	//	//i := big.Int{10}
	//	//input1 := &i
	//	//bytes1, err := a.Pack("foo", uint32(10))
	//	bytes1, err := a.Pack("send", big.NewInt(10))
	//	////bytes1 = common.Hex2Bytes("2a")
	//	////println(string(bytes1))
	//	if (nil != err ) {
	//		println(err.Error())
	//	}
	//	println((common.ToHex(bytes1)))
	//	//println(a.Methods["foo"].Name)
	//	var i *big.Int
	//	err = a.Unpack(&i, "balance", common.Hex2Bytes("a52c101e000000000000000000000000000000000000000000000000000000000000000a"))
	//	if err != nil {
	//		println(err.Error())
	//	}
	//
	//	println(i.Int64())


	//
	//b := &B{}
	//b.test()
	//
	//b.C1 = "aa"
	//b.test()
	//var a *A
	//a = &b.A
	//a.C1 = "bbb"
	//a.test()
	//b.test()

	//client, _ := rpc.Dial("http://127.0.0.1:8545")
	//eth.RPCClient = client
	//
	////eth.Client = eth.NewClient()
	////var account string
	////err := eth.Client.NewAccount(&account, "a")
	////if (nil != err) {
	////	println(err.Error())
	////}
	////println(account)
	//
	//var n rpc.HexNumber
	//err := eth.Client.BlockNumber(&n)
	//if (nil != err) {
	//	println(err.Error())
	//}
	//println(n.Int64())
	//l := &eth.EthListener{}
	//l.Start()

	//
	//erc20 := eth.NewErc20Token("0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520")
	//
	//var amount rpc.HexNumber
	//erc20.BalanceOf.Call(&amount, "latest", common.HexToAddress("0xd86ee51b02c5ac295e59711f4335fed9805c0148"))
	//println(amount.Int())
	//var accounts []string
	//eth.Client.Accounts(&accounts)
	//println(accounts[0])

	//println(erc20.BalanceOf.Name)
	//erc20 := &eth.Erc20Token{}
	//
	//abiMethod := &eth.AbiMethod{}
	//abiMethod.Name = "transfer"
	//e := reflect.ValueOf(erc20).Elem()
	//println(e.NumField())
	//
	////match := func(s string) bool {
	////	return s == "Transfer";
	////}
	//
	//name := strings.ToUpper(abiMethod.Name[0:1]) + abiMethod.Name[1:]
	//println(name)
	//v := e.FieldByName(name)
	//v.Set(reflect.ValueOf(abiMethod))
	//erc20.Transfer.Call()




	//client, _ := rpc.Dial("http://127.0.0.1:8545")
	//
	////ethClient := &eth.EthClient{}
	////ethClient.Client = *client
	////ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	////defer cancel()
	////lastBlockChain := make(chan string)
	////
	////ethClient.EthSubscribe(ctx, lastBlockChain, "newBlocks")
	////
	////select {
	////case s := <- lastBlockChain:
	////	println(s)
	////}
	//
	////arg := map[string]interface{}{
	////	"from": "0xd86ee51b02c5ac295e59711f4335fed9805c0148",
	////	"to":   "0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520",
	////	"data": "0x70a08231000000000000000000000000d86ee51b02c5ac295e59711f4335fed9805c0148",
	////}
	//
	//arg := &eth.CallArgs{}
	//arg.From = "0xd86ee51b02c5ac295e59711f4335fed9805c0148"
	//arg.To = "0x211c9fb2c5ad60a31587a4a11b289e37ed3ea520"
	//arg.Data = "0x70a08231000000000000000000000000d86ee51b02c5ac295e59711f4335fed9805c0148"
	//var lastBlock rpc.HexNumber
	//if err := client.Call(&lastBlock, "eth_call", arg, "latest"); err != nil {
	//	fmt.Println("can't get latest block:", err)
	//	return
	//}
	//println(lastBlock.Int())
	//println(lastBlock.Number.Int())
	//
	//println(lastBlock.Hash)
	//println(lastBlock.Difficulty.Int())
	//println(lastBlock.Transactions[0].Input)

	//var res string
	//s := []byte("transfer")
	//
	//topics := []string{"0x000000000000000000000000000000000000000000000000"+common.Bytes2Hex(s)}
	//println(topics[0])
	//logParameter := &eth.LogParameter{}
	//logParameter.Topics = topics
	//client.Call(&res, "eth_newFilter", nil)
	////res = "0x87682897af4358ffb7f0010213c180cb" //blockfilter
	////res = "0x481cc73339591ea230b06921fbf0ea47" //topic设置为：transfer
	//res = "0xd8c7a9164c78d3715bca025ffbd83b0" // nil
	//println(res)
	//
	//for {
	//	select {
	//	case <- time.Tick(100000):
	//		var logRes []eth.Log
	//		client.Call(&logRes, "eth_getFilterChanges", res)
	//		for _,log := range logRes {
	//			println(log.Address)
	//		}
	//	}
	//}






	//erc20.BalanceOf.Call(erc20.Abi, )

	//
	//	i := 10
	//	var ai AI
	//	ai = AI(i)
	//	println (ai)
	//switch *auth {
	//case "dylenfu":
	//	DebugOrderBook()
	//
	//case "hongyu":
	//	DebugMatch()
	//}
}