// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package converter

import (
	time1 "github.com/angelokurtis/go-laboratory/goverter/converter/time"
	pod "github.com/angelokurtis/go-laboratory/goverter/pod"
	v1 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type Pod struct{}

func (c *Pod) Convert(source v1.Pod) pod.Object {
	var podObject pod.Object
	podObject.Metadata = c.v1ObjectMetaToPPodMetadata(source.ObjectMeta)
	podObject.Status = c.v1PodStatusToPPodStatus(source.Status)
	return podObject
}
func (c *Pod) mapStringStringToPodLabels(source map[string]string) pod.Labels {
	podLabels := make(pod.Labels, len(source))
	for key, value := range source {
		podLabels[key] = value
	}
	return podLabels
}
func (c *Pod) v1ObjectMetaToPPodMetadata(source v11.ObjectMeta) *pod.Metadata {
	podMetadata := c.v1ObjectMetaToPodMetadata(source)
	return &podMetadata
}
func (c *Pod) v1ObjectMetaToPodMetadata(source v11.ObjectMeta) pod.Metadata {
	var podMetadata pod.Metadata
	podMetadata.CreationTimestamp = c.v1TimeToPTimeTime(source.CreationTimestamp)
	podLabels := c.mapStringStringToPodLabels(source.Labels)
	podMetadata.Labels = &podLabels
	podMetadata.Name = source.Name
	podMetadata.Namespace = source.Namespace
	return podMetadata
}
func (c *Pod) v1PodStatusToPPodStatus(source v1.PodStatus) *pod.Status {
	podStatus := c.v1PodStatusToPodStatus(source)
	return &podStatus
}
func (c *Pod) v1PodStatusToPodStatus(source v1.PodStatus) pod.Status {
	var podStatus pod.Status
	podStatus.Phase = string(source.Phase)
	podStatus.PodIP = source.PodIP
	return podStatus
}
func (c *Pod) v1TimeToPTimeTime(source v11.Time) *time.Time {
	timeTime := time1.ToStandard(source)
	return &timeTime
}
