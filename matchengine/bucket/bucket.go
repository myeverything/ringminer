package matchengine

import (
	"github.com/Loopring/ringminer/types"
	"sync"
	"github.com/Loopring/ringminer/matchengine"
)

//负责生成ring，并计算ring相关的所有参数

//按照首字母，对未成环的进行存储
//逻辑为：订单会发送给每一个bucket，每个bucket，根据结尾的coin，进行链接，
//订单开始的coin为对应的bucket的标号时，查询订单结尾的coin的bucket，并进行对应的链接

//同时会监听proxy发送过来的订单环，及时进行订单的删除与修改

//应当尝试更改为node，提高内存的利用率
//type node struct {
//	Buy TokenAddress
//	Sell TokenAddress
//	Orders map[string]types.Order //hash->order
//	ReachPath [][]*node	//可以到达该节点的途径
//}

type order struct {
	types.Order
	postions []*semiRingPos
}


type semiRingPos struct {
	semiRingKey string	//可以到达的途径
	index int	//所在的数组索引
}

type semiRing struct {
	orders []*order	//组成该半环的node
	hash string
	//reduction reductionOrder 	//半环组成的规约后的新的order
}

type Bucket struct {
	ringChan *chan *types.Ring
	token         TokenAddress      //开始的地址
	semiRings     map[string]*semiRing	//每个semiRing都给定一个key
	orders        map[string]*order //order hash -> order
	mtx           sync.RWMutex
}

//新bucket
func NewBucket(token TokenAddress, ringChan *chan *types.Ring) *Bucket {
	bucket := &Bucket{}
	bucket.token = token
	bucket.ringChan = ringChan
	return bucket
}

func (b *Bucket) NewOrder(ord types.Order) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	//最后一个token为当前token，则可以组成环，匹配出最大环，并发送到proxy
	if (ord.InToken == b.token) {
		var ring *types.Ring
		var idx int
		var semiRing *semiRing
		for idx, semiRing = range b.semiRings {
			if (semiRing.orders[-1].OutToken == ord.InToken) {
				//todo：兑换率判断
				//生成环
				ring1 := &types.Ring{}
				ring1.Orders = append(semiRing, ord)
				matchengine.ComputeRing() //todo:计算换的费用、折扣率等
				//进行判断
				if (ring == nil || (ring.FeeRecipient <= ring1.FeeRecipient)) {
					ring = ring1
				}
			}
		}
		//for循环过后，需要提取出收益最高的环，发送给proxy
		//todo：生成新环后，需要proxy将新环对应的各个订单的状态发送给每个bucket，便于修改，, 还有一些过滤条件
		//删除对应的semiRing
		b.semiRings = append(b.semiRings[:idx], b.semiRings[idx + 1:])
		b.ringChan <- ring

	} else if (ord.InToken == b.token) {
		orderWithSemiRing := &order{}
		orderWithSemiRing.Order = ord
		//首先生成包含自己的semiRing
		selfSemiRing := &semiRing{}
		selfSemiRing.orders = []*order{ord}
		b.semiRings = append(b.semiRings, selfSemiRing)
		//相等的话，则为第一层，下面每一层都加过来
		for _,semiRing1 := range b.semiRings {
			//todo：兑换率判断
			if true {
				semiRing2 := &semiRing{}
				semiRing2.orders = append(selfSemiRing.orders, semiRing1.orders[1:]...)
				//todo:semiRing的key
				semiRing2.hash = "dfddddtestetsete"
				b.semiRings = append(b.semiRings, semiRing2)
				//同时需要修改每个订单中保存的semiRing的信息
				for idx,order1 := range semiRing2.orders {
					//生成新的semiring,
					order1.postions = append(order1.postions, &semiRingPos{semiRingKey:semiRing2.hash, index:idx})
				}
			}
		}
	} else {
		//第二层以下，只检测最后的token 即可
		for _, semiRing1 := range b.semiRings {
			if(semiRing1.orders[-1].OutToken == ord.InToken) {
				//todo:生成新的semiRing
				semiRing2 := &semiRing{}
				semiRing2.orders = append(semiRing1.orders, &order{ord})
				semiRing2.hash = "dfddddtestetsete"
				b.semiRings = append(b.semiRings, semiRing2)
				for idx,order1 := range semiRing2.orders {
					order1.postions = append(order1.postions, &semiRingPos{semiRingKey:semiRing2.hash, index:idx})
				}
			}
		}
	}

}

func (b *Bucket) UpdateOrder(order types.Order) {
	//order的更改，除了订单容量和是否取消等，其他的并不能修改
	//修改订单的容量，主要涉及收益，其他的并不需修改
	//订单的新状态
	//todo：修改时，如果已经提交了ring，如何处理，
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	b.orders[order.Id] = order
	//todo：如果环路已经计算了交易量等信息，需要修改对应的环路
	//for _,ring := range b.orders[order.Id].ReachPath {
	//	for
	//}

}

func (b *Bucket) Start() {

}
func (b *Bucket) Stop() {

}

