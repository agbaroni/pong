package main

import (
	"io"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Header().Set("Content-Type", "text/plain")

		_, err := io.WriteString(w, "PING\n")

		if err != nil {
			log.Print(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
