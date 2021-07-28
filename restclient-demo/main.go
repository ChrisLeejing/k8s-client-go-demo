package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string

	// kubeconfig两种获取方式: 1. home家目录(~/.kube/config); 2. 控制台输入绝对路径获取.
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse() // 解析控制台输入的: -kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#list-pod-v1-core
	// 参考: /api/v1/namespaces/{namespace}/pods
	config.APIPath = "api"
	// pod group: apiVersion: v1
	config.GroupVersion = &corev1.SchemeGroupVersion
	// 指定序列化工具
	config.NegotiatedSerializer = scheme.Codecs

	// 根据config获取RESTClient实例
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// 保存获取pods
	result := &corev1.PodList{}
	// 获取指定namespace的pods
	namespace := "kube-system"
	// GET请求
	err = restClient.Get().
		Namespace(namespace). // 指定namespace: 参考: /api/v1/namespaces/{namespace}/pods
		Resource("pods"). // 查找多个pod: 参考: /api/v1/namespaces/{namespace}/pods
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec). // 指定大小限制和序列化工具
		Do(context.TODO()). // 请求
		Into(result) // 结果存入result
	if err != nil {
		panic(err)
	}

	// 打印表头
	fmt.Printf("namespace\t\tstatus\t\tname\n")
	for _, item := range result.Items {
		fmt.Printf("%v\t\t%v\t\t%v\n", item.Namespace, item.Status.Phase, item.Name)
	}

	// namespace		status		name
	// kube-system		Running		calico-kube-controllers-bb478d64-8vqds
	// kube-system		Running		calico-node-5m9dr
	// kube-system		Running		calico-node-bgdc2
	// kube-system		Running		calico-node-bwpl9
	// kube-system		Running		calico-node-k7k44
	// kube-system		Running		calico-node-l27jg
	// kube-system		Running		calico-node-mt4pc
	// kube-system		Running		calico-node-phskq
	// kube-system		Running		calico-node-vr82f
	// kube-system		Running		coredns-b5c4bcdd4-5nbmd
	// kube-system		Running		coredns-b5c4bcdd4-6kmh5
	// kube-system		Running		etcd-k8s-master-1
	// kube-system		Running		etcd-k8s-master-2
	// kube-system		Running		etcd-k8s-master-3
	// kube-system		Running		kube-apiserver-k8s-master-1
	// kube-system		Running		kube-apiserver-k8s-master-2
	// kube-system		Running		kube-apiserver-k8s-master-3
	// kube-system		Running		kube-controller-manager-k8s-master-1
	// kube-system		Running		kube-controller-manager-k8s-master-2
	// kube-system		Running		kube-controller-manager-k8s-master-3
	// kube-system		Running		kube-proxy-42br5
	// kube-system		Running		kube-proxy-58t6h
	// kube-system		Running		kube-proxy-7vkv5
	// kube-system		Running		kube-proxy-bsnhm
	// kube-system		Running		kube-proxy-fdqsj
	// kube-system		Running		kube-proxy-g4zql
	// kube-system		Running		kube-proxy-szxpr
	// kube-system		Running		kube-proxy-xgm5x
	// kube-system		Running		kube-scheduler-k8s-master-1
	// kube-system		Running		kube-scheduler-k8s-master-2
	// kube-system		Running		kube-scheduler-k8s-master-3
	// kube-system		Running		metrics-server-6bff78959f-nn6hw
}
