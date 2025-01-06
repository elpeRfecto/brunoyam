package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
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


    //ЗДЕСЬ ВТОРОЙ ВАРИАНТ

	// pipe := make(chan int)
	// var tickets int64
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		pipe <- i
	// 	}
	// 	close(pipe)
	// }()

    // for i := 0; i < 100; i++ {
    //     <-pipe
    //     atomic.AddInt64(&tickets, 1)
    //     fmt.Printf("Билет %d продан\n", atomic.LoadInt64(&tickets))

    }
