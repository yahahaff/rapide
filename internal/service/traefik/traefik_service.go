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
	dynamic "k8s.io/client-go/dynamic"
)

var Entrance = TraefikGroup{}

type TraefikGroup struct {
	TraefikService
}

type TraefikService struct{}

// GetCRDList 获取Kubernetes CRD列表，支持分页
func (s *TraefikService) GetCRDList(page, pageSize int) ([]apiextensionsv1.CustomResourceDefinition, int, error) {
	// 检查API扩展客户端是否已初始化
	if kubernetes.APIClientset == nil {
		return nil, 0, nil
	}

	// 获取所有CRD
	crdClient := kubernetes.APIClientset.ApiextensionsV1().CustomResourceDefinitions()
	crdList, err := crdClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	total := len(crdList.Items)

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
		return []apiextensionsv1.CustomResourceDefinition{}, total, nil
	}

	if end > total {
		end = total
	}

	return crdList.Items[start:end], total, nil
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
		return nil, 0, fmt.Errorf("kubernetes config is not initialized")
	}

	// 创建DynamicClient
	dynamicClient, err := dynamic.NewForConfig(kubernetes.Config)
	if err != nil {
		return nil, 0, err
	}

	// 构建GroupVersionResource
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: strings.ToLower(kind) + "s",
	}

	// 获取所有资源
	resources, err := dynamicClient.Resource(gvr).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
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
