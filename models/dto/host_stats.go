package dto

type HostK8sStats struct {
	PodCount int `json:"podCount"`
}

type HostStats struct {
	K8sStats HostK8sStats `json:"k8s"`
}
