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
	// lecteur
	for i := 0; i < NBWork; i++ {
		go HeavyWork(ch, &wg)
	}
	// rédacteur
	for i := 0; i < NBWork; i++ {
		ch <- i
	}
	wg.Wait()
	close(ch) // permet de fermer le channel de communication
}

func HeavyWork(workID chan int, wg *sync.WaitGroup) {
	fmt.Printf("HeavyWork called.\n")
	time.Sleep(1 * time.Second) // simulation du temps de travail
	fmt.Printf("work id: %v is finished.\n", <-workID)
	wg.Done()
}
