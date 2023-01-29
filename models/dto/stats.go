package dto

type K8sStats struct {
	PodCount int `json:"podCount"`
}

type Stats struct {
	K8sStats K8sStats `json:"k8s"`
}
