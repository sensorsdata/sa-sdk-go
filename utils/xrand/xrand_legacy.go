//go:build !go1.22
// +build !go1.22

package xrand

import (
	"math/rand"
	"time"
)

// init 在包初始化时设置随机数种子。
// 注意：Go 1.20 之前，math/rand 全局 Source 默认 Seed 为 1（确定性序列），
// 必须手动 Seed 才能获得真正的随机数。Go 1.20+ 已自动 Seed，此处调用冗余但无害。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Int32 返回一个非负的伪随机 int32 值
func Int32() int32 {
	return rand.Int31()
}

// Int64 返回一个非负的伪随机 int64 值
func Int64() int64 {
	return rand.Int63()
}

// IntN 返回 [0, n) 范围内的伪随机 int 值
func IntN(n int) int {
	return rand.Intn(n)
}