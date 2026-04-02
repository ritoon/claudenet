package worker

import (
	"fmt"
	"sync"
	"time"
)

// Worker is a struct that is using a Service to process jobs.
type Worker struct {
	service Service
}

// Service is the interface that wraps the Get method.
type Service interface {
	Get(workID int) ([]byte, error)
}

// New creates a new Worker with the given service.
func New(service Service) *Worker {
	return &Worker{
		service: service,
	}
}

// Run is launching nbWork goroutines to process nbWork jobs.
func (w *Worker) Run(nbWork int) {
	// on utilise un WaitGroup pour attendre la fin de toutes les goroutines
	var wg sync.WaitGroup
	// on indique au WaitGroup le nombre de goroutines à attendre
	wg.Add(nbWork)
	// on crée un channel pour communiquer les jobs aux goroutines
	ch := make(chan int)
	// on lance nbWork goroutines
	for i := 0; i < nbWork; i++ {
		// chaque goroutine exécute la méthode do
		// on passe le channel et le WaitGroup en paramètre
		// le WaitGroup est passé par pointeur pour que toutes les goroutines
		go w.do(ch, &wg)
	}
	// on envoie nbWork jobs dans le channel
	for i := 0; i < nbWork; i++ {
		// chaque job est un entier (workID)
		ch <- i
	}
	close(ch) // permet de fermer le channel de communication
	wg.Wait()
}

// do is executing a request lastly it's name was HeavyWork.
func (w *Worker) do(workID chan int, wg *sync.WaitGroup) {
	fmt.Printf("do called.\n")

	// à la sortir de la fonction, on indique au WaitGroup qu'on a fini
	defer wg.Done()

	// on boucle pour traiter les jobs
	for {
		select {
		// 1) Lecture du job dans le channel workID,
		// on vérifie que le channel n'est pas fermé
		case id, ok := <-workID:
			if !ok { // chan fermé => plus de jobs
				fmt.Println("worker: channel closed, exit")
				return
			}
			w.exec(id) // exécution du job
		// 3) Timeout global d'inactivité du worker
		case <-time.After(5 * time.Second):
			fmt.Println("worker: idle timeout, exit")
			return
		}
	}
}

func (w *Worker) exec(id int) {
	// 2) Timeout autour du service.Get
	done := make(chan struct{})
	go func() {
		// ici on fait le travail de récupération des données
		_, _ = w.service.Get(id)
		// on indique que le travail est fait
		close(done)
	}()
	// on attend soit la fin du travail, soit le timeout
	select {
	// le travail est fini
	case <-done:
		fmt.Printf("job %d: OK\n", id)
	// timeout du service.Get
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("job %d: TIMEOUT\n", id)
	}
}
