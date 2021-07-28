package controller

import (
	"clientset-demo/constant"
	"clientset-demo/util"
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"log"
)

type DeploymentController struct {
}

// todo 优化
func NewDeployment(namespace, name string) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: util.Int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx-demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
					Labels: map[string]string{
						"app": "nginx-demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "harbor.dev.com/test-demo/nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							ImagePullPolicy: apiv1.PullIfNotPresent,
						},
					},
					RestartPolicy: apiv1.RestartPolicyAlways,
				},
			},
		},
	}

	return deployment
}

func (receiver *DeploymentController) CreateDeployment(clientset *kubernetes.Clientset, namespace, name string) (*appsv1.Deployment, error) {
	// Create Deployment
	log.Println("creating deployment...")
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	deployment := NewDeployment(constant.NginxNamespace, name)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

	return result, err
}

func (receiver *DeploymentController) UpdateDeployments(clientset *kubernetes.Clientset, namespace, name string) error {
	log.Println("Updating deployment...")
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	//    You have two options to Update() this Deployment:
	//
	//    1. Modify the "deployment" variable and call: Update(deployment).
	//       This works like the "kubectl replace" command and it overwrites/loses changes
	//       made by other clients between you Create() and Update() the object.
	//    2. Modify the "result" returned by Get() and retry Update(result) until
	//       you no longer get a conflict error. This way, you can preserve changes made
	//       by other clients between Create() and Update(). This is implemented below
	//			 using the retry utility package included with client-go. (RECOMMENDED)
	//
	// More Info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(context.TODO(), name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get the latest version of Deployment(nginx-demo): %v", getErr))
		}
		result.Spec.Replicas = util.Int32Ptr(1)                                               // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = "harbor.dev.com/test-demo/nginx:1.13" // change nginx version
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})

		return updateErr
	})
	if retryErr != nil {
		return retryErr
	}

	log.Println("updated deployment...")
	return nil
}

func (receiver *DeploymentController) ListDeployments(clientset *kubernetes.Clientset, namespace string) (*appsv1.DeploymentList, error) {
	log.Printf("Listing deployments in namespace %q:\n", namespace)

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})

	return list, err
}

func (receiver *DeploymentController) DeleteDeployments(clientset *kubernetes.Clientset, namespace, name string) error {
	log.Println("Deleting deployments...")

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground // 'Foreground' - 删除前台中所有依赖项的级联策略。
	if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{PropagationPolicy: &deletePolicy}); err != nil {
		return err
	}

	log.Println("Deleted deployment.")
	return nil
}
