package main

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/trace"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {

	// Set up the flight recorder.
	fr := trace.NewFlightRecorder()
	fr.Start()

	// Set up and run an HTTP server.
	var once sync.Once
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Do the work...
		fmt.Println("working ...")
		time.Sleep(500 * time.Millisecond)

		// We saw a long request. Take a snapshot!
		if time.Since(start) > 300*time.Millisecond {
			// Do it only once for simplicity, but you can take more than one.
			once.Do(func() {
				// Grab the snapshot.
				var b bytes.Buffer
				_, err := fr.WriteTo(&b)
				if err != nil {
					log.Print(err)
					return
				}
				// Write it to a file.
				if err := os.WriteFile("trace.out", b.Bytes(), 0o755); err != nil {
					log.Print(err)
					return
				}
			})
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
