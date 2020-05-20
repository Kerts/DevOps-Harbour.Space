package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":9090", nil))
}
