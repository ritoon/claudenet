package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const NBWork = 10

func main() {
	var wg sync.WaitGroup
	wg.Add(NBWork)
	var mx sync.Mutex
	var iter atomic.Int32
	for i := 0; i < NBWork; i++ {
		go HeavyWork(&iter, &mx, &wg)
	}
	wg.Wait()
}

func HeavyWork(workID *atomic.Int32, mx *sync.Mutex, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	mx.Lock()
	fmt.Printf("work id: %v is finished.\n", workID.Add(1))
	mx.Unlock()
	wg.Done()
}
