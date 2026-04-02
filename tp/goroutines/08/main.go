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

	ch := make(chan int)
	for i := 0; i < NBWork; i++ {
		go HeavyWork(ch, &wg)
	}
	for i := 0; i < NBWork; i++ {
		ch <- i
	}
	close(ch)
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

var httpClient = http.Client{
	Timeout: time.Second * 3,
}

func subWebserviceCalled(workID int) (body []byte, err error) {
	fmt.Printf("SubWebserviceCalled called id: %d\n", workID)
	defer func() {
		fmt.Println("SubWebserviceCalled finished err:", err, "resp:", string(body))
	}()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://rickandmortyapi.com/api", nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
