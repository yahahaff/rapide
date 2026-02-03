// Package initialize 初始化包
package initialize

import (
	"context"
	"fmt"

	"rapide/pkg/config"
	"rapide/pkg/kubernetes"
	"rapide/pkg/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SetupKubernetes 初始化Kubernetes客户端并测试连接
func SetupKubernetes() {
	// 从配置中获取kubeconfig文件路径
	kubeconfig := config.GetString("KUBECONFIG", "C:/Users/Administrator/Desktop/kubeconfig")

	// 初始化Kubernetes客户端
	err := kubernetes.InitClient(kubeconfig)
	if err != nil {
		// 记录错误，但不中断程序启动
		logger.Error("Failed to initialize Kubernetes client")
		return
	} else {
		logger.Info("Kubernetes client initialized successfully")
	}

	// 测试Kubernetes集群连接
	TestKubernetesConnection()
}

// TestKubernetesConnection 测试Kubernetes集群连接是否成功
func TestKubernetesConnection() {
	// 检查客户端是否已初始化
	if kubernetes.Client == nil {
		logger.Error("Kubernetes client is not initialized")
		return
	}

	// 使用客户端进行简单的API调用，测试连接并获取所有命名空间列表
	namespaces, err := kubernetes.Client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logger.Error("Failed to connect to Kubernetes cluster: " + err.Error())
		return
	}

	logger.Info("Successfully connected to Kubernetes cluster")
	logger.Info("Retrieved namespaces:")
	// 打印返回的命名空间列表
	for _, ns := range namespaces.Items {
		logger.Info("  - " + ns.Name + " (Created: " + ns.CreationTimestamp.Format("2006-01-02 15:04:05") + ")")
	}
	logger.Info("Total namespaces found: " + fmt.Sprint(len(namespaces.Items)))
}
