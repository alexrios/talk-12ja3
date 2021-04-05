package main

import (
	"log"
	"sync"
)

func main() {
	Test()
}

func Datarace(counter int) int{
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				value := counter
				value++
				counter = value
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}

func Test() {
	times := 0
	for {
		times++
		counter := Datarace(0)
		if counter != 4 {
			log.Fatalf("it should be 4 but found %d on execution %d", counter, times)
		}
	}
}