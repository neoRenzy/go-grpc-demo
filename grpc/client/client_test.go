package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkNewGrpcClient(b *testing.B) {
	tests := []struct {
		name        string
		createTimes int
	}{
		{
			name: "测试创建连接速度",
		},
	}
	for _, tt := range tests {
		//  0.0000632 ns/op
		b.ResetTimer()
		b.Run(tt.name, func(b *testing.B) {
			client := NewGrpcClient()
			client.Close()
		})
	}
}

func TestNewGrpcClient(t *testing.T) {
	cnt := 100000
	now := time.Now()
	for i := 0; i < cnt; i++ {
		client := NewGrpcClient()
		client.Close()
	}
	fmt.Printf("单词耗时%d微秒", time.Now().Sub(now).Microseconds()/int64(cnt))
}
