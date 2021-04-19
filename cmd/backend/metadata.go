package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// PodMetadata represents info about an InstanceMetadata in GCE
type PodMetadata struct {
	Name       string
	Namespace  string
	HostIP     string
	PodIP      string
	StartTime  string
	RawRequest string
}

// Populate creates a new instance with info filled out
func (p *PodMetadata) Populate(version string) error {
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
