//go:build go1.22
// +build go1.22

package xrand

import "math/rand/v2"

// Int32 返回一个非负的伪随机 int32 值
func Int32() int32 {
	return rand.Int32()
}

// Int64 返回一个非负的伪随机 int64 值
func Int64() int64 {
	return rand.Int64()
}

// IntN 返回 [0, n) 范围内的伪随机 int 值
func IntN(n int) int {
	return rand.IntN(n)
}