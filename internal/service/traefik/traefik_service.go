package traefik

import (
	"context"
	"fmt"
	"strings"

	"rapide/pkg/kubernetes"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

var Entrance = TraefikGroup{}

type TraefikGroup struct {
	TraefikService
}

type TraefikService struct{}

// GetHTTPRoutes 获取traefik命名空间中的HTTPRoutes，支持分页
func (s *TraefikService) GetHTTPRoutes(page, pageSize int) ([]unstructured.Unstructured, int, error) {
	// 检查Kubernetes客户端是否已初始化
	if kubernetes.Config == nil {
		return nil, 0, nil
	}

	// 创建dynamic client
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return nil, 0, err
	}

	// 定义HTTPRoutes的GroupVersionResource
	httpRouteGVR := schema.GroupVersionResource{
		Group:    "gateway.networking.k8s.io",
		Version:  "v1",
		Resource: "httproutes",
	}

	// 获取traefik命名空间中的HTTPRoutes
	httpRouteList, err := dynamicClient.Resource(httpRouteGVR).Namespace("traefik").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	total := len(httpRouteList.Items)

	// 处理分页
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= total {
		return []unstructured.Unstructured{}, total, nil
	}

	if end > total {
		end = total
	}

	return httpRouteList.Items[start:end], total, nil
}

// GetCRDByName 根据名称获取特定的Kubernetes CRD
func (s *TraefikService) GetCRDByName(name string) (*apiextensionsv1.CustomResourceDefinition, error) {
	// 检查API扩展客户端是否已初始化
	if kubernetes.APIClientset == nil {
		return nil, nil
	}

	// 获取特定名称的CRD
	crdClient := kubernetes.APIClientset.ApiextensionsV1().CustomResourceDefinitions()
	crd, err := crdClient.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return crd, nil
}

// GetCRResources 获取特定CRD下的所有自定义资源，支持分页
func (s *TraefikService) GetCRResources(group, version, kind string, page, pageSize int) ([]unstructured.Unstructured, int, error) {
	// 检查Kubernetes配置是否已初始化
	if kubernetes.Config == nil {
		return []unstructured.Unstructured{}, 0, nil
	}

	// 创建DynamicClient
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return nil, 0, err
	}

	// 构建GroupVersionResource
	resourceName := strings.ToLower(kind)
	// 处理特殊复数形式
	if strings.HasSuffix(resourceName, "class") {
		// 对于以"class"结尾的资源，复数形式为"classes"
		resourceName = strings.TrimSuffix(resourceName, "class") + "classes"
	} else {
		// 一般情况，添加"s"作为复数
		resourceName += "s"
	}
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resourceName,
	}

	// 获取所有资源
	resources, err := dynamicClient.Resource(gvr).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// 如果资源不存在（CRD未安装），返回空列表
		return []unstructured.Unstructured{}, 0, nil
	}

	total := len(resources.Items)

	// 处理分页
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= total {
		return []unstructured.Unstructured{}, total, nil
	}

	if end > total {
		end = total
	}

	return resources.Items[start:end], total, nil
}

// CreateGatewayAPIResource 创建Gateway API资源
func (s *TraefikService) CreateGatewayAPIResource(group, version, kind, namespace string, resource *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	// 检查Kubernetes配置是否已初始化
	if kubernetes.Config == nil {
		return nil, fmt.Errorf("kubernetes config is not initialized")
	}

	// 创建DynamicClient
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return nil, err
	}

	// 构建GroupVersionResource
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: strings.ToLower(kind) + "s",
	}

	// 创建资源
	var createdResource *unstructured.Unstructured
	if namespace == "" {
		// 集群范围资源
		createdResource, err = dynamicClient.Resource(gvr).Create(context.Background(), resource, metav1.CreateOptions{})
	} else {
		// 命名空间范围资源
		createdResource, err = dynamicClient.Resource(gvr).Namespace(namespace).Create(context.Background(), resource, metav1.CreateOptions{})
	}

	if err != nil {
		return nil, err
	}

	return createdResource, nil
}

// GetGatewayAPIResource 获取Gateway API资源
func (s *TraefikService) GetGatewayAPIResource(group, version, kind, namespace, name string) (*unstructured.Unstructured, error) {
	// 检查Kubernetes配置是否已初始化
	if kubernetes.Config == nil {
		return nil, fmt.Errorf("kubernetes config is not initialized")
	}

	// 创建DynamicClient
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return nil, err
	}

	// 构建GroupVersionResource
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: strings.ToLower(kind) + "s",
	}

	// 获取资源
	var resource *unstructured.Unstructured
	if namespace == "" {
		// 集群范围资源
		resource, err = dynamicClient.Resource(gvr).Get(context.Background(), name, metav1.GetOptions{})
	} else {
		// 命名空间范围资源
		resource, err = dynamicClient.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
	}

	if err != nil {
		return nil, err
	}

	return resource, nil
}

// UpdateGatewayAPIResource 更新Gateway API资源
func (s *TraefikService) UpdateGatewayAPIResource(group, version, kind, namespace, name string, resource *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	// 检查Kubernetes配置是否已初始化
	if kubernetes.Config == nil {
		return nil, fmt.Errorf("kubernetes config is not initialized")
	}

	// 创建DynamicClient
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return nil, err
	}

	// 构建GroupVersionResource
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: strings.ToLower(kind) + "s",
	}

	// 更新资源
	var updatedResource *unstructured.Unstructured
	if namespace == "" {
		// 集群范围资源
		updatedResource, err = dynamicClient.Resource(gvr).Update(context.Background(), resource, metav1.UpdateOptions{})
	} else {
		// 命名空间范围资源
		updatedResource, err = dynamicClient.Resource(gvr).Namespace(namespace).Update(context.Background(), resource, metav1.UpdateOptions{})
	}

	if err != nil {
		return nil, err
	}

	return updatedResource, nil
}

// DeleteGatewayAPIResource 删除Gateway API资源
func (s *TraefikService) DeleteGatewayAPIResource(group, version, kind, namespace, name string) error {
	// 检查Kubernetes配置是否已初始化
	if kubernetes.Config == nil {
		return fmt.Errorf("kubernetes config is not initialized")
	}

	// 创建DynamicClient
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return err
	}

	// 构建GroupVersionResource
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: strings.ToLower(kind) + "s",
	}

	// 删除资源
	if namespace == "" {
		// 集群范围资源
		err = dynamicClient.Resource(gvr).Delete(context.Background(), name, metav1.DeleteOptions{})
	} else {
		// 命名空间范围资源
		err = dynamicClient.Resource(gvr).Namespace(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	}

	if err != nil {
		return err
	}

	return nil
}
