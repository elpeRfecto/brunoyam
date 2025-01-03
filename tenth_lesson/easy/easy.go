package main

import (
	"sync"
)

func main() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mux.Lock()
				counter++
				mux.Unlock()
			}
		}()
	}
	wg.Wait()
	println(counter)
}