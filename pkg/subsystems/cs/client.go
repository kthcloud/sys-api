package cs

import "sys-api/pkg/imp/cloudstack"

type Client struct {
	CsClient *cloudstack.CloudStackClient
}

type ClientConfig struct {
	URL    string
	ApiKey string
	Secret string
}

// NewClient creates a new CloudStack client.
func NewClient(config ClientConfig) *Client {
	csClient := cloudstack.NewAsyncClient(
		config.URL,
		config.ApiKey,
		config.Secret,
		true,
	)

	return &Client{
		CsClient: csClient,
	}
}
