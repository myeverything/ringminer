package types_test

import (
	"testing"
	"github.com/Loopring/ringminer/types"
)

func TestStringToAddress(t *testing.T) {
	str := "0xb794f5ea0ba39494ce839613fffba74279579268"
	add := types.StringToAddress(str)
	t.Log(add.Str())
	t.Log(len("0x08935625ce172eb3c6561404c09f130268808d08ba59dda70aefa0016619acbc"))
}