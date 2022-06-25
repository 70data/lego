package kube

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Established(crd *v1.CustomResourceDefinition) bool {
	for _, condition := range crd.Status.Conditions {
		if condition.Type == v1.Established && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}
