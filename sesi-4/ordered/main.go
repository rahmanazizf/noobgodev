package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

// BERHASIL
func main() {
	coba := []string{"coba1", "coba2", "coba3"}
	bisa := []string{"bisa1", "bisa2", "bisa3"}
	for i := 1; i < 5; i++ {
		wg.Add(2) // Menambahkan 2 goroutine ke WaitGroup
		go printer(coba, i)
		go printer(bisa, i)

		wg.Wait() // Menunggu kedua goroutine selesai

	}
}

func printer(data interface{}, id int) {
	for i := 0; i < 1; i++ {
		mutex.Lock()
		fmt.Println(data, id)
		mutex.Unlock()
	}

	wg.Done()
}
