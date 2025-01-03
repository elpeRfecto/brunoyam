package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
	//pipe := make(chan int)
    var tickets uint64
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            for c := 0; c < 10; c++ {
                atomic.AddUint64(&tickets, 1)
				fmt.Printf("Билет %d продан\n", atomic.LoadUint64(&tickets))
            }
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println("Продано билетов:", tickets)
}