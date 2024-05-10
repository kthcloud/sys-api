package k8s

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetNodes() ([]Node, error) {
	nodes, err := c.K8sClient.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var ret []Node
	for _, node := range nodes.Items {
		ret = append(ret, Node{
			Name: node.Name,
			CPU: struct {
				Total int `json:"total"`
			}{
				Total: int(node.Status.Capacity.Cpu().MilliValue()) / 1000,
			},
			RAM: struct {
				Total int `json:"total"`
			}{
				Total: int(float64(node.Status.Capacity.Memory().Value() / 1024 / 1024 / 1024)),
			},
		})
	}

	return ret, nil
}
