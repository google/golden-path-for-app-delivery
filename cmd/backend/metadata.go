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

// Populate creates a new instance with info filled out
func (i *InstanceMetadata) Populate(version string) {
	var err error
	if !metadata.OnGCE() {
		i.Error = "Not running on GCE"
		return
	}

	i.ID, err = metadata.InstanceID()
	if err != nil {
		i.Error += "Unable to populate InstanceID\n"
		return
	}
	i.Zone, err = metadata.Zone()
	if err != nil {
		i.Error += "Unable to populate Zone\n"
		return
	}
	i.Name, err = metadata.InstanceName()
	if err != nil {
		i.Error += "Unable to populate Instance Name\n"
		return
	}
	i.Hostname, err = metadata.Hostname()
	if err != nil {
		i.Error += "Unable to populate Hostname\n"
		return
	}
	i.Project, err = metadata.ProjectID()
	if err != nil {
		i.Error += "Unable to populate Project\n"
		return
	}
	i.InternalIP, err = metadata.InternalIP()
	if err != nil {
		i.Error += "Unable to populate InternalIP\n"
		return
	}
	i.ExternalIP, err = metadata.ExternalIP()
	if err != nil {
		i.Error += "Unable to populate ExternalIP\n"
		return
	}
	i.Version = version
}
