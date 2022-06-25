package kube

import (
	"k8s.io/apimachinery/pkg/types"
)

func containsStringFromSlice(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeStringFromSlice(slice []string, s string) []string {
	var result []string
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return result
}

func GenerateNamespacedName(namespace, name string) types.NamespacedName {
	var namespacedName types.NamespacedName
	namespacedName.Namespace = namespace
	namespacedName.Name = name
	return namespacedName
}
