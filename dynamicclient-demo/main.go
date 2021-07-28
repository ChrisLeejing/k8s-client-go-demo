package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// 需求: 查询指定namespace下的所有pod, 然后在控制台打印出来, 要求用dynamicClient实现.
// 优化: 后期可以使用自定义CRD进行优化.
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

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// dynamicClient唯一关联方法所需要的入参
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	// 返回类型: unstructured非结构化对象
	unstructuredList, err := dynamicClient.Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		panic(err)
	}

	// 实例化podList用于接收unstructured非结构化对象的转换.
	podList := &corev1.PodList{}

	// unstructuredList -> podList
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredList.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}

	// 打印表头
	fmt.Printf("namespace\t\tstatus\t\tname\n")
	for _, item := range podList.Items {
		fmt.Printf("%v\t\t%v\t\t%v\n", item.Namespace, item.Status.Phase, item.Name)
	}
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