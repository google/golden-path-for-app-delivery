package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("VERSION")
	port := 80
	backend := flag.String("backend-service", "http://127.0.0.1:8080", "hostname of backend server")
	flag.Parse()
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", version)
	})

	fmt.Printf("Frontend version: %s\n", version)

	tpl := template.Must(template.New("out").Parse(html))

	transport := http.Transport{DisableKeepAlives: false}
	client := &http.Client{Transport: &transport}
	req, _ := http.NewRequest(
		"GET",
		*backend,
		nil,
	)
	req.Close = false

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s", r.RemoteAddr)
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Error: %s\n", err.Error())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s\n", err.Error())
			return
		}
		var i InstanceMetadata
		err = json.Unmarshal([]byte(body), &i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s\n", err.Error())
			return
		}
		tpl.Execute(w, i)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Backend could not be connected to: %s", err.Error())
			return
		}
		defer resp.Body.Close()
		ioutil.ReadAll(resp.Body)
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
