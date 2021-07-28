package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// 需求: 从kubernetes查询所有的Group, Version, Resource信息, 在控制台打印出来.
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

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}
	// 获取分组和所有资源信息
	apiGroups, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}
	fmt.Printf("apiGroups: %v\n\n", apiGroups)

	for _, apiResourceList := range apiResourceLists {
		// GroupVersion是个字符串，例如"apps/v1"
		groupVersionStr := apiResourceList.GroupVersion
		// groupVersion字符串 -> groupVersion结构体
		groupVersionStruct, err := schema.ParseGroupVersion(groupVersionStr)
		if err != nil {
			panic(err)
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------------")
		fmt.Printf("groupVersionStr: %v, groupVersionStruct: %#v\n", groupVersionStr, groupVersionStruct)

		apiResources := apiResourceList.APIResources
		for _, apiResource := range apiResources {
			fmt.Printf("apiResource Name: %v, Shortname: %v, Kind: %v\n", apiResource.Name, apiResource.ShortNames, apiResource.Kind)
		}
	}
}
// apiGroups: [&APIGroup{Name:,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:v1,Version:v1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:apiregistration.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:apiregistration.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:apiregistration.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:apiregistration.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:extensions,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:extensions/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:extensions/v1beta1,Version:v1beta1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:apps,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:apps/v1,Version:v1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:apps/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:events.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:events.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:events.k8s.io/v1beta1,Version:v1beta1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:authentication.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:authentication.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:authentication.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:authentication.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:authorization.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:authorization.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:authorization.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:authorization.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:autoscaling,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:autoscaling/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:autoscaling/v2beta1,Version:v2beta1,},GroupVersionForDiscovery{GroupVersion:autoscaling/v2beta2,Version:v2beta2,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:autoscaling/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:batch,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:batch/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:batch/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:batch/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:certificates.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:certificates.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:certificates.k8s.io/v1beta1,Version:v1beta1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:networking.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:networking.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:networking.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:networking.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:policy,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:policy/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:policy/v1beta1,Version:v1beta1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:rbac.authorization.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:rbac.authorization.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:rbac.authorization.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:rbac.authorization.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:storage.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:storage.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:storage.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:storage.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:admissionregistration.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:admissionregistration.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:admissionregistration.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:admissionregistration.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:apiextensions.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:apiextensions.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:apiextensions.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:apiextensions.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:scheduling.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:scheduling.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:scheduling.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:scheduling.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:coordination.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:coordination.k8s.io/v1,Version:v1,},GroupVersionForDiscovery{GroupVersion:coordination.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:coordination.k8s.io/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:node.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:node.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:node.k8s.io/v1beta1,Version:v1beta1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:discovery.k8s.io,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:discovery.k8s.io/v1beta1,Version:v1beta1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:discovery.k8s.io/v1beta1,Version:v1beta1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},} &APIGroup{Name:crd.projectcalico.org,Versions:[]GroupVersionForDiscovery{GroupVersionForDiscovery{GroupVersion:crd.projectcalico.org/v1,Version:v1,},},PreferredVersion:GroupVersionForDiscovery{GroupVersion:crd.projectcalico.org/v1,Version:v1,},ServerAddressByClientCIDRs:[]ServerAddressByClientCIDR{},}]
//
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: v1, groupVersionStruct: schema.GroupVersion{Group:"", Version:"v1"}
// apiResource Name: bindings, Shortname: [], Kind: Binding
// apiResource Name: componentstatuses, Shortname: [cs], Kind: ComponentStatus
// apiResource Name: configmaps, Shortname: [cm], Kind: ConfigMap
// apiResource Name: endpoints, Shortname: [ep], Kind: Endpoints
// apiResource Name: events, Shortname: [ev], Kind: Event
// apiResource Name: limitranges, Shortname: [limits], Kind: LimitRange
// apiResource Name: namespaces, Shortname: [ns], Kind: Namespace
// apiResource Name: namespaces/finalize, Shortname: [], Kind: Namespace
// apiResource Name: namespaces/status, Shortname: [], Kind: Namespace
// apiResource Name: nodes, Shortname: [no], Kind: Node
// apiResource Name: nodes/proxy, Shortname: [], Kind: NodeProxyOptions
// apiResource Name: nodes/status, Shortname: [], Kind: Node
// apiResource Name: persistentvolumeclaims, Shortname: [pvc], Kind: PersistentVolumeClaim
// apiResource Name: persistentvolumeclaims/status, Shortname: [], Kind: PersistentVolumeClaim
// apiResource Name: persistentvolumes, Shortname: [pv], Kind: PersistentVolume
// apiResource Name: persistentvolumes/status, Shortname: [], Kind: PersistentVolume
// apiResource Name: pods, Shortname: [po], Kind: Pod
// apiResource Name: pods/attach, Shortname: [], Kind: PodAttachOptions
// apiResource Name: pods/binding, Shortname: [], Kind: Binding
// apiResource Name: pods/eviction, Shortname: [], Kind: Eviction
// apiResource Name: pods/exec, Shortname: [], Kind: PodExecOptions
// apiResource Name: pods/log, Shortname: [], Kind: Pod
// apiResource Name: pods/portforward, Shortname: [], Kind: PodPortForwardOptions
// apiResource Name: pods/proxy, Shortname: [], Kind: PodProxyOptions
// apiResource Name: pods/status, Shortname: [], Kind: Pod
// apiResource Name: podtemplates, Shortname: [], Kind: PodTemplate
// apiResource Name: replicationcontrollers, Shortname: [rc], Kind: ReplicationController
// apiResource Name: replicationcontrollers/scale, Shortname: [], Kind: Scale
// apiResource Name: replicationcontrollers/status, Shortname: [], Kind: ReplicationController
// apiResource Name: resourcequotas, Shortname: [quota], Kind: ResourceQuota
// apiResource Name: resourcequotas/status, Shortname: [], Kind: ResourceQuota
// apiResource Name: secrets, Shortname: [], Kind: Secret
// apiResource Name: serviceaccounts, Shortname: [sa], Kind: ServiceAccount
// apiResource Name: services, Shortname: [svc], Kind: Service
// apiResource Name: services/proxy, Shortname: [], Kind: ServiceProxyOptions
// apiResource Name: services/status, Shortname: [], Kind: Service
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: apiregistration.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"apiregistration.k8s.io", Version:"v1"}
// apiResource Name: apiservices, Shortname: [], Kind: APIService
// apiResource Name: apiservices/status, Shortname: [], Kind: APIService
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: apiregistration.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"apiregistration.k8s.io", Version:"v1beta1"}
// apiResource Name: apiservices, Shortname: [], Kind: APIService
// apiResource Name: apiservices/status, Shortname: [], Kind: APIService
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: extensions/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"extensions", Version:"v1beta1"}
// apiResource Name: ingresses, Shortname: [ing], Kind: Ingress
// apiResource Name: ingresses/status, Shortname: [], Kind: Ingress
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: apps/v1, groupVersionStruct: schema.GroupVersion{Group:"apps", Version:"v1"}
// apiResource Name: controllerrevisions, Shortname: [], Kind: ControllerRevision
// apiResource Name: daemonsets, Shortname: [ds], Kind: DaemonSet
// apiResource Name: daemonsets/status, Shortname: [], Kind: DaemonSet
// apiResource Name: deployments, Shortname: [deploy], Kind: Deployment
// apiResource Name: deployments/scale, Shortname: [], Kind: Scale
// apiResource Name: deployments/status, Shortname: [], Kind: Deployment
// apiResource Name: replicasets, Shortname: [rs], Kind: ReplicaSet
// apiResource Name: replicasets/scale, Shortname: [], Kind: Scale
// apiResource Name: replicasets/status, Shortname: [], Kind: ReplicaSet
// apiResource Name: statefulsets, Shortname: [sts], Kind: StatefulSet
// apiResource Name: statefulsets/scale, Shortname: [], Kind: Scale
// apiResource Name: statefulsets/status, Shortname: [], Kind: StatefulSet
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: events.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"events.k8s.io", Version:"v1beta1"}
// apiResource Name: events, Shortname: [ev], Kind: Event
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: authentication.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"authentication.k8s.io", Version:"v1"}
// apiResource Name: tokenreviews, Shortname: [], Kind: TokenReview
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: authentication.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"authentication.k8s.io", Version:"v1beta1"}
// apiResource Name: tokenreviews, Shortname: [], Kind: TokenReview
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: authorization.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"authorization.k8s.io", Version:"v1"}
// apiResource Name: localsubjectaccessreviews, Shortname: [], Kind: LocalSubjectAccessReview
// apiResource Name: selfsubjectaccessreviews, Shortname: [], Kind: SelfSubjectAccessReview
// apiResource Name: selfsubjectrulesreviews, Shortname: [], Kind: SelfSubjectRulesReview
// apiResource Name: subjectaccessreviews, Shortname: [], Kind: SubjectAccessReview
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: authorization.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"authorization.k8s.io", Version:"v1beta1"}
// apiResource Name: localsubjectaccessreviews, Shortname: [], Kind: LocalSubjectAccessReview
// apiResource Name: selfsubjectaccessreviews, Shortname: [], Kind: SelfSubjectAccessReview
// apiResource Name: selfsubjectrulesreviews, Shortname: [], Kind: SelfSubjectRulesReview
// apiResource Name: subjectaccessreviews, Shortname: [], Kind: SubjectAccessReview
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: autoscaling/v1, groupVersionStruct: schema.GroupVersion{Group:"autoscaling", Version:"v1"}
// apiResource Name: horizontalpodautoscalers, Shortname: [hpa], Kind: HorizontalPodAutoscaler
// apiResource Name: horizontalpodautoscalers/status, Shortname: [], Kind: HorizontalPodAutoscaler
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: autoscaling/v2beta1, groupVersionStruct: schema.GroupVersion{Group:"autoscaling", Version:"v2beta1"}
// apiResource Name: horizontalpodautoscalers, Shortname: [hpa], Kind: HorizontalPodAutoscaler
// apiResource Name: horizontalpodautoscalers/status, Shortname: [], Kind: HorizontalPodAutoscaler
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: autoscaling/v2beta2, groupVersionStruct: schema.GroupVersion{Group:"autoscaling", Version:"v2beta2"}
// apiResource Name: horizontalpodautoscalers, Shortname: [hpa], Kind: HorizontalPodAutoscaler
// apiResource Name: horizontalpodautoscalers/status, Shortname: [], Kind: HorizontalPodAutoscaler
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: batch/v1, groupVersionStruct: schema.GroupVersion{Group:"batch", Version:"v1"}
// apiResource Name: jobs, Shortname: [], Kind: Job
// apiResource Name: jobs/status, Shortname: [], Kind: Job
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: batch/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"batch", Version:"v1beta1"}
// apiResource Name: cronjobs, Shortname: [cj], Kind: CronJob
// apiResource Name: cronjobs/status, Shortname: [], Kind: CronJob
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: certificates.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"certificates.k8s.io", Version:"v1beta1"}
// apiResource Name: certificatesigningrequests, Shortname: [csr], Kind: CertificateSigningRequest
// apiResource Name: certificatesigningrequests/approval, Shortname: [], Kind: CertificateSigningRequest
// apiResource Name: certificatesigningrequests/status, Shortname: [], Kind: CertificateSigningRequest
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: networking.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"networking.k8s.io", Version:"v1"}
// apiResource Name: networkpolicies, Shortname: [netpol], Kind: NetworkPolicy
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: networking.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"networking.k8s.io", Version:"v1beta1"}
// apiResource Name: ingressclasses, Shortname: [], Kind: IngressClass
// apiResource Name: ingresses, Shortname: [ing], Kind: Ingress
// apiResource Name: ingresses/status, Shortname: [], Kind: Ingress
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: policy/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"policy", Version:"v1beta1"}
// apiResource Name: poddisruptionbudgets, Shortname: [pdb], Kind: PodDisruptionBudget
// apiResource Name: poddisruptionbudgets/status, Shortname: [], Kind: PodDisruptionBudget
// apiResource Name: podsecuritypolicies, Shortname: [psp], Kind: PodSecurityPolicy
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: rbac.authorization.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"rbac.authorization.k8s.io", Version:"v1"}
// apiResource Name: clusterrolebindings, Shortname: [], Kind: ClusterRoleBinding
// apiResource Name: clusterroles, Shortname: [], Kind: ClusterRole
// apiResource Name: rolebindings, Shortname: [], Kind: RoleBinding
// apiResource Name: roles, Shortname: [], Kind: Role
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: rbac.authorization.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"rbac.authorization.k8s.io", Version:"v1beta1"}
// apiResource Name: clusterrolebindings, Shortname: [], Kind: ClusterRoleBinding
// apiResource Name: clusterroles, Shortname: [], Kind: ClusterRole
// apiResource Name: rolebindings, Shortname: [], Kind: RoleBinding
// apiResource Name: roles, Shortname: [], Kind: Role
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: storage.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"storage.k8s.io", Version:"v1"}
// apiResource Name: csidrivers, Shortname: [], Kind: CSIDriver
// apiResource Name: csinodes, Shortname: [], Kind: CSINode
// apiResource Name: storageclasses, Shortname: [sc], Kind: StorageClass
// apiResource Name: volumeattachments, Shortname: [], Kind: VolumeAttachment
// apiResource Name: volumeattachments/status, Shortname: [], Kind: VolumeAttachment
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: storage.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"storage.k8s.io", Version:"v1beta1"}
// apiResource Name: csidrivers, Shortname: [], Kind: CSIDriver
// apiResource Name: csinodes, Shortname: [], Kind: CSINode
// apiResource Name: storageclasses, Shortname: [sc], Kind: StorageClass
// apiResource Name: volumeattachments, Shortname: [], Kind: VolumeAttachment
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: admissionregistration.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"admissionregistration.k8s.io", Version:"v1"}
// apiResource Name: mutatingwebhookconfigurations, Shortname: [], Kind: MutatingWebhookConfiguration
// apiResource Name: validatingwebhookconfigurations, Shortname: [], Kind: ValidatingWebhookConfiguration
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: admissionregistration.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"admissionregistration.k8s.io", Version:"v1beta1"}
// apiResource Name: mutatingwebhookconfigurations, Shortname: [], Kind: MutatingWebhookConfiguration
// apiResource Name: validatingwebhookconfigurations, Shortname: [], Kind: ValidatingWebhookConfiguration
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: apiextensions.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"apiextensions.k8s.io", Version:"v1"}
// apiResource Name: customresourcedefinitions, Shortname: [crd crds], Kind: CustomResourceDefinition
// apiResource Name: customresourcedefinitions/status, Shortname: [], Kind: CustomResourceDefinition
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: apiextensions.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"apiextensions.k8s.io", Version:"v1beta1"}
// apiResource Name: customresourcedefinitions, Shortname: [crd crds], Kind: CustomResourceDefinition
// apiResource Name: customresourcedefinitions/status, Shortname: [], Kind: CustomResourceDefinition
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: scheduling.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"scheduling.k8s.io", Version:"v1"}
// apiResource Name: priorityclasses, Shortname: [pc], Kind: PriorityClass
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: scheduling.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"scheduling.k8s.io", Version:"v1beta1"}
// apiResource Name: priorityclasses, Shortname: [pc], Kind: PriorityClass
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: coordination.k8s.io/v1, groupVersionStruct: schema.GroupVersion{Group:"coordination.k8s.io", Version:"v1"}
// apiResource Name: leases, Shortname: [], Kind: Lease
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: coordination.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"coordination.k8s.io", Version:"v1beta1"}
// apiResource Name: leases, Shortname: [], Kind: Lease
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: node.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"node.k8s.io", Version:"v1beta1"}
// apiResource Name: runtimeclasses, Shortname: [], Kind: RuntimeClass
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: discovery.k8s.io/v1beta1, groupVersionStruct: schema.GroupVersion{Group:"discovery.k8s.io", Version:"v1beta1"}
// apiResource Name: endpointslices, Shortname: [], Kind: EndpointSlice
// --------------------------------------------------------------------------------------------------------------
// groupVersionStr: crd.projectcalico.org/v1, groupVersionStruct: schema.GroupVersion{Group:"crd.projectcalico.org", Version:"v1"}
// apiResource Name: ipamblocks, Shortname: [], Kind: IPAMBlock
// apiResource Name: networksets, Shortname: [], Kind: NetworkSet
// apiResource Name: blockaffinities, Shortname: [], Kind: BlockAffinity
// apiResource Name: bgppeers, Shortname: [], Kind: BGPPeer
// apiResource Name: globalnetworksets, Shortname: [], Kind: GlobalNetworkSet
// apiResource Name: ipamhandles, Shortname: [], Kind: IPAMHandle
// apiResource Name: ipamconfigs, Shortname: [], Kind: IPAMConfig
// apiResource Name: ippools, Shortname: [], Kind: IPPool
// apiResource Name: hostendpoints, Shortname: [], Kind: HostEndpoint
// apiResource Name: networkpolicies, Shortname: [], Kind: NetworkPolicy
// apiResource Name: clusterinformations, Shortname: [], Kind: ClusterInformation
// apiResource Name: felixconfigurations, Shortname: [], Kind: FelixConfiguration
// apiResource Name: bgpconfigurations, Shortname: [], Kind: BGPConfiguration
// apiResource Name: globalnetworkpolicies, Shortname: [], Kind: GlobalNetworkPolicy