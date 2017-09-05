package matchengine

import "github.com/Loopring/ringminer/types"


type Filter interface {
	filter(ring *types.RingState)
}

/**
过滤条件
1、inToken != outToken
2、费用
3、
 */