package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const NBWork = 10

var (
	wID = 0
)

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
	time.Sleep(1 * time.Second) // simulation du temps de travail
	wID++                       // on incrémente l'ID du travail
	newID := wID
	fmt.Printf("work id: %v is finished.\n", newID)
	_, err := subWebserviceCalled(newID)
	if err != nil {
		print(err)
	}
	wg.Done()
}

func subWebserviceCalled(workID int) (data []byte, err error) {
	fmt.Printf("SubWebserviceCalled called id: %d\n", workID)
	defer func() {
		fmt.Println("SubWebserviceCalled finished err:", err, "data:", string(data))
	}()
	// ce weekend, je pars en vacances
	// je n'ai pas le temps de faire ça proprement
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		return nil, err
	}

	// bim bam boum, je fais ma requête
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// on fait comme à la maison
	return data, err
}
