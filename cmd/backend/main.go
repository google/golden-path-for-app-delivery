package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gin-gonic/gin"
)

var version = os.Getenv("VERSION")

func main() {
	port := ":8080"

	flag.Parse()

	r := gin.Default()
	log.Printf("Backend version: %s\n", version)

	r.GET("/", handleIndex)
	r.GET("/version", handleVersion)
	r.GET("/healthz", handleHealthz)
	r.Run(port)
}

func handleIndex(c *gin.Context) {
	log.Printf("Received request from %s at %s", c.Request.RemoteAddr, c.Request.URL.EscapedPath())
	p := PodMetadata{}
	err := p.Populate(version)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	raw, _ := httputil.DumpRequest(c.Request, true)
	p.RawRequest = string(raw)
	c.JSON(http.StatusOK, p)
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, "%s", c.Value("version"))
}

func handleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "", "")
}
