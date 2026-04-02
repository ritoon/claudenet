package main

import (
	"fmt"
	"time"
)

const NBWork = 10

func main() {
	for i := 0; i < NBWork; i++ {
		go HeavyWork(i)
	}
}

func HeavyWork(workID int) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	fmt.Printf("work id: %v is finished.\n", workID)
}
