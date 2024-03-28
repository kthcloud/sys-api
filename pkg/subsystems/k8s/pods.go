package k8s

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetTotalPods() (int, error) {
	list, err := c.K8sClient.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(list.Items), nil
}
