package main

import (
	"time"

	"tp/goroutines/09/service"
	"tp/goroutines/09/worker"
)

const NBWork = 10

func main() {

	// initialise les services
	serviceAPI := service.New(3 * time.Second)
	wkr := worker.New(serviceAPI)

	// lance les tâches
	wkr.Run(NBWork)
}
