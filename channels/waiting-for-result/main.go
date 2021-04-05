package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "tarefa"
		fmt.Println("sender : tarefa enviada")
	}()

	p := <-ch
	fmt.Println("receiver : recebido :", p)

	time.Sleep(time.Second)
}