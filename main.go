package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var color = "red"
var version = os.Getenv("VERSION")
var rdb = &redis.Client{}
var rdbCtx = context.Background()

func main() {
	port := ":8080"
	rdb = redis.NewClient(&redis.Options{
		Addr:     getRedisURL(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	flag.Parse()

	r := gin.Default()
	log.Printf("Backend version: %s\n", version)

	r.LoadHTMLGlob("index.html")
	r.GET("/", handleIndex)
	r.GET("/version", handleVersion)
	r.GET("/healthz", handleHealthz)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func handleIndex(c *gin.Context) {
	log.Printf("Received request from %s at %s", c.Request.RemoteAddr, c.Request.URL.EscapedPath())
	p := PodMetadata{}
	counter, err := incrCounter(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	err = p.Populate(version, counter, color)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", p)
}

func incrCounter(c *gin.Context) (string, error) {
	err := rdb.Incr(rdbCtx, "counter").Err()
	if err != nil {
		return "", err
	}
	val, err := rdb.Get(rdbCtx, "counter").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, "%s", c.Value("version"))
}

func handleHealthz(c *gin.Context) {
	status := rdb.Ping(rdbCtx)

	if status.Err() == nil {
		c.String(http.StatusOK, "", "")
	} else {
		errorString := fmt.Sprintf("Unable to reach redis at %v", rdb.Options().Addr)
		c.String(http.StatusInternalServerError, errorString, "")
	}
}

// PodMetadata represents info about an InstanceMetadata in GCE
type PodMetadata struct {
	Name       string
	Namespace  string
	HostIP     string
	PodIP      string
	StartTime  string
	RawRequest string
	Counter    string
	Version    string
	Color      string
	RedisURL   string
}

// Populate creates a new instance with info filled out
func (p *PodMetadata) Populate(version string, counter string, color string) error {
	hostname := os.Getenv("HOSTNAME")
	config, err := rest.InClusterConfig()
	if err != nil {
		return fmt.Errorf("unable to create InClusterConfig client: %v", err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("unable to create kubernetes client: %v", err)
	}

	pod, err := clientset.CoreV1().Pods(getNamespace()).Get(context.TODO(), hostname, v1.GetOptions{})
	if err != nil {
		return fmt.Errorf("unable to find pod %s: %v", hostname, err)
	}
	p.Name = pod.Name
	p.HostIP = pod.Status.HostIP
	p.Namespace = pod.Namespace
	p.PodIP = pod.Status.PodIP
	p.StartTime = pod.Status.StartTime.String()
	p.Counter = counter
	p.Version = version
	p.Color = color
	p.RedisURL = rdb.Options().Addr
	return nil
}

func getNamespace() string {
	// Fall back to the namespace associated with the service account token, if available
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}
	return "default"
}
