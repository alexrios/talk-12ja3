package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForTask()
}

func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("receiver : recebi :", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "tarefa"
	fmt.Println("sender : mandei")

	time.Sleep(time.Second)
}
