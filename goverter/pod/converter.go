//go:generate goverter -output=../converter/pod_gen.go -packageName=converter .

package pod

import (
	corev1 "k8s.io/api/core/v1"
)

// goverter:converter
// goverter:name Pod
// goverter:extend github.com/angelokurtis/go-laboratory/goverter/converter/time:ToStandard
type Converter interface {
	// goverter:map ObjectMeta Metadata
	Convert(source corev1.Pod) Object
}
