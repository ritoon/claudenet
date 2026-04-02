package main

import (
	"fmt"
	"sync"
	"time"
)

const NBWork = 100

func main() {
	var wg sync.WaitGroup
	wg.Add(NBWork)
	for i := 0; i < NBWork; i++ {
		go HeavyWork(i, &wg)
	}
	wg.Wait()
}

func HeavyWork(workID int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	fmt.Printf("work id: %v is finished.\n", workID)
	wg.Done()
}
