package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[int]int
}

func (c *Cache) Get(key int) int {
	c.RLock()
	defer c.RUnlock()
	return c.items[key]
}

func (c *Cache) Set(key, value int) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = value
}

func main() {
	r := rand.IntN
	cache := Cache{
		items: make(map[int]int),
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(i, r(100))
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Get(i)
		}()
	}
	wg.Wait()
	for k, v := range cache.items {
		fmt.Println(k, v)
	}
}
