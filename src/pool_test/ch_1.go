package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	var tickets chan byte
	var tokenWg sync.WaitGroup
	max_tickets := 10
	tx_size := 40 * 10000

	tickets = make(chan byte, max_tickets)
	for i := 0; i < max_tickets; i++ {
		tickets <- 1
	}

	tokenWg.Add(tx_size)
	i := 0
	for i < tx_size {
		<-tickets
		go func() {
			defer func() {
				tokenWg.Done()
				tickets <- 1
			}()
			i = i + 1
			fmt.Println(i)
		}()
	}

	tokenWg.Wait()
}
