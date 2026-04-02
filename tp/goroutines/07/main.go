package main

import (
	"fmt"
	"sync"
	"time"
)

const NBWork = 10

func main() {
	var wg sync.WaitGroup
	wg.Add(NBWork)
	//
	ch := make(chan int)
	for i := 0; i < NBWork; i++ {
		go HeavyWork(ch, &wg)
	}
	for i := 0; i < NBWork; i++ {
		ch <- i
	}
	close(ch) // permet de fermer le channel de communication
	wg.Wait()
}

func HeavyWork(workID chan int, wg *sync.WaitGroup) {
	fmt.Printf("HeavyWork called.\n")
	defer wg.Done()
	time.Sleep(1 * time.Second) // simulation du temps de travail
	for {
		id, ok := <-workID
		if !ok {
			return
		}
		fmt.Printf("work id: %v is finished.\n", id)
	}

}
