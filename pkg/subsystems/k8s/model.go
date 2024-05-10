package k8s

type Node struct {
	Name string `json:"name"`
	CPU  struct {
		Total int `json:"total"`
	} `json:"cpu"`
	RAM struct {
		Total int `json:"total"`
	} `json:"ram"`
}
