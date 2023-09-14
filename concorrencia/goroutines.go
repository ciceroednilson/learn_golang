package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//go escrever("Olá mundo!") //goroutine
	//escrever("Pogramando em Go!")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()
	go func() {
		escrever("Pogramando em Go!")
		waitGroup.Done()
	}()
	waitGroup.Wait()
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
