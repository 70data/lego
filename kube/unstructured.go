package kube

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/klog/v2"
)

func ConvertStructToUnstructured(src interface{}) (*unstructured.Unstructured, error) {
	yamlData, _ := json.Marshal(src)
	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	utd := &unstructured.Unstructured{}
	if _, _, err := decoder.Decode(yamlData, nil, utd); err != nil {
		klog.Errorln(err)
		return nil, err
	}
	return utd, nil
}
