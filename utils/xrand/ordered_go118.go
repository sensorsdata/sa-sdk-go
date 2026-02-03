//go:build !go1.22
// +build !go1.22

package xrand

import "math/rand"

// IntN returns, as an int, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func IntN(n int) int {
	// bearer:disable go_gosec_crypto_weak_random
	return rand.Intn(n)
}

func Int32() int32 {
	// bearer:disable go_gosec_crypto_weak_random
	return rand.Int31()
}

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int64() int64 {
	// bearer:disable go_gosec_crypto_weak_random
	return rand.Int63()
}
