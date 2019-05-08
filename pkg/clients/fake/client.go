package fake

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func GetClient() kubernetes.Interface {
	client := fake.NewSimpleClientset()
	return client

}
