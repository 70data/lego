package kube

import (
	"context"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func GetObjectList(ctx context.Context, c client.Client, obj client.ObjectList, listOption ...client.ListOption) (client.ObjectList, error) {
	err := c.List(ctx, obj, listOption...)
	if err != nil {
		klog.Errorln(err)
	}
	return obj, err
}

func CreateOrUpdate(ctx context.Context, c client.Client, obj client.Object, f controllerutil.MutateFn) (controllerutil.OperationResult, error) {
	return controllerutil.CreateOrUpdate(ctx, c, obj, func() error {
		original := obj.DeepCopyObject()
		if err := f(); err != nil {
			return err
		}
		generateObjectDiff(original, obj)
		return nil
	})
}

func generateObjectDiff(original runtime.Object, modified runtime.Object) {
	diff := cmp.Diff(original, modified)
	if len(diff) != 0 {
		klog.Infoln(diff)
	}
}
