package poll

import (
	"sys-api/pkg/config"
)

// HostFetchWorker fetches hosts from static config
// This includes querying the cloudstack API for the latest host information
// The newer way is for host to register themselves with the Register v2 API
func HostFetchWorker() error {
	return config.SyncCloudStackHosts()
}
