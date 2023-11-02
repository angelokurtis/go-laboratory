package pod

import (
	"time"
)

type Object struct {
	Metadata *Metadata `json:"metadata,omitempty"`
	Status   *Status   `json:"status,omitempty"`
}

type Metadata struct {
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
	Labels            Labels    `json:"labels,omitempty"`
	Name              string    `json:"name,omitempty"`
	Namespace         string    `json:"namespace,omitempty"`
}

type Labels map[string]string

type Status struct {
	Phase string `json:"phase,omitempty"`
	PodIP string `json:"podIP,omitempty"`
}
