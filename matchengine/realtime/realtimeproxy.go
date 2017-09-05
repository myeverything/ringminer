package realtime

import (
	"sync"
	"github.com/Loopring/ringminer/types"
)

/**
todo：功能完整性上，必须要实现的部分
实时计算最小环
有向图的极小强连通分支
 */

type RealtimeProxy struct {
	mtx sync.RWMutex

	OrderChangeChan chan *types.Order   //订单改变的channel，在匹配过程中，订单改变可以及时终止或更改当前匹配




}