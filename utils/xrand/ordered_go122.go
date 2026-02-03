//go:build go1.22
// +build go1.22

package xrand

import "math/rand/v2"

// IntN returns, as an int, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func IntN(n int) int {
	return rand.IntN(n)
}

func Int32() int32 {
	return rand.Int32()
}

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int64() int64 {
	return rand.Int64()
}
