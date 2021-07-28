package controller

import (
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceController struct {
	namespaces []string
}

func (receiver *NamespaceController) ListNamespaces(clientset *kubernetes.Clientset) []string {
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range namespaceList.Items {
		receiver.namespaces = append(receiver.namespaces, item.ObjectMeta.Name)
	}

	return receiver.namespaces
}
