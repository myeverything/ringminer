package matchengine



//代理，控制整个match流程，其中会提供几种实现，如bucket、realtime，etc。

type Ring struct {

}

type Proxy interface {
	Start()  //启动
	Stop() //停止
	NewOrder()
	NewOrderRing()
	UpdateOrder()  //订单的更新
	AddFilter()	//增加order的filter
}
