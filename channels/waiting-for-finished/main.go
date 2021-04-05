package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForFinished()
}

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("fechei o canal")
	}()

	_, withData := <-ch
	fmt.Println("recebi o sinal!", withData)

	time.Sleep(time.Second)
}
