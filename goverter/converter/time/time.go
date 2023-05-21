package time

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ToStandard(t metav1.Time) time.Time {
	return t.Time
}

func ToK8s(t time.Time) metav1.Time {
	return metav1.Time{Time: t}
}
