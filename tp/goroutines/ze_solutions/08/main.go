package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
	time.Sleep(1 * time.Second) // simulation du temps de travail
	newID := <-workID
	fmt.Printf("work id: %v is finished.\n", newID)
	_, err := subWebserviceCalled(newID)
	if err != nil {
		print(err)
	}
	wg.Done()
}

// httpClient is a global HTTP client with a timeout.
// à ne pas faire en prod, préférer un client par service.
var httpClient = http.Client{
	Timeout: time.Second * 3,
}

func subWebserviceCalled(workID int) (body []byte, err error) {
	fmt.Printf("SubWebserviceCalled called id: %d\n", workID)
	// on termine en affichant le résultat.
	defer func() {
		fmt.Println("SubWebserviceCalled finished err:", err, "resp:", string(body))
	}()

	// on crée un contexte avec timeout.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// on crée la requête HTTP avec le contexte.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://rickandmortyapi.com/api", nil)
	if err != nil {
		return nil, err
	}

	// on execute la requête HTTP avec le client.
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// on ferme le body à la sortie de la fonction
	defer resp.Body.Close()

	// on vérifie que le status code est OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	// on lit le body de la réponse HTTP.
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
