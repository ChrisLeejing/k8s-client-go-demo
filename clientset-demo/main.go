package main

import (
	"clientset-demo/constant"
	"clientset-demo/controller"
	"clientset-demo/util"
	"fmt"
	"log"
)

var (
	deploymentController = &controller.DeploymentController{} // deployment controller.
	serviceController    = &controller.ServiceController{}    // service controller.
)

// 参考: https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
func main() {
	// Get clientset.
	configController := controller.ConfigController{}
	clientset, err := configController.GetClientset()
	if err != nil {
		panic(err)
	}

	// namespace controller.
	namespaceController := &controller.NamespaceController{}
	namespaces := namespaceController.ListNamespaces(clientset)
	for i, namespace := range namespaces {
		log.Printf("Cluster namespace %d : %s\n", i, namespace)
	}

	// Get/List deployment by namespace.
	kubeSystemDeployments, err := deploymentController.ListDeployments(clientset, constant.KubeSystemNamespace)
	if err != nil {
		panic(err)
	}
	for _, deployment := range kubeSystemDeployments.Items {
		log.Printf("namespace = %s, name = %s(%d replicas)\n", deployment.Namespace, deployment.Name, *deployment.Spec.Replicas)
	}

	// Create deployment
	util.Prompt()
	deployment, err := deploymentController.CreateDeployment(clientset, constant.NginxNamespace, "nginx-demo")
	if err != nil {
		panic(err)
	}
	log.Printf("created deployment %q.\n", deployment.GetObjectMeta().GetName())

	// Create service
	util.Prompt()
	service, err := serviceController.CreateService(clientset, constant.NginxNamespace, "nginx", 30007)
	if err != nil {
		panic(err)
	}
	log.Printf("created service namespace: %s, name: %s\n", service.GetObjectMeta().GetNamespace(), service.GetObjectMeta().GetName())

	// Update service
	util.Prompt()
	updateService, err := serviceController.UpdateService(clientset, constant.NginxNamespace, "nginx", 30008)
	if err != nil {
		panic(err)
	}
	log.Printf("updated service namespace: %s, name: %s, nodePort: %d\n", updateService.Namespace, updateService.Name, updateService.Spec.Ports[0].NodePort)

	// Delete service
	util.Prompt()
	ok, err := serviceController.DeleteService(clientset, constant.NginxNamespace, "nginx")
	if !ok || err != nil {
		panic(err)
	}
	log.Printf("deleted service namespace: %s, name: %s\n", constant.NginxNamespace, "nginx")

	// Update Deployment
	util.Prompt()
	err = deploymentController.UpdateDeployments(clientset, constant.NginxNamespace, "nginx-demo")
	if err != nil {
		panic(fmt.Errorf("update failed: %v", err))
	}

	// List deployments
	util.Prompt()
	deployments, err := deploymentController.ListDeployments(clientset, constant.NginxNamespace)
	if err != nil {
		panic(err)
	}
	for _, item := range deployments.Items {
		log.Printf(" * namespace = %s, name = %s, replicas = %d\n", item.Namespace, item.Name, *item.Spec.Replicas)
	}

	// Delete deployments
	util.Prompt()
	err = deploymentController.DeleteDeployments(clientset, constant.NginxNamespace, "nginx-demo")
	if err != nil {
		panic(err)
	}
}

// Cluster namespace 0 : default
// Cluster namespace 1 : kube-node-lease
// Cluster namespace 2 : kube-public
// Cluster namespace 3 : kube-system
// Cluster namespace 4 : kubernetes-dashboard
// Cluster namespace 5 : nginx
// Cluster namespace 6 : springboot
// Listing deployments in namespace "kube-system":
// namespace = kube-system, name = calico-kube-controllers(1 replicas)
// namespace = kube-system, name = coredns(2 replicas)
// namespace = kube-system, name = metrics-server(1 replicas)
// -> Press Enter to continue.
//
//
//
// creating deployment...
// created deployment "nginx-demo".
// -> Press Enter to continue.
//
//
// Creating service: namespace: nginx, name: nginx
// created service namespace: nginx, name: nginx
// -> Press Enter to continue.
//
//
// Updating service: namespace: nginx, name: nginx
// updated service namespace: nginx, name: nginx, nodePort: 30008
// -> Press Enter to continue.
//
//
// Deleting service: namespace: nginx, name: nginx
// deleted service namespace: nginx, name: nginx
// -> Press Enter to continue.
//
//
// Updating deployment...
// updated deployment...
// -> Press Enter to continue.
//
//
// Listing deployments in namespace "nginx":
//  * namespace = nginx, name = nginx-demo, replicas = 1
// -> Press Enter to continue.
//
//
// Deleting deployments...
// Deleted deployment.