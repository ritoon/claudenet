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
	var mx sync.Mutex
	var myUnSafeVariable int
	for i := 0; i < NBWork; i++ {
		go Write(&myUnSafeVariable, &mx, &wg)
		go Read(&myUnSafeVariable, &mx, &wg)
	}
	wg.Wait()
}

func Write(myUnSafeVariable *int, mx *sync.Mutex, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	mx.Lock()
	*myUnSafeVariable++
	mx.Unlock()
	wg.Done()
}

func Read(myUnSafeVariable *int, mx *sync.Mutex, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	mx.Lock()
	fmt.Printf("myUnSafeVariable: %v.\n", *myUnSafeVariable)
	mx.Unlock()
	wg.Done()
}
