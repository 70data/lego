package log

import (
	"k8s.io/klog/v2"
)

func Init() {
	klog.InitFlags(nil)
}
