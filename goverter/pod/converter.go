//go:generate go run -mod=mod github.com/jmattheis/goverter/cmd/goverter@v0.17.4 -matchFieldsIgnoreCase -packagePath github.com/angelokurtis/go-laboratory/goverter/pod -packageName pod -output ./converter_gen.go ./

package pod

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var singletonConverter = struct {
	value converter
	once  sync.Once
}{}

func getConverter() converter {
	singletonConverter.once.Do(func() {
		singletonConverter.value = new(converterImpl)
	})

	return singletonConverter.value
}

// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:name converterImpl
// goverter:extend github.com/angelokurtis/go-laboratory/goverter/time:ToStandard
// goverter:extend github.com/angelokurtis/go-laboratory/goverter/time:ToK8s
// goverter:output:file ./converter_gen.go
// goverter:output:package github.com/angelokurtis/go-laboratory/goverter/pod
type converter interface {
	// goverter:map ObjectMeta Metadata
	FromK8s(source corev1.Pod) Object
	// goverter:ignoreMissing
	// goverter:map Metadata ObjectMeta
	ToK8s(source Object) corev1.Pod
	// goverter:ignoreMissing
	MetadataToObjectMeta(source Metadata) metav1.ObjectMeta
	// goverter:ignoreMissing
	StatusToStatus(source Status) corev1.PodStatus
}

func FromK8s(source corev1.Pod) Object {
	conv := getConverter()
	return conv.FromK8s(source)
}

func ToK8s(source Object) corev1.Pod {
	conv := getConverter()
	return conv.ToK8s(source)
}
