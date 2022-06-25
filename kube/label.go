package kube

import (
	"k8s.io/apimachinery/pkg/labels"
)

func GetLabelSelector(key string, value string) labels.Selector {
	// Key is empty
	if len(key) == 0 {
		return labels.SelectorFromSet(map[string]string{})
	}
	return labels.SelectorFromSet(labels.Set{key: value})
}

func AddLabel(labels map[string]string, key string, value string) map[string]string {
	if key == "" {
		return labels
	}
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[key] = value
	return labels
}
