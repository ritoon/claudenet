package main

import (
	"fmt"
	"os"
	"time"

	"runtime/trace"
)

const NBWork = 10

func main() {
	trace.Start(os.Stderr)
	for i := 0; i < NBWork; i++ {
		go HeavyWork(i)
	}
	trace.Stop()
}

func HeavyWork(workID int) {
	time.Sleep(1 * time.Second) // simulation du temps de travail
	fmt.Printf("work id: %v is finished.\n", workID)
}
