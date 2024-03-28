package k8s

import "k8s.io/client-go/kubernetes"

type Client struct {
	K8sClient *kubernetes.Clientset
}

func NewClient(k8sClient *kubernetes.Clientset) *Client {
	return &Client{
		K8sClient: k8sClient,
	}
}
