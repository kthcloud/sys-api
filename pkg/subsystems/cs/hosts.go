package cs

import (
	"strings"
	"sys-api/pkg/imp/cloudstack"
)

func (c *Client) ListHosts() ([]Host, error) {
	csHosts, err := c.CsClient.Host.ListHosts(&cloudstack.ListHostsParams{})
	if err != nil {
		return nil, err
	}

	var hosts []Host
	for _, csHost := range csHosts.Hosts {
		if csHost.Type != "Routing" || csHost.State != "Up" {
			continue
		}

		displayName := csHost.Name
		if csHost.Hosttags != "" {
			hostTags := strings.Split(csHost.Hosttags, ",")
			for _, keyValue := range hostTags {
				keyValueSplit := strings.Split(keyValue, "=")

				if len(keyValueSplit) == 2 {
					key := keyValueSplit[0]
					value := keyValueSplit[1]

					if key == "displayName" {
						displayName = value
					}
				}
			}
		}

		hosts = append(hosts, Host{
			Name:        csHost.Name,
			DisplayName: displayName,
			IP:          csHost.Ipaddress,
			Port:        8081, // TODO: make this configurable
			Enabled:     csHost.Resourcestate == "Enabled",
			Zone:        csHost.Zonename,
		})
	}

	return hosts, nil
}
