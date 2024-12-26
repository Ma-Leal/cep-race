package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ma-Leal/cep-race/internal/entity"
	"github.com/Ma-Leal/cep-race/internal/infra/webserver/handlers"
)

var port = 8000

type Output struct {
	Message string
}

func main() {
	log.Printf("Server running on port %d\n", port)
	log.Printf("Usage example: http://localhost:%d/?cep=01415003\n", port)

	ch1 := make(chan entity.Address)
	ch2 := make(chan entity.Address)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go handlers.GetViaCep(ch1, w, r)
		go handlers.BrasilApi(ch2, w, r)
		select {
		case msg := <-ch1:
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(msg)

		case msg := <-ch2:
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(msg)

		case <-time.After(time.Duration(time.Second)):
			json.NewEncoder(w).Encode("timeout")

		}
	},
	)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
