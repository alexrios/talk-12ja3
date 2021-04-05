package main

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	times := 0
	for times < 10_000 {
		times++
		counter := Datarace(sync.Mutex{}, 0)
		if counter != 4 {
			t.Fatalf("it should be 4 but found %d on execution %d", counter, times)
		}
	}
}

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


