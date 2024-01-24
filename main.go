package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"log"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(req, true)

		if err != nil {
        		log.Print(err)
    		} else {
			log.Print(string(reqDump))

			w.WriteHeader(http.StatusOK)

			w.Header().Set("Content-Type", "text/plain")

			_, err := io.WriteString(w, "PING\n")

			if err != nil {
				log.Print(err)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
