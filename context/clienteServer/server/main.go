package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":9000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Print("Request started")
	defer log.Print("Completed request")
	select {
	case <-time.After(5 * time.Second):
		log.Print("Request successfully processed")
		w.Write([]byte("Request successfully processed"))
	case <-ctx.Done():
		log.Print("Request canceled")
	}
}
