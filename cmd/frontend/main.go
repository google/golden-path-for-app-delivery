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
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	var i InstanceMetadata
	err = json.Unmarshal([]byte(body), &i)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", i)
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, "%s", c.Value("version"))
}

func handleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "", "")
}

// InstanceMetadata stores info about the instance this code is running on.
type InstanceMetadata struct {
	ID         string
	Name       string
	Version    string
	Hostname   string
	Zone       string
	Project    string
	InternalIP string
	ExternalIP string
	LBRequest  string
	ClientIP   string
	Error      string
}
