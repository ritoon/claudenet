package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	// création de la configuration du FlightRecorder
	cfg := trace.FlightRecorderConfig{
		MinAge:   5 * time.Second,
		MaxBytes: 3 << 20, // 3MB
	}
	// Création du FlightRecorder avec la configuration
	fr := trace.NewFlightRecorder(cfg)
	err := fr.Start()
	if err != nil {
		log.Fatalf("unable to trace flight recorder %v", err)
	}
	defer fr.Stop()

	// Création d'un endpoint
	http.HandleFunc("/", handler(fr))

	// Démarrage du serveur sur le port 8080
	log.Println("ouvrez votre navigateur sur l'url http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(fr *trace.FlightRecorder) func(http.ResponseWriter, *http.Request) {
	// permet d'executer une fois une fonction
	var traceWritten sync.Once

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// démarrage de deux goroutines
		var wg sync.WaitGroup
		wg.Add(2)
		go heavyLoad(&wg, 100_000)
		go heavyLoad(&wg, 10_000_000)
		wg.Wait()

		// calcul du temps d'execution
		diff := time.Since(start)

		if diff > 300*time.Millisecond {
			traceWritten.Do(func() {
				err := writeTrace(fr)
				if err != nil {
					log.Println("fail to write to trace : %v", err)
					return
				}
			})
		}
		log.Println("arrêter l'application")
		fmt.Fprintf(w, "work for %f second", diff.Seconds())
	}
}

// heavyLoad est une simulation d'une opération coûteuse.
func heavyLoad(wg *sync.WaitGroup, iterations int) error {
	defer wg.Done()

	for i := 0; i < iterations; i++ {
		_ = fmt.Sprintf("processing %v", i)
	}

	time.Sleep(500 * time.Millisecond)
	return nil
}

// writeTrace permet d'écrir le résultat du trace dans un fichier trace.out.
func writeTrace(fr *trace.FlightRecorder) error {
	if !fr.Enabled() {
		return fmt.Errorf("flight recorder is not open")
	}

	file, err := os.Create("trace.out")
	if err != nil {
		return fmt.Errorf("fail to create trace file: %w", err)
	}
	defer file.Close()

	_, err = fr.WriteTo(file)
	if err != nil {
		return fmt.Errorf("fail to write into trace file: %w", err)
	}
	return nil
}
