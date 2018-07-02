package main

import (
	"cloud.google.com/go/compute/metadata"
)

// InstanceMetadata represents info about an InstanceMetadata in GCE
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

type assigner struct {
	err error
}

func (a *assigner) assign(getVal func() (string, error)) string {
	if a.err != nil {
		return ""
	}
	s, err := getVal()
	if err != nil {
		a.err = err
	}
	return s
}

// New creates a new instance with info filled out
func (i InstanceMetadata) Populate(version string) InstanceMetadata {
	if !metadata.OnGCE() {
		i.Error = "Not running on GCE"
		return i
	}

	a := &assigner{}
	i.ID = a.assign(metadata.InstanceID)
	i.Zone = a.assign(metadata.Zone)
	i.Name = a.assign(metadata.InstanceName)
	i.Hostname = a.assign(metadata.Hostname)
	i.Project = a.assign(metadata.ProjectID)
	i.InternalIP = a.assign(metadata.InternalIP)
	i.ExternalIP = a.assign(metadata.ExternalIP)
	i.Version = version

	if a.err != nil {
		i.Error = a.err.Error()
	}
	return i
}
