package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Final Counter:", Datarace(0))
}

func Datarace(counter int64) int64{
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				atomic.AddInt64(&counter, 1) // <- Ninguém mais atualiza counter até essa linha terminar // HL
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}