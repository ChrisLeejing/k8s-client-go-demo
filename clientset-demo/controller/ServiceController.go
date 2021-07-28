package controller

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"log"
)

type ServiceController struct {
}

func NewService(namespace, name string, nodePort int32) *apiv1.Service {
	svc := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: apiv1.ServiceSpec{
			Type: "NodePort",
			Selector: map[string]string{
				"app": "nginx-demo",
			},
			Ports: []apiv1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{
						IntVal: 80,
					},
					NodePort: nodePort,
				},
			},
		},
	}

	return svc
}

func (receiver *ServiceController) CreateService(clientset *kubernetes.Clientset, namespace, name string, nodePort int32) (*apiv1.Service, error) {
	log.Printf("Creating service: namespace: %s, name: %s\n", namespace, name)
	svc := NewService(namespace, name, nodePort)
	service, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), svc, metav1.CreateOptions{})

	return service, err
}

func (receiver *ServiceController) ListServices(clientset *kubernetes.Clientset, namespace string) ([]string, error) {
	var services []string
	serviceList, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	for _, service := range serviceList.Items {
		log.Printf("Listing service: namespace: %s, name: %s\n", service.Namespace, service.Name)
		services = append(services, service.Name)
	}

	return services, err
}

func (receiver *ServiceController) GetService(clientset *kubernetes.Clientset, namespace, name string) (*apiv1.Service, error) {
	return clientset.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (receiver *ServiceController) UpdateService(clientset *kubernetes.Clientset, namespace, name string, newNodePort int32) (*apiv1.Service, error) {
	log.Printf("Updating service: namespace: %s, name: %s\n", namespace, name)
	service, err := receiver.GetService(clientset, namespace, name)
	if err != nil {
		return nil, fmt.Errorf("GetService err: %v\n", err)
	}

	service.Spec.Ports[0].NodePort = newNodePort
	newService, err := clientset.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})

	return newService, err
}

func (receiver *ServiceController) DeleteService(clientset *kubernetes.Clientset, namespace, name string) (bool, error) {
	log.Printf("Deleting service: namespace: %s, name: %s\n", namespace, name)
	_, err := receiver.GetService(clientset, namespace, name)
	if err != nil {
		return false, fmt.Errorf("GetService err: %v\n", err)
	}
	err = clientset.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return false, err
	}

	return true, nil
}
