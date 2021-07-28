package controller

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type ConfigController struct {
}

func (receiver *ConfigController) GetClientset() (*kubernetes.Clientset, error) {
	// 加载kubeconfig配置, 默认: ~/.kube/config.
	// 通过kubernetes.NewForConfig(config)获取到clientset对象.
	// 通过clientset获取到组信息:
	// 	* AppsV1 <=> apiVersion: v1
	// 	* apiVersion: v1 <=> AppsV1Interface
	// 	* AppsV1Interface <=> [ControllerRevisionsGetter,DaemonSetsGetter,DeploymentsGetter,ReplicaSetsGetter,StatefulSetsGetter]
	// 	* DeploymentsGetter <=> DeploymentInterface
	//  * DeploymentInterface <=> Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (*v1.Deployment, error)
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		// kubeconfig = flag.String("kubeconfig", filepath.Join(constant.KubeConfigHomeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// config -> clientset.
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)

	return clientset, err
}
