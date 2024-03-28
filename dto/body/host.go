package body

type HostBase struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	// Zone is the name of the zone where the host is located.
	// This field might not yet be present in all responses, in which case ZoneID should be used.
	Zone string `json:"zone,omitempty"`
}

type HostInfo struct {
	HostBase
}
