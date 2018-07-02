package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	version := os.Getenv("VERSION")
	port := 8080
	flag.Parse()
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", version)
	})

	log.Printf("Backend version: %s\n", version)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s", r.RemoteAddr)
		var i InstanceMetadata
		i = i.Populate(version)
		raw, _ := httputil.DumpRequest(r, true)
		i.LBRequest = string(raw)
		resp, _ := json.Marshal(i)
		fmt.Fprintf(w, "%s", resp)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
