package tests

import (
	"testing"
	"reflect"
	"math"
	"math/big"
)

// go test -v tests/big_int_test.go
func Test_BigInt(t *testing.T) {
	b := big.NewInt(math.MaxInt64)
	t.Log("test\t-", "initial b", b)

	n := big.NewInt(100)
	t.Log("test\t", "b multiple", b.Mul(n,b))

	t.Log(big.NewInt(46877).Uint64())
	t.Log(reflect.ValueOf(nil).IsValid())
}
