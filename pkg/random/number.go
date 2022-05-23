package random

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Int() int {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt)))
	return int(result.Int64())
}

func Intn(n int) int {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	return int(result.Int64())
}

func Int64() int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	return result.Int64()
}

func Int64n(n int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(n))
	return result.Int64()
}

func Uint64() uint64 {
	n := new(big.Int).SetUint64(math.MaxUint64)
	result, _ := rand.Int(rand.Reader, n)
	return result.Uint64()
}

func Uint64n(n uint64) uint64 {
	result, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(n))
	return result.Uint64()
}
