package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	n *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		n: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "my_counter",
			Help: "An infinite progressing number",
		}, []string{"n"}),
	}

	reg.MustRegister(m.n)

	return m
}

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)

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

			m.n.With(prometheus.Labels{"n":"PING"}).Inc()
		}
	})

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
