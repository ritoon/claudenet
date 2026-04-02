package worker

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	service Service
}

type Service interface {
	Get(workID int) ([]byte, error)
}

func New(service Service) *Worker {
	return &Worker{
		service: service,
	}
}

func (w *Worker) Run(nbWork int) {
	var wg sync.WaitGroup
	wg.Add(nbWork)

	ch := make(chan int)
	for i := 0; i < nbWork; i++ {
		go w.do(ch, &wg)
	}
	for i := 0; i < nbWork; i++ {
		ch <- i
	}
	close(ch) // permet de fermer le channel de communication
	wg.Wait()
}

// do is executing a request lastly it's name was HeavyWork.
func (w *Worker) do(workID chan int, wg *sync.WaitGroup) {
	fmt.Printf("do called.\n")
	time.Sleep(1 * time.Second) // simulation du temps de travail
	newID := <-workID
	fmt.Printf("work id: %v is finished.\n", newID)
	_, err := w.service.Get(newID)
	if err != nil {
		print(err)
	}
	wg.Done()
}
