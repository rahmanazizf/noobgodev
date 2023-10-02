package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	coba := []string{"coba1", "coba2", "coba3"}
	bisa := []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i < 5; i++ {
		wg.Add(2)
		go printer(bisa, i)
		go printer(coba, i)
	}
	wg.Wait()
}

func printer(data interface{}, idx int) {
	defer wg.Done()
	for i := 0; i < 1; i++ {
		fmt.Println(data, idx)
	}
}
