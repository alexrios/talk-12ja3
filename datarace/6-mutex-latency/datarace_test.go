package main

import (
	"sync"
	"testing"
)

func Benchmark(b *testing.B) {
	mutex := sync.Mutex{}
	for i := 0; i < b.N; i++ {
		Datarace(mutex, 0)
	}
}

func BenchmarkMultipleLocks(b *testing.B) {
	mutex := sync.Mutex{}
	for i := 0; i < b.N; i++ {
		DataraceMultiple(mutex, 0)
	}
}


