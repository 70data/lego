package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AddFinalizer(meta *metav1.ObjectMeta, finalizer string) {
	if !HasFinalizer(meta, finalizer) {
		meta.Finalizers = append(meta.Finalizers, finalizer)
	}
}

func HasFinalizer(meta *metav1.ObjectMeta, finalizer string) bool {
	return containsStringFromSlice(meta.Finalizers, finalizer)
}

func RemoveFinalizer(meta *metav1.ObjectMeta, finalizer string) {
	meta.Finalizers = removeStringFromSlice(meta.Finalizers, finalizer)
}
