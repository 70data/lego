package kube

import (
	"k8s.io/apimachinery/pkg/types"
)

func GenerateNamespacedName(namespace, name string) types.NamespacedName {
	var namespacedName types.NamespacedName
	namespacedName.Namespace = namespace
	namespacedName.Name = name
	return namespacedName
}
