package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Final Counter:", Datarace(sync.Mutex{}, 0))
}

func Datarace(m sync.Mutex, counter int64) int64 {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				m.Lock()
				{
					value := counter
					value++
					counter = value
				}
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}

func DataraceMultiple(m sync.Mutex, counter int64) int64 {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				m.Lock()
				value := counter
				m.Unlock()

				m.Lock()
				value++
				m.Unlock()

				m.Lock()
				counter = value
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}
