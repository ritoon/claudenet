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
	close(ch)
	wg.Wait()
}

func (w *Worker) do(workID chan int, wg *sync.WaitGroup) {
	fmt.Printf("do called.\n")

	defer wg.Done()

	for {
		select {
		case id, ok := <-workID:
			if !ok {
				fmt.Println("worker: channel closed, exit")
				return
			}
			w.exec(id)
		case <-time.After(5 * time.Second):
			fmt.Println("worker: idle timeout, exit")
			return
		}
	}
}

func (w *Worker) exec(id int) {
	done := make(chan struct{})
	go func() {
		_, _ = w.service.Get(id)
		close(done)
	}()

	select {
	case <-done:
		fmt.Printf("job %d: OK\n", id)
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("job %d: TIMEOUT\n", id)
	}
}
