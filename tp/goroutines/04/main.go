package main

import (
	"fmt"
	"sync"
	"time"
)

const NBWork = 10

func main() {
	var wg sync.WaitGroup
	wg.Add(NBWork * 2)
	var myUnSafeVariable int
	for i := 0; i < NBWork; i++ {
		go Write(&myUnSafeVariable, &wg)
		go Read(&myUnSafeVariable, &wg)
	}
	wg.Wait()
}

func Write(myUnSafeVariable *int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	*myUnSafeVariable++
	// fmt.Printf("myUnSafeVariable: %v.\n", myUnSafeVariable)
	wg.Done()
}

func Read(myUnSafeVariable *int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	fmt.Printf("myUnSafeVariable: %v.\n", myUnSafeVariable)
	wg.Done()
}
