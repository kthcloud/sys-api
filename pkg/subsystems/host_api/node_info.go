package host_api

import (
	"fmt"
	"sys-api/utils"
)

func (c *Client) GetNodeInfo() (*NodeInfo, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get host node info. details: %s", err)
	}

	response, err := utils.DoJsonGetRequest[NodeInfo](c.URL+"/nodeInfo", nil)
	if err != nil {
		return nil, makeError(err)
	}

	return response, nil
}
