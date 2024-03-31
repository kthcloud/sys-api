package body

type HostRegisterParams struct {
	// Name is the host name of the node
	Name string `json:"name" binding:"required,min=3"`
	// DisplayName is the human readable name of the node
	// This is optional, and is set to Name if not provided
	DisplayName string `json:"displayName" binding:"min=3"`
	IP          string `json:"ip" binding:"required"`
	// Port is the port the node is listening on for API requests
	Port int    `json:"port" binding:"required"`
	Zone string `json:"zone" binding:"required"`

	// Token is the discovery token validated against the config
	Token string `json:"token" binding:"required"`
}

type ClusterRegisterParams struct {
}
