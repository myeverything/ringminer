package matchengine

import "math/big"



//代理，控制整个match流程，其中会提供几种实现，如bucket、realtime，etc。

var orderMinAmount big.Int	//todo：订单的最小金额，可能需要用map，记录每种货币的最小金额，应该定义filter，过滤环的验证规则

type Proxy interface {
	Start()  //启动
	Stop() //停止
	NewOrder()
	UpdateOrder()  //订单的更新
	//NewOrderRing() //todo:并不需要处理新环事件，只是处理新订单、订单修改的事件
	AddFilter()	//增加ring的
}
