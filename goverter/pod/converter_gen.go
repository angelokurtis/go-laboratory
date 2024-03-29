// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package pod

import (
	time "github.com/angelokurtis/go-laboratory/goverter/time"
	v1 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type converterImpl struct{}

func (c *converterImpl) FromK8s(source v1.Pod) Object {
	var podObject Object
	podObject.Metadata = c.v1ObjectMetaToPPodMetadata(source.ObjectMeta)
	podObject.Status = c.v1PodStatusToPPodStatus(source.Status)

	return podObject
}
func (c *converterImpl) MetadataToObjectMeta(source Metadata) v11.ObjectMeta {
	var v1ObjectMeta v11.ObjectMeta
	v1ObjectMeta.Name = source.Name
	v1ObjectMeta.Namespace = source.Namespace
	v1ObjectMeta.CreationTimestamp = time.ToK8s(source.CreationTimestamp)
	v1ObjectMeta.Labels = c.podLabelsToMapStringString(source.Labels)

	return v1ObjectMeta
}
func (c *converterImpl) StatusToStatus(source Status) v1.PodStatus {
	var v1PodStatus v1.PodStatus
	v1PodStatus.Phase = v1.PodPhase(source.Phase)
	v1PodStatus.PodIP = source.PodIP

	return v1PodStatus
}
func (c *converterImpl) ToK8s(source Object) v1.Pod {
	var v1Pod v1.Pod
	v1Pod.ObjectMeta = c.pPodMetadataToV1ObjectMeta(source.Metadata)
	v1Pod.Status = c.pPodStatusToV1PodStatus(source.Status)

	return v1Pod
}
func (c *converterImpl) mapStringStringToPodLabels(source map[string]string) Labels {
	podLabels := make(Labels, len(source))
	for key, value := range source {
		podLabels[key] = value
	}

	return podLabels
}
func (c *converterImpl) pPodMetadataToV1ObjectMeta(source *Metadata) v11.ObjectMeta {
	var v1ObjectMeta v11.ObjectMeta
	if source != nil {
		v1ObjectMeta = c.MetadataToObjectMeta((*source))
	}

	return v1ObjectMeta
}
func (c *converterImpl) pPodStatusToV1PodStatus(source *Status) v1.PodStatus {
	var v1PodStatus v1.PodStatus
	if source != nil {
		v1PodStatus = c.StatusToStatus((*source))
	}

	return v1PodStatus
}
func (c *converterImpl) podLabelsToMapStringString(source Labels) map[string]string {
	mapStringString := make(map[string]string, len(source))
	for key, value := range source {
		mapStringString[key] = value
	}

	return mapStringString
}
func (c *converterImpl) v1ObjectMetaToPPodMetadata(source v11.ObjectMeta) *Metadata {
	podMetadata := c.v1ObjectMetaToPodMetadata(source)
	return &podMetadata
}
func (c *converterImpl) v1ObjectMetaToPodMetadata(source v11.ObjectMeta) Metadata {
	var podMetadata Metadata
	podMetadata.CreationTimestamp = time.ToStandard(source.CreationTimestamp)
	podMetadata.Labels = c.mapStringStringToPodLabels(source.Labels)
	podMetadata.Name = source.Name
	podMetadata.Namespace = source.Namespace

	return podMetadata
}
func (c *converterImpl) v1PodStatusToPPodStatus(source v1.PodStatus) *Status {
	podStatus := c.v1PodStatusToPodStatus(source)
	return &podStatus
}
func (c *converterImpl) v1PodStatusToPodStatus(source v1.PodStatus) Status {
	var podStatus Status
	podStatus.Phase = string(source.Phase)
	podStatus.PodIP = source.PodIP

	return podStatus
}
