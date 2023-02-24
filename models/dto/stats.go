package dto

type K8sStats struct {
	PodCount int `json:"podCount" bson:"podCount"`
}

type Stats struct {
	K8sStats K8sStats `json:"k8s" bson:"k8s"`
}
