package kubernetes

import (
	"log"

	"rapide/pkg/config"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client Kubernetes客户端
var Client *kubernetes.Clientset

// APIClientset API扩展客户端
var APIClientset *apiextensionsv1.Clientset

// Config Kubernetes配置
var Config *rest.Config

// InitClient 初始化Kubernetes客户端
func InitClient(kubeconfig string) error {
	var err error

	// 重置客户端和配置
	Client = nil
	APIClientset = nil
	Config = nil

	// 如果提供了kubeconfig文件路径，则使用该文件
	if kubeconfig != "" {
		Config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Printf("Failed to get Kubernetes config from file: %v", err)
			return err
		}
	} else {
		// 检查是否有环境变量配置
		apiServer := config.GetString("K8S_APISERVER", "")
		token := config.GetString("K8S_TOKEN", "")

		if apiServer != "" {
			// 使用环境变量配置
			Config = &rest.Config{
				Host:        apiServer,
				BearerToken: token,
			}
			log.Println("Kubernetes client using configuration from environment variables")
		} else {
			// 否则使用集群内配置
			Config, err = rest.InClusterConfig()
			if err != nil {
				// 在非Kubernetes集群环境中，rest.InClusterConfig()会返回错误
				// 这里我们不返回错误，而是直接返回，让调用者知道客户端初始化失败
				log.Printf("Failed to get in-cluster Kubernetes config: %v", err)
				return nil
			}
		}
	}

	// 创建核心客户端
	Client, err = kubernetes.NewForConfig(Config)
	if err != nil {
		log.Printf("Failed to create Kubernetes client: %v", err)
		return err
	}

	// 创建API扩展客户端
	APIClientset, err = apiextensionsv1.NewForConfig(Config)
	if err != nil {
		log.Printf("Failed to create API extensions client: %v", err)
		return err
	}

	return nil
}
