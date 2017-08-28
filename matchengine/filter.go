package matchengine

import "github.com/Loopring/ringminer/types"


type Filter interface {
	filter(ring *types.Ring)
}
