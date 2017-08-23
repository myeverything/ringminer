package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"net"
	"log"
	"github.com/Loopring/ringminer/orderbook"
	"github.com/Loopring/ringminer/store"
)

type A int
type B struct {
	A
	User string
}

type MultiplyArgs struct {
	F int
	S int
	Reply int
}

//func (A) Av() {
//	println("ddddd")
//}
//func (*A) Ap() {}
//func (B) Bv() {}
//func (*B) Bp() {}

func (b *B) Multiply(arg MultiplyArgs, replay *MultiplyArgs)  error {
	replay.Reply = arg.F * arg.S
	return nil
}

//func main() {
//	//var b B
//	//t := reflect.TypeOf(&b)
//	//
//	//s := []reflect.Type{t, t.Elem()}
//	//println(reflect.TypeOf(b) == t.Elem())
//	//println(t.NumMethod())
//	////args := []reflect.Value{}
//	////println(method.Type.In(0))
//	////println(reflect.ValueOf(b).Method(0).Call(args))
//	//for _, t2 := range s {
//	//	fmt.Println(t2, ":")
//	//	//println(t2.MethodByName("Av"))
//	//	for i := 0; i < t2.NumMethod(); i++ {
//	//		method := t2.Method(i)
//	//		method.Func.Send(reflect.ValueOf(b))
//	//		fmt.Println(" ", t2.Method(i))
//	//	}
//	//}
//
//	lis, err := net.Listen("tcp", ":1789")
//	if err != nil {
//		return
//	}
//	defer lis.Close()
//
//	srv := rpc.NewServer()
//	if err := srv.RegisterName("Json", new(B)); err != nil {
//		return
//	}
//
//	for {
//		conn, err := lis.Accept()
//		if err != nil {
//			log.Fatalf("lis.Accept(): %v\n", err)
//			continue
//		}
//		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
//	}
//}

func startServer() {
	b := new(B)

	server := rpc.NewServer()
	server.RegisterName("DDD", b)

	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

//func main() {
//	go startServer()
//
//	conn, err := net.Dial("tcp", "localhost:8222")
//
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//
//	args := &MultiplyArgs{7, 8, 1}
//	var reply MultiplyArgs
//
//	c := jsonrpc.NewClient(conn)
//
//	for i := 0; i < 1; i++ {
//
//		err = c.Call("DDD.Multiply", args, &reply)
//		if err != nil {
//			log.Fatal("arith error:", err)
//		}
//		fmt.Printf("Arith: =%d" , reply)
//	}
//}

func channelTest() chan string {
	return make(chan string)
}

func channelLis1(c chan string) {
	select {
	case v := <-c:
		println("channelLis1:" + v)
	}
}

func channelLis2(c chan string) {
	select {
	case v := <-c:
		println("channelLis2:" + v)
	}
}
//
//func main() {
//	c := channelTest()
//
//	//for {
//		go channelLis1(c)
//
//		go channelLis2(c)
//
//		c <- "aaabbb"
//
//
//	//}
//
//}

func main() {
	version := &orderbook.StoreVersion{1, 12}
	var levedbStore orderbook.Store
	levedbStore = &store.LevelDbStore{Version:*version}
	
	levedbStore.Store()
}