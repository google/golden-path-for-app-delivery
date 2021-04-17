package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var version = os.Getenv("VERSION")
var backend string

func main() {

	port := ":80"
	backend = *flag.String("backend-service", "http://gceme-backend:8080", "hostname of backend server")
	flag.Parse()

	log.Printf("Frontend version: %s\n", version)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handleIndex)
	r.GET("/version", handleVersion)
	r.GET("/healthz", handleHealthz)
	r.Run(port)

}

func handleIndex(c *gin.Context) {
	// Call backend for info
	resp, err := http.Get(backend)
	if err != nil {
		c.String(http.StatusInternalServerError, "Request to backend server ()%s failed:\n%v", backend, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to read body from backend request:\n%v", err)
		return
	}
	var p PodMetadata
	err = json.Unmarshal([]byte(body), &p)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to parse JSON from backend request:\n%v", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", p)
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, "%s", c.Value("version"))
}

func handleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "", "")
}

// PodMetadata represents info about an InstanceMetadata in GCE
type PodMetadata struct {
	Name        string
	ClusterName string
	Namespace   string
	HostIP      string
	PodIP       string
	StartTime   string
	RawRequest  string
}
