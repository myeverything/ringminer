package matchengine_test

import (
	"testing"
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/matchengine"
)

func TestRingClient(t *testing.T) {

	ring := &types.RingState{}
	hash := &types.Hash{}
	hash.SetBytes([]byte("testtesthash"))
	ring.Hash = *hash
	ring.FeeMode = 1

	matchengine.NewRing(ring)
}

