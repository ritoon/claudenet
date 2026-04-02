package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DefaultTimeout is the default timeout for the HTTP client and requests.
const DefaultTimeout = time.Second * 3

// BaseURL is the base URL for the API.
const BaseURL = "https://rickandmortyapi.com/api"

// Client is a simple HTTP client for making requests to an external service.
type Client struct {
	cli *http.Client
}

// New creates a new Client with the given timeout.
// If the timeout is less than 3 seconds, it defaults to DefaultTimeout.
func New(timeout time.Duration) *Client {
	if timeout < time.Second*3 {
		timeout = DefaultTimeout
	}
	return &Client{
		cli: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get makes a GET request to the BaseURL and returns the response body.
// It uses a context with a timeout to avoid hanging requests.
func (cs *Client) Get(workID int) (resp []byte, err error) {
	fmt.Printf("SubWebserviceCalled called id: %d\n", workID)
	defer func() {
		fmt.Println("SubWebserviceCalled finished err:", err, "resp:", string(resp))
	}()

	// on crée un contexte avec timeout.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()

	// on crée la requête HTTP avec le contexte.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseURL, nil)
	if err != nil {
		return nil, err
	}

	// on execute la requête HTTP avec le client.
	res, err := cs.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // on ferme le body à la sortie de la fonction

	// on vérifie que le status code est OK.
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	// on lit le body de la réponse HTTP.
	resp, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resp, err
}
