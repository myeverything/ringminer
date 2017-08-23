package matchengine

import "github.com/Loopring/ringminer/models"

//按照首字母，对未成环的进行存储
//逻辑为：订单会发送给每一个bucket，每个bucket，根据结尾的coin，进行链接，
//订单开始的coin为对应的bucket的标号时，查询订单结尾的coin的bucket，并进行对应的链接

//同时会监听proxy发送过来的订单环，及时进行订单的删除与修改

type Bucket struct {
	//OrderChan chan string
	OrderRingChan chan string
	//Filters []string //过滤
	MatchReg string //是否需要保存该环路
}

func (b *Bucket) NewOrder(order models.Order) {


}

func (b *Bucket) UpdateOrder() {

}

func (b *Bucket) NewOrderRing(ring Ring) {

}

func (b *Bucket) Start() {

}

func (b *Bucket) SetPrefix() {

}

func (b *Bucket) AddFilter() {

}

