package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Final Counter:", Datarace(0))
}

func Datarace(counter int) int{
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				value := counter
				runtime.Gosched() // <- chamando scheduler // HL
				value++
				counter = value
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}