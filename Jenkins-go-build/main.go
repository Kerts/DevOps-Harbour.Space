package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Simple struct {
	Name string
	Url string
}

func handler(w http.ResponseWriter, r *http.Request) {
	simple := Simple{"Hello Kirill",r.Host}

	jsonOutput, _ := json.Marshal(simple)

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/", handler)
    http.Handle("/metrics", promhttp.Handler())
    fmt.Println("Monitoring is started")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
