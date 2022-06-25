package kube

import (
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func LoadSecretData(apiReader client.Reader, namespace, secretName, dataKey string) (string, error) {
	secret := &corev1.Secret{}
	err := apiReader.Get(context.TODO(), GenerateNamespacedName(namespace, secretName), secret)
	if err != nil {
		return "", err
	}
	retStr, ok := secret.Data[dataKey]
	if !ok {
		return "", errors.Errorf("secret %s did not contain key %s", secretName, dataKey)
	}
	return string(retStr), nil
}

func LoadSecretDataUsingClient(c client.Client, namespace, secretName, dataKey string) (string, error) {
	secret := &corev1.Secret{}
	err := c.Get(context.TODO(), GenerateNamespacedName(namespace, secretName), secret)
	if err != nil {
		return "", err
	}
	retStr, ok := secret.Data[dataKey]
	if !ok {
		return "", errors.Errorf("secret %s did not contain key %s", secretName, dataKey)
	}
	return string(retStr), nil
}
